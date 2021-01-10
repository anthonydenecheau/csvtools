package cli

import (
	"flag"
	"fmt"
	"log"
	"os"
	"runtime"

	"github.com/anthonydenecheau/csvtools/extract"
)

type CommandLine struct{}

func (cli *CommandLine) printUsage() {
	fmt.Println("Usage:")
	fmt.Println(" [command] -directory DIRECTORY - consume csv files ")
	fmt.Println(" choose command : readCSAU / readTC ")
}

func (cli *CommandLine) validateArgs() {
	if len(os.Args) < 2 {
		cli.printUsage()
		runtime.Goexit()
	}
}

func (cli *CommandLine) find(directory string, prefix string) {
	fmt.Printf("find csv file start with %s\n", prefix)
	var files []string
	scan := extract.ExtractCmd{}
	files, err := scan.FindCSV(prefix, directory)

	if err != nil {
		log.Panic(err)
	}

	for _, file := range files {
		fmt.Println(">", file)
	}

}

func (cli *CommandLine) read(directory string, prefix string) {
	fmt.Printf("read csv file start with %s\n", prefix)
	var files []string
	scan := extract.ExtractCmd{}
	files, err := scan.FindCSV(prefix, directory)

	if err != nil {
		log.Panic(err)
	}

	switch prefix {
	case "csau":
		var lines []extract.LineCsau
		lines, _err := scan.ParseCsauCSV(files)
		if _err != nil {
			log.Panic(_err)
		}
		for _, line := range lines {
			fmt.Println(">", line.DtEpreuve, line.Organisateur, line.TatooChip)
		}
	case "tc":
		var lines []extract.LineTc
		lines, _err := scan.ParseTcCSV(files)
		if _err != nil {
			log.Panic(_err)
		}
		for _, line := range lines {
			fmt.Println(">", line.DtEpreuve, line.Organisateur, line.TatooChip)
		}
	default:
		log.Panic("Couldn't recognized prefix")
	}

}

func (cli *CommandLine) process(directory string, prefix string, out string) {
	fmt.Printf("process csv file start with %s\n", prefix)
	var files []string
	scan := extract.ExtractCmd{}
	files, err := scan.FindCSV(prefix, directory)

	if err != nil {
		log.Panic(err)
	}

	lines, err := scan.ParseCSV(prefix, files)

	if err != nil {
		log.Panic(err)
	}

	for _, line := range lines {
		if csau, ok := line.([]extract.LineCsau); ok {
			for _, innerCsau := range csau {
				log.Printf("%#v", innerCsau)
			}
		}
		if tc, ok := line.([]extract.LineTc); ok {
			for _, innerTc := range tc {
				log.Printf("%#v", innerTc)
			}
		}
	}
	fmt.Printf("process csv %s\n", prefix)

}

func (cli *CommandLine) Run() {
	cli.validateArgs()

	findCmd := flag.NewFlagSet("find", flag.ExitOnError)
	readCmd := flag.NewFlagSet("read", flag.ExitOnError)
	processCmd := flag.NewFlagSet("process", flag.ExitOnError)

	findFrom := findCmd.String("directory", ".", "The directory to parse")
	findPrefix := findCmd.String("prefix", "", "The prefix of the filename")
	readFrom := readCmd.String("directory", ".", "The directory to parse")
	readPrefix := readCmd.String("prefix", "", "The prefix of the filename")
	processFrom := processCmd.String("directory", ".", "The directory to parse")
	processPrefix := processCmd.String("prefix", "", "The prefix of the filename")
	processOut := processCmd.String("out", "", "The filename of results")
	switch os.Args[1] {
	case "find":
		err := findCmd.Parse(os.Args[2:])
		if err != nil {
			log.Panic(err)
		}
	case "read":
		err := readCmd.Parse(os.Args[2:])
		if err != nil {
			log.Panic(err)
		}
	case "process":
		err := processCmd.Parse(os.Args[2:])
		if err != nil {
			log.Panic(err)
		}
	default:
		cli.printUsage()
		runtime.Goexit()
	}

	if findCmd.Parsed() {
		if *findFrom == "" || *findPrefix == "" {
			findCmd.Usage()
			runtime.Goexit()
		}
		cli.find(*findFrom, *findPrefix)
	}
	if readCmd.Parsed() {
		if *readFrom == "" || *readPrefix == "" {
			readCmd.Usage()
			runtime.Goexit()
		}
		cli.read(*readFrom, *readPrefix)
	}
	if processCmd.Parsed() {
		if *processFrom == "" || *processPrefix == "" || *processOut == "" {
			processCmd.Usage()
			runtime.Goexit()
		}
		cli.process(*processFrom, *processPrefix, *processOut)
	}

}
