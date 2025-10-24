// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class CSRInfo */


/* debug [class_header]: Header for CSRInfo */
// The class instance for the [CSRInfo] class.
var (
	CSRInfoClass     _CSRInfoClass
	CSRInfoClassOnce sync.Once
)

func getCSRInfoClass() _CSRInfoClass {
	CSRInfoClassOnce.Do(func() {
		CSRInfoClass = _CSRInfoClass{objc.GetClass("CSRInfo")}
	})
	return CSRInfoClass
}

type _CSRInfoClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CSRInfo */
// An interface definition for the [CSRInfo] class.
type ICSRInfo interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for CSRInfo */
	// properties:
	Csr() objc.IObject /* cross-framework: NSData */
	SetCsr(value objc.IObject /* cross-framework: NSData */)
	Elements() objc.IObject /* cross-framework: NSData */
	SetElements(value objc.IObject /* cross-framework: NSData */)
	ElementsSignature() objc.IObject /* cross-framework: NSData */
	SetElementsSignature(value objc.IObject /* cross-framework: NSData */)
	Nonce() objc.IObject /* cross-framework: NSData */
	SetNonce(value objc.IObject /* cross-framework: NSData */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CSRInfo */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CSRInfo */
// Alloc allocates a new instance without initialization.
func (cc _CSRInfoClass) Alloc() CSRInfo {
	rv := objc.Send[CSRInfo](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CSRInfoClass) New() CSRInfo {
	rv := objc.Send[CSRInfo](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CSRInfo) Init() CSRInfo {
	rv := objc.Send[CSRInfo](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CSRInfo) Autorelease() CSRInfo {
	rv := objc.Send[CSRInfo](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCSRInfo creates a new CSRInfo instance.
func NewCSRInfo() CSRInfo {
	return getCSRInfoClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CSRInfo */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/CSRInfo
type CSRInfo struct {
	objectivec.Object
}

// CSRInfoFrom constructs a [CSRInfo] from an unsafe.Pointer.
func CSRInfoFrom(ptr unsafe.Pointer) CSRInfo {
	return CSRInfo{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CSRInfo */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/CSRInfo/init(nonce:elements:elementsSignature:csr:)
func NewCSRInfoWithNonceElementsElementsSignatureCsr(nonce objc.IObject /* cross-framework: NSData */, elements objc.IObject /* cross-framework: NSData */, elementsSignature objc.IObject /* cross-framework: NSData */, csr objc.IObject /* cross-framework: NSData */) CSRInfo {
	instance := getCSRInfoClass().Alloc()
	rv := objc.Send[CSRInfo](instance.ID, objc.Sel("initWithNonce:elements:elementsSignature:csr:"), nonce, elements, elementsSignature, csr)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewCSRInfoWithNonceElementsElementsSignatureCsr */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CSRInfo */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CSRInfo */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CSRInfo */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CSRInfo */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/CSRInfo/csr
func (c_ CSRInfo) Csr() objc.IObject /* cross-framework: NSData */ {
	rv := objc.Send[foundation.NSData](c_.ID, objc.Sel("csr"))
	return rv
}/* debug [instance_properties/getter]: csr */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/CSRInfo/csr
func (c_ CSRInfo) SetCsr(value objc.IObject /* cross-framework: NSData */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setCsr:"), value)
}/* debug [instance_properties/setter]: csr */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/CSRInfo/elements
func (c_ CSRInfo) Elements() objc.IObject /* cross-framework: NSData */ {
	rv := objc.Send[foundation.NSData](c_.ID, objc.Sel("elements"))
	return rv
}/* debug [instance_properties/getter]: elements */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/CSRInfo/elements
func (c_ CSRInfo) SetElements(value objc.IObject /* cross-framework: NSData */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setElements:"), value)
}/* debug [instance_properties/setter]: elements */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/CSRInfo/elementsSignature
func (c_ CSRInfo) ElementsSignature() objc.IObject /* cross-framework: NSData */ {
	rv := objc.Send[foundation.NSData](c_.ID, objc.Sel("elementsSignature"))
	return rv
}/* debug [instance_properties/getter]: elementsSignature */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/CSRInfo/elementsSignature
func (c_ CSRInfo) SetElementsSignature(value objc.IObject /* cross-framework: NSData */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setElementsSignature:"), value)
}/* debug [instance_properties/setter]: elementsSignature */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/CSRInfo/nonce
func (c_ CSRInfo) Nonce() objc.IObject /* cross-framework: NSData */ {
	rv := objc.Send[foundation.NSData](c_.ID, objc.Sel("nonce"))
	return rv
}/* debug [instance_properties/getter]: nonce */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/CSRInfo/nonce
func (c_ CSRInfo) SetNonce(value objc.IObject /* cross-framework: NSData */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setNonce:"), value)
}/* debug [instance_properties/setter]: nonce */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class CSRInfo */


