package main

import (
	"encoding/json"
	"fastfood/SumPrice"
	"fastfood/WriteAndRead"
	"fmt"
	"io/ioutil"
)

type Food struct {
	Id       int    `jsonning:"id"`
	FoodName string `jsonning:"foodName"`
	Price    int    `jsonning:"price"`
}

type Order struct {
	OrderId  int      `jsonning: "orderId"`
	FoodName []string `jsonning:"foodName"`
	Price    int      `jsonning:"price"`
}

func main() {

	data, err := WriteAndRead.ReadNameAndPrice()
	var foods []Food
	err = json.Unmarshal(data, &foods)
	if err != nil {
		fmt.Println("Error unmarshaling JSON:", err)
		return
	}

	for _, food := range foods {
		fmt.Println(food.Id, food.FoodName, "-", food.Price)
	}
	var orders []Order
	var orderQueue int
	fmt.Print("Ovqat olishmochi bo'lgan odamlar sonini kiriting: ")
	fmt.Scan(&orderQueue)
	for i := 1; i <= orderQueue; i++ {

		var FoodName []string
		var orderNum int
		var order int
		var foodP []int
		fmt.Printf("Salom %v-buyurtmachi qancha mahsulot buyurtma qilmoqchisiz? : ", i)
		fmt.Scanln(&orderNum)
		for _, food := range foods {
			fmt.Println(food.Id, food.FoodName, "-", food.Price)
		}
		for j := 0; j < orderNum; j++ {
			fmt.Print("mahsulot idsini kiriting: ")
			fmt.Scan(&order)
			for _, food := range foods {
				if food.Id == order {
					FoodName = append(FoodName, food.FoodName)
					foodP = append(foodP, food.Price)
				}
			}
		}
		var sumPrice = SumPrice.Sum(foodP)

		orders = append(orders, Order{OrderId: i, FoodName: FoodName, Price: sumPrice})

	}
	ordered, err := json.Marshal(orders)
	if err != nil {
		fmt.Println("Error marshaling JSON:", err)
		return
	}

	err = ioutil.WriteFile("orders.jsonning", ordered, 0644)
	if err != nil {
		fmt.Println("Error writing JSON to file:", err)
		return
	}

}
