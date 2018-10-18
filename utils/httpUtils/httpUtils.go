package httpUtils

import (
	"net/http"
	"io/ioutil"
	"bytes"
	"encoding/json"
		"reflect"
	"strconv"
	"fmt"
	"strings"
)

/**
获取body的data(json)转换为string
*字节数据转string
*/
func GetDataString(req *http.Request) string {
	result, err := ioutil.ReadAll(req.Body)
	if err != nil {
		panic(err)
	} else {
		return bytes.NewBuffer(result).String()
	}
}

/**
将http 请求中的参数转到结构体
 */
func GetJsonToEntity(req *http.Request, b interface{}) (interface{}) {
	dataString := GetDataString(req)
	err := json.Unmarshal([]byte(dataString), &b)
	if err != nil {
		panic(err)
	}
	return b
}

/**
将json字符串转为map
 */
func GetHtppJsonToMap(req *http.Request) (httpMap map[string]interface{}) {
	jsonStr := GetDataString(req)
	var b interface{}
	err := json.Unmarshal([]byte(jsonStr), &b)
	if err != nil {
		panic(err)
	}
	return b.(map[string]interface{})
}

/**
将map转为struct
 */
func MapToStruct(m map[string]interface{} , pointer interface{})  {
	/*err := mapstructure.Decode(a, &b)
	if err != nil {
		panic(err)
	}*/

	// reflect.Ptr类型 *main.Person
	pointertype := reflect.TypeOf(pointer)
	// reflect.Value类型
	pointervalue := reflect.ValueOf(pointer)
	// struct类型  main.Person
	structType := pointertype.Elem()
	// 遍历结构体字段
	for i := 0; i < structType.NumField(); i++ {
		// 获取指定字段的反射值
		f := pointervalue.Elem().Field(i)
		//获取struct的指定字段
		stf := structType.Field(i)
		// 获取tag
		/*fmt.Println(stf.Tag.Get("json"))
		if splitS := strings.Split(stf.Tag.Get("json:"), ",") ; len(splitS) >0 && splitS[0] != "" {

			strings.Trim(splitS[0]," ")
		}*/

		name := strings.Split(stf.Tag.Get("json"), ",")[0]
		// 判断是否为忽略字段
		if name == "-" {
			continue
		}
		// 判断是否为空，若为空则使用字段本身的名称获取value值
		if name == "" {
			name =strings.ToLower( stf.Name)
		}
		//获取value值
		v, ok := m[name]
		if !ok {
			continue
		}

		//获取指定字段的类型
		kind := pointervalue.Elem().Field(i).Kind()
		// 若字段为指针类型
		if kind == reflect.Ptr {
			// 获取对应字段的kind
			kind = f.Type().Elem().Kind()
		}
		// 设置对应字段的值
		switch kind {
		case reflect.Int:
			res, _ := strconv.ParseInt(fmt.Sprint(v), 10, 64)
			pointervalue.Elem().Field(i).SetInt(res)

		case reflect.String:
			pointervalue.Elem().Field(i).SetString(fmt.Sprint(v))
		}
	}
}
