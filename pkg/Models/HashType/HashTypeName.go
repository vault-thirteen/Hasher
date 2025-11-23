package ht

type HashTypeName string

const (
	Name_FileExistence = HashTypeName("EXISTENCE")
	Name_FileSize      = HashTypeName("SIZE")
	Name_CRC32         = HashTypeName("CRC32")
	Name_MD5           = HashTypeName("MD5")
	Name_SHA1          = HashTypeName("SHA1")
	Name_SHA256        = HashTypeName("SHA256")
)
