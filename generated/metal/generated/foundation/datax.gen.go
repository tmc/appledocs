// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [datax] class.
var (
	DataxClass     _dataxClass
	DataxClassOnce sync.Once
)

func getdataxClass() _dataxClass {
	DataxClassOnce.Do(func() {
		DataxClass = _dataxClass{objc.GetClass("datax")}
	})
	return DataxClass
}

type _dataxClass struct {
	class objc.Class
}

// An interface definition for the [datax] class.
type Idatax interface {
	objectivec.IObject
	// properties:
	// methods:
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSUnarchiver/datax
type datax struct {
	objectivec.Object
}

// dataxFrom constructs a [datax] from an unsafe.Pointer.
func dataxFrom(ptr unsafe.Pointer) datax {
	return datax{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (dc _dataxClass) Alloc() datax {
	rv := objc.Send[datax](objc.ID(dc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (dc _dataxClass) New() datax {
	rv := objc.Send[datax](objc.ID(dc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (d_ datax) Init() datax {
	rv := objc.Send[datax](d_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (d_ datax) Autorelease() datax {
	rv := objc.Send[datax](d_.ID, objc.Sel("autorelease"))
	return rv
}

// Newdatax creates a new datax instance.
func Newdatax() datax {
	return getdataxClass().New()
}




