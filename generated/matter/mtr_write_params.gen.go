// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRWriteParams] class.
var (
	MTRWriteParamsClass     _MTRWriteParamsClass
	MTRWriteParamsClassOnce sync.Once
)

func getMTRWriteParamsClass() _MTRWriteParamsClass {
	MTRWriteParamsClassOnce.Do(func() {
		MTRWriteParamsClass = _MTRWriteParamsClass{objc.GetClass("MTRWriteParams")}
	})
	return MTRWriteParamsClass
}

type _MTRWriteParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTRWriteParams] class.
type IMTRWriteParams interface {
	objectivec.IObject
	// properties:
	DataVersion() objc.IObject /* cross-framework: NSNumber */
	SetDataVersion(value objc.IObject /* cross-framework: NSNumber */)
	TimedWriteTimeout() objc.IObject /* cross-framework: NSNumber */
	SetTimedWriteTimeout(value objc.IObject /* cross-framework: NSNumber */)
	// methods:
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRWriteParams
type MTRWriteParams struct {
	objectivec.Object
}

// MTRWriteParamsFrom constructs a [MTRWriteParams] from an unsafe.Pointer.
func MTRWriteParamsFrom(ptr unsafe.Pointer) MTRWriteParams {
	return MTRWriteParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRWriteParamsClass) Alloc() MTRWriteParams {
	rv := objc.Send[MTRWriteParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRWriteParamsClass) New() MTRWriteParams {
	rv := objc.Send[MTRWriteParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRWriteParams) Init() MTRWriteParams {
	rv := objc.Send[MTRWriteParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRWriteParams) Autorelease() MTRWriteParams {
	rv := objc.Send[MTRWriteParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRWriteParams creates a new MTRWriteParams instance.
func NewMTRWriteParams() MTRWriteParams {
	return getMTRWriteParamsClass().New()
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrwriteparams/dataversion
func (m_ MTRWriteParams) DataVersion() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("dataVersion"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrwriteparams/dataversion
func (m_ MTRWriteParams) SetDataVersion(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setDataVersion:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrwriteparams/timedwritetimeout
func (m_ MTRWriteParams) TimedWriteTimeout() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("timedWriteTimeout"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrwriteparams/timedwritetimeout
func (m_ MTRWriteParams) SetTimedWriteTimeout(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedWriteTimeout:"), value)
}



