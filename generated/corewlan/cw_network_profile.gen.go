// Code generated from Apple documentation for CoreWLAN. DO NOT EDIT.

package corewlan

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [CWNetworkProfile] class.
var (
	CWNetworkProfileClass     _CWNetworkProfileClass
	CWNetworkProfileClassOnce sync.Once
)

func getCWNetworkProfileClass() _CWNetworkProfileClass {
	CWNetworkProfileClassOnce.Do(func() {
		CWNetworkProfileClass = _CWNetworkProfileClass{objc.GetClass("CWNetworkProfile")}
	})
	return CWNetworkProfileClass
}

type _CWNetworkProfileClass struct {
	class objc.Class
}

// An interface definition for the [CWNetworkProfile] class.
type ICWNetworkProfile interface {
	objectivec.IObject
	// properties:
	Security() CWSecurity
	Ssid() string /* primitive/slice/pointer. */
	SsidData() foundation.objc.IObject /* cross-framework: NSData */
	// methods:
	IsEqualToNetworkProfile(networkProfile ICWNetworkProfile) bool /* primitive/slice/pointer. */
}

// Encapsulates an immutable network profile entry.


// Encapsulates an immutable network profile entry.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreWLAN/CWNetworkProfile
type CWNetworkProfile struct {
	objectivec.Object
}

// CWNetworkProfileFrom constructs a [CWNetworkProfile] from an unsafe.Pointer.
//
// Encapsulates an immutable network profile entry.
func CWNetworkProfileFrom(ptr unsafe.Pointer) CWNetworkProfile {
	return CWNetworkProfile{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (cc _CWNetworkProfileClass) Alloc() CWNetworkProfile {
	rv := objc.Send[CWNetworkProfile](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (cc _CWNetworkProfileClass) New() CWNetworkProfile {
	rv := objc.Send[CWNetworkProfile](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CWNetworkProfile) Init() CWNetworkProfile {
	rv := objc.Send[CWNetworkProfile](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CWNetworkProfile) Autorelease() CWNetworkProfile {
	rv := objc.Send[CWNetworkProfile](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCWNetworkProfile creates a new CWNetworkProfile instance.
func NewCWNetworkProfile() CWNetworkProfile {
	return getCWNetworkProfileClass().New()
}



// Creates and returns a CWNetworkProfile object initialized with the given CWNetworkProfile object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreWLAN/CWNetworkProfile/init(networkProfile:)
func NewCWNetworkProfileWithNetworkProfile(networkProfile ICWNetworkProfile) CWNetworkProfile {
	instance := getCWNetworkProfileClass().Alloc()
	rv := objc.Send[CWNetworkProfile](instance.ID, objc.Sel("initWithNetworkProfile:"), networkProfile)
	rv.Autorelease()
	return rv
}



// Convenience method for getting a CWNetworkProfile object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreWLAN/CWNetworkProfile/networkProfile
func (cc _CWNetworkProfileClass) NetworkProfile() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(cc.class), objc.Sel("networkProfile"))
	return rv
}


// Convenience method for getting a CWNetworkProfile object initialized with the given CWNetworkProfile object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreWLAN/CWNetworkProfile/networkProfileWithNetworkProfile:
func (cc _CWNetworkProfileClass) NetworkProfileWithNetworkProfile(networkProfile ICWNetworkProfile) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(cc.class), objc.Sel("networkProfileWithNetworkProfile:"), networkProfile)
	return rv
}


// Determine CWNetworkProfile object equality.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreWLAN/CWNetworkProfile/isEqual(to:)
func (c_ CWNetworkProfile) IsEqualToNetworkProfile(networkProfile ICWNetworkProfile) bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](c_.ID, objc.Sel("isEqualToNetworkProfile:"), networkProfile)
	return rv
}


// The security mode for the network profile.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreWLAN/CWNetworkProfile/security
func (c_ CWNetworkProfile) Security() CWSecurity {
	rv := objc.Send[CWSecurity](c_.ID, objc.Sel("security"))
	return rv
}


// The service set identifier (SSID) for the network profile, encoded as a string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreWLAN/CWNetworkProfile/ssid
func (c_ CWNetworkProfile) Ssid() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](c_.ID, objc.Sel("ssid"))
	return rv
}


// The service set identifier (SSID) for the network profile, returned as data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreWLAN/CWNetworkProfile/ssidData
func (c_ CWNetworkProfile) SsidData() foundation.objc.IObject /* cross-framework: NSData */ {
	rv := objc.Send[foundation.NSData](c_.ID, objc.Sel("ssidData"))
	return rv
}


