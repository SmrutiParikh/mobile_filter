package services

import (
	"container/list"
	"fmt"
	"strconv"

	"github.com/mobile_filter_2/models"
)

// FilterModel is ...
func FilterModel(model string, mobileList list.List) []models.Mobile {
	var filteredValues []models.Mobile
	for element := mobileList.Front(); element != nil; element = element.Next() {
		current := element.Value.(models.Mobile)
		if current.Model == model {
			filteredValues = append(filteredValues, current)
		}
	}
	return filteredValues
}

// FilterMemory is ...
func FilterMemory(memory int, mobileList list.List) []models.Mobile {
	var filteredValues []models.Mobile
	for element := mobileList.Front(); element != nil; element = element.Next() {
		current := element.Value.(models.Mobile)
		if current.Memory == memory {
			filteredValues = append(filteredValues, current)
		}
	}
	return filteredValues
}

// FilterBrand is ...
func FilterBrand(brand models.Brand, mobileList list.List) []models.Mobile {
	var filteredValues []models.Mobile
	for element := mobileList.Front(); element != nil; element = element.Next() {
		current := element.Value.(models.Mobile)
		if current.Brand == brand {
			filteredValues = append(filteredValues, current)
		}
	}
	return filteredValues
}

// Filter is ...
func Filter(filterType models.FilterType, keyword string, mobileList list.List) []models.Mobile {
	var filteredValues []models.Mobile
	if filterType == models.MODEL {
		filteredValues = FilterModel(keyword, mobileList)
	} else if filterType == models.BRAND {
		filteredValues = FilterBrand(models.Parse(keyword), mobileList)
	} else if filterType == models.MEMORY {
		memory, error := strconv.Atoi(keyword)
		if error != nil {
			fmt.Println("Error at Filter", error)
			return nil
		}
		filteredValues = FilterMemory(memory, mobileList)
	}
	fmt.Printf("Filter type : %s, keyword : %s\n", filterType, keyword)
	printFilteredValues(filteredValues)
	return filteredValues
}

func printFilteredValues(filteredValues []models.Mobile) {
	for _, mobile := range filteredValues {
		mobile.Print()
	}
}
