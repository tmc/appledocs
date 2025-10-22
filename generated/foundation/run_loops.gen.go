// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [runLoops] class.
var (
	RunLoopsClass     _runLoopsClass
	RunLoopsClassOnce sync.Once
)

func getrunLoopsClass() _runLoopsClass {
	RunLoopsClassOnce.Do(func() {
		RunLoopsClass = _runLoopsClass{objc.GetClass("runLoops")}
	})
	return RunLoopsClass
}

type _runLoopsClass struct {
	class objc.Class
}

// An interface definition for the [runLoops] class.
type IrunLoops interface {
	objectivec.IObject
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSConnection/runLoops

type runLoops struct {
	objectivec.Object
}

// runLoopsFrom constructs a [runLoops] from an unsafe.Pointer.
func runLoopsFrom(ptr unsafe.Pointer) runLoops {
	return runLoops{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (rc _runLoopsClass) Alloc() runLoops {
	rv := objc.Send[runLoops](objc.ID(rc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (rc _runLoopsClass) New() runLoops {
	rv := objc.Send[runLoops](objc.ID(rc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (r_ runLoops) Init() runLoops {
	rv := objc.Send[runLoops](r_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (r_ runLoops) Autorelease() runLoops {
	rv := objc.Send[runLoops](r_.ID, objc.Sel("autorelease"))
	return rv
}

// NewrunLoops creates a new runLoops instance.
func NewrunLoops() runLoops {
	return getrunLoopsClass().New()
}




