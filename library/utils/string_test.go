package utils

import "testing"

func Test_GenerateRandomString(t *testing.T) {
	t.Log(GenerateRandomString(16))
}
