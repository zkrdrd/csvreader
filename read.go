package csvreader

import (
	"io"
	"os"
	"strings"
)

const bufferSizeForFileRead = 1024

// Чтение файла csv
func Read(FilenameFrom string, FilenameTo string) error {

	var buffer = make([]byte, bufferSizeForFileRead)

	dataFileFrom, err := os.ReadFile(FilenameFrom) //os.Open
	if err != nil {
		return err
	}

	dataFileTo, err := os.OpenFile(FilenameTo, os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		return err
	}
	defer func() error {
		if err := dataFileTo.Close(); err != nil {
			return err
		}
		return nil
	}()

	readFrom := strings.NewReader(string(dataFileFrom))
	if _, err := io.CopyBuffer(dataFileTo, readFrom, buffer); err != nil {
		return err
	}

	return nil
}
