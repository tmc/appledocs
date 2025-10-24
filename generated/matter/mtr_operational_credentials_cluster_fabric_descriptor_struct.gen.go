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
	// properties:
	FabricID() objc.IObject /* cross-framework: NSNumber */
	SetFabricID(value objc.IObject /* cross-framework: NSNumber */)
	FabricId() objc.IObject /* cross-framework: NSNumber */
	SetFabricId(value objc.IObject /* cross-framework: NSNumber */)
	FabricIndex() objc.IObject /* cross-framework: NSNumber */
	SetFabricIndex(value objc.IObject /* cross-framework: NSNumber */)
	Label() objc.IObject /* cross-framework: NSString */
	SetLabel(value objc.IObject /* cross-framework: NSString */)
	NodeID() objc.IObject /* cross-framework: NSNumber */
	SetNodeID(value objc.IObject /* cross-framework: NSNumber */)
	NodeId() objc.IObject /* cross-framework: NSNumber */
	SetNodeId(value objc.IObject /* cross-framework: NSNumber */)
	RootPublicKey() objc.IObject /* cross-framework: Data */
	SetRootPublicKey(value objc.IObject /* cross-framework: Data */)
	VendorID() objc.IObject /* cross-framework: NSNumber */
	SetVendorID(value objc.IObject /* cross-framework: NSNumber */)
	VendorId() objc.IObject /* cross-framework: NSNumber */
	SetVendorId(value objc.IObject /* cross-framework: NSNumber */)
	// methods:
}



// [Full Topic]
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



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtroperationalcredentialsclusterfabricdescriptorstruct/fabricid-5teul
func (m_ MTROperationalCredentialsClusterFabricDescriptorStruct) FabricID() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("fabricID"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtroperationalcredentialsclusterfabricdescriptorstruct/fabricid-5teul
func (m_ MTROperationalCredentialsClusterFabricDescriptorStruct) SetFabricID(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setFabricID:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtroperationalcredentialsclusterfabricdescriptorstruct/fabricid-5tetp
func (m_ MTROperationalCredentialsClusterFabricDescriptorStruct) FabricId() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("fabricId"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtroperationalcredentialsclusterfabricdescriptorstruct/fabricid-5tetp
func (m_ MTROperationalCredentialsClusterFabricDescriptorStruct) SetFabricId(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setFabricId:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtroperationalcredentialsclusterfabricdescriptorstruct/fabricindex
func (m_ MTROperationalCredentialsClusterFabricDescriptorStruct) FabricIndex() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("fabricIndex"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtroperationalcredentialsclusterfabricdescriptorstruct/fabricindex
func (m_ MTROperationalCredentialsClusterFabricDescriptorStruct) SetFabricIndex(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setFabricIndex:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtroperationalcredentialsclusterfabricdescriptorstruct/label
func (m_ MTROperationalCredentialsClusterFabricDescriptorStruct) Label() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("label"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtroperationalcredentialsclusterfabricdescriptorstruct/label
func (m_ MTROperationalCredentialsClusterFabricDescriptorStruct) SetLabel(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setLabel:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtroperationalcredentialsclusterfabricdescriptorstruct/nodeid-3hsb0
func (m_ MTROperationalCredentialsClusterFabricDescriptorStruct) NodeID() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("nodeID"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtroperationalcredentialsclusterfabricdescriptorstruct/nodeid-3hsb0
func (m_ MTROperationalCredentialsClusterFabricDescriptorStruct) SetNodeID(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setNodeID:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtroperationalcredentialsclusterfabricdescriptorstruct/nodeid-3hsbw
func (m_ MTROperationalCredentialsClusterFabricDescriptorStruct) NodeId() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("nodeId"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtroperationalcredentialsclusterfabricdescriptorstruct/nodeid-3hsbw
func (m_ MTROperationalCredentialsClusterFabricDescriptorStruct) SetNodeId(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setNodeId:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtroperationalcredentialsclusterfabricdescriptorstruct/rootpublickey
func (m_ MTROperationalCredentialsClusterFabricDescriptorStruct) RootPublicKey() objc.IObject /* cross-framework: Data */ {
	rv := objc.Send[foundation.Data](m_.ID, objc.Sel("rootPublicKey"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtroperationalcredentialsclusterfabricdescriptorstruct/rootpublickey
func (m_ MTROperationalCredentialsClusterFabricDescriptorStruct) SetRootPublicKey(value objc.IObject /* cross-framework: Data */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setRootPublicKey:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtroperationalcredentialsclusterfabricdescriptorstruct/vendorid-3iay9
func (m_ MTROperationalCredentialsClusterFabricDescriptorStruct) VendorID() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("vendorID"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtroperationalcredentialsclusterfabricdescriptorstruct/vendorid-3iay9
func (m_ MTROperationalCredentialsClusterFabricDescriptorStruct) SetVendorID(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setVendorID:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtroperationalcredentialsclusterfabricdescriptorstruct/vendorid-3iaxd
func (m_ MTROperationalCredentialsClusterFabricDescriptorStruct) VendorId() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("vendorId"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtroperationalcredentialsclusterfabricdescriptorstruct/vendorid-3iaxd
func (m_ MTROperationalCredentialsClusterFabricDescriptorStruct) SetVendorId(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setVendorId:"), value)
}



