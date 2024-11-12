package parse

// Walk around logs dir and parse logs in structs.

// in future we can do composition to describe specific logs
// like AuthLog, BootLog...
type Log struct {
	Timestamp string `json:"timestamp"`
	Content   string `json:"content"`
}

func Parse(logsType string) {

}
