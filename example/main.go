package main

import (
	_ "embed"
	"github.com/gin-gonic/gin"
	"github.com/lincaiyong/daemon/common"
	"github.com/lincaiyong/page"
	"github.com/lincaiyong/page/com"
	"net/http"
)

const baseUrl = "http://127.0.0.1:9123"

func main() {
	common.StartServer(
		"page",
		"v1.0.1",
		"",
		func(envs []string, router *gin.RouterGroup) error {
			//router.Use(corsMiddleware())
			//router.Use(cacheMiddleware(""))
			router.GET("/res/*filepath", page.HandleRes(baseUrl))
			router.GET("/debug", debugPage)
			router.GET("/debug2", debug2Page)
			router.GET("/debug3", debug3Page)
			router.GET("/debug4", debug4Page)
			router.GET("/debug5", debug5Page)
			router.GET("/code/get", getCode)
			router.GET("/code/get2", getCode2)
			return nil
		},
	)
}

type ExplorerPaneComponent struct {
	*com.BaseComponent
	handleClickItem com.Method
	onCreated       com.Method
}

func ExplorerPane() *ExplorerPaneComponent {
	ret := &ExplorerPaneComponent{}
	ret.BaseComponent = com.NewBaseComponent("div", ret,
		com.Tree(), // onClickItem=handleClickItem
	)
	return ret
}

func debug3Page(c *gin.Context) {
	comp := com.Div().Contains(
		com.Div().H_(34).BackgroundColor("'black'"),
		com.Div().Y("prev.y2").H("parent.h - prev.h - next.h").Contains(
			ExplorerPane().W("next.x").BackgroundColor("page.theme.grayPaneColor"),
			com.VBar().X("prev.v ? 200 : - .w").BackgroundColor("'red'").Opacity("0.1"),
			com.Editor().X("prev.x2").W("parent.w - prev.x2"),
		),
		com.Div().Y("prev.y2").H_(24).BorderColor("page.theme.grayBorderColor").BorderTop("1").BackgroundColor("page.theme.grayPaneColor"),
	)
	page.MakePage(c, "debug3", comp, baseUrl, map[string]string{
		"ExplorerPaneComponent_handleClickItem.js": `function handleClickItem(ele, ev) {
	if (ele.data.leaf) {
		console.log(ele.data.key);
		const path = ele.data.key.split('/').join('%2f');
		const queryString = window.location.search;
		const url = "<base_url>/code/get" + queryString + "&path=" + path;
		page.util.fetch(url).then(v => {
			this.root.editorEle.setValue(v);
		}).catch(e => page.log.error(e));
	}
}`,
		"ExplorerPaneComponent_onCreated.js": `function onCreated() {
	const queryString = window.location.search;
	page.util.fetch('<base_url>/code/get2' + queryString).then(v => {
		const resp = JSON.parse(v);
		console.log(resp);
		this.treeEle.items = Object.keys(resp);
	}).catch(e => page.log.error(e));
}`,
	})
}

func debug2Page(c *gin.Context) {
	comp := com.Text("'debug2'").H_(200)
	page.MakePage(c, "debug2", comp, baseUrl, nil)
}

func debugPage(c *gin.Context) {
	comp := com.Div().Contains(
		com.Editor().X_(20).Y_(20).W_(800).H("next.y - .y").BackgroundColor("'blue'"),
		com.HBar().Y_(200).W("parent.w").H_(20).Opacity("0.1"),
		com.Div().X_(20).Y("prev.y2").W_(800).H_(400).BackgroundColor("'yellow'").Contains(
			com.Text("'hello world!'").H_(200),
		),
		com.Compare().X_(800).Y_(20).W_(400).H_(400).BackgroundColor("'red'"),
		com.Div().X("prev.x2").Y("prev.y2").W_(40).H_(40).BackgroundColor("'green'"),
		com.Button().Icon("'svg/el/folder.svg'").X("prev.x2").Y("prev.y2 + 100").W_(40).H_(40),
	)
	page.MakePage(c, "debug", comp, baseUrl, nil)
}

//go:embed main.go
var mainGo []byte

func getCode(c *gin.Context) {
	c.String(http.StatusOK, string(mainGo))
}

func getCode2(c *gin.Context) {
	c.String(http.StatusOK, `{
	"test.txt": "hello world!"
}`)
}

func debug4Page(c *gin.Context) {
	comp := com.Container().Scrollable(true).Contains(
		com.Text("'hello world!'").H("400"),
	).BackgroundColor("'#eee'").W("200").H("200").X("parent.w/2-.w/2").Y("parent.h/2-.h/2")
	page.MakePage(c, "debug4", comp, baseUrl, nil)
}

func DemoContainerItem() *DemoContainerItemComponent {
	ret := &DemoContainerItemComponent{}
	ret.BaseComponent = com.NewBaseComponent("div", ret,
		com.Text("''").OnHover(`(ele, hovered) => {
	ele.backgroundColor = hovered ? '#888' : '#eee'; 
}`),
	)
	return ret
}

type DemoContainerItemComponent struct {
	*com.BaseComponent
	data      com.Property `type:"object"`
	compute   com.StaticMethod
	onUpdated com.Method `method:"onUpdated"`
}

func (b *DemoContainerItemComponent) Data(s string) *DemoContainerItemComponent {
	b.Props()["data"] = s
	return b
}

func (b *DemoContainerItemComponent) Compute(s string) *DemoContainerItemComponent {
	b.Props()["compute"] = s
	return b
}

func debug5Page(c *gin.Context) {
	comp := com.Container().List(true).Virtual(true).Scrollable(true).Contains(
		DemoContainerItem(),
	).BackgroundColor("'#eee'").W("200").H("200").X("parent.w/2-.w/2").Y("parent.h/2-.h/2")
	page.MakePage(c, "debug5", comp, baseUrl, map[string]string{
		"DemoContainerItemComponent_compute.js": `function compute(container, idx, prev) {
	return {
		key: ''+idx,
		x: 0,
		y: 20 * idx,
		w: 200,
		h: 20,
		text: 'hello world!' + idx,
	}
}`,
		"DemoContainerItemComponent_onUpdated.js": `function onUpdated(k, v) {
	if (k === 'data') {
		this.children[0].text = v.text;
	}
}`,
		"": `setTimeout(function() {
	const container = page.root.children[0];
	container.items = [1,2,3,4,5,6,7,8,9,10,11,12,13,14,15,16];
}, 1000);`,
	})
}
