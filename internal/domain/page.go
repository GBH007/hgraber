package domain

import "time"

type Page struct {
	URL      string
	Ext      string
	Success  bool
	LoadedAt time.Time
	Rate     int
}

type PageFullInfo struct {
	BookID     int
	PageNumber int
	URL        string
	Ext        string
	Success    bool
	LoadedAt   time.Time
	Rate       int
}
