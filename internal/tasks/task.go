package tasks

type Task interface {
	Execute() error
	GetType() string
	GetPayload() map[string]interface{}
}
