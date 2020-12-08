package common

type Consumer interface {
	Subscribe(topic string) error
	Subscribes(topics []string) error
	Read(cancel <-chan struct{}) (key []byte, value []byte, err error)
	Close()
}
