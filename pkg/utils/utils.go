package utils

import "strings"

func RemoveSubdomainV2(domainName string) (string, error) {

	domainName = strings.TrimRight(domainName, ".")
	domainSlice := strings.Split(domainName, ".")

	if len(domainSlice) < 2 {
		return "", nil
	}

	domainName = strings.Join([]string{domainSlice[len(domainSlice)-2], domainSlice[len(domainSlice)-1]}, ".")

	return domainName, nil
}
