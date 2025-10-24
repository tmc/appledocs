// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTRCertificateInfo */


/* debug [class_header]: Header for MTRCertificateInfo */
// The class instance for the [MTRCertificateInfo] class.
var (
	MTRCertificateInfoClass     _MTRCertificateInfoClass
	MTRCertificateInfoClassOnce sync.Once
)

func getMTRCertificateInfoClass() _MTRCertificateInfoClass {
	MTRCertificateInfoClassOnce.Do(func() {
		MTRCertificateInfoClass = _MTRCertificateInfoClass{objc.GetClass("MTRCertificateInfo")}
	})
	return MTRCertificateInfoClass
}

type _MTRCertificateInfoClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTRCertificateInfo */
// An interface definition for the [MTRCertificateInfo] class.
type IMTRCertificateInfo interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MTRCertificateInfo */
	// properties:
	Issuer() IMTRDistinguishedNameInfo
	NotAfter() objc.IObject /* cross-framework: NSDate */
	NotBefore() objc.IObject /* cross-framework: NSDate */
	PublicKeyData() objc.IObject /* cross-framework: NSData */
	Subject() IMTRDistinguishedNameInfo
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTRCertificateInfo */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTRCertificateInfo */
// Alloc allocates a new instance without initialization.
func (mc _MTRCertificateInfoClass) Alloc() MTRCertificateInfo {
	rv := objc.Send[MTRCertificateInfo](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTRCertificateInfoClass) New() MTRCertificateInfo {
	rv := objc.Send[MTRCertificateInfo](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRCertificateInfo) Init() MTRCertificateInfo {
	rv := objc.Send[MTRCertificateInfo](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRCertificateInfo) Autorelease() MTRCertificateInfo {
	rv := objc.Send[MTRCertificateInfo](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRCertificateInfo creates a new MTRCertificateInfo instance.
func NewMTRCertificateInfo() MTRCertificateInfo {
	return getMTRCertificateInfoClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTRCertificateInfo */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRCertificateInfo
type MTRCertificateInfo struct {
	objectivec.Object
}

// MTRCertificateInfoFrom constructs a [MTRCertificateInfo] from an unsafe.Pointer.
func MTRCertificateInfoFrom(ptr unsafe.Pointer) MTRCertificateInfo {
	return MTRCertificateInfo{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTRCertificateInfo */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRCertificateInfo/init(tlvBytes:)
func NewMTRCertificateInfoWithTLVBytes(bytes unsafe.Pointer) MTRCertificateInfo {
	instance := getMTRCertificateInfoClass().Alloc()
	rv := objc.Send[MTRCertificateInfo](instance.ID, objc.Sel("initWithTLVBytes:"), bytes)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewMTRCertificateInfoWithTLVBytes */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTRCertificateInfo */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTRCertificateInfo */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTRCertificateInfo */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTRCertificateInfo */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRCertificateInfo/issuer
func (m_ MTRCertificateInfo) Issuer() IMTRDistinguishedNameInfo {
	rv := objc.Send[MTRDistinguishedNameInfo](m_.ID, objc.Sel("issuer"))
	return rv
}/* debug [instance_properties/getter]: issuer */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRCertificateInfo/notAfter
func (m_ MTRCertificateInfo) NotAfter() objc.IObject /* cross-framework: NSDate */ {
	rv := objc.Send[foundation.NSDate](m_.ID, objc.Sel("notAfter"))
	return rv
}/* debug [instance_properties/getter]: notAfter */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRCertificateInfo/notBefore
func (m_ MTRCertificateInfo) NotBefore() objc.IObject /* cross-framework: NSDate */ {
	rv := objc.Send[foundation.NSDate](m_.ID, objc.Sel("notBefore"))
	return rv
}/* debug [instance_properties/getter]: notBefore */


// Public key data for this certificate
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRCertificateInfo/publicKeyData
func (m_ MTRCertificateInfo) PublicKeyData() objc.IObject /* cross-framework: NSData */ {
	rv := objc.Send[foundation.NSData](m_.ID, objc.Sel("publicKeyData"))
	return rv
}/* debug [instance_properties/getter]: publicKeyData */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRCertificateInfo/subject
func (m_ MTRCertificateInfo) Subject() IMTRDistinguishedNameInfo {
	rv := objc.Send[MTRDistinguishedNameInfo](m_.ID, objc.Sel("subject"))
	return rv
}/* debug [instance_properties/getter]: subject */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTRCertificateInfo */


