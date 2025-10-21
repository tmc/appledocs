// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [Pasteboard] class.
var (
	PasteboardClass     _PasteboardClass
	PasteboardClassOnce sync.Once
)

func getPasteboardClass() _PasteboardClass {
	PasteboardClassOnce.Do(func() {
		PasteboardClass = _PasteboardClass{objc.GetClass("NSPasteboard")}
	})
	return PasteboardClass
}

type _PasteboardClass struct {
	class objc.Class
}

// An interface definition for the [Pasteboard] class.
type IPasteboard interface {
	objectivec.IObject
	AddTypesOwner(newTypes []string, newOwner objectivec.IObject) int
	AvailableTypeFromArray(types []string) PasteboardType
	CanReadItemWithDataConformingToTypes(types []string) bool
	CanReadObjectForClassesOptions(classArray []objc.IClass, options unsafe.Pointer) bool
	ClearContents() int
	DataForType(dataType PasteboardType) foundation.Data
	DeclareTypesOwner(newTypes []string, newOwner objectivec.IObject) int
	DetectMetadataForTypesCompletionHandler(types unsafe.Pointer, completionHandler unsafe.Pointer)
	DetectPatternsForPatternsCompletionHandler(patterns unsafe.Pointer, completionHandler unsafe.Pointer)
	DetectValuesForPatternsCompletionHandler(patterns unsafe.Pointer, completionHandler unsafe.Pointer)
	IndexOfPasteboardItem(pasteboardItem IPasteboardItem) uint
	PrepareForNewContentsWithOptions(options PasteboardContentsOptions) int
	PropertyListForType(dataType PasteboardType) objc.ID
	ReadFileContentsTypeToFile(type_ PasteboardType, filename string) foundation.String
	ReadFileWrapper() unsafe.Pointer
	ReadObjectsForClassesOptions(classArray []objc.IClass, options unsafe.Pointer) foundation.Array
	ReleaseGlobally()
	SetDataForType(data foundation.IData, dataType PasteboardType) bool
	SetPropertyListForType(plist objectivec.IObject, dataType PasteboardType) bool
	SetStringForType(string_ string, dataType PasteboardType) bool
	StringForType(dataType PasteboardType) foundation.String
	WriteFileWrapper(wrapper unsafe.Pointer) bool
	WriteFileContents(filename string) bool
	WriteObjects(objects []objc.ID) bool
}

// An object that transfers data to and from the pasteboard server.
//
// The pasteboard server is shared by all running apps. It contains data that the user has cut or copied, as well as other data that one application wants to transfer to another. objects are an application’s sole interface to the server and to all pasteboard operations. An object is also used to transfer data between apps and service providers listed in each application’s Services menu. The drag pasteboard is used to transfer data that is being dragged by the user. A pasteboard can contain multiple items. You can directly write or read any object that implements the or   respectively. This allows you to write and read common items such as URLs, colors, images, strings, attributed strings, and sounds without an intermediary object. Your custom classes can also implement these protocols for use with the pasteboard. Writing methods such as provide a convenient means of writing to the first pasteboard item, without having to create the first pasteboard item. You can use code like this, for example: The general pasteboard, available by way of the class method, automatically participates with the Universal Clipboard feature in macOS 10.12 and later and in iOS 10.0 and later. There is no macOS API for interacting with this feature.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPasteboard
type Pasteboard struct {
	objectivec.Object
}

