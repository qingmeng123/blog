/*******
* @Author:qingmeng
* @Description:
* @File:category
* @Date:2022/11/1
 */

package service

import (
	"duryun-blog/dao"
	"duryun-blog/model"
	"duryun-blog/utils/errmsg"
	"gorm.io/gorm"
)

type CategoryService struct {
}

func (s *CategoryService) CheckCategoryByName(name string) int {
	cd := dao.CategoryDao{}
	_, err := cd.GetCategoryByName(name)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return errmsg.ERROR_CATE_NOT_EXIST
		}
		return errmsg.ERRDAO
	}
	return errmsg.ERROR_CATENAME_EXIST
}

func (s *CategoryService) CreateCate(cate *model.Category) int {
	cd := dao.CategoryDao{}
	code := s.CheckCategoryByName(cate.Name)
	if code == errmsg.ERROR_CATE_NOT_EXIST {
		err := cd.CreateCate(cate)
		if err != nil {
			return errmsg.ERRDAO
		}
		return errmsg.SUCCESS
	}
	return errmsg.ERROR_CATENAME_EXIST
}

func (s *CategoryService) GetCateInfo(id int) (cate model.Category, code int) {
	cd := dao.CategoryDao{}
	err := cd.GetCategory(id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return cate, errmsg.ERROR_CATE_NOT_EXIST
		}
		return cate, errmsg.ERRDAO
	}
	return cate, errmsg.SUCCESS
}

func (s *CategoryService) GetCateList(pageSize int, pageNum int) ([]model.Category, int64) {
	cd := dao.CategoryDao{}
	cate, total, err := cd.GetCateList(pageSize, pageNum)
	if err != nil {
		return nil, 0
	}
	return cate, total
}

func (s *CategoryService) UpdateCate(id int, cate *model.Category) int {
	cd := dao.CategoryDao{}
	code := s.CheckCategoryByName(cate.Name)
	if code == errmsg.ERROR_CATE_NOT_EXIST {
		err := cd.UpdateCate(id, cate)
		if err != nil {
			return errmsg.ERRDAO
		}
		return errmsg.SUCCESS
	}
	return errmsg.ERROR_CATENAME_EXIST
}

func (s *CategoryService) DeleteCate(id int) int {
	cd := dao.CategoryDao{}
	err := cd.DeleteCate(id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return errmsg.ERROR_CATE_NOT_EXIST
		}
		return errmsg.ERRDAO
	}
	return errmsg.SUCCESS
}
