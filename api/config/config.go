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

// Load lee las configuraciones del archivo .env
func Load(key string) string {
	//	obtiene el valor del map, si existe regresa el valor
	environmentMutex.RLock()
	val := environment[key]
	environmentMutex.RUnlock()

	if environment[key] != "" {
		return val
	}

	// si el valor no existe, lo obtiene de ENV
	val = os.Getenv(key)
	if val == "" || len(val) <= 0 {
		return val
	}

	// si existe en ENV, lo asigna al map
	environmentMutex.Lock()
	environment[key] = val
	environmentMutex.Unlock()

	return val
}