// PasteboardFrom constructs a [Pasteboard] from an unsafe.Pointer.
//
// An object that transfers data to and from the pasteboard server.
func PasteboardFrom(ptr unsafe.Pointer) Pasteboard {
	return Pasteboard{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (pc _PasteboardClass) Alloc() Pasteboard {
	rv := objc.Send[Pasteboard](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (pc _PasteboardClass) New() Pasteboard {
	rv := objc.Send[Pasteboard](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ Pasteboard) Init() Pasteboard {
	rv := objc.Send[Pasteboard](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ Pasteboard) Autorelease() Pasteboard {
	rv := objc.Send[Pasteboard](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewPasteboard creates a new Pasteboard instance.
func NewPasteboard() Pasteboard {
	return getPasteboardClass().New()
}




// Creates a new pasteboard object that supplies the specified data in as many types as possible based on the available filter services.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPasteboard/init(byFilteringData:ofType:)
func NewPasteboardByFilteringDataOfType(data foundation.IData, type_ PasteboardType) Pasteboard {
	rv := objc.Send[Pasteboard](objc.ID(getPasteboardClass().class), objc.Sel("pasteboardByFilteringData:ofType:"), data, type_)
	return rv
}



// Creates a new pasteboard object that supplies the specified file in as many types as possible based on the available filter services.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPasteboard/init(byFilteringFile:)
func NewPasteboardByFilteringFile(filename string) Pasteboard {
	rv := objc.Send[Pasteboard](objc.ID(getPasteboardClass().class), objc.Sel("pasteboardByFilteringFile:"), objc.String(filename))
	return rv
}



// Creates a new pasteboard object that supplies the specified pasteboard data in as many types as possible based on the available filter services.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPasteboard/init(byFilteringTypesInPasteboard:)
func NewPasteboardByFilteringTypesInPasteboard(pboard IPasteboard) Pasteboard {
	rv := objc.Send[Pasteboard](objc.ID(getPasteboardClass().class), objc.Sel("pasteboardByFilteringTypesInPasteboard:"), pboard)
	return rv
}



// Returns the pasteboard with the specified name.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPasteboard/init(name:)
func NewPasteboardWithName(name IPasteboardName) Pasteboard {
	rv := objc.Send[Pasteboard](objc.ID(getPasteboardClass().class), objc.Sel("pasteboardWithName:"), name)
	return rv
}


// Creates a new pasteboard object that supplies the specified data in as many types as possible based on the available filter services.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPasteboard/init(byFilteringData:ofType:)
func (pc _PasteboardClass) PasteboardByFilteringDataOfType(data foundation.IData, type_ PasteboardType) Pasteboard {
	rv := objc.Send[Pasteboard](objc.ID(pc.class), objc.Sel("pasteboardByFilteringData:ofType:"), data, type_)
	return rv
}

// Creates a new pasteboard object that supplies the specified file in as many types as possible based on the available filter services.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPasteboard/init(byFilteringFile:)
func (pc _PasteboardClass) PasteboardByFilteringFile(filename string) Pasteboard {
	rv := objc.Send[Pasteboard](objc.ID(pc.class), objc.Sel("pasteboardByFilteringFile:"), objc.String(filename))
	return rv
}

// Creates a new pasteboard object that supplies the specified pasteboard data in as many types as possible based on the available filter services.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPasteboard/init(byFilteringTypesInPasteboard:)
func (pc _PasteboardClass) PasteboardByFilteringTypesInPasteboard(pboard IPasteboard) Pasteboard {
	rv := objc.Send[Pasteboard](objc.ID(pc.class), objc.Sel("pasteboardByFilteringTypesInPasteboard:"), pboard)
	return rv
}

// Returns the pasteboard with the specified name.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPasteboard/init(name:)
func (pc _PasteboardClass) PasteboardWithName(name IPasteboardName) Pasteboard {
	rv := objc.Send[Pasteboard](objc.ID(pc.class), objc.Sel("pasteboardWithName:"), name)
	return rv
}

// Returns the data types that can be converted to the specified type using the available filter services.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPasteboard/types(filterableTo:)
func (pc _PasteboardClass) TypesFilterableTo(type_ PasteboardType) []string {
	rv := objc.Send[[]string](objc.ID(pc.class), objc.Sel("typesFilterableTo:"), type_)
	return rv
}

// Creates and returns a new pasteboard with a name that is guaranteed to be unique with respect to other pasteboards in the system.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPasteboard/withUniqueName()
func (pc _PasteboardClass) PasteboardWithUniqueName() Pasteboard {
	rv := objc.Send[Pasteboard](objc.ID(pc.class), objc.Sel("pasteboardWithUniqueName"))
	return rv
}

// The shared pasteboard object to use for general content.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPasteboard/general
func (pc _PasteboardClass) GeneralPasteboard() NSPasteboard {
	rv := objc.Send[NSPasteboard](objc.ID(pc.class), objc.Sel("generalPasteboard"))
	return rv
}
// Adds promises for the specified types to the first pasteboard item.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPasteboard/addTypes(_:owner:)
func (p_ Pasteboard) AddTypesOwner(newTypes []string, newOwner objectivec.IObject) int {
	rv := objc.Send[int](p_.ID, objc.Sel("addTypes:owner:"), newTypes, newOwner)
	return rv
}

// Scans the specified types for a type that the receiver supports.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPasteboard/availableType(from:)
func (p_ Pasteboard) AvailableTypeFromArray(types []string) PasteboardType {
	rv := objc.Send[PasteboardType](p_.ID, objc.Sel("availableTypeFromArray:"), types)
	return rv
}

// Returns a Boolean value that indicates whether the receiver contains any items that conform to the specified UTIs.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPasteboard/canReadItem(withDataConformingToTypes:)
func (p_ Pasteboard) CanReadItemWithDataConformingToTypes(types []string) bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("canReadItemWithDataConformingToTypes:"), types)
	return rv
}

// Returns a Boolean value that indicates whether the receiver contains any items that can be represented as an instance of any class in a given array.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPasteboard/canReadObject(forClasses:options:)
func (p_ Pasteboard) CanReadObjectForClassesOptions(classArray []objc.IClass, options unsafe.Pointer) bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("canReadObjectForClasses:options:"), classArray, options)
	return rv
}

// Clears the existing contents of the pasteboard.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPasteboard/clearContents()
func (p_ Pasteboard) ClearContents() int {
	rv := objc.Send[int](p_.ID, objc.Sel("clearContents"))
	return rv
}

// Returns the data for the specified type from the first item in the receiver that contains the type.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPasteboard/data(forType:)
func (p_ Pasteboard) DataForType(dataType PasteboardType) foundation.Data {
	rv := objc.Send[foundation.Data](p_.ID, objc.Sel("dataForType:"), dataType)
	return rv
}

// Prepares the receiver for a change in its contents by declaring the new types of data it will contain and a new owner.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPasteboard/declareTypes(_:owner:)
func (p_ Pasteboard) DeclareTypesOwner(newTypes []string, newOwner objectivec.IObject) int {
	rv := objc.Send[int](p_.ID, objc.Sel("declareTypes:owner:"), newTypes, newOwner)
	return rv
}

// Determines available metadata from the specified metadata types for the first pasteboard item, without notifying the person using the app.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPasteboard/detectMetadataForTypes:completionHandler:
func (p_ Pasteboard) DetectMetadataForTypesCompletionHandler(types unsafe.Pointer, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("detectMetadataForTypes:completionHandler:"), types, completionHandler)
}

// Determines whether the first pasteboard item matches the specified patterns, without notifying the person using the app.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPasteboard/detectPatternsForPatterns:completionHandler:
func (p_ Pasteboard) DetectPatternsForPatternsCompletionHandler(patterns unsafe.Pointer, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("detectPatternsForPatterns:completionHandler:"), patterns, completionHandler)
}

// Determines whether the first pasteboard item matches the specified patterns, reading the contents if it finds a match.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPasteboard/detectValuesForPatterns:completionHandler:
func (p_ Pasteboard) DetectValuesForPatternsCompletionHandler(patterns unsafe.Pointer, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("detectValuesForPatterns:completionHandler:"), patterns, completionHandler)
}

// Returns the index of the specified pasteboard item.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPasteboard/index(of:)
func (p_ Pasteboard) IndexOfPasteboardItem(pasteboardItem IPasteboardItem) uint {
	rv := objc.Send[uint](p_.ID, objc.Sel("indexOfPasteboardItem:"), pasteboardItem)
	return rv
}

// Prepares the pasteboard to receive new contents, removing the existing pasteboard contents.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPasteboard/prepareForNewContents(with:)
func (p_ Pasteboard) PrepareForNewContentsWithOptions(options PasteboardContentsOptions) int {
	rv := objc.Send[int](p_.ID, objc.Sel("prepareForNewContentsWithOptions:"), options)
	return rv
}

// Returns the property list for the specified type from the first item in the receiver that contains the type.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPasteboard/propertyList(forType:)
func (p_ Pasteboard) PropertyListForType(dataType PasteboardType) objc.ID {
	rv := objc.Send[objc.ID](p_.ID, objc.Sel("propertyListForType:"), dataType)
	return rv
}

// Reads data representing a file’s contents from the receiver and writes it to the specified file.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPasteboard/readFileContentsType(_:toFile:)
func (p_ Pasteboard) ReadFileContentsTypeToFile(type_ PasteboardType, filename string) foundation.String {
	rv := objc.Send[foundation.String](p_.ID, objc.Sel("readFileContentsType:toFile:"), type_, objc.String(filename))
	return rv
}

// Reads data representing a file’s contents from the receiver and returns it as a file wrapper.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPasteboard/readFileWrapper()
func (p_ Pasteboard) ReadFileWrapper() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("readFileWrapper"))
	return rv
}

