package parser

import (
	"fmt"
	"strings"

	"github.com/lfcypo/poel/internal/core/engine"
)

func SyntaxError(message string) {
	panic(fmt.Sprintf("Syntax Error: %s", message))
}

func Participle(expression string) {
	expressionSplits := strings.Split(expression, "")

	isStringLiteral := false
	stringLiteralIndex := 0

	var aParam string
	var bParam string

	index := 0
	for {
		if index >= len(expressionSplits) {
			break
		}
		word := expressionSplits[index]
		//fmt.Println(index, expressionSplits[index])

		if word == " " && !isStringLiteral {
			index++
			continue
		}

		// Parse String Literals
		if strings.HasPrefix(word, "\"") && !isStringLiteral {
			isStringLiteral = true
			stringLiteralIndex++
			index++
			continue
		} else if strings.HasSuffix(word, "\"") && isStringLiteral {
			isStringLiteral = false
			index++
			continue
		} else if isStringLiteral {
			index++
			continue
		}

		// Parse Operators
		if engine.IsOperator(word) {
			j := 0
			for j = 0; j < index; j-- {
				if engine.IsOperator(expressionSplits[j]) {
					continue
				}
				aParam = expressionSplits[j]
				break
			}
			for j = index + 2; j < len(expressionSplits); j++ {
				if engine.IsOperator(expressionSplits[j]) {
					continue
				}
				fmt.Println("bParam")
				bParam = expressionSplits[j]
				break
			}
			index++
			fmt.Println(aParam, word, bParam)

			continue
		}

		index++
	}
	if isStringLiteral {
		SyntaxError("Unterminated String \" Literal")
	}

	//fmt.Println(stringLiterals)

}
