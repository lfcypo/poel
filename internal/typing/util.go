package typing

import "reflect"

func GetType(obj interface{}) string {
	return reflect.TypeOf(obj).String()
}
