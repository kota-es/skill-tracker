package models

import (
	"encoding/json"
	"time"
)

// TODO: ユーザーの設定情報で変えられるようにしたい
const JST_TIME_ZONE = "Asia/Tokyo"

type JstTime struct {
	Time time.Time
}

func (t JstTime) MarshalJSON() ([]byte, error) {
	loc, err := time.LoadLocation(JST_TIME_ZONE)
	if err != nil {
		return nil, err
	}
	return json.Marshal(t.Time.In(loc).Format(time.RFC3339))
}

func (j *JstTime) UnmarshalJSON(data []byte) error {
	var str string
	if err := json.Unmarshal(data, &str); err != nil {
		return err
	}

	t, err := time.Parse(time.RFC3339, str)
	if err != nil {
		return err
	}
	j.Time = t
	return nil
}
