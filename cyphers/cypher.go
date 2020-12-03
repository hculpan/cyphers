package cypher

import (
	"errors"
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"

	"github.com/antonmedv/expr"

	"github.com/hculpan/cyphers/types"
)

const (
	NUMENERA_CYPHERS int = iota
	SUBTLE_CYPHERS
)

var Rnd = rand.New(rand.NewSource(time.Now().UnixNano()))

func GetCyphers(cypherType int, outputRoll bool) (*types.Cypher, error) {
	var cypherList *[]types.Cypher
	var result *types.Cypher = nil // A pointer just so we don't copy when retrieving from list

	switch cypherType {
	case NUMENERA_CYPHERS:
		cypherList = &types.NumeneraCyphers
	case SUBTLE_CYPHERS:
		cypherList = &types.SubtleCyphers
	}

	totalWeight := calculateTotalWeigth(cypherList)

	originalRoll := Rnd.Int() % totalWeight
	if outputRoll {
		fmt.Printf("Made roll: %d\n", originalRoll)
	}
	roll := originalRoll
	for _, v := range *cypherList {
		if v.Weight > 0 {
			roll -= v.Weight
		} else {
			roll--
		}
		if roll < 0 {
			v.ActualLevel = calculateLevel(v.Level)
			effect, err := processEffect(v.Effect, v.ActualLevel)
			v.Effect = effect
			if err != nil {
				return result, err
			}
			result = &v
			break
		}

	}

	if result == nil {
		return result, errors.New(fmt.Sprintf("Rolled %d, found no result", originalRoll))
	}

	return result, nil
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

func calculateLevel(level string) int {
	if strings.Index(level, "1d10") > -1 {
		level = strings.Replace(level, "1d10", strconv.Itoa(Rnd.Intn(10)+1), -1)
	} else {
		level = strings.Replace(level, "1d6", strconv.Itoa(Rnd.Intn(6)+1), -1)
	}

	env := map[string]interface{}{}

	out, err := expr.Eval(level, env)

	if err != nil {
		fmt.Printf("Error: %s", err.Error())
		panic(err)
	}

	return out.(int)
}
