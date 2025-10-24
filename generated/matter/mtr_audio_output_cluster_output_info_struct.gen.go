// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRAudioOutputClusterOutputInfoStruct] class.
var (
	MTRAudioOutputClusterOutputInfoStructClass     _MTRAudioOutputClusterOutputInfoStructClass
	MTRAudioOutputClusterOutputInfoStructClassOnce sync.Once
)

func getMTRAudioOutputClusterOutputInfoStructClass() _MTRAudioOutputClusterOutputInfoStructClass {
	MTRAudioOutputClusterOutputInfoStructClassOnce.Do(func() {
		MTRAudioOutputClusterOutputInfoStructClass = _MTRAudioOutputClusterOutputInfoStructClass{objc.GetClass("MTRAudioOutputClusterOutputInfoStruct")}
	})
	return MTRAudioOutputClusterOutputInfoStructClass
}

type _MTRAudioOutputClusterOutputInfoStructClass struct {
	class objc.Class
}

// An interface definition for the [MTRAudioOutputClusterOutputInfoStruct] class.
type IMTRAudioOutputClusterOutputInfoStruct interface {
	objectivec.IObject
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
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRAudioOutputClusterOutputInfoStruct
type MTRAudioOutputClusterOutputInfoStruct struct {
	objectivec.Object
}

// MTRAudioOutputClusterOutputInfoStructFrom constructs a [MTRAudioOutputClusterOutputInfoStruct] from an unsafe.Pointer.
func MTRAudioOutputClusterOutputInfoStructFrom(ptr unsafe.Pointer) MTRAudioOutputClusterOutputInfoStruct {
	return MTRAudioOutputClusterOutputInfoStruct{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRAudioOutputClusterOutputInfoStructClass) Alloc() MTRAudioOutputClusterOutputInfoStruct {
	rv := objc.Send[MTRAudioOutputClusterOutputInfoStruct](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRAudioOutputClusterOutputInfoStructClass) New() MTRAudioOutputClusterOutputInfoStruct {
	rv := objc.Send[MTRAudioOutputClusterOutputInfoStruct](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRAudioOutputClusterOutputInfoStruct) Init() MTRAudioOutputClusterOutputInfoStruct {
	rv := objc.Send[MTRAudioOutputClusterOutputInfoStruct](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRAudioOutputClusterOutputInfoStruct) Autorelease() MTRAudioOutputClusterOutputInfoStruct {
	rv := objc.Send[MTRAudioOutputClusterOutputInfoStruct](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRAudioOutputClusterOutputInfoStruct creates a new MTRAudioOutputClusterOutputInfoStruct instance.
func NewMTRAudioOutputClusterOutputInfoStruct() MTRAudioOutputClusterOutputInfoStruct {
	return getMTRAudioOutputClusterOutputInfoStructClass().New()
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtraudiooutputclusteroutputinfostruct/index
func (m_ MTRAudioOutputClusterOutputInfoStruct) Index() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("index"))
	return rv
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtraudiooutputclusteroutputinfostruct/index
func (m_ MTRAudioOutputClusterOutputInfoStruct) SetIndex(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setIndex:"), value)
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtraudiooutputclusteroutputinfostruct/name
func (m_ MTRAudioOutputClusterOutputInfoStruct) Name() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("name"))
	return rv
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtraudiooutputclusteroutputinfostruct/name
func (m_ MTRAudioOutputClusterOutputInfoStruct) SetName(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setName:"), value)
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtraudiooutputclusteroutputinfostruct/outputtype
func (m_ MTRAudioOutputClusterOutputInfoStruct) OutputType() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("outputType"))
	return rv
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtraudiooutputclusteroutputinfostruct/outputtype
func (m_ MTRAudioOutputClusterOutputInfoStruct) SetOutputType(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setOutputType:"), value)
}
