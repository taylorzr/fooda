package fooda

import (
	"testing"
)

var called bool

type FakeChat struct {
}

func (FakeChat) Notify(string, string) error {
	called = true

	return nil
}

func TestRemind(t *testing.T) {
	chat = FakeChat{}

	Remind("nothotdog")

	if !called {
		t.Error("Didn't work :(")
	}
}

func TestForgetters(t *testing.T) {
	forgetters, err := Forgetters()

	if err != nil {
		t.Error(err)
	}

	actual := forgetters[0].Handle
	expected := "potato"

	if actual != expected {
		t.Errorf("Expected potato to forget got %s", actual)
	}
}
