package notification

import "time"

type UseCase interface {
	StoreNotification(*Notification) (*Notification, error)
	ReadNotification(int) (*Notification, error)
	GetAllNotifications(uint) ([]Notification, error)
	DeleteAllNotifications(uint) error
}

type notificationUseCase struct {
	notificationRepo Repository
}

func NewUseCase(notificationRepo Repository) *notificationUseCase {
	return &notificationUseCase{
		notificationRepo: notificationRepo,
	}
}

func (u *notificationUseCase) StoreNotification(notification *Notification) (*Notification, error) {
	notification.IsRead = false
	notification.CreatedAt = time.Now()
	notification.UpdatedAt = time.Now()
	return u.notificationRepo.StoreNotification(notification)
}

func (u *notificationUseCase) ReadNotification(id int) (*Notification, error) {
	return u.notificationRepo.ReadNotification(id)
}

func (u *notificationUseCase) GetAllNotifications(userID uint) ([]Notification, error) {
	return u.notificationRepo.GetAllNotifications(userID)
}

func (u *notificationUseCase) DeleteAllNotifications(userID uint) error {
	return u.notificationRepo.DeleteAllNotifications(userID)
}
