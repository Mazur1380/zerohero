package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

const reqestURL = "https://api.open-meteo.com/v1/forecast?latitude=59.9386&longitude=30.3141&hourly=temperature_2m&timezone=Europe%2FMoscow&forecast_days=2"

type Time time.Time

func (t *Time) UnmarshalJSON(data []byte) error {
	layout := "\"2006-01-02T15:04\""
	v, err := time.Parse(layout, string(data))
	if err != nil {
		return err
	}
	*t = Time(v)
	return nil

}
func (t Time) String() string {
	v := time.Time(t)
	return v.Format("2006-Junary-02 15:04")
}

type Forecast struct {
	Hourly ForecastHourly `json:"hourly"`
}

type ForecastHourly struct {
	Time []Time    `json:"time"`
	Temp []float64 `json:"temperature_2m"`
}

func main() {
	w, err := getWeather()
	if err != nil {
		log.Fatal(err)
	}

	for i := range w.Hourly.Time {
		fmt.Println(w.Hourly.Time[i], " - ", w.Hourly.Temp[i], "°C")
	}

}

func getWeather() (Forecast, error) {
	cl := http.Client{}

	req, err := http.NewRequest("GET", reqestURL, nil)
	if err != nil {
		return Forecast{}, err
	}

	res, err := cl.Do(req)
	if err != nil {
		return Forecast{}, err
	}

	defer res.Body.Close()

	b, err := io.ReadAll(res.Body)
	if err != nil {
		return Forecast{}, err
	}

	f := Forecast{}
	err = json.Unmarshal(b, &f)
	if err != nil {
		return Forecast{}, err
	}

	return f, nil
}
