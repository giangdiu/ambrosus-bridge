package server

import (
	"bytes"
	"context"
	"fmt"
	"sync"

	"github.com/ambrosus/ambrosus-bridge/relay/pkg/mpc/networking/common"
	"github.com/ambrosus/ambrosus-bridge/relay/pkg/mpc/tss_wrap"
	ec "github.com/ethereum/go-ethereum/common"
	"github.com/rs/zerolog"
)

type Server struct {
	sync.Mutex
	Tss *tss_wrap.Mpc

	operation []byte
	fullMsg   []byte

	connections      map[string]*common.Conn
	connChangeNotify context.CancelFunc

	results map[string][]byte

	logger *zerolog.Logger
}

// NewServer create and start new server
func NewServer(tss *tss_wrap.Mpc, logger *zerolog.Logger) *Server {
	s := &Server{
		Tss:    tss,
		logger: logger,
	}
	return s
}

// todo if threshold < partyLen, do we need to provide current party or use full party? client doesn't know about current part of party

func (s *Server) Sign(ctx context.Context, partyIDs []string, msg []byte) ([]byte, error) {
	s.logger.Info().Msg("Start sign operation")

	if err := s.startOperation(msg, partyIDs); err != nil {
		return nil, err
	}

	signature, err := s.doOperation(ctx,
		func(ctx context.Context, inCh <-chan []byte, outCh chan<- *tss_wrap.Message) ([]byte, error) {
			return s.Tss.Sign(ctx, partyIDs, inCh, outCh, msg)
		},
	)

	return signature, err
}

func (s *Server) Keygen(ctx context.Context, partyIDs []string) error {
	s.logger.Info().Msg("Start keygen operation")

	if err := s.startOperation(common.KeygenOperation, partyIDs); err != nil {
		return err
	}

	_, err := s.doOperation(ctx,
		func(ctx context.Context, inCh <-chan []byte, outCh chan<- *tss_wrap.Message) ([]byte, error) {
			err := s.Tss.Keygen(ctx, partyIDs, inCh, outCh)
			if err != nil {
				return nil, err
			}
			addr, err := s.Tss.GetAddress()
			return addr.Bytes(), err
		},
	)

	return err
}

func (s *Server) Reshare(ctx context.Context, partyIDsOld, partyIDsNew []string, thresholdNew int) error {
	s.logger.Info().Msg("Start reshare operation")

	if err := s.startOperation(common.ReshareOperation, append(partyIDsOld, partyIDsNew...)); err != nil {
		return err
	}

	_, err := s.doOperation(ctx,
		func(ctx context.Context, inCh <-chan []byte, outCh chan<- *tss_wrap.Message) ([]byte, error) {
			err := s.Tss.Reshare(ctx, partyIDsOld, partyIDsNew, thresholdNew, inCh, outCh)
			if err != nil {
				return nil, err
			}
			addr, err := s.Tss.GetAddress()
			return addr.Bytes(), err
		},
	)

	return err
}

func (s *Server) GetFullMsg() ([]byte, error) {
	// just to implement MpcSigner interface
	panic("can be called only on client")
}

func (s *Server) SetFullMsg(fullMsg []byte) {
	s.fullMsg = fullMsg
}

func (s *Server) GetTssAddress() (ec.Address, error) {
	return s.Tss.GetAddress()
}

func (s *Server) doOperation(
	ctx context.Context,
	tssOperation common.OperationFunc,
) ([]byte, error) {
	defer s.stopOperation()

	if err := s.waitForConnections(ctx); err != nil {
		return nil, fmt.Errorf("wait for connections: %w", err)
	}

	result, err := s.doOperation_(ctx, tssOperation)

	if err != nil {
		s.logger.Error().Err(err).Msg("Operation error")
		s.disconnectAll(fmt.Errorf("server error: %w", err))
		return nil, err
	}

	s.logger.Info().Msg("Operation finished successfully")
	s.disconnectAll(nil)
	return result, nil
}

func (s *Server) doOperation_(
	ctx context.Context,
	tssOperation common.OperationFunc,
) (ownResult []byte, err error) {

	inCh := make(chan []byte, 10)
	outCh := make(chan *tss_wrap.Message, 10)
	errCh := make(chan common.OpError, 3)

	tssWaiter := make(chan interface{}, 1)
	// todo close channels?

	go func() {
		ownResult, err = tssOperation(ctx, inCh, outCh)
		errCh <- common.OpError{"tss", err}
		tssWaiter <- nil
	}()
	go func() { errCh <- common.OpError{"res", s.receiver(outCh)} }()
	go func() { errCh <- common.OpError{"tra", s.transmitter(outCh, inCh)} }()

	for {
		select {
		case <-ctx.Done():
			return nil, ctx.Err()

		case err := <-errCh:
			if err.Err != nil {
				return nil, fmt.Errorf("%s error: %w", err.Type, err.Err)
			}

			// if err is nil, it means that some goroutine successfully finished

			if err.Type == "tss" {
				s.logger.Info().Msg("Tss operation finished successfully")
			}

			if err.Type == "res" {
				// receiver returns nil when all clients send results
				s.logger.Info().Msg("All results received")

				// wait for own result
				<-tssWaiter

				if err := checkResults(s.results, ownResult); err != nil {
					return nil, fmt.Errorf("check results: %w", err)
				}
				s.logger.Info().Msg("Results checked successfully")

				// close outCh so transmitter goroutine will finish (when all queued msgs will be sent)
				close(outCh)
			}
			if err.Type == "tra" {
				// transmitter will return nil when s.operation.OutCh channel closed (when all client send results)
				// at this point we received all results and sends all queued msgs, so finish protocol
				s.logger.Info().Msg("Transmitter finished successfully")
				return ownResult, nil
			}

		}
	}
}

func (s *Server) startOperation(msg []byte, waitForIDs []string) error {
	s.Lock()
	defer s.Unlock()

	if s.operation != nil {
		return fmt.Errorf("operation already started")
	}

	s.operation = msg
	s.results = make(map[string][]byte)
	s.makeNamedConnections(waitForIDs)

	return nil
}

func (s *Server) stopOperation() {
	s.Lock()
	defer s.Unlock()
	s.operation = nil
}

func checkResults(results map[string][]byte, ownResult []byte) error {
	for clientID, v := range results {
		if !bytes.Equal(v, ownResult) {
			return fmt.Errorf("client %v send different result (%v != %v)", clientID, v, ownResult)
		}
	}
	return nil
}
