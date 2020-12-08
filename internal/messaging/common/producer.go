package common

type Producer interface {
	Publish(payload []byte) error
	Close()
}
