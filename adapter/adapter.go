package main

type Adapter interface {
	Duck
	Turkey
}

type flyAnimal interface {
	fly()
}

type realAdapter struct {
	animal flyAnimal
}

func NewAdapter(animal flyAnimal) Adapter {
	return &realAdapter{animal: animal}
}

func (ra realAdapter) quack() {
	duck, ok := ra.animal.(Duck)
	if !ok {
		ra.animal.(Turkey).gobble()
		return
	}
	duck.quack()
}

func (ra *realAdapter) gobble() {
	turkey, ok := ra.animal.(Turkey)
	if !ok {
		ra.animal.(Duck).quack()
		return
	}
	turkey.gobble()
}

func (ra *realAdapter) fly() {
	ra.animal.fly()
}
