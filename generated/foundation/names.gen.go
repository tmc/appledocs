// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [names] class.
var (
	NamesClass     _namesClass
	NamesClassOnce sync.Once
)

func getnamesClass() _namesClass {
	NamesClassOnce.Do(func() {
		NamesClass = _namesClass{objc.GetClass("names")}
	})
	return NamesClass
}

type _namesClass struct {
	class objc.Class
}

// An interface definition for the [names] class.
type Inames interface {
	objectivec.IObject
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSHost/names-c.ivar

type names struct {
	objectivec.Object
}

// namesFrom constructs a [names] from an unsafe.Pointer.
func namesFrom(ptr unsafe.Pointer) names {
	return names{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (nc _namesClass) Alloc() names {
	rv := objc.Send[names](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (nc _namesClass) New() names {
	rv := objc.Send[names](objc.ID(nc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (n_ names) Init() names {
	rv := objc.Send[names](n_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n_ names) Autorelease() names {
	rv := objc.Send[names](n_.ID, objc.Sel("autorelease"))
	return rv
}

// Newnames creates a new names instance.
func Newnames() names {
	return getnamesClass().New()
}




