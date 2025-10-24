// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)





// The class instance for the [OutputStream] class.
var (
	OutputStreamClass     _OutputStreamClass
	OutputStreamClassOnce sync.Once
)

func getOutputStreamClass() _OutputStreamClass {
	OutputStreamClassOnce.Do(func() {
		OutputStreamClass = _OutputStreamClass{objc.GetClass("NSOutputStream")}
	})
	return OutputStreamClass
}

type _OutputStreamClass struct {
	class objc.Class
}





// An interface definition for the [OutputStream] class.
type IOutputStream interface {
	IStream
	

	// properties:
	HasSpaceAvailable() bool
	SetHasSpaceAvailable(value bool)


	

	// methods:
	WriteMaxLength(buffer objectivec.IObject, len_ uint) int


}





// Alloc allocates a new instance without initialization.
func (oc _OutputStreamClass) Alloc() OutputStream {
	rv := objc.Send[OutputStream](objc.ID(oc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (oc _OutputStreamClass) New() OutputStream {
	rv := objc.Send[OutputStream](objc.ID(oc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (o_ OutputStream) Init() OutputStream {
	rv := objc.Send[OutputStream](o_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (o_ OutputStream) Autorelease() OutputStream {
	rv := objc.Send[OutputStream](o_.ID, objc.Sel("autorelease"))
	return rv
}

// NewOutputStream creates a new OutputStream instance.
func NewOutputStream() OutputStream {
	return getOutputStreamClass().New()
}





// A stream that provides write-only stream functionality.
//
// is “toll-free bridged” with its Core Foundation counterpart, . For more information on toll-free bridging, see .


// A stream that provides write-only stream functionality.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/OutputStream
type OutputStream struct {
	Stream
}

// OutputStreamFrom constructs a [OutputStream] from an unsafe.Pointer.
//
// A stream that provides write-only stream functionality.
func OutputStreamFrom(ptr unsafe.Pointer) OutputStream {
	return OutputStream{
		Stream: StreamFrom(ptr),
	}
}




















// Writes the contents of a provided data buffer to the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/OutputStream/write(_:maxLength:)
func (o_ OutputStream) WriteMaxLength(buffer objectivec.IObject, len_ uint) int {
	rv := objc.Send[int](o_.ID, objc.Sel("write:maxLength:"), buffer, len_)
	return rv
}







// A boolean value that indicates whether the receiver can be written to.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/outputstream/hasspaceavailable
func (o_ OutputStream) HasSpaceAvailable() bool {
	rv := objc.Send[bool](o_.ID, objc.Sel("hasSpaceAvailable"))
	return rv
}


// A boolean value that indicates whether the receiver can be written to.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/outputstream/hasspaceavailable
func (o_ OutputStream) SetHasSpaceAvailable(value bool) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setHasSpaceAvailable:"), value)
}








