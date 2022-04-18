package external_logger

type ExternalLogger interface {
	LogError(msg string) error
	LogWarning(msg string) error
}

type Logger struct{}

func (l Logger) LogError(msg string) error {
	return nil
}

func (l Logger) LogWarning(msg string) error {
	return nil
}
