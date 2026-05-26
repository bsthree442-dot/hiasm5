package main

import "fmt"

// Константы для флагов точек
const (
	POINT_FLG_IS_SELECT = 0x01
	POINT_OFF           = 3
	POINT_SPACE         = 7
)

// PointPos представляет точку пути связи между двумя точками элементов
type PointPos struct {
	X    float64   // X позиция точки пути
	Y    float64   // Y позиция точки пути
	Next *PointPos // Указатель на следующую точку пути
	Prev *PointPos // Указатель на предыдущую точку пути
}

// ElementPoint представляет точку элемента
type ElementPoint struct {
	Pos      *PointPos     // Позиция точки элемента
	Name     string        // Имя точки
	Info     string        // Описание точки
	Type     int           // Тип точки (pt_work, pt_event, pt_var, pt_data)
	DataType int           // Тип данных точки (data_int, data_str, data_real, ...)
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
		DataType: DATA_NULL,
		Pos: &PointPos{
			X:    0,
			Y:    0,
			Next: nil,
			Prev: nil,
		},
		Point: nil,
		Flag:  0,
	}
	return ep
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
	// Проверка совместимости типов (abs(type + link->type - 5) == 2)
	diff := ep.Type + link.Type - 5
	if diff < 0 {
		diff = -diff
	}
	return diff == 2
}

// Clear удаляет связь
func (ep *ElementPoint) Clear() {
	if ep.Pos.Next != nil {
		// В оригинале здесь было обновление области перерисовки
		// r := ep.DrawRect()
		for ep.Pos.Next.Next != nil {
			pb := ep.Pos.Next.Next
			// Удаляем промежуточную точку
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
		point2 = ep
	}

	// В оригинале здесь сложная логика трассировки пути с учетом типов элементов
	// Для заглушки просто создадим базовый путь
	_ = point1
	_ = point2

	// В оригинале: tracePath(point1.pos, v1, v2)
	// И обновление области перерисовки: parent->parent->on_redraw_rect.run(&r)
}

// MoveLinePoint перемещает точку связи в позицию (x,y)
func (ep *ElementPoint) MoveLinePoint(lp *PointPos, x, y float64) {
	if lp.Next != nil && lp.Next.Next != nil && lp.X == lp.Next.X && lp.Y != lp.Next.Y {
		lp.X = x
		lp.Next.X = x
	} else if lp.Prev != nil && lp.Prev.Prev != nil && lp.X == lp.Prev.X && lp.Y != lp.Prev.Y {
		lp.X = x
		lp.Prev.X = x
	} else if lp.Next != nil && absFloat64(lp.Next.X-x) < 5 {
		lp.X = lp.Next.X
	} else if lp.Prev != nil && absFloat64(lp.Prev.X-x) < 5 {
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
	} else if lp.Next != nil && absFloat64(lp.Next.Y-y) < 5 {
		lp.Y = lp.Next.Y
	} else if lp.Prev != nil && absFloat64(lp.Prev.Y-y) < 5 {
		lp.Y = lp.Prev.Y
	} else {
		lp.Y = y
	}

	// В оригинале здесь обновление области перерисовки
}

// RemoveLinePoint удаляет точку из связи
func (ep *ElementPoint) RemoveLinePoint(lp *PointPos) {
	if lp.Prev != nil {
		lp.Prev.Next = lp.Next
	}
	if lp.Next != nil {
		lp.Next.Prev = lp.Prev
	}
	// Удаляем точку (в Go сборщик мусора сделает это сам)
	_ = lp
}

// AddLinePoint добавляет точку к связи в позиции (x,y)
func AddLinePoint(lp *PointPos, x, y float64) *PointPos {
	np := &PointPos{
		X:    x,
		Y:    y,
		Next: nil,
		Prev: nil,
	}

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
	return (ep.Flag & POINT_FLG_IS_SELECT) != 0
}

// Move перемещает точку на указанное смещение
func (ep *ElementPoint) Move(dx, dy float64) {
	pb := ep.Pos

	if ep.Point != nil && (ep.Point.Parent.Flag&ELEMENT_FLG_IS_SELECT) != 0 {
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
				if absFloat64(pb.X+dx-pb.Prev.X) < 3 {
					ep.movePos(pb.Prev, pb.Prev.X+dx, pb.Prev.Y)
				}
			} else if pb.Prev.X == pb.X {
				pb.Prev.X += dx
				if absFloat64(pb.Prev.Y-(pb.Y+dy)) < 3 {
					ep.movePos(pb.Prev, pb.Prev.X, pb.Prev.Y+dy)
				}
			}
		}
	}

	pb.Y += dy
	pb.X += dx
}

