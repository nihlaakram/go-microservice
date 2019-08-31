package model

import "database/sql"

type Article struct {
	Id      int64  `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
	Author  string `json:"author"`
}

type ArticleId struct {
	Id int64 `json:"id"`
}

const insertArticleQuery = "INSERT INTO articles VALUES(?, ?, ?, ?)"
const getArticleByIdQuery = "SELECT title, content, author  FROM articles WHERE id=?"
const getAllArticles = "SELECT * FROM articles"

// Adds an article and gets the id
func (article *Article) AddArticle(db *sql.DB) error {
	if res, err := db.Exec(insertArticleQuery, 0, article.Title, article.Content, article.Author); err != nil {
		return err
	} else {
		id, err := res.LastInsertId()
		if err != nil {
			return err
		}
		article.Id = id
	}
	return nil

}

// Gets the article corresponding for the given id
func (article *Article) GetArticleById(db *sql.DB) error {
	return db.QueryRow(getArticleByIdQuery, article.Id).Scan(&article.Title, &article.Content, &article.Author)
}

// Gets all article corresponding
func GetAllArticles(db *sql.DB) ([]Article, error) {

	result, err := db.Query(getAllArticles)
	if err != nil {
		return nil, err
	}
	defer result.Close()
	articles := []Article{}

	for result.Next() {
		var article Article

		if err = result.Scan(&article.Id, &article.Title, &article.Content, &article.Author); err != nil {
			return nil, err
		}
		articles = append(articles, article)
	}

	return articles, nil
}
