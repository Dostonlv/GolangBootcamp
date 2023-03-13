package main

import (
	"fmt"
	"maps/maps"
)

func main() {
	authors := make(map[string]string)
	authors["ArmandoLucas"] = "The Daughter's Tale"
	authors["J.K.Rowling"] = "Harry Potter and the Deathly Hallows"
	authors["StephenieMeyer"] = "Twilight"
	authors["MargaretWiseBrown"] = "Goodnight Moon"
	var author string
	maps.Logger(authors)
	fmt.Print("whether you want to 'search' or 'delete' a book please input Author name: ")
	var FinderOrDeleter string
	fmt.Scan(&FinderOrDeleter)

	switch FinderOrDeleter {
	case "search":
		print("Please input book's author name: ")
		fmt.Scan(&author)
		maps.MapFindLog(authors, author)
	case "delete":
		print("Please input book's author name for delete this list: ")
		fmt.Scan(&author)
		maps.DeleteBooks(authors, author)
		maps.Logger(authors)

	}
}
