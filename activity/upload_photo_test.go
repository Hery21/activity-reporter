package activity

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUserActivity_Upload(t *testing.T) {
	t.Run("Should return true when John upload photo", func(t *testing.T) {
		users := NewUserActivity()

		user1 := new(User)
		user1.Name = "John"
		user2 := new(User)
		user2.Name = "Wick"

		users.Setup(*user1, *user2)

		_ = users.Upload(*user1)

		actual1 := users.Activity[*user1].Upload

		assert.True(t, actual1)
	})

	t.Run("Should return true when John upload photo given John already uploaded photo before", func(t *testing.T) {
		users := NewUserActivity()

		user1 := new(User)
		user1.Name = "John"
		user2 := new(User)
		user2.Name = "Wick"

		users.Setup(*user1, *user2)

		_ = users.Upload(*user1)
		_ = users.Upload(*user1)

		actual1 := users.Activity[*user1].Upload

		assert.True(t, actual1)
	})
}
