/*******
* @Author:qingmeng
* @Description:
* @File:comment
* @Date:2022/11/1
 */

package service

import (
	"duryun-blog/dao"
	"duryun-blog/model"
	"duryun-blog/utils/errmsg"
)

type CommentService struct {
}

func (s *CommentService) AddComment(comment *model.Comment) int {
	cd := dao.CommentDao{}
	err := cd.CreateComment(comment)
	if err != nil {
		return errmsg.ERRDAO
	}
	return errmsg.SUCCESS
}

func (s *CommentService) GetComment(id int) (model.Comment, int) {
	cd := dao.CommentDao{}
	comment, err := cd.GetComment(id)
	if err != nil {
		return comment, errmsg.ERRDAO
	}
	return comment, errmsg.SUCCESS
}

func (s *CommentService) DeleteComment(id uint) int {
	cd := dao.CommentDao{}
	err := cd.DeleteComment(id)
	if err != nil {
		return errmsg.ERRDAO
	}
	return errmsg.SUCCESS
}

func (s *CommentService) GetCommentCount(id int) int64 {
	cd := dao.CommentDao{}
	total, err := cd.GetCommentCount(id)
	if err != nil {
		return 0
	}
	return total
}

func (s *CommentService) GetCommentList(pageSize int, pageNum int) ([]model.Comment, int64, int) {
	cd := dao.CommentDao{}
	commentList, total, err := cd.GetCommentList(pageSize, pageNum)
	if err != nil {
		return nil, 0, errmsg.ERRDAO
	}
	return commentList, total, errmsg.SUCCESS
}

func (s *CommentService) GetCommentListFront(id int, pageSize int, pageNum int) ([]model.Comment, int64, int) {
	cd := dao.CommentDao{}
	commentList, total, err := cd.GetCommentListFront(id, pageSize, pageNum)
	if err != nil {
		return nil, 0, errmsg.ERRDAO
	}
	return commentList, total, errmsg.SUCCESS
}

func (s *CommentService) CheckComment(id int, comment *model.Comment) int {
	cd := dao.CommentDao{}
	err := cd.CheckComment(id, comment)
	if err != nil {
		return errmsg.ERRDAO
	}
	return errmsg.SUCCESS
}

func (s *CommentService) UncheckComment(id int, comment *model.Comment) int {
	cd := dao.CommentDao{}
	err := cd.UncheckComment(id, comment)
	if err != nil {
		return errmsg.ERRDAO
	}
	return errmsg.SUCCESS
}
