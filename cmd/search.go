/*
Copyright © 2022 Darkhound-org

*/
package cmd

import (
	"fmt"

	"io/ioutil"

	"os"
	"time"

	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/op/go-logging"
	"github.com/savioxavier/termlink"
	"github.com/spf13/cobra"
)

var searchCmd = &cobra.Command{
	Use:   "search",
	Short: "Searches your favourite movie or TV series",
	Long:  `Search your favourite movie or TV series by giving this command`,
	Run: func(cmd *cobra.Command, args []string) {
		var mo_name string
		var id_bd string
		var op int
		t := table.NewWriter()
		t.SetOutputMirror(os.Stdout)
		t.AppendHeader(table.Row{"#", "COMMANDS"})
		t.AppendRows([]table.Row{
			{1, "Search by name"},
			{2, "Search by IMDB id"},
			{3, "Exit"},
		})
		t.Render()

		fmt.Print("Enter your choice [1,2 or 3] : ")
		fmt.Scanln(&op)
		switch op {
		case 1:
			fmt.Print("Name of movie: ")
			fmt.Scanln(&mo_name)
			var log = logging.MustGetLogger("data_log")
			data, err := ioutil.ReadFile("data")
			if err != nil {
				log.Panicf("Failed reading api key from data file: %s", err)
			}
			dt_x := string(data)
			url_base := "http://www.omdbapi.com/?t=" + mo_name + "&apikey=" + dt_x
			if err != nil {
				log.Fatal("Failed connecting to api..!!Something went wrong..!!", err)
				return
			}
			f, err := os.Create("movie_log.txt")

			if err != nil {
				log.Fatal("Failed creating log file..!!", err)
			}

			defer f.Close()

			_, err2 := f.WriteString(url_base)

			if err2 != nil {
				log.Fatal("Failed writing to movie_log.txt..!!", err2)
			}

			fmt.Println(termlink.Link("Response url", url_base))
			log.Info("!!..Open response url for detailed description of the movie or series..!!")

			log.Info("Output saved to movie_log.txt..!!")
			return

		case 2:
			// Search by id
			fmt.Print("Enter IMDB id: ")
			fmt.Scanln(&id_bd)
			var log = logging.MustGetLogger("data_log")
			data, err := ioutil.ReadFile("data")
			if err != nil {
				log.Panicf("Failed reading api key from data file: %s", err)
			}
			dt_n := string(data)
			url_base := "http://www.omdbapi.com/?i=" + id_bd + "&apikey=" + dt_n
			if err != nil {
				log.Fatal("Failed connecting to api..!!Something went wrong..!!", err)
				return
			}
			f, err := os.Create("movie_log.txt")

			if err != nil {
				log.Fatal("Failed creating log file..!!", err)
			}

			defer f.Close()

			_, err2 := f.WriteString(url_base)

			if err2 != nil {
				log.Fatal("Failed writing to movie_log.txt..!!", err2)
			}

			fmt.Println(termlink.Link("Response url", url_base))
			log.Info("!!..Open response url for detailed description of the movie or series..!!")

			log.Info("Output saved to movie_log.txt..!!")

		case 3:
			fmt.Println("Exiting...")
			time.Sleep(3 * time.Second)
		}

	},
}

func init() {
	rootCmd.AddCommand(searchCmd)

}
