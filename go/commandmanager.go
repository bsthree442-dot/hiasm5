package main

import "fmt"

// Константы для имен меню
const (
	MENUNAME_ELEMENT  = "el_menu"
	MENUNAME_SDK      = "sdk_menu"
	MENUNAME_LINK     = "line_menu"
	MENUNAME_MAIN     = "menu"
	MENUNAME_TOOLBAR  = "main"
	MENUNAME_PAL_TAB  = "pal_tab_menu"
)

// CMD_TYPE - тип для имени команды
type CMD_TYPE string

// CommandItem представляет описание команды для CommandManager
type CommandItem struct {
	Name     string        // Имя команды (описано в hiasm.db)
	Info     string        // Информация о команде (описано в hiasm.db)
	Icon     interface{}   // Иконка команды (содержится в %hiasm%/int/icons folder)
	Enabled  bool          // Доступна ли команда сейчас
	Checked  bool          // Отмечена ли команда сейчас
	Widgets  []interface{} // Список виджетов, которые вызывают эту команду
}

// CommandManager управляет командами интерфейса и элементами, которые их вызывают
type CommandManager struct {
	Cmds         []*CommandItem // Массив команд окружения
	Widgets      []interface{}  // Массив всех созданных виджетов, содержащих команды интерфейса
	OnCommand    *Event         // Вызывается, когда пользователь запускает команду
	OnCmdEnabled *Event         // Вызывается, когда команды должны обновить свой статус
}

// Event представляет событие
type Event struct {
	Source interface{}
	Type   int
}

// Константы типов событий
const (
	CBT_RUN_COMMAND   = 1
	CBT_CMD_ENABLED   = 2
)

// NewEvent создает новое событие
func NewEvent(source interface{}, eventType int) *Event {
	return &Event{
		Source: source,
		Type:   eventType,
	}
}

// Run запускает событие
func (e *Event) Run(data interface{}) {
	// В оригинале здесь вызывались подписчики события
	// Заглушка: просто логируем
	fmt.Printf("Событие типа %d запущено\n", e.Type)
}

// Global экземпляр CommandManager
var cmdMan *CommandManager

// NewCommandManager создает новый CommandManager
func NewCommandManager() *CommandManager {
	cm := &CommandManager{
		OnCommand:    NewEvent(nil, CBT_RUN_COMMAND),
		OnCmdEnabled: NewEvent(nil, CBT_CMD_ENABLED),
	}
	cm.loadCommandList()
	return cm
}

// loadIcon загружает иконку из директории int
func (cm *CommandManager) loadIcon(name string) interface{} {
	// В оригинале: ustring icon_file = dataDir + ICONS_PATH + name + ".png";
	// if(Glib::file_test(icon_file, Glib::FILE_TEST_EXISTS))
	//     return Gdk::Pixbuf::create_from_file(icon_file);
	
	// Заглушка: возвращаем nil или имя иконки
	iconFile := "/data/int/icons/" + name + ".png"
	fmt.Printf("Попытка загрузки иконки: %s\n", iconFile)
	
	// В реальной реализации здесь была бы проверка существования файла
	// и загрузка изображения
	return nil
}

// loadCommandList загружает список команд из базы данных
func (cm *CommandManager) loadCommandList() {
	// В оригинале здесь было чтение из SQLite через mdb.begin_read_commands()
	// Заглушка: создаем несколько тестовых команд
	
	commands := []struct {
		name string
		info string
	}{
		{"CMD_NEW", "Создать новый проект"},
		{"CMD_OPEN", "Открыть проект"},
		{"CMD_SAVE", "Сохранить проект"},
		{"CMD_BUILD", "Скомпилировать проект"},
		{"CMD_RUN", "Запустить проект"},
		{"CMD_ZOOMIN", "Увеличить масштаб"},
		{"CMD_ZOOMOUT", "Уменьшить масштаб"},
		{"CMD_SELECTALL", "Выделить все"},
		{"CMD_SEARCH", "Поиск"},
		{"CMD_FULLSCREEN", "Полноэкранный режим"},
	}
	
	for _, cmd := range commands {
		ci := &CommandItem{
			Name:     cmd.name,
			Info:     cmd.info,
			Icon:     cm.loadIcon(cmd.name),
			Enabled:  false,
			Checked:  false,
			Widgets:  make([]interface{}, 0),
		}
		cm.Cmds = append(cm.Cmds, ci)
	}
	
	fmt.Println("Список команд загружен")
}

