// Code generated from Apple documentation for GameController. DO NOT EDIT.

package gamecontroller

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/coreml"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [GCPhysicalInputElementCollection] class.
var (
	GCPhysicalInputElementCollectionClass     _GCPhysicalInputElementCollectionClass
	GCPhysicalInputElementCollectionClassOnce sync.Once
)

func getGCPhysicalInputElementCollectionClass() _GCPhysicalInputElementCollectionClass {
	GCPhysicalInputElementCollectionClassOnce.Do(func() {
		GCPhysicalInputElementCollectionClass = _GCPhysicalInputElementCollectionClass{objc.GetClass("GCPhysicalInputElementCollection")}
	})
	return GCPhysicalInputElementCollectionClass
}

type _GCPhysicalInputElementCollectionClass struct {
	class objc.Class
}

// An interface definition for the [GCPhysicalInputElementCollection] class.
type IGCPhysicalInputElementCollection interface {
	objectivec.IObject
	ElementForAlias(alias coreml.IKey) unsafe.Pointer
}

// A collection of physical input elements.
//
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCPhysicalInputElementCollection-c.class
type GCPhysicalInputElementCollection struct {
	objectivec.Object
}

// GCPhysicalInputElementCollectionFrom constructs a [GCPhysicalInputElementCollection] from an unsafe.Pointer.
//
// A collection of physical input elements.
func GCPhysicalInputElementCollectionFrom(ptr unsafe.Pointer) GCPhysicalInputElementCollection {
	return GCPhysicalInputElementCollection{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (gc _GCPhysicalInputElementCollectionClass) Alloc() GCPhysicalInputElementCollection {
	rv := objc.Send[GCPhysicalInputElementCollection](objc.ID(gc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (gc _GCPhysicalInputElementCollectionClass) New() GCPhysicalInputElementCollection {
	rv := objc.Send[GCPhysicalInputElementCollection](objc.ID(gc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (g_ GCPhysicalInputElementCollection) Init() GCPhysicalInputElementCollection {
	rv := objc.Send[GCPhysicalInputElementCollection](g_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (g_ GCPhysicalInputElementCollection) Autorelease() GCPhysicalInputElementCollection {
	rv := objc.Send[GCPhysicalInputElementCollection](g_.ID, objc.Sel("autorelease"))
	return rv
}

// NewGCPhysicalInputElementCollection creates a new GCPhysicalInputElementCollection instance.
func NewGCPhysicalInputElementCollection() GCPhysicalInputElementCollection {
	return getGCPhysicalInputElementCollectionClass().New()
}


// Returns the element in the collection that uses the specified alias.
//
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCPhysicalInputElementCollection-c.class/elementForAlias:
func (g_ GCPhysicalInputElementCollection) ElementForAlias(alias coreml.IKey) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g_.ID, objc.Sel("elementForAlias:"), alias)
	return rv
}



