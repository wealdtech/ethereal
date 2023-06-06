package util

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"golang.org/x/crypto/sha3"
)

// DNSDomainHash hashes a domain name.
func DNSDomainHash(domain string) [32]byte {
	lower := strings.ToLower(domain)
	sha := sha3.NewLegacyKeccak256()
	_, err := sha.Write([]byte(lower))
	if err != nil {
		panic(err)
	}
	var hash [32]byte
	sha.Sum(hash[:0])
	return hash
}

// DNSWireFormatDomainHash hashes a domain name in wire format.
func DNSWireFormatDomainHash(domain string) [32]byte {
	sha := sha3.NewLegacyKeccak256()
	_, err := sha.Write(DNSWireFormat(domain))
	if err != nil {
		panic(err)
	}
	var hash [32]byte
	sha.Sum(hash[:0])
	return hash
}

// DNSWireFormat turns a domain name in to wire format.
func DNSWireFormat(domain string) []byte {
	// Remove leading and trailing dots.
	domain = strings.TrimLeft(domain, ".")
	domain = strings.TrimRight(domain, ".")
	domain = strings.ToLower(domain)

	if domain == "" {
		return []byte{0x00}
	}

	bytes := make([]byte, len(domain)+2)
	pieces := strings.Split(domain, ".")
	offset := 0
	for _, piece := range pieces {
		bytes[offset] = byte(len(piece))
		offset++
		copy(bytes[offset:offset+len(piece)], piece)
		offset += len(piece)
	}
	return bytes
}

// IncrementSerial increments a SOA serial number as per RFC 1912.
func IncrementSerial(serial uint32) uint32 {
	strSerial := fmt.Sprintf("%d", serial)
	datePart := time.Now().Format("20060102")

	if len(strSerial) < 10 {
		// Non-standard format; change to RFC 1912.
		strSerial = fmt.Sprintf("%s00", datePart)
		result, _ := strconv.ParseInt(strSerial, 10, 32)
		return uint32(result)
	}

	// Standard format.
	var nn int
	if datePart == strSerial[:8] {
		// Same day; increment nn.
		nn, _ = strconv.Atoi(strSerial[8:])
		nn++
	} else {
		// Different day; nn to 0.
		nn = 0
	}
	strSerial = fmt.Sprintf("%s%02d", datePart, nn)
	result, _ := strconv.ParseInt(strSerial, 10, 32)

	if uint32(result) > serial {
		// This will be the case if the format was already RFC 1912.
		return uint32(result)
	}

	// If we reach here it means that the serial number given to us is higher
	// than the current date.  We cannot set it in to RFC 1912 format without
	// confusing the nameservers so just increment it.
	return serial + 1
}
