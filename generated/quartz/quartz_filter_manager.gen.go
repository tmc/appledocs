// Code generated from Apple documentation for Quartz. DO NOT EDIT.

package quartz

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [QuartzFilterManager] class.
var (
	QuartzFilterManagerClass     _QuartzFilterManagerClass
	QuartzFilterManagerClassOnce sync.Once
)

func getQuartzFilterManagerClass() _QuartzFilterManagerClass {
	QuartzFilterManagerClassOnce.Do(func() {
		QuartzFilterManagerClass = _QuartzFilterManagerClass{objc.GetClass("QuartzFilterManager")}
	})
	return QuartzFilterManagerClass
}

type _QuartzFilterManagerClass struct {
	class objc.Class
}

// An interface definition for the [QuartzFilterManager] class.
type IQuartzFilterManager interface {
	objectivec.IObject
	// properties:
	// methods:
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Quartz/QuartzFilterManager
type QuartzFilterManager struct {
	objectivec.Object
}

// QuartzFilterManagerFrom constructs a [QuartzFilterManager] from an unsafe.Pointer.
func QuartzFilterManagerFrom(ptr unsafe.Pointer) QuartzFilterManager {
	return QuartzFilterManager{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (qc _QuartzFilterManagerClass) Alloc() QuartzFilterManager {
	rv := objc.Send[QuartzFilterManager](objc.ID(qc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (qc _QuartzFilterManagerClass) New() QuartzFilterManager {
	rv := objc.Send[QuartzFilterManager](objc.ID(qc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (q_ QuartzFilterManager) Init() QuartzFilterManager {
	rv := objc.Send[QuartzFilterManager](q_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (q_ QuartzFilterManager) Autorelease() QuartzFilterManager {
	rv := objc.Send[QuartzFilterManager](q_.ID, objc.Sel("autorelease"))
	return rv
}

// NewQuartzFilterManager creates a new QuartzFilterManager instance.
func NewQuartzFilterManager() QuartzFilterManager {
	return getQuartzFilterManagerClass().New()
}




