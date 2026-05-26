package main

import (
	"encoding/json"
	"fmt"
	"os"
)

// Project представляет проект HiAsm
type Project struct {
	ID          int
	Name        string
	Description string
	Elements    []*Element
	Compilers   *CompilerCollection
	FilePath    string
	Modified    bool
	Version     string
}

// NewProject создает новый проект
func NewProject(name string) *Project {
	return &Project{
		ID:          0,
		Name:        name,
		Description: "",
		Elements:    make([]*Element, 0),
		Compilers:   NewCompilerCollection(),
		FilePath:    "",
		Modified:    false,
		Version:     "5.0",
	}
}

// AddElement добавляет элемент в проект
func (p *Project) AddElement(element *Element) {
	if element == nil {
		return
	}
	p.Elements = append(p.Elements, element)
	p.Modified = true
}

// RemoveElement удаляет элемент из проекта по ID
func (p *Project) RemoveElement(id int) bool {
	for i, elem := range p.Elements {
		if elem.ID == id {
			p.Elements = append(p.Elements[:i], p.Elements[i+1:]...)
			p.Modified = true
			return true
		}
	}
	return false
}

// GetElement возвращает элемент по ID
func (p *Project) GetElement(id int) *Element {
	for _, elem := range p.Elements {
		if elem.ID == id {
			return elem
		}
	}
	return nil
}

// GetElementCount возвращает количество элементов
func (p *Project) GetElementCount() int {
	return len(p.Elements)
}

// Clear очищает проект
func (p *Project) Clear() {
	p.Elements = make([]*Element, 0)
	p.Modified = true
}

// SetModified устанавливает флаг изменений
func (p *Project) SetModified(modified bool) {
	p.Modified = modified
}

// IsModified возвращает флаг изменений
func (p *Project) IsModified() bool {
	return p.Modified
}

// Save сохраняет проект в файл JSON
func (p *Project) Save(filePath string) error {
	data, err := json.MarshalIndent(p, "", "  ")
	if err != nil {
		return err
	}

	err = os.WriteFile(filePath, data, 0644)
	if err != nil {
		return err
	}

	p.FilePath = filePath
	p.Modified = false
	return nil
}

// Load загружает проект из файла JSON
func (p *Project) Load(filePath string) error {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return err
	}

	// Временная структура для размаршаливания
	var temp struct {
		ID          int      `json:"ID"`
		Name        string   `json:"Name"`
		Description string   `json:"Description"`
		Elements    []struct {
			ID       int                    `json:"ID"`
			Name     string                 `json:"Name"`
			Type     string                 `json:"Type"`
			X        int                    `json:"X"`
			Y        int                    `json:"Y"`
			Width    int                    `json:"Width"`
			Height   int                    `json:"Height"`
			Props    map[string]interface{} `json:"props"`
			Points   []Point                `json:"points"`
			Events   []string               `json:"events"`
			ParentID int                    `json:"parentID"`
			Visible  bool                   `json:"visible"`
			Locked   bool                   `json:"locked"`
		} `json:"Elements"`
		Compilers struct {
			Compilers []struct {
				ID   int    `json:"ID"`
				Name string `json:"Name"`
				Cmd  string `json:"Cmd"`
				Path string `json:"Path"`
				Ext  string `json:"Ext"`
			} `json:"Compilers"`
		} `json:"Compilers"`
		FilePath string `json:"FilePath"`
		Modified bool   `json:"Modified"`
		Version  string `json:"Version"`
	}

	err = json.Unmarshal(data, &temp)
	if err != nil {
		return err
	}

	p.ID = temp.ID
	p.Name = temp.Name
	p.Description = temp.Description
	p.FilePath = temp.FilePath
	p.Modified = temp.Modified
	p.Version = temp.Version

	// Восстанавливаем элементы
	p.Elements = make([]*Element, 0)
	for _, e := range temp.Elements {
		elem := NewElement(e.ID, e.Name, e.Type)
		elem.X = e.X
		elem.Y = e.Y
		elem.Width = e.Width
		elem.Height = e.Height
		elem.props = e.Props
		elem.points = e.Points
		elem.events = e.Events
		elem.parentID = e.ParentID
		elem.visible = e.Visible
		elem.locked = e.Locked
		p.Elements = append(p.Elements, elem)
	}

	// Восстанавливаем компиляторы
	p.Compilers = NewCompilerCollection()
	for _, c := range temp.Compilers.Compilers {
		compiler := &Compiler{
			ID:   c.ID,
			Name: c.Name,
			Cmd:  c.Cmd,
			Path: c.Path,
			Ext:  c.Ext,
		}
		p.Compilers.Add(compiler)
	}

	return nil
}

// GetElementsByType возвращает все элементы указанного типа
func (p *Project) GetElementsByType(elemType string) []*Element {
	result := make([]*Element, 0)
	for _, elem := range p.Elements {
		if elem.Type == elemType {
			result = append(result, elem)
		}
	}
	return result
}

// GetAllElementIDs возвращает IDs всех элементов
func (p *Project) GetAllElementIDs() []int {
	ids := make([]int, 0, len(p.Elements))
	for _, elem := range p.Elements {
		ids = append(ids, elem.ID)
	}
	return ids
}

// PrintInfo выводит информацию о проекте
func (p *Project) PrintInfo() {
	fmt.Printf("Project: %s (v%s)\n", p.Name, p.Version)
	fmt.Printf("ID: %d\n", p.ID)
	fmt.Printf("Description: %s\n", p.Description)
	fmt.Printf("Elements count: %d\n", len(p.Elements))
	fmt.Printf("Modified: %v\n", p.Modified)
	if p.FilePath != "" {
		fmt.Printf("File: %s\n", p.FilePath)
	}
}
