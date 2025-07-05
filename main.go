package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	gh "github.com/gorilla/handlers"
	"github.com/softilium/elorm"

	"github.com/flosch/pongo2/v6"

	_ "modernc.org/sqlite"
)

var dbc *DbContext
var err error

func logErr(err error) {
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
}

func handleHttpErr(err error, w http.ResponseWriter) {
	if err != nil {
		log.Println(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		panic(err)
	}
}

func seed(onlyClean bool) {
	recs := 30
	for _, ed := range dbc.EntityDefs {
		_, err := dbc.DB.Exec(fmt.Sprintf("delete from %s", ed.TableName))
		logErr(err)
	}

	if onlyClean {
		return
	}

	for i := range recs {
		user, err := dbc.CreateUser()
		logErr(err)

		user.SetUsername(fmt.Sprintf("User %d", i+1))
		user.SetEmail(fmt.Sprintf("user%d@gmail.com", i+1))
		err = user.Save()
		logErr(err)

		for ii := range recs {
			todo, err := dbc.CreateTodoItem()
			logErr(err)
			todo.SetTitle(fmt.Sprintf("Todo %d for User %d", ii+1, i+1))
			todo.SetDescription(fmt.Sprintf("Description for Todo %d for User %d", ii+1, i+1))
			todo.SetOwner(user)
			todo.SetDueDate(time.Now().Add(time.Duration(ii+1) * time.Hour * 24))
			err = todo.Save()
			logErr(err)

			for iii := range 3 {
				cmt, err := dbc.CreateTodoComment()
				logErr(err)
				cmt.SetTodoItem(todo)
				cmt.SetAuthor(user)
				cmt.SetContent(fmt.Sprintf("Comment %d for Todo %d for User %d", iii+1, ii+1, i+1))

				err = cmt.Save()
				logErr(err)
			}
		}
	}
}

func handleStaticPage(tmplFileName string) func(http.ResponseWriter, *http.Request) {

	return func(w http.ResponseWriter, r *http.Request) {

		pongo2.DefaultSet.CleanCache()
		tmpl, err := pongo2.FromCache(tmplFileName)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		err = tmpl.ExecuteWriter(pongo2.Context{}, w)
		handleHttpErr(err, w)
	}

}

func initServer(listenAddr string) *http.Server {

	router := http.NewServeMux()

	fs := http.FileServer(http.Dir("./assets/"))
	router.Handle("/assets/", http.StripPrefix("/assets", fs))

	usersRestApiConfig := elorm.CreateStdRestApiConfig(
		*dbc.UserDef.EntityDef,
		dbc.LoadUser,
		dbc.UserDef.SelectEntities,
		dbc.CreateUser)
	router.HandleFunc("/api/users", elorm.HandleRestApi(usersRestApiConfig))

	todosRestApiConfig := elorm.CreateStdRestApiConfig(
		*dbc.TodoItemDef.EntityDef,
		dbc.LoadTodoItem,
		dbc.TodoItemDef.SelectEntities,
		dbc.CreateTodoItem)
	router.HandleFunc("/api/todos", elorm.HandleRestApi(todosRestApiConfig))

	commentsRestApiConfig := elorm.CreateStdRestApiConfig(
		*dbc.TodoCommentDef.EntityDef,
		dbc.LoadTodoComment,
		dbc.TodoCommentDef.SelectEntities,
		dbc.CreateTodoComment)
	router.HandleFunc("/api/comments", elorm.HandleRestApi(commentsRestApiConfig))

	router.HandleFunc("/api/seed", func(w http.ResponseWriter, r *http.Request) {
		seed(false)
		w.WriteHeader(http.StatusOK)
	})

	router.HandleFunc("/api/clear", func(w http.ResponseWriter, r *http.Request) {
		seed(true)
	})

	router.HandleFunc("/", handleStaticPage("pages/index.html"))

	server := &http.Server{
		Addr:    listenAddr,
		Handler: router,
	}

	server.Handler = gh.LoggingHandler(os.Stdout, server.Handler)

	return server

}

func gracefullShutdown(server *http.Server, quit <-chan os.Signal, done chan<- bool) {

	<-quit
	log.Println("Server is shutting down...")

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	server.SetKeepAlivesEnabled(false)
	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("Could not gracefully shutdown the server: %v\n", err)
	}
	close(done)

}

func main() {

	dbc, err = CreateDbContext("sqlite", "./todo.db")
	logErr(err)

	err = dbc.EnsureDBStructure()
	logErr(err)

	done := make(chan bool, 1)
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)

	listenAddr := "localhost:8080"

	server := initServer(listenAddr)
	go gracefullShutdown(server, quit, done)

	log.Println("Server is ready to handle requests at", listenAddr)

	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("could not listen on %s: %v\n", listenAddr, err.Error())
	}

	log.Println("Server stopped")

	<-done

}
