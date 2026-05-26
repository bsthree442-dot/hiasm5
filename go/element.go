package main

import "fmt"

// Constants for element flags
const (
	ElementFlgIsSelect   = 0x01
	ElementFlgIsNoMouse  = 0x02
	ElementFlgIsParent   = 0x04
	ElementFlgIsCore     = 0x08
	ElementFlgIsFreeze   = 0x10
	ElementFlgIsNoDelete = 0x20
	ElementFlgOneWidget  = 0x80
	ElementFlgIsSystem   = 0x400
)

// Constants for save/load modes
const (
	ElementSaveChanged  = 0x01
	ElementSaveSelected = 0x02
	ElementLoadFile     = 0x01
	ElementLoadPaste    = 0x02
)

// Grid constants
const (
	GridSpace = 7
	GridFrm   = 5
)

// SDK object types
type SDKObjectType int

const (
	SDKObjNone SDKObjectType = iota
	SDKObjElement
	SDKObjPoint
	SDKObjLinkPoint
	SDKObjLine
	SDKObjLHint
	SDKObjCount
)

// ObjectType описывает объект схемы, выбранный в редакторе
type ObjectType struct {
	Type SDKObjectType
	Obj1 interface{} // Element, ElementPoint, LinkHint
	Obj2 interface{} // PointPos
}

// NewObjectTypeElement создает ObjectType для элемента
func NewObjectTypeElement(elem *Element) *ObjectType {
	return &ObjectType{Type: SDKObjElement, Obj1: elem}
}

// NewObjectTypePoint создает ObjectType для точки
func NewObjectTypePoint(point *ElementPoint) *ObjectType {
	return &ObjectType{Type: SDKObjPoint, Obj1: point}
}

// NewObjectTypeLinkPoint создает ObjectType для точки пути
func NewObjectTypeLinkPoint(pos *PointPos, point *ElementPoint) *ObjectType {
	return &ObjectType{Type: SDKObjLinkPoint, Obj1: point, Obj2: pos}
}

// Element представляет визуальный элемент на схеме
type Element struct {
	ID           int               // Уникальный ID элемента
	Name         string            // Имя элемента
	Type         string            // Тип элемента
	X            float64           // Позиция X
	Y            float64           // Позиция Y
	Width        float64           // Ширина
	Height       float64           // Высота
	Flag         int               // Флаги ELEMENT_FLG_XXX
	Parent       interface{}       // Родительская SDK (interface{} пока без GUI)
	Pack         *PackElement      // Пакет элементов
	CIndex       int               // Индекс конфигурации
	Points       []*ElementPoint   // Список точек элемента
	Properties   map[string]*PropertyData // Свойства элемента
	CoreData     interface{}       // Данные ядра (для core элементов)
	Selected     bool              // Флаг выделения
}

// NewElement создает новый элемент
func NewElement(id int, name, elemType string, x, y float64) *Element {
	return &Element{
		ID:         id,
		Name:       name,
		Type:       elemType,
		X:          x,
		Y:          y,
		Width:      100, // значение по умолчанию
		Height:     50,  // значение по умолчанию
		Flag:       0,
		Points:     make([]*ElementPoint, 0),
		Properties: make(map[string]*PropertyData),
		Selected:   false,
	}
}

// Destroy освобождает ресурсы элемента
func (e *Element) Destroy() {
	for _, p := range e.Points {
		if p != nil {
			p.Destroy()
		}
	}
	e.Points = nil
	e.Properties = nil
}

// AddPoint добавляет точку к элементу
func (e *Element) AddPoint(name, info string, pointType int) *ElementPoint {
	point := NewElementPoint(e, name, info, pointType)
	e.Points = append(e.Points, point)
	return point
}

// GetPoint получает точку по имени
func (e *Element) GetPoint(name string) *ElementPoint {
	for _, p := range e.Points {
		if p.Name == name {
			return p
		}
	}
	return nil
}

// GetPointByIndex получает точку по индексу
func (e *Element) GetPointByIndex(index int) *ElementPoint {
	if index >= 0 && index < len(e.Points) {
		return e.Points[index]
	}
	return nil
}

// PointCount возвращает количество точек
func (e *Element) PointCount() int {
	return len(e.Points)
}

// SetProp устанавливает свойство элемента
func (e *Element) SetProp(name string, value *TData) {
	if e.Properties == nil {
		e.Properties = make(map[string]*PropertyData)
	}
	if _, exists := e.Properties[name]; !exists {
		e.Properties[name] = &PropertyData{}
	}
	e.Properties[name].Value = value
}

