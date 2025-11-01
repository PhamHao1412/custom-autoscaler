package logging

import (
	"io"
	"log"
	"os"
	"path/filepath"
)

func InitLogger(defaultPath string) {
	logPath := os.Getenv("LOG_PATH")
	if logPath == "" {
		logPath = defaultPath
	}
	// Ensure directory exists
	dir := filepath.Dir(logPath)
	if err := os.MkdirAll(dir, 0o755); err != nil {
		log.Fatalf("failed to create log dir %s: %v", dir, err)
	}

	// Open log file for append
	f, err := os.OpenFile(logPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0o644)
	if err != nil {
		log.Fatalf("failed to open log file %s: %v", logPath, err)
	}

	// Write to both stdout and file
	mw := io.MultiWriter(os.Stdout, f)
	log.SetOutput(mw)
	log.SetFlags(log.Ldate | log.Ltime | log.Lmicroseconds | log.Lshortfile)
	log.Printf("🪵 Logger initialized — writing to %s", logPath)
}
