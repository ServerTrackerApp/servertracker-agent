/*
 * Copyright (c) 2024 Luca Fröschke
 */

package log

import (
	"fmt"
	"time"
)

const (
	INFO  = "INFO"
	WARN  = "WARN"
	ERROR = "ERROR"
	DEBUG = "DEBUG"
)

func Log(message string, level string) {
	timestamp := time.Now().Format(time.RFC3339)
	fmt.Println(
		fmt.Sprintf("[%s] [%s] %s", timestamp, level, message),
	)
}
