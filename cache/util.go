package cache

import (
	"encoding/json"
	"log"
	"regexp"
	"strconv"
	"strings"
	_ "strings"
)

const (
	B = 1 << (10 * iota)
	KB
	MB
	GB
	TB
	PB
)

func ParseSize(size string) (int64, string) {
	//默认为100MB
	re, _ := regexp.Compile(`[0-9]+`)                       //定义正则表达式规则
	uint := string(re.ReplaceAll([]byte(size), []byte(""))) //单位 将数字部分用空字符替换
	//num := string(re.FindStringSubmatch(size)) 返回的是一个切片。。。
	num, _ := strconv.ParseInt(strings.Replace(size, uint, "", 1), 10, 64)
	var byteNum int64 = 0
	switch uint {
	case "B":
		byteNum = num
	case "KB":
		byteNum = num * KB
	case "MB":
		byteNum = num * MB
	case "GB":
		byteNum = num * GB
	case "TB":
		byteNum = num * TB
	default:
		num = 0
	}
	if num == 0 {
		log.Println("ParseSize 仅支持 B、KB、MB、GB、TB、PB")
		num = 100
		byteNum = num * MB
		uint = "Mb"
	}
	sizeStr := strconv.FormatInt(num, 10) + uint
	return byteNum, sizeStr
}

func GetValSize(val interface{}) int64 {
	bytes, _ := json.Marshal(val)
	size := int64(len(bytes))
	return size
}
