package time

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNowUTC(t *testing.T) {
	testcase := NewSetTimeModule(1662797475288).String()
	expected := "Saturday 10 September 2022 at 8:11:15.288 AM (+00:00)"

	assert.Equal(t, expected, testcase)
}
