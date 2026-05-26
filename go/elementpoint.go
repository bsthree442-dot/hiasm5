package main

import "fmt"

// Constants for point flags and spacing
const (
	PointFlgIsSelect = 0x01
	PointOff         = 3
	PointSpace       = 7
)

// Point types
const (
	PtWork  = 2
	PtEvent = 0
	PtVar   = 3
	PtData  = 1
)

// PointPos представляет точку пути связи между двумя точками элементов
type PointPos struct {
	X    float64
	Y    float64
	Next *PointPos
	Prev *PointPos
}

// NewPointPos создает новую точку пути
func NewPointPos() *PointPos {
	return &PointPos{}
}

// ElementPoint представляет точку элемента
type ElementPoint struct {
	Pos      *PointPos     // Позиция точки элемента
	Name     string        // Имя точки
	Info     string        // Описание точки
	Type     int           // Тип точки (pt_work, pt_event, pt_var, pt_data)
	DataType DataType      // Тип данных точки
	Parent   *Element      // Указатель на родительский элемент
	Point    *ElementPoint // Указатель на связанную точку (nil по умолчанию)
	Flag     int           // Флаги POINT_FLG_XXX
}

// NewElementPoint создает новую точку элемента
func NewElementPoint(parent *Element, name, info string, pointType int) *ElementPoint {
	ep := &ElementPoint{
		Name:     name,
		Info:     info,
		Type:     pointType,
		Parent:   parent,
		DataType: DataNull,
		Pos:      NewPointPos(),
		Flag:     0,
	}
	ep.Pos.X = 0
	ep.Pos.Y = 0
	return ep
}

// Destroy освобождает ресурсы точки
func (ep *ElementPoint) Destroy() {
	ep.Clear()
	ep.Pos = nil
}

// Connect связывает две точки без трассировки пути
func (ep *ElementPoint) Connect(link *ElementPoint) *ElementPoint {
	var p1, p2 *ElementPoint

	if ep.IsPrimary() {
		p1 = ep
		p2 = link
	} else {
		p1 = link
		p2 = ep
	}
	p1.Point = p2
	p2.Point = p1
	p1.Pos.Next = p2.Pos
	p2.Pos.Prev = p1.Pos

	return ep
}

// CanConnect проверяет возможность соединения двух точек
func (ep *ElementPoint) CanConnect(link *ElementPoint) bool {
	if link == nil {
		return false
	}
	if ep.Point != nil || link.Point != nil {
		return false
	}
	// Проверка: abs(type + link->type - 5) == 2
	sum := ep.Type + link.Type - 5
	if sum < 0 {
		sum = -sum
	}
	return sum == 2
}

// Clear удаляет связь
func (ep *ElementPoint) Clear() {
	if ep.Pos.Next != nil {
		// В оригинале здесь была отрисовка, заглушка
		for ep.Pos.Next.Next != nil {
			pb := ep.Pos.Next.Next
			ep.Pos.Next = pb
		}
		ep.Pos.Next = nil
	} else {
		ep.Pos.Prev = nil
	}

	if ep.Point != nil {
		ep.Point.Point = nil
		ep.Point.Clear()
		ep.Point = nil
	}
}

// CreatePath создает путь между двумя точками (только после connect)
func (ep *ElementPoint) CreatePath() {
	if ep.Point == nil || ep.Point.Point == nil {
		return
	}

	var point1, point2 *ElementPoint
	if ep.IsPrimary() {
		point1 = ep
		point2 = ep.Point
	} else {
		point1 = ep.Point
		point2 = point1.Point
	}

	// Упрощенная версия tracePath без сложной логики отрисовки
	// В полной версии нужно реализовать алгоритм трассировки из C++ кода
	AddLinePoint(point1.Pos, point1.Pos.X, point1.Pos.Y)
	AddLinePoint(point1.Pos, point2.Pos.X, point2.Pos.Y)
}

// RemoveLinePoint удаляет точку из пути связи
func (ep *ElementPoint) RemoveLinePoint(lp *PointPos) {
	if lp.Prev != nil && lp.Next != nil {
		lp.Prev.Next = lp.Next
		lp.Next.Prev = lp.Prev
		// В оригинале здесь была отрисовка
	}
}

