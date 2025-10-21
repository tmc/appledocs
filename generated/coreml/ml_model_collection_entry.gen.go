// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [ModelCollectionEntry] class.
var (
	ModelCollectionEntryClass     _ModelCollectionEntryClass
	ModelCollectionEntryClassOnce sync.Once
)

func getModelCollectionEntryClass() _ModelCollectionEntryClass {
	ModelCollectionEntryClassOnce.Do(func() {
		ModelCollectionEntryClass = _ModelCollectionEntryClass{objc.GetClass("MLModelCollectionEntry")}
	})
	return ModelCollectionEntryClass
}

type _ModelCollectionEntryClass struct {
	class objc.Class
}

// An interface definition for the [ModelCollectionEntry] class.
type IModelCollectionEntry interface {
	objectivec.IObject
	IsEqualToModelCollectionEntry(entry unsafe.Pointer) bool
}

// A model and its identifier within a model collection.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLModelCollection/Entry
type ModelCollectionEntry struct {
	objectivec.Object
}

// ModelCollectionEntryFrom constructs a [ModelCollectionEntry] from an unsafe.Pointer.
//
// A model and its identifier within a model collection.
func ModelCollectionEntryFrom(ptr unsafe.Pointer) ModelCollectionEntry {
	return ModelCollectionEntry{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _ModelCollectionEntryClass) Alloc() ModelCollectionEntry {
	rv := objc.Send[ModelCollectionEntry](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _ModelCollectionEntryClass) New() ModelCollectionEntry {
	rv := objc.Send[ModelCollectionEntry](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ ModelCollectionEntry) Init() ModelCollectionEntry {
	rv := objc.Send[ModelCollectionEntry](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ ModelCollectionEntry) Autorelease() ModelCollectionEntry {
	rv := objc.Send[ModelCollectionEntry](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewModelCollectionEntry creates a new ModelCollectionEntry instance.
func NewModelCollectionEntry() ModelCollectionEntry {
	return getModelCollectionEntryClass().New()
}


// Returns a Boolean value that indicates whether the two entries are equal.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLModelCollection/Entry/isEqual(to:)
func (m_ ModelCollectionEntry) IsEqualToModelCollectionEntry(entry unsafe.Pointer) bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("isEqualToModelCollectionEntry:"), entry)
	return rv
}

// The name of the model, which is unique to the collection.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLModelCollection/Entry/modelIdentifier
func (m_ ModelCollectionEntry) ModelIdentifier() string {
	rv := objc.Send[string](m_.ID, objc.Sel("modelIdentifier"))
	return rv
}

// The compiled model’s location on the device’s file system.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLModelCollection/Entry/modelURL
func (m_ ModelCollectionEntry) ModelURL() foundation.URL {
	rv := objc.Send[foundation.URL](m_.ID, objc.Sel("modelURL"))
	return rv
}

// A dictionary of model entries keyed to the models’ identifiers.
//
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlmodelcollection/entries
func (m_ ModelCollectionEntry) Entries() string {
	rv := objc.Send[string](m_.ID, objc.Sel("entries"))
	return rv
}


// SetEntries sets the value of the entries property.
// A dictionary of model entries keyed to the models’ identifiers.

//
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlmodelcollection/entries
func (m_ ModelCollectionEntry) SetEntries(value string) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setEntries:"), objc.String(value))
}



