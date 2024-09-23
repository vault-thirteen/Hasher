package h

import (
	"os"

	ht "github.com/vault-thirteen/Hasher/pkg/Models/HashType"
	af "github.com/vault-thirteen/auxie/file"
	ah "github.com/vault-thirteen/auxie/hash"
)

func calculateBinaryFileHash(filePath string, hti ht.HashTypeId) (sum []byte, err error) {
	var data []byte
	data, err = af.GetFileContents(filePath)
	if err != nil {
		return nil, err
	}

	switch hti {
	case ht.Id_CRC32:
		x, _ := ah.CalculateCrc32(data)
		sum = x[:]

	case ht.Id_MD5:
		x, _ := ah.CalculateMd5(data)
		sum = x[:]

	case ht.Id_SHA256:
		x, _ := ah.CalculateSha256(data)
		sum = x[:]
	}

	return sum, nil
}

func checkFileExistence(filePath string) (fileExists bool, err error) {
	var exists bool
	exists, err = af.FileExists(filePath)
	if err != nil {
		return false, err
	}

	return exists, nil
}

func getFileSize(filePath string) (fileSize int, err error) {
	var fi os.FileInfo
	fi, err = os.Stat(filePath)
	if err != nil {
		return -1, err
	}

	x := fi.Size()
	fileSize = int(x)

	return fileSize, nil
}
