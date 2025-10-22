// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [PasteboardItem] class.
var (
	PasteboardItemClass     _PasteboardItemClass
	PasteboardItemClassOnce sync.Once
)

func getPasteboardItemClass() _PasteboardItemClass {
	PasteboardItemClassOnce.Do(func() {
		PasteboardItemClass = _PasteboardItemClass{objc.GetClass("NSPasteboardItem")}
	})
	return PasteboardItemClass
}

type _PasteboardItemClass struct {
	class objc.Class
}

// An interface definition for the [PasteboardItem] class.
type IPasteboardItem interface {
	objectivec.IObject
	AvailableTypeFromArray(types []string) PasteboardType
	DataForType(type_ PasteboardType) foundation.Data
	DetectMetadataForTypesCompletionHandler(types unsafe.Pointer, completionHandler unsafe.Pointer)
	DetectPatternsForPatternsCompletionHandler(patterns unsafe.Pointer, completionHandler unsafe.Pointer)
	DetectValuesForPatternsCompletionHandler(patterns unsafe.Pointer, completionHandler unsafe.Pointer)
	PropertyListForType(type_ PasteboardType) objc.ID
	SetDataForType(data foundation.IData, type_ PasteboardType) bool
	SetDataProviderForTypes(dataProvider objectivec.IObject, types []string) bool
	SetPropertyListForType(propertyList objectivec.IObject, type_ PasteboardType) bool
	SetStringForType(string_ string, type_ PasteboardType) bool
	StringForType(type_ PasteboardType) foundation.String
	CollaborationMetadata() unsafe.Pointer
	SetCollaborationMetadata(value unsafe.Pointer)
	Types() []string
	PasteboardItems() NSPasteboardItem
	SetPasteboardItems(value IPasteboardItem)
}

// An item on a pasteboard.
//
// There are three main uses for an object: Providing data on the pasteboard. You can create one or more pasteboard items, set data or data providers for types, and write them to the pasteboard. Customizing data already on the pasteboard. As a delegate or subclass, you can retrieve the pasteboard items currently on the pasteboard, read the existing types and data, and set new data and data providers for types as necessary. Retrieving data from the pasteboard. You can retrieve pasteboard items from the pasteboard and then read the data for types you’re interested in. A pasteboard item can be associated with a single pasteboard. When you create an item, you can write it to any pasteboard. When you pass an item to a pasteboard in , that item becomes bound to the pasteboard it writes to. When you retrieve items from a pasteboard using or , the returned items are associated with the messaged pasteboard. Passing an item that is already associated with a pasteboard into causes an exception. Use pasteboard items during a single pasteboard interaction, rather than retaining and reusing them. A pasteboard item is only valid until the owner of the pasteboard changes.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPasteboardItem
type PasteboardItem struct {
	objectivec.Object
}

