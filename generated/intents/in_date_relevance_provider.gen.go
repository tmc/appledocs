// Code generated from Apple documentation for Intents. DO NOT EDIT.

package intents

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [INDateRelevanceProvider] class.
var (
	INDateRelevanceProviderClass     _INDateRelevanceProviderClass
	INDateRelevanceProviderClassOnce sync.Once
)

func getINDateRelevanceProviderClass() _INDateRelevanceProviderClass {
	INDateRelevanceProviderClassOnce.Do(func() {
		INDateRelevanceProviderClass = _INDateRelevanceProviderClass{objc.GetClass("INDateRelevanceProvider")}
	})
	return INDateRelevanceProviderClass
}

type _INDateRelevanceProviderClass struct {
	class objc.Class
}

// An interface definition for the [INDateRelevanceProvider] class.
type IINDateRelevanceProvider interface {
	objectivec.IObject
}

// The provider class that specifies a relevant day and time.
//
// [Full Topic]: https://developer.apple.com/documentation/Intents/INDateRelevanceProvider
type INDateRelevanceProvider struct {
	objectivec.Object
}

// INDateRelevanceProviderFrom constructs a [INDateRelevanceProvider] from an unsafe.Pointer.
//
// The provider class that specifies a relevant day and time.
func INDateRelevanceProviderFrom(ptr unsafe.Pointer) INDateRelevanceProvider {
	return INDateRelevanceProvider{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (ic _INDateRelevanceProviderClass) Alloc() INDateRelevanceProvider {
	rv := objc.Send[INDateRelevanceProvider](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ic _INDateRelevanceProviderClass) New() INDateRelevanceProvider {
	rv := objc.Send[INDateRelevanceProvider](objc.ID(ic.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (i_ INDateRelevanceProvider) Init() INDateRelevanceProvider {
	rv := objc.Send[INDateRelevanceProvider](i_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i_ INDateRelevanceProvider) Autorelease() INDateRelevanceProvider {
	rv := objc.Send[INDateRelevanceProvider](i_.ID, objc.Sel("autorelease"))
	return rv
}

// NewINDateRelevanceProvider creates a new INDateRelevanceProvider instance.
func NewINDateRelevanceProvider() INDateRelevanceProvider {
	return getINDateRelevanceProviderClass().New()
}


// The relevant end date.
//
// [Full Topic]: https://developer.apple.com/documentation/intents/indaterelevanceprovider/enddate
func (i_ INDateRelevanceProvider) EndDate() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("endDate"))
	return rv
}


// SetEndDate sets the value of the endDate property.
// The relevant end date.

//
// [Full Topic]: https://developer.apple.com/documentation/intents/indaterelevanceprovider/enddate
func (i_ INDateRelevanceProvider) SetEndDate(value unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setEndDate:"), value)
}

// The relevant start date.
//
// [Full Topic]: https://developer.apple.com/documentation/Intents/INDateRelevanceProvider/startDate
func (i_ INDateRelevanceProvider) StartDate() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("startDate"))
	return rv
}



