package main

import (
	"bufio"
	"encoding/base64"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

type flags struct {
	decode       bool
	verbose      bool
	printVersion bool
}

var cmd flags
var Version string = "0.0.1"
var Rev string = ""

func main() {
	flag.Usage = printHelp

	parseCommandLine()

	if cmd.printVersion {
		printVersion()
		os.Exit(0)
	}

	input := readInput()

	if cmd.decode {
		output, err := base64.RawURLEncoding.DecodeString(input)
		checkError(err)
		fmt.Print(string(output))
	} else {
		output := base64.RawURLEncoding.EncodeToString([]byte(input))
		fmt.Print(output)
	}
}

func checkError(err error) {
	if err != nil {
		if cmd.verbose {
			fmt.Print("Error: ")
			fmt.Println(err.Error())
		}
		os.Exit(1)
	}
}

func readInput() string {
	args := flag.Args()
	input := ""
	if len(args) > 1 {
		fmt.Println("too many arguments")
		os.Exit(1)
	} else if len(args) == 0 || args[0] == "-" {
		reader := bufio.NewReader(os.Stdin)
		var err error
		input, err = reader.ReadString('\n')
		checkError(err)
	} else {
		file, err := ioutil.ReadFile(args[0])
		checkError(err)
		input = string(file)
	}
	return strings.Trim(input, " \r\n")
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
	fmt.Println("Version " + Version + " " + Rev)
}
