package myboxsters

import (
	"fmt"
	"log/slog"
	"sync"
)

func Main() int {
	slog.Debug("myboxsters", "test", true)

	main()
	return 0
}

type Singleton struct {
	Data string
}

var (
	instance *Singleton
	once     sync.Once
)

func GetInstance() *Singleton {
	once.Do(func() {
		instance = &Singleton{Data: "Singleton Instance Created"}
	})

	return instance
}

func main() {
	// Get the singleton instance
	singleton := GetInstance()

	// Access and modify the data of the singleton
	fmt.Println("Original Data:", singleton.Data)
	singleton.Data = "Updated Singleton Data"
	fmt.Println("Modified Data:", singleton.Data)

	// Create another reference to the singleton
	anotherReference := GetInstance()

	// Check if both references point to the same instance
	if singleton == anotherReference {
		fmt.Println("Both references point to the same Singleton instance.")
	} else {
		fmt.Println("Both references do not point to the same Singleton instance.")
	}
}
