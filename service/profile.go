/*******
* @Author:qingmeng
* @Description:
* @File:profile
* @Date:2022/11/1
 */

package service

import (
	"duryun-blog/dao"
	"duryun-blog/model"
	"duryun-blog/utils/errmsg"
)

type ProfileService struct {
}

func (s *ProfileService) GetProfile(id int) (model.Profile, int) {
	pd := dao.ProfileDao{}
	profile, err := pd.GetProfile(id)
	if err != nil {
		return model.Profile{}, errmsg.ERRDAO
	}
	return profile, errmsg.SUCCESS
}

func (s *ProfileService) UpdateProfile(id int, profile *model.Profile) int {
	pd := dao.ProfileDao{}
	err := pd.UpdateProfile(id, profile)
	if err != nil {
		return errmsg.ERRDAO
	}
	return errmsg.SUCCESS
}
