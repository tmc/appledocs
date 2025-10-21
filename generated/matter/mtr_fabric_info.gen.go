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
}

//
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


//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrfabricinfo/fabricid
func (m_ MTRFabricInfo) FabricID() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("fabricID"))
	return rv
}


// SetFabricID sets the value of the fabricID property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrfabricinfo/fabricid
func (m_ MTRFabricInfo) SetFabricID(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setFabricID:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrfabricinfo/fabricindex
func (m_ MTRFabricInfo) FabricIndex() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("fabricIndex"))
	return rv
}


// SetFabricIndex sets the value of the fabricIndex property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrfabricinfo/fabricindex
func (m_ MTRFabricInfo) SetFabricIndex(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setFabricIndex:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrfabricinfo/intermediatecertificate
func (m_ MTRFabricInfo) IntermediateCertificate() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("intermediateCertificate"))
	return rv
}


// SetIntermediateCertificate sets the value of the intermediateCertificate property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrfabricinfo/intermediatecertificate
func (m_ MTRFabricInfo) SetIntermediateCertificate(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setIntermediateCertificate:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrfabricinfo/intermediatecertificatetlv
func (m_ MTRFabricInfo) IntermediateCertificateTLV() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("intermediateCertificateTLV"))
	return rv
}


// SetIntermediateCertificateTLV sets the value of the intermediateCertificateTLV property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrfabricinfo/intermediatecertificatetlv
func (m_ MTRFabricInfo) SetIntermediateCertificateTLV(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setIntermediateCertificateTLV:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrfabricinfo/label
func (m_ MTRFabricInfo) Label() string {
	rv := objc.Send[string](m_.ID, objc.Sel("label"))
	return rv
}


// SetLabel sets the value of the label property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrfabricinfo/label
func (m_ MTRFabricInfo) SetLabel(value string) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setLabel:"), objc.String(value))
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrfabricinfo/nodeid
func (m_ MTRFabricInfo) NodeID() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("nodeID"))
	return rv
}


// SetNodeID sets the value of the nodeID property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrfabricinfo/nodeid
func (m_ MTRFabricInfo) SetNodeID(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setNodeID:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrfabricinfo/operationalcertificate
func (m_ MTRFabricInfo) OperationalCertificate() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("operationalCertificate"))
	return rv
}


// SetOperationalCertificate sets the value of the operationalCertificate property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrfabricinfo/operationalcertificate
func (m_ MTRFabricInfo) SetOperationalCertificate(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setOperationalCertificate:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrfabricinfo/operationalcertificatetlv
func (m_ MTRFabricInfo) OperationalCertificateTLV() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("operationalCertificateTLV"))
	return rv
}


// SetOperationalCertificateTLV sets the value of the operationalCertificateTLV property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrfabricinfo/operationalcertificatetlv
func (m_ MTRFabricInfo) SetOperationalCertificateTLV(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setOperationalCertificateTLV:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrfabricinfo/rootcertificate
func (m_ MTRFabricInfo) RootCertificate() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("rootCertificate"))
	return rv
}


// SetRootCertificate sets the value of the rootCertificate property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrfabricinfo/rootcertificate
func (m_ MTRFabricInfo) SetRootCertificate(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setRootCertificate:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrfabricinfo/rootcertificatetlv
func (m_ MTRFabricInfo) RootCertificateTLV() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("rootCertificateTLV"))
	return rv
}


// SetRootCertificateTLV sets the value of the rootCertificateTLV property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrfabricinfo/rootcertificatetlv
func (m_ MTRFabricInfo) SetRootCertificateTLV(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setRootCertificateTLV:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrfabricinfo/rootpublickey
func (m_ MTRFabricInfo) RootPublicKey() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("rootPublicKey"))
	return rv
}


// SetRootPublicKey sets the value of the rootPublicKey property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrfabricinfo/rootpublickey
func (m_ MTRFabricInfo) SetRootPublicKey(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setRootPublicKey:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrfabricinfo/vendorid
func (m_ MTRFabricInfo) VendorID() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("vendorID"))
	return rv
}


// SetVendorID sets the value of the vendorID property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrfabricinfo/vendorid
func (m_ MTRFabricInfo) SetVendorID(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setVendorID:"), value)
}



