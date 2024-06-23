// usecase.go
package plant

import (
    "time"
    "fmt"
)

type PlantEarliestWateringService interface {
    FindEarliestWateringTime(plantID string) (PlantEarliestWatering, error)
}

type plantEarliestWateringService struct {
    WateringScheduleRepo PlantEarliestWateringRepository
}

func NewPlantEarliestWateringService(repo PlantEarliestWateringRepository) PlantEarliestWateringService {
    return &plantEarliestWateringService{WateringScheduleRepo: repo}
}

func ConvertToTime(timeStr string) (time.Time, error) {
    layout := "15:04"
    return time.Parse(layout, timeStr)
}

func (s *plantEarliestWateringService) FindEarliestWateringTime(plantID string) (PlantEarliestWatering, error) {
    schedules, err := s.WateringScheduleRepo.GetWateringSchedules(plantID) // Assuming schedules are filtered by plantID
    if err != nil {
        return PlantEarliestWatering{}, err
    }
    if len(schedules) == 0 {
        return PlantEarliestWatering{}, fmt.Errorf("no schedules provided")
    }

    earliestSchedule := schedules[0]
    earliestTime, err := ConvertToTime(earliestSchedule.WateringTime)
    if err != nil {
        return PlantEarliestWatering{}, err
    }

    for _, schedule := range schedules[1:] {
        currentTime, err := ConvertToTime(schedule.WateringTime)
        if err != nil {
            return PlantEarliestWatering{}, err
        }

        if currentTime.Before(earliestTime) {
            earliestTime = currentTime
            earliestSchedule = schedule
        }
    }
    return earliestSchedule, nil
}
