package cmd

import (
	"fmt"
	"log"
	"strconv"
)

func QueryActivity() string {

	return APICall.UnmarshalResponse(URL)

}

func QueryActivityByType(typeOfActivity string) string {

	activity, err := ValidateType(typeOfActivity)
	if err != nil {
		fmt.Println(err)
	}

	url := URL + "?type=" + activity

	return APICall.UnmarshalResponse(url)

}

func QueryActivityByParticipants(participants int) string {

	if participants == 0 {
		log.Fatalf("error: number of participants cannot be 0")
	}

	url := URL + "?participants=" + strconv.Itoa(participants)

	return APICall.UnmarshalResponse(url)

}

// func

func ValidateType(activityTypeInput string) (string, error) {

	if _, exists := ActivityType.Activity[activityTypeInput]; !exists {
		return activityTypeInput, fmt.Errorf("%s not a valid type", activityTypeInput)
	}

	return activityTypeInput, nil

}
