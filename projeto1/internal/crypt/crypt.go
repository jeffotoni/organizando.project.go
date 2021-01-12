package crypt

import(
	"fmt"
	"hash/crc32"
	"math/rand"
	"time"
	"crypto/md5"
	"io"
	"strconv"
)

func RandowNew() string{
	rand.Seed(time.Now().UnixNano())
	return strconv.Itoa(rand.Intn(1000000) + rand.Intn(100000))
}

func Crc32() string{
	crc32q := crc32.MakeTable(0xD5828281)
	return fmt.Sprintf("%08x", crc32.Checksum([]byte(RandowNew()), crc32q))
}

func Md5(text string) string{
	h := md5.New()
	io.WriteString(h, text)
	return (fmt.Sprintf("%x", h.Sum(nil)))
}

