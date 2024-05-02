package tool

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
)

func ObjectToJson(src interface{}) (string, error) {
	if result, err := json.Marshal(src); err != nil {
		return "", errors.New("Json Str Parse Err: " + err.Error())
	} else {
		return string(result), nil
	}
}

func JsonToObject(src string, target interface{}) error {
	if err := json.Unmarshal([]byte(src), target); err != nil {
		return errors.New("Json Str Parse Err: " + err.Error())
	}
	return nil
}

func JsonRawToObject(src string, target interface{}) error {
	if err := json.Unmarshal([]byte(json.RawMessage(src)), target); err != nil {
		return errors.New("Json Str Parse Err: " + err.Error())
	}
	return nil
}

func JsonToAny(src interface{}, target interface{}) error {
	if src == nil || target == nil {
		return errors.New("Param is empty")
	}
	str, err := ObjectToJson(src)
	if err != nil {
		return err
	}
	if err := JsonToObject(str, target); err != nil {
		return err
	}
	return nil
}

func JsonToObject2(src string, target interface{}) error {
	d := json.NewDecoder(bytes.NewBuffer([]byte(src)))
	d.UseNumber()
	if err := d.Decode(target); err != nil {
		return errors.New("Json Str Parse Err: " + err.Error())
	}
	return nil
}

func JsonToAny2(src interface{}, target interface{}) error {
	if src == nil || target == nil {
		return errors.New("Param is empty")
	}
	str, err := ObjectToJson(src)
	if err != nil {
		return err
	}
	if err := JsonToObject2(str, target); err != nil {
		return err
	}
	return nil
}

func AnyToStr(any interface{}) string {
	if any == nil {
		return ""
	}
	if str, ok := any.(string); ok {
		return str
	} else if str, ok := any.(int); ok {
		return strconv.FormatInt(int64(str), 10)
	} else if str, ok := any.(int8); ok {
		return strconv.FormatInt(int64(str), 10)
	} else if str, ok := any.(int16); ok {
		return strconv.FormatInt(int64(str), 10)
	} else if str, ok := any.(int32); ok {
		return strconv.FormatInt(int64(str), 10)
	} else if str, ok := any.(int64); ok {
		return strconv.FormatInt(int64(str), 10)
	} else if str, ok := any.(float32); ok {
		return strconv.FormatFloat(float64(str), 'f', 0, 64)
	} else if str, ok := any.(float64); ok {
		return strconv.FormatFloat(float64(str), 'f', 0, 64)
	} else if str, ok := any.(uint); ok {
		return strconv.FormatUint(uint64(str), 10)
	} else if str, ok := any.(uint8); ok {
		return strconv.FormatUint(uint64(str), 10)
	} else if str, ok := any.(uint16); ok {
		return strconv.FormatUint(uint64(str), 10)
	} else if str, ok := any.(uint32); ok {
		return strconv.FormatUint(uint64(str), 10)
	} else if str, ok := any.(uint64); ok {
		return strconv.FormatUint(uint64(str), 10)
	} else if str, ok := any.(bool); ok {
		if str {
			return "True"
		}
		return "False"
	} else {
		if ret, err := ObjectToJson(any); err != nil {
			return ""
		} else {
			return ret
		}
	}
	return ""
}
func WriteJSON(fileName string, data map[string]interface{}) error {
	jsonData, err := json.MarshalIndent(data, "", "    ")
	if err != nil {
		fmt.Println("JSON encoding error:", err)
		return err
	}
	err = os.WriteFile(fileName, jsonData, 0644)
	if err != nil {
		fmt.Println("Write file error:", err)
		return err
	}
	return nil
}
func ReadJSON(fileName string) (map[string]interface{}, error) {
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Open file error:", err)
		return nil, err
	}
	defer file.Close()
	jsonData, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println("Read file error:", err)
		return nil, err
	}
	data := make(map[string]interface{})
	err = json.Unmarshal(jsonData, &data)
	if err != nil {
		fmt.Println("JSON decoding error:", err)
		return nil, err
	}
	return data, nil
}
