package main

//import (
//	_ "embed"
//	"github.com/gin-gonic/gin"
//	"github.com/lincaiyong/daemon/common"
//	"github.com/lincaiyong/page"
//	"github.com/lincaiyong/page/com/container"
//	"github.com/lincaiyong/page/com/containeritem"
//	"github.com/lincaiyong/page/com/div"
//	"github.com/lincaiyong/page/com/root"
//	"github.com/lincaiyong/page/com/text"
//	"testing"
//)
//
////go:embed root.js
//var rootJs string
//
//func TestContainer(t *testing.T) {
//	common.StartServer("page", "v1.0.1", "",
//		func(_ []string, r *gin.RouterGroup) error {
//			baseUrl := "http://127.0.0.1:9123"
//			r.GET("/res/*filepath", page.HandleRes(baseUrl))
//			r.GET("/1", func(c *gin.Context) {
//				comp := root.Root(rootJs,
//					container.Container().List(true).Virtual(true).Scrollable(true).NameAs("containerEle").Contains(
//						containeritem.ContainerItem("Root.compute", "Root.onUpdated").Contains(
//							div.Div().OnHover("Root.onHover").Contains(
//								text.Text("''"),
//							),
//						),
//					).BackgroundColor("'#eee'").W("200").H("200").X("parent.w/2-.w/2").Y("parent.h/2-.h/2"),
//				)
//				page.MakePage(c, "debug5", comp, baseUrl, map[string]string{})
//			})
//			r.GET("/2", func(c *gin.Context) {
//				comp := container.Container().Scrollable(true).Contains(
//					text.Text("'hello world!'").H("400"),
//				).BackgroundColor("'#eee'").W("200").H("200").X("parent.w/2-.w/2").Y("parent.h/2-.h/2")
//				page.MakePage(c, "debug4", comp, baseUrl, nil)
//			})
//			return nil
//		},
//	)
//}
