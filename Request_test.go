package gorest
import (
	"testing"
	//"net/url"
	"net/http"
	"strings"
)

func TestGetMethod(t *testing.T) {
	req, _ := http.NewRequest("POST", "http://www.google.com/search?str=foo",
		strings.NewReader("int=2&float=12.4&bool=1"))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded; param=value")


	gorest_request := &Request{req,nil}

	_,err := gorest_request.GetStringValue("sss");

	if err==nil {
		t.Error("Wrong");
	}

	if q := gorest_request.GetStringValueOrDefault("str",""); q != "faa" {
		t.Errorf("GetStringValue Wrong!")
	}
	if ival := gorest_request.GetIntValueOrDefault("int",0); ival != 2 {
		t.Errorf("GetIntValue Wrong!")
	}
	if floatVal := gorest_request.GetFloatValueOrDefault("float",-1); floatVal != 12.4 {
		t.Errorf("GetFloatValue Wrong!")
	}
	if uVal := gorest_request.GetUintValueOrDefault("int",0); uVal != 2 {
		t.Errorf("GetUintValue Wrong!")
	}
	if boolVal := gorest_request.GetBoolValueOrDefault("bool",false); boolVal != true {
		t.Errorf("GetBoolValue Wrong!")
	}
}

func TestPOSTFormValueMethod(t *testing.T) {
	req, _ := http.NewRequest("POST", "http://www.google.com/search?str=foo",
					strings.NewReader("int=2&float=12.4&bool=1"))
			req.Header.Set("Content-Type", "application/x-www-form-urlencoded; param=value")


	gorest_request := &Request{req,nil}

	if q := gorest_request.GetStringValueOrDefault("str",""); q != "foo" {
		t.Errorf("GetStringValue Wrong!")
	}
	if ival := gorest_request.GetIntValueOrDefault("int",0); ival != 2 {
		t.Errorf("GetIntValue Wrong!")
	}
	if floatVal := gorest_request.GetFloatValueOrDefault("float",-1); floatVal != 12.4 {
		t.Errorf("GetFloatValue Wrong!")
	}
	if uVal := gorest_request.GetUintValueOrDefault("int",0); uVal != 2 {
		t.Errorf("GetUintValue Wrong!")
	}
	if boolVal := gorest_request.GetBoolValueOrDefault("bool",false); boolVal != true {
		t.Errorf("GetBoolValue Wrong!")
	}
}

func TestPOSTJsonMethod(t *testing.T) {
	req, _ := http.NewRequest("POST", "http://www.google.com/search?str=foo",
		strings.NewReader("{str:'ssss',int:2,float:12.4,bool:1}"))
	req.Header.Set("Content-Type", "application/json")


	gorest_request := &Request{req,nil}

	body :=gorest_request.GetJsonBody()

	expected := "{str:'ssss',int:2,float:12.4,bool:1}"
	if string(body)!=expected {
		t.Error("GetJsonBody wrong!");
	}
}

func TestPOSTFileMethod(t *testing.T) {

}

func TestPUTMethod(t *testing.T) {
	req, _ := http.NewRequest("PUT", "http://www.google.com/search?str=foo",
		strings.NewReader("int=2&float=12.4&bool=1"))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded; param=value")


	gorest_request := &Request{req,nil}

	if q := gorest_request.GetStringValueOrDefault("str",""); q != "foo" {
		t.Errorf("GetStringValue Wrong!")
	}
	if ival := gorest_request.GetIntValueOrDefault("int",0); ival != 2 {
		t.Errorf("GetIntValue Wrong!")
	}
	if floatVal := gorest_request.GetFloatValueOrDefault("float",-1); floatVal != 12.4 {
		t.Errorf("GetFloatValue Wrong!")
	}
	if uVal := gorest_request.GetUintValueOrDefault("int",0); uVal != 2 {
		t.Errorf("GetUintValue Wrong!")
	}
	if boolVal := gorest_request.GetBoolValueOrDefault("bool",false); boolVal != true {
		t.Errorf("GetBoolValue Wrong!")
	}
}
func TestDELETEMethod(t *testing.T) {
	req, _ := http.NewRequest("DEL", "http://www.google.com/search?str=foo&int=2&float=12.4&bool=1",nil)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded; param=value")


	gorest_request := &Request{req,nil}

	if q := gorest_request.GetStringValueOrDefault("str",""); q != "foo" {
		t.Errorf("GetStringValue Wrong!")
	}
	if ival := gorest_request.GetIntValueOrDefault("int",0); ival != 2 {
		t.Errorf("GetIntValue Wrong!")
	}
	if floatVal := gorest_request.GetFloatValueOrDefault("float",-1); floatVal != 12.4 {
		t.Errorf("GetFloatValue Wrong!")
	}
	if uVal := gorest_request.GetUintValueOrDefault("int",0); uVal != 2 {
		t.Errorf("GetUintValue Wrong!")
	}
	if boolVal := gorest_request.GetBoolValueOrDefault("bool",false); boolVal != true {
		t.Errorf("GetBoolValue Wrong!")
	}
}


