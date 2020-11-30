package main

import (
	"errors"
	"flag"
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"

	"github.com/antonmedv/expr"
	"github.com/atotto/clipboard"

	"github.com/hculpan/cyphers/types"
)

//go:generate go run scripts/includecyphers.go

var rnd = rand.New(rand.NewSource(time.Now().UnixNano()))

var unique bool = false
var fullOutput bool = false
var onClipboard bool = true
var useSubtleCyphers bool = false
var useNumeneraCyphers bool = true

func main() {
	fmt.Println("Cyphers v0.2.0")
	selected := map[string]bool{}

	if count, err := processCommandLine(); err == nil {
		var cypherList *[]types.Cypher

		if useSubtleCyphers {
			cypherList = &types.SubtleCyphers
			fmt.Printf("Selecting %d Subtle cyphers\n", count)
		} else {
			cypherList = &types.NumeneraCyphers
			fmt.Printf("Selecting %d Numenera cyphers\n", count)
		}

		totalWeight := calculateTotalWeigth(cypherList)

		clipText := ""
		for i := 0; i < count; i++ {
			roll := rnd.Int() % totalWeight
			originalRoll := roll
			found := false
			for _, v := range *cypherList {
				if v.Weight > 0 {
					roll -= v.Weight
				} else {
					roll--
				}
				if roll < 0 {
					selected[v.Name] = true
					lvl := calculateLevel(v.Level)
					fmt.Println("----------------------------------------------")
					effect, err := processEffect(v.Effect, lvl)
					if err != nil {
						printUsage(err.Error())
						return
					}

					if fullOutput {
						fmt.Printf(buildCypherOutputForConsole(v, lvl, effect))
					} else {
						fmt.Printf("%s [rolled %d] : Level %d (%s)\n", v.Name, originalRoll, lvl, v.Level)
						if onClipboard {
							clipText += buildCypherOutputForClipboard(v, lvl, effect)
						}
					}
					found = true
					break
				}
			}

			if !found {
				fmt.Printf("Uh-oh, rolled %d, didn't find a result!\n", roll)
				break
			}
		}

		if clipText != "" {
			clipboard.WriteAll(clipText)
		}
	} else {
		printUsage(err.Error())
		return
	}
}

func buildCypherOutputForClipboard(c types.Cypher, lvl int, effect string) string {
	return buildCypherOutput(c, lvl, strings.ReplaceAll(effect, "\n", " "), true)
}

func buildCypherOutputForConsole(c types.Cypher, lvl int, effect string) string {
	return buildCypherOutput(c, lvl, effect, false)
}

func buildCypherOutput(c types.Cypher, lvl int, effect string, roll20Formatting bool) string {
	var result string
	if roll20Formatting {
		result = fmt.Sprintf("**%s : Level %d**\n", c.Name, lvl)
	} else {
		result = fmt.Sprintf("%s : Level %d\n", c.Name, lvl)
	}

	if len(c.Type) > 0 {
		result += fmt.Sprintf("Type: %s\n", c.Type[rand.Int()%len(c.Type)])
	}

	if c.Usable != "" {
		result += fmt.Sprintf("Usable: %s\n", c.Usable)
	}

	result += effect + "\n"

	return result
}

func calculateTotalWeigth(list *[]types.Cypher) int {
	result := 0

	for _, v := range *list {
		if v.Weight > 0 {
			result += v.Weight
		} else {
			result++
		}
	}

	return result
}

func processEffect(effect string, level int) (string, error) {
	env := map[string]interface{}{
		"LEVEL": level,
	}

	for {
		start := strings.Index(effect, "$(")
		if start == -1 {
			break
		}
		end := strings.Index(effect[start:], ")") + start
		if end == -1 {
			break
		}

		out, err := expr.Eval(effect[start+2:end], env)

		if err != nil {
			return effect, err
		}

		effect = strings.Replace(effect, effect[start:end+1], strconv.Itoa(out.(int)), -1)
	}

	return effect, nil
}

func processCommandLine() (int, error) {
	flag.Usage = usage

	flag.BoolVar(&unique, "u", false, "Do not select the same cypher twice")
	flag.BoolVar(&unique, "unique", false, "Do not select the same cypher twice")
	flag.BoolVar(&fullOutput, "f", false, "Full (non-Roll20 formatted) output")
	flag.BoolVar(&fullOutput, "full", false, "Full (non-Roll20 formatted) output")
	flag.BoolVar(&onClipboard, "no-clip", true, "Do not put on clipbaord")
	flag.BoolVar(&useSubtleCyphers, "s", false, "Use the subtle cyphers from the Cypher System core rules")
	flag.BoolVar(&useNumeneraCyphers, "n", true, "Use cypher list from Numenera (DEFAULT)")

	flag.Parse()

	nArgs := len(flag.Args())
	count := 1
	switch {
	case nArgs > 1:
		return count, errors.New("Incorrect number of parameters")
	case nArgs == 1:
		c, err := strconv.Atoi(flag.Arg(0))
		if err != nil {
			return count, errors.New(fmt.Sprintf("Invalid parameter: '%s'", flag.Arg(0)))
		}
		count = c
	}

	return count, nil
}

func calculateLevel(level string) int {
	level = strings.Replace(level, "1d6", strconv.Itoa(rnd.Intn(6)+1), -1)

	env := map[string]interface{}{}

	out, err := expr.Eval(level, env)

	if err != nil {
		panic(err)
	}

	return out.(int)
}

func usage() {
	fmt.Println("Usage:")
	fmt.Println()
	fmt.Println("        cyphers [flags] [n]")
	fmt.Println()
	fmt.Println("Where \"n\" is the number of cyphers to generate.")
	fmt.Println()
	fmt.Println("      -no-clip             Do not put on clipboard")
	fmt.Println("  -f, -full                Full (non-Roll20 formatted) output")
	fmt.Println("  -u, -unique              Do not select the same cypher twice")
	fmt.Println()
	fmt.Println("Cypher lists")
	fmt.Println("  You can select which list of cyphers to use.")
	fmt.Println("  -n                       Use cypher list from Numenera (DEFAULT)")
	fmt.Println("  -s                       Use the subtle cyphers from the Cypher System core rules")
	fmt.Println()
}

func printUsage(errorMsg string) {
	fmt.Println(errorMsg)
	usage()
}
