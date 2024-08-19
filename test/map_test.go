package test

import (
	_ "embed"
	"encoding/json"
	"testing"
	"time"
)

//go:embed dictionaries.json
var dictionaries []byte

type dictionary struct {
	Type  string                      // 字典分类
	Items map[interface{}]interface{} // 字典项
}

func TestMap(t *testing.T) {
	var maps []dictionary
	err := json.Unmarshal(dictionaries, &maps)
	if err != nil {
		panic(err)
	}
	//s := []dictionary{
	//	{Type: "gender", Items: map[interface{}]interface{}{
	//		"1":  "男",
	//		2:    "女",
	//		true: "女",
	//	}},
	//	{Type: "gender", Items: map[interface{}]interface{}{
	//		"1":  "男",
	//		2:    "女",
	//		true: "女",
	//	}},
	//}
	values := withTransferValues(maps)
	for k, v := range values {
		t.Log("key:", k, "value:", v)
	}

	t.Log(values["gender"])
}

func withTransferValues(s []dictionary) map[interface{}]interface{} {
	m := make(map[interface{}]interface{}, len(s))
	for _, d := range s {
		if _, ok := m[d.Type]; ok {
			panic("duplicate type")
		}
		m[d.Type] = d.Items
	}
	return m
}

func TestTime(t *testing.T) {
	t.Log(time.Now().Format(time.RFC3339))

	parse, err := time.Parse(time.RFC3339, "2014-11-12T11:45:26+08:00")
	if err != nil {
		t.Error(err)
		return
	}

	t.Log(parse.Format(time.RFC3339))
}
