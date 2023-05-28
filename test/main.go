package main

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"go-chi-example/api/controller"
	"go-chi-example/api/handler"
	"go-chi-example/api/repository"
	"go-chi-example/api/router"
	"net/http"
	"os"
	"strconv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("Error loading .env file")
	}
	host := os.Getenv("host")
	port, _ := strconv.Atoi(os.Getenv("port"))
	user := os.Getenv("user")
	password := os.Getenv("password")
	dbname := os.Getenv("dbname")
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	boil.SetDB(db)
	r := chi.NewRouter()
	repo := repository.New(db)
	todoCtrl := controller.New(repo)
	todoHandler := handler.New(todoCtrl)
	r.Mount("/todos", router.Router{
		Ctx:         context.Background(),
		TodoHandler: todoHandler,
	}.HandlerTodo())
	fmt.Println("Successfully connected!")
	http.ListenAndServe(":3000", r)
}
