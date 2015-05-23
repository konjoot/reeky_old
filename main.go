package main

import (
	"github.com/konjoot/gin"
	"net/http"
)

type BookUsageStatisticItem struct {
	ProfileId  int `form:"profile_id" binding:"required"`
	BookItemId int `form:"book_item_id" binding:"required"`
}

type BookUsageStatisticItemDB struct {
	Id         int `db:"id"`
	ProfileId  int `db:"profile_id"`
	BookItemId int `db:"book_item_id"`
}

func main() {
	r := gin.Default()

	r.POST("/book_usage_statistic_items", func(c *gin.Context) {

		var form BookUsageStatisticItem

		if errs := c.Bind(&form); errs != nil {
			c.JSON(http.StatusBadRequest, gin.H{"errors": errs})
			return
		}

		c.JSON(200, &form)
		// curl -v --form profile_id=1 --form book_item_id=2 http://localhost:8080/book_usage_statistic_items
	})

	r.Run(":8080")
}
