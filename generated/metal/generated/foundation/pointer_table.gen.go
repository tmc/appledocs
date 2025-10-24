// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [pointerTable] class.
var (
	PointerTableClass     _pointerTableClass
	PointerTableClassOnce sync.Once
)

func getpointerTableClass() _pointerTableClass {
	PointerTableClassOnce.Do(func() {
		PointerTableClass = _pointerTableClass{objc.GetClass("pointerTable")}
	})
	return PointerTableClass
}

type _pointerTableClass struct {
	class objc.Class
}

// An interface definition for the [pointerTable] class.
type IpointerTable interface {
	objectivec.IObject
	// properties:
	// methods:
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSArchiver/pointerTable
type pointerTable struct {
	objectivec.Object
}

// pointerTableFrom constructs a [pointerTable] from an unsafe.Pointer.
func pointerTableFrom(ptr unsafe.Pointer) pointerTable {
	return pointerTable{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (pc _pointerTableClass) Alloc() pointerTable {
	rv := objc.Send[pointerTable](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (pc _pointerTableClass) New() pointerTable {
	rv := objc.Send[pointerTable](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ pointerTable) Init() pointerTable {
	rv := objc.Send[pointerTable](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ pointerTable) Autorelease() pointerTable {
	rv := objc.Send[pointerTable](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewpointerTable creates a new pointerTable instance.
func NewpointerTable() pointerTable {
	return getpointerTableClass().New()
}




