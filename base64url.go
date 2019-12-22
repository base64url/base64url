package main

import (
	"bufio"
	"encoding/base64"
	"flag"
	"fmt"
	"io"
	"os"
)

type flags struct {
	decode       bool
	verbose      bool
	printVersion bool
}

var cmd flags
var version string = "0.0.1"
var rev string = ""

func decode(inStream io.Reader, outStream io.Writer) error {
	// Manually remove newline chars \r and \n before decoding - otherwise the decoder will throw an error
	remover := newNewlineRemovingReader(inStream)
	decoder := base64.NewDecoder(base64.RawURLEncoding, remover)
	_, err := io.Copy(outStream, decoder)
	return err
}

func encode(inStream io.Reader, outStream io.Writer) error {
	scanner := bufio.NewScanner(inStream)
	scanner.Split(bufio.ScanLines)
	encoder := base64.NewEncoder(base64.RawURLEncoding, outStream)
	for scanner.Scan() {
		_, err := encoder.Write(scanner.Bytes())
		checkError(err)
	}
	encoder.Close()
	return scanner.Err()
}

func main() {
	flag.Usage = printHelp

	parseCommandLine()

	if cmd.printVersion {
		printVersion()
		os.Exit(0)
	}

	input := getInputReader()

	if cmd.decode {
		err := decode(input, os.Stdout)
		checkError(err)
	} else {
		err := encode(input, os.Stdout)
		checkError(err)
	}
}

func checkError(err error) {
	if err != nil {
		if cmd.verbose {
			fmt.Print("Error: ")
			fmt.Fprintln(os.Stderr, err.Error())
		}
		os.Exit(1)
	}
}

func getInputReader() io.Reader {
	args := flag.Args()
	if len(args) > 1 {
		fmt.Println("too many arguments")
		os.Exit(1)
	} else if len(args) == 0 || args[0] == "-" {
		return bufio.NewReader(os.Stdin)
	} else {
		reader, err := os.Open(args[0])
		checkError(err)
		return reader
	}
	return nil
}

func parseCommandLine() {
	decodePtr := flag.Bool("d", false, "Decode Data")
	flag.BoolVar(decodePtr, "decode", false, "Decode Data")

	verbosePtr := flag.Bool("v", false, "Print verbose messages")

	versionPtr := flag.Bool("version", false, "Print Version Number")

	flag.Parse()

	cmd = flags{*decodePtr, *verbosePtr, *versionPtr}
}

func printHelp() {
	printVersion()
	flag.PrintDefaults()
}

func printVersion() {
	fmt.Println("base64url - base64 encode/decode data and print to standard output using the URL and Filename Safe Alphabet.")
	fmt.Println("Version " + version + " " + rev)
}
