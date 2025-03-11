package main

import (
	"bufio"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"os"
)

type City struct {
	Name string `json:"name"`
}

type Country struct {
	CountryName string   `json:"country_name"`
	Cities      []string `json:"cities"`
	CityCount   int      `json:"city_count"`
}

func main() {
	inputFile := "cities.csv"
	outputFile := "output.json"

	file, err := os.Open(inputFile)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	reader := csv.NewReader(bufio.NewReader(file))
	reader.Read() // Skip header

	countries := make(map[string]*Country)

	for {
		record, err := reader.Read()
		if err != nil {
			break
		}

		countryName := record[7] // country_name
		cityName := record[1]    // city name

		if _, exists := countries[countryName]; !exists {
			countries[countryName] = &Country{CountryName: countryName, Cities: []string{}}
		}

		countries[countryName].Cities = append(countries[countryName].Cities, cityName)
		countries[countryName].CityCount++
	}

	output, err := os.Create(outputFile)
	if err != nil {
		fmt.Println("Error creating output file:", err)
		return
	}
	defer output.Close()

	encoder := json.NewEncoder(output)
	encoder.SetIndent("", "  ")

	if err := encoder.Encode(countries); err != nil {
		fmt.Println("Error writing JSON:", err)
	}
}
