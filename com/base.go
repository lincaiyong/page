package com

import "strconv"

func NewBaseComponent(tag string, self Component, children ...Component) *BaseComponent {
	return &BaseComponent{
		self:     self,
		tag:      tag,
		children: children,
		props: map[string]string{
			"ch":      ".h - .borderTop - .borderBottom",
			"cw":      ".w - .borderLeft - .borderRight",
			"hovered": ".hoveredByMouse",
			"x2":      ".x + .w",
			"y2":      ".y + .h",
		},
		staticProps: map[string]string{},
	}
}

type BaseComponent struct {
	self            Component
	name            string
	tag             string
	children        []Component
	slots           []Component
	props           map[string]string
	staticProps     map[string]string
	slotsAsChildren bool
}

func (b *BaseComponent) SlotsAsChildren(v bool) {
	b.slotsAsChildren = v
}

func (b *BaseComponent) Name() string {
	if b.name != "" {
		return b.name
	}
	return b.tag
}

func (b *BaseComponent) Tag() string {
	return b.tag
}

func (b *BaseComponent) Children() []Component {
	return b.children
}

func (b *BaseComponent) Slots() []Component {
	return b.slots
}

func (b *BaseComponent) Props() map[string]string {
	return b.props
}

func (b *BaseComponent) StaticProps() map[string]string {
	return b.staticProps
}

func (b *BaseComponent) NameAs(name string) Component {
	b.name = name
	return b.self
}

func (b *BaseComponent) Contains(s ...Component) Component {
	if b.slotsAsChildren {
		b.children = append(b.children, s...)
	} else {
		b.slots = append(b.slots, s...)
	}
	return b.self
}

func (b *BaseComponent) X_(i int) Component {
	b.Props()["x"] = strconv.Itoa(i)
	return b.self
}

func (b *BaseComponent) Y_(i int) Component {
	b.Props()["y"] = strconv.Itoa(i)
	return b.self
}

func (b *BaseComponent) W_(i int) Component {
	b.Props()["w"] = strconv.Itoa(i)
	return b.self
}

func (b *BaseComponent) H_(i int) Component {
	b.Props()["h"] = strconv.Itoa(i)
	return b.self
}

func (b *BaseComponent) Position(s string) Component {
	b.Props()["position"] = s
	return b.self
}

func (b *BaseComponent) X(s string) Component {
	b.Props()["x"] = s
	return b.self
}

func (b *BaseComponent) Y(s string) Component {
	b.Props()["y"] = s
	return b.self
}

func (b *BaseComponent) W(s string) Component {
	b.Props()["w"] = s
	return b.self
}

func (b *BaseComponent) H(s string) Component {
	b.Props()["h"] = s
	return b.self
}

func (b *BaseComponent) V(s string) Component {
	b.Props()["v"] = s
	return b.self
}

func (b *BaseComponent) X2(s string) Component {
	b.Props()["x2"] = s
	return b.self
}

func (b *BaseComponent) Y2(s string) Component {
	b.Props()["y2"] = s
	return b.self
}

func (b *BaseComponent) Cw(s string) Component {
	b.Props()["cw"] = s
	return b.self
}

func (b *BaseComponent) Ch(s string) Component {
	b.Props()["ch"] = s
	return b.self
}

func (b *BaseComponent) BorderRadius(s string) Component {
	b.Props()["borderRadius"] = s
	return b.self
}

func (b *BaseComponent) Color(s string) Component {
	b.Props()["color"] = s
	return b.self
}

func (b *BaseComponent) BackgroundColor(s string) Component {
	b.Props()["backgroundColor"] = s
	return b.self
}

func (b *BaseComponent) BorderColor(s string) Component {
	b.Props()["borderColor"] = s
	return b.self
}

func (b *BaseComponent) BoxShadow(s string) Component {
	b.Props()["boxShadow"] = s
	return b.self
}

func (b *BaseComponent) Background(s string) Component {
	b.Props()["background"] = s
	return b.self
}

func (b *BaseComponent) CaretColor(s string) Component {
	b.Props()["caretColor"] = s
	return b.self
}

func (b *BaseComponent) UserSelect(s string) Component {
	b.Props()["userSelect"] = s
	return b.self
}

func (b *BaseComponent) Cursor(s string) Component {
	b.Props()["cursor"] = s
	return b.self
}

func (b *BaseComponent) ZIndex(s string) Component {
	b.Props()["zIndex"] = s
	return b.self
}

func (b *BaseComponent) Opacity(s string) Component {
	b.Props()["opacity"] = s
	return b.self
}

