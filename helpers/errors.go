package errors

import "log"

// FailOnError fatally fails the application with a log if err is not nil
func FailOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s, %s", msg, err)
	}
}
