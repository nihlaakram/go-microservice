package model

import "database/sql"

type Article struct {
	Id      int64  `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
	Author  string `json:"author"`
}

const insertArticleQuery = "INSERT INTO articles VALUES(?, ?, ?, ?)"
const getArticleByIdQuery = "SELECT title, content, author  FROM articles WHERE id=?"

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

func (article *Article) GetArticleById(db *sql.DB) error {
	return db.QueryRow(getArticleByIdQuery, article.Id).Scan(&article.Title, &article.Content, &article.Author)
}
