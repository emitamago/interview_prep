package main

import (
	"fmt"
	"time"
)

type Alert struct {
	name        string
	query       string
	intervalSec int
	critical    *Thredhold
}

type Thredhold struct {
	value   int
	message string
}

func queryAlerts() []*Alert {
	alerts := []*Alert{
		&Alert{
			name:        "test-alert",
			query:       "test-query",
			intervalSec: 5,
			critical: &Thredhold{
				value:   10,
				message: "critical message",
			},
		},
		&Alert{
			name:        "high-latency",
			query:       "high-latency",
			intervalSec: 1,
			critical: &Thredhold{
				value:   5,
				message: "critical message",
			},
		},
		&Alert{
			name:        "p90-latency",
			query:       "p90-latency",
			intervalSec: 3,
			critical: &Thredhold{
				value:   100,
				message: "critical message",
			},
		},
	}
	return alerts
}

func notify(a *Alert) {
	fmt.Printf("notifying alert! Alert name: %v, Message: %v  ", a.name, a.critical.message)
}

func resolve(a *Alert) {
	fmt.Printf("resolving alert! Alert name: %v", a.name)
}

func query(q string) int {
	alerts := queryAlerts()
	for _, v := range alerts {
		if v.name == q {
			return v.critical.value
		}
	}
	return 0
}

func main() {
	alertingRule := &Alert{
		name:        "test-alert",
		query:       "test-query",
		intervalSec: 1,
		critical: &Thredhold{
			value:   10,
			message: "critical message",
		},
	}
	queryTicker := time.NewTicker(time.Duration(alertingRule.intervalSec) * time.Second)
	i := 0
	for range queryTicker.C {
		value := query(alertingRule.name)
		fmt.Printf("value is %v\n", value)
		i++
		if i > 5 {
			queryTicker.Stop()
			break
		}
	}

}
