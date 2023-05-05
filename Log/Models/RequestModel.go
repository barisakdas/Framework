package log

import "time"

type RequestModel struct {
	TimeStamp  time.Time         `json:"time_stamp"`
	StatusCode int               `json:"status_code"`
	Method     string            `json:"method"`
	Path       string            `json:"path"`
	Body       string            `json:"body"`
	Headers    map[string]string `json:"headers"`
}
