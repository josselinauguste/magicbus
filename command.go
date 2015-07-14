package magicbus

type Command interface {
	Execute() error
}
