package ens

import (
	"fmt"
	"strings"
)

// DomainLevel calculates the level of the domain presented.
// A top-level domain (e.g. 'eth') will be 0, a domain (e.g.
// 'foo.eth') will be 1, a subdomain (e.g. 'bar.foo.eth' will
// be 2, etc.
func DomainLevel(name string) (level int) {
	return len(strings.Split(name, ".")) - 1
}

// NormaliseDomain turns ENS domain in to normal form
func NormaliseDomain(domain string) string {
	wildcard := false
	if strings.HasPrefix(domain, "*.") {
		wildcard = true
		domain = domain[2:]
	}
	output, err := p.ToUnicode(strings.ToLower(domain))
	if err != nil {
		panic("ENS domain normalisation failed")
	}

	// ToUnicode() removes leading periods.  Replace them
	if strings.HasPrefix(domain, ".") && !strings.HasPrefix(output, ".") {
		output = "." + output
	}

	// If we removed a wildcard then add it back
	if wildcard {
		output = "*." + output
	}
	return output
}

// Tld obtains the top-level domain of an ENS name
func Tld(domain string) string {
	domain = NormaliseDomain(domain)
	tld, err := DomainPart(domain, -1)
	if err != nil {
		return domain
	}
	return tld
}

// DomainPart obtains a part of a name
// Positive parts start at the lowest-level of the domain and work towards the
// top-level domain.  Negative parts start at the top-level domain and work
// towards the lowest-level domain.
// For example, with a domain bar.foo.com the following parts will be returned:
// Number | part
//      1 |  bar
//      2 |  foo
//      3 |  com
//     -1 |  com
//     -2 |  foo
//     -3 |  bar
func DomainPart(domain string, part int) (string, error) {
	if part == 0 {
		return "", fmt.Errorf("Invalid part")
	}
	domain = NormaliseDomain(domain)
	parts := strings.Split(domain, ".")
	if len(parts) < abs(part) {
		return "", fmt.Errorf("Not enough parts")
	}
	if part < 0 {
		return parts[len(parts)+part], nil
	}
	return parts[part-1], nil
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
