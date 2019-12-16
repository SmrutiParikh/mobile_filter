package main

import (
	"github.com/mobile_filter_2/models"
	"github.com/mobile_filter_2/services"
)

func main() {
	services.AddToInventory("realme1", models.REALME, 6)
	services.AddToInventory("iphone10", models.APPLE, 8)
	services.AddToInventory("vivov5", models.VIVO, 4)

	mobileList := services.InventoryStatus()

	services.Filter(models.MODEL, "iphone10", mobileList)
	services.Filter(models.BRAND, "REALME", mobileList)
	services.Filter(models.MEMORY, "4", mobileList)

	services.RemoveFromInventory(models.MODEL, "vivov5")

	mobileList = services.InventoryStatus()
}
