package common

import (
	"strconv"
)

var boolDict = map[string]bool{"t": true, "true": true}

type Map map[string]string

func (m Map) BoolDefault(name string, def bool) bool {
	if val, ok := m[name]; ok {
		return boolDict[val]
	}
	return def
}

func (m Map) StringDefault(name, def string) string {
	if val, ok := m[name]; ok {
		return val
	}
	return def
}

func (m Map) IntDefault(name string, def int) int {
	if val, ok := m[name]; ok {
		if v, e := strconv.Atoi(val); e == nil {
			return v
		}
	}
	return def
}

func (m Map) Int64Default(name string, def int64) int64 {
	if val, ok := m[name]; ok {
		if i, e := strconv.ParseInt(val, 10, 64); e == nil {
			return i
		}
	}
	return def
}

func (m Map) Int32Default(name string, def int32) int32 {
	if val, ok := m[name]; ok {
		if i, e := strconv.ParseInt(val, 10, 64); e == nil {
			return int32(i)
		}
	}
	return def
}

func (m Map) UInt64Default(name string, def uint64) uint64 {
	if val, ok := m[name]; ok {
		if i, e := strconv.ParseUint(val, 10, 64); e == nil {
			return i
		}
	}
	return def
}

func (m Map) UInt32Default(name string, def uint32) uint32 {
	if val, ok := m[name]; ok {
		if i, e := strconv.ParseUint(val, 10, 64); e == nil {
			return uint32(i)
		}
	}
	return def
}
