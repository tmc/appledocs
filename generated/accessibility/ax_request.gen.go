// Code generated from Apple documentation for Accessibility. DO NOT EDIT.

package accessibility

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [AXRequest] class.
var (
	AXRequestClass     _AXRequestClass
	AXRequestClassOnce sync.Once
)

func getAXRequestClass() _AXRequestClass {
	AXRequestClassOnce.Do(func() {
		AXRequestClass = _AXRequestClass{objc.GetClass("AXRequest")}
	})
	return AXRequestClass
}

type _AXRequestClass struct {
	class objc.Class
}

// An interface definition for the [AXRequest] class.
type IAXRequest interface {
	objectivec.IObject
	Technology() AXTechnology
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accessibility/AccessibilityRequest
type AXRequest struct {
	objectivec.Object
}

// AXRequestFrom constructs a [AXRequest] from an unsafe.Pointer.
func AXRequestFrom(ptr unsafe.Pointer) AXRequest {
	return AXRequest{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (ac _AXRequestClass) Alloc() AXRequest {
	rv := objc.Send[AXRequest](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ac _AXRequestClass) New() AXRequest {
	rv := objc.Send[AXRequest](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ AXRequest) Init() AXRequest {
	rv := objc.Send[AXRequest](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ AXRequest) Autorelease() AXRequest {
	rv := objc.Send[AXRequest](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAXRequest creates a new AXRequest instance.
func NewAXRequest() AXRequest {
	return getAXRequestClass().New()
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accessibility/AccessibilityRequest/current
func (ac _AXRequestClass) CurrentRequest() AXRequest {
	rv := objc.Send[AXRequest](objc.ID(ac.class), objc.Sel("currentRequest"))
	return rv
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accessibility/AccessibilityRequest/current
func (a_ AXRequest) CurrentRequest() AXRequest {
	rv := objc.Send[AXRequest](a_.ID, objc.Sel("currentRequest"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accessibility/AccessibilityRequest/technology
func (a_ AXRequest) Technology() AXTechnology {
	rv := objc.Send[AXTechnology](a_.ID, objc.Sel("technology"))
	return rv
}




