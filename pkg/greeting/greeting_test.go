package greeting

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHelloWorld(t *testing.T) {

	var expected string = "Hello World!"
	var actual string = HelloWorld()

	assert.Equal(t, expected, actual, "Should print hello world message")

}

func TestGreetMe(t *testing.T) {

	var expected string = "Hello Me!"
	var actual string = GreetMe("Me")

	assert.Equal(t, expected, actual, "Should print message greeting Me")

}
