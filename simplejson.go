// simplejson project simplejson.go
package simplejson

import "encoding/json"

type Json struct {
	data interface{}
}

func (j *Json) Get(key string) *Json {
	m, ok := j.data.(map[string]interface{})
	if ok {
		v, _ok := m[key]
		if _ok {
			vv, isj := v.(*Json)
			if isj {
				return vv
			} else {
				return &Json{data: v}
			}
		}
	}
	return nil
}
func (j *Json) Set(key string, v interface{}) {
	mp := j.data.(map[string]interface{})
	m, ok := v.(*Json)
	if ok {
		mp[key] = m.data
	} else {
		mp[key] = v
	}

}
func New() *Json {
	return &Json{data: make(map[string]interface{})}
}
func NewJson(d []byte) *Json {
	var data interface{}
	json.Unmarshal(d, &data)
	return &Json{data: data}
}

func Unmarshal(d []byte) *Json {
	var data interface{}
	json.Unmarshal(d, &data)
	return &Json{data: data}
}

func NewArray() *Json {
	return &Json{data: make([]interface{}, 0)}
}
func (j *Json) Add(v interface{}) {
	m, ok := v.(*Json)
	if ok {
		j.data = append(j.data.([]interface{}), m.data)
	} else {
		j.data = append(j.data.([]interface{}), v)
	}
}
func (j *Json) Marshal() ([]byte, error) {
	return json.Marshal(j.data)

}

func (j *Json) ArrayLen() int {
	arr, ok := j.data.([]interface{})
	if ok {
		return len(arr)
	}
	return 0
}

func (j *Json) GetIndex(idx int) *Json {
	arr := j.data.([]interface{})
	if len(arr) <= idx {
		return nil
	}
	jo, ok := arr[idx].(*Json)
	if ok {
		return jo
	} else {
		return &Json{data: arr[idx]}
	}
}

func (j *Json) AsDefNum() float64 {
	vf, ok := j.data.(float64)
	if ok {
		return vf
	}
	return 0
}

func (j *Json) AsInt() int {
	vi, ok := j.data.(int)
	if ok {
		return vi
	}
	return int(j.AsDefNum())
}
func (j *Json) AsUint() uint {
	vi, ok := j.data.(uint)
	if ok {
		return vi
	}
	return uint(j.AsDefNum())
}
func (j *Json) AsInt64() int64 {
	vi, ok := j.data.(int64)
	if ok {
		return vi
	}
	return int64(j.AsDefNum())
}
func (j *Json) AsUint64() uint64 {
	vi, ok := j.data.(uint64)
	if ok {
		return vi
	}
	return uint64(j.AsDefNum())
}
func (j *Json) AsFloat32() float32 {

	return float32(j.AsDefNum())
}
func (j *Json) AsFloat64() float64 {

	return j.AsDefNum()
}
func (j *Json) AsString() string {

	vi, ok := j.data.(string)
	if ok {
		return vi
	}
	return ""
}
