package main

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/plugins/cors"
	"github.com/astaxie/beego/session"
	"github.com/xxRanger/music-dat/avatarAndDat/models"
	"github.com/xxRanger/music-dat/avatarAndDat/routers"
	_ "github.com/xxRanger/music-dat/avatarAndDat/routers"
	"os"
	"path"
)

func makeDir(dirname string) error {
	if _,err:= os.Stat(dirname); os.IsNotExist(err) {
		logs.Info("make dir",dirname)
		err:=os.MkdirAll(dirname, os.ModePerm)
		if err!=nil {
			return err
		}
	}
	return nil
}

func createDir() {
	var nftKind []string = []string{"avatar","dat","other"}
	var dirKind []string = []string{"market","public","encryption"}
	var pathPrefix = beego.AppConfig.String("fileBasePath")
	err:=makeDir(pathPrefix)
	if err!=nil {
		panic(err)
	}
	for _,nftPath:=range nftKind {
		for _,dirPath:=range dirKind {
			p:=path.Join(pathPrefix,dirPath,nftPath)
			err:=makeDir(p)
			if err!=nil {
				panic(err)
			}
		}
	}
	err = makeDir(path.Join(pathPrefix,"default"))
	if err!=nil {
		panic(err)
	}
	err=makeDir(path.Join(pathPrefix,"userIcon"))
	if err!=nil {
		panic(err)
	}
}

func main() {
	logs.SetLogFuncCallDepth(3)
	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"PUT", "PATCH", "POST", "GET"},
		AllowHeaders:     []string{"Origin", "content-type", "Access-Control-Allow-Origin"},
		ExposeHeaders:    []string{"Content-Length", "Access-Control-Allow-Origin"},
		AllowCredentials: true,
	}))
	sessionconf := &session.ManagerConfig{
		CookieName: "begoosessionID",
		Gclifetime: 3600,
	}
	beego.GlobalSessions, _ = session.NewManager("memory", sessionconf)
	go beego.GlobalSessions.GC()
	createDir()
	models.InitilizeModel(false,false)
	routers.InitRouter()
	beego.Run()
}