// PasteboardItemFrom constructs a [PasteboardItem] from an unsafe.Pointer.
//
// An item on a pasteboard.
func PasteboardItemFrom(ptr unsafe.Pointer) PasteboardItem {
	return PasteboardItem{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (pc _PasteboardItemClass) Alloc() PasteboardItem {
	rv := objc.Send[PasteboardItem](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (pc _PasteboardItemClass) New() PasteboardItem {
	rv := objc.Send[PasteboardItem](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ PasteboardItem) Init() PasteboardItem {
	rv := objc.Send[PasteboardItem](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ PasteboardItem) Autorelease() PasteboardItem {
	rv := objc.Send[PasteboardItem](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewPasteboardItem creates a new PasteboardItem instance.
func NewPasteboardItem() PasteboardItem {
	return getPasteboardItemClass().New()
}


// Returns from a given array of types the first type within the pasteboard item, according to the ordering of types.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPasteboardItem/availableType(from:)
func (p_ PasteboardItem) AvailableTypeFromArray(types []string) PasteboardType {
	rv := objc.Send[PasteboardType](p_.ID, objc.Sel("availableTypeFromArray:"), types)
	return rv
}

// Returns the value for the specified type as a data object.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPasteboardItem/data(forType:)
func (p_ PasteboardItem) DataForType(type_ PasteboardType) foundation.Data {
	rv := objc.Send[foundation.Data](p_.ID, objc.Sel("dataForType:"), type_)
	return rv
}

// Determines available metadata from the specified metadata types for this pasteboard item, without notifying the person using the app.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPasteboardItem/detectMetadataForTypes:completionHandler:
func (p_ PasteboardItem) DetectMetadataForTypesCompletionHandler(types unsafe.Pointer, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("detectMetadataForTypes:completionHandler:"), types, completionHandler)
}

// Determines whether this pasteboard item matches the specified patterns, without notifying the person using the app.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPasteboardItem/detectPatternsForPatterns:completionHandler:
func (p_ PasteboardItem) DetectPatternsForPatternsCompletionHandler(patterns unsafe.Pointer, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("detectPatternsForPatterns:completionHandler:"), patterns, completionHandler)
}

// Determines whether this pasteboard item matches the specified patterns, reading the contents if it finds a match.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPasteboardItem/detectValuesForPatterns:completionHandler:
func (p_ PasteboardItem) DetectValuesForPatternsCompletionHandler(patterns unsafe.Pointer, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("detectValuesForPatterns:completionHandler:"), patterns, completionHandler)
}

// Returns the value for the specified type as a property list.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPasteboardItem/propertyList(forType:)
func (p_ PasteboardItem) PropertyListForType(type_ PasteboardType) objc.ID {
	rv := objc.Send[objc.ID](p_.ID, objc.Sel("propertyListForType:"), type_)
	return rv
}

// Sets the value for a specified type as a data object.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPasteboardItem/setData(_:forType:)
func (p_ PasteboardItem) SetDataForType(data foundation.IData, type_ PasteboardType) bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("setData:forType:"), data, type_)
	return rv
}

// Sets the data provider for the specified types.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPasteboardItem/setDataProvider(_:forTypes:)
func (p_ PasteboardItem) SetDataProviderForTypes(dataProvider objectivec.IObject, types []string) bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("setDataProvider:forTypes:"), dataProvider, types)
	return rv
}

// Sets the value for a specified type as a property list.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPasteboardItem/setPropertyList(_:forType:)
func (p_ PasteboardItem) SetPropertyListForType(propertyList objectivec.IObject, type_ PasteboardType) bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("setPropertyList:forType:"), propertyList, type_)
	return rv
}

// Sets the value for a specified type as a string.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPasteboardItem/setString(_:forType:)
func (p_ PasteboardItem) SetStringForType(string_ string, type_ PasteboardType) bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("setString:forType:"), objc.String(string_), type_)
	return rv
}

// Returns the value for the specified type as a string.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPasteboardItem/string(forType:)
func (p_ PasteboardItem) StringForType(type_ PasteboardType) foundation.String {
	rv := objc.Send[foundation.String](p_.ID, objc.Sel("stringForType:"), type_)
	return rv
}

// A model object you use for conveying data during a collaboration.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPasteboardItem/collaborationMetadata
func (p_ PasteboardItem) CollaborationMetadata() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("collaborationMetadata"))
	return rv
}


// SetCollaborationMetadata sets the value of the collaborationMetadata property.
// A model object you use for conveying data during a collaboration.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPasteboardItem/collaborationMetadata
func (p_ PasteboardItem) SetCollaborationMetadata(value unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setCollaborationMetadata:"), value)
}

// An array of uniform type identifier strings of the data types that the receiver supports.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPasteboardItem/types
func (p_ PasteboardItem) Types() []string {
	rv := objc.Send[[]string](p_.ID, objc.Sel("types"))
	return rv
}

// An array that contains all the items held by the pasteboard.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspasteboard/pasteboarditems
func (p_ PasteboardItem) PasteboardItems() NSPasteboardItem {
	rv := objc.Send[NSPasteboardItem](p_.ID, objc.Sel("pasteboardItems"))
	return rv
}


// SetPasteboardItems sets the value of the pasteboardItems property.
// An array that contains all the items held by the pasteboard.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspasteboard/pasteboarditems
func (p_ PasteboardItem) SetPasteboardItems(value IPasteboardItem) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setPasteboardItems:"), value)
}



