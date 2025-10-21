// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [MTRClusterBasicInformation] class.
var (
	MTRClusterBasicInformationClass     _MTRClusterBasicInformationClass
	MTRClusterBasicInformationClassOnce sync.Once
)

func getMTRClusterBasicInformationClass() _MTRClusterBasicInformationClass {
	MTRClusterBasicInformationClassOnce.Do(func() {
		MTRClusterBasicInformationClass = _MTRClusterBasicInformationClass{objc.GetClass("MTRClusterBasicInformation")}
	})
	return MTRClusterBasicInformationClass
}

type _MTRClusterBasicInformationClass struct {
	class objc.Class
}

// An interface definition for the [MTRClusterBasicInformation] class.
type IMTRClusterBasicInformation interface {
	IMTRGenericCluster
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterBasicInformation
type MTRClusterBasicInformation struct {
	MTRGenericCluster
}

// MTRClusterBasicInformationFrom constructs a [MTRClusterBasicInformation] from an unsafe.Pointer.
func MTRClusterBasicInformationFrom(ptr unsafe.Pointer) MTRClusterBasicInformation {
	return MTRClusterBasicInformation{
		MTRGenericCluster: MTRGenericClusterFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRClusterBasicInformationClass) Alloc() MTRClusterBasicInformation {
	rv := objc.Send[MTRClusterBasicInformation](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRClusterBasicInformationClass) New() MTRClusterBasicInformation {
	rv := objc.Send[MTRClusterBasicInformation](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRClusterBasicInformation) Init() MTRClusterBasicInformation {
	rv := objc.Send[MTRClusterBasicInformation](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRClusterBasicInformation) Autorelease() MTRClusterBasicInformation {
	rv := objc.Send[MTRClusterBasicInformation](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRClusterBasicInformation creates a new MTRClusterBasicInformation instance.
func NewMTRClusterBasicInformation() MTRClusterBasicInformation {
	return getMTRClusterBasicInformationClass().New()
}




