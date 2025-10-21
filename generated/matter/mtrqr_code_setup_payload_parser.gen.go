// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRQRCodeSetupPayloadParser] class.
var (
	MTRQRCodeSetupPayloadParserClass     _MTRQRCodeSetupPayloadParserClass
	MTRQRCodeSetupPayloadParserClassOnce sync.Once
)

func getMTRQRCodeSetupPayloadParserClass() _MTRQRCodeSetupPayloadParserClass {
	MTRQRCodeSetupPayloadParserClassOnce.Do(func() {
		MTRQRCodeSetupPayloadParserClass = _MTRQRCodeSetupPayloadParserClass{objc.GetClass("MTRQRCodeSetupPayloadParser")}
	})
	return MTRQRCodeSetupPayloadParserClass
}

type _MTRQRCodeSetupPayloadParserClass struct {
	class objc.Class
}

// An interface definition for the [MTRQRCodeSetupPayloadParser] class.
type IMTRQRCodeSetupPayloadParser interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRQRCodeSetupPayloadParser
type MTRQRCodeSetupPayloadParser struct {
	objectivec.Object
}

// MTRQRCodeSetupPayloadParserFrom constructs a [MTRQRCodeSetupPayloadParser] from an unsafe.Pointer.
func MTRQRCodeSetupPayloadParserFrom(ptr unsafe.Pointer) MTRQRCodeSetupPayloadParser {
	return MTRQRCodeSetupPayloadParser{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRQRCodeSetupPayloadParserClass) Alloc() MTRQRCodeSetupPayloadParser {
	rv := objc.Send[MTRQRCodeSetupPayloadParser](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRQRCodeSetupPayloadParserClass) New() MTRQRCodeSetupPayloadParser {
	rv := objc.Send[MTRQRCodeSetupPayloadParser](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRQRCodeSetupPayloadParser) Init() MTRQRCodeSetupPayloadParser {
	rv := objc.Send[MTRQRCodeSetupPayloadParser](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRQRCodeSetupPayloadParser) Autorelease() MTRQRCodeSetupPayloadParser {
	rv := objc.Send[MTRQRCodeSetupPayloadParser](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRQRCodeSetupPayloadParser creates a new MTRQRCodeSetupPayloadParser instance.
func NewMTRQRCodeSetupPayloadParser() MTRQRCodeSetupPayloadParser {
	return getMTRQRCodeSetupPayloadParserClass().New()
}




