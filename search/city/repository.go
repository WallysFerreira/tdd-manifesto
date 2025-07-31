package city

var citiesNames = []string{
	"Paris", "Budapest", "Skopje", "Rotterdam", "Valencia", "Vancouver", "Amsterdam",
	"Vienna", "Sydney", "New York City", "London", "Bangkok", "Hong Kong", "Dubai", "Rome", "Istanbul",
}

func cities() (finalList []City) {
	finalList = []City{}

	for _, cityName := range citiesNames {
		finalList = append(finalList, City{cityName})
	}

	return
}
