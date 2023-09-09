package main

import (
	"encoding/json"
	"fmt"
	"time"
	"sync"
)

type OrderFindResponse struct {
	OrderID     int32    `json:"order_id"`
	OrderStatus string   `json:"order_status"`
	// Add other fields as needed
	Products []string `json:"products"`
}

type OrderFindProductsResponse struct {
	OrderID  int32    `json:"order_id"`
	Products []string `json:"products"`
}

type OrderFindProductIngredientsResponse struct {
	ProductID   int32    `json:"product_id"`
	Ingredients []string `json:"ingredients"`
}

func mockOrderFindResponse(orderID int32 , respch chan any) string {
	// Simulate a 500ms delay
	time.Sleep(100 * time.Millisecond)

	response := OrderFindResponse{
		OrderID:     orderID,
		OrderStatus: "Mocked Status",
		Products:    []string{},
	}

	jsonResponse, _ := json.Marshal(response)
	return string(jsonResponse)
}

func mockOrderFindProductsResponse(orderID int32 , respch chan any ,wg *sync.WaitGroup) string {
	// Simulate a 500ms delay
	time.Sleep(3000 * time.Millisecond)

	response := OrderFindProductsResponse{
		OrderID:  orderID,
		Products: []string{"Product 1", "Product 2"},
	}

	jsonResponse, _ := json.Marshal(response)

	respch <- string(jsonResponse)
	wg.Done()
	return string(jsonResponse)
}

func mockOrderFindProductIngredientsResponse(productID int32 , respch chan any ,wg *sync.WaitGroup) string {
	// Simulate a 500ms delay
	time.Sleep(2900 * time.Millisecond)

	response := OrderFindProductIngredientsResponse{
		ProductID:   productID,
		Ingredients: []string{"Ingredient 1", "Ingredient 2"},
	}

	jsonResponse, _ := json.Marshal(response)
	respch <- string(jsonResponse)
	wg.Done()
	return string(jsonResponse)
}

func main() {
	startTime := time.Now()

	orderID := int32(1)
	productID := int32(2)

	respch := make(chan any , 100)
	mockOrderFindResponse(orderID , respch)
	// fmt.Println(orderResponse , respch)
	wg := &sync.WaitGroup{}
	wg.Add(2)

	go mockOrderFindProductsResponse(orderID , respch  , wg)
	// fmt.Println(productsResponse , respch  , wg)

	go mockOrderFindProductIngredientsResponse(productID , respch  , wg)
	wg.Wait()
	close(respch)
	// fmt.Println(ingredientsResponse , respch)
	 for resp := range respch {
		fmt.Println(resp)
	 }
	elapsedTime := time.Since(startTime)
	fmt.Printf("Total execution time: %s\n", elapsedTime  )
}
