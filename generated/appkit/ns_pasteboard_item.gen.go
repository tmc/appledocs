// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class NSPasteboardItem */


/* debug [class_header]: Header for NSPasteboardItem */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for PasteboardItem */
// An interface definition for the [PasteboardItem] class.
type IPasteboardItem interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for PasteboardItem */
	// properties:
	CollaborationMetadata() objectivec.IObject
	SetCollaborationMetadata(value objectivec.IObject)
	Types() []string
	PasteboardItems() IPasteboardItem
	SetPasteboardItems(value IPasteboardItem)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for PasteboardItem */
	// methods:
	AvailableTypeFromArray(types []string) PasteboardType /* typedef */
	DataForType(type_ PasteboardType /* typedef */) foundation.Data
	DetectMetadataForTypesCompletionHandler(types unsafe.Pointer, completionHandler unsafe.Pointer)
	DetectPatternsForPatternsCompletionHandler(patterns unsafe.Pointer, completionHandler unsafe.Pointer)
	DetectValuesForPatternsCompletionHandler(patterns unsafe.Pointer, completionHandler unsafe.Pointer)
	PropertyListForType(type_ PasteboardType /* typedef */) objc.ID
	SetDataForType(data objc.IObject /* cross-framework: NSData */, type_ PasteboardType /* typedef */) bool
	SetDataProviderForTypes(dataProvider unsafe.Pointer, types []string) bool
	SetPropertyListForType(propertyList objc.IObject, type_ PasteboardType /* typedef */) bool
	SetStringForType(string_ objc.IObject /* cross-framework: NSString */, type_ PasteboardType /* typedef */) bool
	StringForType(type_ PasteboardType /* typedef */) foundation.String
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for PasteboardItem */
// Alloc allocates a new instance without initialization.
func (pc _PasteboardItemClass) Alloc() PasteboardItem {
	rv := objc.Send[PasteboardItem](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for PasteboardItem */
// An item on a pasteboard.
//
// There are three main uses for an object: Providing data on the pasteboard. You can create one or more pasteboard items, set data or data providers for types, and write them to the pasteboard. Customizing data already on the pasteboard. As a delegate or subclass, you can retrieve the pasteboard items currently on the pasteboard, read the existing types and data, and set new data and data providers for types as necessary. Retrieving data from the pasteboard. You can retrieve pasteboard items from the pasteboard and then read the data for types you’re interested in. A pasteboard item can be associated with a single pasteboard. When you create an item, you can write it to any pasteboard. When you pass an item to a pasteboard in , that item becomes bound to the pasteboard it writes to. When you retrieve items from a pasteboard using or , the returned items are associated with the messaged pasteboard. Passing an item that is already associated with a pasteboard into causes an exception. Use pasteboard items during a single pasteboard interaction, rather than retaining and reusing them. A pasteboard item is only valid until the owner of the pasteboard changes.


// An item on a pasteboard.
//
// [Full Topic]
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for PasteboardItem *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for PasteboardItem */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for PasteboardItem */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for PasteboardItem */

// Returns from a given array of types the first type within the pasteboard item, according to the ordering of types.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPasteboardItem/availableType(from:)
func (p_ PasteboardItem) AvailableTypeFromArray(types []string) PasteboardType /* typedef */ {
	rv := objc.Send[foundation.NSString](p_.ID, objc.Sel("availableTypeFromArray:"), types)
	return rv
}/* debug [instance_methods/method]: AvailableTypeFromArray */


// Returns the value for the specified type as a data object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPasteboardItem/data(forType:)
func (p_ PasteboardItem) DataForType(type_ PasteboardType /* typedef */) foundation.Data {
	rv := objc.Send[foundation.Data](p_.ID, objc.Sel("dataForType:"), type_)
	return rv
}/* debug [instance_methods/method]: DataForType */


// Determines available metadata from the specified metadata types for this pasteboard item, without notifying the person using the app.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPasteboardItem/detectMetadataForTypes:completionHandler:
func (p_ PasteboardItem) DetectMetadataForTypesCompletionHandler(types unsafe.Pointer, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("detectMetadataForTypes:completionHandler:"), types, completionHandler)
}/* debug [instance_methods/method]: DetectMetadataForTypesCompletionHandler */


// Determines whether this pasteboard item matches the specified patterns, without notifying the person using the app.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPasteboardItem/detectPatternsForPatterns:completionHandler:
func (p_ PasteboardItem) DetectPatternsForPatternsCompletionHandler(patterns unsafe.Pointer, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("detectPatternsForPatterns:completionHandler:"), patterns, completionHandler)
}/* debug [instance_methods/method]: DetectPatternsForPatternsCompletionHandler */


