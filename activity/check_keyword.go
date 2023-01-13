package activity

import "strings"

func CheckKeyword(s string) error {
	breakString := strings.Split(s, " ")
	if len(breakString) == 3 {
		if breakString[1] == "follows" {
			return nil
		}
		if breakString[1] == "uploaded" {
			if breakString[2] == "photo" {
				return nil
			}
		}
	}

	if len(breakString) == 4 {
		if breakString[1] == "likes" {
			if breakString[3] == "photo" {
				return nil
			}
		}
	}
	return ErrorInvalidKeyword{}
}
