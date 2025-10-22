// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [localPort] class.
var (
	LocalPortClass     _localPortClass
	LocalPortClassOnce sync.Once
)

func getlocalPortClass() _localPortClass {
	LocalPortClassOnce.Do(func() {
		LocalPortClass = _localPortClass{objc.GetClass("localPort")}
	})
	return LocalPortClass
}

type _localPortClass struct {
	class objc.Class
}

// An interface definition for the [localPort] class.
type IlocalPort interface {
	objectivec.IObject
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSPortMessage/localPort

type localPort struct {
	objectivec.Object
}

// localPortFrom constructs a [localPort] from an unsafe.Pointer.
func localPortFrom(ptr unsafe.Pointer) localPort {
	return localPort{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (lc _localPortClass) Alloc() localPort {
	rv := objc.Send[localPort](objc.ID(lc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (lc _localPortClass) New() localPort {
	rv := objc.Send[localPort](objc.ID(lc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (l_ localPort) Init() localPort {
	rv := objc.Send[localPort](l_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (l_ localPort) Autorelease() localPort {
	rv := objc.Send[localPort](l_.ID, objc.Sel("autorelease"))
	return rv
}

// NewlocalPort creates a new localPort instance.
func NewlocalPort() localPort {
	return getlocalPortClass().New()
}




