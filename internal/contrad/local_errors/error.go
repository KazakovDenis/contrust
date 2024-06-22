package local_errors

type DatabaseWriteError struct {
}

func (e *DatabaseWriteError) Error() string {
	return "DatabaseWriteError"
}
