package cmd

import (
	"testing"
)

func TestValidateType(t *testing.T) {

	v, _ := ValidateType("education")
	_, ok := ActivityType.Activity[v]
	if !ok {
		t.Errorf("Validate Type = %v not a valid type in %v", v, ActivityType.Activity[v])
	}
}
