package common

import (
	"bytes"
	"crypto/cipher"
	"crypto/des"
	// "crypto/ras"
	// "crypto/x509"
	"encoding/base64"
	"encoding/json"
	// "encoding/pem"
	// "fmt"
	// "log"
	"github.com/astaxie/beego"
	"io"
	"net"
	"os"
	"regexp"
	"strconv"
	"time"
	// "strings"
	// "bufio"
	// "reflect"
)

/**
 * 请求组件的方法
 */
func Request(url string, data map[string]string) (ret string) {
	var msg = make([]byte, 1024) // 创建一个字节切片去接收服务器返回的消息
	jsonstr, err := json.Marshal(data)
	if err != nil {
		WriteLog("json 转化成 map 系统错误 : (pkg@common)")
		// log.Fatal("json to map error : %d\n", err)
		ret = Response(503, "json转化出错")
	}
	// fmt.Println(reflect.TypeOf(string(jsonstr)))
	// fmt.Println(string(jsonstr))
	/**
	 * 建立socket连接，并向其写入数据
	 */
	conn, err := net.Dial("tcp", url)
	if err != nil {
		WriteLog("Dial 方法想服务器拨号错误 : (pkg@common)" + err.Error())
		// log.Fatal("Dial server : %d\n", err)
		ret = Response(500, "连接服务器失败")
	}
	in, err := conn.Write([]byte(jsonstr))
	if err != nil {
		WriteLog("向服务器发送消息失败 : (pkg@common)" + strconv.Itoa(in) + "error: " + err.Error())
		// log.Fatal("Error when send to server : %d\n", in)
		ret = Response(501, "向服务器发送消息失败")
	}
	length, err := conn.Read(msg)
	if err != nil {
		WriteLog("从服务器接收消息失败 : (pkg@common)" + err.Error())
		// log.Fatal("Error when read msg from server : %d\n", err)
		ret = Response(502, "从服务器接收消息失败")
	}
	// ret = string(msg[0:length])
	encode_msg, _ := EncodeData(msg[0:length])
	the_data := base64.StdEncoding.EncodeToString(encode_msg)
	ret = Response(200, the_data)
	return
}

/**
 * 处理组件响应的方法
 */
func Response(code int, data string) string {
	var ret = make(map[string]string)
	ret["code"] = strconv.Itoa(code)
	ret["msg"] = data
	jsonret, _ := json.Marshal(ret)
	return string(jsonret)
}

/**
 * 数据加密方法 采用RSA方式
 */
func EncodeData(data []byte) ([]byte, error) {
	/*##################des##################*/
	key_str := beego.AppConfig.String("DesKey")
	key := []byte(key_str)
	block, err := des.NewCipher(key)
	if err != nil {
		WriteLog("加密消息出错@common.EncodeData:" + err.Error())
		return nil, err
	}
	data = PKCS5Padding(data, block.BlockSize())
	// data = ZeroPadding(data, block.BlockSize())
	blockMode := cipher.NewCBCEncrypter(block, key)
	crypted := make([]byte, len(data))

	// 根据CryptBlocks方法的说明，如下方式初始化crypted也可以
	// crypted := data

	blockMode.CryptBlocks(crypted, data)

	return crypted, nil
	/*########################################*/
}

/**
 * 数据解密方法
 */
func DecodeData(crypted, key []byte) ([]byte, error) {
	/*##################des##################*/
	block, err := des.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockMode := cipher.NewCBCDecrypter(block, key)
	origData := make([]byte, len(crypted))
	// origData := crypted
	blockMode.CryptBlocks(origData, crypted)

	origData = PKCS5UnPadding(origData)

	// origData = ZeroUnPadding(origData)

	return origData, nil
	/*########################################*/
}

func PKCS5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}
func PKCS5UnPadding(origData []byte) []byte {
	length := len(origData)
	// 去掉最后一个字节 unpadding 次
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}

/**
 * 判断是不是邮箱
 */
func IsEmail(email string) bool {
	pattern := `\w+([-+.]\w+)*@\w+([-.]\w+)*\.\w+([-.]\w+)*`
	reg := regexp.MustCompile(pattern)
	ret := reg.FindAllString(email, -1)
	if ret != nil {
		return true
	} else {
		return false
	}
}

/**
 * 记录log的方法
 */
func WriteLog(log string) {
	var f *os.File
	var err error
	log_file := "log/" + strconv.Itoa(time.Now().Year()) + time.Now().Month().String() + ".log"
	if checkFileIsExist(log_file) { //如果文件存在
		f, err = os.OpenFile(log_file, os.O_APPEND|os.O_WRONLY, os.ModeAppend) //打开文件
		check(err)
		// fmt.Println("文件存在")
	} else {
		f, err = os.Create(log_file) //创建文件
		check(err)
		// fmt.Println("文件不存在")
	}
	_, err = io.WriteString(f, "[@"+time.Now().String()+"#]------"+log+"\r\n") //写入文件(字符串)
	check(err)
	// fmt.Printf("写入 %d 个字节n", n)
}

/**
 * 判断文件是否存在  存在返回 true 不存在返回false
 */
func checkFileIsExist(filename string) bool {
	var exist = true
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		exist = false
	}
	return exist
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
