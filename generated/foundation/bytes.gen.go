// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)





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





// An interface definition for the [bytes] class.
type Ibytes interface {
	objectivec.IObject
	

	// properties:


	

	// methods:


}





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







// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSSimpleCString/bytes
type bytes struct {
	objectivec.Object
}

// bytesFrom constructs a [bytes] from an unsafe.Pointer.
func bytesFrom(ptr unsafe.Pointer) bytes {
	return bytes{objectivec.Object{objc.ID(ptr)}}
}