// FindByName ищет описание команды по её имени
func (cm *CommandManager) FindByName(name string) *CommandItem {
	for _, c := range cm.Cmds {
		if c.Name == name {
			return c
		}
	}
	return nil
}

// createImageByCmd создает изображение для команды
func createImageByCmd(ci *CommandItem) interface{} {
	// В оригинале здесь были проверки на стандартные иконки GTK
	// Заглушка: возвращаем nil
	switch ci.Name {
	case "CMD_ZOOMIN":
		fmt.Println("Создание иконки ZOOM_IN")
	case "CMD_ZOOMOUT":
		fmt.Println("Создание иконки ZOOM_OUT")
	case "CMD_SELECTALL":
		fmt.Println("Создание иконки SELECT_ALL")
	case "CMD_SEARCH":
		fmt.Println("Создание иконки FIND")
	case "CMD_FULLSCREEN":
		fmt.Println("Создание иконки FULLSCREEN")
	}
	
	if ci.Icon != nil {
		return ci.Icon
	}
	
	return nil
}

// CreateToolbar создает виджет Toolbar по имени
func (cm *CommandManager) CreateToolbar(name string) interface{} {
	fmt.Printf("Создание Toolbar: %s\n", name)
	
	// В оригинале здесь было чтение меню команд из базы данных
	// и создание кнопок
	
	// Заглушка: просто добавляем toolbar в список виджетов
	cm.Widgets = append(cm.Widgets, name)
	
	return name
}

// CreateMenuBar создает виджет MenuBar по имени
func (cm *CommandManager) CreateMenuBar(name string) interface{} {
	fmt.Printf("Создание MenuBar: %s\n", name)
	return cm.fillMenu(name)
}

// CreateMenu создает виджет Menu по имени
func (cm *CommandManager) CreateMenu(name string) interface{} {
	fmt.Printf("Создание Menu: %s\n", name)
	return cm.fillMenu(name)
}

// fillMenu заполняет меню командами
func (cm *CommandManager) fillMenu(name string) interface{} {
	fmt.Printf("Заполнение меню: %s\n", name)
	
	// В оригинале здесь было чтение command_menu из базы данных
	// и создание элементов меню
	
	// Заглушка: просто возвращаем имя меню
	cm.Widgets = append(cm.Widgets, name)
	
	return name
}

// Command вызывает пользовательскую команду
func (cm *CommandManager) Command(name string) {
	fmt.Printf("Вызов команды: %s\n", name)
	cm.OnCommand.Run(name)
}

// BeginUpdate инициирует обновление статуса команд
func (cm *CommandManager) BeginUpdate() {
	fmt.Println("Начало обновления статуса команд")
	
	for _, c := range cm.Cmds {
		c.Checked = false
		c.Enabled = false
	}
	
	cm.OnCmdEnabled.Run(cm)
}

// Enable определяет доступность команды для выполнения
func (cm *CommandManager) Enable(cmd CMD_TYPE) bool {
	ci := cm.getCommandByName(string(cmd))
	if ci != nil {
		ci.Enabled = true
		return true
	}
	
	fmt.Printf("WARNING: команда %s не найдена!\n", cmd)
	return false
}

// IsEnabled проверяет, включена ли команда
func (cm *CommandManager) IsEnabled(cmd CMD_TYPE) bool {
	ci := cm.getCommandByName(string(cmd))
	if ci != nil {
		return ci.Enabled
	}
	
	fmt.Printf("WARNING: команда %s не найдена!\n", cmd)
	return false
}

// Check включает команду (устанавливает флаг checked)
func (cm *CommandManager) Check(cmd CMD_TYPE) bool {
	ci := cm.getCommandByName(string(cmd))
	if ci != nil {
		ci.Checked = true
		return true
	}
	
	fmt.Printf("WARNING: команда %s не найдена!\n", cmd)
	return false
}

// EndUpdate завершает обновление статуса команд
func (cm *CommandManager) EndUpdate() {
	fmt.Println("Завершение обновления статуса команд")
	
	for _, c := range cm.Cmds {
		for _, w := range c.Widgets {
			// В оригинале: (*w)->set_sensitive((*c)->enabled);
			fmt.Printf("Обновление виджета: чувствительность=%v, виджет=%v\n", c.Enabled, w)
			
			// Проверка на ToggleToolButton
			// CM_ToggleToolButton *b = dynamic_cast<CM_ToggleToolButton*>(*w);
			// if(b)
			//     b->check((*c)->checked);
		}
	}
}

