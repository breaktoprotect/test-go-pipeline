package shapes

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var r *Rectangle

func setup() {
	r = NewRectangle(10, 15)
}

func TestNewRectangle(t *testing.T) {
	r := Rectangle{Width: 10, Height: 15}
	expectedWidth := 10.0
	gotWidth := r.Width
	assert.Equal(t, expectedWidth, gotWidth, "Expected width to be %v but got %v")
	expectedHeight := 15.0
	gotHeight := r.Height
	assert.Equal(t, expectedHeight, gotHeight, "Expected height to be %v but got %v")

}

func TestGetArea(t *testing.T) {
	setup()
	expected := 150.0
	got := r.getArea()
	assert.Equal(t, expected, got, "Expected area to be %v, but got %v")
}

func TestGetPerimeter(t *testing.T) {
	setup()
	expected := 50.0
	got := r.getPerimeter()
	assert.Equal(t, expected, got, "Expected perimeter to be %v, but got %v")
}
