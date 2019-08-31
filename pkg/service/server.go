package service

import ("database/sql"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

type Server struct {
	Router *mux.Router
	DBCon  *sql.DB
}

func (service *Server) Init(dbUser, dbPass, dbName, hostname, mysqlPort string) {

	var err error
	service.DBCon, err = sql.Open("mysql", dbUser+":"+dbPass+"@tcp("+hostname+":"+mysqlPort+")/"+dbName)
	if err != nil {
		fmt.Println(err)
	}

	service.Router = mux.NewRouter()
	service.initResource()
}


func (service *Server) initResource() {
	service.Router.HandleFunc("/articles", nil).Methods(http.MethodPost)
	service.Router.HandleFunc("/articles", nil).Methods(http.MethodGet)
	service.Router.HandleFunc("/articles/{id:[0-9]+}", nil).Methods(http.MethodGet)
}