// movePoints вспомогательная функция для перемещения всех точек связи
func (ep *ElementPoint) movePoints(pb *PointPos, dx, dy float64) {
	for pb != nil {
		pb.X += dx
		pb.Y += dy
		pb = pb.Next
	}
}

// movePos вспомогательная функция для перемещения точки связи в позицию
func (ep *ElementPoint) movePos(pb *PointPos, x, y float64) {
	if pb.Next != nil && pb.Next.Next != nil && pb.X == pb.Next.X && pb.Y != pb.Next.Y {
		pb.X = x
		pb.Next.X = x
	} else if pb.Prev != nil && pb.Prev.Prev != nil && pb.X == pb.Prev.X && pb.Y != pb.Prev.Y {
		pb.X = x
		pb.Prev.X = x
	} else if pb.Next != nil && absFloat64(pb.Next.X-x) < 5 {
		pb.X = pb.Next.X
	} else if pb.Prev != nil && absFloat64(pb.Prev.X-x) < 5 {
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
	} else if pb.Next != nil && absFloat64(pb.Next.Y-y) < 5 {
		pb.Y = pb.Next.Y
	} else if pb.Prev != nil && absFloat64(pb.Prev.Y-y) < 5 {
		pb.Y = pb.Prev.Y
	} else {
		pb.Y = y
	}
}

// DrawRect вычисляет прямоугольник области отрисовки точки
// Возвращает координаты (x1, y1, x2, y2)
func (ep *ElementPoint) DrawRect() (float64, float64, float64, float64) {
	var p *PointPos
	if ep.Type == PT_EVENT || ep.Type == PT_DATA || ep.Point == nil {
		p = ep.Pos
	} else {
		p = ep.Point.Pos
	}

	x1, y1 := p.X, p.Y
	x2, y2 := p.X, p.Y

	curr := p
	for curr.Next != nil {
		curr = curr.Next
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
	}

	return x1 - POINT_OFF, y1 - POINT_OFF, x2 - x1 + POINT_SPACE, y2 - y1 + POINT_SPACE
}

// SerializePath сериализует данные пути в строку формата: (x1,y1),(x2,y2),...,(xN,yN)
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

	for {
		p := result
		if p.Parent != nil {
			result = p.Parent.GetRealPoint(p)
		} else {
			result = nil
		}

		if p == result || result == nil || root == result {
			break
		}
	}

	if root == result {
		// Возвращаем последнюю успешную точку
		// В упрощенной версии возвращаем себя
		return ep
	}
	return result
}

// OnEvent обрабатывает событие
func (ep *ElementPoint) OnEvent(data *TData) {
	if ep.Point != nil {
		e := ep.Point.Parent
		if e != nil && e.IsCore() {
			// В оригинале: dynamic_cast<ElementCore*>(e)->do_work(point, data)
			// Заглушка для вызова метода do_work
			fmt.Printf("OnEvent вызван для элемента %s\n", e.Name)
		}
	}
}

// GetData получает данные
func (ep *ElementPoint) GetData(data *TData) {
	if ep.Point != nil {
		e := ep.Point.Parent
		if e != nil && e.IsCore() {
			// В оригинале: dynamic_cast<ElementCore*>(e)->read_var(point, data)
			// Заглушка для вызова метода read_var
			fmt.Printf("GetData вызван для элемента %s\n", e.Name)
		}
	}
}

// GetPair получает тип парной точки для подключения
func (ep *ElementPoint) GetPair() int {
	if ep.IsPrimary() {
		return ep.Type - 1
	}
	return ep.Type + 1
}

// GetDirection получает направление точки
// Возвращает 1 для work-event связи и 0 для var-data
func (ep *ElementPoint) GetDirection() int {
	if ep.Type == PT_EVENT || ep.Type == PT_WORK {
		return 1
	}
	return 0
}

// Вспомогательная функция для абсолютного значения float64
func absFloat64(x float64) float64 {
	if x < 0 {
		return -x
	}
	return x
}
