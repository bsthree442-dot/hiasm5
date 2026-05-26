package main

import "fmt"

// DataType определяет тип данных в системе
type DataType int

const (
	DataNull DataType = iota
	DataInt
	DataReal
	DataStr
	DataPixbuf
	DataArray
)

// DataArrayInterface - интерфейс для работы с массивами данных
type DataArrayInterface interface {
	ArrCount() int
	ArrSet(index, data *TData)
	ArrGet(index *TData) *TData
	ArrAdd(data *TData)
}

// TData представляет собой контейнер для передачи данных между элементами
type TData struct {
	Type  DataType
	IData int
	RData float64
	SData string
	Data  interface{} // для object, pixbuf, array
	Next  *TData      // следующая ссылка в MT-цепочке
}

// NewTDataNull создает пустые данные
func NewTDataNull() *TData {
	return &TData{Type: DataNull}
}

// NewTDataInt создает целочисленные данные
func NewTDataInt(value int) *TData {
	return &TData{Type: DataInt, IData: value}
}

// NewTDataReal создает вещественные данные
func NewTDataReal(value float64) *TData {
	return &TData{Type: DataReal, RData: value}
}

// NewTDataStr создает строковые данные
func NewTDataStr(value string) *TData {
	return &TData{Type: DataStr, SData: value}
}

// NewTDataPixbuf создает данные pixbuf
func NewTDataPixbuf(value interface{}) *TData {
	return &TData{Type: DataPixbuf, Data: value}
}

// NewTDataArray создает данные массива
func NewTDataArray(value DataArrayInterface) *TData {
	return &TData{Type: DataArray, Data: value}
}

// NewTDataCopy создает копию данных
func NewTDataCopy(source *TData) *TData {
	if source == nil {
		return NewTDataNull()
	}
	d := &TData{Type: source.Type, IData: source.IData, RData: source.RData, SData: source.SData, Data: source.Data}
	if source.Next != nil {
		d.Next = NewTDataCopy(source.Next)
	}
	return d
}

// Destroy освобождает ресурсы (в Go это делается сборщиком мусора, но оставляем для совместимости)
func (d *TData) Destroy() {
	d.Clear()
}

// SetString устанавливает строковое значение
func (d *TData) SetString(text string) *TData {
	d.Clear()
	d.Type = DataStr
	d.SData = text
	return d
}

// SetInt устанавливает целочисленное значение
func (d *TData) SetInt(value int) *TData {
	d.Clear()
	d.Type = DataInt
	d.IData = value
	return d
}

// SetReal устанавливает вещественное значение
func (d *TData) SetReal(value float64) *TData {
	d.Clear()
	d.Type = DataReal
	d.RData = value
	return d
}

// SetPixbuf устанавливает значение pixbuf
func (d *TData) SetPixbuf(value interface{}) *TData {
	d.Clear()
	d.Type = DataPixbuf
	d.Data = value
	return d
}

// SetArray устанавливает значение массива
func (d *TData) SetArray(value DataArrayInterface) *TData {
	d.Clear()
	d.Type = DataArray
	d.Data = value
	return d
}

// Assign копирует данные из другого TData
func (d *TData) Assign(src *TData) *TData {
	if src == nil {
		d.Clear()
		return d
	}
	d.Clear()
	d.Type = src.Type
	d.IData = src.IData
	d.RData = src.RData
	d.SData = src.SData
	d.Data = src.Data
	return d
}

// AssignMT копирует данные со всей MT-цепочкой
func (d *TData) AssignMT(src *TData) *TData {
	d.Assign(src)
	if src == nil {
		return d
	}
	current := d
	srcCurrent := src.Next
	for srcCurrent != nil {
		current.Next = NewTDataCopy(srcCurrent)
		current = current.Next
		srcCurrent = srcCurrent.Next
	}
	return d
}