// Determines whether this pasteboard item matches the specified patterns, reading the contents if it finds a match.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPasteboardItem/detectValuesForPatterns:completionHandler:
func (p_ PasteboardItem) DetectValuesForPatternsCompletionHandler(patterns unsafe.Pointer, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("detectValuesForPatterns:completionHandler:"), patterns, completionHandler)
}/* debug [instance_methods/method]: DetectValuesForPatternsCompletionHandler */


// Returns the value for the specified type as a property list.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPasteboardItem/propertyList(forType:)
func (p_ PasteboardItem) PropertyListForType(type_ PasteboardType /* typedef */) objc.ID {
	rv := objc.Send[objc.ID](p_.ID, objc.Sel("propertyListForType:"), type_)
	return rv
}/* debug [instance_methods/method]: PropertyListForType */


// Sets the value for a specified type as a data object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPasteboardItem/setData(_:forType:)
func (p_ PasteboardItem) SetDataForType(data objc.IObject /* cross-framework: NSData */, type_ PasteboardType /* typedef */) bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("setData:forType:"), data, type_)
	return rv
}/* debug [instance_methods/method]: SetDataForType */


// Sets the data provider for the specified types.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPasteboardItem/setDataProvider(_:forTypes:)
func (p_ PasteboardItem) SetDataProviderForTypes(dataProvider unsafe.Pointer, types []string) bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("setDataProvider:forTypes:"), dataProvider, types)
	return rv
}/* debug [instance_methods/method]: SetDataProviderForTypes */


// Sets the value for a specified type as a property list.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPasteboardItem/setPropertyList(_:forType:)
func (p_ PasteboardItem) SetPropertyListForType(propertyList objc.IObject, type_ PasteboardType /* typedef */) bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("setPropertyList:forType:"), propertyList, type_)
	return rv
}/* debug [instance_methods/method]: SetPropertyListForType */


// Sets the value for a specified type as a string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPasteboardItem/setString(_:forType:)
func (p_ PasteboardItem) SetStringForType(string_ objc.IObject /* cross-framework: NSString */, type_ PasteboardType /* typedef */) bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("setString:forType:"), string_, type_)
	return rv
}/* debug [instance_methods/method]: SetStringForType */


// Returns the value for the specified type as a string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPasteboardItem/string(forType:)
func (p_ PasteboardItem) StringForType(type_ PasteboardType /* typedef */) foundation.String {
	rv := objc.Send[foundation.String](p_.ID, objc.Sel("stringForType:"), type_)
	return rv
}/* debug [instance_methods/method]: StringForType */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for PasteboardItem */

// A model object you use for conveying data during a collaboration.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPasteboardItem/collaborationMetadata
func (p_ PasteboardItem) CollaborationMetadata() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](p_.ID, objc.Sel("collaborationMetadata"))
	return rv
}/* debug [instance_properties/getter]: collaborationMetadata */


// A model object you use for conveying data during a collaboration.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPasteboardItem/collaborationMetadata
func (p_ PasteboardItem) SetCollaborationMetadata(value objectivec.IObject) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setCollaborationMetadata:"), value)
}/* debug [instance_properties/setter]: collaborationMetadata */


// An array of uniform type identifier strings of the data types that the receiver supports.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPasteboardItem/types
func (p_ PasteboardItem) Types() []string {
	rv := objc.Send[[]string](p_.ID, objc.Sel("types"))
	return rv
}/* debug [instance_properties/getter]: types */


// An array that contains all the items held by the pasteboard.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspasteboard/pasteboarditems
func (p_ PasteboardItem) PasteboardItems() IPasteboardItem {
	rv := objc.Send[PasteboardItem](p_.ID, objc.Sel("pasteboardItems"))
	return rv
}/* debug [instance_properties/getter]: pasteboardItems */


// An array that contains all the items held by the pasteboard.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspasteboard/pasteboarditems
func (p_ PasteboardItem) SetPasteboardItems(value IPasteboardItem) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setPasteboardItems:"), value)
}/* debug [instance_properties/setter]: pasteboardItems */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NSPasteboardItem */



