// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [waitCount] class.
var (
	WaitCountClass     _waitCountClass
	WaitCountClassOnce sync.Once
)

func getwaitCountClass() _waitCountClass {
	WaitCountClassOnce.Do(func() {
		WaitCountClass = _waitCountClass{objc.GetClass("waitCount")}
	})
	return WaitCountClass
}

type _waitCountClass struct {
	class objc.Class
}

// An interface definition for the [waitCount] class.
type IwaitCount interface {
	objectivec.IObject
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSConnection/waitCount

type waitCount struct {
	objectivec.Object
}

// waitCountFrom constructs a [waitCount] from an unsafe.Pointer.
func waitCountFrom(ptr unsafe.Pointer) waitCount {
	return waitCount{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (wc _waitCountClass) Alloc() waitCount {
	rv := objc.Send[waitCount](objc.ID(wc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (wc _waitCountClass) New() waitCount {
	rv := objc.Send[waitCount](objc.ID(wc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (w_ waitCount) Init() waitCount {
	rv := objc.Send[waitCount](w_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (w_ waitCount) Autorelease() waitCount {
	rv := objc.Send[waitCount](w_.ID, objc.Sel("autorelease"))
	return rv
}

// NewwaitCount creates a new waitCount instance.
func NewwaitCount() waitCount {
	return getwaitCountClass().New()
}




