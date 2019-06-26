package commonconfig

import "testing"

func TestTime(t *testing.T) {
	defaulttc := &TimeConfig{}

	timestring := "2019-01-02 01:02:03"

	ts, err := defaulttc.Parse(timestring)
	if err != nil {
		t.Fatal(err)
	}
	datetime := defaulttc.Datetime(ts)
	if datetime != timestring {
		t.Fatal(datetime)
	}

	date := defaulttc.Date(ts)
	if date != "2019-01-02" {
		t.Fatal(date)
	}
	timelabel := defaulttc.Time(ts)
	if timelabel != "01:02:03" {
		t.Fatal(timelabel)
	}
	tc := &TimeConfig{
		DateFormat:     "2006-01-02!",
		DatetimeFormat: "2006-01-02 15:04:05!",
		TimeFormat:     "15:04:05!",
		Timezone:       "Asia/Shanghai",
	}

	timestring = "2019-01-02 01:02:03!"
	ts, err = tc.Parse(timestring)
	if err != nil {
		t.Fatal(err)
	}
	datetime = tc.Datetime(ts)
	if datetime != timestring {
		t.Fatal(datetime)
	}
	datetime = tc.DatetimeUnix(ts.Unix())
	if datetime != timestring {
		t.Fatal(datetime)
	}
	date = tc.Date(ts)
	if date != "2019-01-02!" {
		t.Fatal(date)
	}
	date = tc.DateUnix(ts.Unix())
	if date != "2019-01-02!" {
		t.Fatal(date)
	}

	timelabel = tc.Time(ts)
	if timelabel != "01:02:03!" {
		t.Fatal(timelabel)
	}
	timelabel = tc.TimeUnix(ts.Unix())
	if timelabel != "01:02:03!" {
		t.Fatal(timelabel)
	}
}
