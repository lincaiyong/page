package com

type Component interface {
	ExtraInfo() *ExtraInfo
	Tag() string
	Name() string
	Props() map[string]string
	Methods() map[string]string
	Children() []Component
	Slots() []Component
	SlotsAsChildren() bool
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
