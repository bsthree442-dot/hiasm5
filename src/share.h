/*
 * share.h
 *
 *  Created on: 01.05.2010
 *      Author: dilma
 */

#ifndef SHARE_H_
#define SHARE_H_

#include <iostream>
#include <string>
#include <list>
#include <vector>
#include <map>
#include <memory>

// GTK stubs - placeholder types for GUI-less build
namespace Gtk {
    class Widget {};
    class Window : public Widget {};
    class DrawingArea : public Widget {};
    class VBox : public Widget {};
    class HBox : public Widget {};
    class VPaned : public Widget {};
    class HPaned : public Widget {};
    class Notebook : public Widget {};
    class Toolbar : public Widget {};
    class MenuBar : public Widget {};
    class Menu : public Widget {};
    class MenuItem : public Widget {};
    class ImageMenuItem : public MenuItem {};
    class ToolButton : public Widget {};
    class RadioToolButton : public ToolButton {};
    class ToggleToolButton : public ToolButton {};
    class SeparatorToolItem : public Widget {};
    class ScrolledWindow : public Widget {};
    class TextView : public Widget {};
    class TextBuffer {};
    class TreeView : public Widget {};
    class TreeModel {};
    class ListStore : public TreeModel {};
    class TreeModelColumnRecord {};
    class TreeModelColumn {};
    class TreeViewColumn {};
    class CellRenderer {};
    class CellRendererToggle : public CellRenderer {};
    class EventBox : public Widget {};
    class AboutDialog : public Widget {};
    class FileChooserDialog : public Widget {};
    class RadioButtonGroup {};
    class Image : public Widget {};
    class Label : public Widget {};
    class HScale : public Widget {};
    class StyleContext {};
    class Allocation {
    public:
        int get_width() const { return 0; }
        int get_height() const { return 0; }
    };
    enum PolicyType { POLICY_AUTOMATIC, POLICY_NEVER };
    enum Orientation { ORIENTATION_HORIZONTAL, ORIENTATION_VERTICAL };
    enum PackType { PACK_SHRINK, PACK_EXPAND_WIDGET };
    enum IconSize { ICON_SIZE_MENU };
    enum ToolbarStyle { TOOLBAR_ICONS };
    enum WrapMode { WRAP_WORD };
    enum StateFlags { STATE_FLAG_NORMAL };
    enum WindowPosition { WIN_POS_CENTER };
}

namespace Gdk {
    class Pixbuf {};
    struct RGBA {
        double red = 0, green = 0, blue = 0, alpha = 0;
        
        void set(const std::string& color) {
            // Stub implementation - just set default values
            red = 0.5; green = 0.5; blue = 0.5; alpha = 1.0;
        }
    };
    class Cursor {};
    class Display {};
    struct Rectangle {
        int x, y, width, height;
    };
}

namespace Pango {
    class Layout {};
}

namespace Glib {
    template<typename T> class RefPtr {
    private:
        T* ptr;
    public:
        RefPtr() : ptr(nullptr) {}
        T* operator->() const { return ptr; }
        T& operator*() const { return *ptr; }
        operator bool() const { return ptr != nullptr; }
        static RefPtr<T> create(T* p = nullptr) { RefPtr<T> r; r.ptr = p; return r; }
    };
    
    class Thread {
    public:
        static Thread* create(void (*func)(void*), void* data) { return nullptr; }
    };
}

// Global typedef for ustring (std::string)
typedef std::string ustring;

// Define gdouble as double for stub compatibility
typedef double gdouble;
typedef int gint;
typedef unsigned int guint;
typedef unsigned int guint32;

// File test flags
enum FileTest {
    FILE_TEST_EXISTS = 0,
    FILE_TEST_IS_REGULAR = 1,
    FILE_TEST_IS_DIR = 2,
    FILE_TEST_IS_SYMLINK = 3
};

// Declare file_test function
bool file_test(const std::string& filename, FileTest test);

namespace Cairo {
    class Context {};
    namespace RefPtr {
        template<typename T> class RefPtrT {};
    }
    enum ErrorStatus { STATUS_SUCCESS, STATUS_ERROR };
}

using namespace std;

// Bring GTK stubs into global namespace for compatibility
using Gtk::Widget;
using Gtk::StyleContext;
using Gtk::Entry;
using Gtk::TreeView;
using Gtk::EventBox;
using Gtk::Button;
using Gtk::ListStore;
using Gtk::Fixed;
using Gtk::ScrolledWindow;
using Gtk::Menu;
using Gtk::MenuItem;
using Gtk::ImageMenuItem;
using Gtk::Toolbar;
using Gtk::ToolButton;
using Gtk::Notebook;
using Gtk::VBox;
using Gtk::HBox;
using Gtk::VPaned;
using Gtk::HPaned;
using Gtk::TextView;
using Gtk::Label;
using Gtk::AboutDialog;
using Gtk::FileChooserDialog;
using Gtk::TreeModel;
using Gtk::TreeModelColumnRecord;
using Gtk::TreeModelColumn;
using Gtk::TreeViewColumn;
using Gtk::CellRenderer;
using Gtk::CellRendererToggle;
using Gtk::Image;
using Gtk::HScale;

