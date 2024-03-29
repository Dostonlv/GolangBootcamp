package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type AutoGenerated struct {
	UserID int    `jsonning:"userId"`
	ID     int    `jsonning:"id"`
	Title  string `jsonning:"title"`
	Body   string `jsonning:"body"`
}

func main() {
	client := http.Client{}
	resp, err := client.Get("https://jsonplaceholder.typicode.com/posts")
	if err != nil {
		fmt.Println(err)
		return
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	var post []AutoGenerated
	err = json.Unmarshal(body, &post)
	if err != nil {
		fmt.Println("Error unmarshaling JSON:", err)
		return
	}

	for _, post := range post {
		fmt.Printf("{%v,%v,%v,%v \n }", post.UserID, post.ID, post.Title, post.Body)

	}

}
