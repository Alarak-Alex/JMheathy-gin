package Title

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/Title"
    TitleReq "github.com/flipped-aurora/gin-vue-admin/server/model/Title/request"
    "gorm.io/gorm"
)

type JMheathyTitlesService struct {}
// CreateJMheathyTitles 创建标题表记录
// Author [piexlmax](https://github.com/piexlmax)
func (JMTitleService *JMheathyTitlesService) CreateJMheathyTitles(JMTitle *Title.JMheathyTitles) (err error) {
	err = global.GVA_DB.Create(JMTitle).Error
	return err
}

// DeleteJMheathyTitles 删除标题表记录
// Author [piexlmax](https://github.com/piexlmax)
func (JMTitleService *JMheathyTitlesService)DeleteJMheathyTitles(ID string,userID uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
	    if err := tx.Model(&Title.JMheathyTitles{}).Where("id = ?", ID).Update("deleted_by", userID).Error; err != nil {
              return err
        }
        if err = tx.Delete(&Title.JMheathyTitles{},"id = ?",ID).Error; err != nil {
              return err
        }
        return nil
	})
	return err
}

// DeleteJMheathyTitlesByIds 批量删除标题表记录
// Author [piexlmax](https://github.com/piexlmax)
func (JMTitleService *JMheathyTitlesService)DeleteJMheathyTitlesByIds(IDs []string,deleted_by uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
	    if err := tx.Model(&Title.JMheathyTitles{}).Where("id in ?", IDs).Update("deleted_by", deleted_by).Error; err != nil {
            return err
        }
        if err := tx.Where("id in ?", IDs).Delete(&Title.JMheathyTitles{}).Error; err != nil {
            return err
        }
        return nil
    })
	return err
}

// UpdateJMheathyTitles 更新标题表记录
// Author [piexlmax](https://github.com/piexlmax)
func (JMTitleService *JMheathyTitlesService)UpdateJMheathyTitles(JMTitle Title.JMheathyTitles) (err error) {
	err = global.GVA_DB.Model(&Title.JMheathyTitles{}).Where("id = ?",JMTitle.ID).Updates(&JMTitle).Error
	return err
}

// GetJMheathyTitles 根据ID获取标题表记录
// Author [piexlmax](https://github.com/piexlmax)
func (JMTitleService *JMheathyTitlesService)GetJMheathyTitles(ID string) (JMTitle Title.JMheathyTitles, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&JMTitle).Error
	return
}

// GetJMheathyTitlesInfoList 分页获取标题表记录
// Author [piexlmax](https://github.com/piexlmax)
func (JMTitleService *JMheathyTitlesService)GetJMheathyTitlesInfoList(info TitleReq.JMheathyTitlesSearch) (list []Title.JMheathyTitles, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&Title.JMheathyTitles{})
    var JMTitles []Title.JMheathyTitles
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.StartCreatedAt !=nil && info.EndCreatedAt !=nil {
     db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
    }
    if info.Title != "" {
        db = db.Where("title = ?",info.Title)
    }
    if info.UploadUser != "" {
        db = db.Where("upload_user = ?",info.UploadUser)
    }
    if info.UseTime != nil {
        db = db.Where("use_time = ?",info.UseTime)
    }
    if info.IsBan != "" {
        db = db.Where("is_ban = ?",info.IsBan)
    }
    if info.BanPlatform != "" {
        db = db.Where("ban_platform = ?",info.BanPlatform)
    }
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }
        var OrderStr string
        orderMap := make(map[string]bool)
         	orderMap["title"] = true
         	orderMap["upload_user"] = true
         	orderMap["use_time"] = true
         	orderMap["is_ban"] = true
         	orderMap["ban_platform"] = true
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
	
	err = db.Find(&JMTitles).Error
	return  JMTitles, total, err
}
func (JMTitleService *JMheathyTitlesService)GetJMheathyTitlesPublic() {
    // 此方法为获取数据源定义的数据
    // 请自行实现
}
