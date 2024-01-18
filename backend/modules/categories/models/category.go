package models

import (
	"backend/modules/users/models"
	"gorm.io/gorm"
)

type Category struct {
	gorm.Model
	Name   string      `gorm:"not null"`
	UserID uint        `gorm:"not null"`
	User   models.User `gorm:"foreignKey:UserID"`
}

type CategoryModel struct {
	DB *gorm.DB
}

// GetAll returns all categories for a given user ID.
//
// userID: the ID of the user to retrieve categories for.
// []Category: a slice of Category structs representing the categories.
// error: any error that occurred during the retrieval process.
func (m *CategoryModel) GetAll(userID uint) ([]Category, error) {
	var category []Category

	err := m.DB.Where("user_id = ?", userID).Find(&category).Error
	if err != nil {
		return category, err
	}

	return category, nil
}

// Create creates a new category with the given ID and name.
//
// Parameters:
// - id: The ID of the category.
// - name: The name of the category.
//
// Returns:
// - The created Category.
// - An error if there was a problem creating the category.
func (m *CategoryModel) Create(id uint, name string) (Category, error) {
	category := Category{
		Name:   name,
		UserID: id,
	}

	err := m.DB.Create(&category).Error
	if err != nil {
		return Category{}, err
	}

	return category, nil
}

// Update updates the name of a category with the given ID and user ID.
//
// Parameters:
// - id: the ID of the category to update.
// - userId: the ID of the user who owns the category.
// - name: the new name for the category.
//
// Returns:
// - uint: the ID of the category that was updated.
// - error: an error if the update operation fails.
func (m *CategoryModel) Update(id, userId uint, name string) (uint, error) {
	err := m.DB.Model(&Category{}).Where("id = ? AND user_id = ?", id, userId).Update("name", name).Error

	return id, err
}

// Delete deletes a category from the database.
//
// Parameters:
// - id: the ID of the category to delete.
// - userId: the ID of the user who owns the category.
//
// Returns:
// - error: an error if the deletion fails.
func (m *CategoryModel) Delete(id, userId uint) error {
	return m.DB.Where("id = ? AND user_id = ?", id, userId).Delete(&Category{}).Error
}
