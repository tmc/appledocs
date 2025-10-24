// Code generated from Apple documentation for IOBluetooth. DO NOT EDIT.

package iobluetooth

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class sortedAttributes */


/* debug [class_header]: Header for sortedAttributes */
// The class instance for the [sortedAttributes] class.
var (
	SortedAttributesClass     _sortedAttributesClass
	SortedAttributesClassOnce sync.Once
)

func getsortedAttributesClass() _sortedAttributesClass {
	SortedAttributesClassOnce.Do(func() {
		SortedAttributesClass = _sortedAttributesClass{objc.GetClass("sortedAttributes")}
	})
	return SortedAttributesClass
}

type _sortedAttributesClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for sortedAttributes */
// An interface definition for the [sortedAttributes] class.
type IsortedAttributes interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for sortedAttributes */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for sortedAttributes */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for sortedAttributes */
// Alloc allocates a new instance without initialization.
func (sc _sortedAttributesClass) Alloc() sortedAttributes {
	rv := objc.Send[sortedAttributes](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (sc _sortedAttributesClass) New() sortedAttributes {
	rv := objc.Send[sortedAttributes](objc.ID(sc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (s_ sortedAttributes) Init() sortedAttributes {
	rv := objc.Send[sortedAttributes](s_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s_ sortedAttributes) Autorelease() sortedAttributes {
	rv := objc.Send[sortedAttributes](s_.ID, objc.Sel("autorelease"))
	return rv
}

// NewsortedAttributes creates a new sortedAttributes instance.
func NewsortedAttributes() sortedAttributes {
	return getsortedAttributesClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for sortedAttributes */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothSDPServiceRecord/sortedAttributes-c.ivar
type sortedAttributes struct {
	objectivec.Object
}

// sortedAttributesFrom constructs a [sortedAttributes] from an unsafe.Pointer.
func sortedAttributesFrom(ptr unsafe.Pointer) sortedAttributes {
	return sortedAttributes{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for sortedAttributes *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for sortedAttributes */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for sortedAttributes */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for sortedAttributes */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for sortedAttributes */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class sortedAttributes */



