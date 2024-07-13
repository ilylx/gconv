package gdebug

import (
	"github.com/ilylx/gconv/internal/crypto/gmd5"
	"github.com/ilylx/gconv/internal/encoding/ghash"
	"io/ioutil"
	"strconv"
)

// BinVersion returns the version of current running binary.
// It uses ghash.BKDRHash+BASE36 algorithm to calculate the unique version of the binary.
func BinVersion() string {
	if binaryVersion == "" {
		binaryContent, _ := ioutil.ReadFile(selfPath)
		binaryVersion = strconv.FormatInt(
			int64(ghash.BKDRHash(binaryContent)),
			36,
		)
	}
	return binaryVersion
}

// BinVersionMd5 returns the version of current running binary.
// It uses MD5 algorithm to calculate the unique version of the binary.
func BinVersionMd5() string {
	if binaryVersionMd5 == "" {
		binaryVersionMd5, _ = gmd5.EncryptFile(selfPath)
	}
	return binaryVersionMd5
}
