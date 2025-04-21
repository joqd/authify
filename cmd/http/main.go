package main

import (
	"fmt"

	"github.com/joqd/user-auth/internal/config"
)


func main() {
	config := config.LoadConfig()

	fmt.Println(config)
}