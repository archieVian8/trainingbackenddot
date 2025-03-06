package db

import (
	"trainingbackenddot/domain"

	"gorm.io/gorm"
)

type NotificationRepository struct {
	DB *gorm.DB
}

func NewNotificationRepository(db *gorm.DB) *NotificationRepository {
	return &NotificationRepository{DB: db}
}

func (r *NotificationRepository) CreateNotification(notification *domain.Notification) error {
	return r.DB.Create(notification).Error
}

func (r *NotificationRepository) GetAllNotifications() ([]domain.Notification, error) {
	var notifications []domain.Notification
	err := r.DB.Find(&notifications).Error
	return notifications, err
}
