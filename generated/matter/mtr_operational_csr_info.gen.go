// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTROperationalCSRInfo */


/* debug [class_header]: Header for MTROperationalCSRInfo */
// The class instance for the [MTROperationalCSRInfo] class.
var (
	MTROperationalCSRInfoClass     _MTROperationalCSRInfoClass
	MTROperationalCSRInfoClassOnce sync.Once
)

func getMTROperationalCSRInfoClass() _MTROperationalCSRInfoClass {
	MTROperationalCSRInfoClassOnce.Do(func() {
		MTROperationalCSRInfoClass = _MTROperationalCSRInfoClass{objc.GetClass("MTROperationalCSRInfo")}
	})
	return MTROperationalCSRInfoClass
}

type _MTROperationalCSRInfoClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTROperationalCSRInfo */
// An interface definition for the [MTROperationalCSRInfo] class.
type IMTROperationalCSRInfo interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MTROperationalCSRInfo */
	// properties:
	AttestationSignature() objc.IObject /* cross-framework: NSData */
	Csr() unsafe.Pointer
	CsrElementsTLV() unsafe.Pointer
	CsrNonce() objc.IObject /* cross-framework: NSData */
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTROperationalCSRInfo */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTROperationalCSRInfo */
// Alloc allocates a new instance without initialization.
func (mc _MTROperationalCSRInfoClass) Alloc() MTROperationalCSRInfo {
	rv := objc.Send[MTROperationalCSRInfo](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTROperationalCSRInfoClass) New() MTROperationalCSRInfo {
	rv := objc.Send[MTROperationalCSRInfo](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTROperationalCSRInfo) Init() MTROperationalCSRInfo {
	rv := objc.Send[MTROperationalCSRInfo](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTROperationalCSRInfo) Autorelease() MTROperationalCSRInfo {
	rv := objc.Send[MTROperationalCSRInfo](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTROperationalCSRInfo creates a new MTROperationalCSRInfo instance.
func NewMTROperationalCSRInfo() MTROperationalCSRInfo {
	return getMTROperationalCSRInfoClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTROperationalCSRInfo */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROperationalCSRInfo
type MTROperationalCSRInfo struct {
	objectivec.Object
}

// MTROperationalCSRInfoFrom constructs a [MTROperationalCSRInfo] from an unsafe.Pointer.
func MTROperationalCSRInfoFrom(ptr unsafe.Pointer) MTROperationalCSRInfo {
	return MTROperationalCSRInfo{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTROperationalCSRInfo */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROperationalCSRInfo/init(csr:csrNonce:csrElementsTLV:attestationSignature:)
func NewMTROperationalCSRInfoWithCSRCsrNonceCsrElementsTLVAttestationSignature(csr unsafe.Pointer, csrNonce objc.IObject /* cross-framework: NSData */, csrElementsTLV unsafe.Pointer, attestationSignature objc.IObject /* cross-framework: NSData */) MTROperationalCSRInfo {
	instance := getMTROperationalCSRInfoClass().Alloc()
	rv := objc.Send[MTROperationalCSRInfo](instance.ID, objc.Sel("initWithCSR:csrNonce:csrElementsTLV:attestationSignature:"), csr, csrNonce, csrElementsTLV, attestationSignature)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewMTROperationalCSRInfoWithCSRCsrNonceCsrElementsTLVAttestationSignature */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROperationalCSRInfo/init(csrElementsTLV:attestationSignature:)
func NewMTROperationalCSRInfoWithCSRElementsTLVAttestationSignature(csrElementsTLV unsafe.Pointer, attestationSignature objc.IObject /* cross-framework: NSData */) MTROperationalCSRInfo {
	instance := getMTROperationalCSRInfoClass().Alloc()
	rv := objc.Send[MTROperationalCSRInfo](instance.ID, objc.Sel("initWithCSRElementsTLV:attestationSignature:"), csrElementsTLV, attestationSignature)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewMTROperationalCSRInfoWithCSRElementsTLVAttestationSignature */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROperationalCSRInfo/init(csrNonce:csrElementsTLV:attestationSignature:)
func NewMTROperationalCSRInfoWithCSRNonceCsrElementsTLVAttestationSignature(csrNonce objc.IObject /* cross-framework: NSData */, csrElementsTLV unsafe.Pointer, attestationSignature objc.IObject /* cross-framework: NSData */) MTROperationalCSRInfo {
	instance := getMTROperationalCSRInfoClass().Alloc()
	rv := objc.Send[MTROperationalCSRInfo](instance.ID, objc.Sel("initWithCSRNonce:csrElementsTLV:attestationSignature:"), csrNonce, csrElementsTLV, attestationSignature)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewMTROperationalCSRInfoWithCSRNonceCsrElementsTLVAttestationSignature */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROperationalCSRInfo/init(csrResponseParams:)
func NewMTROperationalCSRInfoWithCSRResponseParams(responseParams IMTROperationalCredentialsClusterCSRResponseParams) MTROperationalCSRInfo {
	instance := getMTROperationalCSRInfoClass().Alloc()
	rv := objc.Send[MTROperationalCSRInfo](instance.ID, objc.Sel("initWithCSRResponseParams:"), responseParams)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewMTROperationalCSRInfoWithCSRResponseParams */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTROperationalCSRInfo */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTROperationalCSRInfo */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTROperationalCSRInfo */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTROperationalCSRInfo */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROperationalCSRInfo/attestationSignature
func (m_ MTROperationalCSRInfo) AttestationSignature() objc.IObject /* cross-framework: NSData */ {
	rv := objc.Send[foundation.NSData](m_.ID, objc.Sel("attestationSignature"))
	return rv
}/* debug [instance_properties/getter]: attestationSignature */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROperationalCSRInfo/csr
func (m_ MTROperationalCSRInfo) Csr() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("csr"))
	return rv
}/* debug [instance_properties/getter]: csr */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROperationalCSRInfo/csrElementsTLV
func (m_ MTROperationalCSRInfo) CsrElementsTLV() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("csrElementsTLV"))
	return rv
}/* debug [instance_properties/getter]: csrElementsTLV */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROperationalCSRInfo/csrNonce
func (m_ MTROperationalCSRInfo) CsrNonce() objc.IObject /* cross-framework: NSData */ {
	rv := objc.Send[foundation.NSData](m_.ID, objc.Sel("csrNonce"))
	return rv
}/* debug [instance_properties/getter]: csrNonce */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTROperationalCSRInfo */


