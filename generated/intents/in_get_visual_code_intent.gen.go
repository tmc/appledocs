// Code generated from Apple documentation for Intents. DO NOT EDIT.

package intents

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [INGetVisualCodeIntent] class.
var (
	INGetVisualCodeIntentClass     _INGetVisualCodeIntentClass
	INGetVisualCodeIntentClassOnce sync.Once
)

func getINGetVisualCodeIntentClass() _INGetVisualCodeIntentClass {
	INGetVisualCodeIntentClassOnce.Do(func() {
		INGetVisualCodeIntentClass = _INGetVisualCodeIntentClass{objc.GetClass("INGetVisualCodeIntent")}
	})
	return INGetVisualCodeIntentClass
}

type _INGetVisualCodeIntentClass struct {
	class objc.Class
}

// An interface definition for the [INGetVisualCodeIntent] class.
type IINGetVisualCodeIntent interface {
	IINIntent
	VisualCodeType() unsafe.Pointer
}

// A request for a visual code to use for exchanging payment and contact information.
//
// Siri creates an object when the user asks for a visual code to use with a scanner. A visual code is a bar code or QR code that embeds information about a transaction. For example, an app might provide a visual code to facilitate payment for services. Apps can also use visual codes to communicate a user’s contact information. To handle this intent, the handler object in your Intents extension must adopt the protocol. Your handler should confirm the request and create an object with an image containing the visual code. Siri handles the display of the provided visual code in order to scan it.
//
// [Full Topic]: https://developer.apple.com/documentation/Intents/INGetVisualCodeIntent
type INGetVisualCodeIntent struct {
	INIntent
}

// INGetVisualCodeIntentFrom constructs a [INGetVisualCodeIntent] from an unsafe.Pointer.
//
// A request for a visual code to use for exchanging payment and contact information.
func INGetVisualCodeIntentFrom(ptr unsafe.Pointer) INGetVisualCodeIntent {
	return INGetVisualCodeIntent{
		INIntent: INIntentFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (ic _INGetVisualCodeIntentClass) Alloc() INGetVisualCodeIntent {
	rv := objc.Send[INGetVisualCodeIntent](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ic _INGetVisualCodeIntentClass) New() INGetVisualCodeIntent {
	rv := objc.Send[INGetVisualCodeIntent](objc.ID(ic.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (i_ INGetVisualCodeIntent) Init() INGetVisualCodeIntent {
	rv := objc.Send[INGetVisualCodeIntent](i_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i_ INGetVisualCodeIntent) Autorelease() INGetVisualCodeIntent {
	rv := objc.Send[INGetVisualCodeIntent](i_.ID, objc.Sel("autorelease"))
	return rv
}

// NewINGetVisualCodeIntent creates a new INGetVisualCodeIntent instance.
func NewINGetVisualCodeIntent() INGetVisualCodeIntent {
	return getINGetVisualCodeIntentClass().New()
}


// The type of visual code requested by the user.
//
// [Full Topic]: https://developer.apple.com/documentation/Intents/INGetVisualCodeIntent/visualCodeType
func (i_ INGetVisualCodeIntent) VisualCodeType() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("visualCodeType"))
	return rv
}



