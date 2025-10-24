// Code generated from Apple documentation for CloudKit. DO NOT EDIT.

package cloudkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class CKShareRequestAccessOperation */


/* debug [class_header]: Header for CKShareRequestAccessOperation */
// The class instance for the [CKShareRequestAccessOperation] class.
var (
	CKShareRequestAccessOperationClass     _CKShareRequestAccessOperationClass
	CKShareRequestAccessOperationClassOnce sync.Once
)

func getCKShareRequestAccessOperationClass() _CKShareRequestAccessOperationClass {
	CKShareRequestAccessOperationClassOnce.Do(func() {
		CKShareRequestAccessOperationClass = _CKShareRequestAccessOperationClass{objc.GetClass("CKShareRequestAccessOperation")}
	})
	return CKShareRequestAccessOperationClass
}

type _CKShareRequestAccessOperationClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CKShareRequestAccessOperation */
// An interface definition for the [CKShareRequestAccessOperation] class.
type ICKShareRequestAccessOperation interface {
	ICKOperation
	
/* debug [class_interface_properties]: Properties for CKShareRequestAccessOperation */
	// properties:
	PerShareAccessRequestCompletionBlock() func(unsafe.Pointer, unsafe.Pointer)
	SetPerShareAccessRequestCompletionBlock(value func(unsafe.Pointer, unsafe.Pointer))
	ShareRequestAccessCompletionBlock() func(unsafe.Pointer)
	SetShareRequestAccessCompletionBlock(value func(unsafe.Pointer))
	ShareURLs() []foundation.URL
	SetShareURLs(value []foundation.URL)
	PerShareAccessRequestResultBlock() objectivec.IObject
	SetPerShareAccessRequestResultBlock(value objectivec.IObject)
	ShareAccessRequestResultBlock() objectivec.IObject
	SetShareAccessRequestResultBlock(value objectivec.IObject)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CKShareRequestAccessOperation */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CKShareRequestAccessOperation */
// Alloc allocates a new instance without initialization.
func (cc _CKShareRequestAccessOperationClass) Alloc() CKShareRequestAccessOperation {
	rv := objc.Send[CKShareRequestAccessOperation](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CKShareRequestAccessOperationClass) New() CKShareRequestAccessOperation {
	rv := objc.Send[CKShareRequestAccessOperation](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CKShareRequestAccessOperation) Init() CKShareRequestAccessOperation {
	rv := objc.Send[CKShareRequestAccessOperation](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CKShareRequestAccessOperation) Autorelease() CKShareRequestAccessOperation {
	rv := objc.Send[CKShareRequestAccessOperation](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCKShareRequestAccessOperation creates a new CKShareRequestAccessOperation instance.
func NewCKShareRequestAccessOperation() CKShareRequestAccessOperation {
	return getCKShareRequestAccessOperationClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CKShareRequestAccessOperation */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKShareRequestAccessOperation
type CKShareRequestAccessOperation struct {
	CKOperation
}

// CKShareRequestAccessOperationFrom constructs a [CKShareRequestAccessOperation] from an unsafe.Pointer.
func CKShareRequestAccessOperationFrom(ptr unsafe.Pointer) CKShareRequestAccessOperation {
	return CKShareRequestAccessOperation{
		CKOperation: CKOperationFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CKShareRequestAccessOperation */

// Creates a share request access operation configured with specified share URLs.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKShareRequestAccessOperation/init(shareURLs:)
func NewCKShareRequestAccessOperationWithShareURLs(shareURLs []foundation.URL) CKShareRequestAccessOperation {
	instance := getCKShareRequestAccessOperationClass().Alloc()
	rv := objc.Send[CKShareRequestAccessOperation](instance.ID, objc.Sel("initWithShareURLs:"), shareURLs)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewCKShareRequestAccessOperationWithShareURLs */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CKShareRequestAccessOperation */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CKShareRequestAccessOperation */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CKShareRequestAccessOperation */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CKShareRequestAccessOperation */

// A completion block called once for each processed share URL.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKShareRequestAccessOperation/perShareAccessRequestCompletionBlock
func (c_ CKShareRequestAccessOperation) PerShareAccessRequestCompletionBlock() func(unsafe.Pointer, unsafe.Pointer) {
	rv := objc.Send[func(unsafe.Pointer, unsafe.Pointer)](c_.ID, objc.Sel("perShareAccessRequestCompletionBlock"))
	return rv
}/* debug [instance_properties/getter]: perShareAccessRequestCompletionBlock */


// A completion block called once for each processed share URL.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKShareRequestAccessOperation/perShareAccessRequestCompletionBlock
func (c_ CKShareRequestAccessOperation) SetPerShareAccessRequestCompletionBlock(value func(unsafe.Pointer, unsafe.Pointer)) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setPerShareAccessRequestCompletionBlock:"), value)
}/* debug [instance_properties/setter]: perShareAccessRequestCompletionBlock */


// A completion block called when the entire operation finishes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKShareRequestAccessOperation/shareRequestAccessCompletionBlock
func (c_ CKShareRequestAccessOperation) ShareRequestAccessCompletionBlock() func(unsafe.Pointer) {
	rv := objc.Send[func(unsafe.Pointer)](c_.ID, objc.Sel("shareRequestAccessCompletionBlock"))
	return rv
}/* debug [instance_properties/getter]: shareRequestAccessCompletionBlock */


// A completion block called when the entire operation finishes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKShareRequestAccessOperation/shareRequestAccessCompletionBlock
func (c_ CKShareRequestAccessOperation) SetShareRequestAccessCompletionBlock(value func(unsafe.Pointer)) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setShareRequestAccessCompletionBlock:"), value)
}/* debug [instance_properties/setter]: shareRequestAccessCompletionBlock */


// The URLs of the shares to request access to.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKShareRequestAccessOperation/shareURLs
func (c_ CKShareRequestAccessOperation) ShareURLs() []foundation.URL {
	rv := objc.Send[[]foundation.URL](c_.ID, objc.Sel("shareURLs"))
	return rv
}/* debug [instance_properties/getter]: shareURLs */


// The URLs of the shares to request access to.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKShareRequestAccessOperation/shareURLs
func (c_ CKShareRequestAccessOperation) SetShareURLs(value []foundation.URL) {
	var nsArray objc.ID
	if len(value) > 0 {
		nsArray = objc.ID(objc.GetClass("NSMutableArray")).Send(objc.Sel("arrayWithCapacity:"), len(value))
		for _, item := range value {
			nsArray.Send(objc.Sel("addObject:"), item)
		}
	} else {
		nsArray = objc.ID(objc.GetClass("NSArray")).Send(objc.Sel("array"))
	}
	objc.Send[objc.ID](c_.ID, objc.Sel("setShareURLs:"), nsArray)
}/* debug [instance_properties/setter]: shareURLs */


// A block called once for each share URL processed by the server.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/cksharerequestaccessoperation/pershareaccessrequestresultblock
func (c_ CKShareRequestAccessOperation) PerShareAccessRequestResultBlock() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("perShareAccessRequestResultBlock"))
	return rv
}/* debug [instance_properties/getter]: perShareAccessRequestResultBlock */


// A block called once for each share URL processed by the server.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/cksharerequestaccessoperation/pershareaccessrequestresultblock
func (c_ CKShareRequestAccessOperation) SetPerShareAccessRequestResultBlock(value objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setPerShareAccessRequestResultBlock:"), value)
}/* debug [instance_properties/setter]: perShareAccessRequestResultBlock */


// A block called when the entire share access request operation completes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/cksharerequestaccessoperation/shareaccessrequestresultblock
func (c_ CKShareRequestAccessOperation) ShareAccessRequestResultBlock() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("shareAccessRequestResultBlock"))
	return rv
}/* debug [instance_properties/getter]: shareAccessRequestResultBlock */


// A block called when the entire share access request operation completes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/cksharerequestaccessoperation/shareaccessrequestresultblock
func (c_ CKShareRequestAccessOperation) SetShareAccessRequestResultBlock(value objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setShareAccessRequestResultBlock:"), value)
}/* debug [instance_properties/setter]: shareAccessRequestResultBlock */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class CKShareRequestAccessOperation */


