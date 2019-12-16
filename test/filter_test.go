package test

import (
	"container/list"
	"fmt"
	"testing"

	"github.com/mobile_filter_2/models"
	"github.com/mobile_filter_2/services"
)

var mobileList list.List

func TestFilters(t *testing.T) {
	testAddToInventory(t)
	testFilterModel(t)
	testFilterBrand(t)
	testFilterMemory(t)
	testRemoveFromInventory(t)
}

func testAddToInventory(t *testing.T) {
	fmt.Println("Init Test......")
	services.AddToInventory("realme1", models.REALME, 6)
	services.AddToInventory("iphone10", models.APPLE, 8)
	services.AddToInventory("vivov5", models.VIVO, 4)
	mobileList = services.InventoryStatus()
	if mobileList.Len() != 3 {
		t.Error("Add to Inventory does not match")
	}
}

func testRemoveFromInventory(t *testing.T) {
	services.RemoveFromInventory(models.MODEL, "realme1")
	mobileList = services.InventoryStatus()
	if mobileList.Len() != 2 {
		t.Error("Remove from Inventory does not match")
	}
}

func testFilterModel(t *testing.T) {
	result := services.Filter(models.MODEL, "realme1", mobileList)
	if !validateResult(result, "realme1", models.REALME, 6) {
		t.Error("Mobile Model does not match")
	}
}

func testFilterBrand(t *testing.T) {
	result := services.Filter(models.BRAND, "APPLE", mobileList)
	if !validateResult(result, "iphone10", models.APPLE, 8) {
		t.Error("Mobile Brand does not match")
	}

}

func testFilterMemory(t *testing.T) {
	result := services.Filter(models.MEMORY, "4", mobileList)
	if !validateResult(result, "vivov5", models.VIVO, 4) {
		t.Error("Mobile Memory does not match")
	}
}

func validateResult(result []models.Mobile, expectedModel string, expectedBrand models.Brand, expectedMemory int) bool {
	return len(result) == 1 && fmt.Sprintf("%T", result[0]) == "models.Mobile" &&
		result[0].Model == expectedModel && result[0].Brand == expectedBrand && result[0].Memory == expectedMemory
}
