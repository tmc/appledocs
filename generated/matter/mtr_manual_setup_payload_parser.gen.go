// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [MTRManualSetupPayloadParser] class.
var (
	MTRManualSetupPayloadParserClass     _MTRManualSetupPayloadParserClass
	MTRManualSetupPayloadParserClassOnce sync.Once
)

func getMTRManualSetupPayloadParserClass() _MTRManualSetupPayloadParserClass {
	MTRManualSetupPayloadParserClassOnce.Do(func() {
		MTRManualSetupPayloadParserClass = _MTRManualSetupPayloadParserClass{objc.GetClass("MTRManualSetupPayloadParser")}
	})
	return MTRManualSetupPayloadParserClass
}

type _MTRManualSetupPayloadParserClass struct {
	class objc.Class
}

// An interface definition for the [MTRManualSetupPayloadParser] class.
type IMTRManualSetupPayloadParser interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRManualSetupPayloadParser
type MTRManualSetupPayloadParser struct {
	objectivec.Object
}

// MTRManualSetupPayloadParserFrom constructs a [MTRManualSetupPayloadParser] from an unsafe.Pointer.
func MTRManualSetupPayloadParserFrom(ptr unsafe.Pointer) MTRManualSetupPayloadParser {
	return MTRManualSetupPayloadParser{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRManualSetupPayloadParserClass) Alloc() MTRManualSetupPayloadParser {
	rv := objc.Send[MTRManualSetupPayloadParser](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRManualSetupPayloadParserClass) New() MTRManualSetupPayloadParser {
	rv := objc.Send[MTRManualSetupPayloadParser](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRManualSetupPayloadParser) Init() MTRManualSetupPayloadParser {
	rv := objc.Send[MTRManualSetupPayloadParser](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRManualSetupPayloadParser) Autorelease() MTRManualSetupPayloadParser {
	rv := objc.Send[MTRManualSetupPayloadParser](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRManualSetupPayloadParser creates a new MTRManualSetupPayloadParser instance.
func NewMTRManualSetupPayloadParser() MTRManualSetupPayloadParser {
	return getMTRManualSetupPayloadParserClass().New()
}




