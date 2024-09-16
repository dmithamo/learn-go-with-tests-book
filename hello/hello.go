package main

import (
	"fmt"
)

const (
	ENG_PREFIX  = "Hello,"
	KIUK_PREFIX = "Geithika,"
	SPA_PREFIX  = "Hola,"
)

func Hello(lang, name string) string {
	if name == "" {
		name = "world"
	}
	return fmt.Sprintf("%s %s!", greetingPrefix(lang), name)
}

func greetingPrefix(lang string) (prefix string) {
	switch lang {
	case "ENG":
		prefix = ENG_PREFIX

	case "KIUK":
		prefix = KIUK_PREFIX

	case "SPA":
		prefix = SPA_PREFIX

	default:
		prefix = "Hiya,"
	}
	return
}
