package main

import (
	"bytes"
	"strings"
	"testing"
)

func TestRun_tokenFlag(t *testing.T) {
	outStream, errStream := new(bytes.Buffer), new(bytes.Buffer)
	cli := &CLI{outStream: outStream, errStream: errStream}
	args := strings.Split("./sacloud-upload-image -token", " ")

	status := cli.Run(args)
	_ = status
}

func TestRun_secretFlag(t *testing.T) {
	outStream, errStream := new(bytes.Buffer), new(bytes.Buffer)
	cli := &CLI{outStream: outStream, errStream: errStream}
	args := strings.Split("./sacloud-upload-image -secret", " ")

	status := cli.Run(args)
	_ = status
}

func TestRun_zoneFlag(t *testing.T) {
	outStream, errStream := new(bytes.Buffer), new(bytes.Buffer)
	cli := &CLI{outStream: outStream, errStream: errStream}
	args := strings.Split("./sacloud-upload-image -zone", " ")

	status := cli.Run(args)
	_ = status
}
