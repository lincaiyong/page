package com

import (
	"github.com/lincaiyong/log"
	"strings"
)

func NewBaseComponent[T any](tag string, self *T, children ...Component) *BaseComponent[T] {
	return &BaseComponent[T]{
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
	}
}

type BaseComponent[T any] struct {
	extraInfo ExtraInfo
	self      *T
	name      string
	tag       string
	children  []Component
	slots     []Component
	props     map[string]string
}

func (b *BaseComponent[T]) Name() string {
	if b.name != "" {
		return b.name
	}
	return b.tag
}

func (b *BaseComponent[T]) ExtraInfo() *ExtraInfo {
	return &b.extraInfo
}

func (b *BaseComponent[T]) Tag() string {
	return b.tag
}

func (b *BaseComponent[T]) Children() []Component {
	return b.children
}

func (b *BaseComponent[T]) Slots() []Component {
	return b.slots
}

func (b *BaseComponent[T]) Props() map[string]string {
	return b.props
}

func (b *BaseComponent[T]) SetProp(k, v string) {
	b.props[k] = v
}

func (b *BaseComponent[T]) NameAs(name string) *T {
	if !strings.HasSuffix(name, "Ele") {
		log.FatalLog("invalid element name: %s", name)
	}
	b.name = name
	return b.self
}

func (b *BaseComponent[T]) Contains(s ...Component) *T {
	b.slots = append(b.slots, s...)
	return b.self
}

func (b *BaseComponent[T]) OnCreated(fn string) *T {
	b.Props()["onCreatedFn"] = fn
	return b.self
}

func (b *BaseComponent[T]) OnUpdated(fn string) *T {
	b.Props()["onUpdated"] = fn
	return b.self
}

func (b *BaseComponent[T]) Position(s string) *T {
	b.Props()["position"] = s
	return b.self
}

func (b *BaseComponent[T]) X(s string) *T {
	b.Props()["x"] = s
	return b.self
}

func (b *BaseComponent[T]) Y(s string) *T {
	b.Props()["y"] = s
	return b.self
}

func (b *BaseComponent[T]) W(s string) *T {
	b.Props()["w"] = s
	return b.self
}

func (b *BaseComponent[T]) H(s string) *T {
	b.Props()["h"] = s
	return b.self
}

func (b *BaseComponent[T]) V(s string) *T {
	b.Props()["v"] = s
	return b.self
}

func (b *BaseComponent[T]) X2(s string) *T {
	b.Props()["x2"] = s
	return b.self
}

func (b *BaseComponent[T]) Y2(s string) *T {
	b.Props()["y2"] = s
	return b.self
}

func (b *BaseComponent[T]) Cw(s string) *T {
	b.Props()["cw"] = s
	return b.self
}

func (b *BaseComponent[T]) Ch(s string) *T {
	b.Props()["ch"] = s
	return b.self
}

func (b *BaseComponent[T]) BorderRadius(s string) *T {
	b.Props()["borderRadius"] = s
	return b.self
}

func (b *BaseComponent[T]) Color(s string) *T {
	b.Props()["color"] = s
	return b.self
}

func (b *BaseComponent[T]) BackgroundColor(s string) *T {
	b.Props()["backgroundColor"] = s
	return b.self
}

func (b *BaseComponent[T]) BorderColor(s string) *T {
	b.Props()["borderColor"] = s
	return b.self
}

func (b *BaseComponent[T]) BoxShadow(s string) *T {
	b.Props()["boxShadow"] = s
	return b.self
}

func (b *BaseComponent[T]) Background(s string) *T {
	b.Props()["background"] = s
	return b.self
}

func (b *BaseComponent[T]) CaretColor(s string) *T {
	b.Props()["caretColor"] = s
	return b.self
}

func (b *BaseComponent[T]) UserSelect(s string) *T {
	b.Props()["userSelect"] = s
	return b.self
}

func (b *BaseComponent[T]) Cursor(s string) *T {
	b.Props()["cursor"] = s
	return b.self
}

func (b *BaseComponent[T]) ZIndex(s string) *T {
	b.Props()["zIndex"] = s
	return b.self
}

func (b *BaseComponent[T]) Opacity(s string) *T {
	b.Props()["opacity"] = s
	return b.self
}

func (b *BaseComponent[T]) BorderStyle(s string) *T {
	b.Props()["borderStyle"] = s
	return b.self
}

func (b *BaseComponent[T]) FontFamily(s string) *T {
	b.Props()["fontFamily"] = s
	return b.self
}

func (b *BaseComponent[T]) FontSize(s string) *T {
	b.Props()["fontSize"] = s
	return b.self
}

