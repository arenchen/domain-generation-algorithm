package main

import "C"

import (
	"dga/src/pkgDga"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	argc := len(os.Args)
	key := pkgDga.GenerateRandomKey(10)
	format := "yyyyMMddHHmm"
	now := time.Now().UTC()
	unixtime := now.Unix()
	count := 10
	suffix := ".com"

	if argc == 1 {
		fmt.Println("Usage: dga [options]")
		fmt.Println("Options:")
		fmt.Println("  -k\tKey, Base32 String")
		fmt.Println("  -t\tUnix-Time Seconds, Default: Now")
		fmt.Println("  -c\tGenerator Count, Default: 10")
		fmt.Println("  -f\tDate Format Pattern, Default: yyyyMMddHHmm")
		fmt.Println("  -s\tDomain Suffix, Default: .com")
		fmt.Println("Example:")
		fmt.Printf("  dga -k %s -t %d\n", key, unixtime)
		return
	}

	for i := 1; i < argc; i += 2 {
		fmt.Println(os.Args[i])
		if strings.EqualFold(os.Args[i], "-k") {
			key = strings.TrimSpace(os.Args[i+1])
		} else if strings.EqualFold(os.Args[i], "-t") {
			num, err := strconv.ParseInt(os.Args[i+1], 10, 64)
			if err != nil {
				fmt.Println(err.Error())
				return
			}
			unixtime = num
		} else if strings.EqualFold(os.Args[i], "-c") {
			num, err := strconv.ParseInt(os.Args[i+1], 10, 64)
			if err != nil {
				fmt.Println(err.Error())
				return
			}
			count = int(num)
		} else if strings.EqualFold(os.Args[i], "-f") {
			format = strings.TrimSpace(os.Args[i+1])
		} else if strings.EqualFold(os.Args[i], "-s") {
			suffix = strings.TrimSpace(os.Args[i+1])
		}
	}

	domains := pkgDga.GenerateDomain(key, unixtime, format, count)

	fmt.Printf("Key: %s\n", key)
	fmt.Printf("Unix-Time Seconds: %d\n", now.Unix())
	fmt.Printf("DateTime: %s\n", time.Unix(unixtime, 0).UTC())
	fmt.Printf("Format: %s\n", format)
	fmt.Printf("Count: %d\n", count)
	fmt.Printf("Suffix: %s\n", suffix)
	fmt.Printf("Domains:")

	for i := 0; i < count; i++ {
		fmt.Printf(" %s%s", domains[i], suffix)
	}

	fmt.Println("")
}
