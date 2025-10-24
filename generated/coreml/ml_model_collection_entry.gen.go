// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MLModelCollectionEntry */


/* debug [class_header]: Header for MLModelCollectionEntry */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for ModelCollectionEntry */
// An interface definition for the [ModelCollectionEntry] class.
type IModelCollectionEntry interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for ModelCollectionEntry */
	// properties:
	Entries() IMLModelCollectionEntry
	SetEntries(value IMLModelCollectionEntry)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for ModelCollectionEntry */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for ModelCollectionEntry */
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for ModelCollectionEntry */
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for ModelCollectionEntry *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for ModelCollectionEntry */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for ModelCollectionEntry */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for ModelCollectionEntry */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for ModelCollectionEntry */

// A dictionary of model entries keyed to the models’ identifiers.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlmodelcollection/entries
func (m_ ModelCollectionEntry) Entries() IMLModelCollectionEntry {
	rv := objc.Send[ModelCollectionEntry](m_.ID, objc.Sel("entries"))
	return rv
}/* debug [instance_properties/getter]: entries */


// A dictionary of model entries keyed to the models’ identifiers.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlmodelcollection/entries
func (m_ ModelCollectionEntry) SetEntries(value IMLModelCollectionEntry) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setEntries:"), value)
}/* debug [instance_properties/setter]: entries */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MLModelCollectionEntry */


