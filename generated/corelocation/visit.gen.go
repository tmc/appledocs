// Code generated from Apple documentation for CoreLocation. DO NOT EDIT.

package corelocation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [Visit] class.
var (
	visitClass     _VisitClass
	visitClassOnce sync.Once
)

func getVisitClass() _VisitClass {
	visitClassOnce.Do(func() {
		visitClass = _VisitClass{objc.GetClass("CLVisit")}
	})
	return visitClass
}

type _VisitClass struct {
	class objc.Class
}

// An interface definition for the [Visit] class.
type IVisit interface {
	objectivec.IObject
}

// Information about the user’s location during a specific period of time. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLVisit
type Visit struct {
	objectivec.Object
}

// VisitFrom constructs a [Visit] from an unsafe.Pointer.
//
// Information about the user’s location during a specific period of time.
func VisitFrom(ptr unsafe.Pointer) Visit {
	return Visit{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (vc _VisitClass) Alloc() Visit {
	rv := objc.Send[Visit](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (vc _VisitClass) New() Visit {
	rv := objc.Send[Visit](objc.ID(vc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (v_ Visit) Init() Visit {
	rv := objc.Send[Visit](v_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v_ Visit) Autorelease() Visit {
	rv := objc.Send[Visit](v_.ID, objc.Sel("autorelease"))
	return rv
}

// NewVisit creates a new Visit instance.
func NewVisit() Visit {
	return getVisitClass().New()
}




