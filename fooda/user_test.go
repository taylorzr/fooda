package fooda

import (
	"reflect"
	"testing"
)

func TestAllUsers(t *testing.T) {
	users, err := AllUsers()

	if err != nil {
		t.Error(err)
	}

	actual := make([]string, len(users))

	for i, user := range users {
		actual[i] = user.Handle
	}

	expected := []string{"nothotdog", "potato"}

	ok := reflect.DeepEqual(actual, expected)

	if !ok {
		t.Errorf("Expected handles to match\n%s\n%s", expected, actual)
	}
}
