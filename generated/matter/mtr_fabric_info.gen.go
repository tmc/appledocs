// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRFabricInfo] class.
var (
	MTRFabricInfoClass     _MTRFabricInfoClass
	MTRFabricInfoClassOnce sync.Once
)

func getMTRFabricInfoClass() _MTRFabricInfoClass {
	MTRFabricInfoClassOnce.Do(func() {
		MTRFabricInfoClass = _MTRFabricInfoClass{objc.GetClass("MTRFabricInfo")}
	})
	return MTRFabricInfoClass
}

type _MTRFabricInfoClass struct {
	class objc.Class
}

// An interface definition for the [MTRFabricInfo] class.
type IMTRFabricInfo interface {
	objectivec.IObject
	// properties:
	FabricID() objc.IObject /* cross-framework: NSNumber */
	SetFabricID(value objc.IObject /* cross-framework: NSNumber */)
	FabricIndex() objc.IObject /* cross-framework: NSNumber */
	SetFabricIndex(value objc.IObject /* cross-framework: NSNumber */)
	IntermediateCertificate() objc.IObject /* cross-framework: Data */
	SetIntermediateCertificate(value objc.IObject /* cross-framework: Data */)
	IntermediateCertificateTLV() objc.IObject /* cross-framework: Data */
	SetIntermediateCertificateTLV(value objc.IObject /* cross-framework: Data */)
	Label() objc.IObject /* cross-framework: NSString */
	SetLabel(value objc.IObject /* cross-framework: NSString */)
	NodeID() objc.IObject /* cross-framework: NSNumber */
	SetNodeID(value objc.IObject /* cross-framework: NSNumber */)
	OperationalCertificate() objc.IObject /* cross-framework: Data */
	SetOperationalCertificate(value objc.IObject /* cross-framework: Data */)
	OperationalCertificateTLV() objc.IObject /* cross-framework: Data */
	SetOperationalCertificateTLV(value objc.IObject /* cross-framework: Data */)
	RootCertificate() objc.IObject /* cross-framework: Data */
	SetRootCertificate(value objc.IObject /* cross-framework: Data */)
	RootCertificateTLV() objc.IObject /* cross-framework: Data */
	SetRootCertificateTLV(value objc.IObject /* cross-framework: Data */)
	RootPublicKey() objc.IObject /* cross-framework: Data */
	SetRootPublicKey(value objc.IObject /* cross-framework: Data */)
	VendorID() objc.IObject /* cross-framework: NSNumber */
	SetVendorID(value objc.IObject /* cross-framework: NSNumber */)
	// methods:
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRFabricInfo
type MTRFabricInfo struct {
	objectivec.Object
}

// MTRFabricInfoFrom constructs a [MTRFabricInfo] from an unsafe.Pointer.
func MTRFabricInfoFrom(ptr unsafe.Pointer) MTRFabricInfo {
	return MTRFabricInfo{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRFabricInfoClass) Alloc() MTRFabricInfo {
	rv := objc.Send[MTRFabricInfo](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRFabricInfoClass) New() MTRFabricInfo {
	rv := objc.Send[MTRFabricInfo](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRFabricInfo) Init() MTRFabricInfo {
	rv := objc.Send[MTRFabricInfo](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRFabricInfo) Autorelease() MTRFabricInfo {
	rv := objc.Send[MTRFabricInfo](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRFabricInfo creates a new MTRFabricInfo instance.
func NewMTRFabricInfo() MTRFabricInfo {
	return getMTRFabricInfoClass().New()
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrfabricinfo/fabricid
func (m_ MTRFabricInfo) FabricID() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("fabricID"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrfabricinfo/fabricid
func (m_ MTRFabricInfo) SetFabricID(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setFabricID:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrfabricinfo/fabricindex
func (m_ MTRFabricInfo) FabricIndex() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("fabricIndex"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrfabricinfo/fabricindex
func (m_ MTRFabricInfo) SetFabricIndex(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setFabricIndex:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrfabricinfo/intermediatecertificate
func (m_ MTRFabricInfo) IntermediateCertificate() objc.IObject /* cross-framework: Data */ {
	rv := objc.Send[foundation.Data](m_.ID, objc.Sel("intermediateCertificate"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrfabricinfo/intermediatecertificate
func (m_ MTRFabricInfo) SetIntermediateCertificate(value objc.IObject /* cross-framework: Data */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setIntermediateCertificate:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrfabricinfo/intermediatecertificatetlv
func (m_ MTRFabricInfo) IntermediateCertificateTLV() objc.IObject /* cross-framework: Data */ {
	rv := objc.Send[foundation.Data](m_.ID, objc.Sel("intermediateCertificateTLV"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrfabricinfo/intermediatecertificatetlv
func (m_ MTRFabricInfo) SetIntermediateCertificateTLV(value objc.IObject /* cross-framework: Data */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setIntermediateCertificateTLV:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrfabricinfo/label
func (m_ MTRFabricInfo) Label() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("label"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrfabricinfo/label
func (m_ MTRFabricInfo) SetLabel(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setLabel:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrfabricinfo/nodeid
func (m_ MTRFabricInfo) NodeID() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("nodeID"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrfabricinfo/nodeid
func (m_ MTRFabricInfo) SetNodeID(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setNodeID:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrfabricinfo/operationalcertificate
func (m_ MTRFabricInfo) OperationalCertificate() objc.IObject /* cross-framework: Data */ {
	rv := objc.Send[foundation.Data](m_.ID, objc.Sel("operationalCertificate"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrfabricinfo/operationalcertificate
func (m_ MTRFabricInfo) SetOperationalCertificate(value objc.IObject /* cross-framework: Data */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setOperationalCertificate:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrfabricinfo/operationalcertificatetlv
func (m_ MTRFabricInfo) OperationalCertificateTLV() objc.IObject /* cross-framework: Data */ {
	rv := objc.Send[foundation.Data](m_.ID, objc.Sel("operationalCertificateTLV"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrfabricinfo/operationalcertificatetlv
func (m_ MTRFabricInfo) SetOperationalCertificateTLV(value objc.IObject /* cross-framework: Data */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setOperationalCertificateTLV:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrfabricinfo/rootcertificate
func (m_ MTRFabricInfo) RootCertificate() objc.IObject /* cross-framework: Data */ {
	rv := objc.Send[foundation.Data](m_.ID, objc.Sel("rootCertificate"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrfabricinfo/rootcertificate
func (m_ MTRFabricInfo) SetRootCertificate(value objc.IObject /* cross-framework: Data */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setRootCertificate:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrfabricinfo/rootcertificatetlv
func (m_ MTRFabricInfo) RootCertificateTLV() objc.IObject /* cross-framework: Data */ {
	rv := objc.Send[foundation.Data](m_.ID, objc.Sel("rootCertificateTLV"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrfabricinfo/rootcertificatetlv
func (m_ MTRFabricInfo) SetRootCertificateTLV(value objc.IObject /* cross-framework: Data */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setRootCertificateTLV:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrfabricinfo/rootpublickey
func (m_ MTRFabricInfo) RootPublicKey() objc.IObject /* cross-framework: Data */ {
	rv := objc.Send[foundation.Data](m_.ID, objc.Sel("rootPublicKey"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrfabricinfo/rootpublickey
func (m_ MTRFabricInfo) SetRootPublicKey(value objc.IObject /* cross-framework: Data */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setRootPublicKey:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrfabricinfo/vendorid
func (m_ MTRFabricInfo) VendorID() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("vendorID"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrfabricinfo/vendorid
func (m_ MTRFabricInfo) SetVendorID(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setVendorID:"), value)
}



