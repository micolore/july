package utils

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strings"
)

func CastStrToMap(jsonStr string) (dat map[string]interface{}) {

	var mapData map[string]interface{}
	json.Unmarshal([]byte(jsonStr), &mapData)
	return mapData
}

func CastByteToMap(data []byte) (dat map[string]interface{}) {
	str := string(data)
	firstIndex := strings.Index(str, ",{")
	lastIndex := strings.LastIndex(str, "}")
	jsonStr := str[firstIndex+1 : lastIndex+1]
	return CastStrToMap(jsonStr)
}

// WriteDataToFile 写数据到文件
func WriteDataToFile(fileName string, content string) error {

	// 以只写的模式，打开文件
	//f, err := os.OpenFile(fileName, os.O_WRONLY, 0644)
	f, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE, 0766)

	if err != nil {
		fmt.Println("file create failed. err: " + err.Error())
	} else {
		// 查找文件末尾的偏移量
		n, _ := f.Seek(0, os.SEEK_END)
		// 从末尾的偏移量开始写入内容
		_, err = f.WriteAt([]byte(content), n)
	}
	defer f.Close()
	return err
}

// ReadLine 读取文件的每一行
func ReadLine(fileName string, handler func(string)) error {
	f, err := os.Open(fileName)
	if err != nil {
		return err
	}
	buf := bufio.NewReader(f)
	for {
		line, err := buf.ReadString('\n')
		line = strings.TrimSpace(line)
		handler(line)
		if err != nil {
			if err == io.EOF {
				return nil
			}
			return err
		}
	}
}
