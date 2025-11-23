package h

import (
	"os"

	ht "github.com/vault-thirteen/Hasher/pkg/Models/HashType"
	"github.com/vault-thirteen/auxie/errors"
	af "github.com/vault-thirteen/auxie/file"
	ah "github.com/vault-thirteen/auxie/hash"
)

func calculateBinaryFileHash(filePath string, hti ht.HashTypeId) (sum []byte, err error) {
	var file *os.File
	file, err = os.Open(filePath)
	if err != nil {
		return nil, err
	}

	defer func() {
		derr := file.Close()
		if derr != nil {
			err = errors.Combine(err, derr)
		}
	}()

	switch hti {
	case ht.Id_CRC32:
		{
			var x ah.Crc32Sum
			x, _, err = ah.CalculateCrc32S(file)
			if err != nil {
				return nil, err
			}

			sum = x[:]
		}

	case ht.Id_MD5:
		{
			var x ah.Md5Sum
			x, _, err = ah.CalculateMd5S(file)
			if err != nil {
				return nil, err
			}

			sum = x[:]
		}

	case ht.Id_SHA1:
		{
			var x ah.Sha1Sum
			x, _, err = ah.CalculateSha1S(file)
			if err != nil {
				return nil, err
			}

			sum = x[:]
		}

	case ht.Id_SHA256:
		{
			var x ah.Sha256Sum
			x, _, err = ah.CalculateSha256S(file)
			if err != nil {
				return nil, err
			}

			sum = x[:]
		}
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
