// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRDeviceControllerStartupParams] class.
var (
	MTRDeviceControllerStartupParamsClass     _MTRDeviceControllerStartupParamsClass
	MTRDeviceControllerStartupParamsClassOnce sync.Once
)

func getMTRDeviceControllerStartupParamsClass() _MTRDeviceControllerStartupParamsClass {
	MTRDeviceControllerStartupParamsClassOnce.Do(func() {
		MTRDeviceControllerStartupParamsClass = _MTRDeviceControllerStartupParamsClass{objc.GetClass("MTRDeviceControllerStartupParams")}
	})
	return MTRDeviceControllerStartupParamsClass
}

type _MTRDeviceControllerStartupParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTRDeviceControllerStartupParams] class.
type IMTRDeviceControllerStartupParams interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceControllerStartupParams
type MTRDeviceControllerStartupParams struct {
	objectivec.Object
}

// MTRDeviceControllerStartupParamsFrom constructs a [MTRDeviceControllerStartupParams] from an unsafe.Pointer.
func MTRDeviceControllerStartupParamsFrom(ptr unsafe.Pointer) MTRDeviceControllerStartupParams {
	return MTRDeviceControllerStartupParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRDeviceControllerStartupParamsClass) Alloc() MTRDeviceControllerStartupParams {
	rv := objc.Send[MTRDeviceControllerStartupParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRDeviceControllerStartupParamsClass) New() MTRDeviceControllerStartupParams {
	rv := objc.Send[MTRDeviceControllerStartupParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRDeviceControllerStartupParams) Init() MTRDeviceControllerStartupParams {
	rv := objc.Send[MTRDeviceControllerStartupParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRDeviceControllerStartupParams) Autorelease() MTRDeviceControllerStartupParams {
	rv := objc.Send[MTRDeviceControllerStartupParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRDeviceControllerStartupParams creates a new MTRDeviceControllerStartupParams instance.
func NewMTRDeviceControllerStartupParams() MTRDeviceControllerStartupParams {
	return getMTRDeviceControllerStartupParamsClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdevicecontrollerstartupparams/vendorid-8ru1s
func (m_ MTRDeviceControllerStartupParams) VendorID() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("vendorID"))
	return rv
}


// SetVendorID sets the value of the vendorID property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdevicecontrollerstartupparams/vendorid-8ru1s
func (m_ MTRDeviceControllerStartupParams) SetVendorID(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setVendorID:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdevicecontrollerstartupparams/nodeid-9iwxr
func (m_ MTRDeviceControllerStartupParams) NodeId() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("nodeId"))
	return rv
}


// SetNodeId sets the value of the nodeId property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdevicecontrollerstartupparams/nodeid-9iwxr
func (m_ MTRDeviceControllerStartupParams) SetNodeId(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setNodeId:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdevicecontrollerstartupparams/fabricid-1cm6z
func (m_ MTRDeviceControllerStartupParams) FabricID() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("fabricID"))
	return rv
}


// SetFabricID sets the value of the fabricID property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdevicecontrollerstartupparams/fabricid-1cm6z
func (m_ MTRDeviceControllerStartupParams) SetFabricID(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setFabricID:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdevicecontrollerstartupparams/operationalkeypair
func (m_ MTRDeviceControllerStartupParams) OperationalKeypair() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("operationalKeypair"))
	return rv
}


// SetOperationalKeypair sets the value of the operationalKeypair property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdevicecontrollerstartupparams/operationalkeypair
func (m_ MTRDeviceControllerStartupParams) SetOperationalKeypair(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setOperationalKeypair:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdevicecontrollerstartupparams/nocsigner
func (m_ MTRDeviceControllerStartupParams) NocSigner() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("nocSigner"))
	return rv
}


// SetNocSigner sets the value of the nocSigner property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdevicecontrollerstartupparams/nocsigner
func (m_ MTRDeviceControllerStartupParams) SetNocSigner(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setNocSigner:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdevicecontrollerstartupparams/vendorid-8ru0w
func (m_ MTRDeviceControllerStartupParams) VendorId() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("vendorId"))
	return rv
}


