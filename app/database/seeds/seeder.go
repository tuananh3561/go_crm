package seeds

import (
	"github.com/tuananh3561/go_crm/app/config"
	"log"
	"reflect"
)

// Seed type
type Seed struct {
	cf config.Config
}

func Seeder(cf config.Config, seedMethodNames ...string) {
	s := Seed{
		cf: cf,
	}

	// Execute all seeders if no method name is given
	if len(seedMethodNames) == 0 {
		log.Println("Running all seeder...")
		// We are looping over the method on a Seed struct
		seedType := reflect.TypeOf(s)
		for i := 0; i < seedType.NumMethod(); i++ {
			// Get the method in the current iteration
			method := seedType.Method(i)
			// Execute seeder
			Execute(s, method.Name)
		}
	} else {
		// Execute only the given method names
		for _, item := range seedMethodNames {
			Execute(s, item)
		}
	}
}

func Execute(s Seed, seedMethodName string) {
	// Get the reflection value of the method
	m := reflect.ValueOf(s).MethodByName(seedMethodName)
	// Exit if the method doesn't exist
	if !m.IsValid() {
		log.Fatal("No method called ", seedMethodName)
	}
	// Execute the method
	log.Println("Seeding", seedMethodName, "...")
	m.Call(nil)
	log.Println("Seed", seedMethodName, "success")
}
