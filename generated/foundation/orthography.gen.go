// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [Orthography] class.
var (
	orthographyClass     _OrthographyClass
	orthographyClassOnce sync.Once
)

func getOrthographyClass() _OrthographyClass {
	orthographyClassOnce.Do(func() {
		orthographyClass = _OrthographyClass{objc.GetClass("NSOrthography")}
	})
	return orthographyClass
}

type _OrthographyClass struct {
	class objc.Class
}

// An interface definition for the [Orthography] class.
type IOrthography interface {
	objectivec.IObject
}

// A description of the linguistic content of natural language text, typically used for spelling and grammar checking.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSOrthography
type Orthography struct {
	objectivec.Object
}

// OrthographyFrom constructs a [Orthography] from an unsafe.Pointer.
//
// A description of the linguistic content of natural language text, typically used for spelling and grammar checking.
func OrthographyFrom(ptr unsafe.Pointer) Orthography {
	return Orthography{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (oc _OrthographyClass) Alloc() Orthography {
	rv := objc.Send[Orthography](objc.ID(oc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (oc _OrthographyClass) New() Orthography {
	rv := objc.Send[Orthography](objc.ID(oc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (o_ Orthography) Init() Orthography {
	rv := objc.Send[Orthography](o_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (o_ Orthography) Autorelease() Orthography {
	rv := objc.Send[Orthography](o_.ID, objc.Sel("autorelease"))
	return rv
}

// NewOrthography creates a new Orthography instance.
func NewOrthography() Orthography {
	return getOrthographyClass().New()
}




