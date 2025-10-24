// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [MTRModeSelectClusterSemanticTag] class.
var (
	MTRModeSelectClusterSemanticTagClass     _MTRModeSelectClusterSemanticTagClass
	MTRModeSelectClusterSemanticTagClassOnce sync.Once
)

func getMTRModeSelectClusterSemanticTagClass() _MTRModeSelectClusterSemanticTagClass {
	MTRModeSelectClusterSemanticTagClassOnce.Do(func() {
		MTRModeSelectClusterSemanticTagClass = _MTRModeSelectClusterSemanticTagClass{objc.GetClass("MTRModeSelectClusterSemanticTag")}
	})
	return MTRModeSelectClusterSemanticTagClass
}

type _MTRModeSelectClusterSemanticTagClass struct {
	class objc.Class
}

// An interface definition for the [MTRModeSelectClusterSemanticTag] class.
type IMTRModeSelectClusterSemanticTag interface {
	IMTRModeSelectClusterSemanticTagStruct
	// properties:
	MfgCode() objc.IObject /* cross-framework: NSNumber */
	SetMfgCode(value objc.IObject /* cross-framework: NSNumber */)
	Value() objc.IObject /* cross-framework: NSNumber */
	SetValue(value objc.IObject /* cross-framework: NSNumber */)
	// methods:
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRModeSelectClusterSemanticTag
type MTRModeSelectClusterSemanticTag struct {
	MTRModeSelectClusterSemanticTagStruct
}

// MTRModeSelectClusterSemanticTagFrom constructs a [MTRModeSelectClusterSemanticTag] from an unsafe.Pointer.
func MTRModeSelectClusterSemanticTagFrom(ptr unsafe.Pointer) MTRModeSelectClusterSemanticTag {
	return MTRModeSelectClusterSemanticTag{
		MTRModeSelectClusterSemanticTagStruct: MTRModeSelectClusterSemanticTagStructFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRModeSelectClusterSemanticTagClass) Alloc() MTRModeSelectClusterSemanticTag {
	rv := objc.Send[MTRModeSelectClusterSemanticTag](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRModeSelectClusterSemanticTagClass) New() MTRModeSelectClusterSemanticTag {
	rv := objc.Send[MTRModeSelectClusterSemanticTag](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRModeSelectClusterSemanticTag) Init() MTRModeSelectClusterSemanticTag {
	rv := objc.Send[MTRModeSelectClusterSemanticTag](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRModeSelectClusterSemanticTag) Autorelease() MTRModeSelectClusterSemanticTag {
	rv := objc.Send[MTRModeSelectClusterSemanticTag](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRModeSelectClusterSemanticTag creates a new MTRModeSelectClusterSemanticTag instance.
func NewMTRModeSelectClusterSemanticTag() MTRModeSelectClusterSemanticTag {
	return getMTRModeSelectClusterSemanticTagClass().New()
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrmodeselectclustersemantictag/mfgcode
func (m_ MTRModeSelectClusterSemanticTag) MfgCode() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("mfgCode"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrmodeselectclustersemantictag/mfgcode
func (m_ MTRModeSelectClusterSemanticTag) SetMfgCode(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMfgCode:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrmodeselectclustersemantictag/value
func (m_ MTRModeSelectClusterSemanticTag) Value() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("value"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrmodeselectclustersemantictag/value
func (m_ MTRModeSelectClusterSemanticTag) SetValue(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setValue:"), value)
}



