package main

import (
	"fmt"
	"log"
	"os"
	"runtime"
	"strconv"
	"sync"
	"time"

	u "github.com/AlienInvasion/util"
)

var utilObj *u.Util

var wg sync.WaitGroup

func main() {
	// Read input from command line
	NumberOfAlians, err := strconv.Atoi(os.Args[1])
	if err != nil {
		log.Fatal(err)
		//os.Exit(0)
	}
	if NumberOfAlians < 2 {
		log.Fatal("Number of alians must be greater than or equal to two")
	}

	// Initiate Util (generate WorldMap)
	utilObj = new(u.Util)
	utilObj.Init()

	wg.Add(NumberOfAlians)
	for i := 1; i <= NumberOfAlians; i++ {
		go alienVisit(i)
	}
	wg.Wait()

	// print remaining WorldMap
	utilObj.PrintWorldMap()

	// trigger garbage collector
	runtime.GC()
}

func alienVisit(alien int) {

	// Get lock on WorldMap for concurent map read-write
	utilObj.WorldMapMutex.Lock()

	// random city from WorldMap
	city := utilObj.GetRandomCity()

	for i := 0; i <= 10000; i++ {

		// If current city is null, alien is trapped.
		if currentCity, ok := utilObj.WorldMap[city]; ok {

			// destroy city condition
			if currentCity.VisitedAt == time.Now() && currentCity.VisitedBy != alien {

				fmt.Printf("%s has been destroyed by alien %d and alien %d!\n", city, alien, currentCity.VisitedBy)

				// Get random direction from city
				nextCity := utilObj.GetRandomDirection(city)

				// Delete city from WorldMap
				delete(utilObj.WorldMap, city)

				// visit next city
				city = nextCity
				continue
			}

			// Update city VisitedBy and VisitedAt
			currentCity.VisitedBy = alien
			currentCity.VisitedAt = time.Now()

			// Visit next city
			city = utilObj.GetRandomDirection(city)

		} else {
			break
		}
	}

	// Unlock WorldMap for concurent map read-write
	utilObj.WorldMapMutex.Unlock()

	// Done concurent goroutine
	wg.Done()
}
