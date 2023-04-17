package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/jutionck/go-sinar-makmur-mongodb/model"
	"github.com/jutionck/go-sinar-makmur-mongodb/usecase"
	"net/http"
)

type BrandController struct {
	router  *gin.Engine
	useCase usecase.BrandUseCase
}

func (b *BrandController) createHandler(c *gin.Context) {
	var payload model.Brand
	err := c.ShouldBind(&payload)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	err = b.useCase.RegisterNewBrand(&payload)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"code":        http.StatusCreated,
		"description": "OK",
		"data":        payload,
	})
}

func (b *BrandController) listHandler(c *gin.Context) {
	brands, err := b.useCase.FindAllBrand()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":        http.StatusOK,
		"description": "OK",
		"data":        brands,
	})
}

func (b *BrandController) getHandler(c *gin.Context) {
	id := c.Param("id")
	brand, err := b.useCase.FindById(id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":        http.StatusOK,
		"description": "OK",
		"data":        brand,
	})
}

func NewBrandController(r *gin.Engine, useCase usecase.BrandUseCase) *BrandController {
	controller := &BrandController{
		router:  r,
		useCase: useCase,
	}
	r.POST("/brands", controller.createHandler)
	r.GET("/brands", controller.listHandler)
	r.GET("/brands/:id", controller.getHandler)
	return controller
}
