package mxj

import (
	"strings"
)

// Sets the value for the path
func (mv Map) SetValueForPath(value interface{}, path string) error {
	pathAry := strings.Split(path, pathSeparator)
	parentPathAry := pathAry[0 : len(pathAry)-1]
	parentPath := strings.Join(parentPathAry, pathSeparator)

	val, err := mv.ValueForPath(parentPath)
	if err != nil {
		return err
	}
	if val == nil {
		return nil // we just ignore the request if there's no val
	}

	key := pathAry[len(pathAry)-1]
	cVal := val.(map[string]interface{})
	cVal[key] = value

	return nil
}