// Reads from the receiver objects that best match the specified array of classes.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPasteboard/readObjects(forClasses:options:)
func (p_ Pasteboard) ReadObjectsForClassesOptions(classArray []objc.IClass, options unsafe.Pointer) foundation.Array {
	rv := objc.Send[foundation.Array](p_.ID, objc.Sel("readObjectsForClasses:options:"), classArray, options)
	return rv
}

// Releases the receiver’s resources in the pasteboard server.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPasteboard/releaseGlobally()
func (p_ Pasteboard) ReleaseGlobally() {
	objc.Send[objc.ID](p_.ID, objc.Sel("releaseGlobally"))
}

// Sets the data as the representation for the specified type for the first item on the receiver.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPasteboard/setData(_:forType:)
func (p_ Pasteboard) SetDataForType(data foundation.IData, dataType PasteboardType) bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("setData:forType:"), data, dataType)
	return rv
}

// Sets the given property list as the representation for the specified type for the first item on the receiver.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPasteboard/setPropertyList(_:forType:)
func (p_ Pasteboard) SetPropertyListForType(plist objectivec.IObject, dataType PasteboardType) bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("setPropertyList:forType:"), plist, dataType)
	return rv
}

// Sets the given string as the representation for the specified type for the first item on the receiver.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPasteboard/setString(_:forType:)
func (p_ Pasteboard) SetStringForType(string_ string, dataType PasteboardType) bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("setString:forType:"), objc.String(string_), dataType)
	return rv
}

