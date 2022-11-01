/*******
* @Author:qingmeng
* @Description:
* @File:category
* @Date:2022/11/1
 */

package dao

import "duryun-blog/model"

type CategoryDao struct {
}

func (d *CategoryDao) GetCategoryByName(name string) (model.Category, error) {
	var cate model.Category
	err := model.Db.Select("id").Where("name = ?", name).First(&cate).Error
	return cate, err
}

func (d *CategoryDao) CreateCate(cate *model.Category) error {
	return model.Db.Create(&cate).Error

}

func (d *CategoryDao) GetCategory(id int) error {
	return model.Db.Where("id = ?", id).First(&model.Category{}).Error

}

func (d *CategoryDao) GetCateList(pageSize int, pageNum int) (cateList []model.Category, total int64, err error) {
	err = model.Db.Find(&cateList).Limit(pageSize).Offset((pageNum - 1) * pageSize).Error
	if err != nil {
		return nil, 0, err
	}
	err = model.Db.Model(&cateList).Count(&total).Error
	return cateList, total, err
}

func (d *CategoryDao) UpdateCate(id int, cate *model.Category) error {
	return model.Db.Model(&cate).Where("id = ? ", id).Updates(cate).Error
}

func (d *CategoryDao) DeleteCate(id int) error {
	return model.Db.Where("id = ? ", id).Delete(&model.Category{}).Error
}
