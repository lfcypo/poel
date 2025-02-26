package engine

import (
	"math"
	"strings"

	"github.com/lfcypo/poel/internal/typing"
)

func IsOperator(s string) bool {
	for _, op := range OperatingSymbols {
		if s == op {
			return true
		}
	}
	return false
}

var OperatingSymbols []string = []string{"+", "-", "**", "/", "^", "%", "//", "==", "!=", ">", "<", ">=", "<="}

func AddInteger(left typing.Integer, right typing.Integer) typing.Integer {
	return left + right
}

func AddFloat(left typing.Float, right typing.Float) typing.Float {
	return left + right
}

func AddString(left typing.String, right typing.String) typing.String {
	return left + right
}

func AddBoolean(left typing.Boolean, right typing.Boolean) typing.Boolean {
	panic("Cannot add rightoolean values")
}

func SubtractionInteger(left typing.Integer, right typing.Integer) typing.Integer {
	return left - right
}

func SubtractionFloat(left typing.Float, right typing.Float) typing.Float {
	return left - right
}

func SubtractionString(left typing.String, right typing.String) typing.String {
	panic("Cannot subtract string values")
}

func SubtractionBoolean(left typing.Boolean, right typing.Boolean) typing.Boolean {
	panic("Cannot subtract rightoolean values")
}

func MultiplicationInteger(left typing.Integer, right typing.Integer) typing.Integer {
	return left * right
}

func MultiplicationFloat(left typing.Float, right typing.Float) typing.Float {
	return left * right
}

func MultiplicationString(left typing.String, right typing.String) typing.String {
	panic("Cannot multiply string values")
}

func RepeatString(left typing.String, right typing.Integer) typing.String {
	return typing.String(strings.Repeat(string(left), int(right)))
}

func MultiplicationBoolean(left typing.Boolean, right typing.Boolean) typing.Boolean {
	panic("Cannot multiply rightoolean values")
}

func DivisionInteger(left typing.Integer, right typing.Integer) typing.Integer {
	return left / right
}

func DivisionFloat(left typing.Float, right typing.Float) typing.Float {
	return left / right
}

func DivisionString(left typing.String, right typing.String) typing.String {
	panic("Cannot divide string values")
}

func DivisionBoolean(left typing.Boolean, right typing.Boolean) typing.Boolean {
	panic("Cannot divide rightoolean values")
}

func PowerInteger(left typing.Integer, right typing.Integer) typing.Integer {
	return typing.Integer(math.Pow(float64(left), float64(right)))
}

func PowerFloat(left typing.Float, right typing.Float) typing.Float {
	return typing.Float(math.Pow(float64(left), float64(right)))
}

func PowerString(left typing.String, right typing.String) typing.String {
	panic("Cannot raise string values to left power")
}

func PowerBoolean(left typing.Boolean, right typing.Boolean) typing.Boolean {
	panic("Cannot raise rightoolean values to left power")
}

func ModuloInteger(left typing.Integer, right typing.Integer) typing.Integer {
	return left % right
}

func ModuloFloat(left typing.Float, right typing.Float) typing.Float {
	panic("Cannot take modulo of float values")
}

func ModuloString(left typing.String, right typing.String) typing.String {
	panic("Cannot take modulo of string values")
}

func ModuloBoolean(left typing.Boolean, right typing.Boolean) typing.Boolean {
	panic("Cannot take modulo of rightoolean values")
}

func FloorDivisionInteger(left typing.Integer, right typing.Integer) typing.Integer {
	return left // right
}

func FloorDivisionFloat(left typing.Float, right typing.Float) typing.Float {
	return left // right
}

func FloorDivisionString(left typing.String, right typing.String) typing.String {
	panic("Cannot take floor division of string values")
}

func FloorDivisionBoolean(left typing.Boolean, right typing.Boolean) typing.Boolean {
	panic("Cannot take floor division of rightoolean values")
}

func EqualInteger(left typing.Integer, right typing.Integer) typing.Boolean {
	return left == right
}

func EqualFloat(left typing.Float, right typing.Float) typing.Boolean {
	return left == right
}

func EqualString(left typing.String, right typing.String) typing.Boolean {
	return left == right
}

func EqualBoolean(left typing.Boolean, right typing.Boolean) typing.Boolean {
	return left == right
}

func NotEqualInteger(left typing.Integer, right typing.Integer) typing.Boolean {
	return left != right
}

func NotEqualFloat(left typing.Float, right typing.Float) typing.Boolean {
	return left != right
}

func NotEqualString(left typing.String, right typing.String) typing.Boolean {
	return left != right
}

func NotEqualBoolean(left typing.Boolean, right typing.Boolean) typing.Boolean {
	return left != right
}

func GreaterThanInteger(left typing.Integer, right typing.Integer) typing.Boolean {
	return left > right
}

func GreaterThanFloat(left typing.Float, right typing.Float) typing.Boolean {
	return left > right
}

func GreaterThanString(left typing.String, right typing.String) typing.Boolean {
	panic("Cannot compare string values")
}

func GreaterThanBoolean(left typing.Boolean, right typing.Boolean) typing.Boolean {
	panic("Cannot compare rightoolean values")
}

func GreaterThanOrEqualInteger(left typing.Integer, right typing.Integer) typing.Boolean {
	return left >= right
}

func GreaterThanOrEqualFloat(left typing.Float, right typing.Float) typing.Boolean {
	return left >= right
}

func GreaterThanOrEqualString(left typing.String, right typing.String) typing.Boolean {
	panic("Cannot compare string values")
}

func GreaterThanOrEqualBoolean(left typing.Boolean, right typing.Boolean) typing.Boolean {
	panic("Cannot compare rightoolean values")
}

func LessThanInteger(left typing.Integer, right typing.Integer) typing.Boolean {
	return left < right
}

func LessThanFloat(left typing.Float, right typing.Float) typing.Boolean {
	return left < right
}

func LessThanString(left typing.String, right typing.String) typing.Boolean {
	panic("Cannot compare string values")
}

func LessThanBoolean(left typing.Boolean, right typing.Boolean) typing.Boolean {
	panic("Cannot compare rightoolean values")
}

func LessThanOrEqualInteger(left typing.Integer, right typing.Integer) typing.Boolean {
	return left <= right
}

func LessThanOrEqualFloat(left typing.Float, right typing.Float) typing.Boolean {
	return left <= right
}

func LessThanOrEqualString(left typing.String, right typing.String) typing.Boolean {
	panic("Cannot compare string values")
}

func LessThanOrEqualBoolean(left typing.Boolean, right typing.Boolean) typing.Boolean {
	panic("Cannot compare rightoolean values")
}
