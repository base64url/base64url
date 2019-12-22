package main

import (
	"bytes"
	"encoding/hex"
	"io"
	"strings"
	"testing"
)

func prepareStreams(input string) (io.Reader, *bytes.Buffer) {
	reader := strings.NewReader(input)

	writer := new(bytes.Buffer)
	return reader, writer
}

func checkOutput(t *testing.T, expected string, buffer *bytes.Buffer, err error) {
	if err != nil {
		t.Fatal(err)
	}
	actual := buffer.String()
	if actual != expected {
		t.Logf("actual:\n%s\n", hex.Dump([]byte(actual)))
		t.Logf("expected:\n%s\n", hex.Dump([]byte(expected)))
		t.Fatal("Actual output is not equal to expected")
	}
}

func checkNewlineRemoval(t *testing.T, input string, expected string) {
	r, w := prepareStreams(input)

	reader := newNewlineRemovingReader(r)
	_, err := io.Copy(w, reader)

	checkOutput(t, expected, w, err)
}

func TestEncode(t *testing.T) {
	r, w := prepareStreams("Hello, World")

	output := "SGVsbG8sIFdvcmxk"
	err := encode(r, w)

	checkOutput(t, output, w, err)
}

func TestDecode(t *testing.T) {
	r, w := prepareStreams("SGVsbG8sIFdvcmxk")

	output := "Hello, World"
	err := decode(r, w)

	checkOutput(t, output, w, err)
}

func TestDecodeWithNewLine(t *testing.T) {
	r, w := prepareStreams("SG\nVsb\rG8sIFd\r\nvcmxk")

	output := "Hello, World"
	err := decode(r, w)

	checkOutput(t, output, w, err)
}

func TestNewlineRemovingReader(t *testing.T) {
	checkNewlineRemoval(t,
		"A \rpiece of text with \nvarious \r\ntypes of new lines",
		"A piece of text with various types of new lines")
}

func TestCyrillicLatinNewlineRemoval(t *testing.T) {
	checkNewlineRemoval(t,
		"Лорем ипсум долор \rсит амет, \nелигенди ехпетенда \r\nиус еа ан",
		"Лорем ипсум долор сит амет, елигенди ехпетенда иус еа ан")
}

func TestArmenianLatinNewlineRemoval(t *testing.T) {
	checkNewlineRemoval(t,
		"լոռեմ իպսում դոլոռ \nսիթ ամեթ, \rսանծթուս ծոնսթիթութո \r\nթե հաս ծում",
		"լոռեմ իպսում դոլոռ սիթ ամեթ, սանծթուս ծոնսթիթութո թե հաս ծում")
}

func TestHindiLatinNewlineRemoval(t *testing.T) {
	checkNewlineRemoval(t,
		"संस्था वेबजाल विभाजनक्षमता \rकरता। अन्य \nप्राण आपके गुजरना \r\nआंतरकार्यक्षमता सके।",
		"संस्था वेबजाल विभाजनक्षमता करता। अन्य प्राण आपके गुजरना आंतरकार्यक्षमता सके।")
}
