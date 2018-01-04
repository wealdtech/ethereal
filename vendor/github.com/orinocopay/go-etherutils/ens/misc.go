package ens

import "strings"

// DomainLevel calculates the level of the domain presented.
// A top-level domain (e.g. 'eth') will be 0, a domain (e.g.
// 'foo.eth') will be 1, a subdomain (e.g. 'bar.foo.eth' will
// be 2, etc.
func DomainLevel(name string) (level int) {
	return len(strings.Split(name, ".")) - 1
}