#define HIASM_VERSION_MAJOR 5
#define HIASM_VERSION_MINOR 0
#define HIASM_VERSION_BUILD 12

// Define gchar as char for stub compatibility
typedef char gchar;
gchar *HIASM_VERSION();

#ifdef G_OS_WIN32
	#define LINE_END "\n"
	#define PATH_SLASH "\\"
#else
	#define LINE_END "\n"
	#define PATH_SLASH "/"
#endif

// point types ----------------------------------------------------------------------------------------

#define pt_work  1
#define pt_event 2
#define pt_var   3
#define pt_data  4

// data types -----------------------------------------------------------------------------------------
enum DataType {
	data_null = 0,
	data_int,
	data_str,
	data_data,
	data_combo,
	data_list,
	data_icon,
	data_real,
	data_color,
	data_script,
	data_stream,
	data_bitmap,
	data_wave,
	data_array,
	data_comboEx,
	data_font,
	data_matr,
	data_jpeg,
	data_menu,
	data_code,
	data_element,
	data_flags,
	data_stock,
	data_pixbuf,

	data_count
};

extern const char *dataNames[];

#ifndef max
  #define max(x,y) (((x) > (y)) ? (x) : (y))
  #define min(x,y) (((x) < (y)) ? (x) : (y))
#endif

// element types ----------------------------------------------------------------------------------
enum ElementType {
	CI_DPElement = 1,
	CI_MultiElement,
	CI_EditMulti,
	CI_EditMultiEx,
	CI_InlineCode,
	CI_DrawElement,
	CI_AS_Special,
	CI_DPLElement,
	CI_UseHiDLL,
	CI_WinElement,
	CI_PointHint,
};

#define str_to_int(s) atoi(s.c_str())
extern ustring int_to_str(int value);
extern ustring double_to_str(double value);

extern char *getTok(char **buf, char tok);

// path ----------------------------------------------------------------------------------------------
#define INT_PATH     "int"PATH_SLASH
#define ICONS_PATH     INT_PATH"icons"PATH_SLASH
#define CURSORS_PATH   INT_PATH"cur"PATH_SLASH
#define DATABASE_FILE  "hiasm.db"
#define LANG_PATH      INT_PATH"lang"PATH_SLASH
#define ELEMENTS_PATH  "elements"PATH_SLASH
#define ELEMENTS_ICON_PATH  "icon"PATH_SLASH
#define ELEMENTS_CODE_PATH  "code"PATH_SLASH
#define ELEMENTS_CONF_PATH  "conf"PATH_SLASH
#define ELEMENTS_NEW_PATH   "new"PATH_SLASH
#define ELEMENTS_MAKE_PATH  "make"PATH_SLASH
#define PACK_ICON_FILE  "icon.png"
#define SPLASH_LOGO_FILE  "splash.png"

#define TABS_STATE_PATH	"tabs"PATH_SLASH
#define APP_SETTINGS "settings.ini"

#define ELEMENTS_NIL_FILE ELEMENTS_PATH"_base"PATH_SLASH""ELEMENTS_ICON_PATH"nil.png"

extern ustring databaseFile;
extern ustring dataDir;
extern ustring homeDir;

// internal types --------------------------------------------------------------------------------------
typedef void* DrawContext;  // Stub for Cairo context
typedef void* TypePixbuf;   // Stub for Gdk::Pixbuf
struct TypeColor {          // Stub for Gdk::RGBA
    double red = 0, green = 0, blue = 0, alpha = 0;
    
    std::string to_string() const {
        return std::to_string(red) + "," + std::to_string(green) + "," + std::to_string(blue);
    }
    
    void set(const std::string& color) {
        // Stub implementation - just set default values
        red = 0.5; green = 0.5; blue = 0.5; alpha = 1.0;
    }
};

// macro -----------------------------------------------------------------------------------------------
#define TRACE_PROC //std::cout << "[" << __FILE__ << "]:" << __FUNCTION__ << std::endl;
#define DEBUG_MSG(v) {}// std::cout << v << std::endl;
#define ERROR_MSG(v) std::cout << v << std::endl;

// buttons ---------------------------------------------------------------------------------------------
#define BTN_LEFT   1
#define BTN_MIDDLE 2
#define BTN_RIGHT  3

#define KEY_SHIFT    1
#define KEY_CONTROL  4
#define KEY_ALT      8

// cross classes callback ------------------------------------------------------------------------------
/*! \enum CallbackType
    \brief callback event index

    Provides methods for the management teams of environment and interface elements that cause them
*/

