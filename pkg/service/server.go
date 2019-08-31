package service

import ("database/sql"
	"encoding/json"
	"fmt"
	"github.com/nihlaakram/go-microservice/pkg/model"
	"github.com/nihlaakram/go-microservice/pkg/util"
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
	service.Router.HandleFunc("/articles", service.addArticle).Methods(http.MethodPost)
	service.Router.HandleFunc("/articles", nil).Methods(http.MethodGet)
	service.Router.HandleFunc("/articles/{id:[0-9]+}", nil).Methods(http.MethodGet)
}

func (service *Server) addArticle(w http.ResponseWriter, r *http.Request) {
	var article model.Article
	body := json.NewDecoder(r.Body)
	if err := body.Decode(&article); err != nil {
		failureResponse(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()

	if err := article.AddArticle(service.DBCon); err != nil {
		failureResponse(w, http.StatusInternalServerError, err.Error())
		return
	}
	writeResponse(w, http.StatusCreated, util.SuccessMsg, article.Id)
}

func failureResponse(w http.ResponseWriter, code int, message string) {
	writeResponse(w, code, message, nil)
}

func writeResponse(w http.ResponseWriter, code int, message string, data interface{}) {
	response := model.Response{code, message, data}
	payload, _ := json.Marshal(response)

	w.Header().Set(util.ContentType, util.ApplicationJson)
	w.WriteHeader(code)
	w.Write(payload)
}