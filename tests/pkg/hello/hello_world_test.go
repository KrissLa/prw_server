package hello_test

import (
	"prw_server/app/pkg/hello"
	"testing"
)

func TestHelloWorld(t *testing.T) {
	t.Run("Hello world", func(t *testing.T) {
		got := hello.World()
		want := "Hello world!"

		if got != want {
			t.Errorf("fail")
		}
	})
}
