
package main

import (
	"fmt"

	"bufio"
	"os"
	"encoding/json"
	// "io/ioutil"

	"go-lottery/model"
	"go-lottery/logic"
)

func main() {
  facilitator := loadJson("facilitator.json")

	var facilitatorName []string
	for _, f := range facilitator {
		facilitatorName = append(facilitatorName, f.Name)
	}

	scanner := bufio.NewScanner(os.Stdin)
	
	for i := 0;; i++ {
		fmt.Print("num: %s", i)

		scanner.Scan()

		in := scanner.Text()

		switch in {
		case "Exit":
			goto Exit
		default:
			logic.Rotation(facilitatorName)
			fmt.Print(facilitatorName[logic.GetMemory()])
		}
	}
	Exit:
}


func loadJson(filename string) []model.Facilitator {
	bytes, err := os.ReadFile(filename)

	if err != nil {
		fmt.Print("cannot load json file", err)
	}

	var facilitator []model.Facilitator
	json.Unmarshal(bytes, &facilitator)
	return facilitator
}