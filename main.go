package main

import (
  // "net/http"
  "github.com/gin-gonic/gin"
  "github.com/gin-gonic/gin/binding"
)

type BookUsageStatisticItemForm struct {
  ProfileId int `form:"profile_id" json:"ProfileId" binding required`
  BookItemId int `form:"book_item_id" json:"BookItemId" binding required`
}

type BookUsageStatisticItem struct {
  Id int `db:"id"`
  ProfileId int `db:"profile_id"`
  BookItemId int `db:"book_item_id"`
}

func main() {
  r := gin.Default()

  r.POST("/book_usage_statistic_items", func(c *gin.Context){

    var form BookUsageStatisticItemForm

    c.BindWith(&form, binding.MultipartForm)

    c.JSON(200, &form)

    // var Busi BookUsageStatisticItem
    // curl -v --form profile_id=1 --form book_item_id=2 http://localhost:8080/book_usage_statistic_items
  })

  r.Run(":8080")
}



