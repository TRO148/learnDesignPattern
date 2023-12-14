package command

type Group struct {
	command []Interface
}

func NewGroup(commands ...Interface) *Group {
	return &Group{command: commands}
}

func (G *Group) Execute() {
	for _, v := range G.command {
		v.Execute()
	}
}

func (G *Group) Undo() {
	for _, v := range G.command {
		v.Undo()
	}
}
