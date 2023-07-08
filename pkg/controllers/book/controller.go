package book

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mrtc0/seccamp-2023/pkg/usecases"
)

type BookController struct {
	listBooks *usecases.ListBooks
}

func NewBookController(listBooks *usecases.ListBooks) *BookController {
	return &BookController{listBooks: listBooks}
}

func (con *BookController) List(c *gin.Context) {
	books, err := con.listBooks.Execute(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "ok",
		"data":    books,
	})
}
