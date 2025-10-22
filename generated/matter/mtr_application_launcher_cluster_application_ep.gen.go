// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [MTRApplicationLauncherClusterApplicationEP] class.
var (
	MTRApplicationLauncherClusterApplicationEPClass     _MTRApplicationLauncherClusterApplicationEPClass
	MTRApplicationLauncherClusterApplicationEPClassOnce sync.Once
)

func getMTRApplicationLauncherClusterApplicationEPClass() _MTRApplicationLauncherClusterApplicationEPClass {
	MTRApplicationLauncherClusterApplicationEPClassOnce.Do(func() {
		MTRApplicationLauncherClusterApplicationEPClass = _MTRApplicationLauncherClusterApplicationEPClass{objc.GetClass("MTRApplicationLauncherClusterApplicationEP")}
	})
	return MTRApplicationLauncherClusterApplicationEPClass
}

type _MTRApplicationLauncherClusterApplicationEPClass struct {
	class objc.Class
}

// An interface definition for the [MTRApplicationLauncherClusterApplicationEP] class.
type IMTRApplicationLauncherClusterApplicationEP interface {
	IMTRApplicationLauncherClusterApplicationEPStruct
	Application() MTRApplicationLauncherClusterApplicationStruct
	SetApplication(value IMTRApplicationLauncherClusterApplicationStruct)
	Endpoint() foundation.Number
	SetEndpoint(value foundation.INumber)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRApplicationLauncherClusterApplicationEP
type MTRApplicationLauncherClusterApplicationEP struct {
	MTRApplicationLauncherClusterApplicationEPStruct
}

// MTRApplicationLauncherClusterApplicationEPFrom constructs a [MTRApplicationLauncherClusterApplicationEP] from an unsafe.Pointer.
func MTRApplicationLauncherClusterApplicationEPFrom(ptr unsafe.Pointer) MTRApplicationLauncherClusterApplicationEP {
	return MTRApplicationLauncherClusterApplicationEP{
		MTRApplicationLauncherClusterApplicationEPStruct: MTRApplicationLauncherClusterApplicationEPStructFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRApplicationLauncherClusterApplicationEPClass) Alloc() MTRApplicationLauncherClusterApplicationEP {
	rv := objc.Send[MTRApplicationLauncherClusterApplicationEP](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRApplicationLauncherClusterApplicationEPClass) New() MTRApplicationLauncherClusterApplicationEP {
	rv := objc.Send[MTRApplicationLauncherClusterApplicationEP](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRApplicationLauncherClusterApplicationEP) Init() MTRApplicationLauncherClusterApplicationEP {
	rv := objc.Send[MTRApplicationLauncherClusterApplicationEP](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRApplicationLauncherClusterApplicationEP) Autorelease() MTRApplicationLauncherClusterApplicationEP {
	rv := objc.Send[MTRApplicationLauncherClusterApplicationEP](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRApplicationLauncherClusterApplicationEP creates a new MTRApplicationLauncherClusterApplicationEP instance.
func NewMTRApplicationLauncherClusterApplicationEP() MTRApplicationLauncherClusterApplicationEP {
	return getMTRApplicationLauncherClusterApplicationEPClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrapplicationlauncherclusterapplicationep/application
func (m_ MTRApplicationLauncherClusterApplicationEP) Application() MTRApplicationLauncherClusterApplicationStruct {
	rv := objc.Send[MTRApplicationLauncherClusterApplicationStruct](m_.ID, objc.Sel("application"))
	return rv
}


// SetApplication sets the value of the application property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrapplicationlauncherclusterapplicationep/application
func (m_ MTRApplicationLauncherClusterApplicationEP) SetApplication(value IMTRApplicationLauncherClusterApplicationStruct) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setApplication:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrapplicationlauncherclusterapplicationep/endpoint
func (m_ MTRApplicationLauncherClusterApplicationEP) Endpoint() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("endpoint"))
	return rv
}


// SetEndpoint sets the value of the endpoint property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrapplicationlauncherclusterapplicationep/endpoint
func (m_ MTRApplicationLauncherClusterApplicationEP) SetEndpoint(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setEndpoint:"), value)
}



