// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
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
	

	// properties:
	Entries() IMLModelCollectionEntry
	SetEntries(value IMLModelCollectionEntry)


	

	// methods:


}





// Alloc allocates a new instance without initialization.
func (mc _ModelCollectionEntryClass) Alloc() ModelCollectionEntry {
	rv := objc.Send[ModelCollectionEntry](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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