// Equal сравнивает два TData
func (d *TData) Equal(other *TData) bool {
	if d.Type != other.Type {
		return false
	}
	switch d.Type {
	case DataNull:
		return true
	case DataInt:
		return d.IData == other.IData
	case DataReal:
		return d.RData == other.RData
	case DataStr:
		return d.SData == other.SData
	default:
		return false
	}
}

// Compare сравнивает данные с преобразованием типов
func (d *TData) Compare(other *TData) bool {
	switch d.Type {
	case DataNull:
		return other.Empty()
	case DataInt:
		return d.IData == other.ToInt()
	case DataReal:
		return d.RData == other.ToReal()
	case DataStr:
		return d.SData == other.ToStr()
	default:
		return false
	}
}

// Empty проверяет, пусто ли значение
func (d *TData) Empty() bool {
	switch d.Type {
	case DataInt:
		return d.IData == 0
	case DataReal:
		return d.RData == 0.0
	case DataStr:
		return d.SData == ""
	default:
		return true
	}
}

// Clear очищает данные
func (d *TData) Clear() {
	d.Type = DataNull
	d.IData = 0
	d.RData = 0.0
	d.SData = ""
	d.Data = nil
	d.Next = nil
}

// IsNull проверяет, является ли тип null
func (d *TData) IsNull() bool {
	return d.Type == DataNull
}

// ToStr преобразует данные в строку
func (d *TData) ToStr() string {
	switch d.Type {
	case DataInt:
		return intToStr(d.IData)
	case DataStr:
		return d.SData
	case DataReal:
		return floatToStr(d.RData)
	default:
		return ""
	}
}

// ToInt преобразует данные в целое число
func (d *TData) ToInt() int {
	switch d.Type {
	case DataInt:
		return d.IData
	case DataStr:
		return strToInt(d.SData)
	case DataReal:
		return int(d.RData)
	default:
		return 0
	}
}

// ToReal преобразует данные в вещественное число
func (d *TData) ToReal() float64 {
	switch d.Type {
	case DataInt:
		return float64(d.IData)
	case DataStr:
		return strToFloat(d.SData)
	case DataReal:
		return d.RData
	default:
		return 0.0
	}
}

// ToObj возвращает объект
func (d *TData) ToObj() interface{} {
	return d.Data
}

// ToPixbuf возвращает pixbuf
func (d *TData) ToPixbuf() interface{} {
	if d.Type == DataPixbuf {
		return d.Data
	}
	return nil
}

// ToArray возвращает массив
func (d *TData) ToArray() DataArrayInterface {
	if d.Type == DataArray {
		if arr, ok := d.Data.(DataArrayInterface); ok {
			return arr
		}
	}
	return nil
}

// Push добавляет MT-кольцо к данным
func (d *TData) Push(value *TData) *TData {
	if d.IsNull() {
		d.Assign(value)
	} else {
		n := &d.Next
		for *n != nil {
			n = &(*n).Next
		}
		*n = NewTDataCopy(value)
	}
	return d
}

// Shift отдает следующее MT-кольцо и удаляет его из цепочки
func (d *TData) Shift() *TData {
	if d.Next != nil {
		d.Type = d.Next.Type
		d.IData = d.Next.IData
		d.RData = d.Next.RData
		d.SData = d.Next.SData
		d.Data = d.Next.Data
		next := d.Next.Next
		d.Next.Next = nil
		d.Next = next
	} else {
		d.Clear()
	}
	return d
}

// Pop возвращает MT-кольцо
func (d *TData) Pop() *TData {
	result := NewTDataCopy(d)
	d.Shift()
	return result
}

// Вспомогательные функции для преобразования типов
func intToStr(i int) string {
	return fmt.Sprintf("%d", i)
}

func floatToStr(f float64) string {
	return fmt.Sprintf("%g", f)
}

func strToInt(s string) int {
	var i int
	fmt.Sscanf(s, "%d", &i)
	return i
}

func strToFloat(s string) float64 {
	var f float64
	fmt.Sscanf(s, "%g", &f)
	return f
}
