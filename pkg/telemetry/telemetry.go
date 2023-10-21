/*
Copyright (C) 2021-2023, Kubefirst

This program is licensed under MIT.
See the LICENSE file for more details.
*/
package telemetry

import (
	"strings"
	"time"

	"github.com/kubefirst/metrics-client/pkg/segment"
)



func RemoveSubdomainV2(domainName string) (string, error) {

	domainName = strings.TrimRight(domainName, ".")
	domainSlice := strings.Split(domainName, ".")

	if len(domainSlice) < 2 {
		return "", nil
	}

	domainName = strings.Join([]string{domainSlice[len(domainSlice)-2], domainSlice[len(domainSlice)-1]}, ".")

	return domainName, nil
}
