// Code generated from Apple documentation for Intents. DO NOT EDIT.

package intents

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/corelocation"
	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [INLocationRelevanceProvider] class.
var (
	INLocationRelevanceProviderClass     _INLocationRelevanceProviderClass
	INLocationRelevanceProviderClassOnce sync.Once
)

func getINLocationRelevanceProviderClass() _INLocationRelevanceProviderClass {
	INLocationRelevanceProviderClassOnce.Do(func() {
		INLocationRelevanceProviderClass = _INLocationRelevanceProviderClass{objc.GetClass("INLocationRelevanceProvider")}
	})
	return INLocationRelevanceProviderClass
}

type _INLocationRelevanceProviderClass struct {
	class objc.Class
}

// An interface definition for the [INLocationRelevanceProvider] class.
type IINLocationRelevanceProvider interface {
	IINRelevanceProvider
	// properties:
	Region() objc.IObject /* cross-framework: Region */
	SetRegion(value objc.IObject /* cross-framework: Region */)
	// methods:
}

// The provider class that specifies a relevant location.
//
// Ask the user for permission to use their location before providing shortcuts to that include a location relevance provider. If the user gives your app permission to access their location While In Use or Always_,_ shortcuts your app provides can influence widget stacks and the Siri watch face.

// The provider class that specifies a relevant location.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Intents/INLocationRelevanceProvider
type INLocationRelevanceProvider struct {
	INRelevanceProvider
}

// INLocationRelevanceProviderFrom constructs a [INLocationRelevanceProvider] from an unsafe.Pointer.
//
// The provider class that specifies a relevant location.
func INLocationRelevanceProviderFrom(ptr unsafe.Pointer) INLocationRelevanceProvider {
	return INLocationRelevanceProvider{
		INRelevanceProvider: INRelevanceProviderFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (ic _INLocationRelevanceProviderClass) Alloc() INLocationRelevanceProvider {
	rv := objc.Send[INLocationRelevanceProvider](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ic _INLocationRelevanceProviderClass) New() INLocationRelevanceProvider {
	rv := objc.Send[INLocationRelevanceProvider](objc.ID(ic.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (i_ INLocationRelevanceProvider) Init() INLocationRelevanceProvider {
	rv := objc.Send[INLocationRelevanceProvider](i_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i_ INLocationRelevanceProvider) Autorelease() INLocationRelevanceProvider {
	rv := objc.Send[INLocationRelevanceProvider](i_.ID, objc.Sel("autorelease"))
	return rv
}

// NewINLocationRelevanceProvider creates a new INLocationRelevanceProvider instance.
func NewINLocationRelevanceProvider() INLocationRelevanceProvider {
	return getINLocationRelevanceProviderClass().New()
}

// The region representing the relevant location.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/intents/inlocationrelevanceprovider/region
func (i_ INLocationRelevanceProvider) Region() objc.IObject /* cross-framework: Region */ {
	rv := objc.Send[corelocation.Region](i_.ID, objc.Sel("region"))
	return rv
}

// The region representing the relevant location.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/intents/inlocationrelevanceprovider/region
func (i_ INLocationRelevanceProvider) SetRegion(value objc.IObject /* cross-framework: Region */) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setRegion:"), value)
}
