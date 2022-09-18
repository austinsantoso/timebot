package time

import (
	gotime "time"
)

type TimeModule struct {
	Timestamp int64
	time      gotime.Time
}

// https://stackoverflow.com/questions/20234104/how-to-format-current-time-using-a-yyyymmddhhmmss-format
const timeFormat = "Monday 02 January 2006 at 3:04:05.999 PM (-07:00)"

func NewSetTimeModule(ms int64) *TimeModule {
	return &TimeModule{Timestamp: ms, time: gotime.UnixMilli(ms).UTC()}
}

func NewTimeModule(t gotime.Time) *TimeModule {
	return NewSetTimeModule((t.UnixMilli()))
}

// returns go time but at UTC +00
func Now() *TimeModule {
	return NewSetTimeModule(gotime.Now().UnixMilli())
}

func (t *TimeModule) String() string {
	return t.time.Format(timeFormat)
}

func (t *TimeModule) Add(d gotime.Duration) *TimeModule {
	newTime := t.time.Add(d)
	return NewTimeModule(newTime)
}
