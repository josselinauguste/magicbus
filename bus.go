package magicbus

type Bus interface {
	Send(command Command) error
}

type SynchronousBus struct {
}

func NewSynchronousBus() *SynchronousBus {
	return &SynchronousBus{}
}

func (bus SynchronousBus) Send(command Command) error {
	return command.Execute()
}