func (b *BaseComponent) BorderStyle(s string) Component {
	b.Props()["borderStyle"] = s
	return b.self
}

func (b *BaseComponent) FontFamily(s string) Component {
	b.Props()["fontFamily"] = s
	return b.self
}

func (b *BaseComponent) FontSize(s string) Component {
	b.Props()["fontSize"] = s
	return b.self
}

func (b *BaseComponent) Outline(s string) Component {
	b.Props()["outline"] = s
	return b.self
}

func (b *BaseComponent) LineHeight(s string) Component {
	b.Props()["lineHeight"] = s
	return b.self
}

func (b *BaseComponent) FontVariantLigatures(s string) Component {
	b.Props()["fontVariantLigatures"] = s
	return b.self
}

func (b *BaseComponent) InnerText(s string) Component {
	b.Props()["innerText"] = s
	return b.self
}

func (b *BaseComponent) ScrollTop(s string) Component {
	b.Props()["scrollTop"] = s
	return b.self
}

func (b *BaseComponent) ScrollLeft(s string) Component {
	b.Props()["scrollLeft"] = s
	return b.self
}

func (b *BaseComponent) BorderLeft(s string) Component {
	b.Props()["borderLeft"] = s
	return b.self
}

func (b *BaseComponent) BorderRight(s string) Component {
	b.Props()["borderRight"] = s
	return b.self
}

func (b *BaseComponent) BorderTop(s string) Component {
	b.Props()["borderTop"] = s
	return b.self
}

func (b *BaseComponent) BorderBottom(s string) Component {
	b.Props()["borderBottom"] = s
	return b.self
}

func (b *BaseComponent) Hovered(s string) Component {
	b.Props()["hovered"] = s
	return b.self
}

func (b *BaseComponent) HoveredByMouse(s string) Component {
	b.Props()["hoveredByMouse"] = s
	return b.self
}

func (b *BaseComponent) OnClick(s string) Component {
	b.Props()["onClick"] = s
	return b.self
}

func (b *BaseComponent) OnDoubleClick(s string) Component {
	b.Props()["onDoubleClick"] = s
	return b.self
}

func (b *BaseComponent) OnContextMenu(s string) Component {
	b.Props()["onContextMenu"] = s
	return b.self
}

func (b *BaseComponent) OnMouseDown(s string) Component {
	b.Props()["onMouseDown"] = s
	return b.self
}

func (b *BaseComponent) OnMouseMove(s string) Component {
	b.Props()["onMouseMove"] = s
	return b.self
}

func (b *BaseComponent) OnMouseUp(s string) Component {
	b.Props()["onMouseUp"] = s
	return b.self
}

func (b *BaseComponent) OnWheel(s string) Component {
	b.Props()["onWheel"] = s
	return b.self
}

func (b *BaseComponent) OnInput(s string) Component {
	b.Props()["onInput"] = s
	return b.self
}

func (b *BaseComponent) OnKeyUp(s string) Component {
	b.Props()["onKeyUp"] = s
	return b.self
}

func (b *BaseComponent) OnKeyDown(s string) Component {
	b.Props()["onKeyDown"] = s
	return b.self
}

func (b *BaseComponent) OnCompositionStart(s string) Component {
	b.Props()["onCompositionStart"] = s
	return b.self
}

func (b *BaseComponent) OnCompositionUpdate(s string) Component {
	b.Props()["onCompositionUpdate"] = s
	return b.self
}

func (b *BaseComponent) OnCompositionEnd(s string) Component {
	b.Props()["onCompositionEnd"] = s
	return b.self
}

func (b *BaseComponent) OnPaste(s string) Component {
	b.Props()["onPaste"] = s
	return b.self
}

func (b *BaseComponent) OnCopy(s string) Component {
	b.Props()["onCopy"] = s
	return b.self
}

func (b *BaseComponent) OnCut(s string) Component {
	b.Props()["onCut"] = s
	return b.self
}

func (b *BaseComponent) OnActive(s string) Component {
	b.Props()["onActive"] = s
	return b.self
}

func (b *BaseComponent) OnFocus(s string) Component {
	b.Props()["onFocus"] = s
	return b.self
}

func (b *BaseComponent) OnHover(s string) Component {
	b.Props()["onHover"] = s
	return b.self
}

func (b *BaseComponent) OnClickOutside(s string) Component {
	b.Props()["onClickOutside"] = s
	return b.self
}

func (b *BaseComponent) OnScrollTop(s string) Component {
	b.Props()["onScrollTop"] = s
	return b.self
}

func (b *BaseComponent) OnScrollLeft(s string) Component {
	b.Props()["onScrollLeft"] = s
	return b.self
}
