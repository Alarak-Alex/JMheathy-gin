package Cookie

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/Cookie"
    CookieReq "github.com/flipped-aurora/gin-vue-admin/server/model/Cookie/request"
    "gorm.io/gorm"
)

type CookieService struct {}
// CreateCookie 创建Cookie记录
// Author [piexlmax](https://github.com/piexlmax)
func (JMCookieService *CookieService) CreateCookie(JMCookie *Cookie.Cookie) (err error) {
	err = global.GVA_DB.Create(JMCookie).Error
	return err
}

// DeleteCookie 删除Cookie记录
// Author [piexlmax](https://github.com/piexlmax)
func (JMCookieService *CookieService)DeleteCookie(ID string,userID uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
	    if err := tx.Model(&Cookie.Cookie{}).Where("id = ?", ID).Update("deleted_by", userID).Error; err != nil {
              return err
        }
        if err = tx.Delete(&Cookie.Cookie{},"id = ?",ID).Error; err != nil {
              return err
        }
        return nil
	})
	return err
}

// DeleteCookieByIds 批量删除Cookie记录
// Author [piexlmax](https://github.com/piexlmax)
func (JMCookieService *CookieService)DeleteCookieByIds(IDs []string,deleted_by uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
	    if err := tx.Model(&Cookie.Cookie{}).Where("id in ?", IDs).Update("deleted_by", deleted_by).Error; err != nil {
            return err
        }
        if err := tx.Where("id in ?", IDs).Delete(&Cookie.Cookie{}).Error; err != nil {
            return err
        }
        return nil
    })
	return err
}

// UpdateCookie 更新Cookie记录
// Author [piexlmax](https://github.com/piexlmax)
func (JMCookieService *CookieService)UpdateCookie(JMCookie Cookie.Cookie) (err error) {
	err = global.GVA_DB.Model(&Cookie.Cookie{}).Where("id = ?",JMCookie.ID).Updates(&JMCookie).Error
	return err
}

// GetCookie 根据ID获取Cookie记录
// Author [piexlmax](https://github.com/piexlmax)
func (JMCookieService *CookieService)GetCookie(ID string) (JMCookie Cookie.Cookie, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&JMCookie).Error
	return
}

// GetCookieInfoList 分页获取Cookie记录
// Author [piexlmax](https://github.com/piexlmax)
func (JMCookieService *CookieService)GetCookieInfoList(info CookieReq.CookieSearch) (list []Cookie.Cookie, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&Cookie.Cookie{})
    var JMCookies []Cookie.Cookie
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.StartCreatedAt !=nil && info.EndCreatedAt !=nil {
     db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
    }
    if info.CookieType != "" {
        db = db.Where("cookie_type = ?",info.CookieType)
    }
    if info.SystemUserId != nil {
        db = db.Where("system_user_id = ?",info.SystemUserId)
    }
    if info.CookieName != "" {
        db = db.Where("cookie_name LIKE ?","%"+ info.CookieName+"%")
    }
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }
        var OrderStr string
        orderMap := make(map[string]bool)
         	orderMap["user_cookie"] = true
         	orderMap["cookie_type"] = true
         	orderMap["enable"] = true
         	orderMap["system_user_id"] = true
         	orderMap["cookie_name"] = true
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
	
	err = db.Find(&JMCookies).Error
	return  JMCookies, total, err
}
func (JMCookieService *CookieService)GetCookiePublic() {
    // 此方法为获取数据源定义的数据
    // 请自行实现
}
