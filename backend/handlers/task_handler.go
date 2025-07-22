package handlers

import (
	"errors"
	"github.com/fouradithep/todolist/models"
	"gorm.io/gorm"
	"log"
)

func CreateTask(db *gorm.DB, task *models.Task, userID uint) error {
	task.UserID = userID
	result := db.Create(task)

	if result.Error != nil {
		return result.Error
	}
	return nil
}

func GetTask(db *gorm.DB, id int, userID uint) (*models.Task, error) {
	var task models.Task
	// preload User และเช็คเงื่อนไข user_id ด้วยใน WHERE
    result := db.Preload("User").Where("id = ? AND user_id = ?", id, userID).First(&task)
    if result.Error != nil {
        return nil, result.Error
    }

    return &task, nil
}

func UpdateTask(db *gorm.DB, task *models.Task, userID uint) error {
    var currentTask models.Task
    if err := db.First(&currentTask, "id = ? AND user_id = ?", task.ID, userID).Error; err != nil {
        return errors.New("task not found or unauthorized")
    }
    result := db.Model(&currentTask).Updates(task)
    return result.Error
}

func DeletedTask(db *gorm.DB, id int, userID uint) error {
	var task models.Task
	result := db.First(&task, id)
	if result.Error != nil {
		return result.Error
	}

	if task.UserID != userID {
		return errors.New("unauthorized to delete task")
	}

	return db.Delete(&task).Error
}

func SearchTask(db *gorm.DB, taskStatus string, userID uint) []models.Task {
	var tasks []models.Task

	result := db.Where("status = ? AND user_id = ?", taskStatus, userID).Find(&tasks)
	if result.Error != nil {
		log.Fatalf("Search task failed: %v", result.Error)
	}
	return tasks
}

func GetTasks(db *gorm.DB, userID uint) []models.Task {
	var tasks []models.Task
	db.Preload("User").Where("user_id = ?", userID).Find(&tasks)
    return tasks
}
