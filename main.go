package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"net/http"
	"time"
	"unicode"
)

type busiForm struct {
	ProfileId  int `form:"profile_id" binding:"required"`
	BookItemId int `form:"book_item_id" binding:"required"`
}

type timestamps struct {
	CreatedAt time.Time
	UpdatedAt time.Time
}

type pkey struct {
	Id int64
}

type busiView struct {
	pkey
	busiForm
	timestamps
}

func main() {
	r := gin.Default()

	db, err := sqlx.Connect("postgres", "user=konjoot dbname=reeky sslmode=disable")
	db.MapperFunc(toUnderscore)

	if err != nil {
		panic("Can't establish database connection!")
	}

	r.POST("/book_usage_statistic_items", func(c *gin.Context) {

		var form busiForm

		if errs := c.Bind(&form); errs != nil {
			c.JSON(http.StatusBadRequest, gin.H{"Errors": errs})
			return
		}

		stmt, err := db.Preparex(`
		INSERT INTO book_usage_statistic_items (profile_id, book_item_id)
			VALUES ($1, $2) RETURNING id, created_at, updated_at`)

		if err != nil {
			c.JSON(http.StatusNotAcceptable, gin.H{"Error": err})
			return
		}

		view := busiView{busiForm: form}

		stmt.QueryRowx(form.ProfileId, form.BookItemId).StructScan(&view)

		c.JSON(200, &view)
	})

	r.GET("/book_usage_statistic_items/:id", func(c *gin.Context) {
		stmt, err := db.Preparex(`
		SELECT * FROM book_usage_statistic_items
			WHERE id = $1 LIMIT 1`)

		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"Error": err})
			return
		}

		var view busiView

		if err = stmt.Get(&view, c.Param("id")); err != nil {
			c.JSON(http.StatusNotFound, gin.H{"Error": err.Error()})
			return
		}

		c.JSON(200, &view)
	})

	r.Run(":8080")
}

func toUnderscore(s string) string {
	runes := []rune(s)
	length := len(runes)

	var out []rune

	for i, s := range runes {
		if i > 0 && unicode.IsUpper(s) && ((i+1 < length && unicode.IsLower(runes[i+1])) || unicode.IsLower(runes[i-1])) {
			out = append(out, '_')
		}
		out = append(out, unicode.ToLower(runes[i]))
	}

	return string(out)
}
