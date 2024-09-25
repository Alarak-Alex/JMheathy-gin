package Promt

import (
	"encoding/json"
	"fmt"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/Promt"
	PromtReq "github.com/flipped-aurora/gin-vue-admin/server/model/Promt/request"
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type PromtService struct{}

// CreatePromt 创建提示词记录
// Author [piexlmax](https://github.com/piexlmax)
func (pmtService *PromtService) CreatePromt(pmt *Promt.Promt) (err error) {
	err = global.GVA_DB.Create(pmt).Error
	return err
}

// DeletePromt 删除提示词记录
// Author [piexlmax](https://github.com/piexlmax)
func (pmtService *PromtService) DeletePromt(ID string, userID uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&Promt.Promt{}).Where("id = ?", ID).Update("deleted_by", userID).Error; err != nil {
			return err
		}
		if err = tx.Delete(&Promt.Promt{}, "id = ?", ID).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

// DeletePromtByIds 批量删除提示词记录
// Author [piexlmax](https://github.com/piexlmax)
func (pmtService *PromtService) DeletePromtByIds(IDs []string, deleted_by uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&Promt.Promt{}).Where("id in ?", IDs).Update("deleted_by", deleted_by).Error; err != nil {
			return err
		}
		if err := tx.Where("id in ?", IDs).Delete(&Promt.Promt{}).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

// UpdatePromt 更新提示词记录
// Author [piexlmax](https://github.com/piexlmax)
func (pmtService *PromtService) UpdatePromt(pmt Promt.Promt) (err error) {
	err = global.GVA_DB.Model(&Promt.Promt{}).Where("id = ?", pmt.ID).Updates(&pmt).Error
	return err
}

// GetPromt 根据ID获取提示词记录
// Author [piexlmax](https://github.com/piexlmax)
func (pmtService *PromtService) GetPromt(ID string) (pmt Promt.Promt, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&pmt).Error
	return
}

// GetPromtInfoList 分页获取提示词记录
// Author [piexlmax](https://github.com/piexlmax)
func (pmtService *PromtService) GetPromtInfoList(info PromtReq.PromtSearch) (list []Promt.Promt, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&Promt.Promt{})
	var pmts []Promt.Promt
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	if info.SystemUserId != nil {
		db = db.Where("system_user_id = ?", info.SystemUserId)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	var OrderStr string
	orderMap := make(map[string]bool)
	orderMap["promt_data"] = true
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

	err = db.Find(&pmts).Error
	if err != nil {
		return
	}

	return pmts, total, err
}
func (pmtService *PromtService) GetPromtPublic() {
	// 此方法为获取数据源定义的数据
	// 请自行实现
}

func ParseJson(pd datatypes.JSON, result []string) []string {
	// 接受json数据返回字符串切片
	err := json.Unmarshal([]byte(pd), &result)
	if err != nil {
		fmt.Println("json解析失败", err)
	}
	return result
}
