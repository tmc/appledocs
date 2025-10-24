// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)





// The class instance for the [CaptionConversionValidator] class.
var (
	CaptionConversionValidatorClass     _CaptionConversionValidatorClass
	CaptionConversionValidatorClassOnce sync.Once
)

func getCaptionConversionValidatorClass() _CaptionConversionValidatorClass {
	CaptionConversionValidatorClassOnce.Do(func() {
		CaptionConversionValidatorClass = _CaptionConversionValidatorClass{objc.GetClass("AVCaptionConversionValidator")}
	})
	return CaptionConversionValidatorClass
}

type _CaptionConversionValidatorClass struct {
	class objc.Class
}





// An interface definition for the [CaptionConversionValidator] class.
type ICaptionConversionValidator interface {
	objectivec.IObject
	

	// properties:
	Captions() []Caption
	Status() CaptionConversionValidatorStatus
	TimeRange() TimeRange /* not a class type */
	Warnings() []CaptionConversionWarning


	

	// methods:
	StopValidating()
	ValidateCaptionConversionWithWarningHandler(handler unsafe.Pointer)


}





// Alloc allocates a new instance without initialization.
func (cc _CaptionConversionValidatorClass) Alloc() CaptionConversionValidator {
	rv := objc.Send[CaptionConversionValidator](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CaptionConversionValidatorClass) New() CaptionConversionValidator {
	rv := objc.Send[CaptionConversionValidator](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CaptionConversionValidator) Init() CaptionConversionValidator {
	rv := objc.Send[CaptionConversionValidator](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CaptionConversionValidator) Autorelease() CaptionConversionValidator {
	rv := objc.Send[CaptionConversionValidator](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCaptionConversionValidator creates a new CaptionConversionValidator instance.
func NewCaptionConversionValidator() CaptionConversionValidator {
	return getCaptionConversionValidatorClass().New()
}





// An object that validates captions for a conversion operation.


// An object that validates captions for a conversion operation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptionConversionValidator
type CaptionConversionValidator struct {
	objectivec.Object
}

// CaptionConversionValidatorFrom constructs a [CaptionConversionValidator] from an unsafe.Pointer.
//
// An object that validates captions for a conversion operation.
func CaptionConversionValidatorFrom(ptr unsafe.Pointer) CaptionConversionValidator {
	return CaptionConversionValidator{objectivec.Object{objc.ID(ptr)}}
}






// Creates an object that validates captions for a conversion operation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptionConversionValidator/init(captions:timeRange:conversionSettings:)
func NewCaptionConversionValidatorWithCaptionsTimeRangeConversionSettings(captions []Caption, timeRange TimeRange /* not a class type */, conversionSettings foundation.IDictionary) CaptionConversionValidator {
	instance := getCaptionConversionValidatorClass().Alloc()
	rv := objc.Send[CaptionConversionValidator](instance.ID, objc.Sel("initWithCaptions:timeRange:conversionSettings:"), captions, timeRange, conversionSettings)
	rv.Autorelease()
	return rv
}







// A convenience initializer to create an object that validates captions for a conversion operation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptionConversionValidator/captionConversionValidatorWithCaptions:timeRange:conversionSettings:
func (cc _CaptionConversionValidatorClass) CaptionConversionValidatorWithCaptionsTimeRangeConversionSettings(captions []Caption, timeRange TimeRange /* not a class type */, conversionSettings foundation.IDictionary) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(cc.class), objc.Sel("captionConversionValidatorWithCaptions:timeRange:conversionSettings:"), captions, timeRange, conversionSettings)
	return rv
}












// Stops the active validation operation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptionConversionValidator/stopValidating()
func (c_ CaptionConversionValidator) StopValidating() {
	objc.Send[objc.ID](c_.ID, objc.Sel("stopValidating"))
}


// Validates the object’s captions.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptionConversionValidator/validateCaptionConversion(warningHandler:)
func (c_ CaptionConversionValidator) ValidateCaptionConversionWithWarningHandler(handler unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("validateCaptionConversionWithWarningHandler:"), handler)
}







// The array of captions that the system validates.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptionConversionValidator/captions
func (c_ CaptionConversionValidator) Captions() []Caption {
	rv := objc.Send[[]Caption](c_.ID, objc.Sel("captions"))
	return rv
}


// A value that indicates the status of validation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptionConversionValidator/status-swift.property
func (c_ CaptionConversionValidator) Status() CaptionConversionValidatorStatus {
	rv := objc.Send[CaptionConversionValidatorStatus](c_.ID, objc.Sel("status"))
	return rv
}


// The time range of the media timeline in which the captions must exist.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptionConversionValidator/timeRange
func (c_ CaptionConversionValidator) TimeRange() TimeRange /* not a class type */ {
	rv := objc.Send[TimeRange](c_.ID, objc.Sel("timeRange"))
	return rv
}


// The collection of warnings the validator encountered.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptionConversionValidator/warnings
func (c_ CaptionConversionValidator) Warnings() []CaptionConversionWarning {
	rv := objc.Send[[]CaptionConversionWarning](c_.ID, objc.Sel("warnings"))
	return rv
}







