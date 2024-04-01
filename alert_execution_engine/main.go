package main

import (
	"fmt"
	"math/rand"
	"time"
)

// TODO: Add the engine code here.
func main() {
	newEngine := &AlertEngine{}
	alertService := NewAlertService()
	state := make(map[string]string)
	newEngine.Service = alertService
	newEngine.State = state

	alerts, err := alertService.GetAlerts()
	if err == nil {
		for _, v := range alerts {
			go newEngine.IntervalQuery(v)
		}
	}
	time.Sleep(time.Duration(30) * time.Second)
}

// Interfaces / structures for the alerts service.

// Alert is the alert structure.
type Alert struct {
	Name         string
	Query        string
	IntervalSecs int
	Critical     Threshold
}

// Threshold is the structure for a particular threshold.
type Threshold struct {
	Value   float64
	Message string
}

type AlertService interface {
	GetAlerts() ([]*Alert, error)
	ExecuteQuery(query string) (float64, error)
	Notify(alertName, thresholdMessage string) error
	Resolve(alertName string) error
}

// This is a fake implementation of an alerts service to help
// test the program. You are free to change the fake outputs to
// exercise different parts of the code.
type alertService struct{}

// Alertengine struct hold alerts service and statei(if anything is firing) of the alerts service
type AlertEngine struct {
	Service AlertService
	State   map[string]string
}

// By returning alert service struct as AlertService, alertService can call AlertService interface
func NewAlertService() AlertService {
	return &alertService{}
}

func (a *alertService) GetAlerts() ([]*Alert, error) {
	return []*Alert{
		{
			Name:         "test-alert",
			Query:        "test-query",
			IntervalSecs: 2,
			Critical: Threshold{
				Value:   10,
				Message: "critical-alert",
			},
		},
	}, nil
}

func (a *alertService) ExecuteQuery(_ string) (float64, error) {
	max := 11
	min := 5
	rand.Seed(time.Now().UnixNano())
	result := rand.Intn(max-min) + min
	fmt.Printf("ramdon number is %v\n", result)
	return float64(result), nil
}

func (a *alertService) Notify(alertName, thresholdMessage string) error {
	log("notifying alert %s -- %s", alertName, thresholdMessage)
	return nil
}

func (a *alertService) Resolve(alertName string) error {
	log("resolving alert %s", alertName)
	return nil
}

func log(s string, args ...interface{}) {
	fmt.Printf("[%v] %s\n", time.Now().Format("03:04:05"), fmt.Sprintf(s, args...))
}

// Make query in interval
func (e *AlertEngine) IntervalQuery(a *Alert) {
	queryTicker := time.NewTicker(time.Duration(a.IntervalSecs) * time.Second)
	for range queryTicker.C {
		result, err := e.Service.ExecuteQuery(a.Query)
		if err == nil {
			if result >= a.Critical.Value {
				e.State[a.Name] = "firing"
				e.Service.Notify(a.Name, a.Critical.Message)
			} else {
				_, ok := e.State[a.Name]
				if ok {
					e.Service.Resolve(a.Name)
					delete(e.State, a.Name)
				}

			}

		}
	}
}
