// Code generated from Apple documentation for GameController. DO NOT EDIT.

package gamecontroller

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/coreml"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class GCPhysicalInputElementCollection */


/* debug [class_header]: Header for GCPhysicalInputElementCollection */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for GCPhysicalInputElementCollection */
// An interface definition for the [GCPhysicalInputElementCollection] class.
type IGCPhysicalInputElementCollection interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for GCPhysicalInputElementCollection */
	// properties:
	Count() uint
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for GCPhysicalInputElementCollection */
	// methods:
	ElementEnumerator() unsafe.Pointer
	ElementForAlias(alias coreml.Key) unsafe.Pointer
	ObjectForKeyedSubscript(key coreml.Key) unsafe.Pointer
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for GCPhysicalInputElementCollection */
// Alloc allocates a new instance without initialization.
func (gc _GCPhysicalInputElementCollectionClass) Alloc() GCPhysicalInputElementCollection {
	rv := objc.Send[GCPhysicalInputElementCollection](objc.ID(gc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for GCPhysicalInputElementCollection */
// A collection of physical input elements.


// A collection of physical input elements.
//
// [Full Topic]
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for GCPhysicalInputElementCollection *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for GCPhysicalInputElementCollection */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for GCPhysicalInputElementCollection */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for GCPhysicalInputElementCollection */

// Returns an enumerator to iterate the elements in the collection.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCPhysicalInputElementCollection-c.class/elementEnumerator
func (g_ GCPhysicalInputElementCollection) ElementEnumerator() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g_.ID, objc.Sel("elementEnumerator"))
	return rv
}/* debug [instance_methods/method]: ElementEnumerator */


// Returns the element in the collection that uses the specified alias.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCPhysicalInputElementCollection-c.class/elementForAlias:
func (g_ GCPhysicalInputElementCollection) ElementForAlias(alias coreml.Key) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g_.ID, objc.Sel("elementForAlias:"), alias)
	return rv
}/* debug [instance_methods/method]: ElementForAlias */


// Returns the element in the collection for the specified key.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCPhysicalInputElementCollection-c.class/objectForKeyedSubscript:
func (g_ GCPhysicalInputElementCollection) ObjectForKeyedSubscript(key coreml.Key) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g_.ID, objc.Sel("objectForKeyedSubscript:"), key)
	return rv
}/* debug [instance_methods/method]: ObjectForKeyedSubscript */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for GCPhysicalInputElementCollection */

// The number of elements in the collection.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCPhysicalInputElementCollection-c.class/count
func (g_ GCPhysicalInputElementCollection) Count() uint {
	rv := objc.Send[uint](g_.ID, objc.Sel("count"))
	return rv
}/* debug [instance_properties/getter]: count */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class GCPhysicalInputElementCollection */



