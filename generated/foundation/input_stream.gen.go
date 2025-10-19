// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [InputStream] class.
var (
	inputStreamClass     _InputStreamClass
	inputStreamClassOnce sync.Once
)

func getInputStreamClass() _InputStreamClass {
	inputStreamClassOnce.Do(func() {
		inputStreamClass = _InputStreamClass{objc.GetClass("NSInputStream")}
	})
	return inputStreamClass
}

type _InputStreamClass struct {
	class objc.Class
}

// An interface definition for the [InputStream] class.
type IInputStream interface {
	IStream
}

// A stream that provides read-only stream functionality. [Full Topic]
//
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

// Alloc allocates a new instance without initialization.
func (ic _InputStreamClass) Alloc() InputStream {
	rv := objc.Send[InputStream](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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




