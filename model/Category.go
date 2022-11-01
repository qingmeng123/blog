package model

import (
	"duryun-blog/utils/errmsg"
	"gorm.io/gorm"
)

type Category struct {
	ID   uint   `gorm:"primary_key;auto_increment" json:"id"`
	Name string `gorm:"type:varchar(20);not null" json:"name"`
}

// CheckCategory 查询分类是否存在
func CheckCategory(name string) (code int) {
	var cate Category
	Db.Select("id").Where("name = ?", name).First(&cate)
	if cate.ID > 0 {
		return errmsg.ERROR_CATENAME_EXIST //2001
	}
	return errmsg.SUCCESS
}

// CreateCate 新增分类
func CreateCate(data *Category) int {
	err := Db.Create(&data).Error
	if err != nil {
		return errmsg.ERROR // 500
	}
	return errmsg.SUCCESS
}

// GetCateInfo 查询单个分类信息
func GetCateInfo(id int) (Category, int) {
	var cate Category
	Db.Where("id = ?", id).First(&cate)
	return cate, errmsg.SUCCESS
}

// GetCate 查询分类列表
func GetCate(pageSize int, pageNum int) ([]Category, int64) {
	var cate []Category
	var total int64
	Err = Db.Find(&cate).Limit(pageSize).Offset((pageNum - 1) * pageSize).Error
	Db.Model(&cate).Count(&total)
	if Err != nil && Err != gorm.ErrRecordNotFound {
		return nil, 0
	}
	return cate, total
}

// EditCate 编辑分类信息
func EditCate(id int, data *Category) int {
	var cate Category
	var maps = make(map[string]interface{})
	maps["name"] = data.Name

	Err = Db.Model(&cate).Where("id = ? ", id).Updates(maps).Error
	if Err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// DeleteCate 删除分类
func DeleteCate(id int) int {
	var cate Category
	Err = Db.Where("id = ? ", id).Delete(&cate).Error
	if Err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}
