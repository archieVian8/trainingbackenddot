package usecase

import (
	"fmt"
	"time"
	"trainingbackenddot/domain"
	"trainingbackenddot/infrastructure/db"
)

type NotificationUsecase struct {
	NotificationRepo *db.NotificationRepository
}

func NewNotificationUsecase(notificationRepo *db.NotificationRepository) *NotificationUsecase {
	return &NotificationUsecase{NotificationRepo: notificationRepo}
}

// Function for Send Notification
func (u *NotificationUsecase) SendNotification(message string) {
	go func() {
		time.Sleep(2 * time.Second)

		notification := &domain.Notification{
			Message: message,
		}

		err := u.NotificationRepo.CreateNotification(notification)
		if err != nil {
			fmt.Println("Failed to save notification:", err)
		} else {
			fmt.Println("Notification sent:", message)
		}
	}()
}

func (u *NotificationUsecase) GetNotifications() ([]domain.Notification, error) {
	return u.NotificationRepo.GetAllNotifications()
}
