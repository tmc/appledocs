// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [MTROTAHeaderParser] class.
var (
	MTROTAHeaderParserClass     _MTROTAHeaderParserClass
	MTROTAHeaderParserClassOnce sync.Once
)

func getMTROTAHeaderParserClass() _MTROTAHeaderParserClass {
	MTROTAHeaderParserClassOnce.Do(func() {
		MTROTAHeaderParserClass = _MTROTAHeaderParserClass{objc.GetClass("MTROTAHeaderParser")}
	})
	return MTROTAHeaderParserClass
}

type _MTROTAHeaderParserClass struct {
	class objc.Class
}

// An interface definition for the [MTROTAHeaderParser] class.
type IMTROTAHeaderParser interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROTAHeaderParser
type MTROTAHeaderParser struct {
	objectivec.Object
}

// MTROTAHeaderParserFrom constructs a [MTROTAHeaderParser] from an unsafe.Pointer.
func MTROTAHeaderParserFrom(ptr unsafe.Pointer) MTROTAHeaderParser {
	return MTROTAHeaderParser{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTROTAHeaderParserClass) Alloc() MTROTAHeaderParser {
	rv := objc.Send[MTROTAHeaderParser](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTROTAHeaderParserClass) New() MTROTAHeaderParser {
	rv := objc.Send[MTROTAHeaderParser](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTROTAHeaderParser) Init() MTROTAHeaderParser {
	rv := objc.Send[MTROTAHeaderParser](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTROTAHeaderParser) Autorelease() MTROTAHeaderParser {
	rv := objc.Send[MTROTAHeaderParser](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTROTAHeaderParser creates a new MTROTAHeaderParser instance.
func NewMTROTAHeaderParser() MTROTAHeaderParser {
	return getMTROTAHeaderParserClass().New()
}




