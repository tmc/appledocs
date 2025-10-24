// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRModeSelectClusterModeOptionStruct] class.
var (
	MTRModeSelectClusterModeOptionStructClass     _MTRModeSelectClusterModeOptionStructClass
	MTRModeSelectClusterModeOptionStructClassOnce sync.Once
)

func getMTRModeSelectClusterModeOptionStructClass() _MTRModeSelectClusterModeOptionStructClass {
	MTRModeSelectClusterModeOptionStructClassOnce.Do(func() {
		MTRModeSelectClusterModeOptionStructClass = _MTRModeSelectClusterModeOptionStructClass{objc.GetClass("MTRModeSelectClusterModeOptionStruct")}
	})
	return MTRModeSelectClusterModeOptionStructClass
}

type _MTRModeSelectClusterModeOptionStructClass struct {
	class objc.Class
}

// An interface definition for the [MTRModeSelectClusterModeOptionStruct] class.
type IMTRModeSelectClusterModeOptionStruct interface {
	objectivec.IObject
	// properties:
	Label() objc.IObject /* cross-framework: NSString */
	SetLabel(value objc.IObject /* cross-framework: NSString */)
	Mode() objc.IObject /* cross-framework: NSNumber */
	SetMode(value objc.IObject /* cross-framework: NSNumber */)
	SemanticTags() unsafe.Pointer
	SetSemanticTags(value unsafe.Pointer)
	// methods:
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRModeSelectClusterModeOptionStruct
type MTRModeSelectClusterModeOptionStruct struct {
	objectivec.Object
}

// MTRModeSelectClusterModeOptionStructFrom constructs a [MTRModeSelectClusterModeOptionStruct] from an unsafe.Pointer.
func MTRModeSelectClusterModeOptionStructFrom(ptr unsafe.Pointer) MTRModeSelectClusterModeOptionStruct {
	return MTRModeSelectClusterModeOptionStruct{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRModeSelectClusterModeOptionStructClass) Alloc() MTRModeSelectClusterModeOptionStruct {
	rv := objc.Send[MTRModeSelectClusterModeOptionStruct](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRModeSelectClusterModeOptionStructClass) New() MTRModeSelectClusterModeOptionStruct {
	rv := objc.Send[MTRModeSelectClusterModeOptionStruct](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRModeSelectClusterModeOptionStruct) Init() MTRModeSelectClusterModeOptionStruct {
	rv := objc.Send[MTRModeSelectClusterModeOptionStruct](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRModeSelectClusterModeOptionStruct) Autorelease() MTRModeSelectClusterModeOptionStruct {
	rv := objc.Send[MTRModeSelectClusterModeOptionStruct](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRModeSelectClusterModeOptionStruct creates a new MTRModeSelectClusterModeOptionStruct instance.
func NewMTRModeSelectClusterModeOptionStruct() MTRModeSelectClusterModeOptionStruct {
	return getMTRModeSelectClusterModeOptionStructClass().New()
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrmodeselectclustermodeoptionstruct/label
func (m_ MTRModeSelectClusterModeOptionStruct) Label() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("label"))
	return rv
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrmodeselectclustermodeoptionstruct/label
func (m_ MTRModeSelectClusterModeOptionStruct) SetLabel(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setLabel:"), value)
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrmodeselectclustermodeoptionstruct/mode
func (m_ MTRModeSelectClusterModeOptionStruct) Mode() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("mode"))
	return rv
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrmodeselectclustermodeoptionstruct/mode
func (m_ MTRModeSelectClusterModeOptionStruct) SetMode(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMode:"), value)
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrmodeselectclustermodeoptionstruct/semantictags
func (m_ MTRModeSelectClusterModeOptionStruct) SemanticTags() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("semanticTags"))
	return rv
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrmodeselectclustermodeoptionstruct/semantictags
func (m_ MTRModeSelectClusterModeOptionStruct) SetSemanticTags(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setSemanticTags:"), value)
}
