package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Vikash", "")
		want := "Hello world, I am Vikash here!"

		assertCorrectMessage(t, got, want)
	})
	t.Run("say 'Hello world, I am World here!' when an empty string is supplied", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello world, I am World here!"

		assertCorrectMessage(t, got, want)
	})
	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Vikash", "Spanish")
		want := "Hola mundo, ¡soy Vikash aquí!"
		assertCorrectMessage(t, got, want)
	})
	t.Run("in French", func(t *testing.T) {
		got := Hello("Vikash", "French")
		want := "Bonjour tout le monde, je suis Vikash ici!"
		assertCorrectMessage(t, got, want)
	})
	t.Run("in Bhojpuri", func(t *testing.T) {
		got := Hello("Vikash", "Bhojpuri")
		want := "नमस्ते दुनिया, हम Vikash यहाँ बा!"
		assertCorrectMessage(t, got, want)
	})
	t.Run("in Kannada", func(t *testing.T) {
		got := Hello("Vikash", "Kannada")
		want := "ಹಲೋ ವರ್ಲ್ಡ್, ನಾನು Vikash ಇಲ್ಲಿ!"
		assertCorrectMessage(t, got, want)
	})
	t.Run("in Japanese", func(t *testing.T) {
		got := Hello("Vikash", "Japanese")
		want := "こんにちは世界、私はここでVikashです!"
		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
