package main

import (
	"fmt"
	"os"
	"flag"
	"bufio"
	"strconv"
	"strings"
	"gopkg.in/yaml.v3"
)

type MemorySource interface {
	GetData() map[string]any
}

func FormatAddressMap(input map[string]uint32) map[string]string {
	result := make(map[string]string)
	for key, val := range input {
		result[key] = fmt.Sprintf("0x%08X", val)
	}
	return result
}
func FormatRangeMap(input map[string]AddressRange) map[string]map[string]string {
	result := make(map[string]map[string]string)
	for key, val := range input {
		result[key] = map[string]string{
			"start": fmt.Sprintf("0x%08X", val.Start),
			"end": fmt.Sprintf("0x%08X", val.End),
		}
	}
	return result
}
func writeYAML(filename string, data map[string]any) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := yaml.NewEncoder(file)
	defer encoder.Close()

	return encoder.Encode(data)
}

//need to do this bc go doesnt like outputting raw hex
func PatchAddressPrefix(path string) error {
	inputFile, err := os.Open(path)
	if err != nil {
		return err
	}
	defer inputFile.Close()

	var lines []string
	scanner := bufio.NewScanner(inputFile)
	for scanner.Scan() {
		line := scanner.Text()
		colonIdx := strings.Index(line, ":")
		if colonIdx != -1 {
			value := strings.TrimSpace(line[colonIdx + 1:])
			if num, err := strconv.ParseUint(value, 10, 64); err == nil {
				line = line[:colonIdx + 1] + " 0x" + fmt.Sprintf("%X", num)
			}
		}
		lines = append(lines, line)
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	outputFile, err := os.Create(path)
	if err != nil {
		return err
	}
	defer outputFile.Close()

	for _, line := range lines {
		_, err := outputFile.WriteString(line + "\n")
		if err != nil {
			return err
		}
	}

	return nil
}

func main() {
	input := flag.String("version", "", "TWWHD Version (example: aroma)")
	output := flag.String("output", "output.yaml", "The output file")
	flag.Parse()

	var source MemorySource

	switch *input {
	case "aroma":
		source = AromaSource{}
	default:
		fmt.Println("Version not found: ", *input)
		os.Exit(1)
	}

	data := source.GetData()

	if err := writeYAML(*output, data); err != nil {
		fmt.Println("Failed to write YAML:", err)
		os.Exit(1)
	}

	//if err := PatchAddressPrefix(*output); err != nil {
	//	fmt.Println("Failed to patch prefixes: ", err)
	//	os.Exit(1)
	//}

	fmt.Println("Wrote YAML to", *output)
}
