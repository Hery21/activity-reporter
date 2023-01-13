package activity

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUserActivity_Display(t *testing.T) {
	t.Run("Should display John's activities John(you) upload a photo", func(t *testing.T) {
		users := NewUserActivity()

		user1 := new(User)
		user1.Name = "John"
		user2 := new(User)
		user2.Name = "Wick"

		users.Setup(*user1, *user2)

		_ = users.Upload(*user1)

		display, _ := users.Display(*user1)

		expected := fmt.Sprintf(
			"%v activities:\n"+
				"You uploaded a photo\n", user1.Name)

		assert.Equal(t, expected, *display)
	})

	t.Run("Should display John's activities John(you) upload a photo and You likes your photo", func(t *testing.T) {
		users := NewUserActivity()

		user1 := new(User)
		user1.Name = "John"
		user2 := new(User)
		user2.Name = "Wick"

		users.Setup(*user1, *user2)

		_ = users.Upload(*user1)
		_ = users.Like(*user1, *user1)

		display, _ := users.Display(*user1)

		expected := fmt.Sprintf(
			"%v activities:\n"+
				"You uploaded a photo\n"+
				"You liked your photo\n", user1.Name)

		assert.Equal(t, expected, *display)
	})

	t.Run("Should display John's activities John(you) upload a photo and Wick likes your photo", func(t *testing.T) {
		users := NewUserActivity()

		user1 := new(User)
		user1.Name = "John"
		user2 := new(User)
		user2.Name = "Wick"

		users.Setup(*user2, *user1)

		_ = users.Upload(*user1)
		_ = users.Like(*user2, *user1)

		display, _ := users.Display(*user1)

		expected := fmt.Sprintf(
			"%v activities:\n"+
				"You uploaded a photo\n"+
				"%v liked your photo\n", user1.Name, user2.Name)

		assert.Equal(t, expected, *display)
	})

	t.Run("Should display Wick's activities John upload a photo and Wick(you) likes John's photo", func(t *testing.T) {
		users := NewUserActivity()

		user1 := new(User)
		user1.Name = "John"
		user2 := new(User)
		user2.Name = "Wick"

		users.Setup(*user1, *user2)
		users.Setup(*user2, *user1)

		_ = users.Upload(*user1)
		_ = users.Like(*user2, *user1)

		display, _ := users.Display(*user2)

		expected := fmt.Sprintf(
			"%v activities:\n"+
				"%v uploaded a photo\n"+
				"You liked %v's photo\n", user2.Name, user1.Name, user1.Name)

		assert.Equal(t, expected, *display)
	})

	t.Run("Should display Cena's activities Wick likes John's photo", func(t *testing.T) {
		users := NewUserActivity()

		user1 := new(User)
		user1.Name = "John"
		user2 := new(User)
		user2.Name = "Wick"
		user3 := new(User)
		user3.Name = "Cena"

		users.Setup(*user2, *user1)
		users.Setup(*user3, *user2)

		_ = users.Upload(*user1)
		_ = users.Like(*user2, *user1)

		display, _ := users.Display(*user3)

		expected := fmt.Sprintf(
			"%v activities:\n"+
				"%v liked %v's photo\n", user3.Name, user2.Name, user1.Name)

		assert.Equal(t, expected, *display)
	})

	t.Run("Should display error when user has not been setup", func(t *testing.T) {
		users := NewUserActivity()

		user1 := new(User)
		user1.Name = "John"
		user2 := new(User)
		user2.Name = "Wick"

		_, err := users.Display(*user1)
		expected := fmt.Sprintf("Unknown user %v", user1.Name)

		assert.ErrorContains(t, err, expected)
	})
}
