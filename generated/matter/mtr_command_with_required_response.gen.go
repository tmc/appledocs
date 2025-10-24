// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTRCommandWithRequiredResponse */


/* debug [class_header]: Header for MTRCommandWithRequiredResponse */
// The class instance for the [MTRCommandWithRequiredResponse] class.
var (
	MTRCommandWithRequiredResponseClass     _MTRCommandWithRequiredResponseClass
	MTRCommandWithRequiredResponseClassOnce sync.Once
)

func getMTRCommandWithRequiredResponseClass() _MTRCommandWithRequiredResponseClass {
	MTRCommandWithRequiredResponseClassOnce.Do(func() {
		MTRCommandWithRequiredResponseClass = _MTRCommandWithRequiredResponseClass{objc.GetClass("MTRCommandWithRequiredResponse")}
	})
	return MTRCommandWithRequiredResponseClass
}

type _MTRCommandWithRequiredResponseClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTRCommandWithRequiredResponse */
// An interface definition for the [MTRCommandWithRequiredResponse] class.
type IMTRCommandWithRequiredResponse interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MTRCommandWithRequiredResponse */
	// properties:
	CommandFields() objc.IObject /* cross-framework: NSString */
	SetCommandFields(value objc.IObject /* cross-framework: NSString */)
	Path() IMTRCommandPath
	SetPath(value IMTRCommandPath)
	RequiredResponse() objc.IObject /* cross-framework: NSString */
	SetRequiredResponse(value objc.IObject /* cross-framework: NSString */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTRCommandWithRequiredResponse */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTRCommandWithRequiredResponse */
// Alloc allocates a new instance without initialization.
func (mc _MTRCommandWithRequiredResponseClass) Alloc() MTRCommandWithRequiredResponse {
	rv := objc.Send[MTRCommandWithRequiredResponse](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTRCommandWithRequiredResponseClass) New() MTRCommandWithRequiredResponse {
	rv := objc.Send[MTRCommandWithRequiredResponse](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRCommandWithRequiredResponse) Init() MTRCommandWithRequiredResponse {
	rv := objc.Send[MTRCommandWithRequiredResponse](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRCommandWithRequiredResponse) Autorelease() MTRCommandWithRequiredResponse {
	rv := objc.Send[MTRCommandWithRequiredResponse](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRCommandWithRequiredResponse creates a new MTRCommandWithRequiredResponse instance.
func NewMTRCommandWithRequiredResponse() MTRCommandWithRequiredResponse {
	return getMTRCommandWithRequiredResponseClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTRCommandWithRequiredResponse */
// An object representing a single command to be invoked and the response required for the invoke to be considered successful.


// An object representing a single command to be invoked and the response required for the invoke to be considered successful.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRCommandWithRequiredResponse
type MTRCommandWithRequiredResponse struct {
	objectivec.Object
}

// MTRCommandWithRequiredResponseFrom constructs a [MTRCommandWithRequiredResponse] from an unsafe.Pointer.
//
// An object representing a single command to be invoked and the response required for the invoke to be considered successful.
func MTRCommandWithRequiredResponseFrom(ptr unsafe.Pointer) MTRCommandWithRequiredResponse {
	return MTRCommandWithRequiredResponse{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTRCommandWithRequiredResponse */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRCommandWithRequiredResponse/init(path:commandFields:requiredResponse:)
func NewMTRCommandWithRequiredResponseWithPathCommandFieldsRequiredResponse(path IMTRCommandPath, commandFields foundation.IDictionary, requiredResponse foundation.IDictionary) MTRCommandWithRequiredResponse {
	instance := getMTRCommandWithRequiredResponseClass().Alloc()
	rv := objc.Send[MTRCommandWithRequiredResponse](instance.ID, objc.Sel("initWithPath:commandFields:requiredResponse:"), path, commandFields, requiredResponse)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewMTRCommandWithRequiredResponseWithPathCommandFieldsRequiredResponse */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTRCommandWithRequiredResponse */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTRCommandWithRequiredResponse */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTRCommandWithRequiredResponse */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTRCommandWithRequiredResponse */

// The command fields to pass for the command invoke. nil if this command does
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcommandwithrequiredresponse/commandfields
func (m_ MTRCommandWithRequiredResponse) CommandFields() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("commandFields"))
	return rv
}/* debug [instance_properties/getter]: commandFields */


// The command fields to pass for the command invoke. nil if this command does
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcommandwithrequiredresponse/commandfields
func (m_ MTRCommandWithRequiredResponse) SetCommandFields(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setCommandFields:"), value)
}/* debug [instance_properties/setter]: commandFields */


// The path of the command being invoked.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcommandwithrequiredresponse/path
func (m_ MTRCommandWithRequiredResponse) Path() IMTRCommandPath {
	rv := objc.Send[MTRCommandPath](m_.ID, objc.Sel("path"))
	return rv
}/* debug [instance_properties/getter]: path */


// The path of the command being invoked.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcommandwithrequiredresponse/path
func (m_ MTRCommandWithRequiredResponse) SetPath(value IMTRCommandPath) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setPath:"), value)
}/* debug [instance_properties/setter]: path */


// The response that represents this command succeeding.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcommandwithrequiredresponse/requiredresponse
func (m_ MTRCommandWithRequiredResponse) RequiredResponse() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("requiredResponse"))
	return rv
}/* debug [instance_properties/getter]: requiredResponse */


// The response that represents this command succeeding.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrcommandwithrequiredresponse/requiredresponse
func (m_ MTRCommandWithRequiredResponse) SetRequiredResponse(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setRequiredResponse:"), value)
}/* debug [instance_properties/setter]: requiredResponse */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTRCommandWithRequiredResponse */


