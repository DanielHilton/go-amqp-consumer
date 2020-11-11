package errors

import "log"

// ExitOnFail fatally fails the application with a log if err is not nil
func ExitOnFail(err error, msg string) {
	if err != nil {
		log.Fatalf("%s, %s", msg, err)
	}
}
