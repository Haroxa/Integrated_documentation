package model

import "gorm.io/gorm"

type Apply struct {
	Id          int    `json:"id"`
	Carshareid  int    `json:"carshareid"`
	Userid      int    `json:"userid"`
	Username    string `json:"username"`
	Contact     string `json:"contact"`
	Createdtime string `json:"createdtime"`
	Status      string `json:"status"`
}

func CreateApply(apply *Apply) error {
	return db.Create(apply).Error
}

func GetApplyById(ApplyId int) (Apply, error) {
	var apply Apply
	err := db.Where("id=?", ApplyId).First(&apply).Error
	if err == gorm.ErrRecordNotFound {
		err = nil
	}
	return apply, err
}

func GetApplyByTime(time string) ([]Apply, int, error) {
	applys := make([]Apply, 10)
	var c int64
	err := db.Where("Createdtime LIKE ?", "%"+time+"%").Find(&applys).Count(&c).Error
	if err == gorm.ErrRecordNotFound {
		err = nil
	}
	return applys, int(c), err
}

func GetApplyByUser(Userid int) ([]Apply, int, error) {
	applys := make([]Apply, 10)
	var c int64
	err := db.Where("userid=?", Userid).Find(&applys).Count(&c).Error
	if err == gorm.ErrRecordNotFound {
		err = nil
	}
	return applys, int(c), err
}

func GetApplyByCarShare(Carshareid int) ([]Apply, int, error) {
	applys := make([]Apply, 10)
	var c int64
	err := db.Where("carshareid=?", Carshareid).Find(&applys).Count(&c).Error
	if err == gorm.ErrRecordNotFound {
		err = nil
	}
	return applys, int(c), err
}

func GetAllApply() ([]Apply, int, error) {
	applys := make([]Apply, 1000) //Apply模型切片
	var c int64
	err := db.Find(&applys).Count(&c).Error
	//Count函数，直接返回查询匹配的行数。
	return applys, int(c), err
}

func UpdateApply(apply Apply, mp interface{}) error {
	return db.Model(&apply).Updates(mp).Error
}

func DeleteApply(apply Apply) error {
	//删除指定用户以及数据库中的删除记录
	err := db.Unscoped().Delete(&apply).Error
	return err
}

func DeleteAllApply(apply []Apply) error {
	//删除指定用户以及数据库中的删除记录
	err := db.Unscoped().Delete(&apply).Error
	return err
}
