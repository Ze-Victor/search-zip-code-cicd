package util

import "fmt"

func ErrParamIsRequired(param, typ string) error {
	return fmt.Errorf("param: %s (type: %s) is requered", param, typ)
}
