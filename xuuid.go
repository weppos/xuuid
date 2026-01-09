package main

import (
	"encoding/hex"
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/google/uuid"
)

func main() {
	// Define flags
	v7Flag := flag.Bool("v7", false, "Generate UUID v7 (default)")
	v6Flag := flag.Bool("v6", false, "Generate UUID v6")
	v4Flag := flag.Bool("v4", false, "Generate UUID v4")
	md5Flag := flag.Bool("md5", false, "Generate MD5-based UUID")
	sha1Flag := flag.Bool("sha1", false, "Generate SHA1-based UUID")
	countFlag := flag.Int("count", 1, "Number of UUIDs to generate")
	hexFlag := flag.Bool("hex", false, "Output UUIDs as 32-char hex (no dashes)")
	urnFlag := flag.Bool("urn", false, "Output UUIDs in urn:uuid: format")

	// Custom usage banner
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "xuuid - UUID Generator\n\n")
		fmt.Fprintf(os.Stderr, "Usage:\n")
		fmt.Fprintf(os.Stderr, "  xuuid [flags]\n\n")
		fmt.Fprintf(os.Stderr, "Flags:\n\n")
		flag.PrintDefaults()
		fmt.Fprintf(os.Stderr, "\nExamples:\n\n")
		fmt.Fprintf(os.Stderr, "  xuuid                 Generate a single UUID v7 (default)\n")
		fmt.Fprintf(os.Stderr, "  xuuid --v4            Generate a UUID v4\n")
		fmt.Fprintf(os.Stderr, "  xuuid --count 5       Generate 5 UUID v7s\n")
		fmt.Fprintf(os.Stderr, "  xuuid --v4 --count 3  Generate 3 UUID v4s\n")
	}

	flag.Parse()

	// Determine which UUID type to generate
	var generateFunc func() (uuid.UUID, error)

	switch {
	case *v6Flag:
		generateFunc = uuid.NewV6
	case *v4Flag:
		generateFunc = func() (uuid.UUID, error) {
			return uuid.New(), nil
		}
	case *md5Flag:
		generateFunc = func() (uuid.UUID, error) {
			name := fmt.Sprintf("%d", time.Now().UnixNano())
			return uuid.NewMD5(uuid.NameSpaceDNS, []byte(name)), nil
		}
	case *sha1Flag:
		generateFunc = func() (uuid.UUID, error) {
			name := fmt.Sprintf("%d", time.Now().UnixNano())
			return uuid.NewSHA1(uuid.NameSpaceDNS, []byte(name)), nil
		}
	case *v7Flag:
		generateFunc = uuid.NewV7
	default:
		// Default to v7
		generateFunc = uuid.NewV7
	}

	// Generate UUID(s)
	for i := 0; i < *countFlag; i++ {
		id, err := generateFunc()
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error generating UUID: %v\n", err)
			os.Exit(1)
		}
		if *urnFlag {
			fmt.Println(id.URN())
		} else if *hexFlag {
			fmt.Println(hex.EncodeToString(id[:]))
		} else {
			fmt.Println(id.String())
		}
	}
}
