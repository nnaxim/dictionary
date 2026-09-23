package data

import "time"

type Words struct {
	ID        int64     `json:"id"`
	CreatedAt time.Time `json:"-"`
	Title     string    `json:"title"`
	Example   string    `json:",omitempty"`
	Synonims  []string  `json:"synonims"`
	Version   int32     `json:"version,string"`
}
