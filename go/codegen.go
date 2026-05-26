package main

import "fmt"

// CodeGen module - Go translation of CodeGen.h/cpp
// Original C++ code from HiAsm project, translated to Go

// CGTParams represents the parameter indices for code generation
type CGTParams int

const (
	ParamCodePath CGTParams = iota
	ParamDebugMode
	ParamDebugServerPort
	ParamDebugClientPort
	ParamProjectPath
	ParamHiasmVersion
	ParamUserName
	ParamUserMail
	ParamProjectName
	ParamSdeWidth
	ParamSdeHeight
	ParamCompiler
)

const (
	CGTSize = 85
)

// Resources represents a collection of resource files
type Resources struct {
	Items []string
	Icons int
}

// Init initializes the resources collection
func (r *Resources) Init() {
	r.Items = []string{}
	r.Icons = 0
}

// NextIcon returns the next icon index
func (r *Resources) NextIcon() int {
	index := r.Icons
	r.Icons++
	return index
}

// Global resources instance
var ResourcesGlobal = &Resources{}

// Event callback type for debug events
type DebugCallback func(text string)

// cgtOnDebug is the global debug event handler
var cgtOnDebug DebugCallback

// SDK functions (stubs - original C++ had complex SDK interactions)

// SDKGetCount returns the count of elements in SDK (stub)
func SDKGetCount(sdk interface{}) int {
	return 0
}

// SDKGetElement returns an element by index (stub)
func SDKGetElement(sdk interface{}, index uint) interface{} {
	return nil
}

// SDKGetElementName returns an element by name (stub)
func SDKGetElementName(sdk interface{}, name string) interface{} {
	return nil
}

// SDKGetParent returns the parent SDK (stub)
func SDKGetParent(sdk interface{}) interface{} {
	return nil
}

// Element functions (stubs - original C++ had complex Element class interactions)

// ElGetFlag returns the element flag (stub)
func ElGetFlag(e interface{}) int {
	return 0
}

// ElGetPropCount returns the count of properties (stub)
func ElGetPropCount(e interface{}) int {
	return 0
}

// ElGetProperty returns a property by index (stub)
func ElGetProperty(e interface{}, index uint) interface{} {
	return nil
}

// ElIsDefProp checks if property is default (stub)
func ElIsDefProp(e interface{}, index int) bool {
	return false
}

// ElSetCodeName sets the code name for an element (stub)
func ElSetCodeName(e interface{}, name string) interface{} {
	return e
}

// ElGetCodeName returns the code name (stub)
func ElGetCodeName(e interface{}) string {
	return ""
}

// ElGetClassName returns the class name (stub)
func ElGetClassName(e interface{}) string {
	return ""
}

// ElGetInfSub returns the info sub (stub)
func ElGetInfSub(e interface{}) string {
	return ""
}

// ElGetPtCount returns the point count (stub)
func ElGetPtCount(e interface{}) int {
	return 0
}

// ElGetPt returns a point by index (stub)
func ElGetPt(e interface{}, index int) interface{} {
	return nil
}

// ElGetPtName returns a point by name (stub)
func ElGetPtName(e interface{}, name string) interface{} {
	return nil
}

// ElGetClassIndex returns the class index (stub)
func ElGetClassIndex(e interface{}) byte {
	return 0
}

// ElGetSDK returns the SDK for an element (stub)
func ElGetSDK(e interface{}) interface{} {
	return nil
}

// ElGetSDKByIndex returns SDK by index (stub)
func ElGetSDKByIndex(e interface{}, index int) interface{} {
	return nil
}

// ElGetSDKCount returns the SDK count (stub)
func ElGetSDKCount(e interface{}) int {
	return 0
}

// ElGetSDKBName returns SDK base name (stub)
func ElGetSDKBName(e interface{}, index int) string {
	return ""
}

// ElLinkIs checks if element is a link (stub)
func ElLinkIs(e interface{}) bool {
	return false
}

// ElLinkMain returns the main linked element (stub)
func ElLinkMain(e interface{}) interface{} {
	return nil
}

// ElGetEID returns the element ID (stub)
func ElGetEID(e interface{}) int {
	return 0
}

// ElGetGroup returns the element group (stub)
func ElGetGroup(e interface{}) int {
	return 0
}

// ElGetPos returns the element position (stub)
func ElGetPos(e interface{}) (x, y int) {
	return 0, 0
}

// ElGetSize returns the element size (stub)
func ElGetSize(e interface{}) (w, h int) {
	return 0, 0
}

// ElGetData returns user data (stub)
func ElGetData(e interface{}) interface{} {
	return nil
}

// ElSetData sets user data (stub)
func ElSetData(e interface{}, data interface{}) {
}

// IsDebug checks if in debug mode (stub)
func IsDebug(e interface{}) bool {
	return false
}

// ElGetParent returns the parent element (stub)
func ElGetParent(e interface{}) interface{} {
	return nil
}

// ElGetPropertyListCount returns property list count (stub)
func ElGetPropertyListCount(e interface{}) int {
	return 0
}

// ElGetPropertyListItem returns a property list item (stub)
func ElGetPropertyListItem(e interface{}, index int) interface{} {
	return nil
}

// ElGetInterface returns the interface string (stub)
func ElGetInterface(e interface{}) string {
	return ""
}

// ElGetInherit returns the inherit string (stub)
func ElGetInherit(e interface{}) string {
	return ""
}

// PropertyList functions (stubs)

// PlGetName returns property name (stub)
func PlGetName(p interface{}) string {
	return ""
}

// PlGetInfo returns property info (stub)
func PlGetInfo(p interface{}) string {
	return ""
}

// PlGetGroup returns property group (stub)
func PlGetGroup(p interface{}) string {
	return ""
}

