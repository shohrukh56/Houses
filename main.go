package main

import (
	"sort"
)

type house struct {
	id                 int64
	name               string
	price              int64
	distanceFromCenter int64
	region             string
}

func sortBy(houses []house, less func(a, b house) bool) []house {
	result := make([]house, len(houses))
	copy(result, houses)
	sort.Slice(result, func(i, j int) bool {
		return less(result[i], result[j])
	})
	return result
}
func sortByAscPrice(houses []house) []house {
	return sortBy(houses, func(a, b house) bool {
		return a.price < b.price
	})
}
func sortByDescPrice(houses []house) []house {
	return sortBy(houses, func(a, b house) bool {
		return a.price > b.price
	})
}
func sortByNearestFromCenter(houses []house) []house {
	return sortBy(houses, func(a, b house) bool {
		return a.distanceFromCenter < b.distanceFromCenter
	})
}
func sortByFurthestFromCenter(houses []house) []house {
	return sortBy(houses, func(a, b house) bool {
		return a.distanceFromCenter > b.distanceFromCenter
	})
}
func SearchBy(houses []house, predicate func(houses house) bool) []house {
	result := make([]house, 0)
	for _, house := range houses {
		if predicate(house) {
			result = append(result, house)
		}
	}
	return result
}
func searchByMaxPrice(houses []house, maxPrice int64) []house {
	return SearchBy(houses, func(a house) bool {
		return a.price <= maxPrice
	})
}

func searchByWantedPrice(houses []house, price int64) []house {
	return SearchBy(houses, func(a house) bool {
		return a.price == price
	})
}
func searchByMaxAndMinPrice(houses []house, minPrice, maxPrice int64) []house {
	return SearchBy(houses, func(a house) bool {
		return a.price >= minPrice && a.price <= maxPrice
	})
}
func searchByRegion(houses []house, region string) []house {
	return SearchBy(houses, func(a house) bool {
		return a.region == region
	})
}
func searchByRegions(houses []house, regions []string) []house {
	result := make([]house, 0)
	for _, house := range houses{
		for _, region := range regions {
			if house.region == region {
				result = append(result, house)
			}
		}
	}
	return result
}

func main() {
}


