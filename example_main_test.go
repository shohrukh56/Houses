package main

import (
	"fmt"
)

var (
	houses []house = []house{
		{id: 1, name: "2-х ком квартира в центре города", price: 100, distanceFromCenter: 100, region: "Рудаки"},
		{id: 2, name: "3-х ком. квартира возле рынка Саховат", price: 80, distanceFromCenter: 200, region: "Фирдавси"},
		{id: 3, name: "Земельный участок 20 соток", price: 500, distanceFromCenter: 300, region: "Сино"},
		{id: 4, name: "Пентхаус в центре города", price: 500, distanceFromCenter: 50, region: "Исмоили Сомони"},
		{id: 5, name: "Дом возле школы", price: 90, distanceFromCenter: 250, region: "Сино"},
	}
)

func ExampleSortByAscPrice() {
	ascPrice := sortByAscPrice(houses)
	fmt.Println(ascPrice)
	//Output:[{2 3-х ком. квартира возле рынка Саховат 80 200 Фирдавси} {5 Дом возле школы 90 250 Сино} {1 2-х ком квартира в центре города 100 100 Рудаки} {3 Земельный участок 20 соток 500 300 Сино} {4 Пентхаус в центре города 500 50 Исмоили Сомони}]
}
func ExampleSortByDescPrice() {
	descPrice := sortByDescPrice(houses)
	fmt.Println(descPrice)
	//Output:[{3 Земельный участок 20 соток 500 300 Сино} {4 Пентхаус в центре города 500 50 Исмоили Сомони} {1 2-х ком квартира в центре города 100 100 Рудаки} {5 Дом возле школы 90 250 Сино} {2 3-х ком. квартира возле рынка Саховат 80 200 Фирдавси}]
}
func ExampleSortByNearestFromCenter() {
	nearFromCenter := sortByNearestFromCenter(houses)
	fmt.Println(nearFromCenter)
	//Output:[{4 Пентхаус в центре города 500 50 Исмоили Сомони} {1 2-х ком квартира в центре города 100 100 Рудаки} {2 3-х ком. квартира возле рынка Саховат 80 200 Фирдавси} {5 Дом возле школы 90 250 Сино} {3 Земельный участок 20 соток 500 300 Сино}]
}
func ExampleSortByFurthestFromCenter() {
	farFromCenter := sortByFurthestFromCenter(houses)
	fmt.Println(farFromCenter)
	//Output:[{3 Земельный участок 20 соток 500 300 Сино} {5 Дом возле школы 90 250 Сино} {2 3-х ком. квартира возле рынка Саховат 80 200 Фирдавси} {1 2-х ком квартира в центре города 100 100 Рудаки} {4 Пентхаус в центре города 500 50 Исмоили Сомони}]
}
func ExampleSearchByMaxPrice_NoResult() {
	result := searchByMaxPrice(houses, 20)
	fmt.Println(result)
	//Output:[]
}
func ExampleSearchByMaxPrice_OneResult() {
	result := searchByMaxPrice(houses, 80)
	fmt.Println(result)
	//Output:[{2 3-х ком. квартира возле рынка Саховат 80 200 Фирдавси}]
}
func ExampleSearchByMaxPrice_TwoOrMoreResults() {
	result := searchByMaxPrice(houses, 500)
	fmt.Println(result)
	//Output:[{1 2-х ком квартира в центре города 100 100 Рудаки} {2 3-х ком. квартира возле рынка Саховат 80 200 Фирдавси} {3 Земельный участок 20 соток 500 300 Сино} {4 Пентхаус в центре города 500 50 Исмоили Сомони} {5 Дом возле школы 90 250 Сино}]
}
func ExampleSearchByWantedPrice_NoResult() {
	result := searchByWantedPrice(houses, 0)
	fmt.Println(result)
	//Output:[]
}
func ExampleSearchByWantedPrice_OneResult() {
	result := searchByWantedPrice(houses, 80)
	fmt.Println(result)
	//Output:[{2 3-х ком. квартира возле рынка Саховат 80 200 Фирдавси}]
}
func ExampleSearchByWantedPrice_TwoOrMoreResults() {
	result := searchByWantedPrice(houses, 500)
	fmt.Println(result)
	//Output:[{3 Земельный участок 20 соток 500 300 Сино} {4 Пентхаус в центре города 500 50 Исмоили Сомони}]
}
func ExampleSearchByMaxAndMinPrice() {
	result := searchByMaxAndMinPrice(houses, 80, 150)
	fmt.Println(result)
	//Output:[{1 2-х ком квартира в центре города 100 100 Рудаки} {2 3-х ком. квартира возле рынка Саховат 80 200 Фирдавси} {5 Дом возле школы 90 250 Сино}]
}
func ExampleSearchByRegion_NoResult() {
	result := searchByRegion(houses, "Помир")
	fmt.Println(result)
	//Output:[]
}
func ExampleSearchByRegion_OneResult() {
	result := searchByRegion(houses, "Исмоили Сомони")
	fmt.Println(result)
	//Output:[{4 Пентхаус в центре города 500 50 Исмоили Сомони}]
}
func ExampleSearchByRegion_TwoOrMoreResult() {
	result := searchByRegion(houses, "Сино")
	fmt.Println(result)
	//Output:[{3 Земельный участок 20 соток 500 300 Сино} {5 Дом возле школы 90 250 Сино}]
}
func ExampleSearchByRegions_NoResult() {
	result := searchByRegions(houses, []string{"Исмоил", "Фирдавс"})
	fmt.Println(result)
	//Output:[]
}

func ExampleSearchByRegions_OneOrMoreResult() {
	result := searchByRegions(houses, []string{"Исмоили Сомони", "Фирдавси", "Рудаки"})
	fmt.Println(result)
	//Output:[{4 Пентхаус в центре города 500 50 Исмоили Сомони} {2 3-х ком. квартира возле рынка Саховат 80 200 Фирдавси} {1 2-х ком квартира в центре города 100 100 Рудаки}]
}
