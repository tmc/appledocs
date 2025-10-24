// Code generated from Apple documentation for OpenDirectory. DO NOT EDIT.

package opendirectory

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [connectionSetupTimeoutInSeconds] class.
var (
	ConnectionSetupTimeoutInSecondsClass     _connectionSetupTimeoutInSecondsClass
	ConnectionSetupTimeoutInSecondsClassOnce sync.Once
)

func getconnectionSetupTimeoutInSecondsClass() _connectionSetupTimeoutInSecondsClass {
	ConnectionSetupTimeoutInSecondsClassOnce.Do(func() {
		ConnectionSetupTimeoutInSecondsClass = _connectionSetupTimeoutInSecondsClass{objc.GetClass("connectionSetupTimeoutInSeconds")}
	})
	return ConnectionSetupTimeoutInSecondsClass
}

type _connectionSetupTimeoutInSecondsClass struct {
	class objc.Class
}

// An interface definition for the [connectionSetupTimeoutInSeconds] class.
type IconnectionSetupTimeoutInSeconds interface {
	objectivec.IObject
	// properties:
	// methods:
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODConfiguration/connectionSetupTimeoutInSeconds-c.ivar
type connectionSetupTimeoutInSeconds struct {
	objectivec.Object
}

// connectionSetupTimeoutInSecondsFrom constructs a [connectionSetupTimeoutInSeconds] from an unsafe.Pointer.
func connectionSetupTimeoutInSecondsFrom(ptr unsafe.Pointer) connectionSetupTimeoutInSeconds {
	return connectionSetupTimeoutInSeconds{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (cc _connectionSetupTimeoutInSecondsClass) Alloc() connectionSetupTimeoutInSeconds {
	rv := objc.Send[connectionSetupTimeoutInSeconds](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (cc _connectionSetupTimeoutInSecondsClass) New() connectionSetupTimeoutInSeconds {
	rv := objc.Send[connectionSetupTimeoutInSeconds](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ connectionSetupTimeoutInSeconds) Init() connectionSetupTimeoutInSeconds {
	rv := objc.Send[connectionSetupTimeoutInSeconds](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ connectionSetupTimeoutInSeconds) Autorelease() connectionSetupTimeoutInSeconds {
	rv := objc.Send[connectionSetupTimeoutInSeconds](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewconnectionSetupTimeoutInSeconds creates a new connectionSetupTimeoutInSeconds instance.
func NewconnectionSetupTimeoutInSeconds() connectionSetupTimeoutInSeconds {
	return getconnectionSetupTimeoutInSecondsClass().New()
}




