// Code generated from Apple documentation for ThreadNetwork. DO NOT EDIT.

package threadnetwork

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [THCredentials] class.
var (
	THCredentialsClass     _THCredentialsClass
	THCredentialsClassOnce sync.Once
)

func getTHCredentialsClass() _THCredentialsClass {
	THCredentialsClassOnce.Do(func() {
		THCredentialsClass = _THCredentialsClass{objc.GetClass("THCredentials")}
	})
	return THCredentialsClass
}

type _THCredentialsClass struct {
	class objc.Class
}

// An interface definition for the [THCredentials] class.
type ITHCredentials interface {
	objectivec.IObject
}

// A class that contains credentials for a Thread network.
//
// A Thread network defines parameters that all connected devices use. provides these parameters.
//
// [Full Topic]: https://developer.apple.com/documentation/ThreadNetwork/THCredentials
type THCredentials struct {
	objectivec.Object
}

// THCredentialsFrom constructs a [THCredentials] from an unsafe.Pointer.
//
// A class that contains credentials for a Thread network.
func THCredentialsFrom(ptr unsafe.Pointer) THCredentials {
	return THCredentials{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (tc _THCredentialsClass) Alloc() THCredentials {
	rv := objc.Send[THCredentials](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (tc _THCredentialsClass) New() THCredentials {
	rv := objc.Send[THCredentials](objc.ID(tc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (t_ THCredentials) Init() THCredentials {
	rv := objc.Send[THCredentials](t_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t_ THCredentials) Autorelease() THCredentials {
	rv := objc.Send[THCredentials](t_.ID, objc.Sel("autorelease"))
	return rv
}

// NewTHCredentials creates a new THCredentials instance.
func NewTHCredentials() THCredentials {
	return getTHCredentialsClass().New()
}


// The essential operational parameters for the Thread network.
//
// [Full Topic]: https://developer.apple.com/documentation/ThreadNetwork/THCredentials/activeOperationalDataSet
func (t_ THCredentials) ActiveOperationalDataSet() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t_.ID, objc.Sel("activeOperationalDataSet"))
	return rv
}

// The identifier of an active Thread network Border Agent.
//
// [Full Topic]: https://developer.apple.com/documentation/ThreadNetwork/THCredentials/borderAgentID
func (t_ THCredentials) BorderAgentID() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t_.ID, objc.Sel("borderAgentID"))
	return rv
}

// The Thread network radio channel.
//
// [Full Topic]: https://developer.apple.com/documentation/ThreadNetwork/THCredentials/channel
func (t_ THCredentials) Channel() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t_.ID, objc.Sel("channel"))
	return rv
}


// SetChannel sets the value of the channel property.
// The Thread network radio channel.

//
// [Full Topic]: https://developer.apple.com/documentation/ThreadNetwork/THCredentials/channel
func (t_ THCredentials) SetChannel(value unsafe.Pointer) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setChannel:"), value)
}

// The date and time that the framework stored the credential in the database.
//
// [Full Topic]: https://developer.apple.com/documentation/ThreadNetwork/THCredentials/creationDate
func (t_ THCredentials) CreationDate() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t_.ID, objc.Sel("creationDate"))
	return rv
}

// The Thread network extended PAN identifier.
//
// [Full Topic]: https://developer.apple.com/documentation/ThreadNetwork/THCredentials/extendedPANID
func (t_ THCredentials) ExtendedPANID() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t_.ID, objc.Sel("extendedPANID"))
	return rv
}

// The date and time that the framework updated the credential in the database.
//
// [Full Topic]: https://developer.apple.com/documentation/ThreadNetwork/THCredentials/lastModificationDate
func (t_ THCredentials) LastModificationDate() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t_.ID, objc.Sel("lastModificationDate"))
	return rv
}

// The Thread network key.
//
// [Full Topic]: https://developer.apple.com/documentation/ThreadNetwork/THCredentials/networkKey
func (t_ THCredentials) NetworkKey() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t_.ID, objc.Sel("networkKey"))
	return rv
}

// The Thread network name.
//
// [Full Topic]: https://developer.apple.com/documentation/ThreadNetwork/THCredentials/networkName
func (t_ THCredentials) NetworkName() string {
	rv := objc.Send[string](t_.ID, objc.Sel("networkName"))
	return rv
}

// The Thread network PAN identifier.
//
// [Full Topic]: https://developer.apple.com/documentation/ThreadNetwork/THCredentials/panID
func (t_ THCredentials) PanID() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t_.ID, objc.Sel("panID"))
	return rv
}

// The Thread network pre-shared key (PSKC) for the Commissioner.
//
// [Full Topic]: https://developer.apple.com/documentation/ThreadNetwork/THCredentials/pskc
func (t_ THCredentials) PSKC() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t_.ID, objc.Sel("PSKC"))
	return rv
}




