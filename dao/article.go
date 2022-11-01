/*******
* @Author:qingmeng
* @Description:
* @File:article
* @Date:2022/10/31
 */

package dao

import (
	"duryun-blog/model"
	"gorm.io/gorm"
)

type ArticleDao struct {
}

// CreateArt 新增文章
func (d *ArticleDao) CreateArt(data *model.Article) error {
	return model.Db.Create(&data).Error
}

// GetCateArtList 查询分类下的所有文章
func (d *ArticleDao) GetCateArtList(size int, num int, id int) (cateArtList []model.Article, err error) {
	err = model.Db.Preload("Category").Limit(size).Offset((num-1)*size).Where(
		"cid =?", id).Find(&cateArtList).Error
	return cateArtList, err
}

// GetCateArtNum 查询分类的文章数量
func (d *ArticleDao) GetCateArtNum(id int) (total int64, err error) {
	err = model.Db.Model(&model.Article{}).Where("cid =?", id).Count(&total).Error
	return total, err
}

func (d *ArticleDao) GetArt(id int) (art model.Article, err error) {
	err = model.Db.Where("id = ?", id).Preload("Category").First(&art).Error
	return art, err
}

//增加文章阅读数
func (d *ArticleDao) AddArtReadCount(id int) error {
	art := model.Article{}
	return model.Db.Model(&art).Where("id = ?", id).UpdateColumn("read_count", gorm.Expr("read_count + ?", 1)).Error
}

func (d *ArticleDao) GetArtList(pageSize int, pageNum int) (articleList []model.Article, err error) {
	err = model.Db.Select("article.id, title, img, created_at, updated_at, `desc`, comment_count, read_count, Category.name").Limit(pageSize).Offset((pageNum - 1) * pageSize).Order("Created_At DESC").Joins("Category").Find(&articleList).Error
	return articleList, err
}

func (d *ArticleDao) GetArtNum() (total int64, err error) {
	err = model.Db.Model(&model.Article{}).Count(&total).Error
	return total, err
}

func (d *ArticleDao) SearchArticle(title string, pageSize int, pageNum int) ([]model.Article, int64, error) {
	var articleList []model.Article
	var err error
	var total int64
	err = model.Db.Select("article.id,title, img, created_at, updated_at, `desc`, comment_count, read_count, Category.name").Order("Created_At DESC").Joins("Category").Where("title LIKE ?",
		title+"%",
	).Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&articleList).Error
	if err != nil {
		return nil, 0, err
	}

	//单独计数
	err = model.Db.Model(&articleList).Where("title LIKE ?", title+"%").Count(&total).Error
	if err != nil {
		return nil, 0, err
	}
	return articleList, total, err
}

func (d *ArticleDao) UpdateArt(id int, maps map[string]interface{}) error {
	return model.Db.Model(&model.Article{}).Where("id = ? ", id).Updates(&maps).Error
}

func (d *ArticleDao) DeleteArt(id int) error {
	return model.Db.Where("id = ? ", id).Delete(&model.Article{}).Error
}
