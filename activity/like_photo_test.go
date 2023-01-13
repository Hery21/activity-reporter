package activity

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUserActivity_Like(t *testing.T) {
	t.Run("Should return nil when Wick followed John and like John's photo", func(t *testing.T) {
		users := NewUserActivity()

		user1 := new(User)
		user1.Name = "John"
		user2 := new(User)
		user2.Name = "Wick"

		users.Setup(*user2, *user1)
		_ = users.Upload(*user1)
		err := users.Like(*user2, *user1)

		assert.Nil(t, err)
	})

	t.Run("Should return error when Cena has not been setup and Wick likes Cena's photo", func(t *testing.T) {
		users := NewUserActivity()

		user1 := new(User)
		user1.Name = "John"
		user2 := new(User)
		user2.Name = "Wick"
		user3 := new(User)
		user3.Name = "Cena"

		users.Setup(*user1, *user2)
		err := users.Like(*user2, *user3)
		expected := fmt.Sprintf("Unknown user %v", user3.Name)

		assert.ErrorContains(t, err, expected)
	})

	t.Run("Should return error when Cena has not been setup and Cena likes Wick's photo", func(t *testing.T) {
		users := NewUserActivity()

		user1 := new(User)
		user1.Name = "John"
		user2 := new(User)
		user2.Name = "Wick"
		user3 := new(User)
		user3.Name = "Cena"

		users.Setup(*user1, *user2)
		err := users.Like(*user3, *user2)
		expected := fmt.Sprintf("Unknown user %v", user3.Name)

		assert.ErrorContains(t, err, expected)
	})

	t.Run("Should return nil when Cena likes Wick's photo given Cena has liked Wick's photo", func(t *testing.T) {
		users := NewUserActivity()

		user2 := new(User)
		user2.Name = "Wick"
		user3 := new(User)
		user3.Name = "Cena"

		users.Setup(*user3, *user2)
		_ = users.Upload(*user2)
		_ = users.Like(*user3, *user2)
		err := users.Like(*user3, *user2)

		assert.Nil(t, err)
	})

	t.Run("Should return error when Cena likes Wick's photo given Wick has not uploaded photo", func(t *testing.T) {
		users := NewUserActivity()

		user2 := new(User)
		user2.Name = "Wick"
		user3 := new(User)
		user3.Name = "Cena"

		users.Setup(*user3, *user2)
		err := users.Like(*user3, *user2)

		assert.Equal(t, err, ErrorNoPhoto{})
	})
}
