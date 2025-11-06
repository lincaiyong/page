package com

type Component interface {
	Tag() string
	Name() string
	Props() map[string]string
	StaticProps() map[string]string
	Children() []Component
	Slots() []Component
	Contains(s ...Component) Component
	X_(i int) Component
	Y_(i int) Component
	W_(i int) Component
	H_(i int) Component
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
