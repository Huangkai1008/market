package utils

import (
	"database/sql/driver"
	"fmt"
	"time"
)

type JsonTime struct {
	time.Time
}

func (jsonTime JsonTime) MarshalJSON() ([]byte, error) {
	formatted := fmt.Sprintf("\"%s\"", jsonTime.Format("2006-01-02 15:04:05"))
	return []byte(formatted), nil
}

// Value insert timestamp into mysql need this function.
func (jsonTime JsonTime) Value() (driver.Value, error) {
	var zeroTime time.Time
	if jsonTime.Time.UnixNano() == zeroTime.UnixNano() {
		return nil, nil
	}
	return jsonTime.Time, nil
}

// Scan valueOf time.Time
func (jsonTime *JsonTime) Scan(v interface{}) error {
	value, ok := v.(time.Time)
	if ok {
		*jsonTime = JsonTime{Time: value}
		return nil
	}
	return fmt.Errorf("can not convert %v to timestamp", v)
}
