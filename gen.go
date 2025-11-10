package page

import (
	"embed"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/lincaiyong/log"
	"github.com/lincaiyong/page/com"
	"github.com/lincaiyong/page/js"
	"github.com/lincaiyong/page/parser"
	"github.com/lincaiyong/page/printer"
	"github.com/lincaiyong/page/visit"
	"net/http"
	"reflect"
	"sort"
	"strings"
)

//go:embed js/*
var jsEmbed embed.FS

func sortedKeys(m map[string]string) []string {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	return keys
}

func getAllInstances(comp com.Component, result []com.Component) []com.Component {
	result = append(result, comp)
	for _, tmp := range comp.Slots() {
		result = getAllInstances(tmp, result)
	}
	for _, tmp := range comp.Children() {
		result = getAllInstances(tmp, result)
	}
	return result
}

type StructInfo struct {
	name             string
	properties       []string
	staticProperties []string
	methods          []string
	staticMethods    []string
	bindJs           map[string]string
	defaultValue     map[string]string
}

func (info *StructInfo) readFunction(name, source string) string {
	var code string
	if strings.HasPrefix(source, "function ") {
		code = source
	} else {
		b, err := jsEmbed.ReadFile(fmt.Sprintf("js/%s", source))
		if err != nil {
			log.FatalLog("xx")
		}
		code = strings.TrimSpace(string(b))
	}
	if !strings.HasPrefix(code, fmt.Sprintf("function %s(", name)) {
		log.FatalLog("xx")
	}
	return code[9:]
}

func (info *StructInfo) GenClass(pr *printer.Printer) {
	pr.Put("class %s extends Component {", info.name).Push()
	{
		pr.Put("constructor(parent, model) {").Push()
		{
			if len(info.properties) == 0 {
				pr.Put("model.properties = Object.assign({}, model.properties);")
			} else {
				pr.Put("model.properties = Object.assign({").Push()
				for _, property := range info.properties {
					pr.Put("%s: [e => %s, []],", property, info.defaultValue[property])
				}
				pr.Pop().Put("}, model.properties);")
			}
			pr.Put("super(parent, model);")
		}
		pr.Pop().Put("}")
		for _, method := range info.methods {
			code := info.readFunction(method, info.bindJs[method])
			pr.Put(code)
		}
		for _, method := range info.staticMethods {
			code := info.readFunction(method, info.bindJs[method])
			pr.Put("static " + code)
		}
		for _, property := range info.properties {
			pr.Put("get %s() { return this._properties.%s.value; }", property, property)
			pr.Put("set %s(v) { this._properties.%s.value = v; }", property, property)
		}
	}
	pr.Pop().Put("}")
}

func defaultValue(type_, default_ string) string {
	if default_ != "" {
		return default_
	}
	switch type_ {
	case "bool":
		return "false"
	case "number":
		return "0"
	case "string":
		return "''"
	case "array", "element":
		return "undefined"
	case "object":
		return "{}"
	default:
		log.FatalLog("xx")
		return ""
	}
}

