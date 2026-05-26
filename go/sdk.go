package main

import (
	"fmt"
	"sync"
)

// SDKObjectType определяет типы объектов в SDK
type SDKObjectType int

const (
	SDKObjectElement SDKObjectType = iota
	SDKObjectProperty
	SDKObjectPoint
)

// ElementDescription описывает элемент SDK
type ElementDescription struct {
	ID          string
	Name        string
	Category    string
	Description string
	IconID      int
	Properties  []PropertyDescription
	Points      []PointDescription
}

// PropertyDescription описывает свойство элемента
type PropertyDescription struct {
	Name    string
	Type    int
	Default interface{}
}

// PointDescription описывает точку соединения
type PointDescription struct {
	Name string
	Type int // Вход/Выход
}

// SDK управляет набором доступных элементов
type SDK struct {
	mu       sync.RWMutex
	elements map[string]*ElementDescription
	loaded   bool
	version  string
}

// NewSDK создает новый экземпляр SDK
func NewSDK() *SDK {
	return &SDK{
		elements: make(map[string]*ElementDescription),
		loaded:   false,
		version:  "5.0",
	}
}

// Load загружает описания элементов (заглушка без файловой системы/SQLite)
func (s *SDK) Load(path string) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	fmt.Printf("[SDK] Загрузка элементов из %s (заглушка)\n", path)
	
	// Имитация загрузки нескольких тестовых элементов
	s.elements["Start"] = &ElementDescription{
		ID:       "Start",
		Name:     "Старт",
		Category: "System",
		Properties: []PropertyDescription{
			{Name: "Caption", Type: 1, Default: "Start"},
		},
	}
	
	s.elements["End"] = &ElementDescription{
		ID:       "End",
		Name:     "Конец",
		Category: "System",
	}

	s.loaded = true
	return nil
}

// GetElement возвращает описание элемента по ID
func (s *SDK) GetElement(id string) *ElementDescription {
	s.mu.RLock()
	defer s.mu.RUnlock()
	
	if el, ok := s.elements[id]; ok {
		return el
	}
	return nil
}

// GetElementByName ищет элемент по имени
func (s *SDK) GetElementByName(name string) *ElementDescription {
	s.mu.RLock()
	defer s.mu.RUnlock()
	
	for _, el := range s.elements {
		if el.Name == name {
			return el
		}
	}
	return nil
}

// Count возвращает количество загруженных элементов
func (s *SDK) Count() int {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return len(s.elements)
}

// IsLoaded проверяет, загружен ли SDK
func (s *SDK) IsLoaded() bool {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.loaded
}

// GetVersion возвращает версию SDK
func (s *SDK) GetVersion() string {
	return s.version
}

// RegisterElement регистрирует новый элемент (для расширения)
func (s *SDK) RegisterElement(desc *ElementDescription) {
	s.mu.Lock()
	defer s.mu.Unlock()
	if desc != nil && desc.ID != "" {
		s.elements[desc.ID] = desc
	}
}

// GetAllElements возвращает список всех элементов
func (s *SDK) GetAllElements() []*ElementDescription {
	s.mu.RLock()
	defer s.mu.RUnlock()
	
	list := make([]*ElementDescription, 0, len(s.elements))
	for _, el := range s.elements {
		list = append(list, el)
	}
	return list
}

// Заглушки для GUI-методов
func (s *SDK) LoadIcons(path string) error {
	fmt.Println("[SDK] Загрузка иконок отключена (нет GUI)")
	return nil
}

func (s *SDK) RefreshUI() {
	fmt.Println("[SDK] Обновление UI отключено (нет GUI)")
}
