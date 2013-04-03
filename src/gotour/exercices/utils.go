package exercices

import (
    "reflect"
    "runtime"
)

/**
* Return a func's name
*/
func getFunctionName(i interface{}) string {
    return runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name()
}
