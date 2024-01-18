package actions

import (
	"backend/modules/categories/services"
	"backend/modules/users/models"
	services2 "backend/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

type CreateOrUpdateCategoryRequest struct {
	Name string `json:"name"`
}

type CategoryRequestAndResponse struct {
	ID uint `json:"id" uri:"id" binding:"required"`
}

// GetCategories retrieves the categories using the given gin.Context.
//
// It expects a *gin.Context parameter and returns nothing.
// @Summary Get list of categories
// @Description Retrieves the categories for the logged-in user
// @Tags Categories
// @Accept  json
// @Produce  json
// @Security BearerAuth
// @Success 200 {array} services.Category
// @Failure 500 {object} services2.ErrorResponse
// @Router /category/all [get]
func GetCategories(c *gin.Context) {
	categoryService, user := getServiceAndUser(c)

	categories, err := categoryService.CreateCategories(user.User.ID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, services2.ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, categories)
}

// CreateCategory handles the creation of a new category.
//
// It expects a gin.Context parameter to access the HTTP request and response.
// It does not have any return values.
// @Summary Create a new category
// @Description Creates a new category with the provided name
// @Tags Categories
// @Accept  json
// @Produce  json
// @Security BearerAuth
// @Param   name     body    CreateOrUpdateCategoryRequest     true        "Category name"
// @Success 200 {object} services.Category
// @Failure 400 {object} services2.ErrorResponse
// @Failure 500 {object} services2.ErrorResponse
// @Router /category/create [post]
func CreateCategory(c *gin.Context) {
	var request CreateOrUpdateCategoryRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, services2.ErrorResponse{Error: err.Error()})
		return
	}

	categoryService, user := getServiceAndUser(c)

	category, err := categoryService.CreateCategory(user.User.ID, request.Name)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, services2.ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, category)
}

// UpdateCategory updates a category.
//
// It takes a gin context as input and performs the following steps:
// 1. Binds the URI parameters to the CategoryRequestAndResponse struct.
// 2. If there is an error binding the URI parameters, it aborts the request with a bad request status and returns the error.
// 3. Binds the JSON body to the CreateOrUpdateCategoryRequest struct.
// 4. If there is an error binding the JSON body, it aborts the request with a bad request status and returns the error.
// 5. Gets the categoryService and user from the gin context.
// 6. Calls the UpdateCategory method of the categoryService with the ID, userID, and name from the request.
// 7. If there is an error updating the category, it aborts the request with an internal server error status and returns the error.
// 8. Returns the updated category ID in the response body.
// @Summary Update category
// @Description Updates the category with the given ID
// @Tags Categories
// @Accept  json
// @Produce  json
// @Security BearerAuth
// @Param   id       path     int                             true        "Category ID"
// @Param   name     body    CreateOrUpdateCategoryRequest     true        "Category name"
// @Success 200 {object} CategoryRequestAndResponse
// @Failure 400 {object} services2.ErrorResponse
// @Failure 500 {object} services2.ErrorResponse
// @Router /category/update/{id} [put]
func UpdateCategory(c *gin.Context) {
	var request CategoryRequestAndResponse
	var json CreateOrUpdateCategoryRequest

	if err := c.ShouldBindUri(&request); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, services2.ErrorResponse{Error: err.Error()})
		return
	}

	if err := c.ShouldBindJSON(&json); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, services2.ErrorResponse{Error: err.Error()})
		return
	}

	categoryService, user := getServiceAndUser(c)

	id, err := categoryService.UpdateCategory(request.ID, user.User.ID, json.Name)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, services2.ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, CategoryRequestAndResponse{ID: id})
}

// DeleteCategory deletes a category.
//
// The function takes a gin.Context pointer as a parameter.
// It returns nothing.
// @Summary Delete category
// @Description Deletes the category with the given ID
// @Tags Categories
// @Accept  json
// @Produce  json
// @Security BearerAuth
// @Param   id       path     int                             true        "Category ID"
// @Success 200 {object} CategoryRequestAndResponse
// @Failure 400 {object} services2.ErrorResponse
// @Failure 500 {object} services2.ErrorResponse
// @Router /category/delete/{id} [delete]
func DeleteCategory(c *gin.Context) {
	var request CategoryRequestAndResponse

	if err := c.ShouldBindUri(&request); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, services2.ErrorResponse{Error: err.Error()})
		return
	}

	categoryService, user := getServiceAndUser(c)

	err := categoryService.DeleteCategory(request.ID, user.User.ID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, services2.ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, CategoryRequestAndResponse{ID: request.ID})
}

// getServiceAndUser returns the category service and user token.
//
// It takes a Gin context as a parameter.
// It returns a CategoryService and a Token.
func getServiceAndUser(c *gin.Context) (services.CategoryService, models.Token) {
	service := services.CategoryService{DB: services2.GetDBConnection()}
	user := services2.GetUserFromContext(c)

	return service, user
}
