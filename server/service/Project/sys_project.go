package Project

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/duke-git/lancet/v2/fileutil"
	"github.com/duke-git/lancet/v2/system"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	Pic "github.com/flipped-aurora/gin-vue-admin/server/model/Picture"
	"github.com/flipped-aurora/gin-vue-admin/server/model/Project"
	ProjectReq "github.com/flipped-aurora/gin-vue-admin/server/model/Project/request"
	Prompt "github.com/flipped-aurora/gin-vue-admin/server/model/Promt"
	UserUtils "github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/gin-gonic/gin"
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
func (ProjectsService *SystemProjectService) WriteWord(ID string, c *gin.Context) (err error) {

	var Projects Project.SystemProject
	err = global.GVA_DB.Model(&Project.SystemProject{}).Where("id = ?", ID).First(&Projects).Error
	if err != nil {
		return err
	}
	// PromtId为地址，需要取地址(*Projects.PromtId)
	// absPath := fileutil.CurrentPath()
	// fmt.Println(absPath)
	PromtId := *Projects.PromtId
	PIcType := Projects.PicType
	var Promp Prompt.Promt
	err = global.GVA_DB.Model(&Prompt.Promt{}).Where("id = ?", PromtId).First(&Promp).Error
	if err != nil {
		fmt.Println("没有找到提示词")
	}
	var Pictures []*Pic.Picture
	err = global.GVA_DB.Model(&Pic.Picture{}).Where("type = ?", PIcType).Find(&Pictures).Error
	if err != nil {
		fmt.Println("没有找到图片")
	}
	dir, err := os.Getwd()
	if err != nil {
		fmt.Println("获取当前工作目录失败:", err)
		return err
	}

	var PicList []string

	for _, picture := range Pictures {
		var PicPath string
		// 替换字符串
		PicPath = strings.Replace(picture.Pic, `\`, `/`, 2)
		PicPath = filepath.Join(dir, PicPath)
		if !fileutil.IsExist(PicPath) {
			fmt.Println(PicPath)
		}
		picture.Pic = PicPath
		PicList = append(PicList, PicPath)
	}
	// os.Mkdir("..\\..\\..\\demo", 0777)
	UserName := UserUtils.GetUserName(c)
	fmt.Println(UserName)
	filepath.Join(dir, UserName)
	CreatePath := "UserWord/" + UserName
	if fileutil.IsExist(CreatePath) {

		fmt.Println("文件夹已创建")

	} else {

		err := fileutil.CreateDir(CreatePath)
		if err != nil {
			fmt.Println("创建目录失败:", err)
			return err
		}

	}
	Titles, err := UserUtils.JsonArrayToStringSlice(Projects.TitleList)
	if err != nil {
		fmt.Println("转换json数组失败:", err)
		return err
	}
	// 调试文件生成代码
	var Part1 []string
	var Part2 []string
	for _, title := range Titles {
		// 写入标题
		// fmt.Println(title)
		part1 := title + "_part1.docx"
		part2 := title + "_part2.docx"
		Part2 = append(Part2, part2)
		Part1 = append(Part1, part1)
	}
	// fmt.Println(Part1)
	// fmt.Println(Part2)
	stdout, stderr, err := system.ExecCommand("pandoc --version")
	fmt.Println("std out: ", stdout)
	fmt.Println("std err: ", stderr)
	// 打印图片路径（为了调试）
	// fmt.Println(PicList)

	return err
}

// PublishArticle 发布文章
// Author [AlarakStark](https://github.com/AlarakStark)
func (ProjectsService *SystemProjectService) PublishArticle(ID string) (err error) {
	var Projects Project.SystemProject
	err = global.GVA_DB.Model(&Project.SystemProject{}).Where("id = ?", ID).First(&Projects).Error
	PromtId := *Projects.PromtId
	CookieType := Projects.CookieType
	fmt.Println(PromtId, CookieType)
	return err
}

// SyncTitle 写入标题列表
// Author [AlarakStark](https://github.com/AlarakStark)
func (ProjectsService *SystemProjectService) SyncTitle(Projects Project.SystemProject) (err error) {
	err = global.GVA_DB.Model(&Project.SystemProject{}).Where("id = ?", Projects.ID).Updates(&Projects).Error
	return err
}
