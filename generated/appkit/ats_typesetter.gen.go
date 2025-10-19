// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [ATSTypesetter] class.
var (
	aTSTypesetterClass     _ATSTypesetterClass
	aTSTypesetterClassOnce sync.Once
)

func getATSTypesetterClass() _ATSTypesetterClass {
	aTSTypesetterClassOnce.Do(func() {
		aTSTypesetterClass = _ATSTypesetterClass{objc.GetClass("NSATSTypesetter")}
	})
	return aTSTypesetterClass
}

type _ATSTypesetterClass struct {
	class objc.Class
}

// An interface definition for the [ATSTypesetter] class.
type IATSTypesetter interface {
	ITypesetter
}

// A concrete typesetter object that places glyphs during the text layout process. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSATSTypesetter

type ATSTypesetter struct {
	Typesetter
}

// ATSTypesetterFrom constructs a [ATSTypesetter] from an unsafe.Pointer.
//
// A concrete typesetter object that places glyphs during the text layout process.
func ATSTypesetterFrom(ptr unsafe.Pointer) ATSTypesetter {
	return ATSTypesetter{
		Typesetter: TypesetterFrom(ptr),
	}
}
// Alloc allocates a new instance without initialization.
func (ac _ATSTypesetterClass) Alloc() ATSTypesetter {
	rv := objc.Send[ATSTypesetter](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new instance with a +1 retain count.
func (ac _ATSTypesetterClass) New() ATSTypesetter {
	rv := objc.Send[ATSTypesetter](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ ATSTypesetter) Init() ATSTypesetter {
	rv := objc.Send[ATSTypesetter](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ ATSTypesetter) Autorelease() ATSTypesetter {
	rv := objc.Send[ATSTypesetter](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewATSTypesetter creates a new ATSTypesetter instance.
func NewATSTypesetter() ATSTypesetter {
	return getATSTypesetterClass().New()
}




