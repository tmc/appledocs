// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTROperationalCredentialsClusterFabricDescriptorStruct] class.
var (
	MTROperationalCredentialsClusterFabricDescriptorStructClass     _MTROperationalCredentialsClusterFabricDescriptorStructClass
	MTROperationalCredentialsClusterFabricDescriptorStructClassOnce sync.Once
)

func getMTROperationalCredentialsClusterFabricDescriptorStructClass() _MTROperationalCredentialsClusterFabricDescriptorStructClass {
	MTROperationalCredentialsClusterFabricDescriptorStructClassOnce.Do(func() {
		MTROperationalCredentialsClusterFabricDescriptorStructClass = _MTROperationalCredentialsClusterFabricDescriptorStructClass{objc.GetClass("MTROperationalCredentialsClusterFabricDescriptorStruct")}
	})
	return MTROperationalCredentialsClusterFabricDescriptorStructClass
}

type _MTROperationalCredentialsClusterFabricDescriptorStructClass struct {
	class objc.Class
}

// An interface definition for the [MTROperationalCredentialsClusterFabricDescriptorStruct] class.
type IMTROperationalCredentialsClusterFabricDescriptorStruct interface {
	objectivec.IObject
	FabricID() foundation.Number
	SetFabricID(value foundation.INumber)
	FabricId() foundation.Number
	SetFabricId(value foundation.INumber)
	FabricIndex() foundation.Number
	SetFabricIndex(value foundation.INumber)
	Label() string
	SetLabel(value string)
	NodeID() foundation.Number
	SetNodeID(value foundation.INumber)
	NodeId() foundation.Number
	SetNodeId(value foundation.INumber)
	RootPublicKey() foundation.Data
	SetRootPublicKey(value foundation.IData)
	VendorID() foundation.Number
	SetVendorID(value foundation.INumber)
	VendorId() foundation.Number
	SetVendorId(value foundation.INumber)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROperationalCredentialsClusterFabricDescriptorStruct
type MTROperationalCredentialsClusterFabricDescriptorStruct struct {
	objectivec.Object
}

// MTROperationalCredentialsClusterFabricDescriptorStructFrom constructs a [MTROperationalCredentialsClusterFabricDescriptorStruct] from an unsafe.Pointer.
func MTROperationalCredentialsClusterFabricDescriptorStructFrom(ptr unsafe.Pointer) MTROperationalCredentialsClusterFabricDescriptorStruct {
	return MTROperationalCredentialsClusterFabricDescriptorStruct{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTROperationalCredentialsClusterFabricDescriptorStructClass) Alloc() MTROperationalCredentialsClusterFabricDescriptorStruct {
	rv := objc.Send[MTROperationalCredentialsClusterFabricDescriptorStruct](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTROperationalCredentialsClusterFabricDescriptorStructClass) New() MTROperationalCredentialsClusterFabricDescriptorStruct {
	rv := objc.Send[MTROperationalCredentialsClusterFabricDescriptorStruct](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTROperationalCredentialsClusterFabricDescriptorStruct) Init() MTROperationalCredentialsClusterFabricDescriptorStruct {
	rv := objc.Send[MTROperationalCredentialsClusterFabricDescriptorStruct](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTROperationalCredentialsClusterFabricDescriptorStruct) Autorelease() MTROperationalCredentialsClusterFabricDescriptorStruct {
	rv := objc.Send[MTROperationalCredentialsClusterFabricDescriptorStruct](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTROperationalCredentialsClusterFabricDescriptorStruct creates a new MTROperationalCredentialsClusterFabricDescriptorStruct instance.
func NewMTROperationalCredentialsClusterFabricDescriptorStruct() MTROperationalCredentialsClusterFabricDescriptorStruct {
	return getMTROperationalCredentialsClusterFabricDescriptorStructClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtroperationalcredentialsclusterfabricdescriptorstruct/fabricid-5teul
func (m_ MTROperationalCredentialsClusterFabricDescriptorStruct) FabricID() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("fabricID"))
	return rv
}


// SetFabricID sets the value of the fabricID property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtroperationalcredentialsclusterfabricdescriptorstruct/fabricid-5teul
func (m_ MTROperationalCredentialsClusterFabricDescriptorStruct) SetFabricID(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setFabricID:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtroperationalcredentialsclusterfabricdescriptorstruct/fabricid-5tetp
func (m_ MTROperationalCredentialsClusterFabricDescriptorStruct) FabricId() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("fabricId"))
	return rv
}


// SetFabricId sets the value of the fabricId property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtroperationalcredentialsclusterfabricdescriptorstruct/fabricid-5tetp
func (m_ MTROperationalCredentialsClusterFabricDescriptorStruct) SetFabricId(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setFabricId:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtroperationalcredentialsclusterfabricdescriptorstruct/fabricindex
func (m_ MTROperationalCredentialsClusterFabricDescriptorStruct) FabricIndex() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("fabricIndex"))
	return rv
}


// SetFabricIndex sets the value of the fabricIndex property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtroperationalcredentialsclusterfabricdescriptorstruct/fabricindex
func (m_ MTROperationalCredentialsClusterFabricDescriptorStruct) SetFabricIndex(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setFabricIndex:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtroperationalcredentialsclusterfabricdescriptorstruct/label
func (m_ MTROperationalCredentialsClusterFabricDescriptorStruct) Label() string {
	rv := objc.Send[string](m_.ID, objc.Sel("label"))
	return rv
}


// SetLabel sets the value of the label property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtroperationalcredentialsclusterfabricdescriptorstruct/label
func (m_ MTROperationalCredentialsClusterFabricDescriptorStruct) SetLabel(value string) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setLabel:"), objc.String(value))
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtroperationalcredentialsclusterfabricdescriptorstruct/nodeid-3hsb0
func (m_ MTROperationalCredentialsClusterFabricDescriptorStruct) NodeID() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("nodeID"))
	return rv
}


// SetNodeID sets the value of the nodeID property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtroperationalcredentialsclusterfabricdescriptorstruct/nodeid-3hsb0
func (m_ MTROperationalCredentialsClusterFabricDescriptorStruct) SetNodeID(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setNodeID:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtroperationalcredentialsclusterfabricdescriptorstruct/nodeid-3hsbw
func (m_ MTROperationalCredentialsClusterFabricDescriptorStruct) NodeId() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("nodeId"))
	return rv
}


// SetNodeId sets the value of the nodeId property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtroperationalcredentialsclusterfabricdescriptorstruct/nodeid-3hsbw
func (m_ MTROperationalCredentialsClusterFabricDescriptorStruct) SetNodeId(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setNodeId:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtroperationalcredentialsclusterfabricdescriptorstruct/rootpublickey
func (m_ MTROperationalCredentialsClusterFabricDescriptorStruct) RootPublicKey() foundation.Data {
	rv := objc.Send[foundation.Data](m_.ID, objc.Sel("rootPublicKey"))
	return rv
}


// SetRootPublicKey sets the value of the rootPublicKey property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtroperationalcredentialsclusterfabricdescriptorstruct/rootpublickey
func (m_ MTROperationalCredentialsClusterFabricDescriptorStruct) SetRootPublicKey(value foundation.IData) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setRootPublicKey:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtroperationalcredentialsclusterfabricdescriptorstruct/vendorid-3iay9
func (m_ MTROperationalCredentialsClusterFabricDescriptorStruct) VendorID() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("vendorID"))
	return rv
}


// SetVendorID sets the value of the vendorID property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtroperationalcredentialsclusterfabricdescriptorstruct/vendorid-3iay9
func (m_ MTROperationalCredentialsClusterFabricDescriptorStruct) SetVendorID(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setVendorID:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtroperationalcredentialsclusterfabricdescriptorstruct/vendorid-3iaxd
func (m_ MTROperationalCredentialsClusterFabricDescriptorStruct) VendorId() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("vendorId"))
	return rv
}


// SetVendorId sets the value of the vendorId property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtroperationalcredentialsclusterfabricdescriptorstruct/vendorid-3iaxd
func (m_ MTROperationalCredentialsClusterFabricDescriptorStruct) SetVendorId(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setVendorId:"), value)
}



