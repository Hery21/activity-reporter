package activity

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUserActivity_Trending(t *testing.T) {
	t.Run("Should return true if user exist and false if not exist", func(t *testing.T) {
		users := NewUserActivity()
		firstTrendLen := 0
		secondTrendLen := 0
		thirdTrendLen := 0

		user1 := new(User)
		user1.Name = "John"
		user2 := new(User)
		user2.Name = "Wick"
		user3 := new(User)
		user3.Name = "Cena"
		user4 := new(User)
		user4.Name = "Felix"
		user5 := new(User)
		user5.Name = "Anthony"

		users.Setup(*user1, *user2)
		users.Setup(*user1, *user3)
		users.Setup(*user1, *user4)
		users.Setup(*user1, *user5)
		users.Setup(*user2, *user1)
		users.Setup(*user2, *user3)
		users.Setup(*user2, *user4)
		users.Setup(*user2, *user5)
		users.Setup(*user3, *user1)
		users.Setup(*user3, *user2)
		users.Setup(*user3, *user4)
		users.Setup(*user3, *user5)
		users.Setup(*user4, *user1)
		users.Setup(*user4, *user2)
		users.Setup(*user4, *user3)
		users.Setup(*user4, *user5)
		users.Setup(*user5, *user1)
		users.Setup(*user5, *user2)
		users.Setup(*user5, *user3)
		users.Setup(*user5, *user4)

		_ = users.Upload(*user1)
		_ = users.Upload(*user2)
		_ = users.Upload(*user3)
		_ = users.Upload(*user4)
		_ = users.Upload(*user5)

		_ = users.Like(*user2, *user1)
		firstTrendLen += 1
		_ = users.Like(*user3, *user1)
		firstTrendLen += 1
		_ = users.Like(*user4, *user1)
		firstTrendLen += 1
		_ = users.Like(*user5, *user1)
		firstTrendLen += 1
		_ = users.Like(*user3, *user2)
		secondTrendLen += 1
		_ = users.Like(*user4, *user2)
		secondTrendLen += 1
		_ = users.Like(*user5, *user2)
		secondTrendLen += 1
		_ = users.Like(*user4, *user3)
		thirdTrendLen += 1
		_ = users.Like(*user5, *user3)
		thirdTrendLen += 1
		_ = users.Like(*user5, *user4)

		trend := users.Trending()

		expected := fmt.Sprintf(
			"Trending photos:\n"+
				"1. %v photo got %v likes\n"+
				"2. %v photo got %v likes\n"+
				"3. %v photo got %v likes\n", user1.Name, firstTrendLen, user2.Name, secondTrendLen, user3.Name, thirdTrendLen)

		assert.Equal(t, expected, trend)
	})
}
