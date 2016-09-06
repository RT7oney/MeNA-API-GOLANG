package models

import (
	"MeNA-Api/common"
	"MeNA-Api/models/database"
	"errors"
	// "fmt"
	"github.com/astaxie/beego/orm"
	// "time"
)

/**
 * 之后注释该代码，因为九型人格添加完成后不需要再添加了
 * @param {[type]} content string) (character_id int64, err error [description]
 */
func AddCharacter(content string) (character_id int64, err error) {
	var db_orm = orm.NewOrm() //orm不能定义为全局的
	var res []orm.Params
	num, err := db_orm.Raw("SELECT * FROM character").Values(&res)
	if err == nil && num >= 9 {
		return -1, nil
	}
	character := &db.Character{
		Introduction: content,
	}
	v, err := db_orm.Insert(character)
	if err != nil {
		common.WriteLog("add new character error @characterMod.AddOne() : " + err.Error())
		return 0, err
	}
	return v, nil
}

func GetOneCharacter(character_id int64) (character db.Character, err error) {
	var db_orm = orm.NewOrm()
	character = db.Character{
		Id: character_id,
	}
	err = db_orm.Read(&character)
	if err != nil {
		// common.WriteLog("get one character error @characterMod.GetOneCharacter() : " + e.Error())
		return character, errors.New("CharacterId Not Exist")
	}
	return character, nil
}

// func GetAllCharacter() map[string]*db.Character {
// 	return Characters
// }

// func UpdateCharacter(character_id string, content string) (err error) {
// 	if v, ok := Characters[CharacterId]; ok {
// 		v.Score = Score
// 		return nil
// 	}
// 	return errors.New("CharacterId Not Exist")
// }

// func DeleteCharacter(character_id string) {
// 	delete(Characters, CharacterId)
// }
