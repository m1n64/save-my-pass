package services

import (
	"backend/modules/categories/models"
	"gorm.io/gorm"
)

type CategoryService struct {
	DB *gorm.DB
}

type Category struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

// getModel returns a CategoryModel.
//
// No parameters.
// Returns a models.CategoryModel.
func (s *CategoryService) getModel() models.CategoryModel {
	return models.CategoryModel{DB: s.DB}
}

// CreateCategories creates categories for a given user.
//
// It takes in a userId of type uint as a parameter.
// It returns a slice of Category and an error.
func (s *CategoryService) CreateCategories(userId uint) ([]Category, error) {
	categoryModel := s.getModel()

	categories, err := categoryModel.GetAll(userId)

	categoriesList := []Category{}

	for _, category := range categories {
		categoriesList = append(categoriesList, Category{ID: category.ID, Name: category.Name})
	}

	return categoriesList, err
}

// CreateCategory creates a new category for a given user.
//
// Parameters:
// - userId: the ID of the user.
// - name: the name of the category.
//
// Returns:
// - Category: the created category with its ID and name.
// - error: any error that occurred during the creation process.
func (s *CategoryService) CreateCategory(userId uint, name string) (Category, error) {
	categoryModel := s.getModel()

	category, err := categoryModel.Create(userId, name)

	return Category{
		ID:   category.ID,
		Name: category.Name,
	}, err
}

// UpdateCategory updates a category with the given ID, user ID, and name.
//
// Parameters:
// - id: The ID of the category to update.
// - userId: The ID of the user performing the update.
// - name: The new name to assign to the category.
//
// Returns:
// - uint: The ID of the updated category.
// - error: An error if the update operation fails.
func (s *CategoryService) UpdateCategory(id, userId uint, name string) (uint, error) {
	categoryModel := s.getModel()

	return categoryModel.Update(id, userId, name)
}

// DeleteCategory deletes a category by its ID and user ID.
//
// Parameters:
// - id: the ID of the category to be deleted.
// - userId: the ID of the user requesting the deletion.
//
// Return type: error.
func (s *CategoryService) DeleteCategory(id, userId uint) error {
	categoryModel := s.getModel()

	return categoryModel.Delete(id, userId)
}
