// Code generated from Apple documentation for CoreTelephony. DO NOT EDIT.

package coretelephony

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class CTCellularPlanStatus */


/* debug [class_header]: Header for CTCellularPlanStatus */
// The class instance for the [CellularPlanStatus] class.
var (
	CellularPlanStatusClass     _CellularPlanStatusClass
	CellularPlanStatusClassOnce sync.Once
)

func getCellularPlanStatusClass() _CellularPlanStatusClass {
	CellularPlanStatusClassOnce.Do(func() {
		CellularPlanStatusClass = _CellularPlanStatusClass{objc.GetClass("CTCellularPlanStatus")}
	})
	return CellularPlanStatusClass
}

type _CellularPlanStatusClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CellularPlanStatus */
// An interface definition for the [CellularPlanStatus] class.
type ICellularPlanStatus interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for CellularPlanStatus */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CellularPlanStatus */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CellularPlanStatus */
// Alloc allocates a new instance without initialization.
func (cc _CellularPlanStatusClass) Alloc() CellularPlanStatus {
	rv := objc.Send[CellularPlanStatus](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CellularPlanStatusClass) New() CellularPlanStatus {
	rv := objc.Send[CellularPlanStatus](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CellularPlanStatus) Init() CellularPlanStatus {
	rv := objc.Send[CellularPlanStatus](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CellularPlanStatus) Autorelease() CellularPlanStatus {
	rv := objc.Send[CellularPlanStatus](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCellularPlanStatus creates a new CellularPlanStatus instance.
func NewCellularPlanStatus() CellularPlanStatus {
	return getCellularPlanStatusClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CellularPlanStatus */
// An object used for retrieving and checking the validity of a token.
//
// checks if the Integrated Circuit Card Identifier (ICCID) on a device is associated with a given token. Use the method to configure the instance of a view for Unified Payments Interface (UPI) device validation. This process generates your token, which you can use to help determine if there are any changes to the underlying ICCID. After you generate a token, use to retrieve the token associated with your app. Your app has 30 seconds to call before the system invalidates the token. If called in time, the framework sends and stores a token associated with your app. The token is mapped to the ICCID associated with the original instance. Use to check the status of the token. If the ICCID is present and turned-on, returns . For more information on configuring an instance for UPI device validation, see .


// An object used for retrieving and checking the validity of a token.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreTelephony/CTCellularPlanStatus
type CellularPlanStatus struct {
	objectivec.Object
}

// CellularPlanStatusFrom constructs a [CellularPlanStatus] from an unsafe.Pointer.
//
// An object used for retrieving and checking the validity of a token.
func CellularPlanStatusFrom(ptr unsafe.Pointer) CellularPlanStatus {
	return CellularPlanStatus{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CellularPlanStatus *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CellularPlanStatus */

// Checks for a valid ICCID associated with the token.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreTelephony/CTCellularPlanStatus/checkValidity(ofToken:completionHandler:)
func (cc _CellularPlanStatusClass) CheckValidityOfTokenCompletionHandler(token objc.IObject /* cross-framework: NSString */, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(cc.class), objc.Sel("checkValidityOfToken:completionHandler:"), token, completionHandler)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=CheckValidityOfTokenCompletionHandler) */


// A method you use to retrieve and store the token.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreTelephony/CTCellularPlanStatus/getTokenWithCompletion(_:)
func (cc _CellularPlanStatusClass) GetTokenWithCompletion(completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(cc.class), objc.Sel("getTokenWithCompletion:"), completionHandler)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=GetTokenWithCompletion) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CellularPlanStatus */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CellularPlanStatus */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CellularPlanStatus */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class CTCellularPlanStatus */


