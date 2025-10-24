// Code generated from Apple documentation for FSKit. DO NOT EDIT.

package fskit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class FSMutableFileDataBuffer */


/* debug [class_header]: Header for FSMutableFileDataBuffer */
// The class instance for the [FSMutableFileDataBuffer] class.
var (
	FSMutableFileDataBufferClass     _FSMutableFileDataBufferClass
	FSMutableFileDataBufferClassOnce sync.Once
)

func getFSMutableFileDataBufferClass() _FSMutableFileDataBufferClass {
	FSMutableFileDataBufferClassOnce.Do(func() {
		FSMutableFileDataBufferClass = _FSMutableFileDataBufferClass{objc.GetClass("FSMutableFileDataBuffer")}
	})
	return FSMutableFileDataBufferClass
}

type _FSMutableFileDataBufferClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for FSMutableFileDataBuffer */
// An interface definition for the [FSMutableFileDataBuffer] class.
type IFSMutableFileDataBuffer interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for FSMutableFileDataBuffer */
	// properties:
	Length() uint
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for FSMutableFileDataBuffer */
	// methods:
	MutableBytes()
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for FSMutableFileDataBuffer */
// Alloc allocates a new instance without initialization.
func (fc _FSMutableFileDataBufferClass) Alloc() FSMutableFileDataBuffer {
	rv := objc.Send[FSMutableFileDataBuffer](objc.ID(fc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (fc _FSMutableFileDataBufferClass) New() FSMutableFileDataBuffer {
	rv := objc.Send[FSMutableFileDataBuffer](objc.ID(fc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (f_ FSMutableFileDataBuffer) Init() FSMutableFileDataBuffer {
	rv := objc.Send[FSMutableFileDataBuffer](f_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (f_ FSMutableFileDataBuffer) Autorelease() FSMutableFileDataBuffer {
	rv := objc.Send[FSMutableFileDataBuffer](f_.ID, objc.Sel("autorelease"))
	return rv
}

// NewFSMutableFileDataBuffer creates a new FSMutableFileDataBuffer instance.
func NewFSMutableFileDataBuffer() FSMutableFileDataBuffer {
	return getFSMutableFileDataBufferClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for FSMutableFileDataBuffer */
// A wrapper object for a data buffer.
//
// This object provides a “zero-copy” buffer, for use when reading data from files. By not requiring additional buffer copying, this object reduces the extension’s memory footprint and improves performance. The behaves similarly to a in the kernel.


// A wrapper object for a data buffer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/FSKit/FSMutableFileDataBuffer
type FSMutableFileDataBuffer struct {
	objectivec.Object
}

// FSMutableFileDataBufferFrom constructs a [FSMutableFileDataBuffer] from an unsafe.Pointer.
//
// A wrapper object for a data buffer.
func FSMutableFileDataBufferFrom(ptr unsafe.Pointer) FSMutableFileDataBuffer {
	return FSMutableFileDataBuffer{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for FSMutableFileDataBuffer *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for FSMutableFileDataBuffer */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for FSMutableFileDataBuffer */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for FSMutableFileDataBuffer */

// The byte data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/FSKit/FSMutableFileDataBuffer/mutableBytes
func (f_ FSMutableFileDataBuffer) MutableBytes() {
	objc.Send[objc.ID](f_.ID, objc.Sel("mutableBytes"))
}/* debug [instance_methods/method]: MutableBytes */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for FSMutableFileDataBuffer */

// The data length of the buffer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/FSKit/FSMutableFileDataBuffer/length
func (f_ FSMutableFileDataBuffer) Length() uint {
	rv := objc.Send[uint](f_.ID, objc.Sel("length"))
	return rv
}/* debug [instance_properties/getter]: length */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class FSMutableFileDataBuffer */



