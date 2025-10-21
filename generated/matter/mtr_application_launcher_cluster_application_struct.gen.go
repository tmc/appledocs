// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [MTRApplicationLauncherClusterApplicationStruct] class.
var (
	MTRApplicationLauncherClusterApplicationStructClass     _MTRApplicationLauncherClusterApplicationStructClass
	MTRApplicationLauncherClusterApplicationStructClassOnce sync.Once
)

func getMTRApplicationLauncherClusterApplicationStructClass() _MTRApplicationLauncherClusterApplicationStructClass {
	MTRApplicationLauncherClusterApplicationStructClassOnce.Do(func() {
		MTRApplicationLauncherClusterApplicationStructClass = _MTRApplicationLauncherClusterApplicationStructClass{objc.GetClass("MTRApplicationLauncherClusterApplicationStruct")}
	})
	return MTRApplicationLauncherClusterApplicationStructClass
}

type _MTRApplicationLauncherClusterApplicationStructClass struct {
	class objc.Class
}

// An interface definition for the [MTRApplicationLauncherClusterApplicationStruct] class.
type IMTRApplicationLauncherClusterApplicationStruct interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRApplicationLauncherClusterApplicationStruct
type MTRApplicationLauncherClusterApplicationStruct struct {
	objectivec.Object
}

// MTRApplicationLauncherClusterApplicationStructFrom constructs a [MTRApplicationLauncherClusterApplicationStruct] from an unsafe.Pointer.
func MTRApplicationLauncherClusterApplicationStructFrom(ptr unsafe.Pointer) MTRApplicationLauncherClusterApplicationStruct {
	return MTRApplicationLauncherClusterApplicationStruct{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRApplicationLauncherClusterApplicationStructClass) Alloc() MTRApplicationLauncherClusterApplicationStruct {
	rv := objc.Send[MTRApplicationLauncherClusterApplicationStruct](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRApplicationLauncherClusterApplicationStructClass) New() MTRApplicationLauncherClusterApplicationStruct {
	rv := objc.Send[MTRApplicationLauncherClusterApplicationStruct](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRApplicationLauncherClusterApplicationStruct) Init() MTRApplicationLauncherClusterApplicationStruct {
	rv := objc.Send[MTRApplicationLauncherClusterApplicationStruct](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRApplicationLauncherClusterApplicationStruct) Autorelease() MTRApplicationLauncherClusterApplicationStruct {
	rv := objc.Send[MTRApplicationLauncherClusterApplicationStruct](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRApplicationLauncherClusterApplicationStruct creates a new MTRApplicationLauncherClusterApplicationStruct instance.
func NewMTRApplicationLauncherClusterApplicationStruct() MTRApplicationLauncherClusterApplicationStruct {
	return getMTRApplicationLauncherClusterApplicationStructClass().New()
}




