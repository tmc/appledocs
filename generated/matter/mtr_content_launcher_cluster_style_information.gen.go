// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/foundation"
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
	// properties:
	Color() objc.IObject /* cross-framework: NSString */
	SetColor(value objc.IObject /* cross-framework: NSString */)
	Size() IMTRContentLauncherClusterDimensionStruct
	SetSize(value IMTRContentLauncherClusterDimensionStruct)
	// methods:
}

// [Full Topic]
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

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcontentlauncherclusterstyleinformation/color
func (m_ MTRContentLauncherClusterStyleInformation) Color() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("color"))
	return rv
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcontentlauncherclusterstyleinformation/color
func (m_ MTRContentLauncherClusterStyleInformation) SetColor(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setColor:"), value)
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcontentlauncherclusterstyleinformation/size
func (m_ MTRContentLauncherClusterStyleInformation) Size() IMTRContentLauncherClusterDimensionStruct {
	rv := objc.Send[MTRContentLauncherClusterDimensionStruct](m_.ID, objc.Sel("size"))
	return rv
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcontentlauncherclusterstyleinformation/size
func (m_ MTRContentLauncherClusterStyleInformation) SetSize(value IMTRContentLauncherClusterDimensionStruct) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setSize:"), value)
}
