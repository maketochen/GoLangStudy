package main

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Order struct { //tag
	ID         string      `json:"id,omitempty"`
	Items      []OrderItem `json:"items"`
	Quantity   int         `json:"quantity,omitempty"`
	TotalPrice float64     `json:"total_price,omitempty"`
}

type OrderItem struct {
	ID    string  `json:"id,omitempty"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

func main() {
	r := gin.Default()

	r.GET("/", func(context *gin.Context) {
		context.String(http.StatusOK, "OK")
		context.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run()
}
func marshal() {
	o := Order{
		ID:         "1234",
		Quantity:   3,
		TotalPrice: 30.3,
		Items: []OrderItem{
			{ID: "item1",
				Name:  "learn go",
				Price: 30},

			{ID: "item1",
				Name:  "learn go",
				Price: 30},
		},
	}

	b, err := json.Marshal(o)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", b)
}
func unmarshal() {
	s := `{"id":"1234","item":{"id":"item1","name":"learn go","price":30},"quantity":3,"total_price":30.3}`
	var o Order
	err := json.Unmarshal([]byte(s), &o)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", o)
}
