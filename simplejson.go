// simplejson project simplejson.go
package simplejson

import "encoding/json"

type Json struct {
	data interface{}
}

func (j *Json) Data() interface{} {
	return j.data
}
func (j *Json) Interface(key string) interface{} {
	m, ok := j.data.(map[string]interface{})
	if ok {
		v, _ok := m[key]
		if _ok {
			return v
		}
	}
	return nil
}
func (j *Json) IndexInterface(idx int) interface{} {
	arr := j.data.([]interface{})
	if len(arr) <= idx {
		return nil
	}
	return arr[idx]
}

func (j *Json) Get(key string) *Json {

	return &Json{data: j.Interface(key)}
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
func NewJson(d []byte) (*Json, error) {
	var data interface{}
	err := json.Unmarshal(d, &data)
	if err != nil {
		return nil, err
	}
	return &Json{data: data}, nil
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

	return &Json{data: j.IndexInterface(idx)}
}

func (j *Json) AsDefNum() float64 {
	vf, ok := j.data.(float64)
	if ok {
		return vf
	}
	return 0
}
func (j *Json) AsBool() bool {
	vi, ok := j.data.(bool)
	if ok {
		return vi
	}
	return false
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

func (j *Json) ParseDefNum(d interface{}) float64 {

	if d == nil {
		return 0
	}

	vf, ok := d.(float64)
	if ok {
		return vf
	}
	return 0
}
func (j *Json) GetBool(key string) bool {
	d := j.Interface(key)
	if d == nil {
		return false
	}
	vi, ok := d.(bool)
	if ok {
		return vi
	}
	return false
}
func (j *Json) GetInt(key string) int {
	d := j.Interface(key)
	if d == nil {
		return 0
	}
	vi, ok := d.(int)
	if ok {
		return vi
	}
	return int(j.ParseDefNum(d))
}
func (j *Json) GetUint(key string) uint {
	d := j.Interface(key)
	if d == nil {
		return 0
	}
	vi, ok := d.(uint)
	if ok {
		return vi
	}
	return uint(j.ParseDefNum(d))
}
func (j *Json) GetInt64(key string) int64 {
	d := j.Interface(key)
	if d == nil {
		return 0
	}
	vi, ok := d.(int64)
	if ok {
		return vi
	}
	return int64(j.ParseDefNum(d))
}
func (j *Json) GetUint64(key string) uint64 {
	d := j.Interface(key)
	if d == nil {
		return 0
	}
	vi, ok := d.(uint64)
	if ok {
		return vi
	}
	return uint64(j.ParseDefNum(d))
}
func (j *Json) GetFloat32(key string) float32 {

	d := j.Interface(key)
	if d == nil {
		return 0
	}
	vi, ok := d.(float32)
	if ok {
		return vi
	}
	return float32(j.ParseDefNum(d))
}
func (j *Json) GetFloat64(key string) float64 {

	d := j.Interface(key)

	return j.ParseDefNum(d)
}
func (j *Json) GetString(key string) string {

	d := j.Interface(key)
	if d == nil {
		return ""
	}
	vi, ok := d.(string)
	if ok {
		return vi
	}
	return ""
}
