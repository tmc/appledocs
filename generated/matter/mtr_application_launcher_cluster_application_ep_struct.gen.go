// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRApplicationLauncherClusterApplicationEPStruct] class.
var (
	MTRApplicationLauncherClusterApplicationEPStructClass     _MTRApplicationLauncherClusterApplicationEPStructClass
	MTRApplicationLauncherClusterApplicationEPStructClassOnce sync.Once
)

func getMTRApplicationLauncherClusterApplicationEPStructClass() _MTRApplicationLauncherClusterApplicationEPStructClass {
	MTRApplicationLauncherClusterApplicationEPStructClassOnce.Do(func() {
		MTRApplicationLauncherClusterApplicationEPStructClass = _MTRApplicationLauncherClusterApplicationEPStructClass{objc.GetClass("MTRApplicationLauncherClusterApplicationEPStruct")}
	})
	return MTRApplicationLauncherClusterApplicationEPStructClass
}

type _MTRApplicationLauncherClusterApplicationEPStructClass struct {
	class objc.Class
}

// An interface definition for the [MTRApplicationLauncherClusterApplicationEPStruct] class.
type IMTRApplicationLauncherClusterApplicationEPStruct interface {
	objectivec.IObject
	Application() MTRApplicationLauncherClusterApplicationStruct
	SetApplication(value IMTRApplicationLauncherClusterApplicationStruct)
	Endpoint() foundation.Number
	SetEndpoint(value foundation.INumber)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRApplicationLauncherClusterApplicationEPStruct
type MTRApplicationLauncherClusterApplicationEPStruct struct {
	objectivec.Object
}

// MTRApplicationLauncherClusterApplicationEPStructFrom constructs a [MTRApplicationLauncherClusterApplicationEPStruct] from an unsafe.Pointer.
func MTRApplicationLauncherClusterApplicationEPStructFrom(ptr unsafe.Pointer) MTRApplicationLauncherClusterApplicationEPStruct {
	return MTRApplicationLauncherClusterApplicationEPStruct{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRApplicationLauncherClusterApplicationEPStructClass) Alloc() MTRApplicationLauncherClusterApplicationEPStruct {
	rv := objc.Send[MTRApplicationLauncherClusterApplicationEPStruct](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRApplicationLauncherClusterApplicationEPStructClass) New() MTRApplicationLauncherClusterApplicationEPStruct {
	rv := objc.Send[MTRApplicationLauncherClusterApplicationEPStruct](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRApplicationLauncherClusterApplicationEPStruct) Init() MTRApplicationLauncherClusterApplicationEPStruct {
	rv := objc.Send[MTRApplicationLauncherClusterApplicationEPStruct](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRApplicationLauncherClusterApplicationEPStruct) Autorelease() MTRApplicationLauncherClusterApplicationEPStruct {
	rv := objc.Send[MTRApplicationLauncherClusterApplicationEPStruct](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRApplicationLauncherClusterApplicationEPStruct creates a new MTRApplicationLauncherClusterApplicationEPStruct instance.
func NewMTRApplicationLauncherClusterApplicationEPStruct() MTRApplicationLauncherClusterApplicationEPStruct {
	return getMTRApplicationLauncherClusterApplicationEPStructClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrapplicationlauncherclusterapplicationepstruct/application
func (m_ MTRApplicationLauncherClusterApplicationEPStruct) Application() MTRApplicationLauncherClusterApplicationStruct {
	rv := objc.Send[MTRApplicationLauncherClusterApplicationStruct](m_.ID, objc.Sel("application"))
	return rv
}


// SetApplication sets the value of the application property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrapplicationlauncherclusterapplicationepstruct/application
func (m_ MTRApplicationLauncherClusterApplicationEPStruct) SetApplication(value IMTRApplicationLauncherClusterApplicationStruct) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setApplication:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrapplicationlauncherclusterapplicationepstruct/endpoint
func (m_ MTRApplicationLauncherClusterApplicationEPStruct) Endpoint() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("endpoint"))
	return rv
}


// SetEndpoint sets the value of the endpoint property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrapplicationlauncherclusterapplicationepstruct/endpoint
func (m_ MTRApplicationLauncherClusterApplicationEPStruct) SetEndpoint(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setEndpoint:"), value)
}