// PlGetProperty returns property value (stub)
func PlGetProperty(p interface{}) interface{} {
	return nil
}

// PlGetOwner returns property owner (stub)
func PlGetOwner(p interface{}) interface{} {
	return nil
}

// Point functions (stubs)

// PtGetLinkPoint returns the linked point (stub)
func PtGetLinkPoint(p interface{}) interface{} {
	return nil
}

// PtGetRLinkPoint returns the real linked point (stub)
func PtGetRLinkPoint(p interface{}) interface{} {
	return nil
}

// PtGetType returns the point type (stub)
func PtGetType(p interface{}) int {
	return 0
}

// PtGetName returns the point name (stub)
func PtGetName(p interface{}) string {
	return ""
}

// PtGetParent returns the point parent (stub)
func PtGetParent(p interface{}) interface{} {
	return nil
}

// PtGetIndex returns the point index (stub)
func PtGetIndex(p interface{}) int {
	return 0
}

// PtDpeGetName returns the DPE point name (stub)
func PtDpeGetName(p interface{}) string {
	return ""
}

// PtGetDataType returns the point data type (stub)
func PtGetDataType(p interface{}) int {
	return 0
}

// PtGetInfo returns the point info (stub)
func PtGetInfo(p interface{}) string {
	return ""
}

// Property functions (stubs)

// PropGetType returns the property type (stub)
func PropGetType(prop interface{}) int {
	return 0
}

// PropGetName returns the property name (stub)
func PropGetName(prop interface{}) string {
	return ""
}

// PropGetValue returns the property value (stub)
func PropGetValue(prop interface{}) interface{} {
	return nil
}

// PropToByte converts property to byte (stub)
func PropToByte(prop interface{}) byte {
	return 0
}

// PropToInteger converts property to integer (stub)
func PropToInteger(prop interface{}) int {
	return 0
}

// PropToReal converts property to real (stub)
func PropToReal(prop interface{}) float64 {
	return 0.0
}

// PropToString converts property to string (stub)
func PropToString(p interface{}) string {
	return ""
}

// PropGetLinkedElement returns the linked element by property name (stub)
func PropGetLinkedElement(e interface{}, propName string) interface{} {
	return nil
}

// PropGetLinkedElementInfo returns linked element info (stub)
func PropGetLinkedElementInfo(e interface{}, prop interface{}, hint string) interface{} {
	return nil
}

// PropIsTranslate checks if property is translatable (stub)
func PropIsTranslate(e interface{}, p interface{}) int {
	return 0
}

// PropSaveToFile saves property to file (stub)
func PropSaveToFile(p interface{}, fileName string) int {
	return 0
}

// Resource functions

// ResAddFile adds a file to resources
func ResAddFile(name string) int {
	ResourcesGlobal.Items = append(ResourcesGlobal.Items, name)
	return 0
}

// ResAddIcon adds an icon to resources (stub - no GUI available)
func ResAddIcon(p interface{}) string {
	iconName := fmt.Sprintf("icon%d.png", ResourcesGlobal.NextIcon())
	ResourcesGlobal.Items = append(ResourcesGlobal.Items, iconName)
	return iconName
}

// ResAddStr adds a string to resources (stub)
func ResAddStr(p string) string {
	return ""
}

// ResAddStream adds a stream to resources (stub)
func ResAddStream(p interface{}) string {
	return ""
}

// ResAddWave adds a wave to resources (stub)
func ResAddWave(p interface{}) string {
	return ""
}

// ResAddBitmap adds a bitmap to resources (stub)
func ResAddBitmap(p interface{}) string {
	return ""
}

// ResAddMenu adds a menu to resources (stub)
func ResAddMenu(p interface{}) string {
	return ""
}

// ResEmpty checks if resources are empty (stub)
func ResEmpty() int {
	if len(ResourcesGlobal.Items) == 0 {
		return 1
	}
	return 0
}

// ResSetPref sets the resource prefix (stub)
func ResSetPref(pref string) int {
	return 0
}

// Other functions

// Debug sends a debug message
func Debug(text string, color int) int {
	if cgtOnDebug != nil {
		if color != 0 {
			prefix := "@"
			if color == 0x0000FF {
				prefix = "#"
			} else if color == 0xFF0000 {
				prefix = "!"
			}
			cgtOnDebug(prefix + text)
		} else {
			cgtOnDebug(text)
		}
	}
	return 0
}

// GetParam retrieves a parameter by index (stub)
func GetParam(index int16, value interface{}) int {
	switch CGTParams(index) {
	case ParamCodePath:
		return 0
	case ParamDebugMode:
		return 0
	case ParamDebugServerPort:
		return 0
	case ParamDebugClientPort:
		return 0
	case ParamProjectPath:
		return 0
	case ParamHiasmVersion:
		return 0
	case ParamUserName:
		return 0
	case ParamUserMail:
		return 0
	case ParamProjectName:
		return 0
	case ParamSdeWidth:
		return 0
	case ParamSdeHeight:
		return 0
	case ParamCompiler:
		return 0
	}
	return 1
}

// Error reports an error (stub)
func Error(line int, e interface{}, text string) int {
	if cgtOnDebug != nil {
		cgtOnDebug("!" + text)
	}
	return 0
}

// Array functions (stubs)

// ArrCount returns array count (stub)
func ArrCount(a interface{}) int {
	return 0
}

// ArrType returns array type (stub)
func ArrType(a interface{}) int {
	return 0
}

// ArrItemName returns array item name (stub)
func ArrItemName(a interface{}, index int) string {
	return ""
}

// ArrItemData returns array item data (stub)
func ArrItemData(a interface{}, index int) interface{} {
	return nil
}
