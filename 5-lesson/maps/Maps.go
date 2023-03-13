package maps

import "fmt"

func MapFindLog(authors map[string]string, a string) {

	for author, books := range authors {
		if author == a {
			fmt.Println(" Book: ", books)
		}
	}

}

func Logger(authors map[string]string) {
	for author, books := range authors {
		fmt.Println("Author:", author, " Book: ", books)
	}
}

func DeleteBooks(authors map[string]string, a string) {
	for author, _ := range authors {
		if author == a {
			delete(authors, author)
		}
	}
}
