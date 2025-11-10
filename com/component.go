package com

type Component interface {
	SetSlotsAsChildren()
	SlotsAsChildren() bool
	ExtraInfo() *ExtraInfo
	Tag() string
	Name() string
	Props() map[string]string
	Children() []Component
	Slots() []Component
	Contains(s ...Component) Component
	Position(s string) Component
	X(s string) Component
	Y(s string) Component
	W(s string) Component
	H(s string) Component
	V(s string) Component
	X2(s string) Component
	Y2(s string) Component
	Cw(s string) Component
	Ch(s string) Component
	BorderRadius(s string) Component
	Color(s string) Component
	BackgroundColor(s string) Component
	BorderColor(s string) Component
	BoxShadow(s string) Component
	Background(s string) Component
	CaretColor(s string) Component
	UserSelect(s string) Component
	Cursor(s string) Component
	ZIndex(s string) Component
	Opacity(s string) Component
	BorderStyle(s string) Component
	FontFamily(s string) Component
	FontSize(s string) Component
	Outline(s string) Component
	LineHeight(s string) Component
	FontVariantLigatures(s string) Component
	InnerText(s string) Component
	ScrollTop(s string) Component
	ScrollLeft(s string) Component
	BorderLeft(s string) Component
	BorderRight(s string) Component
	BorderTop(s string) Component
	BorderBottom(s string) Component
	Hovered(s string) Component
	HoveredByMouse(s string) Component
	OnClick(s string) Component
	OnDoubleClick(s string) Component
	OnContextMenu(s string) Component
	OnMouseDown(s string) Component
	OnMouseMove(s string) Component
	OnMouseUp(s string) Component
	OnWheel(s string) Component
	OnInput(s string) Component
	OnKeyUp(s string) Component
	OnKeyDown(s string) Component
	OnCompositionStart(s string) Component
	OnCompositionUpdate(s string) Component
	OnCompositionEnd(s string) Component
	OnPaste(s string) Component
	OnCopy(s string) Component
	OnCut(s string) Component
	OnActive(s string) Component
	OnFocus(s string) Component
	OnHover(s string) Component
	OnClickOutside(s string) Component
	OnScrollTop(s string) Component
	OnScrollLeft(s string) Component
}

type ExtraInfo struct {
	name          string
	properties    []string
	methods       []string
	staticMethods []string
	bindJs        map[string]string
	defaultValue  map[string]string
	thisComponent Component
	selfIndex     []int
}

func (e *ExtraInfo) Name() string {
	return e.name
}

func (e *ExtraInfo) SetName(name string) {
	e.name = name
}

func (e *ExtraInfo) Properties() []string {
	return e.properties
}

func (e *ExtraInfo) SetProperties(properties []string) {
	e.properties = properties
}

func (e *ExtraInfo) Methods() []string {
	return e.methods
}

func (e *ExtraInfo) SetMethods(methods []string) {
	e.methods = methods
}

func (e *ExtraInfo) StaticMethods() []string {
	return e.staticMethods
}

func (e *ExtraInfo) SetStaticMethods(staticMethods []string) {
	e.staticMethods = staticMethods
}

func (e *ExtraInfo) BindJs() map[string]string {
	return e.bindJs
}

func (e *ExtraInfo) SetBindJs(bindJs map[string]string) {
	e.bindJs = bindJs
}

func (e *ExtraInfo) DefaultValue() map[string]string {
	return e.defaultValue
}

func (e *ExtraInfo) SetDefaultValue(defaultValue map[string]string) {
	e.defaultValue = defaultValue
}

func (e *ExtraInfo) ThisComponent() Component {
	return e.thisComponent
}

func (e *ExtraInfo) SetThisComponent(thisComponent Component) {
	e.thisComponent = thisComponent
}

func (e *ExtraInfo) SelfIndex() []int {
	return e.selfIndex
}

func (e *ExtraInfo) SetSelfIndex(selfIndex []int) {
	e.selfIndex = selfIndex
}