func (b *BaseComponent[T]) Outline(s string) *T {
	b.Props()["outline"] = s
	return b.self
}

func (b *BaseComponent[T]) LineHeight(s string) *T {
	b.Props()["lineHeight"] = s
	return b.self
}

func (b *BaseComponent[T]) FontVariantLigatures(s string) *T {
	b.Props()["fontVariantLigatures"] = s
	return b.self
}

func (b *BaseComponent[T]) InnerText(s string) *T {
	b.Props()["innerText"] = s
	return b.self
}

func (b *BaseComponent[T]) ScrollTop(s string) *T {
	b.Props()["scrollTop"] = s
	return b.self
}

func (b *BaseComponent[T]) ScrollLeft(s string) *T {
	b.Props()["scrollLeft"] = s
	return b.self
}

func (b *BaseComponent[T]) BorderLeft(s string) *T {
	b.Props()["borderLeft"] = s
	return b.self
}

func (b *BaseComponent[T]) BorderRight(s string) *T {
	b.Props()["borderRight"] = s
	return b.self
}

func (b *BaseComponent[T]) BorderTop(s string) *T {
	b.Props()["borderTop"] = s
	return b.self
}

func (b *BaseComponent[T]) BorderBottom(s string) *T {
	b.Props()["borderBottom"] = s
	return b.self
}

func (b *BaseComponent[T]) Hovered(s string) *T {
	b.Props()["hovered"] = s
	return b.self
}

func (b *BaseComponent[T]) HoveredByMouse(s string) *T {
	b.Props()["hoveredByMouse"] = s
	return b.self
}

func (b *BaseComponent[T]) OnClick(s string) *T {
	b.Props()["onClick"] = s
	return b.self
}

func (b *BaseComponent[T]) OnDoubleClick(s string) *T {
	b.Props()["onDoubleClick"] = s
	return b.self
}

func (b *BaseComponent[T]) OnContextMenu(s string) *T {
	b.Props()["onContextMenu"] = s
	return b.self
}

func (b *BaseComponent[T]) OnMouseDown(s string) *T {
	b.Props()["onMouseDown"] = s
	return b.self
}

func (b *BaseComponent[T]) OnMouseMove(s string) *T {
	b.Props()["onMouseMove"] = s
	return b.self
}

func (b *BaseComponent[T]) OnMouseUp(s string) *T {
	b.Props()["onMouseUp"] = s
	return b.self
}

func (b *BaseComponent[T]) OnWheel(s string) *T {
	b.Props()["onWheel"] = s
	return b.self
}

func (b *BaseComponent[T]) OnInput(s string) *T {
	b.Props()["onInput"] = s
	return b.self
}

func (b *BaseComponent[T]) OnKeyUp(s string) *T {
	b.Props()["onKeyUp"] = s
	return b.self
}

func (b *BaseComponent[T]) OnKeyDown(s string) *T {
	b.Props()["onKeyDown"] = s
	return b.self
}

func (b *BaseComponent[T]) OnCompositionStart(s string) *T {
	b.Props()["onCompositionStart"] = s
	return b.self
}

func (b *BaseComponent[T]) OnCompositionUpdate(s string) *T {
	b.Props()["onCompositionUpdate"] = s
	return b.self
}

func (b *BaseComponent[T]) OnCompositionEnd(s string) *T {
	b.Props()["onCompositionEnd"] = s
	return b.self
}

func (b *BaseComponent[T]) OnPaste(s string) *T {
	b.Props()["onPaste"] = s
	return b.self
}

func (b *BaseComponent[T]) OnCopy(s string) *T {
	b.Props()["onCopy"] = s
	return b.self
}

func (b *BaseComponent[T]) OnCut(s string) *T {
	b.Props()["onCut"] = s
	return b.self
}

func (b *BaseComponent[T]) OnActive(s string) *T {
	b.Props()["onActive"] = s
	return b.self
}

func (b *BaseComponent[T]) OnFocus(s string) *T {
	b.Props()["onFocus"] = s
	return b.self
}

func (b *BaseComponent[T]) OnHover(s string) *T {
	b.Props()["onHover"] = s
	return b.self
}

func (b *BaseComponent[T]) OnClickOutside(s string) *T {
	b.Props()["onClickOutside"] = s
	return b.self
}

func (b *BaseComponent[T]) OnScrollTop(s string) *T {
	b.Props()["onScrollTop"] = s
	return b.self
}

func (b *BaseComponent[T]) OnScrollLeft(s string) *T {
	b.Props()["onScrollLeft"] = s
	return b.self
}
