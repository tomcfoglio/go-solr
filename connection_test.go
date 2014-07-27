package solr

import "testing"
import "fmt"

func TestConnection(t *testing.T) {
	fmt.Println("Running TestConnection")
	/*body,_ := HTTPGet("http://igeonote.com/api/geoip/country/66.249.66.20")

	res,_ := bytes2Json(&body)
	fmt.Println(fmt.Sprintf("%s", *res))
	*/
}

func TestBytes2Json(t *testing.T) {
	data := []byte(`{"t":"s","two":2,"obj":{"c":"b","j":"F"},"a":[1,2,3]}`)
	d, _ := bytes2json(&data)
	if d["t"] != "s" {
		t.Errorf("t should have s as value")
	}

	if d["two"].(float64) != 2 {
		t.Errorf("two should have 2 as value")
	}

	PrintMapInterface(d)
}

func PrintMapInterface(d map[string]interface{}) {
	for k, v := range d {
		switch vv := v.(type) {
		case string:
			fmt.Println(fmt.Sprintf("%s:%s", k, v))
		case int:
			fmt.Println(k, "is int", vv)
		case float64:
			fmt.Println(k, "is float", vv)
		case map[string]interface{}:
			fmt.Println(k, "type is map[string]interface{}")
			for i, u := range vv {
				fmt.Println(i, u)
			}
		case []interface{}:
			fmt.Println(k, "type is []interface{}")
			for i, u := range vv {
				fmt.Println(i, u)
			}
		default:
			fmt.Println(k, "is of a type I don't know how to handle", vv)
		}
	}
}

func TestJson2Bytes(t *testing.T) {

	test_json := map[string]interface{}{
		"t":   "s",
		"two": 2,
		"obj": map[string]interface{}{"c": "b", "j": "F"},
		"a":   []interface{}{1, 2, 3},
	}

	b, err := json2bytes(test_json)
	if err != nil {
		fmt.Println(err)
	}
	d, _ := bytes2json(b)

	if d["t"] != "s" {
		t.Errorf("t should have s as value")
	}

	if d["two"].(float64) != 2 {
		t.Errorf("two should have 2 as value")
	}

	PrintMapInterface(d)
}

func TestHasError(t *testing.T) {
	data := map[string]interface{}{
		"responseHeader": map[string]interface{}{
			"status": 400,
			"QTime":  30,
			"params": map[string]interface{}{
				"indent": "true",
				"q":      "**",
				"wt":     "json"}},
		"error": map[string]interface{}{
			"msg":  "no field name specified in query and no default specified via 'df' param",
			"code": 400}}

	if hasError(data) != true {
		t.Errorf("Should have an error")
	}

	data2 := map[string]interface{}{
		"responseHeader": map[string]interface{}{
			"status": 400,
			"QTime":  30,
			"params": map[string]interface{}{
				"indent": "true",
				"q":      "**",
				"wt":     "json"}},
		"response": map[string]interface{}{
			"numFound": 1,
			"start":    0,
			"docs": []map[string]interface{}{{
				"id":        "change.me",
				"title":     "change.me",
				"_version_": 14}}}}

	if hasError(data2) != false {
		t.Errorf("Should not has an error")
	}
}
