package ui

import (
	"net/http"

	"github.com/SealTV/handmade-shope/model"
	"github.com/gin-gonic/gin"
)

// Define the route for the index page and display the index.html template
// To start with, we'll use an inline route handler. Later on, we'll create
// standalone functions that will be used as route handlers.
func InitIndexPage(m *model.Model) func(c *gin.Context) {
	return func(c *gin.Context) {
		users, err := m.GetAllUsers()
		if err != nil {
			c.HTML(http.StatusInternalServerError, "idnex.html", gin.H{
				"title": "Error",
			})

			return
		}

		// Call the HTML method of the Context to render a template
		c.HTML(
			// Set the HTTP status to 200 (OK)
			http.StatusOK,
			// Use the index.html template
			"index.html",
			// Pass the data that the page uses (in this case, 'title')
			gin.H{
				"title": "Home Page",
				"users": users,
			},
		)
	}
}
