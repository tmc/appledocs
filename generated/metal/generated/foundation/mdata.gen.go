// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [mdata] class.
var (
	MdataClass     _mdataClass
	MdataClassOnce sync.Once
)

func getmdataClass() _mdataClass {
	MdataClassOnce.Do(func() {
		MdataClass = _mdataClass{objc.GetClass("mdata")}
	})
	return MdataClass
}

type _mdataClass struct {
	class objc.Class
}

// An interface definition for the [mdata] class.
type Imdata interface {
	objectivec.IObject
	// properties:
	// methods:
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSArchiver/mdata
type mdata struct {
	objectivec.Object
}

// mdataFrom constructs a [mdata] from an unsafe.Pointer.
func mdataFrom(ptr unsafe.Pointer) mdata {
	return mdata{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _mdataClass) Alloc() mdata {
	rv := objc.Send[mdata](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _mdataClass) New() mdata {
	rv := objc.Send[mdata](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ mdata) Init() mdata {
	rv := objc.Send[mdata](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ mdata) Autorelease() mdata {
	rv := objc.Send[mdata](m_.ID, objc.Sel("autorelease"))
	return rv
}

// Newmdata creates a new mdata instance.
func Newmdata() mdata {
	return getmdataClass().New()
}




