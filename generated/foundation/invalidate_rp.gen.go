// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [invalidateRP] class.
var (
	InvalidateRPClass     _invalidateRPClass
	InvalidateRPClassOnce sync.Once
)

func getinvalidateRPClass() _invalidateRPClass {
	InvalidateRPClassOnce.Do(func() {
		InvalidateRPClass = _invalidateRPClass{objc.GetClass("invalidateRP")}
	})
	return InvalidateRPClass
}

type _invalidateRPClass struct {
	class objc.Class
}

// An interface definition for the [invalidateRP] class.
type IinvalidateRP interface {
	objectivec.IObject
	// properties:
	// methods:
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSConnection/invalidateRP
type invalidateRP struct {
	objectivec.Object
}

// invalidateRPFrom constructs a [invalidateRP] from an unsafe.Pointer.
func invalidateRPFrom(ptr unsafe.Pointer) invalidateRP {
	return invalidateRP{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (ic _invalidateRPClass) Alloc() invalidateRP {
	rv := objc.Send[invalidateRP](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ic _invalidateRPClass) New() invalidateRP {
	rv := objc.Send[invalidateRP](objc.ID(ic.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (i_ invalidateRP) Init() invalidateRP {
	rv := objc.Send[invalidateRP](i_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i_ invalidateRP) Autorelease() invalidateRP {
	rv := objc.Send[invalidateRP](i_.ID, objc.Sel("autorelease"))
	return rv
}

// NewinvalidateRP creates a new invalidateRP instance.
func NewinvalidateRP() invalidateRP {
	return getinvalidateRPClass().New()
}




