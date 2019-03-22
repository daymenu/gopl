package main

import (
	"encoding/json"
	"fmt"
)

type Movie struct {
	Title  string
	Year   int  `json:"released"`
	Color  bool `json:"color,omitempty"`
	Actors []string
}

func main() {
	movies := []Movie{
		Movie{"无名之辈", 2018, true, []string{"黄渤", "徐峥"}},
		Movie{"无名之辈2", 2018, false, []string{"黄渤", "徐峥"}},
	}

	moviesJson, err := json.MarshalIndent(movies, "", " ")
	if err != nil {
		fmt.Println(err)
	}

	var newMovies []Movie

	json.Unmarshal(moviesJson, &newMovies)
	fmt.Println(string(moviesJson))
	fmt.Println(newMovies[0].Year)
}
