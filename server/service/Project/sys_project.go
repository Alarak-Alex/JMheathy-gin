package Project

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
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
func (ProjectsService *SystemProjectService) WriteWord(ID string, c *gin.Context) error {
	var project Project.SystemProject
	// 从数据库获取项目
	if err := global.GVA_DB.Model(&Project.SystemProject{}).Where("id = ?", ID).First(&project).Error; err != nil {
		return err
	}

	promptID := *project.PromtId
	picType := project.PicType

	// 从数据库获取提示词
	var prompt Prompt.Promt
	if err := global.GVA_DB.Model(&Prompt.Promt{}).Where("id = ?", promptID).First(&prompt).Error; err != nil {
		fmt.Println("没有找到提示词")
	}

	// 从数据库获取图片
	var pictures []*Pic.Picture
	if err := global.GVA_DB.Model(&Pic.Picture{}).Where("type = ?", picType).Find(&pictures).Error; err != nil {
		fmt.Println("没有找到图片")
	}

	// 获取当前工作目录
	dir, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("获取当前工作目录失败: %w", err)
	}

	var picList []string
	for _, picture := range pictures {
		picPath := filepath.Join(dir, strings.Replace(picture.Pic, `\`, `/`, 2))
		if !fileutil.IsExist(picPath) {
			fmt.Println(picPath)
		}
		picture.Pic = picPath
		picList = append(picList, picPath)
	}

	// 获取用户名并创建目录
	userName := UserUtils.GetUserName(c)
	createPath := filepath.Join("UserWord", userName)
	if !fileutil.IsExist(createPath) {
		if err := fileutil.CreateDir(createPath); err != nil {
			return fmt.Errorf("创建目录失败: %w", err)
		}
		fmt.Println("文件夹已创建")
	}

	// 转换标题列表
	fmt.Println(project.TitleList)
	titles, err := UserUtils.JsonArrayToStringSlice(project.TitleList)
	if err != nil {
		return fmt.Errorf("转换json数组失败: %w", err)
	}
	// 初始化通道
	Titlechan := make(chan string)

	// 发送标题到通道
	go func() {
		for _, title := range titles {
			Titlechan <- title
		}
		close(Titlechan)
	}()

	// CHAN 通道
	fmt.Println("通道已创建")

	// 根据系统的 CPU 核心数开启多个协程
	numGoroutines := runtime.NumCPU()
	maxGoroutines := len(titles) // 可以用 titles 长度作为最大并发数

	for i := 0; i < numGoroutines && i < maxGoroutines; i++ {
		fmt.Println("开启协程", i)
		go UserUtils.Chatmain(prompt, Titlechan, createPath)
	}

	// 生成文档名称
	var part1, part2 []string
	for _, title := range titles {
		part1 = append(part1, title+"_part1.docx")
		part2 = append(part2, title+"_part2.docx")
	}

	// 打印Pandoc版本信息
	stdout, stderr, err := system.ExecCommand("pandoc --version")
	if err != nil {
		return fmt.Errorf("执行命令失败: %w", err)
	}
	fmt.Println("std out: ", stdout)
	fmt.Println("std err: ", stderr)

	return nil
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
