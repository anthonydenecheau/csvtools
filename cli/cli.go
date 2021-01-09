package cli

import (
	"flag"
	"fmt"
	"log"
	"os"
	"runtime"

	"github.com/anthonydenecheau/golang-scc/extract"
)

type CommandLine struct{}

func (cli *CommandLine) printUsage() {
	fmt.Println("Usage:")
	fmt.Println(" readCSAU -directory DIRECTORY - parse [tc] files")
}

func (cli *CommandLine) validateArgs() {
	if len(os.Args) < 2 {
		cli.printUsage()
		runtime.Goexit()
	}
}

func (cli *CommandLine) readCSAU(directory string) {
	fmt.Printf("readCSAU of %s\n", directory)
	var files []string
	scan := extract.ExtractCmd{}
	files, err := scan.FindCSV("csau", directory)

	if err != nil {
		log.Panic(err)
	}

	var lines []extract.LineCsau
	lines, _err := scan.ParseCsauCSV(files)

	if _err != nil {
		log.Panic(_err)
	}

	for _, line := range lines {
		fmt.Println(">", line.DtEpreuve, line.Organisateur, line.TatooChip)
	}

}

func (cli *CommandLine) readTC(directory string) {
	fmt.Printf("readTC of %s\n", directory)
	var files []string
	scan := extract.ExtractCmd{}
	files, err := scan.FindCSV("tc", directory)

	if err != nil {
		log.Panic(err)
	}

	for _, file := range files {
		log.Printf(file)
	}

	var lines []extract.LineTc
	lines, _err := scan.ParseTcCSV(files)

	if _err != nil {
		log.Panic(_err)
	}

	for _, line := range lines {
		fmt.Println(">", line.DtEpreuve, line.Organisateur, line.TatooChip)
	}

}

func (cli *CommandLine) Run() {
	cli.validateArgs()

	readCsauCmd := flag.NewFlagSet("readCSAU", flag.ExitOnError)
	readCSAU := readCsauCmd.String("directory", ".", "The directory to parse")

	readTcCmd := flag.NewFlagSet("readTC", flag.ExitOnError)
	readTC := readTcCmd.String("directory", ".", "The directory to parse")

	switch os.Args[1] {
	case "readCSAU":
		err := readCsauCmd.Parse(os.Args[2:])
		if err != nil {
			log.Panic(err)
		}
	case "readTC":
		err := readTcCmd.Parse(os.Args[2:])
		if err != nil {
			log.Panic(err)
		}
	default:
		cli.printUsage()
		runtime.Goexit()
	}

	if readCsauCmd.Parsed() {
		if *readCSAU == "" {
			readCsauCmd.Usage()
			runtime.Goexit()
		}
		cli.readCSAU(*readCSAU)
	}

	if readTcCmd.Parsed() {
		if *readTC == "" {
			readTcCmd.Usage()
			runtime.Goexit()
		}
		cli.readTC(*readTC)
	}

}
