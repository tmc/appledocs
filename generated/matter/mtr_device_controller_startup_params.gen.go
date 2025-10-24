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
	// properties:
	CaseAuthenticatedTags() objc.IObject /* cross-framework: NSNumber */
	SetCaseAuthenticatedTags(value objc.IObject /* cross-framework: NSNumber */)
	FabricID() objc.IObject /* cross-framework: NSNumber */
	SetFabricID(value objc.IObject /* cross-framework: NSNumber */)
	FabricId() uint64
	SetFabricId(value uint64)
	IntermediateCertificate() objc.IObject /* cross-framework: Data */
	SetIntermediateCertificate(value objc.IObject /* cross-framework: Data */)
	Ipk() objc.IObject /* cross-framework: Data */
	SetIpk(value objc.IObject /* cross-framework: Data */)
	NocSigner() unsafe.Pointer
	SetNocSigner(value unsafe.Pointer)
	NodeID() objc.IObject /* cross-framework: NSNumber */
	SetNodeID(value objc.IObject /* cross-framework: NSNumber */)
	NodeId() objc.IObject /* cross-framework: NSNumber */
	SetNodeId(value objc.IObject /* cross-framework: NSNumber */)
	OperationalCertificate() objc.IObject /* cross-framework: Data */
	SetOperationalCertificate(value objc.IObject /* cross-framework: Data */)
	OperationalCertificateIssuer() unsafe.Pointer
	SetOperationalCertificateIssuer(value unsafe.Pointer)
	OperationalCertificateIssuerQueue() unsafe.Pointer
	SetOperationalCertificateIssuerQueue(value unsafe.Pointer)
	OperationalKeypair() unsafe.Pointer
	SetOperationalKeypair(value unsafe.Pointer)
	RootCertificate() objc.IObject /* cross-framework: Data */
	SetRootCertificate(value objc.IObject /* cross-framework: Data */)
	VendorID() objc.IObject /* cross-framework: NSNumber */
	SetVendorID(value objc.IObject /* cross-framework: NSNumber */)
	VendorId() objc.IObject /* cross-framework: NSNumber */
	SetVendorId(value objc.IObject /* cross-framework: NSNumber */)
	// methods:
}



