package main

import "container/list"

type Item string

type DinerMenu struct {
	menuItems []Item
	position  int
}

func (dm *DinerMenu) HasNext() bool {
	return dm.position < len(dm.menuItems) && dm.menuItems[dm.position] != ""
}

func (dm *DinerMenu) GetNext() (result Item) {
	result = dm.menuItems[dm.position]
	dm.position++
	return
}

type PancakeHouseMenu struct {
	menuItems list.List
	position  int
}

func (phm *PancakeHouseMenu) HasNext() bool {
	return phm.position <= phm.menuItems.Len()
}

func (phm *PancakeHouseMenu) GetNext() (result Item) {
	point := phm.position
	for e := phm.menuItems.Front(); e != nil; e = e.Next() {
		point--
		if point == 0 {
			result = e.Value.(Item)
			break
		}
	}
	phm.position++
	return
}

type CafeMenu struct {
	menuItems map[string]Item
	position  int
}

func (cm *CafeMenu) HasNext() bool {
	return cm.position < len(cm.menuItems)
}

func (cm *CafeMenu) GetNext() (result Item) {
	point := cm.position
	for _, v := range cm.menuItems {
		if point == 0 {
			result = v
			break
		}
		point--
	}
	cm.position++
	return
}
