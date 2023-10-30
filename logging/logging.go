package logging

import "github.com/go-logr/logr"

var Logger = logr.Discard()

func SetLogger(l logr.Logger) {
	Logger = l
}
