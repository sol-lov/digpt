package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"digpt/model"
)

func main() {
	url := "https://api.digikala.com/v2/product/16175190/"
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error fetching data:", err)
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response:", err)
		return
	}

	err = ioutil.WriteFile("response.json", body, 0644)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}
	 

	var response model.Response
	err = json.Unmarshal(body, &response)
	if err != nil {
		fmt.Println("Error parsing JSON:", err)
		return
	}

	if response.Data.Product.IsInactive {
		fmt.Println("Product is inactive.")
	} else {

		fmt.Println("Product ID:", response.Data.Product.ID)
		fmt.Println("Title (FA):", response.Data.Product.TitleFa)
		fmt.Println("Category:", response.Data.Product.DataLayer.Category)

	}
}
