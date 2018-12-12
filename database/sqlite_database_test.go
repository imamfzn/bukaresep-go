package database_test

import (
	"github.com/imamfzn/bukaresep-go/database"
	"testing"
)

func TestNewDatabase(t *testing.T) {
	t.Run("right db filename", func(t *testing.T) {
		_, err := database.CreateDatabase("file::memory:")

		if err != nil {
			t.Fail()
		}
	})

	t.Run("wrong db filename", func(t *testing.T) {
		_, err := database.CreateDatabase("/dev/null")

		if err == nil {
			t.Fail()
		}
	})
}
