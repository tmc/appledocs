// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class NSPurgeableData */


/* debug [class_header]: Header for NSPurgeableData */
// The class instance for the [PurgeableData] class.
var (
	PurgeableDataClass     _PurgeableDataClass
	PurgeableDataClassOnce sync.Once
)

func getPurgeableDataClass() _PurgeableDataClass {
	PurgeableDataClassOnce.Do(func() {
		PurgeableDataClass = _PurgeableDataClass{objc.GetClass("NSPurgeableData")}
	})
	return PurgeableDataClass
}

type _PurgeableDataClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for PurgeableData */
// An interface definition for the [PurgeableData] class.
type IPurgeableData interface {
	IMutableData
	
/* debug [class_interface_properties]: Properties for PurgeableData */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for PurgeableData */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for PurgeableData */
// Alloc allocates a new instance without initialization.
func (pc _PurgeableDataClass) Alloc() PurgeableData {
	rv := objc.Send[PurgeableData](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (pc _PurgeableDataClass) New() PurgeableData {
	rv := objc.Send[PurgeableData](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ PurgeableData) Init() PurgeableData {
	rv := objc.Send[PurgeableData](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ PurgeableData) Autorelease() PurgeableData {
	rv := objc.Send[PurgeableData](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewPurgeableData creates a new PurgeableData instance.
func NewPurgeableData() PurgeableData {
	return getPurgeableDataClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for PurgeableData */
// A mutable data object containing bytes that can be discarded when they’re no longer needed.
//
// objects inherit their creation methods from their superclass, while providing a default implementation of the protocol. All objects begin “accessed” to ensure that they are not instantly discarded. The method marks the object’s bytes as “accessed,” thus protecting them from being discarded, and must be called before accessing the object, or else an exception will be raised. This method returns if the bytes have not been discarded and if they have been successfully marked as “accessed”. Any method that directly or indirectly accesses these bytes or their length when they are not “accessed” will raise an exception. When you are done with the data, call to allow them to be discarded in order to quickly free up memory. You may use these objects by themselves, and do not necessarily have to use them in conjunction with to get the purging behavior. The class incorporates a caching mechanism with some auto-removal policies to ensure that its memory footprint does not get too large. objects should not be used as keys in hashing-based collections, because the value of the bytes pointer can change after every mutation of the data.


// A mutable data object containing bytes that can be discarded when they’re no longer needed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSPurgeableData
type PurgeableData struct {
	MutableData
}

// PurgeableDataFrom constructs a [PurgeableData] from an unsafe.Pointer.
//
// A mutable data object containing bytes that can be discarded when they’re no longer needed.
func PurgeableDataFrom(ptr unsafe.Pointer) PurgeableData {
	return PurgeableData{
		MutableData: MutableDataFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for PurgeableData *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for PurgeableData */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for PurgeableData */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for PurgeableData */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for PurgeableData */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NSPurgeableData */



