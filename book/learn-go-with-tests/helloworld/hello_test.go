package helloworld

import "testing"

// ===================== No helper =====================
// func TestHello(t *testing.T) {

// 	t.Run("saying hello to people", func(t *testing.T) {
// 		got := Hello("Chris")
// 		want := "Hello, Chris"

// 		if got != want {
// 			t.Errorf("got '%q' want '%q'", got, want)
// 		}
// 	})

// 	t.Run("say hello world when an empty string is supplied", func(t *testing.T) {
// 		got := Hello("")
// 		want := "Hello, World"

// 		if got != want {
// 			t.Errorf("got '%q' want '%q'", got, want)
// 		}
// 	})

// }

func TestHello(t *testing.T) {

	assertCorrectMessage := func(t *testing.T, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got '%q' want '%q'", got, want)
		}
	}

	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Chris", "English")
		want := "Hello, Chris"
		assertCorrectMessage(t, got, want)
	})

	t.Run("empty string defaults to 'world'", func(t *testing.T) {
		got := Hello("", "English")
		want := "Hello, World"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Elodie", "Spanish")
		want := "Hola, Elodie"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in Franch", func(t *testing.T) {
		got := Hello("beau", "French")
		want := "Bonjour, beau"
		assertCorrectMessage(t, got, want)
	})

}
