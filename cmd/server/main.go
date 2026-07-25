package main

import (
	"fmt"
	"grpc/internal/database"
)

func main() {
	db := database.Connect()

	fmt.Println("databse cncted",db!=nil)
}