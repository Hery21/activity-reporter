package activity

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUserActivity_CheckKeyword(t *testing.T) {
	t.Run("Should return error if syntax is not exact", func(t *testing.T) {
		setup1 := "Alice follows Bob"
		setup2 := "Alice follow Bob"
		setup3 := "Alice Follows Bob"
		action1 := "Alice uploaded photo"
		action2 := "Alice upload photo"
		action3 := "Alice likes Bill photo"
		action4 := "Alice likes Bill"

		err1 := CheckKeyword(setup1)
		err2 := CheckKeyword(setup2)
		err3 := CheckKeyword(setup3)
		err4 := CheckKeyword(action1)
		err5 := CheckKeyword(action2)
		err6 := CheckKeyword(action3)
		err7 := CheckKeyword(action4)

		assert.Nil(t, err1)
		assert.Equal(t, ErrorInvalidKeyword{}, err2)
		assert.Equal(t, ErrorInvalidKeyword{}, err3)
		assert.Nil(t, err4)
		assert.Equal(t, ErrorInvalidKeyword{}, err5)
		assert.Nil(t, err6)
		assert.ErrorContains(t, err7, "Invalid keyword")
	})
}
