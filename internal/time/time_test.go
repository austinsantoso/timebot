package time

import (
	"testing"
	gotime "time"

	"github.com/stretchr/testify/assert"
)

func TestNowUTC(t *testing.T) {
	testcase := NewSetTimeModule(1662797475288).String()
	expected := "Saturday 10 September 2022 at 8:11:15.288 AM (+00:00)"

	assert.Equal(t, expected, testcase)
}

func TestAdd(t *testing.T) {
	testCases := map[string]struct {
		input *TimeModule
		want  string
	}{
		"one hour ahead":  {input: NewSetTimeModule(1662797475288).Add(gotime.Hour * gotime.Duration(1)), want: "Saturday 10 September 2022 at 9:11:15.288 AM (+00:00)"},
		"one hour behind": {input: NewSetTimeModule(1662797475288).Add(gotime.Hour * gotime.Duration(-1)), want: "Saturday 10 September 2022 at 7:11:15.288 AM (+00:00)"},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			res := tc.input.String()
			assert.Equal(t, tc.want, res)
		})
	}
}
