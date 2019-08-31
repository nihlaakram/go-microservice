/*
 * Copyright (c) 2019, Nihla Akram. All Rights Reserved.
 */

package service

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/nihlaakram/go-microservice/pkg/model"
	"github.com/nihlaakram/go-microservice/pkg/util"
	"log"
	"net/http"
	"strconv"
)

type Server struct {
	Router *mux.Router
	DBCon  *sql.DB
}

const tableCreationQuery = `CREATE TABLE IF NOT EXISTS articles (
		id INT NOT NULL AUTO_INCREMENT,
		title VARCHAR(45) NULL,
		content TEXT NULL,
		author VARCHAR(45) NULL,
	PRIMARY KEY (id));`

// Initialize database connection and server route
func (service *Server) Init(dbUser, dbPass, dbName, hostname, mysqlPort string) {

	var err error
	log.Println(util.ConnectingDB)
	service.DBCon, err = sql.Open("mysql", dbUser+":"+dbPass+"@tcp("+hostname+":"+mysqlPort+")/"+dbName)
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println(util.DBConnectionSuccess)
		if _, err := service.DBCon.Exec(tableCreationQuery); err != nil {
			log.Fatal(err)
		}
	}

	service.Router = mux.NewRouter()
	service.initResource()
}

// Start the server in given port
func (service *Server) Start(port int) {
	log.Println(fmt.Sprintf(util.StartingServer, port))
	err := http.ListenAndServe(fmt.Sprintf(":%v", port), service.Router)
	if err != nil {
		log.Fatal(err)
	}
	http.Handle("/", service.Router)
}

// Initialize resource URLs
func (service *Server) initResource() {
	log.Println(util.DeployingResources)
	service.Router.HandleFunc(fmt.Sprintf("/%v", util.ArticlesResource), service.addArticle).Methods(http.MethodPost)
	service.Router.HandleFunc(fmt.Sprintf("/%v", util.ArticlesResource), service.getAllArticles).Methods(http.MethodGet)
	service.Router.HandleFunc(fmt.Sprintf("/%v/{id}", util.ArticlesResource), service.getArticleByID).Methods(http.MethodGet)
	log.Println(util.ResourcesDepSuccess)
}

// Handle POST requests of pattern /articles
func (service *Server) addArticle(w http.ResponseWriter, r *http.Request) {
	var article model.Article
	body := json.NewDecoder(r.Body)
	if err := body.Decode(&article); err != nil {
		failureResponse(w, http.StatusBadRequest, util.BadRequestMsg)
		return
	}

	defer r.Body.Close()

	if article.Title == "" || article.Content == "" || article.Author == "" {
		failureResponse(w, http.StatusBadRequest, util.BadRequestMsg)
		return
	}
	if err := article.AddArticle(service.DBCon); err != nil {
		failureResponse(w, http.StatusInternalServerError, err.Error())
		return
	}
	writeResponse(w, http.StatusCreated, util.SuccessMsg, model.ArticleId{Id: article.Id})
}

// Handle GET requests of pattern /articles/{id}
func (service *Server) getArticleByID(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 0, 0)
	if err != nil {
		failureResponse(w, http.StatusBadRequest, util.InvalidArticleIdMsg)
		return
	}
	article := model.Article{Id: id}
	err = article.GetArticleById(service.DBCon)
	if err != nil {
		switch err {
		case sql.ErrNoRows:
			failureResponse(w, http.StatusNotFound, util.ArticleNotFoundMsg)
			break
		default:
			failureResponse(w, http.StatusInternalServerError, err.Error())
			break
		}
		return
	}
	writeResponse(w, http.StatusOK, util.SuccessMsg, article)
}

// Handle GET requests of pattern /articles
func (service *Server) getAllArticles(w http.ResponseWriter, r *http.Request) {
	articles, err := model.GetAllArticles(service.DBCon)
	if err != nil {
		failureResponse(w, http.StatusInternalServerError, err.Error())
		return
	}
	writeResponse(w, http.StatusOK, util.SuccessMsg, articles)

}

// Generate failed response
func failureResponse(w http.ResponseWriter, code int, message string) {
	writeResponse(w, code, message, nil)
}

// write response
func writeResponse(w http.ResponseWriter, code int, message string, data interface{}) {
	response := model.Response{Status: code, Message: message, Data: data}
	payload, _ := json.Marshal(response)

	w.Header().Set(util.ContentType, util.ApplicationJson)
	w.WriteHeader(code)
	_, err := w.Write(payload)
	if err != nil {
		log.Fatal(err)
	}
}
