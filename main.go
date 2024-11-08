package main

import (
	"context"
	"log"

	"go-dao/dao"
	"go-dao/routes"
	"go-dao/services"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
)

func main() {

	db, err := pgx.Connect(context.Background(), "postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable")
	if err != nil {
		log.Fatal("Unable to connect to database:", err)
	}
	defer db.Close(context.Background())

	personDAO := dao.NewPersonDao(db)
	personService := services.NewPersonService(personDAO)

	KelasDAO := dao.NewKelasDao(db)
	KelasService := services.NewKelasService(KelasDAO)

	r := gin.Default()
	routes.SetupPersonRoutes(r, personService)
	routes.SetupKelasRoutes(r, KelasService)

	r.Run(":8080")
}
