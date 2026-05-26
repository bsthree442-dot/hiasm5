package main

import (
	"fmt"
	"strconv"
)

// DataType представляет тип данных
type DataType int

// Константы типов данных
const (
	DATA_NULL = iota
	DATA_INT
	DATA_REAL
	DATA_STR
)

// Вспомогательные функции для преобразования типов

// IntToStr преобразует целое число в строку
func IntToStr(value int) string {
	return strconv.Itoa(value)
}

// FloatToStr преобразует вещественное число в строку
func FloatToStr(value float64) string {
	return strconv.FormatFloat(value, 'f', -1, 64)
}

// StrToInt преобразует строку в целое число
func StrToInt(value string) int {
	v, err := strconv.Atoi(value)
	if err != nil {
		return 0
	}
	return v
}

// StrToFloat преобразует строку в вещественное число
func StrToFloat(value string) float64 {
	v, err := strconv.ParseFloat(value, 64)
	if err != nil {
		return 0.0
	}
	return v
}

// DebugMsg выводит отладочное сообщение (заглушка)
func DebugMsg(msg interface{}) {
	fmt.Println(msg)
}

// DataArrayInterface - интерфейс для доступа к массивам
type DataArrayInterface interface {
	ArrCount() int
	ArrSet(index, data *TData)
	ArrGet(index *TData) *TData
	ArrAdd(data *TData)
}

// TData - внутренняя структура передачи данных
// Предоставляет универсальный контейнер для передачи данных между элементами
type TData struct {
	Type  DataType      // тип данных
	IData int           // целочисленные данные
	RData float64       // вещественные данные
	SData string        // строковые данные
	Data  interface{}   // объектные данные (pixbuf, array и т.д.)
	Next  *TData        // указатель на следующее кольцо в MT данных
}

// NewTData создает пустые данные (null)
func NewTData() *TData {
	return &TData{Type: DATA_NULL}
}

// NewTDataInt создает целочисленные данные
func NewTDataInt(value int) *TData {
	return &TData{Type: DATA_INT, IData: value}
}

// NewTDataReal создает вещественные данные
func NewTDataReal(value float64) *TData {
	return &TData{Type: DATA_REAL, RData: value}
}

// NewTDataStr создает строковые данные
func NewTDataStr(value string) *TData {
	return &TData{Type: DATA_STR, SData: value}
}

// NewTDataCopy создает копию данных
func NewTDataCopy(source *TData) *TData {
	d := &TData{}
	if source != nil {
		d.Assign(source)
	} else {
		d.Type = DATA_NULL
	}
	return d
}

// Destroy уничтожает TData и все MT кольца
func (d *TData) Destroy() {
	if d.Next != nil {
		d.Next.Destroy()
		d.Next = nil
	}
}

// Assign копирует данные из источника
func (d *TData) Assign(src *TData) *TData {
	d.Clear()
	d.Type = src.Type
	d.RData = src.RData
	d.SData = src.SData
	d.Data = src.Data
	return d
}

// AssignMt копирует данные со всеми MT кольцами
func (d *TData) AssignMt(src *TData) *TData {
	d.Assign(src)
	dstNext := &d.Next
	srcNext := src.Next
	for srcNext != nil {
		*dstNext = NewTDataCopy(srcNext)
		dstNext = &(*dstNext).Next
		srcNext = srcNext.Next
	}
	return d
}

// Clear очищает данные (устанавливает тип в DATA_NULL)
func (d *TData) Clear() {
	d.Type = DATA_NULL
	if d.Next != nil {
		d.Next.Destroy()
		d.Next = nil
	}
}

// Empty проверяет, пусто ли значение данных
func (d *TData) Empty() bool {
	switch d.Type {
	case DATA_INT:
		return d.IData == 0
	case DATA_REAL:
		return d.RData == 0.0
	case DATA_STR:
		return d.SData == ""
	default:
		return true
	}
}

// IsNull проверяет, является ли тип данных null
func (d *TData) IsNull() bool {
	return d.Type == DATA_NULL
}

// ToStr преобразует данные в строку
func (d *TData) ToStr() string {
	switch d.Type {
	case DATA_INT:
		return IntToStr(d.IData)
	case DATA_STR:
		return d.SData
	case DATA_REAL:
		return FloatToStr(d.RData)
	default:
		return ""
	}
}

// ToInt преобразует данные в целое число
func (d *TData) ToInt() int {
	switch d.Type {
	case DATA_INT:
		return d.IData
	case DATA_STR:
		return StrToInt(d.SData)
	case DATA_REAL:
		return int(d.RData)
	default:
		return 0
	}
}

// ToReal преобразует данные в вещественное число
func (d *TData) ToReal() float64 {
	switch d.Type {
	case DATA_INT:
		return float64(d.IData)
	case DATA_STR:
		return StrToFloat(d.SData)
	case DATA_REAL:
		return d.RData
	default:
		return 0.0
	}
}

// ToObj возвращает объектные данные
func (d *TData) ToObj() interface{} {
	return d.Data
}

// ToPixbuf возвращает pixbuf данные (заглушка)
func (d *TData) ToPixbuf() interface{} {
	// Заглушка: тип DataPixbuf удален
	return nil
}

// ToArray возвращает массив данных
func (d *TData) ToArray() DataArrayInterface {
	// Заглушка: тип DataArray удален
	return nil
}

// Compare сравнивает два TData с преобразованием типов
func (d *TData) Compare(other *TData) bool {
	switch d.Type {
	case DATA_NULL:
		return other.Empty()
	case DATA_INT:
		return d.IData == other.ToInt()
	case DATA_REAL:
		return d.RData == other.ToReal()
	case DATA_STR:
		return d.SData == other.ToStr()
	default:
		return false
	}
}

// Equal сравнивает два TData без преобразования типов
func (d *TData) Equal(other *TData) bool {
	if d.Type == other.Type {
		switch d.Type {
		case DATA_NULL:
			return other.Type == DATA_NULL
		case DATA_INT:
			return other.IData == d.IData
		case DATA_REAL:
			return other.RData == d.RData
		case DATA_STR:
			return other.SData == d.SData
		default:
			return false
		}
	}
	return false
}

// Push добавляет MT кольцо к данным
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

// Shift извлекает следующее MT кольцо и удаляет его из цепочки
func (d *TData) Shift() *TData {
	if d.Next != nil {
		d.Type = d.Next.Type
		d.RData = d.Next.RData
		d.SData = d.Next.SData
		d.Data = d.Next.Data
		n := d.Next.Next
		d.Next.Next = nil
		d.Next.Destroy()
		d.Next = n
	} else {
		d.Clear()
	}
	return d
}

// Pop извлекает MT кольцо и возвращает его
func (d *TData) Pop() *TData {
	ret := NewTDataCopy(d)
	d.Shift()
	return ret
}

// Dump выводит отладочную информацию (заглушка)
func (d *TData) Dump() {
	// Заглушка для отладки
}
