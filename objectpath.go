package objectpath

import (
	"reflect"
	"regexp"
	"strconv"
	"strings"
)

var listReg = regexp.MustCompile(`\[(\d+)\]`)

// path is something like follow
// a.b.c
// a.[0].b
func Get[T any](obj any, path string) (T, bool) {
	var _default T
	cur := obj

	for _, key := range strings.Split(path, ".") {
		if matches := listReg.FindStringSubmatch(key); matches != nil {
			idx, err := strconv.Atoi(matches[1])
			if err != nil {
				return _default, false
			}
			rv := reflect.ValueOf(cur)
			ok := rv.Kind() == reflect.Slice
			if !ok {
				return _default, false
			}
			if idx > rv.Len() {
				return _default, false
			}
			cur = rv.Index(idx).Interface()
			continue
		}
		dict, ok := cur.(map[string]any)
		if !ok {
			return _default, false
		}
		v, ok := dict[key]
		if !ok {
			return _default, false
		}
		cur = v
	}
	ret, ok := cur.(T)
	return ret, ok
}
