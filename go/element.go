package main

// Element представляет визуальный элемент на схеме
type Element struct {
	ID       int
	Name     string
	Type     string
	X        int
	Y        int
	Width    int
	Height   int
	Flag     int // Флаги элемента
	props    map[string]interface{}
	points   []Point
	events   []string
	parentID int
	visible  bool
	locked   bool
}

// Константы флагов элемента
const (
	ELEMENT_FLG_IS_SELECT = 0x01
)

// Константы типов точек
const (
	PT_EVENT = iota
	PT_DATA
	PT_WORK
	PT_VAR
)

// Point представляет точку соединения элемента
type Point struct {
	ID   int
	Name string
	Type int // 0 - вход, 1 - выход
	X    int
	Y    int
}

// NewElement создает новый элемент
func NewElement(id int, name, elemType string) *Element {
	return &Element{
		ID:       id,
		Name:     name,
		Type:     elemType,
		X:        0,
		Y:        0,
		Width:    64,
		Height:   64,
		props:    make(map[string]interface{}),
		points:   make([]Point, 0),
		events:   make([]string, 0),
		parentID: -1,
		visible:  true,
		locked:   false,
	}
}

// SetProp устанавливает свойство элемента
func (e *Element) SetProp(name string, value interface{}) {
	if e.props == nil {
		e.props = make(map[string]interface{})
	}
	e.props[name] = value
}

// GetProp получает свойство элемента
func (e *Element) GetProp(name string) interface{} {
	if e.props == nil {
		return nil
	}
	return e.props[name]
}

// AddPoint добавляет точку соединения
func (e *Element) AddPoint(name string, pType int, x, y int) int {
	id := len(e.points)
	point := Point{
		ID:   id,
		Name: name,
		Type: pType,
		X:    x,
		Y:    y,
	}
	e.points = append(e.points, point)
	return id
}

// GetPoint возвращает точку по ID
func (e *Element) GetPoint(id int) *Point {
	if id < 0 || id >= len(e.points) {
		return nil
	}
	return &e.points[id]
}

// AddEvent добавляет событие
func (e *Element) AddEvent(eventName string) {
	e.events = append(e.events, eventName)
}

// SetPosition устанавливает позицию элемента
func (e *Element) SetPosition(x, y int) {
	e.X = x
	e.Y = y
}

// SetSize устанавливает размер элемента
func (e *Element) SetSize(width, height int) {
	e.Width = width
	e.Height = height
}

// IsVisible возвращает видимость элемента
func (e *Element) IsVisible() bool {
	return e.visible
}

// SetVisible устанавливает видимость элемента
func (e *Element) SetVisible(visible bool) {
	e.visible = visible
}

// IsLocked возвращает состояние блокировки
func (e *Element) IsLocked() bool {
	return e.locked
}

// SetLocked устанавливает блокировку элемента
func (e *Element) SetLocked(locked bool) {
	e.locked = locked
}

// GetParentID возвращает ID родительского элемента
func (e *Element) GetParentID() int {
	return e.parentID
}

// SetParentID устанавливает ID родительского элемента
func (e *Element) SetParentID(id int) {
	e.parentID = id
}

// GetEventCount возвращает количество событий
func (e *Element) GetEventCount() int {
	return len(e.events)
}

// GetEvent возвращает событие по индексу
func (e *Element) GetEvent(index int) string {
	if index < 0 || index >= len(e.events) {
		return ""
	}
	return e.events[index]
}

// IsCore проверяет, является ли элемент ядровым
func (e *Element) IsCore() bool {
	// Заглушка: в оригинале здесь была проверка на ElementCore
	return false
}

// GetRealPoint получает реальную точку для данной точки
func (e *Element) GetRealPoint(p *ElementPoint) *ElementPoint {
	// Заглушка: в оригинале здесь была сложная логика получения реальной точки
	return p
}
