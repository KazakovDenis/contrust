package local_errors

type DatabaseReadError struct {
}

func (e *DatabaseReadError) Error() string {
	return "DatabaseReadError"
}

type DatabaseWriteError struct {
}

func (e *DatabaseWriteError) Error() string {
	return "DatabaseWriteError"
}
