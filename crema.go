package crema

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type server struct {
	DB     *sql.DB
	Router *mux.Router
}

func InitServer() *server {
	InitLogger()

	LogPrintf("[MAIN] Initializing Server ...")

	db := InitDatabase()
	router := InitRoutes()

	return &server{db, router}
}

/**
 * set up DB configuration (PostgreSQL) in conf
 * copy db.json.example as db.json
 */
func InitDatabase() *sql.DB {
	LogPrintf("[MAIN] Initializing Database ...")

	db := InitDB()
	SetDB(db)

	return db
}

func InitLogger() {
	log.Println("[MAIN] Initializing Logger ...")

	InitLogFiles()
	Printf("[MAIN] Initializing Logger ...")
}

func InitRoutes() *mux.Router {
	LogPrintf("[MAIN] Initializing Endpoints ...")

	router := mux.NewRouter()

	return router
}

func (s *server) AddRoutes(method string, routes string, handler func(http.ResponseWriter, *http.Request)) {
	s.Router.HandleFunc(routes, handler).Methods(method)
}
