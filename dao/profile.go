/*******
* @Author:qingmeng
* @Description:
* @File:profile
* @Date:2022/11/1
 */

package dao

import "duryun-blog/model"

type ProfileDao struct {
}

func (d ProfileDao) GetProfile(id int) (profile model.Profile, err error) {
	err = model.Db.Where("ID = ?", id).First(&profile).Error
	return profile, err
}

func (d ProfileDao) UpdateProfile(id int, profile *model.Profile) error {
	return model.Db.Model(&profile).Where("ID = ?", id).Updates(&profile).Error
}