// [Full Topic]
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



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdevicecontrollerstartupparams/caseauthenticatedtags
func (m_ MTRDeviceControllerStartupParams) CaseAuthenticatedTags() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("caseAuthenticatedTags"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdevicecontrollerstartupparams/caseauthenticatedtags
func (m_ MTRDeviceControllerStartupParams) SetCaseAuthenticatedTags(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setCaseAuthenticatedTags:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdevicecontrollerstartupparams/fabricid-1cm6z
func (m_ MTRDeviceControllerStartupParams) FabricID() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("fabricID"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdevicecontrollerstartupparams/fabricid-1cm6z
func (m_ MTRDeviceControllerStartupParams) SetFabricID(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setFabricID:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdevicecontrollerstartupparams/fabricid-1cm7v
func (m_ MTRDeviceControllerStartupParams) FabricId() uint64 {
	rv := objc.Send[uint64](m_.ID, objc.Sel("fabricId"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdevicecontrollerstartupparams/fabricid-1cm7v
func (m_ MTRDeviceControllerStartupParams) SetFabricId(value uint64) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setFabricId:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdevicecontrollerstartupparams/intermediatecertificate
func (m_ MTRDeviceControllerStartupParams) IntermediateCertificate() objc.IObject /* cross-framework: Data */ {
	rv := objc.Send[foundation.Data](m_.ID, objc.Sel("intermediateCertificate"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdevicecontrollerstartupparams/intermediatecertificate
func (m_ MTRDeviceControllerStartupParams) SetIntermediateCertificate(value objc.IObject /* cross-framework: Data */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setIntermediateCertificate:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdevicecontrollerstartupparams/ipk
func (m_ MTRDeviceControllerStartupParams) Ipk() objc.IObject /* cross-framework: Data */ {
	rv := objc.Send[foundation.Data](m_.ID, objc.Sel("ipk"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdevicecontrollerstartupparams/ipk
func (m_ MTRDeviceControllerStartupParams) SetIpk(value objc.IObject /* cross-framework: Data */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setIpk:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdevicecontrollerstartupparams/nocsigner
func (m_ MTRDeviceControllerStartupParams) NocSigner() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("nocSigner"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdevicecontrollerstartupparams/nocsigner
func (m_ MTRDeviceControllerStartupParams) SetNocSigner(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setNocSigner:"), value)
}


// Node id for this controller.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdevicecontrollerstartupparams/nodeid-9iwwv
func (m_ MTRDeviceControllerStartupParams) NodeID() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("nodeID"))
	return rv
}


// Node id for this controller.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdevicecontrollerstartupparams/nodeid-9iwwv
func (m_ MTRDeviceControllerStartupParams) SetNodeID(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setNodeID:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdevicecontrollerstartupparams/nodeid-9iwxr
func (m_ MTRDeviceControllerStartupParams) NodeId() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("nodeId"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdevicecontrollerstartupparams/nodeid-9iwxr
func (m_ MTRDeviceControllerStartupParams) SetNodeId(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setNodeId:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdevicecontrollerstartupparams/operationalcertificate
func (m_ MTRDeviceControllerStartupParams) OperationalCertificate() objc.IObject /* cross-framework: Data */ {
	rv := objc.Send[foundation.Data](m_.ID, objc.Sel("operationalCertificate"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdevicecontrollerstartupparams/operationalcertificate
func (m_ MTRDeviceControllerStartupParams) SetOperationalCertificate(value objc.IObject /* cross-framework: Data */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setOperationalCertificate:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdevicecontrollerstartupparams/operationalcertificateissuer
func (m_ MTRDeviceControllerStartupParams) OperationalCertificateIssuer() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("operationalCertificateIssuer"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdevicecontrollerstartupparams/operationalcertificateissuer
func (m_ MTRDeviceControllerStartupParams) SetOperationalCertificateIssuer(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setOperationalCertificateIssuer:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdevicecontrollerstartupparams/operationalcertificateissuerqueue
func (m_ MTRDeviceControllerStartupParams) OperationalCertificateIssuerQueue() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("operationalCertificateIssuerQueue"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdevicecontrollerstartupparams/operationalcertificateissuerqueue
func (m_ MTRDeviceControllerStartupParams) SetOperationalCertificateIssuerQueue(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setOperationalCertificateIssuerQueue:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdevicecontrollerstartupparams/operationalkeypair
func (m_ MTRDeviceControllerStartupParams) OperationalKeypair() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("operationalKeypair"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdevicecontrollerstartupparams/operationalkeypair
func (m_ MTRDeviceControllerStartupParams) SetOperationalKeypair(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setOperationalKeypair:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdevicecontrollerstartupparams/rootcertificate
func (m_ MTRDeviceControllerStartupParams) RootCertificate() objc.IObject /* cross-framework: Data */ {
	rv := objc.Send[foundation.Data](m_.ID, objc.Sel("rootCertificate"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdevicecontrollerstartupparams/rootcertificate
func (m_ MTRDeviceControllerStartupParams) SetRootCertificate(value objc.IObject /* cross-framework: Data */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setRootCertificate:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdevicecontrollerstartupparams/vendorid-8ru1s
func (m_ MTRDeviceControllerStartupParams) VendorID() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("vendorID"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdevicecontrollerstartupparams/vendorid-8ru1s
func (m_ MTRDeviceControllerStartupParams) SetVendorID(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setVendorID:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdevicecontrollerstartupparams/vendorid-8ru0w
func (m_ MTRDeviceControllerStartupParams) VendorId() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("vendorId"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdevicecontrollerstartupparams/vendorid-8ru0w
func (m_ MTRDeviceControllerStartupParams) SetVendorId(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setVendorId:"), value)
}



