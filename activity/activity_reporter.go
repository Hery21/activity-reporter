package activity

import (
	"fmt"
)

type User struct {
	Name string
}

type Notifications struct {
	Following *[]User
	Upload    *[]User
	Likes     *[][]User
}

type Action struct {
	Upload   bool
	LikeBy   *[]User
	Likes    *[]User
	Follower *[]User
	Notify   *Notifications
}

type UserActivity struct {
	Activity map[User]*Action
}

type ErrorNotFollowed struct{}

type ErrorUnknownUser struct {
	user *User
}

type ErrorInvalidKeyword struct{}

type ErrorNoPhoto struct{}

func (e ErrorNotFollowed) Error() string {
	return "User is not followed"
}

func (e ErrorUnknownUser) Error() string {
	errorString := fmt.Sprintf("Unknown user %v", e.user.Name)
	return errorString
}

func (e ErrorInvalidKeyword) Error() string {
	return "Invalid keyword"
}

func (e ErrorNoPhoto) Error() string {
	return "No photo"
}

func NewUserActivity() *UserActivity {
	return &UserActivity{map[User]*Action{}}
}

func (u *UserActivity) CheckUserExist(user User) bool {
	if _, ok := u.Activity[user]; !ok {
		return false
	}
	return true
}

func (u *UserActivity) Setup(initUser, targetUser User) {
	if u.CheckUserExist(initUser) == false {
		notifications := &Notifications{&[]User{targetUser, initUser}, &[]User{}, &[][]User{}}
		u.Activity[initUser] = &Action{false, &[]User{}, &[]User{}, &[]User{initUser}, notifications}
	} else {
		*u.Activity[initUser].Notify.Following = append(*u.Activity[initUser].Notify.Following, targetUser)
	}

	if u.CheckUserExist(targetUser) == false {
		notifications := &Notifications{&[]User{targetUser}, &[]User{}, &[][]User{}}
		u.Activity[targetUser] = &Action{false, &[]User{}, &[]User{}, &[]User{initUser, targetUser}, notifications}
	} else {
		*u.Activity[targetUser].Follower = append(*u.Activity[targetUser].Follower, initUser)
	}
}
