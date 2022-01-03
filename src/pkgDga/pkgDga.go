package pkgDga

import (
	"bytes"
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base32"
	"encoding/binary"
	"fmt"
	"math"
	"strconv"
	"strings"
	"time"
)

const ASCII_START = 97
const POW = 25
const DIGITS = 10

//export GenerateRandomKey
func GenerateRandomKey(length int) string {
	if length < 1 {
		length = 10
	}

	token := make([]byte, length)
	rand.Read(token)
	key := base32.StdEncoding.EncodeToString(token)

	return key
}

//export GenerateDomain
func GenerateDomain(token string, unixSeconds int64, formatPattern string, count int) []string {
	securityToken, err := base32.StdEncoding.DecodeString(token)

	if err != nil {
		fmt.Println(err.Error())
		return nil
	}

	var domains []string

	if count < 1 {
		count = 1
	}

	dateTimeOffset := time.Unix(unixSeconds, 0).UTC()
	year := fmt.Sprintf("%04d", dateTimeOffset.Year())
	month := fmt.Sprintf("%02d", dateTimeOffset.Month())
	day := fmt.Sprintf("%02d", dateTimeOffset.Day())
	hour := fmt.Sprintf("%02d", dateTimeOffset.Hour())
	minute := fmt.Sprintf("%02d", dateTimeOffset.Minute())
	second := fmt.Sprintf("%02d", dateTimeOffset.Second())

	pattern := strings.ReplaceAll(formatPattern, "yyyy", year)
	pattern = strings.ReplaceAll(pattern, "yy", year[2:])
	pattern = strings.ReplaceAll(pattern, "MM", month)
	pattern = strings.ReplaceAll(pattern, "M", strings.TrimLeft(month, "0"))
	pattern = strings.ReplaceAll(pattern, "dd", day)
	pattern = strings.ReplaceAll(pattern, "d", strings.TrimLeft(day, "0"))
	pattern = strings.ReplaceAll(pattern, "HH", hour)
	pattern = strings.ReplaceAll(pattern, "H", strings.TrimLeft(hour, "0"))
	pattern = strings.ReplaceAll(pattern, "hh", hour)
	pattern = strings.ReplaceAll(pattern, "h", strings.TrimLeft(hour, "0"))
	pattern = strings.ReplaceAll(pattern, "mm", minute)
	pattern = strings.ReplaceAll(pattern, "m", strings.TrimLeft(minute, "0"))
	pattern = strings.ReplaceAll(pattern, "ss", second)
	pattern = strings.ReplaceAll(pattern, "s", strings.TrimLeft(second, "0"))

	for i := 0; i < count; i++ {
		code, err := computeCode(securityToken, pattern+"-"+strconv.Itoa(i))

		if err != nil {
			fmt.Println(err.Error())
			return nil
		}

		domain := ""
		for code > 0 {
			div, rem := code/POW, code%POW
			code = div
			character := string(rune(rem + ASCII_START))
			domain += character
		}

		domains = append(domains, domain)
	}

	return domains
}

func computeCode(securityToken []byte, dateTimePattern string) (int64, error) {
	mod := int64(math.Pow(POW, DIGITS))
	mac := hmac.New(sha256.New, securityToken)
	mac.Write([]byte(dateTimePattern))
	hash := mac.Sum(nil)
	binaryCode := bytesToInt(hash)
	rem := float64(binaryCode % mod)
	result := int64(math.Abs(rem))
	return result, nil
}

func intToBytes(num int64) []byte {
	data := int64(num)
	buffers := bytes.NewBuffer([]byte{})
	binary.Write(buffers, binary.BigEndian, data)
	return buffers.Bytes()
}

func bytesToInt(arr []byte) int64 {
	bytebuff := bytes.NewBuffer(arr)
	var data int64
	binary.Read(bytebuff, binary.BigEndian, &data)
	return data
}
