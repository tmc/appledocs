// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [isDead] class.
var (
	IsDeadClass     _isDeadClass
	IsDeadClassOnce sync.Once
)

func getisDeadClass() _isDeadClass {
	IsDeadClassOnce.Do(func() {
		IsDeadClass = _isDeadClass{objc.GetClass("isDead")}
	})
	return IsDeadClass
}

type _isDeadClass struct {
	class objc.Class
}

// An interface definition for the [isDead] class.
type IisDead interface {
	objectivec.IObject
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSConnection/isDead
type isDead struct {
	objectivec.Object
}

// isDeadFrom constructs a [isDead] from an unsafe.Pointer.
func isDeadFrom(ptr unsafe.Pointer) isDead {
	return isDead{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (ic _isDeadClass) Alloc() isDead {
	rv := objc.Send[isDead](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ic _isDeadClass) New() isDead {
	rv := objc.Send[isDead](objc.ID(ic.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (i_ isDead) Init() isDead {
	rv := objc.Send[isDead](i_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i_ isDead) Autorelease() isDead {
	rv := objc.Send[isDead](i_.ID, objc.Sel("autorelease"))
	return rv
}

// NewisDead creates a new isDead instance.
func NewisDead() isDead {
	return getisDeadClass().New()
}




