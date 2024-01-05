package main

type Waitress struct {
	dinerMenu        *DinerMenu
	pancakeHouseMenu *PancakeHouseMenu
	cafeMenu         *CafeMenu
}

func NewWaitress(dinerMenu *DinerMenu, pancakeHouseMenu *PancakeHouseMenu, cafeMenu *CafeMenu) *Waitress {
	return &Waitress{
		dinerMenu:        dinerMenu,
		pancakeHouseMenu: pancakeHouseMenu,
		cafeMenu:         cafeMenu,
	}
}

func (w *Waitress) PrintMenu() {
	dinerIterator := w.dinerMenu
	pancakeIterator := w.pancakeHouseMenu
	cafeIterator := w.cafeMenu
	println("MENU\n----\nBREAKFAST")

	w.printMenu(pancakeIterator)
	println("\nLUNCH")
	w.printMenu(dinerIterator)
	println("\nDINNER")
	w.printMenu(cafeIterator)
}

func (w *Waitress) printMenu(iterator Iterator) {
	for iterator.HasNext() {
		item := iterator.GetNext()
		println(item)
	}
}
