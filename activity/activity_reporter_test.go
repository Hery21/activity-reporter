package activity_test

import (
	"fmt"
	. " hery-ciaputra/assignment-activity-reporter/activity"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestErrorNotFollowed_Error(t *testing.T) {
	t.Run("Should return error if given John likes Wick's photo but didn't follow Wick", func(t *testing.T) {
		users := NewUserActivity()

		user1 := new(User)
		user1.Name = "John"
		user2 := new(User)
		user2.Name = "Wick"
		user3 := new(User)
		user3.Name = "Cena"

		users.Setup(*user1, *user2)
		users.Setup(*user2, *user3)
		_ = users.Upload(*user3)

		err := users.Like(*user1, *user3)

		assert.ErrorContains(t, err, "User is not followed")
	})
}

func TestErrorUnknownUser_Error(t *testing.T) {
	t.Run("Should return error when Cena is not setup", func(t *testing.T) {
		users := NewUserActivity()

		user := new(User)
		user.Name = "Cena"
		err := users.Upload(*user)
		expected := fmt.Sprintf("Unknown user %v", user.Name)

		assert.ErrorContains(t, err, expected)
	})
}

func TestErrorNoPhoto_Error(t *testing.T) {
	t.Run("Should return error when Cena has not uploaded photo", func(t *testing.T) {
		users := NewUserActivity()

		user1 := new(User)
		user1.Name = "John"
		user2 := new(User)
		user2.Name = "Wick"

		users.Setup(*user2, *user1)
		err := users.Like(*user2, *user1)
		expected := fmt.Sprintf("No photo")

		assert.ErrorContains(t, err, expected)
	})
}

func TestNewUserActivity(t *testing.T) {
	t.Run("Should create new user map when initiated", func(t *testing.T) {
		users := NewUserActivity()

		assert.NotNil(t, users)
	})
}

func TestUserActivity_Setup(t *testing.T) {
	t.Run("Should setup John and Wick when John and Wick has not been set up", func(t *testing.T) {
		users := NewUserActivity()

		user1 := new(User)
		user1.Name = "John"
		user2 := new(User)
		user2.Name = "Wick"

		users.Setup(*user1, *user2)

		actual1 := users.CheckUserExist(*user1)
		actual2 := users.CheckUserExist(*user2)

		assert.True(t, actual1)
		assert.True(t, actual2)
	})

	t.Run("Should setup Cena when John follows Cena given John has been set up but Cena has not been set up", func(t *testing.T) {
		users := NewUserActivity()

		user1 := new(User)
		user1.Name = "John"
		user2 := new(User)
		user2.Name = "Wick"
		user3 := new(User)
		user3.Name = "Cena"

		users.Setup(*user1, *user2)
		users.Setup(*user1, *user3)

		actual := users.CheckUserExist(*user3)

		assert.True(t, actual)
	})

	t.Run("Should setup Cena when Cena follows John given John has been set up but Cena has not been set up", func(t *testing.T) {
		users := NewUserActivity()

		user1 := new(User)
		user1.Name = "John"
		user2 := new(User)
		user2.Name = "Wick"
		user3 := new(User)
		user3.Name = "Cena"

		users.Setup(*user1, *user2)
		users.Setup(*user3, *user1)

		actual := users.CheckUserExist(*user3)

		assert.True(t, actual)
	})
}

func TestUserActivity_checkUserExist(t *testing.T) {
	t.Run("Should return true if user exist and false if not exist", func(t *testing.T) {
		users := NewUserActivity()

		user1 := new(User)
		user1.Name = "John"
		user2 := new(User)
		user2.Name = "Wick"
		user3 := new(User)
		user3.Name = "Cena"

		users.Setup(*user1, *user2)

		actual1 := users.CheckUserExist(*user1)
		actual2 := users.CheckUserExist(*user2)
		actual3 := users.CheckUserExist(*user3)

		assert.True(t, actual1)
		assert.True(t, actual2)
		assert.False(t, actual3)
	})
}
