package main

import "testing"

type article struct {
	ID		int		`json:"id"`
	Title	string	`json:"title"`
	Content	string	`json:"content"`
}

var articleList = []article{
	article{ID: 1, Title: "Article_1", Content: "Article 1 body"},
	article{ID: 2, Title: "Article 2", Content: "Lorem ipsum et dolor sit amet"},
}

func TestGetAllArticles(t *testing.T) {
	alist := getAllArticles()

	// check that the length of the list of articles returned is the same as the length of the global varialbe holding the list
	if len(alist) != len(articleList) {
		t.Fail()
	}

	// check that each member is identical
	for i, v := range alist{
		if v.Content != articleList[i].Content ||
		   v.ID != articleList[i].ID ||
		   v.Title != articleList[i].Title {

			t.Fail()
			break
		}
	}
}

func getAllArticles() []article {
	return articleList
}