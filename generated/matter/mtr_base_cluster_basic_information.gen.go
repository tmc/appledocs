// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [MTRBaseClusterBasicInformation] class.
var (
	MTRBaseClusterBasicInformationClass     _MTRBaseClusterBasicInformationClass
	MTRBaseClusterBasicInformationClassOnce sync.Once
)

func getMTRBaseClusterBasicInformationClass() _MTRBaseClusterBasicInformationClass {
	MTRBaseClusterBasicInformationClassOnce.Do(func() {
		MTRBaseClusterBasicInformationClass = _MTRBaseClusterBasicInformationClass{objc.GetClass("MTRBaseClusterBasicInformation")}
	})
	return MTRBaseClusterBasicInformationClass
}

type _MTRBaseClusterBasicInformationClass struct {
	class objc.Class
}

// An interface definition for the [MTRBaseClusterBasicInformation] class.
type IMTRBaseClusterBasicInformation interface {
	IMTRGenericBaseCluster
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterBasicInformation
type MTRBaseClusterBasicInformation struct {
	MTRGenericBaseCluster
}

// MTRBaseClusterBasicInformationFrom constructs a [MTRBaseClusterBasicInformation] from an unsafe.Pointer.
func MTRBaseClusterBasicInformationFrom(ptr unsafe.Pointer) MTRBaseClusterBasicInformation {
	return MTRBaseClusterBasicInformation{
		MTRGenericBaseCluster: MTRGenericBaseClusterFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRBaseClusterBasicInformationClass) Alloc() MTRBaseClusterBasicInformation {
	rv := objc.Send[MTRBaseClusterBasicInformation](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRBaseClusterBasicInformationClass) New() MTRBaseClusterBasicInformation {
	rv := objc.Send[MTRBaseClusterBasicInformation](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRBaseClusterBasicInformation) Init() MTRBaseClusterBasicInformation {
	rv := objc.Send[MTRBaseClusterBasicInformation](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRBaseClusterBasicInformation) Autorelease() MTRBaseClusterBasicInformation {
	rv := objc.Send[MTRBaseClusterBasicInformation](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRBaseClusterBasicInformation creates a new MTRBaseClusterBasicInformation instance.
func NewMTRBaseClusterBasicInformation() MTRBaseClusterBasicInformation {
	return getMTRBaseClusterBasicInformationClass().New()
}