// Returns a concatenation of the strings for the specified type from all the items in the receiver that contain the type.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPasteboard/string(forType:)
func (p_ Pasteboard) StringForType(dataType PasteboardType) foundation.String {
	rv := objc.Send[foundation.String](p_.ID, objc.Sel("stringForType:"), dataType)
	return rv
}

// Writes the serialized contents of the specified file wrapper to the pasteboard.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPasteboard/write(_:)
func (p_ Pasteboard) WriteFileWrapper(wrapper unsafe.Pointer) bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("writeFileWrapper:"), wrapper)
	return rv
}

// Writes the contents of the specified file to the pasteboard.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPasteboard/writeFileContents(_:)
func (p_ Pasteboard) WriteFileContents(filename string) bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("writeFileContents:"), objc.String(filename))
	return rv
}

// Writes an array of objects to the receiver.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPasteboard/writeObjects(_:)
func (p_ Pasteboard) WriteObjects(objects []objc.ID) bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("writeObjects:"), objects)
	return rv
}

// The current pasteboard access behavior. The user can customize this behavior per-app in System Settings for any app that has triggered a pasteboard access alert in the past.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPasteboard/accessBehavior-86972
func (p_ Pasteboard) AccessBehavior() PasteboardAccessBehavior {
	rv := objc.Send[PasteboardAccessBehavior](p_.ID, objc.Sel("accessBehavior"))
	return rv
}

// The receiver’s change count.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPasteboard/changeCount
func (p_ Pasteboard) ChangeCount() int {
	rv := objc.Send[int](p_.ID, objc.Sel("changeCount"))
	return rv
}

// The shared pasteboard object to use for general content.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPasteboard/general
func (p_ Pasteboard) GeneralPasteboard() NSPasteboard {
	rv := objc.Send[NSPasteboard](p_.ID, objc.Sel("generalPasteboard"))
	return rv
}

// The receiver’s name.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPasteboard/name-swift.property
func (p_ Pasteboard) Name() PasteboardName {
	rv := objc.Send[PasteboardName](p_.ID, objc.Sel("name"))
	return rv
}

// An array that contains all the items held by the pasteboard.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPasteboard/pasteboardItems
func (p_ Pasteboard) PasteboardItems() []PasteboardItem {
	rv := objc.Send[[]PasteboardItem](p_.ID, objc.Sel("pasteboardItems"))
	return rv
}

// An array of the receiver’s supported data types.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPasteboard/types
func (p_ Pasteboard) Types() []string {
	rv := objc.Send[[]string](p_.ID, objc.Sel("types"))
	return rv
}


