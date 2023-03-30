package main

import (
	"encoding/json"
	"fmt"
	"github.com/TwiN/go-color"
	"log"
	"os"
	"shop/models"
)

func main() {
	prod, err := os.ReadFile("../data/products.json")
	if err != nil {
		log.Fatal(err)
	}

	var product []models.Product

	err = json.Unmarshal(prod, &product)
	if err != nil {
		fmt.Println("Error unmarshalling JSON:", err)
		return
	}

	user, err := os.ReadFile("../data/users.json")
	if err != nil {
		log.Fatal(err)
	}
	var users []models.User

	err = json.Unmarshal(user, &users)
	if err != nil {
		fmt.Println("Error unmarshalling JSON:", err)
		return
	}

	Users := &models.Users{Users: users}

	var have bool
	var name string
	fmt.Println(`----------------- Welcome to Shop------------------`)
	fmt.Print("Enter name: ")
	fmt.Scan(&name)

	for _, us := range Users.SendUsers() {
		if us.Name == name {
			have = true
		}
	}

	Products := &models.Products{Products: product}
	if have {
		fmt.Printf("hello %v what do you want to buy \n ", name)
		print(color.Ize(color.Red, "**************************************\n"))
		for _, pro := range Products.PrintProduct() {

			fmt.Printf(`---------------------
product id: %v
product name: %v
product price: %v
`, pro.ID, pro.Name, pro.Price)
		}
		var NameArray []string
		var ProductName string
	START:
		fmt.Print("enter product name if you buy products enter stop: ")

		fmt.Scan(&ProductName)
		for _, v := range Products.PrintProduct() {
			if v.Name == ProductName {
				NameArray = append(NameArray, ProductName)
				if ProductName == "stop" {
					break
				}
				goto START
			}
		}
		sum := 0
		for _, v := range Products.PrintProduct() {
			for _, k := range NameArray {
				if v.Name == k {
					sum += v.Price
				}
			}
		}
		var UserArray []models.User
		var user int
		var users models.User

		for _, v := range Users.SendUsers() {
			if v.Name == name {
				user = v.Balance
				users = v
			}
		}
		if user <= 0 {
			fmt.Println("there is no money on your card")
			goto STOP
		} else if user-sum <= 0 {
			fmt.Println("the money on your card is not enough to buy goods")
			goto STOP
		} else {
			users.Balance = user - sum
		}

		for _, v := range Users.SendUsers() {
			if v.Name != name {
				UserArray = append(UserArray, v)
			}
		}

		UserArray = append(UserArray, users)

		js, err := json.Marshal(UserArray)
		if err != nil {
			fmt.Errorf("failed to marshal tasks to JSON: %v", err)
		}
		err = os.WriteFile("../data/users.json", js, 0644)

		fmt.Printf("%v soums have been debited from your account", sum)
		da, err := os.ReadFile("../data/benifit.json")
		if err != nil {
			log.Fatal(err)
		}
		var b models.Benefit

		err = json.Unmarshal(da, &b)
		b.Benefit += sum
		jsonData, err := json.Marshal(b)
		if err != nil {
			fmt.Errorf("failed to marshal tasks to JSON: %v", err)
		}
		err = os.WriteFile("../data/benifit.json", jsonData, 0644)
	} else {
		fmt.Println("This user not")
	}
STOP:
	fmt.Println("please complete your account")

}
