// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class bytes */


/* debug [class_header]: Header for bytes */
// The class instance for the [bytes] class.
var (
	BytesClass     _bytesClass
	BytesClassOnce sync.Once
)

func getbytesClass() _bytesClass {
	BytesClassOnce.Do(func() {
		BytesClass = _bytesClass{objc.GetClass("bytes")}
	})
	return BytesClass
}

type _bytesClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for bytes */
// An interface definition for the [bytes] class.
type Ibytes interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for bytes */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for bytes */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for bytes */
// Alloc allocates a new instance without initialization.
func (bc _bytesClass) Alloc() bytes {
	rv := objc.Send[bytes](objc.ID(bc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (bc _bytesClass) New() bytes {
	rv := objc.Send[bytes](objc.ID(bc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (b_ bytes) Init() bytes {
	rv := objc.Send[bytes](b_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (b_ bytes) Autorelease() bytes {
	rv := objc.Send[bytes](b_.ID, objc.Sel("autorelease"))
	return rv
}

// Newbytes creates a new bytes instance.
func Newbytes() bytes {
	return getbytesClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for bytes */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSSimpleCString/bytes
type bytes struct {
	objectivec.Object
}

// bytesFrom constructs a [bytes] from an unsafe.Pointer.
func bytesFrom(ptr unsafe.Pointer) bytes {
	return bytes{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for bytes *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for bytes */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for bytes */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for bytes */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for bytes */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class bytes */



