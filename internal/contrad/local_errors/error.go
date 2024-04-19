package local_errors

import "fmt"

type DatabaseWriteError struct {
	error
}

func (e *DatabaseWriteError) Error() string {
	return fmt.Sprintf("%s", e)
}
