package activity

import "fmt"

func (u *UserActivity) Display(user User) (*string, error) {
	if u.CheckUserExist(user) == false {
		return nil, ErrorUnknownUser{&user}
	}

	display := fmt.Sprintf("%v activities:\n", user.Name)

	if u.Activity[user].Upload {
		display += "You uploaded a photo\n"
		for _, val := range *u.Activity[user].LikeBy {
			if val == user {
				display += fmt.Sprintf("You liked your photo\n")
			} else {
				display += fmt.Sprintf("%v liked your photo\n", val.Name)
			}
		}
	}

	for _, val := range *u.Activity[user].Notify.Upload {
		if val != user {
			display += fmt.Sprintf("%v uploaded a photo\n", val.Name)
		}
	}

	for _, val := range *u.Activity[user].Likes {
		if val != user {
			display += fmt.Sprintf("You liked %v's photo\n", val.Name)
		}
	}

	for _, val := range *u.Activity[user].Notify.Likes {
		if val[0].Name != user.Name && val[1].Name != user.Name {
			display += fmt.Sprintf("%v liked %v's photo\n", val[0].Name, val[1].Name)
		}
	}

	return &display, nil
}
