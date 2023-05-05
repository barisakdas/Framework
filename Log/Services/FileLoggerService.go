package log

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"time"

	. "github.com/barisakdas/Framework/Log/Models"
)

type FileLoggerService struct {
	LogDir string
}

func (f *FileLoggerService) LogRequest(request RequestModel) {
	today := time.Now().Format("2006-01-02")
	logFilePath := filepath.Join(f.LogDir, today+"_Requests.json")

	logEntry := map[string]interface{}{
		"request": request,
		"time":    time.Now().Format(time.RFC3339Nano),
	}

	logData, err := json.Marshal(logEntry)
	if err != nil {
		fmt.Printf("Error marshalling log entry: %s\n", err)
		return
	}

	if _, err := os.Stat(f.LogDir); os.IsNotExist(err) {
		err = os.MkdirAll(f.LogDir, 0755)
		if err != nil {
			fmt.Printf("Error creating log directory: %s\n", err)
			return
		}
	}

	err = ioutil.WriteFile(logFilePath, logData, 0644)
	if err != nil {
		fmt.Printf("Error writing log entry: %s\n", err)
		return
	}
}

func (f *FileLoggerService) LogResponse(response ResponseModel) {
	today := time.Now().Format("2006-01-02")
	logFilePath := filepath.Join(f.LogDir, today+"_Responses.json")

	logEntry := map[string]interface{}{
		"response": response,
		"time":     time.Now().Format(time.RFC3339Nano),
	}

	logData, err := json.Marshal(logEntry)
	if err != nil {
		fmt.Printf("Error marshalling log entry: %s\n", err)
		return
	}

	if _, err := os.Stat(f.LogDir); os.IsNotExist(err) {
		err = os.MkdirAll(f.LogDir, 0755)
		if err != nil {
			fmt.Printf("Error creating log directory: %s\n", err)
			return
		}
	}

	err = ioutil.WriteFile(logFilePath, logData, 0644)
	if err != nil {
		fmt.Printf("Error writing log entry: %s\n", err)
		return
	}
}
