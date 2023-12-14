package command

type Interface interface {
	Execute()
	Undo()
}
