package epoch

import (
	"database/sql/driver"
	"github.com/GagulProject/go-whisky/internal/shared/errors"
	"strconv"
	"strings"
	"time"

	"github.com/samber/lo"
)

type Milli struct {
	time.Time
}

func ParseMilliFromTime(time time.Time) Milli {
	return Milli{time.UTC()}
}

func ParseMilliFromTimePtr(time *time.Time) *Milli {
	if time == nil {
		return nil
	}
	return &Milli{time.UTC()}
}

func ParseMilliFromInt64(milliSec int64) Milli {
	return ParseMilliFromTime(time.UnixMilli(milliSec))
}

func ParseMilliFromString(str string) (Milli, error) {
	sinceInt, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		return Milli{}, errors.Wrap(err)
	}

	return ParseMilliFromInt64(sinceInt), nil
}

func (m *Milli) ToTime() *time.Time {
	if m == nil {
		return nil
	}
	return lo.ToPtr(m.Time)
}

func (m Milli) String() string {
	return strconv.FormatInt(m.UnixMilli(), 10)
}

func (m Milli) MarshalJSON() ([]byte, error) {
	return []byte(strconv.FormatInt(m.UnixMilli(), 10)), nil
}

// UnmarshalJSON
// timestamp 값을 받으면서 종종 끼어들어오는 ""를 제거해주기 위하여 작업.
func (m *Milli) UnmarshalJSON(data []byte) error {
	t := strings.Trim(string(data), `"`)
	milliSec, err := strconv.ParseInt(t, 10, 64)
	if err != nil {
		return errors.Wrap(err)
	}

	*m = ParseMilliFromInt64(milliSec)
	return nil
}

// UnmarshalText implements encoding.TextUnmarshaler.
// It is used to parse string into epoch.Milli.
// Parameter data is byte of unix milliseconds.
func (m *Milli) UnmarshalText(data []byte) error {
	t := string(data)
	milliSec, err := strconv.ParseInt(t, 10, 64)
	if err != nil {
		return errors.Wrap(err)
	}

	*m = ParseMilliFromInt64(milliSec)
	return nil
}

func (m Milli) Value() (driver.Value, error) {
	return m.Time, nil
}

func (m *Milli) Scan(value any) error {
	if value == nil {
		return nil
	}

	t, ok := value.(time.Time)
	if !ok {
		return errors.New("invalid time type")
	}

	*m = ParseMilliFromTime(t)
	return nil
}
