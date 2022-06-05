package main

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

type LogEntry struct {
	ClientIP   string
	Time       string
	Method     string
	Path       string
	Proto      string
	StatusCode int
	Latency    time.Duration
	UserAgent  string
	Msg        string
}

func JSONLogger(param gin.LogFormatterParams) string {
	entry := LogEntry{
		param.ClientIP,
		param.TimeStamp.Format(time.RFC3339),
		param.Method,
		param.Path,
		param.Request.Proto,
		param.StatusCode,
		param.Latency,
		param.Request.UserAgent(),
		param.ErrorMessage,
	}

	entryJson, _ := json.Marshal(entry)

	// your custom format
	return fmt.Sprintln(string(entryJson))
}
