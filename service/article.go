/*******
* @Author:qingmeng
* @Description:
* @File:article
* @Date:2022/10/31
 */

package service

import (
	"duryun-blog/dao"
	"duryun-blog/model"
	"duryun-blog/utils/errmsg"
	"gorm.io/gorm"
)

type ArticleService struct {
}

// CreateArt 新增文章
func (s *ArticleService) CreateArt(data *model.Article) int {
	ad := dao.ArticleDao{}
	err := ad.CreateArt(data)
	if err != nil {
		return errmsg.ERROR // 500
	}
	return errmsg.SUCCESS
}

// GetCateArt 查询分类下的所有文章
func (s *ArticleService) GetCateArt(id int, size int, num int) ([]model.Article, int, int64) {
	ad := dao.ArticleDao{}
	cateArtList, err := ad.GetCateArtList(size, num, id)
	if err != nil {
		return nil, errmsg.ERROR_CATE_NOT_EXIST, 0
	}
	total, err := ad.GetCateArtNum(id)
	if err != nil {
		return nil, errmsg.ERROR_CATE_NOT_EXIST, 0
	}
	return cateArtList, errmsg.SUCCESS, total
}

// GetArtInfo 获取文章详情
func (s *ArticleService) GetArtInfo(id int) (model.Article, int) {
	ad := dao.ArticleDao{}
	art, err := ad.GetArt(id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return art, errmsg.ERROR_ART_NOT_EXIST
		}
		return art, errmsg.ERRDAO
	}
	err = ad.AddArtReadCount(id)
	if err != nil {
		return art, errmsg.ERRDAO
	}
	return art, errmsg.SUCCESS
}

// GetArtList 获取所有文章
func (s *ArticleService) GetArtList(pageSize int, pageNum int) ([]model.Article, int, int64) {
	var articleList []model.Article
	var err error
	var total int64
	ad := dao.ArticleDao{}
	articleList, err = ad.GetArtList(pageSize, pageNum)
	if err != nil {
		return nil, errmsg.ERRDAO, 0
	}
	// 单独计数
	total, err = ad.GetArtNum()
	if err != nil {
		return nil, errmsg.ERRDAO, 0
	}
	return articleList, errmsg.SUCCESS, total
}

func (s *ArticleService) SearchArticle(title string, size int, num int) ([]model.Article, int, int64) {
	ad := dao.ArticleDao{}
	articleList, total, err := ad.SearchArticle(title, size, num)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errmsg.ERROR_ART_NOT_EXIST, 0
		}
		return nil, errmsg.ERRDAO, 0
	}
	return articleList, errmsg.SUCCESS, total
}

func (s *ArticleService) EditArt(id int, data *model.Article) int {
	var maps = make(map[string]interface{})
	maps["title"] = data.Title
	maps["cid"] = data.Cid
	maps["desc"] = data.Desc
	maps["content"] = data.Content
	maps["img"] = data.Img
	ad := dao.ArticleDao{}
	err := ad.UpdateArt(id, maps)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return errmsg.ERROR_ART_NOT_EXIST
		}
		return errmsg.ERRDAO
	}
	return errmsg.SUCCESS
}

func (s *ArticleService) DeleteArt(id int) int {
	ad := dao.ArticleDao{}
	err := ad.DeleteArt(id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return errmsg.ERROR_ART_NOT_EXIST
		}
		return errmsg.ERRDAO
	}
	return errmsg.SUCCESS
}
