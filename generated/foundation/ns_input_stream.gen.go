// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class NSInputStream */


/* debug [class_header]: Header for NSInputStream */
// The class instance for the [InputStream] class.
var (
	InputStreamClass     _InputStreamClass
	InputStreamClassOnce sync.Once
)

func getInputStreamClass() _InputStreamClass {
	InputStreamClassOnce.Do(func() {
		InputStreamClass = _InputStreamClass{objc.GetClass("NSInputStream")}
	})
	return InputStreamClass
}

type _InputStreamClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for InputStream */
// An interface definition for the [InputStream] class.
type IInputStream interface {
	IStream
	
/* debug [class_interface_properties]: Properties for InputStream */
	// properties:
	HasBytesAvailable() bool
	SetHasBytesAvailable(value bool)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for InputStream */
	// methods:
	ReadMaxLength(buffer objectivec.IObject, len_ uint) int
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for InputStream */
// Alloc allocates a new instance without initialization.
func (ic _InputStreamClass) Alloc() InputStream {
	rv := objc.Send[InputStream](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (ic _InputStreamClass) New() InputStream {
	rv := objc.Send[InputStream](objc.ID(ic.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (i_ InputStream) Init() InputStream {
	rv := objc.Send[InputStream](i_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i_ InputStream) Autorelease() InputStream {
	rv := objc.Send[InputStream](i_.ID, objc.Sel("autorelease"))
	return rv
}

// NewInputStream creates a new InputStream instance.
func NewInputStream() InputStream {
	return getInputStreamClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for InputStream */
// A stream that provides read-only stream functionality.
//
// is “toll-free bridged” with its Core Foundation counterpart, . For more information on toll-free bridging, see .


// A stream that provides read-only stream functionality.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/InputStream
type InputStream struct {
	Stream
}

// InputStreamFrom constructs a [InputStream] from an unsafe.Pointer.
//
// A stream that provides read-only stream functionality.
func InputStreamFrom(ptr unsafe.Pointer) InputStream {
	return InputStream{
		Stream: StreamFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for InputStream *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for InputStream */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for InputStream */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for InputStream */

// Reads up to a given number of bytes into a given buffer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/InputStream/read(_:maxLength:)
func (i_ InputStream) ReadMaxLength(buffer objectivec.IObject, len_ uint) int {
	rv := objc.Send[int](i_.ID, objc.Sel("read:maxLength:"), buffer, len_)
	return rv
}/* debug [instance_methods/method]: ReadMaxLength */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for InputStream */

// A Boolean value that indicates whether the receiver has bytes available to read.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/inputstream/hasbytesavailable
func (i_ InputStream) HasBytesAvailable() bool {
	rv := objc.Send[bool](i_.ID, objc.Sel("hasBytesAvailable"))
	return rv
}/* debug [instance_properties/getter]: hasBytesAvailable */


// A Boolean value that indicates whether the receiver has bytes available to read.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/inputstream/hasbytesavailable
func (i_ InputStream) SetHasBytesAvailable(value bool) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setHasBytesAvailable:"), value)
}/* debug [instance_properties/setter]: hasBytesAvailable */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NSInputStream */



