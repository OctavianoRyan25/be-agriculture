package notification

import "gorm.io/gorm"

type Repository interface {
	StoreNotification(*Notification) (*Notification, error)
	ReadNotification(int) (*Notification, error)
	GetAllNotifications(uint) ([]Notification, error)
	DeleteAllNotifications(uint) error
}

type notificationRepository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *notificationRepository {
	return &notificationRepository{
		db: db,
	}
}

func (r *notificationRepository) StoreNotification(notification *Notification) (*Notification, error) {
	err := r.db.Create(notification).Error
	if err != nil {
		return nil, err
	}

	return notification, nil
}

func (r *notificationRepository) ReadNotification(id int) (*Notification, error) {
	var notification Notification
	err := r.db.Where("id = ?", id).First(&notification).Error
	if err != nil {
		return nil, err
	}
	notification.IsRead = true
	err = r.db.Save(&notification).Error
	if err != nil {
		return nil, err
	}

	return &notification, nil
}

func (r *notificationRepository) GetAllNotifications(userID uint) ([]Notification, error) {
	var notifications []Notification
	err := r.db.Where("user_id = ?", userID).Find(&notifications).Error
	if err != nil {
		return nil, err
	}

	return notifications, nil
}

func (r *notificationRepository) DeleteAllNotifications(userID uint) error {
	err := r.db.Where("user_id = ?", userID).Delete(&Notification{}).Error
	if err != nil {
		return err
	}

	return nil
}