// SetVendorId sets the value of the vendorId property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdevicecontrollerstartupparams/vendorid-8ru0w
func (m_ MTRDeviceControllerStartupParams) SetVendorId(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setVendorId:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdevicecontrollerstartupparams/rootcertificate
func (m_ MTRDeviceControllerStartupParams) RootCertificate() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("rootCertificate"))
	return rv
}


// SetRootCertificate sets the value of the rootCertificate property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdevicecontrollerstartupparams/rootcertificate
func (m_ MTRDeviceControllerStartupParams) SetRootCertificate(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setRootCertificate:"), value)
}

// Node id for this controller.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdevicecontrollerstartupparams/nodeid-9iwwv
func (m_ MTRDeviceControllerStartupParams) NodeID() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("nodeID"))
	return rv
}


// SetNodeID sets the value of the nodeID property.
// Node id for this controller.

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdevicecontrollerstartupparams/nodeid-9iwwv
func (m_ MTRDeviceControllerStartupParams) SetNodeID(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setNodeID:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdevicecontrollerstartupparams/caseauthenticatedtags
func (m_ MTRDeviceControllerStartupParams) CaseAuthenticatedTags() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("caseAuthenticatedTags"))
	return rv
}


// SetCaseAuthenticatedTags sets the value of the caseAuthenticatedTags property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdevicecontrollerstartupparams/caseauthenticatedtags
func (m_ MTRDeviceControllerStartupParams) SetCaseAuthenticatedTags(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setCaseAuthenticatedTags:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdevicecontrollerstartupparams/intermediatecertificate
func (m_ MTRDeviceControllerStartupParams) IntermediateCertificate() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("intermediateCertificate"))
	return rv
}


// SetIntermediateCertificate sets the value of the intermediateCertificate property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdevicecontrollerstartupparams/intermediatecertificate
func (m_ MTRDeviceControllerStartupParams) SetIntermediateCertificate(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setIntermediateCertificate:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdevicecontrollerstartupparams/ipk
func (m_ MTRDeviceControllerStartupParams) Ipk() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("ipk"))
	return rv
}


// SetIpk sets the value of the ipk property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdevicecontrollerstartupparams/ipk
func (m_ MTRDeviceControllerStartupParams) SetIpk(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setIpk:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdevicecontrollerstartupparams/operationalcertificateissuerqueue
func (m_ MTRDeviceControllerStartupParams) OperationalCertificateIssuerQueue() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("operationalCertificateIssuerQueue"))
	return rv
}


// SetOperationalCertificateIssuerQueue sets the value of the operationalCertificateIssuerQueue property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdevicecontrollerstartupparams/operationalcertificateissuerqueue
func (m_ MTRDeviceControllerStartupParams) SetOperationalCertificateIssuerQueue(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setOperationalCertificateIssuerQueue:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdevicecontrollerstartupparams/operationalcertificateissuer
func (m_ MTRDeviceControllerStartupParams) OperationalCertificateIssuer() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("operationalCertificateIssuer"))
	return rv
}


// SetOperationalCertificateIssuer sets the value of the operationalCertificateIssuer property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdevicecontrollerstartupparams/operationalcertificateissuer
func (m_ MTRDeviceControllerStartupParams) SetOperationalCertificateIssuer(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setOperationalCertificateIssuer:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdevicecontrollerstartupparams/operationalcertificate
func (m_ MTRDeviceControllerStartupParams) OperationalCertificate() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("operationalCertificate"))
	return rv
}


// SetOperationalCertificate sets the value of the operationalCertificate property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdevicecontrollerstartupparams/operationalcertificate
func (m_ MTRDeviceControllerStartupParams) SetOperationalCertificate(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setOperationalCertificate:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdevicecontrollerstartupparams/fabricid-1cm7v
func (m_ MTRDeviceControllerStartupParams) FabricId() uint64 {
	rv := objc.Send[uint64](m_.ID, objc.Sel("fabricId"))
	return rv
}


// SetFabricId sets the value of the fabricId property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdevicecontrollerstartupparams/fabricid-1cm7v
func (m_ MTRDeviceControllerStartupParams) SetFabricId(value uint64) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setFabricId:"), value)
}



