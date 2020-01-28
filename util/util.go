package util

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"
)

// City struct
type City struct {
	Name      string
	South     *City
	North     *City
	East      *City
	West      *City
	VisitedBy int
	VisitedAt time.Time
}

// Util struct
type Util struct {
	WorldMap      map[string]*City
	Directions    []string
	WorldMapMutex sync.RWMutex
}

// Init : constructor
func (u *Util) Init() {
	u.Directions = []string{"North", "South", "East", "West"}
	u.WorldMapMutex = sync.RWMutex{}
	u.GetWorldMap()

	return
}

// GetRandomCity : select available random city
func (u *Util) GetRandomCity() string {
	for k := range u.WorldMap {
		return k
	}
	return ""
}

// GetRandomDirection : select available random direction
func (u *Util) GetRandomDirection(city string) string {
	if u.WorldMap[city].North != nil {
		if u.WorldMap[u.WorldMap[city].North.Name] != nil {
			return u.WorldMap[city].North.Name
		}
	}
	if u.WorldMap[city].South != nil {
		if u.WorldMap[u.WorldMap[city].South.Name] != nil {
			return u.WorldMap[city].South.Name
		}
	}
	if u.WorldMap[city].East != nil {
		if u.WorldMap[u.WorldMap[city].East.Name] != nil {
			return u.WorldMap[city].East.Name
		}
	}
	if u.WorldMap[city].West != nil {
		if u.WorldMap[u.WorldMap[city].West.Name] != nil {
			return u.WorldMap[city].West.Name
		}
	}
	return ""
}

// PrintWorldMap : print world map
func (u *Util) PrintWorldMap() {
	for city := range u.WorldMap {
		var printString = city
		if u.WorldMap[city].North != nil {
			if u.WorldMap[u.WorldMap[city].North.Name] != nil {
				printString = printString + " north=" + u.WorldMap[city].North.Name
			}
		}
		if u.WorldMap[city].South != nil {
			if u.WorldMap[u.WorldMap[city].South.Name] != nil {
				printString = printString + " south=" + u.WorldMap[city].South.Name
			}
		}
		if u.WorldMap[city].East != nil {
			if u.WorldMap[u.WorldMap[city].East.Name] != nil {
				printString = printString + " east=" + u.WorldMap[city].East.Name
			}
		}
		if u.WorldMap[city].West != nil {
			if u.WorldMap[u.WorldMap[city].West.Name] != nil {
				printString = printString + " west=" + u.WorldMap[city].West.Name
			}
		}
		fmt.Println(printString)
	}
}

// GetWorldMap : read world_map.txt and create WorldMap object
func (u *Util) GetWorldMap() {
	u.WorldMap = make(map[string]*City)
	// read world map file
	absPath, err := filepath.Abs("./util/world_map.txt")
	if err != nil {
		log.Fatal(err)
	}
	file, err := os.Open(absPath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// parse file line by line
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		row := scanner.Text()
		words := strings.Fields(row)

		// if city not present in WorldMap, create new city
		if _, ok := u.WorldMap[words[0]]; !ok {
			city := City{Name: words[0]}
			u.WorldMap[words[0]] = &city
		}

		// assign directions to the city
		for i := 1; i <= len(words)-1; i++ {
			splittedString := strings.Split(words[i], "=")
			direction := splittedString[0]
			directedCityName := splittedString[1]

			// if city not present in WorldMap, create new city
			if _, ok := u.WorldMap[directedCityName]; !ok {
				city := City{Name: directedCityName}
				u.WorldMap[directedCityName] = &city
			}

			// create route between cities
			switch direction {
			case "north":
				u.WorldMap[words[0]].North = u.WorldMap[directedCityName]
			case "south":
				u.WorldMap[words[0]].South = u.WorldMap[directedCityName]
			case "east":
				u.WorldMap[words[0]].East = u.WorldMap[directedCityName]
			case "west":
				u.WorldMap[words[0]].West = u.WorldMap[directedCityName]
			}
		}
	}
	return
}