// MoveLinePoint перемещает точку пути в позицию (x, y)
func (ep *ElementPoint) MoveLinePoint(lp *PointPos, x, y float64) {
	if lp.Next != nil && lp.Next.Next != nil && lp.X == lp.Next.X && lp.Y != lp.Next.Y {
		lp.X = x
		lp.Next.X = x
	} else if lp.Prev != nil && lp.Prev.Prev != nil && lp.X == lp.Prev.X && lp.Y != lp.Prev.Y {
		lp.X = x
		lp.Prev.X = x
	} else if lp.Next != nil && absFloat(lp.Next.X-x) < 5 {
		lp.X = lp.Next.X
	} else if lp.Prev != nil && absFloat(lp.Prev.X-x) < 5 {
		lp.X = lp.Prev.X
	} else {
		lp.X = x
	}

	if lp.Next != nil && lp.Next.Next != nil && lp.Y == lp.Next.Y && lp.X != lp.Next.X {
		lp.Y = y
		lp.Next.Y = y
	} else if lp.Prev != nil && lp.Prev.Prev != nil && lp.Y == lp.Prev.Y && lp.X != lp.Prev.X {
		lp.Y = y
		lp.Prev.Y = y
	} else if lp.Next != nil && absFloat(lp.Next.Y-y) < 5 {
		lp.Y = lp.Next.Y
	} else if lp.Prev != nil && absFloat(lp.Prev.Y-y) < 5 {
		lp.Y = lp.Prev.Y
	} else {
		lp.Y = y
	}
}

// AddLinePoint добавляет точку в путь связи
func AddLinePoint(lp *PointPos, x, y float64) *PointPos {
	np := NewPointPos()
	np.X = x
	np.Y = y

	if lp != nil {
		np.Next = lp.Next
		np.Prev = lp
		lp.Next = np
		if np.Next != nil {
			np.Next.Prev = np
		}
	}
	return np
}

// IsPrimary проверяет, является ли точка первичной (pt_event или pt_data)
func (ep *ElementPoint) IsPrimary() bool {
	return ep.Type%2 == 0
}

// IsSelect проверяет флаг POINT_FLG_IS_SELECT
func (ep *ElementPoint) IsSelect() bool {
	return (ep.Flag & PointFlgIsSelect) != 0
}

// Move перемещает точку на указанный оффсет
func (ep *ElementPoint) Move(dx, dy float64) {
	pb := ep.Pos
	if ep.Point != nil && (ep.Point.Parent.Flag&ElementFlgIsSelect) != 0 {
		if pb.Next != nil && ep.IsPrimary() {
			ep.movePoints(pb.Next, dx, dy)
		}
	} else {
		if pb.Next != nil && pb.Next.Next != nil {
			if pb.Next.X == pb.X {
				pb.Next.X += dx
				if pb.Y+dy-pb.Next.Y < 3 {
					ep.movePos(pb.Next, pb.Next.X, pb.Next.Y+dy)
				}
			} else if pb.Next.Y == pb.Y {
				pb.Next.Y += dy
				if pb.Next.X-(pb.X+dx) < 3 {
					ep.movePos(pb.Next, pb.Next.X+dx, pb.Next.Y)
				}
			}
		}

		if pb.Prev != nil && pb.Prev.Prev != nil {
			if pb.Prev.Y == pb.Y {
				pb.Prev.Y += dy
				if absFloat(pb.X+dx-pb.Prev.X) < 3 {
					ep.movePos(pb.Prev, pb.Prev.X+dx, pb.Prev.Y)
				}
			} else if pb.Prev.X == pb.X {
				pb.Prev.X += dx
				if absFloat(pb.Prev.Y-(pb.Y+dy)) < 3 {
					ep.movePos(pb.Prev, pb.Prev.X, pb.Prev.Y+dy)
				}
			}
		}
	}
	pb.Y += dy
	pb.X += dx
}

// movePoints сдвигает все точки связи на (dx, dy)
func (ep *ElementPoint) movePoints(pb *PointPos, dx, dy float64) {
	for pb.Next != nil {
		pb.X += dx
		pb.Y += dy
		pb = pb.Next
	}
}

