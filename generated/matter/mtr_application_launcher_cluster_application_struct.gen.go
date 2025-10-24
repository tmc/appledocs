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
	// properties:
	ApplicationID() objc.IObject /* cross-framework: NSString */
	SetApplicationID(value objc.IObject /* cross-framework: NSString */)
	ApplicationId() objc.IObject /* cross-framework: NSString */
	SetApplicationId(value objc.IObject /* cross-framework: NSString */)
	CatalogVendorID() objc.IObject /* cross-framework: NSNumber */
	SetCatalogVendorID(value objc.IObject /* cross-framework: NSNumber */)
	CatalogVendorId() objc.IObject /* cross-framework: NSNumber */
	SetCatalogVendorId(value objc.IObject /* cross-framework: NSNumber */)
	// methods:
}



// [Full Topic]
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



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrapplicationlauncherclusterapplicationstruct/applicationid-6t05a
func (m_ MTRApplicationLauncherClusterApplicationStruct) ApplicationID() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("applicationID"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrapplicationlauncherclusterapplicationstruct/applicationid-6t05a
func (m_ MTRApplicationLauncherClusterApplicationStruct) SetApplicationID(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setApplicationID:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrapplicationlauncherclusterapplicationstruct/applicationid-6t04e
func (m_ MTRApplicationLauncherClusterApplicationStruct) ApplicationId() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("applicationId"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrapplicationlauncherclusterapplicationstruct/applicationid-6t04e
func (m_ MTRApplicationLauncherClusterApplicationStruct) SetApplicationId(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setApplicationId:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrapplicationlauncherclusterapplicationstruct/catalogvendorid-rb5w
func (m_ MTRApplicationLauncherClusterApplicationStruct) CatalogVendorID() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("catalogVendorID"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrapplicationlauncherclusterapplicationstruct/catalogvendorid-rb5w
func (m_ MTRApplicationLauncherClusterApplicationStruct) SetCatalogVendorID(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setCatalogVendorID:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrapplicationlauncherclusterapplicationstruct/catalogvendorid-rb6s
func (m_ MTRApplicationLauncherClusterApplicationStruct) CatalogVendorId() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("catalogVendorId"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrapplicationlauncherclusterapplicationstruct/catalogvendorid-rb6s
func (m_ MTRApplicationLauncherClusterApplicationStruct) SetCatalogVendorId(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setCatalogVendorId:"), value)
}