/** test */
typedef enum {
	CBT_SELECT_ELEMENT_PALETTE,	/**< do when user select element in Palette */
	CBT_BEGIN_OPERATION,		/**< do when call beginOperation in SDK_Editor */
	CBT_END_OPERATION,			/**< do when call endOperation in SDK_Editor */
	CBT_POPUP_MENU,				/**< do when necessary to show the context menu in SDK_Editor */
	CBT_SDK_REDRAW_RECT,		/**< do when SDK require redraw rect */
	CBT_SDK_ADD_ELEMENT,		/**< do when element added to SDK */
	CBT_SDK_REMOVE_ELEMENT,		/**< do when element removed from SDK */
	CBT_SDK_CHANGE_SELECT,		/**< do when selectionset changes in SelectManager */
	CBT_SDK_RUN_CHANGE,			/**< do when call run() or stop() methods in SDK */
	CBT_RUN_COMMAND,			/**< do when user call internal command */
	CBT_CMD_ENABLED,			/**< do when commands need to refresh her status */
	CBT_CHANGE_PROPERTY,		/**< do when user changed property value in PropEditor */
	CBT_SELECT_PROPERTY,		/**< do when user selected property in PropEditor */
	CBT_CHECK_PROPERTY,			/**< do when user clicked on checkbox of property in PropEditor */
	CBT_EDIT_PROPERTY,			/**< do when user clicked on edit button in PropEditor */
	CBT_TAB_OPEN,				/**< do when successfully opened a new file in WindowManager (data is WM_Tab) */
	CBT_TAB_CHANGE,				/**< do when selected a new current tab in WindowManager (data is WM_Tab) */
	CBT_TAB_CLOSE,				/**< do when tab has been closed in WindowManager (data is WM_Tab) */
	CBT_KEY_DOWN,				/**< do when user press any key */
	CBT_ELEMENT_PROP_CHANGE,	/**< do when element property changed */
	CBT_CHANGE_SDK,				/**< do when changed current SDK in SDK_Editor */
	CBT_CHANGE_SCHEME,			/**< do when scheme is changed */
	CBT_CHANGE_RECENT_LIST,		/**< do when recent file list is change */
	CBT_CHANGE_SETTINGS,		/**< do when user change application settings */
	CBT_DEBUG_INFO,				/**< do when module generates some information for user */
} CallbackType;

/*! \class CallBack
    \brief interface for callback events

    Provides an interface for organizing callbacks between classes.
    To subscribe to the events use a field of type Event. Example: myObject->on_some_event += this
*/

class CallBack {
	public:
		virtual ~CallBack() { }
		/**
		 * Emit for all subscribers then user call on_some_event.run(..)
		 * @param owner class who emit this event
		 * @param type index of event from CallbackType enum
		 * @param data pointer to user data
		 */
		virtual void callback(void *owner, CallbackType type, const void *data) = 0;
};

/*! \class Event
    \brief list of subscribers

    Provides an interface to create a list of subscribers to callback between classes
*/

class Event {
	private:
		std::list<CallBack*> objs;
		CallbackType type;
		void *owner;
	public:
		bool enabled;	/**< Event enabled. If false, the execution method run have no effect */

		/**
		 * Create event object
		 * @param owner class who contain this event (usually this)
		 * @param type index of event
		 */
		Event(void *owner, CallbackType type) { this->type = type; this->owner = owner; enabled = true; }
		~Event() { objs.clear(); }
		/**
		 * Subscribe to event
		 * @param obj class who receive callback (must inherit from class CallBack)
		 */
		Event& operator +=(CallBack *obj) {
			objs.push_back(obj);
			return *this;
		}
		/**
		 * Unsubscribe from the event
		 * @param obj class that was previously added to the list
		 */
		Event& operator -=(CallBack *obj) {
			objs.remove(obj);
			return *this;
		}
		/**
		 * Run event for all subscribers
		 * @param data user data
		 */
		void run(const void *data) {
			if(enabled)
				for(std::list<CallBack*>::iterator event = objs.begin(); event != objs.end(); event++)
					(*event)->callback(owner, type, data);
		}
};

class ObjectEvent {
	public:
		typedef void (*event)(void *obj);
	private:
		void *obj;
		event proc;
	public:
		ObjectEvent() { clear(); }
		ObjectEvent(void *obj, event proc) { this->obj = obj; this->proc = proc; }
		void operator ()() {
			if(proc)
				proc(obj);
		}
		void clear() { obj = NULL; proc = NULL; }
		inline bool empty() { return obj == NULL; }
};

#define EVENT(x) ObjectEvent(this, (ObjectEvent::event)x)

Glib::RefPtr<Gdk::Pixbuf> getPointIcon(int type);

extern void initDirs();

extern void parseHintName(const ustring &text, ustring &name, ustring &hint);

#endif

/* SHARE_H_ */
