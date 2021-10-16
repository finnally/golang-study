package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"reflect"
	"strconv"
	"strings"
)

// mysql config
type MysqlConfig struct {
	Address  string `ini:"address"`
	Port     int    `ini:"port"`
	Username string `ini:"username"`
	Password string `ini:"password"`
}

// redis config
type RedisConfig struct {
	Host     string `ini:"host"`
	Port     int    `ini:"port"`
	Password string `ini:"password"`
	Database int    `ini:"database"`
}

type Config struct {
	MysqlConfig `ini:"mysql"`
	RedisConfig `ini:"redis"`
}

func loadIni(fileName string, data interface{}) (err error) {
	// 1. 参数较验
	// 1.1 传入的data参数必须是指针类型（需要在函数中对其赋值）
	t := reflect.TypeOf(data)
	fmt.Println("接收参数data结构体名称为:", t.Elem().Name())
	if t.Kind() != reflect.Ptr {
		err = errors.New("data param should be a pointer")
		return
	}
	// 1.2 传入的data参数必须为结构体类型（配置文件中键值对需要赋值给结构体的字段）
	if t.Elem().Kind() != reflect.Struct {
		err = errors.New("data param should be a struct pointer")
		return
	}
	// 2. 读文件得到字节类型数据,将字节类型数据转换成字符串
	b, err := ioutil.ReadFile(fileName)
	if err != nil {
		return
	}
	lineSlice := strings.Split(string(b), "\r\n")
	// fmt.Printf("%#v\n", lineSlice)
	// 3. 逐行读取数据
	var structName string
	for idx, line := range lineSlice {
		// 去掉字符串首尾空格
		line = strings.TrimSpace(line)
		// 跳过空行
		if len(line) == 0 {
			continue
		}
		// 3.1 跳过注释
		if strings.HasPrefix(line, ";") || strings.HasPrefix(line, "#") {
			continue
		}
		// 3.2 "["开头表示节(section)
		if strings.HasPrefix(line, "[") {
			// fmt.Println(line[0], '[', line[len(line)-1])
			if line[0] != '[' || line[len(line)-1] != ']' {
				err = fmt.Errorf("line:%d syntax error", idx+1)
				return
			}
			// 获取[]里内容并将首尾空格去掉
			sectionName := strings.TrimSpace(line[1 : len(line)-1])
			if len(sectionName) == 0 {
				err = fmt.Errorf("line:%d syntax error", idx+1)
				return
			}
			// 根据字符串sectionName去data里通过反射找到对应的结构体
			for i := 0; i < t.Elem().NumField(); i++ {
				// 获取结构体Config内字段：MysqlConfig RedisConfig
				field := t.Elem().Field(i)
				// 获取结构体Config内tag：mysql redis，判断与sectionName是否相等
				if sectionName == field.Tag.Get("ini") {
					// 获取结构体Config内字段名称：MysqlConfig RedisConfig
					structName = field.Name
					fmt.Printf("%s对应的嵌套结构体：%s\n", sectionName, structName)
				}
			}
		} else {
			// 3.3 非"["开头则为"="分割的键值对
			// 3.3.1 以"="分割，左右为key，右边为value
			if !strings.Contains(line, "=") {
				err = fmt.Errorf("line:%d syntax error", idx+1)
				return
			}
			index := strings.Index(line, "=")
			key := strings.TrimSpace(line[:index])
			value := strings.TrimSpace(line[index+1:])
			// 3.3.2 根据structName去data里把对应的嵌套结构体取出
			v := reflect.ValueOf(data)
			fmt.Println("嵌套结构体内值：", v)
			// 获取嵌套结构体的值信息
			sValue := v.Elem().FieldByName(structName)
			fmt.Println("结构体内值：", sValue)
			// 获取嵌套结构体内值的类型，Config结构体内MysqlConfig值类型为MysqlConfig struct
			sType := sValue.Type()
			fmt.Println("嵌套结构体内值类型为：", sType)
			// 嵌套结构体Config内值类型必须为Struct
			if sType.Kind() != reflect.Struct {
				err = fmt.Errorf("%s must be a struct", structName)
				return
			}
			var fieldName string
			var fileType reflect.StructField
			// 3.3.3 遍历结构体（MysqlConfig或RedisConfig）的每一个字段，判断tag是否等于key
			for i := 0; i < sValue.NumField(); i++ {
				field := sType.Field(i)
				fileType = field
				if field.Tag.Get("ini") == key {
					// 找到对应的字段
					fmt.Println("字段：", field.Name)
					fieldName = field.Name
					break
				}
			}
			// 3.3.4 如果key = tag，给这个字段赋值
			// 4. 根据fileName取出这个字段
			if len(fileName) == 0 {
				// 结构体中找不到对应的字符
				continue
			}
			fileObj := sValue.FieldByName(fieldName)
			fmt.Println("fileObj:", fileObj)
			// 5. 赋值
			// fmt.Println(fieldName, fileType.Type.Kind())
			switch fileType.Type.Kind() {
			case reflect.String:
				fmt.Printf("将%s写入到fileObj\n", value)
				fileObj.SetString(value)
			case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
				fmt.Printf("将%s写入到fileObj\n", value)
				var valueInt int64
				valueInt, err = strconv.ParseInt(value, 10, 64)
				if err != nil {
					err = fmt.Errorf("line:%d value type error", idx+1)
					return
				}
				fileObj.SetInt(valueInt)
			}
		}
	}
	return
}

const (
	test = "dd"
)

func main() {
	fmt.Print(reflect.TypeOf(test))
	// var cfg Config
	// err := loadIni("./reflect_demo/conf.ini", &cfg)
	// if err != nil {
	// 	fmt.Printf("load ini faild, err:%v", err)
	// 	return
	// }
	// fmt.Printf("%#v", cfg)
}

// func LoadIni(b []byte, v interface{}) error {
// 	m := MysqlConfig{}
// 	f, err := os.OpenFile("reflect_demo/mysql.ini", os.O_RDONLY, 0644)
// 	if err != nil {
// 		panic("file not found")
// 	}
// 	defer f.Close()
// 	buf := bufio.NewReader(f)
// 	for {
// 		line, err := buf.ReadString('\n')
// 		line = strings.TrimSpace(line)
// 		if err != nil {
// 			if err == io.EOF {
// 				return nil
// 			}
// 			return err
// 		}
// 		t := reflect.TypeOf(v)
// 		for i := 0; i < t.NumField(); i++ {
// 			stuct_field := t.Field(i)
// 			ini_field := strings.Split(line, "=")
// 			if stuct_field.Tag.Get("ini") == ini_field[0] {
// 				_, value := ini_field[0], ini_field[1]
// 				m.Address = value
// 				fmt.Println(m)
// 			}
// 			// fmt.Printf("name:%s index:%d type:%v json tag:%v\n", stuct_field.Name, stuct_field.Index, stuct_field.Type, stuct_field.Tag.Get("ini"))
// 		}
// 		// fmt.Println(reflect.ValueOf(v).FieldByName(line).IsValid())
// 	}
// }

// func main() {
// 	M := MysqlConfig{}
// 	LoadMysqlIni(M)
// }
