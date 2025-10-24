// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [FilterGenerator] class.
var (
	FilterGeneratorClass     _FilterGeneratorClass
	FilterGeneratorClassOnce sync.Once
)

func getFilterGeneratorClass() _FilterGeneratorClass {
	FilterGeneratorClassOnce.Do(func() {
		FilterGeneratorClass = _FilterGeneratorClass{objc.GetClass("CIFilterGenerator")}
	})
	return FilterGeneratorClass
}

type _FilterGeneratorClass struct {
	class objc.Class
}

// An interface definition for the [FilterGenerator] class.
type IFilterGenerator interface {
	objectivec.IObject
	// properties:
	ClassAttributes() objc.IObject /* cross-framework: NSDictionary */
	SetClassAttributes(value objc.IObject /* cross-framework: NSDictionary */)
	ExportedKeys() objc.IObject /* cross-framework: NSDictionary */
	// methods:
	ConnectObjectWithKeyToObjectWithKey(sourceObject objectivec.IObject, sourceKey objc.IObject /* cross-framework: NSString */, targetObject objectivec.IObject, targetKey objc.IObject /* cross-framework: NSString */)
	DisconnectObjectWithKeyToObjectWithKey(sourceObject objectivec.IObject, sourceKey objc.IObject /* cross-framework: NSString */, targetObject objectivec.IObject, targetKey objc.IObject /* cross-framework: NSString */)
	ExportKeyFromObjectWithName(key objc.IObject /* cross-framework: NSString */, targetObject objectivec.IObject, exportedKeyName objc.IObject /* cross-framework: NSString */)
	Filter() IFilter
	RegisterFilterName(name objc.IObject /* cross-framework: NSString */)
	RemoveExportedKey(exportedKeyName objc.IObject /* cross-framework: NSString */)
	SetAttributesForExportedKey(attributes objc.IObject /* cross-framework: NSDictionary */, key objc.IObject /* cross-framework: NSString */)
	WriteToURLAtomically(aURL objc.IObject /* cross-framework: NSURL */, flag bool) bool
}

// An object that creates and configures chains of individual image filters.
//
// The class provides methods for creating a object by chaining together existing objects to create complex effects. (A refers to the objects that are connected in the object.) The complex effect can be encapsulated as a object and saved as a file so that it can be used again. The contains an archived instance of all the objects that are chained together. Any filter generator files that you copy to are loaded when any of the loading methods provided by the class are invoked. A object is registered by its filename or, if present, by a class attribute that you supply in its description. You can create a object programmatically, using the methods provided by the class, or by using the editor view provided by Core Image.


// An object that creates and configures chains of individual image filters.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIFilterGenerator
type FilterGenerator struct {
	objectivec.Object
}

