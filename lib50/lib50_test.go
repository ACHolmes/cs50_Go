package lib50

import "testing"

func TestGetInt(t *testing.T) {
	t.Run("test", func(t *testing.T) {
		_ = GetInt("Test: ")
	})
}
