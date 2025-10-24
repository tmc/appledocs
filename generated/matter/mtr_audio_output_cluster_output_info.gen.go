// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objc"
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
	// properties:
	Index() objc.IObject /* cross-framework: NSNumber */
	SetIndex(value objc.IObject /* cross-framework: NSNumber */)
	Name() objc.IObject /* cross-framework: NSString */
	SetName(value objc.IObject /* cross-framework: NSString */)
	OutputType() objc.IObject /* cross-framework: NSNumber */
	SetOutputType(value objc.IObject /* cross-framework: NSNumber */)
	// methods:
}

// [Full Topic]
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

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtraudiooutputclusteroutputinfo/index
func (m_ MTRAudioOutputClusterOutputInfo) Index() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("index"))
	return rv
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtraudiooutputclusteroutputinfo/index
func (m_ MTRAudioOutputClusterOutputInfo) SetIndex(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setIndex:"), value)
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtraudiooutputclusteroutputinfo/name
func (m_ MTRAudioOutputClusterOutputInfo) Name() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("name"))
	return rv
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtraudiooutputclusteroutputinfo/name
func (m_ MTRAudioOutputClusterOutputInfo) SetName(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setName:"), value)
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtraudiooutputclusteroutputinfo/outputtype
func (m_ MTRAudioOutputClusterOutputInfo) OutputType() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("outputType"))
	return rv
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtraudiooutputclusteroutputinfo/outputtype
func (m_ MTRAudioOutputClusterOutputInfo) SetOutputType(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setOutputType:"), value)
}
