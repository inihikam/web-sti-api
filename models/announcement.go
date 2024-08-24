package models

import (
	"fmt"
	"strings"
	"time"
)

type CustomTime time.Time

const customTimeFormat = "2006-01-02 15:04:05"

func (ct *CustomTime) UnmarshalJSON(b []byte) error {
    s := strings.Trim(string(b), "\"")
    t, err := time.Parse(customTimeFormat, s)
    if err != nil {
        return err
    }
    *ct = CustomTime(t)
    return nil
}

func (ct CustomTime) MarshalJSON() ([]byte, error) {
    return []byte(fmt.Sprintf("\"%s\"", time.Time(ct).Format(customTimeFormat))), nil
}

func (ct CustomTime) Format(format string) string {
    t := time.Time(ct)
    return t.Format(format)
}


type Announcement struct {
	ID           uint   `json:"id"`
	Judul        string `json:"judul"`
	Penulis      string `json:"penulis"`
	Konten       string `json:"konten"`
	Published_at CustomTime `json:"published_at"`
}