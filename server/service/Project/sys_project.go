package Project

import (
	"fmt"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/Project"
	ProjectReq "github.com/flipped-aurora/gin-vue-admin/server/model/Project/request"
	"gorm.io/gorm"
)

type SystemProjectService struct{}

// CreateSystemProject 创建项目记录
// Author [piexlmax](https://github.com/piexlmax)
func (ProjectsService *SystemProjectService) CreateSystemProject(Projects *Project.SystemProject) (err error) {
	err = global.GVA_DB.Create(Projects).Error
	fmt.Println("创建项目记录成功", Projects.TitleList)
	return err
}

// DeleteSystemProject 删除项目记录
// Author [piexlmax](https://github.com/piexlmax)
func (ProjectsService *SystemProjectService) DeleteSystemProject(ID string, userID uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&Project.SystemProject{}).Where("id = ?", ID).Update("deleted_by", userID).Error; err != nil {
			return err
		}
		if err = tx.Delete(&Project.SystemProject{}, "id = ?", ID).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

// DeleteSystemProjectByIds 批量删除项目记录
// Author [piexlmax](https://github.com/piexlmax)
func (ProjectsService *SystemProjectService) DeleteSystemProjectByIds(IDs []string, deleted_by uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&Project.SystemProject{}).Where("id in ?", IDs).Update("deleted_by", deleted_by).Error; err != nil {
			return err
		}
		if err := tx.Where("id in ?", IDs).Delete(&Project.SystemProject{}).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

// UpdateSystemProject 更新项目记录
// Author [piexlmax](https://github.com/piexlmax)

func (ProjectsService *SystemProjectService) UpdateSystemProject(Projects Project.SystemProject) (err error) {
	err = global.GVA_DB.Model(&Project.SystemProject{}).Where("id = ?", Projects.ID).Updates(&Projects).Error
	return err
}

// GetSystemProject 根据ID获取项目记录
// Author [piexlmax](https://github.com/piexlmax)
func (ProjectsService *SystemProjectService) GetSystemProject(ID string) (Projects Project.SystemProject, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&Projects).Error
	return
}

// GetSystemProjectInfoList 分页获取项目记录
// Author [piexlmax](https://github.com/piexlmax)
func (ProjectsService *SystemProjectService) GetSystemProjectInfoList(info ProjectReq.SystemProjectSearch) (list []Project.SystemProject, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&Project.SystemProject{})
	var Projectss []Project.SystemProject
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	if info.PicType != "" {
		db = db.Where("pic_type = ?", info.PicType)
	}
	if info.PromtId != nil {
		db = db.Where("promt_id = ?", info.PromtId)
	}
	if info.CookieType != "" {
		db = db.Where("cookie_type = ?", info.CookieType)
	}
	if info.SystemUserId != nil {
		db = db.Where("system_user_id = ?", info.SystemUserId)
	}
	if info.Status != "" {
		db = db.Where("status = ?", info.Status)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	var OrderStr string
	orderMap := make(map[string]bool)
	orderMap["title_list"] = true
	orderMap["pic_type"] = true
	orderMap["promt_id"] = true
	orderMap["cookie_type"] = true
	orderMap["system_user_id"] = true
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

	err = db.Find(&Projectss).Error
	return Projectss, total, err
}
func (ProjectsService *SystemProjectService) GetSystemProjectPublic() {
	// 此方法为获取数据源定义的数据
	// 请自行实现
}

// WriteWord 写文
// Author [AlarakStark](https://github.com/AlarakStark)
func (ProjectsService *SystemProjectService) WriteWord() (err error) {
	// 请在这里实现自己的业务逻辑
	db := global.GVA_DB.Model(&Project.SystemProject{})
	return db.Error
}

// PublishArticle 发布文章
// Author [AlarakStark](https://github.com/AlarakStark)
func (ProjectsService *SystemProjectService) PublishArticle() (err error) {
	// 请在这里实现自己的业务逻辑
	db := global.GVA_DB.Model(&Project.SystemProject{})
	return db.Error
}

// // SyncTitle 写入标题列表
// // Author [yourname](https://github.com/yourname)
func (ProjectsService *SystemProjectService) SyncTitle(Projects Project.SystemProject) (err error) {
	// 请在这里实现自己的业务逻辑
	err = global.GVA_DB.Model(&Project.SystemProject{}).Where("id = ?", Projects.ID).Updates(&Projects).Error
	return err
}
