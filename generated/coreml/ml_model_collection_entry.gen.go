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
	ModelIdentifier() string
	SetModelIdentifier(value string)
	ModelURL() foundation.URL
	SetModelURL(value foundation.URL)
	Entries() IMLModelCollectionEntry
	SetEntries(value IMLModelCollectionEntry)
}

// A model and its identifier within a model collection.


// A model and its identifier within a model collection.
//
// [Full Topic]
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



// The name of the model, which is unique to the collection.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlmodelcollection/entry/modelidentifier
func (m_ ModelCollectionEntry) ModelIdentifier() string {
	rv := objc.Send[string](m_.ID, objc.Sel("modelIdentifier"))
	return rv
}


// The name of the model, which is unique to the collection.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlmodelcollection/entry/modelidentifier
func (m_ ModelCollectionEntry) SetModelIdentifier(value string) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setModelIdentifier:"), objc.String(value))
}


// The compiled model’s location on the device’s file system.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlmodelcollection/entry/modelurl
func (m_ ModelCollectionEntry) ModelURL() foundation.URL {
	rv := objc.Send[foundation.URL](m_.ID, objc.Sel("modelURL"))
	return rv
}


// The compiled model’s location on the device’s file system.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlmodelcollection/entry/modelurl
func (m_ ModelCollectionEntry) SetModelURL(value foundation.URL) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setModelURL:"), value)
}


// A dictionary of model entries keyed to the models’ identifiers.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlmodelcollection/entries
func (m_ ModelCollectionEntry) Entries() IMLModelCollectionEntry {
	rv := objc.Send[ModelCollectionEntry](m_.ID, objc.Sel("entries"))
	return rv
}


// A dictionary of model entries keyed to the models’ identifiers.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlmodelcollection/entries
func (m_ ModelCollectionEntry) SetEntries(value IMLModelCollectionEntry) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setEntries:"), value)
}



