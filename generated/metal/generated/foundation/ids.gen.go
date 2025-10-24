// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [ids] class.
var (
	IdsClass     _idsClass
	IdsClassOnce sync.Once
)

func getidsClass() _idsClass {
	IdsClassOnce.Do(func() {
		IdsClass = _idsClass{objc.GetClass("ids")}
	})
	return IdsClass
}

type _idsClass struct {
	class objc.Class
}

// An interface definition for the [ids] class.
type Iids interface {
	objectivec.IObject
	// properties:
	// methods:
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSArchiver/ids
type ids struct {
	objectivec.Object
}

// idsFrom constructs a [ids] from an unsafe.Pointer.
func idsFrom(ptr unsafe.Pointer) ids {
	return ids{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (ic _idsClass) Alloc() ids {
	rv := objc.Send[ids](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ic _idsClass) New() ids {
	rv := objc.Send[ids](objc.ID(ic.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (i_ ids) Init() ids {
	rv := objc.Send[ids](i_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i_ ids) Autorelease() ids {
	rv := objc.Send[ids](i_.ID, objc.Sel("autorelease"))
	return rv
}

// Newids creates a new ids instance.
func Newids() ids {
	return getidsClass().New()
}




