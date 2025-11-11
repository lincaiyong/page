package page

import (
	"embed"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/lincaiyong/log"
	"github.com/lincaiyong/page/com"
	"github.com/lincaiyong/page/com/root"
	"github.com/lincaiyong/page/js"
	"github.com/lincaiyong/page/parser"
	"github.com/lincaiyong/page/printer"
	"github.com/lincaiyong/page/utils"
	"github.com/lincaiyong/page/visit"
	"net/http"
	"reflect"
	"sort"
	"strconv"
	"strings"
)

func sortedKeys(m map[string]string) []string {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	return keys
}

func getAllInstances(comp com.Component, comps []com.Component, thisComp com.Component, selfIndex []int) []com.Component {
	struct_ := reflect.TypeOf(comp).Elem()
	name := struct_.Name()
	if name == "Component" {
		name = struct_.String()
		name = name[:strings.Index(name, ".")]
		name = utils.PascalCase(name)
	}
	info := comp.ExtraInfo()
	info.SetName(name)
	if thisComp == nil {
		thisComp = comp
	}
	info.SetThisComponent(thisComp)
	info.SetSelfIndex(selfIndex)
	comps = append(comps, comp)
	for i, tmp := range comp.Slots() {
		comps = getAllInstances(tmp, comps, thisComp, append(selfIndex, i))
	}
	for i, tmp := range comp.Children() {
		comps = getAllInstances(tmp, comps, comp, []int{i})
	}
	return comps
}

func readFunction(name, code string) string {
	if !strings.HasPrefix(code, fmt.Sprintf("function %s(", name)) {
		log.FatalLog("invalid function code: %s", code)
	}
	return code[9:]
}

