package main

import (
	"fmt"
	"time"
)

const (
	colorRed    = "\033[31m"
	colorGreen  = "\033[32m"
	colorYellow = "\033[33m"
	colorBlue   = "\033[34m"
	colorPurple = "\033[35m"
	colorCyan   = "\033[36m"
	colorWhite  = "\033[37m"
)

func LogInfo(log string) {

	log_time := fmt.Sprintf("%s [%s]", colorWhite, time.Now().Format("2006-02-01"))
	log_type := fmt.Sprintf("%s [LOG INFO]", colorCyan)
	log_content := fmt.Sprintf("%s : %s", colorWhite, log)

	fmt.Printf("%s %s %s\n", log_time, log_type, log_content)
}

func LogError(err error, args map[string]interface{}) {

	log_time := fmt.Sprintf("%s [%s]", colorWhite, time.Now().Format("2006-02-01"))
	log_type := fmt.Sprintf("%s [LOG ERROR]", colorRed)
	log_error := fmt.Sprintf("Error : %s %s", colorRed, err.Error())

	fmt.Println()
	fmt.Printf("%s %s\n", log_time, log_type)
	fmt.Printf("%s  %s\n", log_time, log_error)
	for key, el := range args {
		fmt.Printf("%s %s %s : %v \n", colorBlue, key, colorWhite, el)
	}
	fmt.Println()
}

func LogWarn(args map[string]interface{}) {

	log_time := fmt.Sprintf("%s [%s]", colorWhite, time.Now().Format("2006-02-01"))
	log_type := fmt.Sprintf("%s [LOG WARN]", colorYellow)

	fmt.Println()
	fmt.Printf("%s %s\n", log_time, log_type)
	for key, el := range args {
		fmt.Printf("%s %s %s : %v \n", colorBlue, key, colorWhite, el)
	}
	fmt.Println()
}
