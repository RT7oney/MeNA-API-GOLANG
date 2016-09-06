package models

import (
	"MeNA-Api/common"
	"MeNA-Api/models/database"
	"errors"
	"github.com/astaxie/beego/orm"
	"time"
)

func AddMusicTag(content string) (tag_id int64, err error) {
	var db_orm = orm.NewOrm() //orm不能定义为全局的
	music_tag := &db.AllMusicTags{
		TagName: content,
	}
	err = db_orm.Read(music_tag, "TagName")
	if err != nil || music_tag.TagName == "" {
		music_tag = &db.AllMusicTags{
			TagName: content,
			Created: time.Now(),
			Update:  time.Now(),
		}
		v, e := db_orm.Insert(music_tag)
		if e != nil {
			common.WriteLog("add new music_tag error @MusicMod.AddMusicTag() : " + err.Error())
			return 0, e
		}
		return v, nil
	}
	return -1, nil // 已经存在该tag不能添加重复的tag

}

func GetOneMusicTag(tag_id int64) (music_tag db.AllMusicTags, err error) {
	var db_orm = orm.NewOrm()
	music_tag = db.AllMusicTags{
		Id: tag_id,
	}
	err = db_orm.Read(&music_tag)
	if err != nil {
		// common.WriteLog("get one character error @characterMod.GetOneCharacter() : " + e.Error())
		return music_tag, errors.New("Music-Tag-Id Not Exist")
	}
	return music_tag, nil
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
