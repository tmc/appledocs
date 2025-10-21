// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [MTRAudioOutputClusterOutputInfo] class.
var (
	MTRAudioOutputClusterOutputInfoClass     _MTRAudioOutputClusterOutputInfoClass
	MTRAudioOutputClusterOutputInfoClassOnce sync.Once
)

func getMTRAudioOutputClusterOutputInfoClass() _MTRAudioOutputClusterOutputInfoClass {
	MTRAudioOutputClusterOutputInfoClassOnce.Do(func() {
		MTRAudioOutputClusterOutputInfoClass = _MTRAudioOutputClusterOutputInfoClass{objc.GetClass("MTRAudioOutputClusterOutputInfo")}
	})
	return MTRAudioOutputClusterOutputInfoClass
}

type _MTRAudioOutputClusterOutputInfoClass struct {
	class objc.Class
}

// An interface definition for the [MTRAudioOutputClusterOutputInfo] class.
type IMTRAudioOutputClusterOutputInfo interface {
	IMTRAudioOutputClusterOutputInfoStruct
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRAudioOutputClusterOutputInfo
type MTRAudioOutputClusterOutputInfo struct {
	MTRAudioOutputClusterOutputInfoStruct
}

// MTRAudioOutputClusterOutputInfoFrom constructs a [MTRAudioOutputClusterOutputInfo] from an unsafe.Pointer.
func MTRAudioOutputClusterOutputInfoFrom(ptr unsafe.Pointer) MTRAudioOutputClusterOutputInfo {
	return MTRAudioOutputClusterOutputInfo{
		MTRAudioOutputClusterOutputInfoStruct: MTRAudioOutputClusterOutputInfoStructFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRAudioOutputClusterOutputInfoClass) Alloc() MTRAudioOutputClusterOutputInfo {
	rv := objc.Send[MTRAudioOutputClusterOutputInfo](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRAudioOutputClusterOutputInfoClass) New() MTRAudioOutputClusterOutputInfo {
	rv := objc.Send[MTRAudioOutputClusterOutputInfo](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRAudioOutputClusterOutputInfo) Init() MTRAudioOutputClusterOutputInfo {
	rv := objc.Send[MTRAudioOutputClusterOutputInfo](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRAudioOutputClusterOutputInfo) Autorelease() MTRAudioOutputClusterOutputInfo {
	rv := objc.Send[MTRAudioOutputClusterOutputInfo](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRAudioOutputClusterOutputInfo creates a new MTRAudioOutputClusterOutputInfo instance.
func NewMTRAudioOutputClusterOutputInfo() MTRAudioOutputClusterOutputInfo {
	return getMTRAudioOutputClusterOutputInfoClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtraudiooutputclusteroutputinfo/index
func (m_ MTRAudioOutputClusterOutputInfo) Index() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("index"))
	return rv
}


// SetIndex sets the value of the index property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtraudiooutputclusteroutputinfo/index
func (m_ MTRAudioOutputClusterOutputInfo) SetIndex(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setIndex:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtraudiooutputclusteroutputinfo/name
func (m_ MTRAudioOutputClusterOutputInfo) Name() appkit.string {
	rv := objc.Send[appkit.string](m_.ID, objc.Sel("name"))
	return rv
}


// SetName sets the value of the name property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtraudiooutputclusteroutputinfo/name
func (m_ MTRAudioOutputClusterOutputInfo) SetName(value appkit.string) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setName:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtraudiooutputclusteroutputinfo/outputtype
func (m_ MTRAudioOutputClusterOutputInfo) OutputType() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("outputType"))
	return rv
}


// SetOutputType sets the value of the outputType property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtraudiooutputclusteroutputinfo/outputtype
func (m_ MTRAudioOutputClusterOutputInfo) SetOutputType(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setOutputType:"), value)
}



