package main

import (
	"testing"
)

func TestNone(t *testing.T) {
	t.Log("current version is ", Version)
	Version = "1.0.1"
	t.Log("current version is ", Version)
}

func TestV2(t *testing.T) {
	t.Log("current version is ", Version)
	Version = "2.0.1"
	t.Log("current version is ", Version)
}
