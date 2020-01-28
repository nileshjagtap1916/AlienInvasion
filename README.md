# AlienInvasion
Alien Invasion in unknown world. Random number of aliens will visit the random cities in the world map. If two aliens will meet in same city, they will destroy the city. If alien goes inside the city which don't have any routes, then alien will be trapped. This way (considering one alien will only move 10000 times) at a some point, all aliens will get trapped or entire world will get destroyed. 

## Installation steps :
1. Goto your `$GOPATH\src\github.com`
2. Clone the repository 
    `git clone https://github.com/nileshjagtap1916/AlienInvasion.git`
3. Goto project directory
    `cd AlienInvasion`
4. Build the project 
    `go build .\main.go`
5. Run the project with command line argument as number of aliens (Number of aliens must be greater than two)
    `.\main {number_of_aliens}`

## Output format : 
```
{city name} has been destroyed by alien {int} and alien {int}
.
.
.
{Remaining world map after all aliens are trapped. If all cities get destroyed by aliens, then no records will shown}
```

## world map settings :
To define your own world map, update the file `util/world_map.txt`

    rules of world_map.txt :
    1. Format : `{cityName} north={cityName} south={cityName} east={cityName} west={cityName}`
    1. Each row start with distinct city name and directions in above format 
    2. city name should not contain special character except `-`