func buildClasses(page com.Component, mm map[string]string) string {
	result := getAllInstances(page, nil)
	result2 := map[string]reflect.Type{}
	for _, tmp := range result {
		t := reflect.ValueOf(tmp).Type().Elem()
		result2[t.Name()] = t
	}
	result3 := map[string]StructInfo{}
	for n, t := range result2 {
		info := StructInfo{
			name:         n,
			bindJs:       map[string]string{},
			defaultValue: map[string]string{},
		}
		for i := 0; i < t.NumField(); i++ {
			field := t.Field(i)
			if !field.Anonymous {
				tn := field.Type.Name()
				switch tn {
				case "Property":
					info.properties = append(info.properties, field.Name)
					info.defaultValue[field.Name] = defaultValue(field.Tag.Get("type"), field.Tag.Get("default"))
				case "StaticProperty":
					info.staticProperties = append(info.staticProperties, field.Name)
					info.defaultValue[field.Name] = defaultValue(field.Tag.Get("type"), field.Tag.Get("default"))
				case "Method":
					info.methods = append(info.methods, field.Name)
					code := field.Tag.Get("bind")
					if code == "" {
						code = mm[fmt.Sprintf("%s_%s.js", n, field.Name)]
						if code == "" {
							code = js.Get(n, field.Name)
							if code == "" {
								log.FatalLog("xx")
							}
						}
					}
					info.bindJs[field.Name] = code
				case "StaticMethod":
					info.staticMethods = append(info.staticMethods, field.Name)
					code := field.Tag.Get("bind")
					if code == "" {
						code = mm[fmt.Sprintf("%s_%s.js", n, field.Name)]
						if code == "" {
							code = js.Get(n, field.Name)
							if code == "" {
								log.FatalLog("xx")
							}
						}
					}
					info.bindJs[field.Name] = code
				default:
					log.FatalLog("xx")
				}
			}
		}
		result3[n] = info
	}
	keys := make([]string, 0, len(result3))
	for k := range result3 {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	pr := printer.NewPrinter()
	for _, k := range keys {
		info := result3[k]
		info.GenClass(pr)
	}
	return pr.Code()
}

func convertExpr(s string) (string, string) {
	tokens, err := parser.Tokenize(s)
	if err != nil {
		log.WarnLog("invalid expr: %s, %v", s, err)
		return s, "[]"
	}
	node, err := parser.Parse(tokens)
	if err != nil {
		log.FatalLog("invalid expr: %s, %v", s, err)
	}
	s1, s2, err := visit.Visit(node)
	if err != nil {
		log.FatalLog("invalid expr: %s, %v", s, err)
	}
	return s1, s2
}

func buildModel(page com.Component, depth, modelDepth int, pr *printer.Printer) {
	t := reflect.ValueOf(page).Type().Elem()
	pr.Put("{").Push()
	{
		pr.Put("Component: %s,", t.Name())
		pr.Put("tag: '%s',", page.Tag())
		pr.Put("overflow: 'hidden',")
		pr.Put("name: '%s',", page.Name())
		pr.Put("depth: %d,", depth)
		props := make(map[string]string)
		if modelDepth > 0 {
			props["h"] = "parent.ch"
			props["v"] = "parent.v"
			props["w"] = "parent.cw"
			props["x"] = "-parent.scrollLeft"
			props["y"] = "-parent.scrollTop"
			props["zIndex"] = "parent.zIndex"
		}
		for k, v := range page.Props() {
			props[k] = v
		}
		if len(props) == 0 {
			pr.Put("properties: {},")
		} else {
			pr.Put("properties: {").Push()
			for _, k := range sortedKeys(props) {
				v1, v2 := convertExpr(props[k])
				pr.Put("%s: [e => %s, %s],", k, v1, v2)
			}
			pr.Pop().Put("},")
		}
		if len(page.StaticProps()) == 0 {
			pr.Put("staticProperties: {},")
		} else {
			pr.Put("staticProperties: {").Push()
			for _, k := range sortedKeys(page.StaticProps()) {
				pr.Put("%s: [e => %s, []],", k, page.StaticProps()[k])
			}
			pr.Pop().Put("},")
		}
		if len(page.Children()) == 0 {
			pr.Put("children: [],")
		} else {
			pr.Put("children: [").Push()
			for _, tmp := range page.Children() {
				buildModel(tmp, depth+1, modelDepth+1, pr)
			}
			pr.Pop().Put("],")
		}
		if len(page.Slots()) == 0 {
			pr.Put("slot: null,")
		} else {
			pr.Put("slot: [").Push()
			for _, tmp := range page.Slots() {
				buildModel(tmp, depth+1, modelDepth+1, pr)
			}
			pr.Pop().Put("],")
		}
	}
	pr.Pop().Put("},")
}

func buildPageModel(page com.Component) string {
	pr := printer.NewPrinter()
	buildModel(page, 0, 0, pr)
	code := pr.Code()
	code = strings.TrimRight(code, ",")
	return "page.model = " + code + ";"
}

func MakePage(c *gin.Context, title string, page com.Component, baseUrl string, mm map[string]string) {
	if mm == nil {
		mm = make(map[string]string)
	}
	page = com.Div().Contains(page)
	eventJs, _ := jsEmbed.ReadFile("js/_event.js")
	propertyJs, _ := jsEmbed.ReadFile("js/_property.js")
	scrollbarJs, _ := jsEmbed.ReadFile("js/_scrollbar.js")
	componentJs, _ := jsEmbed.ReadFile("js/_component.js")
	pageJs, _ := jsEmbed.ReadFile("js/_page.js")
	classes := buildClasses(page, mm)
	model := buildPageModel(page)
	s := []string{string(eventJs), string(propertyJs), string(scrollbarJs), string(componentJs), string(pageJs), classes, model}
	ss := strings.Join(strings.Split(strings.Join(s, "\n"), "\n"), "\n        ")
	template := `<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="utf-8"/>
    <title><ttt></title>
    <link rel="stylesheet" href="<base_url>/res/vs/editor/editor.main.css" />
</head>
<body>
    <script src="<base_url>/res/vs/loader.js"></script>
    <script>
        <xxx>
        require.config({paths: {'vs': '<base_url>/res/vs'}});
        require(['vs/editor/editor.main'], () =>page.create());
    </script>
</body>
</html>`
	if !strings.Contains(ss, "monaco") {
		template = `<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="utf-8"/>
    <title><ttt></title>
</head>
<body>
    <script>
        <xxx>
        page.create();
		<xxx2>
    </script>
</body>
</html>`
	}
	template = strings.ReplaceAll(template, "<ttt>", title)
	ss = strings.ReplaceAll(template, "<xxx>", ss)
	ss = strings.ReplaceAll(ss, "<xxx2>", mm[""])
	html := strings.ReplaceAll(ss, "<base_url>", baseUrl)
	c.Data(http.StatusOK, "text/html; charset=utf-8", []byte(html))
}
