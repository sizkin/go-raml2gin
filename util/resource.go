package grgutil

import (
    "regexp"
)

var (
    regRAMLParam = regexp.MustCompile(`{(\w+)}`)
)

func ToGinResource(resource string) string {
	return regRAMLParam.ReplaceAllString(resource, ":$1")
}
