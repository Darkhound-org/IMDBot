/*
Copyright © 2022 Drakhound-org

*/
package cmd

import (
	"fmt"
	"io/ioutil"

	"github.com/op/go-logging"
	"github.com/savioxavier/termlink"
	"github.com/spf13/cobra"
)

// startCmd represents the start command
var testCmd = &cobra.Command{
	Use:   "test",
	Short: "Test run",
	Long: `By giving this command you can check if the bot is having any bug or not.
	Also a test run of the api would be done [movie: Interstellar ]
	If the bot is okay '!!..IMDBot bot is ready..!!' message will be displayed`,
	Run: func(cmd *cobra.Command, args []string) {
		var log = logging.MustGetLogger("data_log")

		data, err := ioutil.ReadFile("data")
		if err != nil {
			log.Panicf("Failed reading api key from data file: %s", err)
		}
		dt_x := string(data)
		url_base := "http://www.omdbapi.com/?t=Interstellar" + "&apikey=" + dt_x
		fmt.Println(termlink.Link("Response url", url_base))
		if err != nil {
			log.Fatal("Failed connecting to IMDB api..!!", err)
			return
		}
		log.Info("!!..Open response url for detailed description of the movie or series..!! ")
		log.Debugf("!!..No bugs found. Successfully connected to IMDB api..!!")
		log.Info("[ !!..IMDBot is okay..!! ]")
	},
}

func init() {
	rootCmd.AddCommand(testCmd)
}
