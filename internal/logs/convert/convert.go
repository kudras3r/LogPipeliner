package convert

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"github.com/kudras3r/LogPipeliner/internal/logs/parse"
)

const (
	dpkgLayout = "2006-01-02 15:04:05"
)

type ResourceLogs struct {
	ResourceLogs []ResourceLog `json:"resourceLogs"`
}

type ResourceLog struct {
	Resource  Resource   `json:"resource"`
	ScopeLogs []ScopeLog `json:"scopeLogs"`
}

type Resource struct {
}

type ScopeLog struct {
	Scope      Scope       `json:"scope"`
	LogRecords []LogRecord `json:"logRecords"`
}

type Scope struct {
}

type LogRecord struct {
	TimeUnixNano string  `json:"timeUnixNano"`
	Body         LogBody `json:"body"`
	TraceID      string  `json:"traceId"`
	SpanID       string  `json:"spanId"`
}

type LogBody struct {
	StringValue string `json:"stringValue"`
}

func toUnixNano(layout, timeStr string) string {
	t, _ := time.Parse(layout, timeStr)
	return strconv.Itoa(int(t.UnixNano()))
}

func DpkgsToJson(logs []parse.DpkgLog) ([]byte, error) {
	var resLogsTmp []ResourceLog
	var resLogs ResourceLogs

	for _, l := range logs {
		convLog := ResourceLog{
			Resource: Resource{},
			ScopeLogs: []ScopeLog{
				{
					Scope: Scope{},
					LogRecords: []LogRecord{
						{
							TimeUnixNano: toUnixNano(dpkgLayout, l.Timestamp),
							Body: LogBody{
								StringValue: fmt.Sprintf("{\"message\":\"%s\"}", l.Content),
							},
							TraceID: "",
							SpanID:  "",
						},
					},
				},
			},
		}
		resLogsTmp = append(resLogsTmp, convLog)
	}

	resLogs = ResourceLogs{
		ResourceLogs: resLogsTmp,
	}

	jsonData, err := json.Marshal(resLogs)
	if err != nil {
		return nil, err
	}

	return jsonData, nil
}
