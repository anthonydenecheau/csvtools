package extract

import (
	"bufio"
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/jszwec/csvutil"
)

type ExtractCmd struct{}

type LineCsau struct {
	DtEpreuve    string
	Organisateur string
	CdRace       string
	CdLO         string
	NumLO        string
	TatooChip    string
	CdResultat   string
}

type LineTc struct {
	DtEpreuve    string
	Organisateur string
	CdFci        string
	CdLO         string
	NumLO        string
	TatooChip    string
	CdResultat   string
	Niveau       string
}

func visit(prefix *string, files *[]string) filepath.WalkFunc {
	return func(path string, info os.FileInfo, err error) error {
		if err != nil {
			log.Fatal(err)
		}
		if filepath.Ext(path) == ".csv" && strings.HasPrefix(info.Name(), *prefix) {
			*files = append(*files, path)

		}
		return nil
	}
}

func (extract *ExtractCmd) FindCSV(prefix string, directory string) (files []string, _err error) {

	err := filepath.Walk(directory, visit(&prefix, &files))
	if err != nil {
		return files, err
	}
	if len(files) == 0 {
		return files, errors.New("Aucun fichier trouvé")
	}
	return files, nil
}

func setDecoder(csvfile *os.File, lineType interface{}) (dec *csvutil.Decoder, _err error) {

	// Skip first row (line)
	row1, err := bufio.NewReader(csvfile).ReadSlice('\n')
	if err != nil {
		return nil, errors.New("Couldn't read the csv file")
	}
	_, err = csvfile.Seek(int64(len(row1)), io.SeekStart)
	if err != nil {
		return nil, errors.New("Couldn't skip first line the csv file")
	}

	// Read remaining rows
	r := csv.NewReader(csvfile)
	r.Comma = ';'

	csvHeader, err := csvutil.Header(lineType, "csv")
	if err != nil {
		return nil, errors.New("Couldn't fix header")
	}

	dec, err1 := csvutil.NewDecoder(r, csvHeader...)
	if err1 != nil {
		return nil, errors.New("Couldn't configure decoder")
	}

	return dec, nil

}

func parseTcCSV(files []string) (lines []LineTc, _err error) {

	for _, file := range files {

		log.Printf("ParseCSV >  %s", file)
		csvfile, err := os.Open(file)
		if err != nil {
			return lines, errors.New("Couldn't open the csv file")
		}
		defer csvfile.Close()

		dec, err := setDecoder(csvfile, LineTc{})
		if err != nil {
			return lines, errors.New("Couldn't configure decoder")
		}

		for {
			var l LineTc
			if err := dec.Decode(&l); err == io.EOF {
				break
			} else if err != nil {
				return lines, errors.New("Couldn't decode csv line")
			}
			lines = append(lines, l)
		}
	}

	return lines, nil
}

func parseCsauCSV(files []string) (lines []LineCsau, _err error) {

	for _, file := range files {

		log.Printf("ParseCSV >  %s", file)
		csvfile, err := os.Open(file)
		if err != nil {
			return lines, errors.New("Couldn't open the csv file")
		}
		defer csvfile.Close()

		dec, err := setDecoder(csvfile, LineCsau{})
		if err != nil {
			return lines, errors.New("Couldn't configure decoder")
		}

		for {
			var l LineCsau
			if err := dec.Decode(&l); err == io.EOF {
				break
			} else if err != nil {
				return lines, fmt.Errorf("Couldn't decode csv line %s", err)
			}
			lines = append(lines, l)
		}

	}
	return lines, nil
}

func (extrac *ExtractCmd) ParseCSV(prefix string, files []string) (lines []interface{}, _err error) {
	switch prefix {
	case "csau":
		var lines []LineCsau
		lines, _err := parseCsauCSV(files)
		return []interface{}{lines, nil}, _err
	case "tc":
		var lines []LineTc
		lines, _err := parseTcCSV(files)
		return []interface{}{nil, lines}, _err
	default:
		return []interface{}{nil, nil}, fmt.Errorf("Couldn't recognized prefix %s", prefix)
	}

}