func genClassCode(info *com.ExtraInfo, namedChildren map[string]map[string][]int, pr *printer.Printer) {
	pr.Put("class %s extends Component {", info.Name()).Push()
	{
		pr.Put("constructor(parent, model) {").Push()
		{
			if len(info.Properties()) == 0 {
				pr.Put("model.properties = Object.assign({}, model.properties);")
			} else {
				pr.Put("model.properties = Object.assign({").Push()
				for _, property := range info.Properties() {
					pr.Put("%s: [e => %s, []],", property, info.DefaultValue()[property])
				}
				pr.Pop().Put("}, model.properties);")
			}
			pr.Put("super(parent, model);")
		}
		pr.Pop().Put("}")
		for _, method := range info.Methods() {
			code := readFunction(method, info.BindJs()[method])
			pr.Put(code)
		}
		for _, method := range info.StaticMethods() {
			code := readFunction(method, info.BindJs()[method])
			pr.Put("static " + code)
		}
		for _, property := range info.Properties() {
			pr.Put("get %s() { return this._properties.%s.value; }", property, property)
			pr.Put("set %s(v) { this._properties.%s.value = v; }", property, property)
		}
		if m, ok := namedChildren[info.Name()]; ok {
			keys := make([]string, 0, len(m))
			for k := range m {
				keys = append(keys, k)
			}
			sort.Strings(keys)
			for _, k := range keys {
				items := make([]string, 0, len(m[k]))
				for _, i := range m[k] {
					items = append(items, strconv.Itoa(i))
				}
				pr.Put("get %s() { return [%s].reduce((prev, i) => prev.children[i], this); }", k, strings.Join(items, ", "))
			}
		}
	}
	if info.Name() == "Root" {
		m := js.GetAll(info.Name())
		keys := make([]string, 0, len(m))
		for k := range m {
			keys = append(keys, k)
		}
		sort.Strings(keys)
		for _, k := range keys {
			if strings.HasPrefix(m[k], "function ") {
				code := m[k][9:]
				pr.Put("static " + code)
			}
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
		log.FatalLog("invalid property type: %s", type_)
		return ""
	}
}

func buildClasses(page com.Component, mm map[string]string) string {
	comps := getAllInstances(page, nil, nil, nil)
	namedChildren := make(map[string]map[string][]int)
	for _, comp := range comps {
		if strings.HasSuffix(comp.Name(), "Ele") {
			thisComp := comp.ExtraInfo().ThisComponent()
			if thisComp != comp {
				thisCompName := thisComp.ExtraInfo().Name()
				if m, ok := namedChildren[thisCompName]; ok {
					m[comp.Name()] = comp.ExtraInfo().SelfIndex()
				} else {
					m = make(map[string][]int)
					m[comp.Name()] = comp.ExtraInfo().SelfIndex()
					namedChildren[thisCompName] = m
				}
			}
		}
	}
	compMap := make(map[string]com.Component)
	for _, comp := range comps {
		name := comp.ExtraInfo().Name()
		if _, ok := compMap[name]; !ok {
			compMap[name] = comp
		}
	}
	keys := make([]string, 0, len(compMap))
	for n, comp := range compMap {
		keys = append(keys, n)
		info := comp.ExtraInfo()
		info.SetBindJs(map[string]string{})
		info.SetDefaultValue(map[string]string{})
		struct_ := reflect.TypeOf(comp).Elem()
		for i := 0; i < struct_.NumField(); i++ {
			field := struct_.Field(i)
			if !field.Anonymous {
				tn := field.Type.Name()
				switch tn {
				case "Property":
					info.SetProperties(append(info.Properties(), field.Name))
					info.DefaultValue()[field.Name] = defaultValue(field.Tag.Get("type"), field.Tag.Get("default"))
				case "Method":
					info.SetMethods(append(info.Methods(), field.Name))
					code := mm[fmt.Sprintf("%s_%s", n, field.Name)]
					if code == "" {
						code = js.Get(n, field.Name)
						if code == "" {
							log.FatalLog("fail to get js code: %s %s", n, field.Name)
						}
					}
					info.BindJs()[field.Name] = code
				case "StaticMethod":
					info.SetStaticMethods(append(info.StaticMethods(), field.Name))
					code := mm[fmt.Sprintf("%s_%s", n, field.Name)]
					if code == "" {
						code = js.Get(n, field.Name)
						if code == "" {
							log.FatalLog("fail to get js code: %s %s", n, field.Name)
						}
					}
					info.BindJs()[field.Name] = code
				default:
					log.FatalLog("invalid field type: %s", tn)
				}
			}
		}
	}
	sort.Strings(keys)
	pr := printer.NewPrinter()
	for _, k := range keys {
		genClassCode(compMap[k].ExtraInfo(), namedChildren, pr)
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

func buildModel(comp com.Component, depth, modelDepth int, pr *printer.Printer) {
	t := reflect.ValueOf(comp).Type().Elem()
	pr.Put("{").Push()
	{
		s := t.Name()
		if s == "Component" {
			s = t.String()
			s = s[:strings.Index(s, ".")]
			s = utils.PascalCase(s)
		}
		pr.Put("Component: %s,", s)
		pr.Put("tag: '%s',", comp.Tag())
		pr.Put("overflow: 'hidden',")
		pr.Put("name: '%s',", comp.Name())
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
		for k, v := range comp.Props() {
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
		children := comp.Children()
		slots := comp.Slots()
		if comp.SlotsAsChildren() {
			children = slots
			slots = nil
		}
		if len(children) == 0 {
			pr.Put("children: [],")
		} else {
			pr.Put("children: [").Push()
			for _, tmp := range children {
				buildModel(tmp, depth+1, modelDepth+1, pr)
			}
			pr.Pop().Put("],")
		}
		if len(slots) == 0 {
			pr.Put("slot: null,")
		} else {
			pr.Put("slot: [").Push()
			for _, tmp := range slots {
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

//go:embed js/*.js
var jsEmbed embed.FS

func MakePage(c *gin.Context, title string, page *root.Component, baseUrl string, mm map[string]string) {
	if mm == nil {
		mm = make(map[string]string)
	}
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
