// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

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

// An interface definition for the [MTRCommandWithRequiredResponse] class.
type IMTRCommandWithRequiredResponse interface {
	objectivec.IObject
}

// An object representing a single command to be invoked and the response required for the invoke to be considered successful.
//
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

// Alloc allocates a new instance without initialization.
func (mc _MTRCommandWithRequiredResponseClass) Alloc() MTRCommandWithRequiredResponse {
	rv := objc.Send[MTRCommandWithRequiredResponse](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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


//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRCommandWithRequiredResponse/init(path:commandFields:requiredResponse:)
func NewMTRCommandWithRequiredResponseWithPathCommandFieldsRequiredResponse(path unsafe.Pointer, commandFields unsafe.Pointer, requiredResponse unsafe.Pointer) MTRCommandWithRequiredResponse {
	instance := getMTRCommandWithRequiredResponseClass().Alloc()
	rv := objc.Send[MTRCommandWithRequiredResponse](instance.ID, objc.Sel("initWithPath:commandFields:requiredResponse:"), path, commandFields, requiredResponse)
	rv.Autorelease()
	return rv
}


// The command fields to pass for the command invoke. nil if this command does not have any fields. If not nil, this should be a data-value dictionary of MTRStructureValueType.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRCommandWithRequiredResponse/commandFields
func (m_ MTRCommandWithRequiredResponse) CommandFields() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("commandFields"))
	return rv
}


// SetCommandFields sets the value of the commandFields property.
// The command fields to pass for the command invoke. nil if this command does not have any fields. If not nil, this should be a data-value dictionary of MTRStructureValueType.

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRCommandWithRequiredResponse/commandFields
func (m_ MTRCommandWithRequiredResponse) SetCommandFields(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setCommandFields:"), value)
}
// The path of the command being invoked.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRCommandWithRequiredResponse/path
func (m_ MTRCommandWithRequiredResponse) Path() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("path"))
	return rv
}


// SetPath sets the value of the path property.
// The path of the command being invoked.

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRCommandWithRequiredResponse/path
func (m_ MTRCommandWithRequiredResponse) SetPath(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setPath:"), value)
}
// The response that represents this command succeeding.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRCommandWithRequiredResponse/requiredResponse
func (m_ MTRCommandWithRequiredResponse) RequiredResponse() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("requiredResponse"))
	return rv
}


// SetRequiredResponse sets the value of the requiredResponse property.
// The response that represents this command succeeding.

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRCommandWithRequiredResponse/requiredResponse
func (m_ MTRCommandWithRequiredResponse) SetRequiredResponse(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setRequiredResponse:"), value)
}

