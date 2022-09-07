package encode

import (
	xtime "time"
)

// Duration be used toml,json unmarshal string time, like 1s, 500ms.
type Duration xtime.Duration

func (d *Duration) UnmarshalText(text []byte) error {
	tmp, err := xtime.ParseDuration(string(text))
	if err == nil {
		*d = Duration(tmp)
	}
	return err
}

func (d *Duration) MarshalText() ([]byte, error) {
	tmp := xtime.Duration(*d).String()
	return []byte(tmp), nil
}
