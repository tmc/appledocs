// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [padding] class.
var (
	PaddingClass     _paddingClass
	PaddingClassOnce sync.Once
)

func getpaddingClass() _paddingClass {
	PaddingClassOnce.Do(func() {
		PaddingClass = _paddingClass{objc.GetClass("padding")}
	})
	return PaddingClass
}

type _paddingClass struct {
	class objc.Class
}

// An interface definition for the [padding] class.
type Ipadding interface {
	objectivec.IObject
	// properties:
	// methods:
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSXMLDocument/padding
type padding struct {
	objectivec.Object
}

// paddingFrom constructs a [padding] from an unsafe.Pointer.
func paddingFrom(ptr unsafe.Pointer) padding {
	return padding{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (pc _paddingClass) Alloc() padding {
	rv := objc.Send[padding](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (pc _paddingClass) New() padding {
	rv := objc.Send[padding](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ padding) Init() padding {
	rv := objc.Send[padding](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ padding) Autorelease() padding {
	rv := objc.Send[padding](p_.ID, objc.Sel("autorelease"))
	return rv
}

// Newpadding creates a new padding instance.
func Newpadding() padding {
	return getpaddingClass().New()
}




