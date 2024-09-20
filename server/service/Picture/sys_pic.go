package Picture

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/Picture"
	PictureReq "github.com/flipped-aurora/gin-vue-admin/server/model/Picture/request"
	"gorm.io/gorm"
)

type PictureService struct{}

// CreatePicture 创建图片记录
// Author [piexlmax](https://github.com/piexlmax)
func (PicService *PictureService) CreatePicture(Pic *Picture.Picture) (err error) {
	err = global.GVA_DB.Create(Pic).Error
	return err
}

// DeletePicture 删除图片记录
// Author [piexlmax](https://github.com/piexlmax)
func (PicService *PictureService) DeletePicture(ID string, userID uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&Picture.Picture{}).Where("id = ?", ID).Update("deleted_by", userID).Error; err != nil {
			return err
		}
		if err = tx.Delete(&Picture.Picture{}, "id = ?", ID).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

// DeletePictureByIds 批量删除图片记录
// Author [piexlmax](https://github.com/piexlmax)
func (PicService *PictureService) DeletePictureByIds(IDs []string, deleted_by uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&Picture.Picture{}).Where("id in ?", IDs).Update("deleted_by", deleted_by).Error; err != nil {
			return err
		}
		if err := tx.Where("id in ?", IDs).Delete(&Picture.Picture{}).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

// UpdatePicture 更新图片记录
// Author [piexlmax](https://github.com/piexlmax)
func (PicService *PictureService) UpdatePicture(Pic Picture.Picture) (err error) {
	err = global.GVA_DB.Model(&Picture.Picture{}).Where("id = ?", Pic.ID).Updates(&Pic).Error
	return err
}

// GetPicture 根据ID获取图片记录
// Author [piexlmax](https://github.com/piexlmax)
func (PicService *PictureService) GetPicture(ID string) (Pic Picture.Picture, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&Pic).Error
	return
}

// GetPictureInfoList 分页获取图片记录
// Author [piexlmax](https://github.com/piexlmax)
func (PicService *PictureService) GetPictureInfoList(info PictureReq.PictureSearch) (list []Picture.Picture, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&Picture.Picture{})
	var Pics []Picture.Picture
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	if info.Type != "" {
		db = db.Where("type = ?", info.Type)
	}
	if info.UpLoadUser != nil {
		db = db.Where("up_load_user = ?", info.UpLoadUser)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	var OrderStr string
	orderMap := make(map[string]bool)
	orderMap["type"] = true
	orderMap["pic"] = true
	orderMap["up_load_user"] = true
	if orderMap[info.Sort] {
		OrderStr = info.Sort
		if info.Order == "descending" {
			OrderStr = OrderStr + " desc"
		}
		db = db.Order(OrderStr)
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Find(&Pics).Error
	return Pics, total, err
}
func (PicService *PictureService) GetPicturePublic() {
	// 此方法为获取数据源定义的数据
	// 请自行实现
}
