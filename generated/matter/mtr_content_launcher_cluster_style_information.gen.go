// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [MTRContentLauncherClusterStyleInformation] class.
var (
	MTRContentLauncherClusterStyleInformationClass     _MTRContentLauncherClusterStyleInformationClass
	MTRContentLauncherClusterStyleInformationClassOnce sync.Once
)

func getMTRContentLauncherClusterStyleInformationClass() _MTRContentLauncherClusterStyleInformationClass {
	MTRContentLauncherClusterStyleInformationClassOnce.Do(func() {
		MTRContentLauncherClusterStyleInformationClass = _MTRContentLauncherClusterStyleInformationClass{objc.GetClass("MTRContentLauncherClusterStyleInformation")}
	})
	return MTRContentLauncherClusterStyleInformationClass
}

type _MTRContentLauncherClusterStyleInformationClass struct {
	class objc.Class
}

// An interface definition for the [MTRContentLauncherClusterStyleInformation] class.
type IMTRContentLauncherClusterStyleInformation interface {
	IMTRContentLauncherClusterStyleInformationStruct
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRContentLauncherClusterStyleInformation
type MTRContentLauncherClusterStyleInformation struct {
	MTRContentLauncherClusterStyleInformationStruct
}

// MTRContentLauncherClusterStyleInformationFrom constructs a [MTRContentLauncherClusterStyleInformation] from an unsafe.Pointer.
func MTRContentLauncherClusterStyleInformationFrom(ptr unsafe.Pointer) MTRContentLauncherClusterStyleInformation {
	return MTRContentLauncherClusterStyleInformation{
		MTRContentLauncherClusterStyleInformationStruct: MTRContentLauncherClusterStyleInformationStructFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRContentLauncherClusterStyleInformationClass) Alloc() MTRContentLauncherClusterStyleInformation {
	rv := objc.Send[MTRContentLauncherClusterStyleInformation](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRContentLauncherClusterStyleInformationClass) New() MTRContentLauncherClusterStyleInformation {
	rv := objc.Send[MTRContentLauncherClusterStyleInformation](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRContentLauncherClusterStyleInformation) Init() MTRContentLauncherClusterStyleInformation {
	rv := objc.Send[MTRContentLauncherClusterStyleInformation](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRContentLauncherClusterStyleInformation) Autorelease() MTRContentLauncherClusterStyleInformation {
	rv := objc.Send[MTRContentLauncherClusterStyleInformation](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRContentLauncherClusterStyleInformation creates a new MTRContentLauncherClusterStyleInformation instance.
func NewMTRContentLauncherClusterStyleInformation() MTRContentLauncherClusterStyleInformation {
	return getMTRContentLauncherClusterStyleInformationClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcontentlauncherclusterstyleinformation/size
func (m_ MTRContentLauncherClusterStyleInformation) Size() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("size"))
	return rv
}


// SetSize sets the value of the size property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcontentlauncherclusterstyleinformation/size
func (m_ MTRContentLauncherClusterStyleInformation) SetSize(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setSize:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcontentlauncherclusterstyleinformation/color
func (m_ MTRContentLauncherClusterStyleInformation) Color() string {
	rv := objc.Send[string](m_.ID, objc.Sel("color"))
	return rv
}


// SetColor sets the value of the color property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcontentlauncherclusterstyleinformation/color
func (m_ MTRContentLauncherClusterStyleInformation) SetColor(value string) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setColor:"), objc.String(value))
}



