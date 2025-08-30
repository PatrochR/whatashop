package main

import (
	"database/sql"
	"os"

	"github.com/PatrochR/whatashop/handler"
	"github.com/PatrochR/whatashop/repository"
	"github.com/PatrochR/whatashop/router"
	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/log"
	_ "github.com/lib/pq"
)

func main() {
	style := log.DefaultStyles()
	style.Keys["Error"] = lipgloss.NewStyle().Foreground(lipgloss.Color("204"))
	logger := log.NewWithOptions(os.Stderr, log.Options{
		ReportTimestamp: true,
	})
	logger.SetStyles(style)

	db, err := ConnectionDatabase()
	if err != nil {
		logger.Fatal(err)
	}
	err = db.Ping()
	if err != nil {
		logger.Fatal(err)
	}
	log.Info("database connected")
	userRepo := repository.NewUserPostgres(db)
	userHandler := handler.NewUserHandler(logger, userRepo)
	router := router.NewRouter(":3000", userHandler)
	if err := router.Run(); err != nil {
		logger.Fatal("server problem")
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
