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
	

	// properties:
	ActiveOperationalDataSet() foundation.foundation.INSData
	BorderAgentID() foundation.foundation.INSData
	Channel() uint8 /* not a class type */
	SetChannel(value uint8 /* not a class type */)
	CreationDate() foundation.foundation.INSDate
	ExtendedPANID() foundation.foundation.INSData
	LastModificationDate() foundation.foundation.INSDate
	NetworkKey() foundation.foundation.INSData
	NetworkName() foundation.foundation.INSString
	PanID() foundation.foundation.INSData
	PSKC() foundation.foundation.INSData


	

	// methods:


}





// Alloc allocates a new instance without initialization.
func (tc _THCredentialsClass) Alloc() THCredentials {
	rv := objc.Send[THCredentials](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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





// A class that contains credentials for a Thread network.
//
// A Thread network defines parameters that all connected devices use. provides these parameters.


// A class that contains credentials for a Thread network.
//
// [Full Topic]
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

























// The essential operational parameters for the Thread network.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ThreadNetwork/THCredentials/activeOperationalDataSet
func (t_ THCredentials) ActiveOperationalDataSet() foundation.foundation.INSData {
	rv := objc.Send[foundation.NSData](t_.ID, objc.Sel("activeOperationalDataSet"))
	return rv
}


// The identifier of an active Thread network Border Agent.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ThreadNetwork/THCredentials/borderAgentID
func (t_ THCredentials) BorderAgentID() foundation.foundation.INSData {
	rv := objc.Send[foundation.NSData](t_.ID, objc.Sel("borderAgentID"))
	return rv
}


// The Thread network radio channel.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ThreadNetwork/THCredentials/channel
func (t_ THCredentials) Channel() uint8 /* not a class type */ {
	rv := objc.Send[uint8](t_.ID, objc.Sel("channel"))
	return rv
}


// The Thread network radio channel.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ThreadNetwork/THCredentials/channel
func (t_ THCredentials) SetChannel(value uint8 /* not a class type */) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setChannel:"), value)
}


// The date and time that the framework stored the credential in the database.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ThreadNetwork/THCredentials/creationDate
func (t_ THCredentials) CreationDate() foundation.foundation.INSDate {
	rv := objc.Send[foundation.NSDate](t_.ID, objc.Sel("creationDate"))
	return rv
}


// The Thread network extended PAN identifier.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ThreadNetwork/THCredentials/extendedPANID
func (t_ THCredentials) ExtendedPANID() foundation.foundation.INSData {
	rv := objc.Send[foundation.NSData](t_.ID, objc.Sel("extendedPANID"))
	return rv
}


// The date and time that the framework updated the credential in the database.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ThreadNetwork/THCredentials/lastModificationDate
func (t_ THCredentials) LastModificationDate() foundation.foundation.INSDate {
	rv := objc.Send[foundation.NSDate](t_.ID, objc.Sel("lastModificationDate"))
	return rv
}


// The Thread network key.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ThreadNetwork/THCredentials/networkKey
func (t_ THCredentials) NetworkKey() foundation.foundation.INSData {
	rv := objc.Send[foundation.NSData](t_.ID, objc.Sel("networkKey"))
	return rv
}


// The Thread network name.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ThreadNetwork/THCredentials/networkName
func (t_ THCredentials) NetworkName() foundation.foundation.INSString {
	rv := objc.Send[foundation.NSString](t_.ID, objc.Sel("networkName"))
	return rv
}


// The Thread network PAN identifier.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ThreadNetwork/THCredentials/panID
func (t_ THCredentials) PanID() foundation.foundation.INSData {
	rv := objc.Send[foundation.NSData](t_.ID, objc.Sel("panID"))
	return rv
}


// The Thread network pre-shared key (PSKC) for the Commissioner.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ThreadNetwork/THCredentials/pskc
func (t_ THCredentials) PSKC() foundation.foundation.INSData {
	rv := objc.Send[foundation.NSData](t_.ID, objc.Sel("PSKC"))
	return rv
}











