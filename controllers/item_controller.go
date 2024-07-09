// リクエストデータのハンドリングやレスポンスに関する処理
package controllers

import (
	"gin-fleamarket/dto"
	"gin-fleamarket/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type IItemContoroller interface {
	FindAll(ctx *gin.Context)
	FindById(ctx *gin.Context)
	Create(ctx *gin.Context)
	Update(ctx *gin.Context)
	Delete(ctx *gin.Context)
}

type ItemContoroller struct {
	service services.IItemService
}

func NewItemController(service services.IItemService) IItemContoroller {
	return &ItemContoroller{service: service}
}

func (c *ItemContoroller) FindAll(ctx *gin.Context) {
	items, err := c.service.FindAll()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Unexpected error"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": items})
}

func (c *ItemContoroller) FindById(ctx *gin.Context) {
	itemId, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	item, err := c.service.FindById(uint(itemId))
	if err != nil {
		if err.Error() == "Item not found" {
			ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Unexpected error"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": item})
}

func (c *ItemContoroller) Create(ctx *gin.Context) {
	var input dto.CreateItemInput
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	newItem, err := c.service.Create(input)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{"data": newItem})
}

func (c *ItemContoroller) Update(ctx *gin.Context) {
	itemId, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid id"})
		return
	}
	var input dto.UpdateItemInput
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updatedItem, err := c.service.Update(uint(itemId), input)
	if err != nil {
		if err.Error() == "Item not found" {
			ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Unexpected error"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"date": updatedItem})
}

func (c *ItemContoroller) Delete(ctx *gin.Context) {
	itemId, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid id"})
		return
	}
	err = c.service.Delete(uint(itemId))
	if err != nil {
		if err.Error() == "Item not found" {
			ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Unexpected error"})
		return
	}

	ctx.Status(http.StatusOK)
}
