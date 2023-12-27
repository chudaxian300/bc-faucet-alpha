package utils

import (
	"api/conf"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/patrickmn/go-cache"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
	"time"
)

var Cache *cache.Cache = cache.New(time.Duration(conf.Conf.Cache.Expire)*time.Second, time.Duration(conf.Conf.Cache.Clear)*time.Second)

const DataFormat = "2006-01-02 15:04:05"

func GetDate(format string) string {
	return time.Now().Format(format)
}

func CheckErr(err error) {
	if err != nil {
		panic(err)
	}
}

func CreateDir(filePath string) error {
	if !IsExist(filePath) {
		err := os.MkdirAll(filePath, os.ModePerm) //os.ModePerm -> 文件打开格式
		return err
	}
	return nil
}

func IsExist(path string) bool {
	_, err := os.Stat(path) //os.Stat -> 获取文件信息
	if err != nil {
		if os.IsExist(err) { //目录已存在返回false
			return true
		}
	}
	return false
}

func Md5(s string) string {
	h := md5.New()
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))
}

func GetFileContent(address, dirname string) (string, error) {
	data, err := ioutil.ReadDir(dirname) //得到路径中所有文件名
	if err != nil {
		fmt.Println("read dir err", err)
		return "", err
	}
	for _, v := range data {
		if strings.Index(strings.ToLower(v.Name()), strings.ToLower(address)[2:]) > 0 {
			//代表找到文件
			file, err := os.Open(dirname + "\\" + v.Name())
			if err != nil {
				fmt.Printf("Failed to open file %v, err === %v\n", v.Name(), err)
				return "", err
			}
			data := make([]byte, 1024)
			count, err := file.Read(data)
			if err != nil || count < 1 {
				fmt.Printf("Failed to read file %v, err === %v\n", v.Name(), err)
				return "", err
			}
			return string(data[:count]), nil
		}
	}
	return "", nil
}

func StrVal(value interface{}) string {
	var key string
	if value == nil {
		return key
	}

	switch value.(type) {
	case float64:
		ft := value.(float64)
		key = strconv.FormatFloat(ft, 'f', -1, 64)
	case float32:
		ft := value.(float32)
		key = strconv.FormatFloat(float64(ft), 'f', -1, 64)
	case int:
		it := value.(int)
		key = strconv.Itoa(it)
	case uint:
		it := value.(uint)
		key = strconv.Itoa(int(it))
	case int8:
		it := value.(int8)
		key = strconv.Itoa(int(it))
	case uint8:
		it := value.(uint8)
		key = strconv.Itoa(int(it))
	case int16:
		it := value.(int16)
		key = strconv.Itoa(int(it))
	case uint16:
		it := value.(uint16)
		key = strconv.Itoa(int(it))
	case int32:
		it := value.(int32)
		key = strconv.Itoa(int(it))
	case uint32:
		it := value.(uint32)
		key = strconv.Itoa(int(it))
	case int64:
		it := value.(int64)
		key = strconv.FormatInt(it, 10)
	case uint64:
		it := value.(uint64)
		key = strconv.FormatUint(it, 10)
	case string:
		key = value.(string)
	case []byte:
		key = string(value.([]byte))
	default:
		newValue, _ := json.Marshal(value)
		key = string(newValue)
	}
	return key
}
