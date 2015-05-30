package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"net/http"
)

func main() {
	r := gin.Default()

	db, err := sqlx.Connect("postgres", "user=konjoot dbname=reeky sslmode=disable")

	if err != nil {
		panic("Can't establish database connection!")
	}

	r.POST("/book_usage_statistic_items", func(c *gin.Context) {
		type busiForm struct {
			ProfileId  int `form:"profile_id" binding:"required"`
			BookItemId int `form:"book_item_id" binding:"required"`
		}

		type busiView struct {
			Id int64
			busiForm
		}

		var form busiForm

		if errs := c.Bind(&form); errs != nil {
			c.JSON(http.StatusBadRequest, gin.H{"Errors": errs})
			return
		}

		var view = busiView{0, form}

		query := "INSERT INTO book_usage_statistic_items (profile_id, book_item_id) VALUES ($1, $2) RETURNING id;"

		if err := db.QueryRowx(query, form.ProfileId, form.BookItemId).Scan(&view.Id); err != nil {
			c.JSON(http.StatusNotAcceptable, gin.H{"Error": err})
			return
		}

		c.JSON(200, &view)
	})

	r.Run(":8080")
}
