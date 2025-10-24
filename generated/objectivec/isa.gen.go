// Code generated from Apple documentation for ObjectiveC. DO NOT EDIT.

package objectivec

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [isa] class.
var (
	IsaClass     _isaClass
	IsaClassOnce sync.Once
)

func getisaClass() _isaClass {
	IsaClassOnce.Do(func() {
		IsaClass = _isaClass{objc.GetClass("isa")}
	})
	return IsaClass
}

type _isaClass struct {
	class objc.Class
}

// An interface definition for the [isa] class.
type Iisa interface {
	IObject
	// properties:
	// methods:
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/isa
type isa struct {
	Object
}

// isaFrom constructs a [isa] from an unsafe.Pointer.
func isaFrom(ptr unsafe.Pointer) isa {
	return isa{Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (ic _isaClass) Alloc() isa {
	rv := objc.Send[isa](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ic _isaClass) New() isa {
	rv := objc.Send[isa](objc.ID(ic.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (i_ isa) Init() isa {
	rv := objc.Send[isa](i_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i_ isa) Autorelease() isa {
	rv := objc.Send[isa](i_.ID, objc.Sel("autorelease"))
	return rv
}

// Newisa creates a new isa instance.
func Newisa() isa {
	return getisaClass().New()
}




