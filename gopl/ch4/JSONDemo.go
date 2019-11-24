package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Movie struct {
	Title  string
	Year   int  `json:"released"`
	Color  bool `json:"color,omitempty"`
	Actors []string
}

func main() {
	var movies = []Movie{
		{Title: "我不是药神", Year: 2018, Color: true,
			Actors: []string{"徐峥", "王传君"}},
		{Title: "地道战", Year: 1965, Color: false,
			Actors: []string{"朱龙广"}},
	}
	//data, err := json.Marshal(movies)
	data, err := json.MarshalIndent(movies, "", " ") // indent: 缩进
	if err != nil{
		log.Fatal("JSON marshaling failed: %s\n", err)
	}
	fmt.Printf("%s\n",data)

	var newMovies []Movie
	if err := json.Unmarshal(data, &newMovies); err!=nil{
		log.Printf("JSON unmarshaling failed: %s\n", err)
	}
	fmt.Println(newMovies)
}
