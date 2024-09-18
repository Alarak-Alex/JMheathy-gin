package initialize

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/Cookie"
	"github.com/flipped-aurora/gin-vue-admin/server/model/Picture"
	"github.com/flipped-aurora/gin-vue-admin/server/model/Project"
	"github.com/flipped-aurora/gin-vue-admin/server/model/Promt"
	"github.com/flipped-aurora/gin-vue-admin/server/model/Title"
)

func bizModel() error {
	db := global.GVA_DB
	err := db.AutoMigrate(Title.JMheathyTitles{}, Cookie.Cookie{}, Picture.Picture{}, Promt.Promt{}, Project.SystemProject{})
	if err != nil {
		return err
	}
	return nil
}
