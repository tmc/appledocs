// Code generated from Apple documentation for CoreWLAN. DO NOT EDIT.

package corewlan

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [CWMutableNetworkProfile] class.
var (
	CWMutableNetworkProfileClass     _CWMutableNetworkProfileClass
	CWMutableNetworkProfileClassOnce sync.Once
)

func getCWMutableNetworkProfileClass() _CWMutableNetworkProfileClass {
	CWMutableNetworkProfileClassOnce.Do(func() {
		CWMutableNetworkProfileClass = _CWMutableNetworkProfileClass{objc.GetClass("CWMutableNetworkProfile")}
	})
	return CWMutableNetworkProfileClass
}

type _CWMutableNetworkProfileClass struct {
	class objc.Class
}

// An interface definition for the [CWMutableNetworkProfile] class.
type ICWMutableNetworkProfile interface {
	ICWNetworkProfile
}

// Encapsulates a mutable network profile entry.
//
// Use this class to change profile properties. To commit Wi-Fi network profile changes, use and .
//
// [Full Topic]: https://developer.apple.com/documentation/CoreWLAN/CWMutableNetworkProfile
type CWMutableNetworkProfile struct {
	CWNetworkProfile
}

// CWMutableNetworkProfileFrom constructs a [CWMutableNetworkProfile] from an unsafe.Pointer.
//
// Encapsulates a mutable network profile entry.
func CWMutableNetworkProfileFrom(ptr unsafe.Pointer) CWMutableNetworkProfile {
	return CWMutableNetworkProfile{
		CWNetworkProfile: CWNetworkProfileFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (cc _CWMutableNetworkProfileClass) Alloc() CWMutableNetworkProfile {
	rv := objc.Send[CWMutableNetworkProfile](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (cc _CWMutableNetworkProfileClass) New() CWMutableNetworkProfile {
	rv := objc.Send[CWMutableNetworkProfile](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CWMutableNetworkProfile) Init() CWMutableNetworkProfile {
	rv := objc.Send[CWMutableNetworkProfile](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CWMutableNetworkProfile) Autorelease() CWMutableNetworkProfile {
	rv := objc.Send[CWMutableNetworkProfile](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCWMutableNetworkProfile creates a new CWMutableNetworkProfile instance.
func NewCWMutableNetworkProfile() CWMutableNetworkProfile {
	return getCWMutableNetworkProfileClass().New()
}


// The security type.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreWLAN/CWMutableNetworkProfile/security
func (c_ CWMutableNetworkProfile) Security() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("security"))
	return rv
}


// SetSecurity sets the value of the security property.
// The security type.

//
// [Full Topic]: https://developer.apple.com/documentation/CoreWLAN/CWMutableNetworkProfile/security
func (c_ CWMutableNetworkProfile) SetSecurity(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setSecurity:"), value)
}
// The service set identifier (SSID).
//
// [Full Topic]: https://developer.apple.com/documentation/CoreWLAN/CWMutableNetworkProfile/ssidData
func (c_ CWMutableNetworkProfile) SsidData() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("ssidData"))
	return rv
}


// SetSsidData sets the value of the ssidData property.
// The service set identifier (SSID).

//
// [Full Topic]: https://developer.apple.com/documentation/CoreWLAN/CWMutableNetworkProfile/ssidData
func (c_ CWMutableNetworkProfile) SetSsidData(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setSsidData:"), value)
}


