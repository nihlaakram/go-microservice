package service

import ("database/sql"
	"fmt"
	"log"
	"github.com/gorilla/mux"
	"net/http"
	_ "github.com/go-sql-driver/mysql"
)

type Server struct {
	Router *mux.Router
	DBCon  *sql.DB
}

func (service *Server) Init(dbUser, dbPass, dbName, hostname, mysqlPort string) {

	var err error
	service.DBCon, err = sql.Open("mysql", dbUser+":"+dbPass+"@tcp("+hostname+":"+mysqlPort+")/"+dbName)
	if err != nil {
		log.Fatal(err)
	}

	service.Router = mux.NewRouter()
	service.initResource()
}

func (service *Server) Start(port int) {

	err := http.ListenAndServe(fmt.Sprintf(":%v", port), service.Router)
	if err != nil {
		log.Fatal(err)
	}
}

func (service *Server) initResource() {
	service.Router.HandleFunc("/articles", nil).Methods(http.MethodPost)
	service.Router.HandleFunc("/articles", nil).Methods(http.MethodGet)
	service.Router.HandleFunc("/articles/{id:[0-9]+}", nil).Methods(http.MethodGet)
}