// AttachMenu прикрепляет popup меню к команде
func (cm *CommandManager) AttachMenu(name CMD_TYPE, menu interface{}) {
	cmd := cm.getCommandByName(string(name))
	if cmd != nil {
		// Только одно прикрепление...
		for _, w := range cmd.Widgets {
			// Попытка привести к CM_MenuToolButton
			// CM_MenuToolButton *b = dynamic_cast<CM_MenuToolButton*>(*w);
			// if(b) {
			//     b->set_menu(menu);
			//     return;
			// }
			
			// Попытка привести к ImageMenuItem
			// ImageMenuItem *m = dynamic_cast<ImageMenuItem*>(*w);
			// if(m) {
			//     m->set_submenu(menu);
			//     return;
			// }
			
			_ = w
		}
	}
}

// getCommandByName получает команду по имени
func (cm *CommandManager) getCommandByName(codeName string) *CommandItem {
	for _, c := range cm.Cmds {
		if c.Name == codeName {
			return c
		}
	}
	return nil
}

// ============================================================================
// Специализированные классы кнопок и элементов меню
// ============================================================================

// CM_ToolButton - специальная кнопка для CommandManager
type CM_ToolButton struct {
	Cmd    string
	Parent *CommandManager
	Label  string
}

// NewCM_ToolButton создает новую кнопку инструмента
func NewCM_ToolButton(parent *CommandManager, name, info string) *CM_ToolButton {
	return &CM_ToolButton{
		Cmd:    name,
		Parent: parent,
		Label:  info,
	}
}

// OnClicked обрабатывает клик по кнопке
func (b *CM_ToolButton) OnClicked() {
	b.Parent.Command(b.Cmd)
}

// CM_ToggleToolButton - специальная переключаемая кнопка для CommandManager
type CM_ToggleToolButton struct {
	Cmd     string
	Parent  *CommandManager
	Label   string
	Change  bool
	Active  bool
}

// NewCM_ToggleToolButton создает новую переключаемую кнопку инструмента
func NewCM_ToggleToolButton(parent *CommandManager, name, info string) *CM_ToggleToolButton {
	return &CM_ToggleToolButton{
		Cmd:    name,
		Parent: parent,
		Label:  info,
		Change: false,
		Active: false,
	}
}

// OnClicked обрабатывает клик по переключаемой кнопке
func (b *CM_ToggleToolButton) OnClicked() {
	if !b.Change {
		b.Parent.Command(b.Cmd)
	}
}

// Check устанавливает состояние кнопки
func (b *CM_ToggleToolButton) Check(value bool) {
	b.Change = true
	b.Active = value
	b.Change = false
}

// CM_MenuToolButton - специальная кнопка с меню для CommandManager
type CM_MenuToolButton struct {
	Cmd    string
	Parent *CommandManager
	Label  string
	Menu   interface{}
}

// NewCM_MenuToolButton создает новую кнопку с меню
func NewCM_MenuToolButton(parent *CommandManager, name, info string) *CM_MenuToolButton {
	return &CM_MenuToolButton{
		Cmd:    name,
		Parent: parent,
		Label:  info,
		Menu:   nil,
	}
}

// OnClicked обрабатывает клик по кнопке с меню
func (b *CM_MenuToolButton) OnClicked() {
	b.Parent.Command(b.Cmd)
}

// SetMenu устанавливает меню для кнопки
func (b *CM_MenuToolButton) SetMenu(menu interface{}) {
	b.Menu = menu
}

// CM_ImageMenuItem - специальный элемент меню с изображением для CommandManager
type CM_ImageMenuItem struct {
	Cmd    string
	Parent *CommandManager
	Label  string
	Image  interface{}
}

// NewCM_ImageMenuItem создает новый элемент меню с изображением
func NewCM_ImageMenuItem(parent *CommandManager, name, info string) *CM_ImageMenuItem {
	return &CM_ImageMenuItem{
		Cmd:    name,
		Parent: parent,
		Label:  info,
		Image:  nil,
	}
}

// OnActivate обрабатывает активацию элемента меню
func (m *CM_ImageMenuItem) OnActivate() {
	m.Parent.Command(m.Cmd)
}

// SetImage устанавливает изображение для элемента меню
func (m *CM_ImageMenuItem) SetImage(image interface{}) {
	m.Image = image
}

// SetSubmenu устанавливает подменю для элемента меню
func (m *CM_ImageMenuItem) SetSubmenu(menu interface{}) {
	// В оригинале здесь было бы установление подменю
	fmt.Println("Установка подменю для элемента меню")
}
