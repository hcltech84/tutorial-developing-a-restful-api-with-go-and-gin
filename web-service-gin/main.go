package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

type Page struct {
	ID string `json:"id"`
	Title string `json:"title"`
	Links []string `json:"links"`
}

var pages = []Page{
	{
		ID: "1",
		Title: "Tutorial: Get started with Go",
		Links: []string{
			"https://go.dev/doc/tutorial/getting-started",
			"https://github.com/hcltech84/tutorial-get-started-with-go",
		},
	},
	{
		ID: "2",
		Title: "Tutorial: Create a Go module",
		Links: []string{
			"https://go.dev/doc/tutorial/create-module",
			"https://github.com/hcltech84/tutorial-create-a-go-module",
		},
	},
	{
		ID: "3",
		Title: "Tutorial: Getting started with multi-module workspaces",
		Links: []string{
			"https://go.dev/doc/tutorial/workspaces",
			"https://github.com/hcltech84/tutorial-getting-started-with-multi-module-workspace",
		},
	},
	{
		ID: "4",
		Title: "Tutorial: Developing a RESTful API with Go and Gin",
		Links: []string{
			"https://go.dev/doc/tutorial/web-service-gin",
			"https://github.com/hcltech84/tutorial-developing-a-restful-api-with-go-and-gin",
		},
	},
}

func main() {
	router := gin.Default()
	router.GET("/pages", getPages)
	router.POST("/pages", postPages)
	router.GET("/pages/:id", getPage)
	router.Run("localhost:8080")
}

func getPages(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, pages)
}

func postPages(c *gin.Context) {
	var newPage Page
	c.BindJSON(&newPage)
	pages = append(pages, newPage)
	c.IndentedJSON(http.StatusCreated, gin.H{"message": "Successfully added a new page"})
}

func getPage(c *gin.Context) {
	for _, page := range pages {
		if page.ID == c.Param("id") {
			c.IndentedJSON(http.StatusOK, page)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "The page you request is not found"})
}
