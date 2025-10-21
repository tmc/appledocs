// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
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


//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrapplicationlauncherclusterapplicationstruct/applicationid-6t05a
func (m_ MTRApplicationLauncherClusterApplicationStruct) ApplicationID() string {
	rv := objc.Send[string](m_.ID, objc.Sel("applicationID"))
	return rv
}


// SetApplicationID sets the value of the applicationID property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrapplicationlauncherclusterapplicationstruct/applicationid-6t05a
func (m_ MTRApplicationLauncherClusterApplicationStruct) SetApplicationID(value string) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setApplicationID:"), objc.String(value))
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrapplicationlauncherclusterapplicationstruct/applicationid-6t04e
func (m_ MTRApplicationLauncherClusterApplicationStruct) ApplicationId() string {
	rv := objc.Send[string](m_.ID, objc.Sel("applicationId"))
	return rv
}


// SetApplicationId sets the value of the applicationId property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrapplicationlauncherclusterapplicationstruct/applicationid-6t04e
func (m_ MTRApplicationLauncherClusterApplicationStruct) SetApplicationId(value string) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setApplicationId:"), objc.String(value))
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrapplicationlauncherclusterapplicationstruct/catalogvendorid-rb5w
func (m_ MTRApplicationLauncherClusterApplicationStruct) CatalogVendorID() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("catalogVendorID"))
	return rv
}


// SetCatalogVendorID sets the value of the catalogVendorID property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrapplicationlauncherclusterapplicationstruct/catalogvendorid-rb5w
func (m_ MTRApplicationLauncherClusterApplicationStruct) SetCatalogVendorID(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setCatalogVendorID:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrapplicationlauncherclusterapplicationstruct/catalogvendorid-rb6s
func (m_ MTRApplicationLauncherClusterApplicationStruct) CatalogVendorId() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("catalogVendorId"))
	return rv
}


// SetCatalogVendorId sets the value of the catalogVendorId property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrapplicationlauncherclusterapplicationstruct/catalogvendorid-rb6s
func (m_ MTRApplicationLauncherClusterApplicationStruct) SetCatalogVendorId(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setCatalogVendorId:"), value)
}



