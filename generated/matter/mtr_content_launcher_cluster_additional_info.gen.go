// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [MTRContentLauncherClusterAdditionalInfo] class.
var (
	MTRContentLauncherClusterAdditionalInfoClass     _MTRContentLauncherClusterAdditionalInfoClass
	MTRContentLauncherClusterAdditionalInfoClassOnce sync.Once
)

func getMTRContentLauncherClusterAdditionalInfoClass() _MTRContentLauncherClusterAdditionalInfoClass {
	MTRContentLauncherClusterAdditionalInfoClassOnce.Do(func() {
		MTRContentLauncherClusterAdditionalInfoClass = _MTRContentLauncherClusterAdditionalInfoClass{objc.GetClass("MTRContentLauncherClusterAdditionalInfo")}
	})
	return MTRContentLauncherClusterAdditionalInfoClass
}

type _MTRContentLauncherClusterAdditionalInfoClass struct {
	class objc.Class
}

// An interface definition for the [MTRContentLauncherClusterAdditionalInfo] class.
type IMTRContentLauncherClusterAdditionalInfo interface {
	IMTRContentLauncherClusterAdditionalInfoStruct
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRContentLauncherClusterAdditionalInfo
type MTRContentLauncherClusterAdditionalInfo struct {
	MTRContentLauncherClusterAdditionalInfoStruct
}

// MTRContentLauncherClusterAdditionalInfoFrom constructs a [MTRContentLauncherClusterAdditionalInfo] from an unsafe.Pointer.
func MTRContentLauncherClusterAdditionalInfoFrom(ptr unsafe.Pointer) MTRContentLauncherClusterAdditionalInfo {
	return MTRContentLauncherClusterAdditionalInfo{
		MTRContentLauncherClusterAdditionalInfoStruct: MTRContentLauncherClusterAdditionalInfoStructFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRContentLauncherClusterAdditionalInfoClass) Alloc() MTRContentLauncherClusterAdditionalInfo {
	rv := objc.Send[MTRContentLauncherClusterAdditionalInfo](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRContentLauncherClusterAdditionalInfoClass) New() MTRContentLauncherClusterAdditionalInfo {
	rv := objc.Send[MTRContentLauncherClusterAdditionalInfo](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRContentLauncherClusterAdditionalInfo) Init() MTRContentLauncherClusterAdditionalInfo {
	rv := objc.Send[MTRContentLauncherClusterAdditionalInfo](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRContentLauncherClusterAdditionalInfo) Autorelease() MTRContentLauncherClusterAdditionalInfo {
	rv := objc.Send[MTRContentLauncherClusterAdditionalInfo](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRContentLauncherClusterAdditionalInfo creates a new MTRContentLauncherClusterAdditionalInfo instance.
func NewMTRContentLauncherClusterAdditionalInfo() MTRContentLauncherClusterAdditionalInfo {
	return getMTRContentLauncherClusterAdditionalInfoClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcontentlauncherclusteradditionalinfo/name
func (m_ MTRContentLauncherClusterAdditionalInfo) Name() appkit.string {
	rv := objc.Send[appkit.string](m_.ID, objc.Sel("name"))
	return rv
}


// SetName sets the value of the name property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcontentlauncherclusteradditionalinfo/name
func (m_ MTRContentLauncherClusterAdditionalInfo) SetName(value appkit.string) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setName:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcontentlauncherclusteradditionalinfo/value
func (m_ MTRContentLauncherClusterAdditionalInfo) Value() appkit.string {
	rv := objc.Send[appkit.string](m_.ID, objc.Sel("value"))
	return rv
}


// SetValue sets the value of the value property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcontentlauncherclusteradditionalinfo/value
func (m_ MTRContentLauncherClusterAdditionalInfo) SetValue(value appkit.string) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setValue:"), value)
}



