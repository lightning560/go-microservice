package server

import (
	"fmt"
	"log"
	"net/http"
	v1 "redbook/api/interface/demo/v1"
	"redbook/app/interface/demo/internal/biz"
	"redbook/app/interface/demo/internal/data"

	"github.com/gin-gonic/gin"
)

type DemoInterface struct {
	dbiz *biz.DemoBiz
}

// 这里组装biz\repo中的组件
// newRepo NewUseCase
func NewDemoInterface() *DemoInterface {
	demoRepo := data.NewDemoRepo()
	dbiz := biz.NewDemoBiz(demoRepo)
	return &DemoInterface{dbiz: dbiz}
}

func (p *DemoInterface) CreateDemo(c *gin.Context) {
	v := new(v1.CreateDemoReq)
	//使用pb约束对外开放的端口和内容。开发随意暴露的话会出现很多安全问题。输出
	// TODO:use validation增加验证
	err := c.Bind(v)
	if err != nil {
		log.Println("CreateDemo bind err--", err)
	}
	log.Println("pre p.pc.savademo")
	rv, err := p.dbiz.CreateDemo(c, int32(v.Age), v.Name)
	fmt.Println("rv:", rv)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusNotImplemented, gin.H{})
	}
	//使用pb约束对外开放的端口和内容。开发随意暴露的话会出现很多安全问题。
	replay := &v1.CreateDemoReply{Id: int64(rv.Id)}
	c.JSON(http.StatusOK, replay)
}
