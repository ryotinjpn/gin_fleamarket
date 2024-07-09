package main

import (
	"gin-fleamarket/controllers"
	"gin-fleamarket/infra"
	"gin-fleamarket/repositories"
	"gin-fleamarket/services"

	"github.com/gin-gonic/gin"
)

func main() {
	infra.Initialize()
	db := infra.SetupDB()

	itemRepository := repositories.NewItemRepository(db)
	itemService := services.NewItemService(itemRepository)
	itemContoroller := controllers.NewItemController(itemService)

	r := gin.Default()
	itemRouter := r.Group("/items")
	itemRouter.GET("", itemContoroller.FindAll)
	itemRouter.GET("/:id", itemContoroller.FindById)
	itemRouter.POST("", itemContoroller.Create)
	itemRouter.PUT("/:id", itemContoroller.Update)
	itemRouter.DELETE("/:id", itemContoroller.Delete)
	r.Run("localhost:8080")
}
