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
	"fmt"
	// "log"
	"crypto/md5"
	"encoding/hex"
	// "github.com/astaxie/beego"
	"io"
	"net"
	"os"
	"regexp"
	"strconv"
	"time"
	// "strings"
	// "bufio"
	// "errors"
	"reflect"
	"unsafe"
)

/**
 * 获取接口app_key
 */
func MakeAppKey(content map[string]string) (app_key string) {
	timestamp := time.Now().Unix()
	tmpstr := content["api_token"] + content["dev_name"] + strconv.FormatInt(timestamp, 10)
	h := md5.New()
	h.Write([]byte(tmpstr))
	app_key_str := hex.EncodeToString(h.Sum(nil))
	app_key = Substr(app_key_str, 0, 8)
	return
}

/**
 * 检验接口权限
 */
func CheckSign(check map[string]string, app_key map[string]string) int {
	timenow := time.Now().Unix()
	expire_time, _ := strconv.ParseInt(app_key["expire_time"], 10, 64)
	// fmt.Println("==============过期时间================")
	// fmt.Println(expire_time)
	if expire_time < timenow {
		return -1 //app_key过期
	}
	tmpstr := app_key["app_key"] + check["timestamp"]
	h := md5.New()
	h.Write([]byte(tmpstr))
	server_sign := hex.EncodeToString(h.Sum(nil))
	if server_sign != check["sign"] {
		return 0 // 签名不匹配
	}
	return 1 //成功
}

/**
 * 请求组件的方法
 */
func Request(url string, data map[string]string, app_key string) map[string]string {
	var msg = make([]byte, 1024) // 创建一个字节切片去接收服务器返回的消息
	var ret = make(map[string]string)
	jsonstr, err := json.Marshal(data)
	if err != nil {
		WriteLog("json to map sync error @common.Request")
		// log.Fatal("json to map error : %d\n", err)
		// ret = Response(503, "json转化出错")
		ret["code"] = "503"
		ret["msg"] = "json转化出错"
	}
	// fmt.Println(reflect.TypeOf(string(jsonstr)))
	// fmt.Println(string(jsonstr))
	/**
	 * 建立socket连接，并向其写入数据
	 */
	conn, err := net.Dial("tcp", url)
	if err != nil {
		WriteLog("Dial the server error @common.Request" + err.Error())
		// log.Fatal("Dial server : %d\n", err)
		// ret = Response(500, "连接服务器失败")
		ret["code"] = "500"
		ret["msg"] = "连接服务器失败"
	}
	in, err := conn.Write([]byte(jsonstr))
	if err != nil {
		WriteLog("send msg to server error @common.Request" + strconv.Itoa(in) + "error: " + err.Error())
		// log.Fatal("Error when send to server : %d\n", in)
		// ret = Response(501, "向服务器发送消息失败")
		ret["code"] = "501"
		ret["msg"] = "向服务器发送消息失败"
	}
	length, err := conn.Read(msg)
	if err != nil {
		WriteLog("get msg from server error @common.Request" + err.Error())
		// log.Fatal("Error when read msg from server : %d\n", err)
		// ret = Response(502, "从服务器接收消息失败")
		ret["code"] = "502"
		ret["msg"] = "从服务器接收消息失败"
	}
	// ret = string(msg[0:length])
	if app_key != "" {
		encode_msg, _ := EncodeData(msg[0:length], app_key)
		// ret = Response(200, encode_msg)
		ret["code"] = "200"
		ret["msg"] = encode_msg
	} else {
		// ret = Response(200, string(msg[0:length]))
		ret["code"] = "200"
		ret["msg"] = string(msg[0:length])
	}
	// the_data := base64.StdEncoding.EncodeToString(encode_msg)
	return ret
}

/**
 * 处理组件响应的方法
 */
func Response(code int, msg string, data interface{}) string {
	var ret = make(map[string]interface{})
	ret["code"] = strconv.Itoa(code)
	ret["msg"] = msg
	if data != nil {
		ret["data"] = data
	}
	jsonret, _ := json.Marshal(ret)
	return string(jsonret)
}

/**
 * 数据加密方法 采用RSA方式
 */
func EncodeData(data []byte, key_str string) (string, error) {
	/*##################des##################*/
	// key_str := beego.AppConfig.String("DesKey")
	key := []byte(key_str)
	block, err := des.NewCipher(key)
	if err != nil {
		WriteLog("encode data error @common.EncodeData:" + err.Error())
		return "", err
	}
	data = PKCS5Padding(data, block.BlockSize())
	// data = ZeroPadding(data, block.BlockSize())
	blockMode := cipher.NewCBCEncrypter(block, key)
	crypted := make([]byte, len(data))

	// 根据CryptBlocks方法的说明，如下方式初始化crypted也可以
	// crypted := data

	blockMode.CryptBlocks(crypted, data)
	encode_data := base64.StdEncoding.EncodeToString(crypted)
	return encode_data, nil
	/*########################################*/
}

/**
 * 数据解密方法
 */
func DecodeData(crypted, key []byte) ([]byte, error) {
	/*##################des##################*/
	block, err := des.NewCipher(key)
	if err != nil {
		WriteLog("decode data error @common.EncodeData:" + err.Error())
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
 * 字符串截取
 */
func Substr(str string, start, length int) string {
	rs := []rune(str)
	rl := len(rs)
	end := 0
	if start < 0 {
		start = rl - 1 + start
	}
	end = start + length
	if start > end {
		start, end = end, start
	}
	if start < 0 {
		start = 0
	}
	if start > rl {
		start = rl
	}
	if end < 0 {
		end = 0
	}
	if end > rl {
		end = rl
	}

	return string(rs[start:end])
}

/**
 * 将不同类型的切片转化成[]byte
 */
func ByteSlice(slice interface{}) (data []byte) {
	sv := reflect.ValueOf(slice)
	if sv.Kind() != reflect.Slice {
		WriteLog("ByteSlice called with non-slice value of type @ common.ByteSlice")
		// return nil, errors.New("ByteSlice called with non-slice value of type")
		panic("ByteSlice called with non-slice value of type error")
	}
	h := (*reflect.SliceHeader)((unsafe.Pointer(&data)))
	h.Cap = sv.Cap() * int(sv.Type().Elem().Size())
	h.Len = sv.Len() * int(sv.Type().Elem().Size())
	h.Data = sv.Pointer()
	return
}

/**
 * 将[]x转换成[]y
 */
func SliceTrans(slice interface{}, newSliceType reflect.Type) interface{} {
	sv := reflect.ValueOf(slice)
	if sv.Kind() != reflect.Slice {
		WriteLog("Slice called with non-slice value of type @ common.Slice")
		panic(fmt.Sprintf("Slice called with non-slice value of type %T", slice))
	}
	if newSliceType.Kind() != reflect.Slice {
		WriteLog("Slice called with non-slice value of type @ common.Slice")
		panic(fmt.Sprintf("Slice called with non-slice type of type %T", newSliceType))
	}
	newSlice := reflect.New(newSliceType)
	hdr := (*reflect.SliceHeader)(unsafe.Pointer(newSlice.Pointer()))
	hdr.Cap = sv.Cap() * int(sv.Type().Elem().Size()) / int(newSliceType.Elem().Size())
	hdr.Len = sv.Len() * int(sv.Type().Elem().Size()) / int(newSliceType.Elem().Size())
	hdr.Data = uintptr(sv.Pointer())
	return newSlice.Elem().Interface()
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
