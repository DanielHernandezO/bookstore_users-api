package main

import "github.com/DanielHernandezO/bookstore_users-api/config"

func main() {
	strategy := config.NewStrategy()
	strategy.Start()
}
