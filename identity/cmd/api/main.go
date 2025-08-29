package main

import (
	"database/sql"

	"github.com/PatrochR/whatashop/handler"
	"github.com/PatrochR/whatashop/repository"
	"github.com/PatrochR/whatashop/router"
	"github.com/charmbracelet/log"
	_ "github.com/lib/pq"
)

func main() {

	db, err := ConnectionDatabase()
	if err != nil {
		log.Fatal(err)
	}
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	log.Info("database connected")
	userRepo := repository.NewUserPostgres(db)
	userHandler := handler.NewUserHandler(userRepo)
	router := router.NewRouter(":3000", userHandler)
	if err := router.Run(); err != nil {
		log.Fatal("server problem")
	}


}

func ConnectionDatabase() (*sql.DB, error) {
	connectionString := "user=postgres dbname=postgres port=5432 password=dune sslmode=disable host=localhost"
	db, err := sql.Open("postgres", connectionString)

	if err != nil {
		return nil, err
	}
	return db, nil
}
