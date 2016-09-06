package models

import (
	"MeNA-Api/common"
	"MeNA-Api/models/database"
	// "errors"
	"fmt"
	"github.com/astaxie/beego/orm"
	"strconv"
	"time"
)

func AddAppKey(content map[string]string) (id int64, err error) {
	var o = orm.NewOrm() //orm不能定义为全局的
	dev := db.DevAppKey{
		DevName: content["dev_name"],
	}
	err = o.Read(&dev, "DevName") // 读取成功的话 err 为空
	if err != nil || dev.ApiToken == "" {
		app_key := common.MakeAppKey(content)
		timestamp := time.Now().Unix()
		expire_time := strconv.FormatInt(timestamp+604800, 10)
		dev = db.DevAppKey{
			DevName:    content["dev_name"],
			AppKey:     app_key,
			ApiToken:   content["api_token"],
			ExpireTime: expire_time,
			Created:    time.Now(),
			Update:     time.Now(),
		}
		v, err := o.Insert(&dev)
		if err != nil {
			// common.WriteLog("add new dev_key error @DevMod.AddAppKey() : " + err.Error())
			return 0, err // 插入出错
		}
		return v, nil
	}
	return -1, nil // 已经存在数据库，不能insert只能update
}

// func AdminGetAppKey(api_token string) (app_key map[string]string, err error) {
// 	/*####################使用原生sql查询出错#####################*/
// 	type Res struct {
// 		AppKey     string
// 		ExpireTime string
// 	}
// 	var ret = make(map[string]string)
// 	var o = orm.NewOrm()
// 	var res Res
// 	sql := "select app_key,expire_time from dev_app_key where api_token=`" + api_token + "`"
// 	e := o.Raw(sql).QueryRow(&res)
// 	if e == nil {
// 		if res.AppKey != "" {
// 			// fmt.Println(res)
// 			ret["app_key"] = res.AppKey
// 			ret["expire_time"] = res.ExpireTime
// 			return ret, nil
// 		} else {
// 			return nil, errors.New("没有查询到开发者信息") //查询不到api_token的对应的值
// 		}
// 	}
// 	return nil, e
// 	/*####################使用原生sql查询出错#####################*/
// }

func GetAppKey(api_token string) (app_key map[string]string, err error) {
	res := make(map[string]string)
	o := orm.NewOrm()

	dev := db.DevAppKey{
		ApiToken: api_token,
	}

	err = o.Read(&dev, "ApiToken")

	if err != nil {
		return nil, err
	} else {
		res["app_key"] = dev.AppKey
		res["expire_time"] = dev.ExpireTime
		return res, nil
	}
}

// func GetAllCharacter() map[string]*db.Character {
// 	return Characters
// }

func UpdateAppKey(api_token string) (app_key map[string]string, err error) {
	var o = orm.NewOrm()
	dev := db.DevAppKey{
		ApiToken: api_token,
	}
	err = o.Read(&dev, "ApiToken")
	if err != nil || dev.DevName == "" {
		// common.WriteLog("get one character error @characterMod.GetOneCharacter() : " + e.Error())
		return nil, err
	}
	content := make(map[string]string)
	content["dev_name"] = dev.DevName
	content["api_token"] = dev.ApiToken
	app_key_str := common.MakeAppKey(content)
	timestamp := time.Now().Unix()
	expire_time := strconv.FormatInt(timestamp+604800, 10)
	// expire_time := strconv.FormatInt(timestamp+20, 10)

	// dev = db.DevAppKey{
	// 	DevName:    content["dev_name"],
	// 	AppKey:     app_key_str,
	// 	ApiToken:   content["api_token"],
	// 	ExpireTime: expire_time,
	// 	Update:     time.Now(),
	// }

	dev.AppKey = app_key_str
	dev.ExpireTime = expire_time
	dev.Update = time.Now()
	_, e := o.Update(&dev)
	if e != nil {
		fmt.Println(e)
		// common.WriteLog("add new dev_key error @DevMod.AddAppKey() : " + err.Error())
		return nil, e
	}
	res := make(map[string]string)
	res["app_key"] = app_key_str
	res["expire_time"] = expire_time
	return res, nil
}

// func DeleteCharacter(character_id string) {
// 	delete(Characters, CharacterId)
// }