// GetProp получает свойство элемента
func (e *Element) GetProp(name string) *TData {
	if e.Properties == nil {
		return NewTDataNull()
	}
	if prop, exists := e.Properties[name]; exists {
		if prop.Value != nil {
			return prop.Value
		}
	}
	return NewTDataNull()
}

// GetPropStr получает свойство как строку
func (e *Element) GetPropStr(name string) string {
	data := e.GetProp(name)
	return data.ToStr()
}

// GetPropInt получает свойство как целое число
func (e *Element) GetPropInt(name string) int {
	data := e.GetProp(name)
	return data.ToInt()
}

// GetPropReal получает свойство как вещественное число
func (e *Element) GetPropReal(name string) float64 {
	data := e.GetProp(name)
	return data.ToReal()
}

// IsCore проверяет, является ли элемент ядром
func (e *Element) IsCore() bool {
	return (e.Flag & ElementFlgIsCore) != 0
}

// IsSelect проверяет флаг выделения
func (e *Element) IsSelect() bool {
	return (e.Flag & ElementFlgIsSelect) != 0
}

// IsFreeze проверяет флаг заморозки
func (e *Element) IsFreeze() bool {
	return (e.Flag & ElementFlgIsFreeze) != 0
}

// Select выделяет элемент
func (e *Element) Select() {
	e.Flag |= ElementFlgIsSelect
	e.Selected = true
}

// UnSelect снимает выделение с элемента
func (e *Element) UnSelect() {
	e.Flag &= ^ElementFlgIsSelect
	e.Selected = false
}

// Move перемещает элемент на (dx, dy)
func (e *Element) Move(dx, dy float64) {
	e.X += dx
	e.Y += dy
	// Перемещение всех точек элемента
	for _, p := range e.Points {
		if p != nil {
			p.Move(dx, dy)
		}
	}
}

// MoveTo перемещает элемент в позицию (x, y)
func (e *Element) MoveTo(x, y float64) {
	dx := x - e.X
	dy := y - e.Y
	e.Move(dx, dy)
}

// GetRect возвращает прямоугольник элемента
func (e *Element) GetRect() (x, y, w, h float64) {
	return e.X, e.Y, e.Width, e.Height
}

// SetRect устанавливает прямоугольник элемента
func (e *Element) SetRect(x, y, w, h float64) {
	e.X = x
	e.Y = y
	e.Width = w
	e.Height = h
}

// GetRealPoint получает реальную точку без связанных элементов кода
func (e *Element) GetRealPoint(point *ElementPoint) *ElementPoint {
	// Заглушка - в полной версии нужно реализовать логику обхода hub/getdata элементов
	return point
}

// DoWork выполняет работу ядра элемента (для core элементов)
func (e *Element) DoWork(point *ElementPoint, data *TData) {
	if !e.IsCore() {
		return
	}
	// В полной версии здесь вызывается логика конкретного core элемента
}

// ReadVar читает переменную ядра элемента (для core элементов)
func (e *Element) ReadVar(point *ElementPoint, data *TData) {
	if !e.IsCore() {
		return
	}
	// В полной версии здесь читается значение переменной из конкретного core элемента
}

// Serialize сериализует элемент в строку (для сохранения)
func (e *Element) Serialize() string {
	result := fmt.Sprintf("Element:%s,ID:%d,X:%.0f,Y:%.0f,W:%.0f,H:%.0f",
		e.Name, e.ID, e.X, e.Y, e.Width, e.Height)
	return result
}

// Deserialize десериализует элемент из строки (заглушка)
func (e *Element) Deserialize(data string) bool {
	// В полной версии нужно реализовать парсинг строки
	return true
}

// PropertyData хранит данные свойства элемента
type PropertyData struct {
	Name  string
	Value *TData
	Type  int
}

// NewPropertyData создает новые данные свойства
func NewPropertyData(name string, value *TData, propType int) *PropertyData {
	return &PropertyData{Name: name, Value: value, Type: propType}
}

// PackElement представляет элемент пакета (упрощенная версия)
type PackElement struct {
	Name    string
	Type    string
	Icon    string
	Points  []PackPoint
	Props   []PackProperty
}

// PackPoint представляет точку в пакете
type PackPoint struct {
	Name string
	Info string
	Type int
}

// PackProperty представляет свойство в пакете
type PackProperty struct {
	Name    string
	Type    int
	Default *TData
}
