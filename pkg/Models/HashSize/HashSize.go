package hs

type HashSize int

const (
	HashSize_None   = HashSize(0)
	HashSize_CRC32  = HashSize(4)
	HashSize_MD5    = HashSize(16)
	HashSize_SHA1   = HashSize(20)
	HashSize_SHA256 = HashSize(32)
)
