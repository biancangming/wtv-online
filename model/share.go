package model

import (
	"github.com/duke-git/lancet/v2/random"
)

type Share struct {
	Uuid      string `json:"uuid" gorm:"type:varchar(100);primaryKey"`
	Content   string `json:"content" gorm:"type:text"` // 填入的内容
	FileType  string `json:"fileType" gorm:"type:text"`
	Mark      string `json:"mark" gorm:"type:text"` // 备注
	UseStatus int    `json:"useStatus" gorm:"type:text"`
}

func GetShareUrl(uuid string) (Share, error) {
	s := Share{}
	var err error
	err = dbs.Debug().Model(&Share{}).Select([]string{"content", "uuid", "use_status", "file_type", "mark"}).Find(&s, "uuid = ?", uuid).Limit(1).Error

	if err != nil {
		return s, err
	}
	return s, nil
}

func GetShareUrls(us int) ([]Share, error) {
	s := new([]Share)
	if err := dbs.Model(&Share{}).
		Select([]string{"uuid", "file_type", "mark"}).
		Find(&s, "use_status = ?", us).Error; err != nil {
		return *s, err
	}
	return *s, nil
}

func UpdateOrAddShare(s Share) (Share, error) {
	var err error

	if s.Uuid == "" {
		s.Uuid, _ = random.UUIdV4()
		s.UseStatus = 0
		// 存储内容到sqlite
		if err := dbs.Model(&Share{}).Create(s).Error; err != nil {
			return s, err
		}
		return s, nil
	} else {
		// 更新数据到sqlite
		if err := dbs.Model(&Share{}).Where("uuid = ?", s.Uuid).Updates(s).Error; err != nil {
			return s, err
		}
	}

	if err != nil {
		return s, err
	}

	return s, nil
}

// UpdateUseStatus 修改使用状态
func UpdateUseStatus(s Share) (Share, error) {
	table := dbs.Model(&Share{})
	if err := table.Where("uuid = ?", s.Uuid).Update("use_status", s.UseStatus).Error; err != nil {
		return s, err
	}
	return s, nil
}
