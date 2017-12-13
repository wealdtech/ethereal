package ens

import (
	"strings"
)

// DomainLevel calculates the level of the domain presented.
// A top-level domain (e.g. 'eth') will be 0, a domain (e.g.
// 'foo.eth') will be 1, a subdomain (e.g. 'bar.foo.eth' will
// be 2, etc.
func DomainLevel(name string) (level int) {
	return len(strings.Split(name, ".")) - 1
}

// Normalise an ENS domain
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

// Obtain the TLD of an ENS domain
func Tld(domain string) string {
	domain = NormaliseDomain(domain)
	lastPeriodLoc := strings.LastIndex(domain, ".")
	if lastPeriodLoc == -1 {
		return domain
	} else if lastPeriodLoc == len(domain) {
		return ""
	} else {
		return domain[lastPeriodLoc+1:]
	}
}
