/*******
* @Author:qingmeng
* @Description:
* @File:article
* @Date:2022/10/31
 */

package service

import (
	"duryun-blog/model"
	"duryun-blog/utils/errmsg"
)

type ArticleService struct {
}

func (s *ArticleService) CreateArt(data *model.Article) int {
	err := model.Db.Create(&data).Error
	if err != nil {
		return errmsg.ERROR // 500
	}
	return errmsg.SUCCESS
}
