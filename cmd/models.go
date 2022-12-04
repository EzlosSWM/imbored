package cmd

import (
	"encoding/json"
	"fmt"
	"log"
)

const URL string = "http://www.boredapi.com/api/activity"

var MadeBy = `Made by @EzlosSWM using https://www.boredapi.com/`
var UsageString = "imbored : \n\treturns a random activity to do\n\nimbored type [type of activity] : \n\treturns an activity based on type of activity\n\nimbored people [number of people] : \n\treturns an activity based on the number of entered people\n"
var APICall = BoredAPI{}
var ActivityType = ActivityList{
	map[string]int{
		"recreational": 0,
		"social":       1,
		"diy":          2,
		"charity":      3,
		"cooking":      4,
		"relaxation":   5,
		"music":        6,
		"busywork":     7,
		"education":    8,
	},
}

type ActivityList struct {
	Activity map[string]int
}

type BoredAPI struct {
	Activity      string  `json:"activity"`
	Accessibility float32 `json:"accessibility"`
	Type          string  `json:"type"`
	Participants  int     `json:"participants"`
	Price         float32 `json:"price"`
	Link          string  `json:"link"`
	Key           string  `json:"key"`
}

func (*ActivityList) Print() {
	fmt.Println("Accepted types of activities: ")
	for a := range ActivityType.Activity {
		fmt.Println("\t", a)
	}
}

func (*BoredAPI) UnmarshalResponse(url string) string {
	responseBytes := FetchActivityData(url)

	if err := json.Unmarshal(responseBytes, &APICall); err != nil {
		log.Printf("could not unmarshal response - %v\n", err)
	}

	return string(APICall.Activity)
}
