// package main

// import (
// 	"fmt"
// 	"slices"
// )

// type LogEntry struct {
// 	timestamp string
// 	severity  string
// 	message   string
// }

// // Definindo uma ordem de severidade
// var severityOrder = map[string]int{
// 	"INFO":  1,
// 	"WARN":  2,
// 	"ERROR": 3,
// }

// func main() {
// 	logs := []LogEntry{
// 		{"2024-01-10 14:01", "ERROR", "Failed to load module"},
// 		{"2024-01-10 14:00", "INFO", "System started"},
// 		{"2024-01-10 14:02", "WARN", "Memory usage high"},
// 	}

// 	severityCmp := func(a, b LogEntry) int {
// 		return severityOrder[a.severity] - severityOrder[b.severity]
// 	}

//		slices.SortFunc(logs, severityCmp)
//		fmt.Println("Logs sorted by severity:", logs)
//	}
//
// ###################################################
// package main

// import (
// 	"fmt"
// 	"slices"
// 	"time"
// )

// type Activity struct {
// 	name     string
// 	duration time.Duration
// }

// func main() {
// 	activities := []Activity{
// 		{"Read email", 15 * time.Minute},
// 		{"Meeting", 2 * time.Hour},
// 		{"Code review", 45 * time.Minute},
// 	}

// 	durationCmp := func(a, b Activity) int {
// 		if a.duration < b.duration {
// 			return -1
// 		} else if a.duration > b.duration {
// 			return 1
// 		}
// 		return 0
// 	}

//		slices.SortFunc(activities, durationCmp)
//		fmt.Println("Activities sorted by duration:", activities)
//	}
//
// ###########################################################
package main

import (
	"fmt"
	"slices"
)

type MenuItem struct {
	name     string
	calories int
}

func main() {
	menu := []MenuItem{
		{"Salad", 200},
		{"Cheeseburger", 700},
		{"Smoothie", 300},
	}

	caloriesCmp := func(a, b MenuItem) int {
		return a.calories - b.calories
	}

	slices.SortFunc(menu, caloriesCmp)
	fmt.Println("Menu items sorted by calories:", menu)
}
