package weather

import (
	"encoding/json"
	"os"
	"time"

	"github.com/go-resty/resty/v2"
)

type WeatherService interface {
	GetCurrentWeather(city string) (*Weather, error)
	GetHourlyWeather(city string) ([]HourlyWeather, error)
	GetDailyWeather(city string) ([]DailyWeather, error)
}

type weatherService struct{}

func NewWeatherService() WeatherService {
	return &weatherService{}
}

func (s *weatherService) GetCurrentWeather(city string) (*Weather, error) {
	client := resty.New()
	//Production
	apiKey := os.Getenv("OPENWEATHER_API_KEY")
	resp, err := client.R().
		SetQueryParams(map[string]string{
			"q":     city,
			"appid": apiKey,
			"units": "metric",
		}).
		Get("http://api.openweathermap.org/data/2.5/weather")

	if err != nil {
		return nil, err
	}

	var apiResponse map[string]interface{}
	if err := json.Unmarshal(resp.Body(), &apiResponse); err != nil {
		return nil, err
	}

	main := apiResponse["main"].(map[string]interface{})
	weatherDesc := apiResponse["weather"].([]interface{})[0].(map[string]interface{})
	wind := apiResponse["wind"].(map[string]interface{})
	sys := apiResponse["sys"].(map[string]interface{})

	weather := &Weather{
		ID:          1,
		City:        city,
		Temperature: main["temp"].(float64),
		RealFeel:    main["feels_like"].(float64),
		Pressure:    int(main["pressure"].(float64)),
		Humidity:    int(main["humidity"].(float64)),
		WindSpeed:   wind["speed"].(float64),
		Main:        weatherDesc["main"].(string),
		Description: weatherDesc["description"].(string),
		Icon:        weatherDesc["icon"].(string),
		Sunrise:     int64(sys["sunrise"].(float64)),
		CreatedAt:   time.Now(),
	}

	return weather, nil
}

func (s *weatherService) GetHourlyWeather(city string) ([]HourlyWeather, error) {
	client := resty.New()
	//Production
	apiKey := os.Getenv("OPENWEATHER_API_KEY")

	resp, err := client.R().
		SetQueryParams(map[string]string{
			"q":     city,
			"appid": apiKey,
			"units": "metric",
		}).
		Get("http://api.openweathermap.org/data/2.5/forecast")

	if err != nil {
		return nil, err
	}

	var apiResponse map[string]interface{}
	if err := json.Unmarshal(resp.Body(), &apiResponse); err != nil {
		return nil, err
	}

	list := apiResponse["list"].([]interface{})
	var hourlyWeathers []HourlyWeather

	for i, item := range list {
		forecastItem := item.(map[string]interface{})
		main := forecastItem["main"].(map[string]interface{})
		weatherDesc := forecastItem["weather"].([]interface{})[0].(map[string]interface{})
		wind := forecastItem["wind"].(map[string]interface{})

		hourlyWeather := HourlyWeather{
			ID:          uint(i + 1),
			City:        city,
			Timestamp:   time.Unix(int64(forecastItem["dt"].(float64)), 0),
			Temperature: main["temp"].(float64),
			RealFeel:    main["feels_like"].(float64),
			Pressure:    int(main["pressure"].(float64)),
			Humidity:    int(main["humidity"].(float64)),
			WindSpeed:   wind["speed"].(float64),
			Main:        weatherDesc["main"].(string),
			Description: weatherDesc["description"].(string),
			Icon:        weatherDesc["icon"].(string),
		}

		hourlyWeathers = append(hourlyWeathers, hourlyWeather)
	}

	return hourlyWeathers, nil
}

func (s *weatherService) GetDailyWeather(city string) ([]DailyWeather, error) {
	client := resty.New()
	//Production
	apiKey := os.Getenv("OPENWEATHER_API_KEY")
	resp, err := client.R().
		SetQueryParams(map[string]string{
			"q":     city,
			"appid": apiKey,
			"units": "metric",
		}).
		Get("http://api.openweathermap.org/data/2.5/forecast/daily")

	if err != nil {
		return nil, err
	}

	var apiResponse map[string]interface{}
	if err := json.Unmarshal(resp.Body(), &apiResponse); err != nil {
		return nil, err
	}

	list := apiResponse["list"].([]interface{})
	var dailyWeathers []DailyWeather

	for i, item := range list {
		forecastItem := item.(map[string]interface{})
		temp := forecastItem["temp"].(map[string]interface{})
		weatherDesc := forecastItem["weather"].([]interface{})[0].(map[string]interface{})

		dailyWeather := DailyWeather{
			ID:          uint(i + 1),
			City:        city,
			Date:        time.Unix(int64(forecastItem["dt"].(float64)), 0),
			Temperature: temp["day"].(float64),
			RealFeel:    temp["eve"].(float64),
			Pressure:    int(forecastItem["pressure"].(float64)),
			Humidity:    int(forecastItem["humidity"].(float64)),
			WindSpeed:   forecastItem["speed"].(float64),
			Main:        weatherDesc["main"].(string),
			Description: weatherDesc["description"].(string),
			Icon:        weatherDesc["icon"].(string),
			Sunrise:     int64(forecastItem["sunrise"].(float64)),
		}

		dailyWeathers = append(dailyWeathers, dailyWeather)
	}

	return dailyWeathers, nil
}
