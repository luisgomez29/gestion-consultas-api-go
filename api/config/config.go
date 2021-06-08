package config

import (
	"os"
	"sync"

	_ "github.com/joho/godotenv/autoload"
)

var (
	environment      = map[string]string{}
	environmentMutex = sync.RWMutex{}
)

// Load read the settings from the .env file
func Load(key string) string {
	// Gets the value of the map, if it exists returns the value
	environmentMutex.RLock()
	val := environment[key]
	environmentMutex.RUnlock()

	if environment[key] != "" {
		return val
	}

	// If the value does not exist, it gets from ENV
	val = os.Getenv(key)
	if val == "" || len(val) <= 0 {
		return val
	}

	// If the value exists in ENV it is assigned to the map
	environmentMutex.Lock()
	environment[key] = val
	environmentMutex.Unlock()

	return val
}
