// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [MTRContentLauncherClusterBrandingInformation] class.
var (
	MTRContentLauncherClusterBrandingInformationClass     _MTRContentLauncherClusterBrandingInformationClass
	MTRContentLauncherClusterBrandingInformationClassOnce sync.Once
)

func getMTRContentLauncherClusterBrandingInformationClass() _MTRContentLauncherClusterBrandingInformationClass {
	MTRContentLauncherClusterBrandingInformationClassOnce.Do(func() {
		MTRContentLauncherClusterBrandingInformationClass = _MTRContentLauncherClusterBrandingInformationClass{objc.GetClass("MTRContentLauncherClusterBrandingInformation")}
	})
	return MTRContentLauncherClusterBrandingInformationClass
}

type _MTRContentLauncherClusterBrandingInformationClass struct {
	class objc.Class
}

// An interface definition for the [MTRContentLauncherClusterBrandingInformation] class.
type IMTRContentLauncherClusterBrandingInformation interface {
	IMTRContentLauncherClusterBrandingInformationStruct
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRContentLauncherClusterBrandingInformation
type MTRContentLauncherClusterBrandingInformation struct {
	MTRContentLauncherClusterBrandingInformationStruct
}

// MTRContentLauncherClusterBrandingInformationFrom constructs a [MTRContentLauncherClusterBrandingInformation] from an unsafe.Pointer.
func MTRContentLauncherClusterBrandingInformationFrom(ptr unsafe.Pointer) MTRContentLauncherClusterBrandingInformation {
	return MTRContentLauncherClusterBrandingInformation{
		MTRContentLauncherClusterBrandingInformationStruct: MTRContentLauncherClusterBrandingInformationStructFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRContentLauncherClusterBrandingInformationClass) Alloc() MTRContentLauncherClusterBrandingInformation {
	rv := objc.Send[MTRContentLauncherClusterBrandingInformation](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRContentLauncherClusterBrandingInformationClass) New() MTRContentLauncherClusterBrandingInformation {
	rv := objc.Send[MTRContentLauncherClusterBrandingInformation](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRContentLauncherClusterBrandingInformation) Init() MTRContentLauncherClusterBrandingInformation {
	rv := objc.Send[MTRContentLauncherClusterBrandingInformation](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRContentLauncherClusterBrandingInformation) Autorelease() MTRContentLauncherClusterBrandingInformation {
	rv := objc.Send[MTRContentLauncherClusterBrandingInformation](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRContentLauncherClusterBrandingInformation creates a new MTRContentLauncherClusterBrandingInformation instance.
func NewMTRContentLauncherClusterBrandingInformation() MTRContentLauncherClusterBrandingInformation {
	return getMTRContentLauncherClusterBrandingInformationClass().New()
}




