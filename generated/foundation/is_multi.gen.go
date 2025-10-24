// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [isMulti] class.
var (
	IsMultiClass     _isMultiClass
	IsMultiClassOnce sync.Once
)

func getisMultiClass() _isMultiClass {
	IsMultiClassOnce.Do(func() {
		IsMultiClass = _isMultiClass{objc.GetClass("isMulti")}
	})
	return IsMultiClass
}

type _isMultiClass struct {
	class objc.Class
}

// An interface definition for the [isMulti] class.
type IisMulti interface {
	objectivec.IObject
	// properties:
	// methods:
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSConnection/isMulti
type isMulti struct {
	objectivec.Object
}

// isMultiFrom constructs a [isMulti] from an unsafe.Pointer.
func isMultiFrom(ptr unsafe.Pointer) isMulti {
	return isMulti{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (ic _isMultiClass) Alloc() isMulti {
	rv := objc.Send[isMulti](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ic _isMultiClass) New() isMulti {
	rv := objc.Send[isMulti](objc.ID(ic.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (i_ isMulti) Init() isMulti {
	rv := objc.Send[isMulti](i_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i_ isMulti) Autorelease() isMulti {
	rv := objc.Send[isMulti](i_.ID, objc.Sel("autorelease"))
	return rv
}

// NewisMulti creates a new isMulti instance.
func NewisMulti() isMulti {
	return getisMultiClass().New()
}




