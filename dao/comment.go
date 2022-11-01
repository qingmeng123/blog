/*******
* @Author:qingmeng
* @Description:
* @File:comment
* @Date:2022/11/1
 */

package dao

import (
	"duryun-blog/model"
	"gorm.io/gorm"
)

type CommentDao struct {
}

func (d *CommentDao) CreateComment(comment *model.Comment) error {
	return model.Db.Create(&comment).Error
}

func (d *CommentDao) GetComment(id int) (comment model.Comment, err error) {
	err = model.Db.Where("id = ?", id).First(&comment).Error
	return comment, err
}

func (d *CommentDao) DeleteComment(id uint) error {
	return model.Db.Where("id = ?", id).Delete(&model.Comment{}).Error
}

func (d *CommentDao) GetCommentCount(id int) (int64, error) {
	var total int64
	err := model.Db.Find(&model.Comment{}).Where("article_id = ?", id).Where("status = ?", 1).Count(&total).Error
	return total, err
}

func (d *CommentDao) GetCommentList(pageSize int, pageNum int) (commentList []model.Comment, total int64, err error) {
	err = model.Db.Model(&commentList).Limit(pageSize).Offset((pageNum - 1) * pageSize).Order("Created_At DESC").Select("comment.id, article.title,user_id,article_id, user.username, comment.content, comment.status,comment.created_at,comment.deleted_at").Joins("LEFT JOIN article ON comment.article_id = article.id").Joins("LEFT JOIN user ON comment.user_id = user.id").Scan(&commentList).Error
	if err != nil {
		return nil, 0, err
	}
	err = model.Db.Find(&commentList).Count(&total).Error
	return commentList, total, err
}

func (d *CommentDao) GetCommentListFront(id int, pageSize int, pageNum int) (commentList []model.Comment, total int64, err error) {
	err = model.Db.Model(&model.Comment{}).Limit(pageSize).Offset((pageNum-1)*pageSize).Order("Created_At DESC").
		Select("comment.id, article.title, user_id, article_id, user.username, comment.content, comment.status,comment.created_at,comment.deleted_at").Joins("LEFT JOIN article ON comment.article_id = article.id").Joins("LEFT JOIN user ON comment.user_id = user.id").Where("article_id = ?",
		id).Where("status = ?", 1).Scan(&commentList).Error
	if err != nil {
		return nil, 0, err
	}
	err = model.Db.Find(&commentList).Where("article_id = ?", id).Where("status = ?", 1).Count(&total).Error
	return commentList, total, err
}

func (d *CommentDao) CheckComment(id int, comment *model.Comment) error {
	var res model.Comment
	var article model.Article
	err := model.Db.Model(&comment).Where("id = ?", id).Updates(comment.Status).First(&res).Error
	if err != nil {
		return err
	}
	return model.Db.Model(&article).Where("id = ?", res.ArticleId).UpdateColumn("comment_count", gorm.Expr("comment_count + ?", 1)).Error
}

func (d *CommentDao) UncheckComment(id int, comment *model.Comment) error {
	var res model.Comment
	var article model.Article
	err := model.Db.Model(&comment).Where("id = ?", id).Updates(comment.Status).First(&res).Error
	if err != nil {
		return err
	}
	return model.Db.Model(&article).Where("id = ?", res.ArticleId).UpdateColumn("comment_count", gorm.Expr("comment_count - ?", 1)).Error
}
