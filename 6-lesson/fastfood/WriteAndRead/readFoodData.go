package WriteAndRead

import (
	"fmt"
	"io/ioutil"
)

func ReadNameAndPrice() ([]byte, error) {
	data, err := ioutil.ReadFile("foodnamesandprice.json")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return nil, nil
	}
	return data, err
}
