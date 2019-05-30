// simplejson_test
package simplejson

import (
	"encoding/json"
	"log"
	"testing"
	"time"
)

type t1tp struct {
	A json.Number `json:"a1"`
}

var teststr1 []byte = []byte(`
{
	"a1":922337203685477,
	"a2":"a22",
	"a3":["arr1","arr2",333],
	"a4":{"a41":51.1}
}
`)

var teststr2 []byte = []byte(`["a1",295,"a2","a22","a3",["arr1","arr2",333],"a4",{"a41":51.1}]`)

func Test_simpleJson(t *testing.T) {
	aa := &t1tp{}

	json.Unmarshal(teststr1, aa)
	staaa, _ := json.Marshal(aa)
	sss := string(staaa)
	_ = sss
	j1, e := NewJson(teststr1)
	if e != nil {
		t.Errorf("newjson err")
		log.Println(e)
	}

	j2, e := NewJson(teststr2)
	if e != nil {
		t.Errorf("newjson err")
		log.Println(e)
	}

	if j1.Get("a3").GetIndex(2).AsInt() != 333 {
		t.Errorf("sub arr error")
	}
	if j2.GetIndex(5).GetIndex(1).AsString() != "arr2" {
		t.Errorf("sub arr error")
	}

	barr, e := j1.Marshal()

	if e != nil {
		t.Errorf("Marshal error")
	}
	j3, e := NewJson(barr)
	if e != nil {
		t.Errorf("Unmarshal error")
	}
	if j3.Get("a1").AsInt64() != 922337203685477 ||
		j3.Get("a4").Get("a41").AsFloat32() != 51.1 {
		t.Errorf("Unmarshal afte get error")
		log.Println(j3.GetInt64("a1"))
		log.Println(j3.Get("a4").GetFloat32("a41"))

	}
	barr2, e := j3.Marshal()
	if len(barr) != len(barr2) {
		t.Errorf("Marshal NewJson err")
	}
	tm1 := time.Now()

	for i := 0; i < 10000; i++ {

		jj, _ := NewJson(teststr1)
		teststr1, _ = jj.Marshal()
	}

	log.Println(time.Now().Sub(tm1))
}

func Test_New(t *testing.T) {
	f := New()
	f.Set("a", 1)
	f.Set("b", "aa")
	f2 := New()
	f.Set("c", f2)
	f2.Set("f2", 122.2)

	if f.Get("c").Get("f2").AsFloat64() != 122.2 {
		t.Errorf("test new eroo")
	}

}
