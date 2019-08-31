package service

import ("database/sql"
	"github.com/gorilla/mux")

type Server struct {
	Router *mux.Router
	DBCon  *sql.DB
}