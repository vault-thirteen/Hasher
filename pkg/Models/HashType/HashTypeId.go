package ht

type HashTypeId byte

const (
	Id_FileExistence = HashTypeId(1)
	Id_FileSize      = HashTypeId(2)
	Id_CRC32         = HashTypeId(3)
	Id_MD5           = HashTypeId(4)
	Id_SHA256        = HashTypeId(5)
)
