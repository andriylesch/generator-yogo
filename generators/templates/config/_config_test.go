package config

import (
	"testing"

	"os"

	"github.com/stretchr/testify/assert"
)

var portConfigTests = []struct {
	key   string
	value string
	out   string
	env   string
	desc  string
}{
	{"APP_PORT", "421", "421", "", "Customized configuration expect"},
}

//TestGet_Port Test the port configuration
func TestGet_Port(t *testing.T) {
	for _, test := range portConfigTests {
		// Arrange
		os.Setenv("ENVIRONMENT", test.env)
		os.Setenv(test.key, test.value)

		//Act
		InitConfig()

		//Assert
		assert.Equal(t, test.out, AppPort, test.desc)
	}
}
