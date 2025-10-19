// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [FilterGenerator] class.
var (
	filterGeneratorClass     _FilterGeneratorClass
	filterGeneratorClassOnce sync.Once
)

func getFilterGeneratorClass() _FilterGeneratorClass {
	filterGeneratorClassOnce.Do(func() {
		filterGeneratorClass = _FilterGeneratorClass{objc.GetClass("CIFilterGenerator")}
	})
	return filterGeneratorClass
}

type _FilterGeneratorClass struct {
	class objc.Class
}

// An interface definition for the [FilterGenerator] class.
type IFilterGenerator interface {
	objectivec.IObject
	ConnectObjectWithKeyToObjectWithKey(sourceObject objc.ID, sourceKey string, targetObject objc.ID, targetKey string)
	DisconnectObjectWithKeyToObjectWithKey(sourceObject objc.ID, sourceKey string, targetObject objc.ID, targetKey string)
	ExportKeyFromObjectWithName(key string, targetObject objc.ID, exportedKeyName string)
	Filter() unsafe.Pointer
	RegisterFilterName(name string)
	RemoveExportedKey(exportedKeyName string)
	SetAttributesForExportedKey(attributes unsafe.Pointer, key string)
	WriteToURLAtomically(aURL unsafe.Pointer, flag bool) bool
}

// An object that creates and configures chains of individual image filters. [Full Topic]
//
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


// Initializes a filter generator object with the contents of a filter generator file. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIFilterGenerator/init(contentsOf:)
func NewFilterGeneratorWithContentsOfURL(aURL unsafe.Pointer) FilterGenerator {
	instance := getFilterGeneratorClass().Alloc()
	rv := objc.Send[FilterGenerator](instance.ID, objc.Sel("initWithContentsOfURL:"), aURL)
	rv.Autorelease()
	return rv
}


// Creates and returns an empty filter generator object. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIFilterGenerator/filterGenerator
func (fc _FilterGeneratorClass) FilterGenerator() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(fc.class), objc.Sel("filterGenerator"))
	return rv
}
// Creates and returns a filter generator object and initializes it with the contents of a filter generator file. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIFilterGenerator/filterGeneratorWithContentsOfURL:
func (fc _FilterGeneratorClass) FilterGeneratorWithContentsOfURL(aURL unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(fc.class), objc.Sel("filterGeneratorWithContentsOfURL:"), aURL)
	return rv
}
// Adds an object to the filter chain. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIFilterGenerator/connect(_:withKey:to:withKey:)
func (f_ FilterGenerator) ConnectObjectWithKeyToObjectWithKey(sourceObject objc.ID, sourceKey string, targetObject objc.ID, targetKey string) {
	objc.Send[objc.ID](f_.ID, objc.Sel("connectObject:withKey:toObject:withKey:"), sourceObject, objc.String(sourceKey), targetObject, objc.String(targetKey))
}
// Removes the connection between two objects in the filter chain. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIFilterGenerator/disconnectObject(_:withKey:to:withKey:)
func (f_ FilterGenerator) DisconnectObjectWithKeyToObjectWithKey(sourceObject objc.ID, sourceKey string, targetObject objc.ID, targetKey string) {
	objc.Send[objc.ID](f_.ID, objc.Sel("disconnectObject:withKey:toObject:withKey:"), sourceObject, objc.String(sourceKey), targetObject, objc.String(targetKey))
}
// Exports an input or output key of an object in the filter chain. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIFilterGenerator/exportKey(_:from:withName:)
func (f_ FilterGenerator) ExportKeyFromObjectWithName(key string, targetObject objc.ID, exportedKeyName string) {
	objc.Send[objc.ID](f_.ID, objc.Sel("exportKey:fromObject:withName:"), objc.String(key), targetObject, objc.String(exportedKeyName))
}
// Creates a filter object based on the filter chain. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIFilterGenerator/filter()
func (f_ FilterGenerator) Filter() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](f_.ID, objc.Sel("filter"))
	return rv
}
// Registers the name associated with a filter chain. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIFilterGenerator/registerFilterName(_:)
func (f_ FilterGenerator) RegisterFilterName(name string) {
	objc.Send[objc.ID](f_.ID, objc.Sel("registerFilterName:"), objc.String(name))
}
// Removes a key that was previously exported. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIFilterGenerator/removeExportedKey(_:)
func (f_ FilterGenerator) RemoveExportedKey(exportedKeyName string) {
	objc.Send[objc.ID](f_.ID, objc.Sel("removeExportedKey:"), objc.String(exportedKeyName))
}
// Sets a dictionary of attributes for an exported key. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIFilterGenerator/setAttributes(_:forExportedKey:)
func (f_ FilterGenerator) SetAttributesForExportedKey(attributes unsafe.Pointer, key string) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setAttributes:forExportedKey:"), attributes, objc.String(key))
}
// Archives a filter generator object to a filter generator file. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIFilterGenerator/write(to:atomically:)
func (f_ FilterGenerator) WriteToURLAtomically(aURL unsafe.Pointer, flag bool) bool {
	rv := objc.Send[bool](f_.ID, objc.Sel("writeToURL:atomically:"), aURL, flag)
	return rv
}

