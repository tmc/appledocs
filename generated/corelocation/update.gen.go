// Code generated from Apple documentation for CoreLocation. DO NOT EDIT.

package corelocation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [Update] class.
var (
	updateClass     _UpdateClass
	updateClassOnce sync.Once
)

func getUpdateClass() _UpdateClass {
	updateClassOnce.Do(func() {
		updateClass = _UpdateClass{objc.GetClass("CLUpdate")}
	})
	return updateClass
}

type _UpdateClass struct {
	class objc.Class
}

// An interface definition for the [Update] class.
type IUpdate interface {
	objectivec.IObject
}

// An object that represents a location update. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLUpdate
type Update struct {
	objectivec.Object
}

// UpdateFrom constructs a [Update] from an unsafe.Pointer.
//
// An object that represents a location update.
func UpdateFrom(ptr unsafe.Pointer) Update {
	return Update{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (uc _UpdateClass) Alloc() Update {
	rv := objc.Send[Update](objc.ID(uc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (uc _UpdateClass) New() Update {
	rv := objc.Send[Update](objc.ID(uc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (u_ Update) Init() Update {
	rv := objc.Send[Update](u_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (u_ Update) Autorelease() Update {
	rv := objc.Send[Update](u_.ID, objc.Sel("autorelease"))
	return rv
}

// NewUpdate creates a new Update instance.
func NewUpdate() Update {
	return getUpdateClass().New()
}




