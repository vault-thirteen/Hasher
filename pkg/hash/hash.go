package hash

import (
	"crypto/md5"
	"crypto/sha256"
	"hash/crc32"
	"io"
	"os"

	ae "github.com/vault-thirteen/auxie/errors"
	"github.com/vault-thirteen/auxie/file"
)

const (
	ReadBufferSize = 1_048_576 // 1 MiB.
)

func GetFileHashCRC32(filePath string) (sum uint32, err error) {
	var f *os.File
	f, err = os.Open(filePath)
	if err != nil {
		return sum, err
	}
	defer func() {
		derr := f.Close()
		if derr != nil {
			err = ae.Combine(err, derr)
		}
	}()

	h := crc32.New(crc32.IEEETable)

	var n int
	var buf = make([]byte, ReadBufferSize)
	var readErr, writeErr error

	for {
		n, readErr = f.Read(buf)
		if readErr == nil {
			_, writeErr = h.Write(buf[:n])
			if writeErr != nil {
				return sum, writeErr
			}
			continue
		}

		if readErr == io.EOF {
			_, writeErr = h.Write(buf[:n])
			if writeErr != nil {
				return sum, writeErr
			}
			break
		}

		return sum, readErr
	}

	return h.Sum32(), nil
}

func GetFileHashMD5(filePath string) (sum []byte, err error) {
	var f *os.File
	f, err = os.Open(filePath)
	if err != nil {
		return sum, err
	}
	defer func() {
		derr := f.Close()
		if derr != nil {
			err = ae.Combine(err, derr)
		}
	}()

	h := md5.New()

	var n int
	var buf = make([]byte, ReadBufferSize)
	var readErr, writeErr error

	for {
		n, readErr = f.Read(buf)
		if readErr == nil {
			_, writeErr = h.Write(buf[:n])
			if writeErr != nil {
				return sum, writeErr
			}
			continue
		}

		if readErr == io.EOF {
			_, writeErr = h.Write(buf[:n])
			if writeErr != nil {
				return sum, writeErr
			}
			break
		}

		return sum, readErr
	}

	return h.Sum(nil), nil
}

func GetFileHashSHA256(filePath string) (sum []byte, err error) {
	var f *os.File
	f, err = os.Open(filePath)
	if err != nil {
		return sum, err
	}
	defer func() {
		derr := f.Close()
		if derr != nil {
			err = ae.Combine(err, derr)
		}
	}()

	h := sha256.New()

	var n int
	var buf = make([]byte, ReadBufferSize)
	var readErr, writeErr error

	for {
		n, readErr = f.Read(buf)
		if readErr == nil {
			_, writeErr = h.Write(buf[:n])
			if writeErr != nil {
				return sum, writeErr
			}
			continue
		}

		if readErr == io.EOF {
			_, writeErr = h.Write(buf[:n])
			if writeErr != nil {
				return sum, writeErr
			}
			break
		}

		return sum, readErr
	}

	return h.Sum(nil), nil
}

func GetFileHashFileSize(filePath string) (sum int64, err error) {
	var fi os.FileInfo
	fi, err = os.Stat(filePath)
	if err != nil {
		return sum, err
	}

	return fi.Size(), nil
}

func GetFileHashFileExistence(filePath string) (sum bool, err error) {
	var exists bool
	exists, err = file.FileExists(filePath)
	if err != nil {
		return false, err
	}

	return exists, nil
}
