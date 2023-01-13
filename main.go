package main

import (
	"bufio"
	"fmt"
	. " hery-ciaputra/assignment-activity-reporter/activity"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var choice int
	users := NewUserActivity()

	for {
		fmt.Println("Activity Reporter\n" +
			"1. Setup\n" +
			"2. Action\n" +
			"3. Display\n" +
			"4. Trending\n" +
			"5. Exit")

		fmt.Print("Enter menu: ")
		_, errorInput := fmt.Scanln(&choice)

		if errorInput != nil {
			fmt.Println("Invalid input")
			fmt.Println("—------------------------------------------------------")
			continue
		}

		switch {
		case choice == 1:
			var setupString string
			fmt.Print("Setup social graph: ")
			if scanner.Scan() {
				setupString = scanner.Text()
			}

			err := CheckKeyword(setupString)

			if err == nil {
				initUser := new(User)
				initUser.Name = strings.Split(setupString, " ")[0]
				targetUser := new(User)
				targetUser.Name = strings.Split(setupString, " ")[2]

				users.Setup(*initUser, *targetUser)
			} else {
				fmt.Println(err)
			}
			fmt.Println("—------------------------------------------------------")

		case choice == 2:
			var actionString string
			fmt.Print("Enter user Actions: ")
			if scanner.Scan() {
				actionString = scanner.Text()
			}

			err := CheckKeyword(actionString)

			if err == nil {
				if strings.Split(actionString, " ")[1] == "likes" {
					initUser := new(User)
					initUser.Name = strings.Split(actionString, " ")[0]
					targetUser := new(User)
					targetUser.Name = strings.Split(actionString, " ")[2]

					err2 := users.Like(*initUser, *targetUser)
					if err2 != nil {
						fmt.Println(err2)
					}
				}
				if strings.Split(actionString, " ")[1] == "uploaded" {
					initUser := new(User)
					initUser.Name = strings.Split(actionString, " ")[0]

					err2 := users.Upload(*initUser)
					if err2 != nil {
						fmt.Println(err2)
					}
				}
			} else {
				fmt.Println(err)
			}
			fmt.Println("—------------------------------------------------------")

		case choice == 3:
			var displayString string
			fmt.Print("Display activity for: ")
			if scanner.Scan() {
				displayString = scanner.Text()
			}

			initUser := new(User)
			initUser.Name = strings.Split(displayString, " ")[0]

			displayed, err := users.Display(*initUser)

			if err != nil {
				fmt.Println(err)
				continue
			}
			fmt.Println(*displayed)

			fmt.Println("—------------------------------------------------------")

		case choice == 4:
			trend := users.Trending()
			fmt.Println(trend)
			fmt.Println("—------------------------------------------------------")

		case choice == 5:
			fmt.Println("Good bye!")
			fmt.Println("—------------------------------------------------------")
			return

		default:
			fmt.Println("Invalid menu!")
			fmt.Println("—------------------------------------------------------")
		}
	}
}
