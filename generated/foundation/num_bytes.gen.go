// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class numBytes */


/* debug [class_header]: Header for numBytes */
// The class instance for the [numBytes] class.
var (
	NumBytesClass     _numBytesClass
	NumBytesClassOnce sync.Once
)

func getnumBytesClass() _numBytesClass {
	NumBytesClassOnce.Do(func() {
		NumBytesClass = _numBytesClass{objc.GetClass("numBytes")}
	})
	return NumBytesClass
}

type _numBytesClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for numBytes */
// An interface definition for the [numBytes] class.
type InumBytes interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for numBytes */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for numBytes */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for numBytes */
// Alloc allocates a new instance without initialization.
func (nc _numBytesClass) Alloc() numBytes {
	rv := objc.Send[numBytes](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (nc _numBytesClass) New() numBytes {
	rv := objc.Send[numBytes](objc.ID(nc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (n_ numBytes) Init() numBytes {
	rv := objc.Send[numBytes](n_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n_ numBytes) Autorelease() numBytes {
	rv := objc.Send[numBytes](n_.ID, objc.Sel("autorelease"))
	return rv
}

// NewnumBytes creates a new numBytes instance.
func NewnumBytes() numBytes {
	return getnumBytesClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for numBytes */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSSimpleCString/numBytes
type numBytes struct {
	objectivec.Object
}

// numBytesFrom constructs a [numBytes] from an unsafe.Pointer.
func numBytesFrom(ptr unsafe.Pointer) numBytes {
	return numBytes{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for numBytes *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for numBytes */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for numBytes */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for numBytes */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for numBytes */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class numBytes */



