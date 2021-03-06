// Copyright (c) 2020 Xiaozhe Yao & AICAMP.CO.,LTD
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package daemon

import (
	"github.com/gin-gonic/gin"
	"os"
)

func getRouter() *gin.Engine {
	if os.Getenv("AID_PROD") == "true" {
		gin.SetMode(gin.ReleaseMode)
	}
	r := gin.Default()
	r.Use(beforeResponse())
	r.Use(gin.Recovery())
	// All get requests
	r.GET("/packages", getPackages)
	r.GET("/logs", getLogs)
	r.GET("/images", getImages)
	r.GET("/containers", getContainers)
	r.GET("/logs/:logid", getLog)
	r.GET("/solvers", getSolvers)
	r.GET("/configs", getConfigs)
	r.GET("/packages/:vendorName/:packageName/meta", getMetaInfo)
	// all put requests
	r.PUT("/packages/:vendorName/:packageName/:solverName/images", buildSolverImage)
	r.PUT("/images/:imageId/containers")
	// all post requests
	r.POST("/packages", installPackage)
	r.POST("/configs", updateConfigs)
	// all delete requests
	r.DELETE("/logs/:logid", deleteLog)
	// websocket
	r.GET("/socket/:logid", func(c *gin.Context) {
		serveWs(c)
	})
	return r
}
