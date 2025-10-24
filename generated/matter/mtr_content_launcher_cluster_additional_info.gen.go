// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
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
	// properties:
	Name() objc.IObject /* cross-framework: NSString */
	SetName(value objc.IObject /* cross-framework: NSString */)
	Value() objc.IObject /* cross-framework: NSString */
	SetValue(value objc.IObject /* cross-framework: NSString */)
	// methods:
}



// [Full Topic]
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



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcontentlauncherclusteradditionalinfo/name
func (m_ MTRContentLauncherClusterAdditionalInfo) Name() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("name"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcontentlauncherclusteradditionalinfo/name
func (m_ MTRContentLauncherClusterAdditionalInfo) SetName(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setName:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcontentlauncherclusteradditionalinfo/value
func (m_ MTRContentLauncherClusterAdditionalInfo) Value() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("value"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcontentlauncherclusteradditionalinfo/value
func (m_ MTRContentLauncherClusterAdditionalInfo) SetValue(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setValue:"), value)
}



