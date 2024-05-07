package csvreader

import (
	"io"
	"os"
)

const bufferSizeForFileRead = 1

// Чтение файла csv
func Read(FilenameFrom string, FilenameTo string) error {

	var buffer = make([]byte, bufferSizeForFileRead)

	dataFileFrom, err := os.OpenFile(FilenameFrom, os.O_RDONLY, 0644)
	if err != nil {
		return err
	}
	defer dataFileFrom.Close()

	dataFileTo, err := os.OpenFile(FilenameTo, os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		return err
	}
	defer dataFileTo.Close()

	if _, err := io.CopyBuffer(dataFileTo, dataFileFrom, buffer); err != nil {
		return err
	}

	return nil
}
