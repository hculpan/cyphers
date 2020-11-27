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

var rnd = rand.New(rand.NewSource(time.Now().UnixNano()))

var unique bool = false
var fullOutput bool = false
var onClipboard bool = true

func main() {
	fmt.Println("Cyphers v0.1.0")
	selected := map[string]bool{}

	count, err := processCommandLine()
	if count > 1 && unique {
		fmt.Printf("Selecting %d unique results\n", count)
	}

	if err != nil {
		printUsage(err.Error())
		return
	}

	clipText := ""
	for i := 0; i < count; i++ {
		roll := 29 //(rnd.Int() % 100) + 1
		fmt.Printf("The roll is %d\n", roll)
		found := false
		for k, v := range types.SubtleCyphers {
			if isBetween(k, roll) {
				if unique && selected[v.Name] {
					fmt.Println("Got here but shouldn't have", unique)
				} else {
					selected[v.Name] = true
					lvl := calculateLevel(v.Level)
					fmt.Println("----------------------------------------------")
					s, err := processEffect(v.Effect, lvl)
					if err != nil {
						printUsage(err.Error())
						return
					}

					if fullOutput {
						fmt.Printf("%s [rolled %d] : Level %d (%s)%s\n", v.Name, roll, lvl, v.Level, s)
					} else {
						fmt.Printf("%s [rolled %d] : Level %d (%s)\n", v.Name, roll, lvl, v.Level)
						if onClipboard {
							s = strings.ReplaceAll(s, "\n", "")
							//							s = strings.ReplaceAll(s, "\r", "")
							clipText += fmt.Sprintf("\n**%s : Level %d**\n%s", v.Name, lvl, s)
						}
					}
					found = true
					break
				}
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

func isBetween(rollRange string, roll int) bool {
	nums := strings.Split(rollRange, "-")
	n1, _ := strconv.Atoi(nums[0])
	n2, _ := strconv.Atoi(nums[1])
	return roll >= n1 && roll <= n2
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
	fmt.Println("      -on-clip             Do not put on clipboard")
	fmt.Println("  -f, -full                Full (non-Roll20 formatted) output")
	fmt.Println("  -u, -unique              Do not select the same cypher twice")
	fmt.Println()
}

func printUsage(errorMsg string) {
	fmt.Println(errorMsg)
	usage()
}
