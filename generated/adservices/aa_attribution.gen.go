// Code generated from Apple documentation for AdServices. DO NOT EDIT.

package adservices

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [AAAttribution] class.
var (
	AAAttributionClass     _AAAttributionClass
	AAAttributionClassOnce sync.Once
)

func getAAAttributionClass() _AAAttributionClass {
	AAAttributionClassOnce.Do(func() {
		AAAttributionClass = _AAAttributionClass{objc.GetClass("AAAttribution")}
	})
	return AAAttributionClass
}

type _AAAttributionClass struct {
	class objc.Class
}

// An interface definition for the [AAAttribution] class.
type IAAAttribution interface {
	objectivec.IObject
}

// The parent class that the framework uses to request a token.
//
// [Full Topic]: https://developer.apple.com/documentation/AdServices/AAAttribution
type AAAttribution struct {
	objectivec.Object
}

// AAAttributionFrom constructs a [AAAttribution] from an unsafe.Pointer.
//
// The parent class that the framework uses to request a token.
func AAAttributionFrom(ptr unsafe.Pointer) AAAttribution {
	return AAAttribution{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (ac _AAAttributionClass) Alloc() AAAttribution {
	rv := objc.Send[AAAttribution](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ac _AAAttributionClass) New() AAAttribution {
	rv := objc.Send[AAAttribution](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ AAAttribution) Init() AAAttribution {
	rv := objc.Send[AAAttribution](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ AAAttribution) Autorelease() AAAttribution {
	rv := objc.Send[AAAttribution](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAAAttribution creates a new AAAttribution instance.
func NewAAAttribution() AAAttribution {
	return getAAAttributionClass().New()
}


// Generates a token.
//
// [Full Topic]: https://developer.apple.com/documentation/AdServices/AAAttribution/attributionToken()
func (ac _AAAttributionClass) AttributionTokenWithError(error_ unsafe.Pointer) string {
	rv := objc.Send[string](objc.ID(ac.class), objc.Sel("attributionTokenWithError:"), error_)
	return rv
}




