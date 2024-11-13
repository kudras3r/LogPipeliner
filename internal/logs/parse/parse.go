package parse

import (
	"bufio"
	"os"
	"path/filepath"
	"strings"
)

// Walk around logs dir and parse logs in structs.

type BaseLog struct {
	Timestamp string `json:"timestamp"`
	Content   string `json:"content"`
}

type DpkgLog struct {
	BaseLog
}

func Dpkg(path string) ([]DpkgLog, error) {
	var result []DpkgLog
	var timestamp, content string

	err := filepath.Walk(path, func(filePath string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() {
			filename := info.Name()
			if strings.Contains(filename, ".log") && strings.HasPrefix(filename, "dpkg") {
				file, err := os.Open(filePath)
				if err != nil {

				}
				defer file.Close()

				scanner := bufio.NewScanner(file)
				for scanner.Scan() {
					lineParts := strings.Split(scanner.Text(), " ")

					if len(lineParts) < 3 {
						break
					}

					timestamp = lineParts[0] + " " + lineParts[1]
					content = strings.Join(lineParts[2:], " ")

					log := DpkgLog{
						BaseLog: BaseLog{
							Timestamp: timestamp,
							Content:   content,
						},
					}
					result = append(result, log)
				}
			}
		}
		return nil
	})

	if err != nil {
		return result, err
	}

	return result, nil
}
