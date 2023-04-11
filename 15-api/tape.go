package api

import "os"

type tape struct {
	file *os.File
}

func (t *tape) Write(p []byte) (n int, err error) {
	err = t.file.Truncate(0)
	if err != nil {
		return 0, err
	}

	_, err = t.file.Seek(0, 0)
	if err != nil {
		return 0, err
	}

	return t.file.Write(p)
}
