/*
Copyright © 2022 CARLOS GOMEZ ezlosswm@gmail.com
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// randomCmd represents the random command
var randomCmd = &cobra.Command{
	Use:   "random",
	Short: "Returns a random dad joke.",
	Long:  `Returns a random dad joke.`,
	Run: func(cmd *cobra.Command, args []string) {
		var activity string

		typeFilter, _ := cmd.Flags().GetString("type")
		if typeFilter != "" {
			activity = QueryActivityByType(typeFilter)
		}

		participantFilter, _ := cmd.Flags().GetInt("people")
		if participantFilter != 0 {
			activity = QueryActivityByParticipants(participantFilter)
		}
		activityList, _ := cmd.Flags().GetBool("activities")

		if activityList {
			ActivityType.Print()
		} else {
			fmt.Println(activity)
		}

	},
}

func init() {
	rootCmd.AddCommand(randomCmd)

	randomCmd.PersistentFlags().String("type", "", "filter by type")
	randomCmd.PersistentFlags().Int("people", 1, "filter by number of people")
	randomCmd.PersistentFlags().Bool("activities", false, "prints all accepts activity types")

}