// FilterGeneratorFrom constructs a [FilterGenerator] from an unsafe.Pointer.
//
// An object that creates and configures chains of individual image filters.
func FilterGeneratorFrom(ptr unsafe.Pointer) FilterGenerator {
	return FilterGenerator{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (fc _FilterGeneratorClass) Alloc() FilterGenerator {
	rv := objc.Send[FilterGenerator](objc.ID(fc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (fc _FilterGeneratorClass) New() FilterGenerator {
	rv := objc.Send[FilterGenerator](objc.ID(fc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (f_ FilterGenerator) Init() FilterGenerator {
	rv := objc.Send[FilterGenerator](f_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (f_ FilterGenerator) Autorelease() FilterGenerator {
	rv := objc.Send[FilterGenerator](f_.ID, objc.Sel("autorelease"))
	return rv
}

// NewFilterGenerator creates a new FilterGenerator instance.
func NewFilterGenerator() FilterGenerator {
	return getFilterGeneratorClass().New()
}



// Initializes a filter generator object with the contents of a filter generator file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIFilterGenerator/init(contentsOf:)
func NewFilterGeneratorWithContentsOfURL(aURL objc.IObject /* cross-framework: NSURL */) FilterGenerator {
	instance := getFilterGeneratorClass().Alloc()
	rv := objc.Send[FilterGenerator](instance.ID, objc.Sel("initWithContentsOfURL:"), aURL)
	rv.Autorelease()
	return rv
}



// Creates and returns an empty filter generator object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIFilterGenerator/filterGenerator
func (fc _FilterGeneratorClass) FilterGenerator() IFilterGenerator {
	rv := objc.Send[FilterGenerator](objc.ID(fc.class), objc.Sel("filterGenerator"))
	return rv
}


// Creates and returns a filter generator object and initializes it with the contents of a filter generator file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIFilterGenerator/filterGeneratorWithContentsOfURL:
func (fc _FilterGeneratorClass) FilterGeneratorWithContentsOfURL(aURL objc.IObject /* cross-framework: NSURL */) IFilterGenerator {
	rv := objc.Send[FilterGenerator](objc.ID(fc.class), objc.Sel("filterGeneratorWithContentsOfURL:"), aURL)
	return rv
}


// Adds an object to the filter chain.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIFilterGenerator/connect(_:withKey:to:withKey:)
func (f_ FilterGenerator) ConnectObjectWithKeyToObjectWithKey(sourceObject objectivec.IObject, sourceKey objc.IObject /* cross-framework: NSString */, targetObject objectivec.IObject, targetKey objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](f_.ID, objc.Sel("connectObject:withKey:toObject:withKey:"), sourceObject, sourceKey, targetObject, targetKey)
}


// Removes the connection between two objects in the filter chain.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIFilterGenerator/disconnectObject(_:withKey:to:withKey:)
func (f_ FilterGenerator) DisconnectObjectWithKeyToObjectWithKey(sourceObject objectivec.IObject, sourceKey objc.IObject /* cross-framework: NSString */, targetObject objectivec.IObject, targetKey objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](f_.ID, objc.Sel("disconnectObject:withKey:toObject:withKey:"), sourceObject, sourceKey, targetObject, targetKey)
}


// Exports an input or output key of an object in the filter chain.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIFilterGenerator/exportKey(_:from:withName:)
func (f_ FilterGenerator) ExportKeyFromObjectWithName(key objc.IObject /* cross-framework: NSString */, targetObject objectivec.IObject, exportedKeyName objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](f_.ID, objc.Sel("exportKey:fromObject:withName:"), key, targetObject, exportedKeyName)
}


// Creates a filter object based on the filter chain.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIFilterGenerator/filter()
func (f_ FilterGenerator) Filter() IFilter {
	rv := objc.Send[Filter](f_.ID, objc.Sel("filter"))
	return rv
}


// Registers the name associated with a filter chain.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIFilterGenerator/registerFilterName(_:)
func (f_ FilterGenerator) RegisterFilterName(name objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](f_.ID, objc.Sel("registerFilterName:"), name)
}


// Removes a key that was previously exported.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIFilterGenerator/removeExportedKey(_:)
func (f_ FilterGenerator) RemoveExportedKey(exportedKeyName objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](f_.ID, objc.Sel("removeExportedKey:"), exportedKeyName)
}


// Sets a dictionary of attributes for an exported key.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIFilterGenerator/setAttributes(_:forExportedKey:)
func (f_ FilterGenerator) SetAttributesForExportedKey(attributes objc.IObject /* cross-framework: NSDictionary */, key objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setAttributes:forExportedKey:"), attributes, key)
}


// Archives a filter generator object to a filter generator file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIFilterGenerator/write(to:atomically:)
func (f_ FilterGenerator) WriteToURLAtomically(aURL objc.IObject /* cross-framework: NSURL */, flag bool) bool {
	rv := objc.Send[bool](f_.ID, objc.Sel("writeToURL:atomically:"), aURL, flag)
	return rv
}


// The class attributes associated with the filter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIFilterGenerator/classAttributes
func (f_ FilterGenerator) ClassAttributes() objc.IObject /* cross-framework: NSDictionary */ {
	rv := objc.Send[foundation.NSDictionary](f_.ID, objc.Sel("classAttributes"))
	return rv
}


// The class attributes associated with the filter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIFilterGenerator/classAttributes
func (f_ FilterGenerator) SetClassAttributes(value objc.IObject /* cross-framework: NSDictionary */) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setClassAttributes:"), value)
}


// Returns an array of the exported keys.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIFilterGenerator/exportedKeys
func (f_ FilterGenerator) ExportedKeys() objc.IObject /* cross-framework: NSDictionary */ {
	rv := objc.Send[foundation.NSDictionary](f_.ID, objc.Sel("exportedKeys"))
	return rv
}


