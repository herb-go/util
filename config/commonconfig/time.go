package commonconfig

import "time"

var defaultDateLayout = "2006-01-02"
var defaultDatetimeLayout = "2006-01-02 15:04:05"
var defaultTimeLayout = "15:04:05"

// TimeConfig app time config
type TimeConfig struct {
	//Timezone time zone.
	Timezone string
	//TimeLayout  format used when converting time in day.
	TimeLayout string
	//DateLayout  format used when converting date.
	DateLayout string
	//DatetimeLayout format used when converting date and time.
	DatetimeLayout string
	location       *time.Location
}

//Parse time string to local time.
//Panic if  Timezone error.
func (c *TimeConfig) Parse(s string) (time.Time, error) {
	format := c.DatetimeLayout
	if format == "" {
		format = defaultDatetimeLayout
	}
	return time.ParseInLocation(format, s, c.loadLocation())
}

func (c *TimeConfig) MustParse(s string) time.Time {
	t, err := c.Parse(s)
	if err != nil {
		panic(err)
	}
	return t
}

func (c *TimeConfig) Location() *time.Location {
	return c.loadLocation()
}
func (c *TimeConfig) loadLocation() *time.Location {
	if c.location == nil {
		if c.Timezone == "" {
			c.location = time.Local
		} else {
			var err error
			c.location, err = time.LoadLocation(c.Timezone)
			if err != nil {
				panic(err)
			}
		}
	}
	return c.location
}

//TimeInLocation set time location to given time zone.
//Panic if  Timezone error.
func (c *TimeConfig) TimeInLocation(t time.Time) time.Time {
	return t.In(c.loadLocation())
}

//DateUnix format date from unix timestamp
func (c *TimeConfig) DateUnix(ts int64) string {
	return c.Date(time.Unix(ts, 0))
}

//Date format date.
//Panic if  Timezone error.
func (c *TimeConfig) Date(t time.Time) string {
	localTime := c.TimeInLocation(t)
	if c.DateLayout == "" {
		return localTime.Format(defaultDateLayout)
	}
	return localTime.Format(c.DateLayout)
}

//TimeUnix format time from unix timestamp
func (c *TimeConfig) TimeUnix(ts int64) string {
	return c.Time(time.Unix(ts, 0))
}

//Time format time.
func (c *TimeConfig) Time(t time.Time) string {
	localTime := c.TimeInLocation(t)

	if c.TimeLayout == "" {
		return localTime.Format(defaultTimeLayout)
	}
	return localTime.Format(c.TimeLayout)
}

//DatetimeUnix format date and time from unix timestamp
//Panic if  Timezone error.
func (c *TimeConfig) DatetimeUnix(ts int64) string {
	return c.Datetime(time.Unix(ts, 0))
}

//Datetime format date and time
func (c *TimeConfig) Datetime(t time.Time) string {
	localTime := c.TimeInLocation(t)

	if c.DatetimeLayout == "" {
		return localTime.Format(defaultDatetimeLayout)
	}
	return localTime.Format(c.DatetimeLayout)
}

//FormatNow format with given time
func (c *TimeConfig) FormatNow(format string) string {
	localTime := c.TimeInLocation(time.Now())
	return localTime.Format(format)
}
