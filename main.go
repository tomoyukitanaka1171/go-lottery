
package main

import (
	"fmt"

	"bufio"
	"os"
	"encoding/json"

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
		scanner.Scan()

		in := scanner.Text()

		switch in {
		case "Exit":
			goto Exit
		default:
			logic.Rotation(facilitatorName)
			fmt.Printf("%[1]d : %[2]s", i, facilitatorName[logic.GetMemory()])
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