// movePos перемещает точку связи в позицию (x, y)
func (ep *ElementPoint) movePos(pb *PointPos, x, y float64) {
	if pb.Next != nil && pb.Next.Next != nil && pb.X == pb.Next.X && pb.Y != pb.Next.Y {
		pb.X = x
		pb.Next.X = x
	} else if pb.Prev != nil && pb.Prev.Prev != nil && pb.X == pb.Prev.X && pb.Y != pb.Prev.Y {
		pb.X = x
		pb.Prev.X = x
	} else if pb.Next != nil && absFloat(pb.Next.X-x) < 5 {
		pb.X = pb.Next.X
	} else if pb.Prev != nil && absFloat(pb.Prev.X-x) < 5 {
		pb.X = pb.Prev.X
	} else {
		pb.X = x
	}

	if pb.Next != nil && pb.Next.Next != nil && pb.Y == pb.Next.Y && pb.X != pb.Next.X {
		pb.Y = y
		pb.Next.Y = y
	} else if pb.Prev != nil && pb.Prev.Prev != nil && pb.Y == pb.Prev.Y && pb.X != pb.Prev.X {
		pb.Y = y
		pb.Prev.Y = y
	} else if pb.Next != nil && absFloat(pb.Next.Y-y) < 5 {
		pb.Y = pb.Next.Y
	} else if pb.Prev != nil && absFloat(pb.Prev.Y-y) < 5 {
		pb.Y = pb.Prev.Y
	} else {
		pb.Y = y
	}
}

// DrawRect вычисляет прямоугольник отрисовки точки
func (ep *ElementPoint) DrawRect() (x1, y1, x2, y2 float64) {
	var p *PointPos
	if ep.Type == PtEvent || ep.Type == PtData || ep.Point == nil {
		p = ep.Pos
	} else {
		p = ep.Point.Pos
	}
	x1 = p.X
	y1 = p.Y
	x2 = p.X
	y2 = p.Y

	curr := p.Next
	for curr != nil {
		if curr.X < x1 {
			x1 = curr.X
		}
		if curr.Y < y1 {
			y1 = curr.Y
		}
		if curr.X > x2 {
			x2 = curr.X
		}
		if curr.Y > y2 {
			y2 = curr.Y
		}
		curr = curr.Next
	}
	return x1 - PointOff, y1 - PointOff, x2 - x1 + PointSpace, y2 - y1 + PointSpace
}

// SerializePath сохраняет данные пути в строку в формате: (x1,y1),(x2,y2),...,(xN,yN)
func (ep *ElementPoint) SerializePath() string {
	path := ""
	p := ep.Pos.Next
	for p != nil && p.Next != nil {
		path += fmt.Sprintf("(%.0f,%.0f)", p.X, p.Y)
		p = p.Next
	}
	return path
}

// GetRealPointWithPath получает реальную точку без связанных элементов кода
func (ep *ElementPoint) GetRealPointWithPath() *ElementPoint {
	root := ep
	result := ep
	var p *ElementPoint

	for {
		p = result
		if p.Parent != nil {
			result = p.Parent.GetRealPoint(p)
		} else {
			break
		}
		if p == result || result == nil || root == result {
			break
		}
	}

	if root == result {
		result = p
	}
	return result
}

// GetPair возвращает тип парной точки для подключения
func (ep *ElementPoint) GetPair() int {
	if ep.IsPrimary() {
		return ep.Type - 1
	}
	return ep.Type + 1
}

// GetDirection возвращает направление точки (1 для work-event, 0 для var-data)
func (ep *ElementPoint) GetDirection() int {
	if ep.Type == PtEvent || ep.Type == PtWork {
		return 1
	}
	return 0
}

// OnEvent обрабатывает событие с данными
func (ep *ElementPoint) OnEvent(data *TData) {
	if ep.Point != nil {
		e := ep.Point.Parent
		if e != nil && e.IsCore() {
			// Вызов do_work у ядра элемента
			e.DoWork(ep.Point, data)
		}
	}
}

// OnEventString обрабатывает событие со строковым значением
func (ep *ElementPoint) OnEventString(value string) {
	data := NewTDataStr(value)
	ep.OnEvent(data)
}

// OnEventInt обрабатывает событие с целочисленным значением
func (ep *ElementPoint) OnEventInt(value int) {
	data := NewTDataInt(value)
	ep.OnEvent(data)
}

// OnEventReal обрабатывает событие с вещественным значением
func (ep *ElementPoint) OnEventReal(value float64) {
	data := NewTDataReal(value)
	ep.OnEvent(data)
}

// GetData получает данные
func (ep *ElementPoint) GetData(data *TData) {
	if ep.Point != nil {
		e := ep.Point.Parent
		if e != nil && e.IsCore() {
			// Вызов read_var у ядра элемента
			e.ReadVar(ep.Point, data)
		}
	}
}

// Invalidate обновляет область отрисовки (заглушка без GUI)
func (ep *ElementPoint) Invalidate() {
	// Заглушка - в оригинале вызывалась перерисовка
}

// Вспомогательная функция
func absFloat(x float64) float64 {
	if x < 0 {
		return -x
	}
	return x
}
