package main

import (
	"golang/projects/go-hexagonal/entities"
	infra_database "golang/projects/go-hexagonal/infra/database"
	infra_secrets "golang/projects/go-hexagonal/infra/secrets"
)

func init() {
	infra_secrets.LoadEnv()
	infra_database.LoadDatabase()
}

func main() {
	infra_database.DB.AutoMigrate(&entities.Post{})
}
