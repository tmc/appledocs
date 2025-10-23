// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [isValid] class.
var (
	IsValidClass     _isValidClass
	IsValidClassOnce sync.Once
)

func getisValidClass() _isValidClass {
	IsValidClassOnce.Do(func() {
		IsValidClass = _isValidClass{objc.GetClass("isValid")}
	})
	return IsValidClass
}

type _isValidClass struct {
	class objc.Class
}

// An interface definition for the [isValid] class.
type IisValid interface {
	objectivec.IObject
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSConnection/isValid
type isValid struct {
	objectivec.Object
}

// isValidFrom constructs a [isValid] from an unsafe.Pointer.
func isValidFrom(ptr unsafe.Pointer) isValid {
	return isValid{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (ic _isValidClass) Alloc() isValid {
	rv := objc.Send[isValid](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ic _isValidClass) New() isValid {
	rv := objc.Send[isValid](objc.ID(ic.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (i_ isValid) Init() isValid {
	rv := objc.Send[isValid](i_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i_ isValid) Autorelease() isValid {
	rv := objc.Send[isValid](i_.ID, objc.Sel("autorelease"))
	return rv
}

// NewisValid creates a new isValid instance.
func NewisValid() isValid {
	return getisValidClass().New()
}




