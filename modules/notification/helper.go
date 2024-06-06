package notification

import (
	"context"
	"fmt"
	"log"
	"time"

	firebase "firebase.google.com/go/v4"
	"github.com/OctavianoRyan25/be-agriculture/modules/plant"
	"github.com/OctavianoRyan25/be-agriculture/modules/user" // Updated import
	"github.com/robfig/cron/v3"
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

// Send reminder notification and store it in the database
func SendReminder(user user.User, plant plant.Plant, useCase UseCase) error {
	// Simulating FCM messaging part is commented out
	/*
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
	*/

	// Store the notification in the database
	notification := &Notification{
		Title:     "Watering Reminder",
		Body:      fmt.Sprintf("It's time to water your plant: %s", plant.Name),
		UserId:    user.ID,
		IsRead:    false,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	_, err := useCase.StoreNotification(notification)
	if err != nil {
		log.Printf("Error storing notification: %v", err)
		return err
	}
	fmt.Printf("Notification stored for user %s\n", user.Email)
	return nil
}

// Schedule watering reminders based on PlantReminder.WateringTime
func StartScheduler(db *gorm.DB, useCase UseCase) {
	// Define desired location for time zone (Asia/Jakarta)
	location, err := time.LoadLocation("Asia/Jakarta")
	if err != nil {
		fmt.Printf("Error loading location: %v\n", err)
		return
	}

	c := cron.New()
	c.AddFunc("@hourly", func() {
		fmt.Println("Checking for plants to water...")
		var plantsToWater []plant.Plant
		currentTime := time.Now().In(location)
		formattedTime := currentTime.Format("15:04")

		// Fetch all plants that need watering at the current time
		err := db.
			Preload("WateringSchedule").
			Joins("JOIN plant_reminders ON plant_reminders.plant_id = plants.id").
			Where("plant_reminders.watering_time = ?", formattedTime).
			Find(&plantsToWater).Error
		if err != nil {
			fmt.Printf("Failed to fetch plants to water: %v\n", err)
			return
		}

		// Check if any plants need watering
		if len(plantsToWater) == 0 {
			fmt.Println("No plants found for watering at this time.")
			return
		}

		// Iterate over each plant to water
		for _, plantToWater := range plantsToWater {
			var usersWithPlant []user.User

			// Find users who have this plant
			err := db.Model(&user.User{}).
				Joins("JOIN user_plants ON users.id = user_plants.user_id").
				Where("user_plants.plant_id = ?", plantToWater.ID).
				Find(&usersWithPlant).Error
			if err != nil {
				fmt.Printf("Failed to fetch users with plant %s: %v\n", plantToWater.Name, err)
				continue
			}

			// Notify each user
			for _, user := range usersWithPlant {
				err := SendReminder(user, plantToWater, useCase)
				if err != nil {
					fmt.Printf("Error sending reminder to user %s: %v\n", user.Email, err)
				}
			}
		}
	})
	c.Start()
}
