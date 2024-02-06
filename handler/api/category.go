package api

import (
	"a21hc3NpZ25tZW50/model"
	"a21hc3NpZ25tZW50/service"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CategoryAPI interface {
	AddCategory(c *gin.Context)
	UpdateCategory(c *gin.Context)
	DeleteCategory(c *gin.Context)
	GetCategoryByID(c *gin.Context)
	GetCategoryList(c *gin.Context)
}

type categoryAPI struct {
	categoryService service.CategoryService
}

func NewCategoryAPI(categoryRepo service.CategoryService) *categoryAPI {
	return &categoryAPI{categoryRepo}
}

func (ct *categoryAPI) AddCategory(c *gin.Context) {
	var newCategory model.Category
	if err := c.ShouldBindJSON(&newCategory); err != nil {
		c.JSON(http.StatusBadRequest, model.ErrorResponse{Error: err.Error()})
		return
	}

	err := ct.categoryService.Store(&newCategory)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, model.SuccessResponse{Message: "add category success"})
}

func (ct *categoryAPI) UpdateCategory(c *gin.Context) {
	// TODO: answer here
	categoryId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, model.ErrorResponse{
			Error: "invalid Category ID",
		})
		return
	}
	
	category := model.Category{}
	if err := c.ShouldBindJSON(&category); err != nil {
		c.JSON(http.StatusBadRequest, model.ErrorResponse{
			Error: err.Error(),
		})
		return
	}
	category.ID = categoryId 
	errUpdateCategory := ct.categoryService.Update(category.ID, category)
	if errUpdateCategory != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorResponse {
			Error: "Internal Server Error",
		})
		return
	}
	c.JSON(http.StatusOK, model.SuccessResponse{
		Message: "category update success",
	})
}

func (ct *categoryAPI) DeleteCategory(c *gin.Context) {
	// TODO: answer here
	categoryId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadGateway, model.ErrorResponse{
			Error: "Invalid category ID",
		})
		return
	}

	errorDeleteCategory := ct.categoryService.Delete(categoryId)
	if errorDeleteCategory != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorResponse {
			Error: "Error Delete!...",
		})
		return
	}

	c.JSON(http.StatusOK, model.SuccessResponse{
		Message: "category delete success",
	})
}

func (ct *categoryAPI) GetCategoryByID(c *gin.Context) {
	categoryID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, model.ErrorResponse{Error: "Invalid category ID"})
		return
	}

	category, err := ct.categoryService.GetByID(categoryID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, category)
}

func (ct *categoryAPI) GetCategoryList(c *gin.Context) {
	// TODO: answer here
	categorys, err := ct.categoryService.GetList()
	fmt.Println("=====API Categorys")
	fmt.Println(categorys)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorResponse{
			Error: "Internal Server Error",
		})
		return
	}
	c.JSON(http.StatusOK, categorys)
}
