package main

import (
	"encoding/csv"
	"log"
	"os"
)

//main function
func main() {
	csvFile, err := os.Open("netflix_titles.csv")
	if err != nil {
		log.Println(err)
	}
	log.Println("Successfully Opened CSV file")
	defer csvFile.Close()

	netflix_titles, err := csv.NewReader(csvFile).ReadAll()
	if err != nil {
		log.Println(err)
	}

	netflix_titles = sort_csv(netflix_titles)
	tv_shows(netflix_titles)
	horror_movies(netflix_titles)
	indian_movies(netflix_titles)
}
