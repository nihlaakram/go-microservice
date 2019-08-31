package model

import "database/sql"

type Article struct {
	Id      int64  `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
	Author  string `json:"author"`
}

func (article *Article) AddArticle(db *sql.DB) error {
	if res, err := db.Exec("INSERT INTO articles VALUES(?, ?, ?, ?)", 0, article.Title, article.Content, article.Author); err != nil {
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