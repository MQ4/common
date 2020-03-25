package common

import "strconv"

var boolDict = map[string]bool{"t": true, "true": true}

type Map map[string]string

func (m Map) Bool(name string, def bool) bool {
	if val, ok := m[name]; ok {
		return boolDict[val]
	}
	return def
}

func (m Map) String(name, def string) string {
	if val, ok := m[name]; ok {
		return val
	}
	return def
}

func (m Map) Int(name string, def int) int {
	if val, ok := m[name]; ok {
		if v, e := strconv.Atoi(val); e == nil {
			return v
		}
	}
	return def
}

func (m Map) Int64(name string, def int64) int64 {
	if val, ok := m[name]; ok {
		if i, e := strconv.ParseInt(val, 10, 64); e == nil {
			return i
		}
	}
	return def
}

func (m Map) Int32(name string, def int32) int32 {
	if val, ok := m[name]; ok {
		if i, e := strconv.ParseInt(val, 10, 64); e == nil {
			return int32(i)
		}
	}
	return def
}

func (m Map) UInt64(name string, def uint64) uint64 {
	if val, ok := m[name]; ok {
		if i, e := strconv.ParseUint(val, 10, 64); e == nil {
			return i
		}
	}
	return def
}

func (m Map) UInt32(name string, def uint32) uint32 {
	if val, ok := m[name]; ok {
		if i, e := strconv.ParseUint(val, 10, 64); e == nil {
			return uint32(i)
		}
	}
	return def
}
