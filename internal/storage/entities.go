package storage

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"time"
)

type Data map[string]any

type Event struct {
	Timestamp time.Time
	Source    string
	Data      Data
}

func (a Data) Value() (driver.Value, error) {
	return json.Marshal(a)
}

func (a *Data) Scan(value any) error {
	b, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}

	return json.Unmarshal(b, &a)
}
