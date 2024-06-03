package notification

import (
	"context"
	"fmt"
	"log"
	"time"

	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/messaging"
	"github.com/OctavianoRyan25/be-agriculture/modules/plant"
	"github.com/OctavianoRyan25/be-agriculture/modules/user"
	"github.com/robfig/cron"
	"google.golang.org/api/option"
	"gorm.io/gorm"
)

// Initialize Firebase app
func InitFirebase() *firebase.App {
	opt := option.WithCredentialsFile("path/to/your-firebase-adminsdk.json")
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		log.Fatalf("error initializing app: %v", err)
	}
	return app
}

// Send reminder notification using FCM
func SendReminder(user user.User, plant plant.Plant, client *messaging.Client) {
	message := &messaging.Message{
		Token: user.FCMToken,
		Notification: &messaging.Notification{
			Title: "Watering Reminder",
			Body:  fmt.Sprintf("It's time to water your plant: %s", plant.Name),
		},
	}

	_, err := client.Send(context.Background(), message)
	if err != nil {
		log.Printf("Error sending FCM message: %v", err)
	} else {
		fmt.Printf("Reminder sent to %s for watering plant %s\n", user.Email, plant.Name)
	}
}

// Schedule watering reminders based on PlantReminder.WateringTime
func ScheduleWateringReminders(c *cron.Cron, db *gorm.DB, firebaseApp *firebase.App) {
	var userPlants []plant.UserPlant
	err := db.Preload("Plant.WateringSchedule").Preload("User").Find(&userPlants).Error
	if err != nil {
		log.Fatalf("Failed to fetch user plants: %v", err)
	}

	client, err := firebaseApp.Messaging(context.Background())
	if err != nil {
		log.Fatalf("error getting Messaging client: %v", err)
	}

	for _, userPlant := range userPlants {
		plant := userPlant.Plant
		wateringSchedule := plant.WateringSchedule
		user := userPlant.User

		// Parse the WateringTime to schedule the reminder
		wateringTime := wateringSchedule.WateringTime
		parsedTime, err := time.Parse("15:04", wateringTime)
		if err != nil {
			log.Printf("Invalid watering time format for plant %s: %v", plant.Name, err)
			continue
		}

		// Create cron schedule based on parsed time
		cronSchedule := fmt.Sprintf("%d %d * * *", parsedTime.Minute(), parsedTime.Hour())
		c.AddFunc(cronSchedule, func() {
			SendReminder(user, plant, client)
		})

		fmt.Printf("Scheduled reminder for plant %s at %s\n", plant.Name, cronSchedule)
	}
}
