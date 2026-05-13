package scraper

import "testing"

func TestCleanText(t *testing.T) {
	input := " Robotics   is \n the   study\t of robots. "
	expected := "Robotics is the study of robots."

	actual := CleanText(input)

	if actual != expected {
		t.Errorf("expected %q but got %q", expected, actual)
	}
}

func TestCountWords(t *testing.T) {
	input := "Robotics is the study of robots"
	expected := 6

	actual := CountWords(input)

	if actual != expected {
		t.Errorf("expected %d but got %d", expected, actual)
	}
}
