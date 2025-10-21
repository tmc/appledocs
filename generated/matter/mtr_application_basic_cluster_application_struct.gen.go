// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRApplicationBasicClusterApplicationStruct] class.
var (
	MTRApplicationBasicClusterApplicationStructClass     _MTRApplicationBasicClusterApplicationStructClass
	MTRApplicationBasicClusterApplicationStructClassOnce sync.Once
)

func getMTRApplicationBasicClusterApplicationStructClass() _MTRApplicationBasicClusterApplicationStructClass {
	MTRApplicationBasicClusterApplicationStructClassOnce.Do(func() {
		MTRApplicationBasicClusterApplicationStructClass = _MTRApplicationBasicClusterApplicationStructClass{objc.GetClass("MTRApplicationBasicClusterApplicationStruct")}
	})
	return MTRApplicationBasicClusterApplicationStructClass
}

type _MTRApplicationBasicClusterApplicationStructClass struct {
	class objc.Class
}

// An interface definition for the [MTRApplicationBasicClusterApplicationStruct] class.
type IMTRApplicationBasicClusterApplicationStruct interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRApplicationBasicClusterApplicationStruct
type MTRApplicationBasicClusterApplicationStruct struct {
	objectivec.Object
}

// MTRApplicationBasicClusterApplicationStructFrom constructs a [MTRApplicationBasicClusterApplicationStruct] from an unsafe.Pointer.
func MTRApplicationBasicClusterApplicationStructFrom(ptr unsafe.Pointer) MTRApplicationBasicClusterApplicationStruct {
	return MTRApplicationBasicClusterApplicationStruct{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRApplicationBasicClusterApplicationStructClass) Alloc() MTRApplicationBasicClusterApplicationStruct {
	rv := objc.Send[MTRApplicationBasicClusterApplicationStruct](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRApplicationBasicClusterApplicationStructClass) New() MTRApplicationBasicClusterApplicationStruct {
	rv := objc.Send[MTRApplicationBasicClusterApplicationStruct](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRApplicationBasicClusterApplicationStruct) Init() MTRApplicationBasicClusterApplicationStruct {
	rv := objc.Send[MTRApplicationBasicClusterApplicationStruct](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRApplicationBasicClusterApplicationStruct) Autorelease() MTRApplicationBasicClusterApplicationStruct {
	rv := objc.Send[MTRApplicationBasicClusterApplicationStruct](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRApplicationBasicClusterApplicationStruct creates a new MTRApplicationBasicClusterApplicationStruct instance.
func NewMTRApplicationBasicClusterApplicationStruct() MTRApplicationBasicClusterApplicationStruct {
	return getMTRApplicationBasicClusterApplicationStructClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrapplicationbasicclusterapplicationstruct/applicationid-1jzu
func (m_ MTRApplicationBasicClusterApplicationStruct) ApplicationID() string {
	rv := objc.Send[string](m_.ID, objc.Sel("applicationID"))
	return rv
}


// SetApplicationID sets the value of the applicationID property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrapplicationbasicclusterapplicationstruct/applicationid-1jzu
func (m_ MTRApplicationBasicClusterApplicationStruct) SetApplicationID(value string) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setApplicationID:"), objc.String(value))
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrapplicationbasicclusterapplicationstruct/applicationid-1jyy
func (m_ MTRApplicationBasicClusterApplicationStruct) ApplicationId() string {
	rv := objc.Send[string](m_.ID, objc.Sel("applicationId"))
	return rv
}


// SetApplicationId sets the value of the applicationId property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrapplicationbasicclusterapplicationstruct/applicationid-1jyy
func (m_ MTRApplicationBasicClusterApplicationStruct) SetApplicationId(value string) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setApplicationId:"), objc.String(value))
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrapplicationbasicclusterapplicationstruct/catalogvendorid-16o17
func (m_ MTRApplicationBasicClusterApplicationStruct) CatalogVendorID() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("catalogVendorID"))
	return rv
}


// SetCatalogVendorID sets the value of the catalogVendorID property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrapplicationbasicclusterapplicationstruct/catalogvendorid-16o17
func (m_ MTRApplicationBasicClusterApplicationStruct) SetCatalogVendorID(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setCatalogVendorID:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrapplicationbasicclusterapplicationstruct/catalogvendorid-16o0b
func (m_ MTRApplicationBasicClusterApplicationStruct) CatalogVendorId() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("catalogVendorId"))
	return rv
}


// SetCatalogVendorId sets the value of the catalogVendorId property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrapplicationbasicclusterapplicationstruct/catalogvendorid-16o0b
func (m_ MTRApplicationBasicClusterApplicationStruct) SetCatalogVendorId(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setCatalogVendorId:"), value)
}



