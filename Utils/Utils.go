package Utils
import (
	"strconv"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"unicode/utf8"
	"reflect"
	"fmt"
)

func SliceString2SliceInterface(strInputs []string) []interface{} {
	vals := make([]interface{}, len(strInputs))
	for i, v := range strInputs {
		vals[i] = v
	}

	return vals
}
func SliceFloat2SliceInterface(strInputs []float64) []interface{} {
	vals := make([]interface{}, len(strInputs))
	for i, v := range strInputs {
		vals[i] = v
	}

	return vals
}
func SliceBool2SliceInterface(strInputs []bool) []interface{} {
	vals := make([]interface{}, len(strInputs))
	for i, v := range strInputs {
		vals[i] = v
	}

	return vals
}

func Int2String(intVal int) string {
	return strconv.Itoa(intVal)
}
func Int642String(intVal int64) string {
	return strconv.FormatInt(intVal,10)
}
func Float2String(intVal float64) string {
	return strconv.FormatFloat(intVal,'f',4,64)
}

func String2IntOrZero(strVal string) int64  {
	val,err:= strconv.ParseInt(strVal,0,64)

	if err!=nil {
		return 0
	} else {
		return val
	}
}
func String2FloatOrZero(strVal string) float64  {
	val,err:= strconv.ParseFloat(strVal,64)

	if err!=nil {
		return 0
	} else {
		return val
	}
}
func String2BoolOrFalse(strVal string) bool  {
	val,err:= strconv.ParseBool(strVal)

	if err!=nil {
		return false
	} else {
		return val
	}
}

func GetMD5Hash(text string) string {
	hasher := md5.New()
	hasher.Write([]byte(text))
	return hex.EncodeToString(hasher.Sum(nil))
}

func TokenToScore(token string) int64 {
	return String2IntOrZero(token)
}

// define how to convert between score to token
func ScoreToToken(score int64) string  {
	return Int642String(score)
}

func Utf8StrLen(str string) int  {
	return utf8.RuneCountInString(str)
}

func ObjectToJson(v interface{}) string  {
	data,err := json.Marshal(v)

	if err != nil {
		return ""
	} else {
		return  string(data)
	}
}
func JsonToObject(strJson string,v interface{}) error {
	err := json.Unmarshal([]byte(strJson),v)

	if err != nil {
		return err
	}
	return nil
}

func BytesToObject(src []byte,v interface{}) error  {
	err := json.Unmarshal(src,v)

	if err != nil {
		return err
	}
	return nil
}

func Deref(t reflect.Type) reflect.Type {
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	return t
}

func BaseType(t reflect.Type, expected reflect.Kind) (reflect.Type, error) {
	t = Deref(t)
	if t.Kind() != expected {
		return nil, fmt.Errorf("expected %s but got %s", expected, t.Kind())
	}
	return t, nil
}
