package main

import (
	"log"
	"github.com/gin-gonic/gin"
)

type Cloth struct {
	ID	     string `json:"id"`
	Name     string `json:"name"`
	Price    string `json:"price"`
	Quantity string `json:"quantity"`
	Material string `json:"material"`
}

var clothes = []Cloth{
	{ID: "1", Name: "Round-neck shirt", Price: "$5", Quantity: "10", Material: "cotton"},
	{ID: "2", Name: "Polo-shirt", Price: "$8", Quantity: "20", Material: "cotton"},
	{ID: "3", Name: "Hoodie", Price: "$15", Quantity: "4", Material: "wool"},
	{ID: "4", Name: "Turtle-neck", Price: "$5", Quantity: "10", Material: "cotton"},
	{ID: "5", Name: "Jacket", Price: "$30", Quantity: "3", Material: "leather"},
}

func getClothes(c *gin.Context) {
	c.JSON(200, clothes)
}

func getclothbyID(c *gin.Context){
	id := c.Param("id")
	var index int
	for k, v := range(clothes){
		if id == v.ID{
			index = k
		}  
	}
	c.JSON(200, clothes[index])
}

func addcloth(c *gin.Context){
	var newCloth Cloth
	err := c.BindJSON(&newCloth)
	if err != nil{
		log.Fatal(err)
		return
	} 
	clothes = append(clothes, newCloth)
	c.JSON(201, clothes)
}

func deletecloth(c *gin.Context){
	id := c.Param("id")
	var index int
	for k, v := range(clothes){
		if id == v.ID{
			index = k
		}
	} 
	clothes = append(clothes[:index], clothes[index+1:]...)
	c.JSON(200, clothes)
}

func updateCloth(c *gin.Context){
	id := c.Param("id")
	var index int
	for k, v := range(clothes){
		if id == v.ID{
			index = k 
		} 
	}
	clothes[index].Price = "$12"
	c.JSON(200, clothes[index])
}

func requesthandler() {
	r := gin.Default()
	r.GET("/clothes", getClothes) 
	r.GET("/cloth/:id", getclothbyID)
	r.POST("/newcloth", addcloth)
	r.DELETE("/cloth/:id", deletecloth)
	r.PATCH("/cloth/:id", updateCloth)
	err := r.Run("localhost:8080")
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	requesthandler()
}
