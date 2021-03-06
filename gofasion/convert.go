package gofasion

import (
	"strconv"
	"strings"
)

//Init current json node as single node
func (self *Fasion) initCur() {
	if self.current == nil {
		curNode, err := self.parseJson()
		self.current = curNode
		self.errInfo = err
	}
}

//Init current json node as array
func (self *Fasion) initArray() {
	if self.current == nil {
		curNode, err := self.parseArray()
		self.current = curNode
		self.errInfo = err
	}
}

//Parse current node value to string
func (self *Fasion) ValueStr() string {
	self.initCur()
	if val, ok := self.current.(map[string]interface{}); ok {
		if x, ok := val[self.currentKey].(string); ok {
			return x
		} else {
			return self.rawJson
		}
	}
	return ""
}

//Parse current node value to int64
func (self *Fasion) ValueInt64() int64 {
	self.initCur()
	if val, ok := self.current.(map[string]interface{}); ok {
		if x, ok := val[self.currentKey].(string); ok {
			if n, ok := strconv.ParseInt(x, 10, 64); ok == nil {
				return n
			}
		}
	}
	return 0
}

//Parse current node value to int32
func (self *Fasion) ValueInt32() int32 {
	self.initCur()
	if val, ok := self.current.(map[string]interface{}); ok {
		if x, ok := val[self.currentKey].(string); ok {
			if n, ok := strconv.ParseInt(x, 10, 64); ok == nil {
				return int32(n)
			}
		}
	}
	return 0
}

//Parse current node value to int16
func (self *Fasion) ValueInt16() int16 {
	self.initCur()
	if val, ok := self.current.(map[string]interface{}); ok {
		if x, ok := val[self.currentKey].(string); ok {
			if n, ok := strconv.ParseInt(x, 10, 64); ok == nil {
				return int16(n)
			}
		}
	}
	return 0
}

//Parse current node value to int
func (self *Fasion) ValueInt() int {
	self.initCur()
	if val, ok := self.current.(map[string]interface{}); ok {
		if x, ok := val[self.currentKey].(string); ok {
			if n, ok := strconv.ParseInt(x, 10, 64); ok == nil {
				return int(n)
			}
		}
	}
	return 0
}

//Parse current node value to float32
func (self *Fasion) ValueFloat32() float32 {
	self.initCur()
	if val, ok := self.current.(map[string]interface{}); ok {
		if x, ok := val[self.currentKey].(string); ok {
			if n, ok := strconv.ParseFloat(x, 32); ok == nil {
				return float32(n)
			}
		}
	}
	return 0
}

//Retained specifc decimals
//parse 1.1115 and spec is 3 and return 1.112
func (self *Fasion) ValueFloat32N(spec int) float32 {
	ret := self.ValueFloat32()
	if spec <= 0 {
		return ret
	}
	return float32(round(float64(ret), spec))
}

//Parse current node value to float64
func (self *Fasion) ValueFloat64() float64 {
	self.initCur()
	if val, ok := self.current.(map[string]interface{}); ok {
		if x, ok := val[self.currentKey].(string); ok {
			if n, ok := strconv.ParseFloat(x, 64); ok == nil {
				return n
			}
		}
	}
	return 0
}

//Retained specifc decimals
//parse 1.1115 and spec is 3 and return 1.112
func (self *Fasion) ValueFloat64N(spec int) float64 {
	ret := self.ValueFloat64()
	if spec <= 0 {
		return ret
	}
	return round(ret, spec)
}

//Marshal current node to json string
func (self *Fasion) Json() string {
	return strings.Trim(self.rawJson, " ")
}

//Parse current node value to []*Fasion nodes
func (self *Fasion) Array() []*Fasion {
	self.initArray()
	result := make([]*Fasion, 0)
	if val, ok := self.current.([]interface{}); ok {
		for _, v := range val {
			rawJson, err := self.toJson(v)
			if err == nil {
				result = append(result, NewFasion(rawJson))
			}
		}
	}
	return result
}

//Parse current node Value to []*Fasion and iterate it via job function
func (self *Fasion) ArrayForEach(job func(int, *Fasion)) {
	elements := self.Array()
	for i, v := range elements {
		job(i, v)
	}
}

//Parse current node value to bool
func (self *Fasion) ValueBool() bool {
	self.initCur()
	valueStr := self.ValueStr()
	if valueStr == "" {
		return false
	}
	if strings.ToLower(valueStr) == "false" {
		return false
	}
	if strings.ToLower(valueStr) == "true" {
		return true
	}
	return false
}

//Parse current node value to specific interface
func (self *Fasion) Value(dest interface{}) error {
	return _unmarshalFunc([]byte(self.rawJson), &dest)
}

//List All keys of this node
func (self *Fasion) Keys() []string {
	var keys []string
	curMap, err := self.parseJson()
	if err != nil {
		return keys
	}
	for k, _ := range curMap {
		if k == self.currentKey {
			continue
		}
		keys = append(keys, k)
	}
	return keys
}

//Judge if specific key exists or not
func (self *Fasion) HasKey(key string) bool {
	if key == self.currentKey {
		return false
	}
	curMap, err := self.parseJson()
	if err != nil {
		return false
	}
	_, hasKey := curMap[key]
	return hasKey
}

func (self *Fasion) ValueDefaultStr(defaultValue string) string {
	ret := self.ValueStr()
	if ret == "" {
		return defaultValue
	}
	return ret
}

func (self *Fasion) ValueDefaultInt(defaultValue int) int {
	if self.current == nil && self.currentKey == "" && self.rawJson == "" {
		return defaultValue
	}
	return self.ValueInt()
}

func (self *Fasion) ValueDefaultInt16(defaultValue int16) int16 {
	if self.current == nil && self.currentKey == "" && self.rawJson == "" {
		return defaultValue
	}
	return self.ValueInt16()
}

func (self *Fasion) ValueDefaultInt32(defaultValue int32) int32 {
	if self.current == nil && self.currentKey == "" && self.rawJson == "" {
		return defaultValue
	}
	return self.ValueInt32()
}

func (self *Fasion) ValueDefaultInt64(defaultValue int64) int64 {
	if self.current == nil && self.currentKey == "" && self.rawJson == "" {
		return defaultValue
	}
	return self.ValueInt64()
}

func (self *Fasion) ValueDefaultFloat32(defaultValue float32) float32 {
	if self.current == nil && self.currentKey == "" && self.rawJson == "" {
		return defaultValue
	}
	return self.ValueFloat32()
}

func (self *Fasion) ValueDefaultFloat64(defaultValue float64) float64 {
	if self.current == nil && self.currentKey == "" && self.rawJson == "" {
		return defaultValue
	}
	return self.ValueFloat64()
}

func (self *Fasion) ValueDefaultBool(defaultValue bool) bool {
	if self.current == nil && self.currentKey == "" && self.rawJson == "" {
		return defaultValue
	}
	return self.ValueBool()
}
