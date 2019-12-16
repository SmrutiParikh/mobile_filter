package services

import (
	"container/list"
	"fmt"

	"github.com/mobile_filter_2/models"
)

var mobileList list.List

//AddToInventory is ...
func AddToInventory(model string, brand models.Brand, memory int) {
	fmt.Println("Adding Mobile to Inventory ...")

	mobileList.PushBack(models.Mobile{Model: model, Brand: brand, Memory: memory})
}

//RemoveFromInventory is ...
func RemoveFromInventory(filterType models.FilterType, keyword string) {
	fmt.Println("Removing Mobile from Inventory ...")

	var filteredValue = Filter(filterType, keyword, mobileList)[0]
	for element := mobileList.Front(); element != nil; element = element.Next() {
		current := element.Value.(models.Mobile)
		if current.Equals(filteredValue) {
			mobileList.Remove(element)
		}
	}
}

//InventoryStatus is ...
func InventoryStatus() list.List {
	fmt.Println("Status : ")
	for element := mobileList.Front(); element != nil; element = element.Next() {
		current := element.Value.(models.Mobile)
		current.Print()
	}
	return mobileList
}
