package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"unicode"

	"github.com/hculpan/cyphers/types"
)

// Reads all .txt files in the current folder
// and encodes them as strings literals in textfiles.go
func main() {
	fs, _ := ioutil.ReadDir("./cypher-data")
	for _, f := range fs {
		if strings.HasSuffix(f.Name(), ".txt") {
			if file, err := os.Open(filepath.Join(".", "cypher-data", f.Name())); err == nil {
				defer file.Close()

				out, _ := os.Create(filepath.Join(".", "types", removeExtension(f.Name())+"Cyphers.go"))
				defer out.Close()

				out.WriteString("package types\n\n")

				out.WriteString("")
				cypherLines := []string{}

				out.WriteString("// " + file.Name() + "\n\n")

				out.WriteString("var NumeneraCyphers []Cypher = []Cypher {\n")

				s := bufio.NewScanner(file)
				for s.Scan() {
					line := s.Text()
					if isAllUpper(line) {
						if len(cypherLines) > 0 {
							outputCypher(cypherLines, out)
						}
						cypherLines = []string{line}
					}
					cypherLines = append(cypherLines, line)
				}

				// make sure to write out final cypher
				if len(cypherLines) > 0 {
					outputCypher(cypherLines, out)
				}

				out.Write([]byte("}\n"))
			} else {
				fmt.Printf("Unable to open file '%s': %s\n", f.Name(), err.Error())
				break
			}
		}
	}
}

func newCypher(lines []string) types.Cypher {
	return types.Cypher{Name: convertToMixedCase(lines[0])}
}

func outputCypher(lines []string, out *os.File) {
	out.WriteString(fmt.Sprint("\tCypher{ \n"))
	out.WriteString(fmt.Sprintf("\t\tName: \"%s\",\n", convertToMixedCase(lines[0])))

	// Have to deal with situations where we multiple types
	// A little klunky, but iterate through all lines and collect all Types
	typeLines := []string{}
	for i := 0; i < len(lines); i++ {
		if strings.HasPrefix(lines[i], "Type:") {
			typeLines = append(typeLines, lines[i])
		}
	}

	typesReadyToOutput := true

	for i := 0; i < len(lines); i++ {
		s := lines[i]
		switch {
		case strings.HasPrefix(s, "Level:"):
			out.WriteString(fmt.Sprintf("\t\tLevel: \"%s\",\n", strings.Trim(s[6:], " ")))
		case strings.HasPrefix(s, "Type:") && typesReadyToOutput:
			out.WriteString(fmt.Sprint("\t\tType: []string{\n"))
			for _, l := range typeLines {
				out.WriteString(fmt.Sprintf("\t\t\t\"%s\",\n", strings.Trim(l[5:], " ")))
			}
			out.WriteString(fmt.Sprint("\t\t},\n"))
			typesReadyToOutput = false
		case strings.HasPrefix(s, "Usable:"):
			out.WriteString(fmt.Sprintf("\t\tUsable: \"%s\",\n", strings.Trim(s[7:], " ")))
		case strings.HasPrefix(s, "Selection weighting:"):
			out.WriteString(fmt.Sprintf("\t\tWeight: %s,\n", strings.Trim(s[20:], " ")))
		case strings.HasPrefix(s, "Effect:"): // This should be the last item, so iterate and get them all
			if i == len(lines)-1 {
				out.WriteString(fmt.Sprintf("\t\tEffect: `%s`,\n", strings.Trim(s[7:], " ")))
			} else {
				out.WriteString(fmt.Sprintf("\t\tEffect: `%s\n", strings.Trim(lines[i][7:], " ")))
				for i < len(lines)-2 {
					i++
					out.WriteString(fmt.Sprintf("%s\n", lines[i]))
				}
				i++
				out.WriteString(fmt.Sprintf("%s`,\n", lines[i]))
			}
		}
	}
	out.WriteString("\t},\n")
}

func removeExtension(s string) string {
	return strings.Split(s, ".")[0]
}

func convertToMixedCase(s string) string {
	result := ""

	lastWasSpace := true
	index := 0
	for _, b := range strings.ToLower(s) {
		if lastWasSpace {
			result += strings.ToUpper(string(b))
			lastWasSpace = false
		} else if b == ' ' {
			result += " "
			lastWasSpace = true
		} else if b == '-' {
			result += "-"
			lastWasSpace = true
		} else if b == '(' {
			result += "("
			lastWasSpace = true
		} else {
			result += string(b)
		}
		index++
	}
	return strings.Trim(result, " ")
}

func isAllUpper(s string) bool {
	for _, r := range s {
		if !unicode.IsUpper(r) && unicode.IsLetter(r) {
			return false
		}
	}

	return true
}
