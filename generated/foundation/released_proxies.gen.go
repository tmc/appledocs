// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [releasedProxies] class.
var (
	ReleasedProxiesClass     _releasedProxiesClass
	ReleasedProxiesClassOnce sync.Once
)

func getreleasedProxiesClass() _releasedProxiesClass {
	ReleasedProxiesClassOnce.Do(func() {
		ReleasedProxiesClass = _releasedProxiesClass{objc.GetClass("releasedProxies")}
	})
	return ReleasedProxiesClass
}

type _releasedProxiesClass struct {
	class objc.Class
}

// An interface definition for the [releasedProxies] class.
type IreleasedProxies interface {
	objectivec.IObject
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSConnection/releasedProxies

type releasedProxies struct {
	objectivec.Object
}

// releasedProxiesFrom constructs a [releasedProxies] from an unsafe.Pointer.
func releasedProxiesFrom(ptr unsafe.Pointer) releasedProxies {
	return releasedProxies{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (rc _releasedProxiesClass) Alloc() releasedProxies {
	rv := objc.Send[releasedProxies](objc.ID(rc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (rc _releasedProxiesClass) New() releasedProxies {
	rv := objc.Send[releasedProxies](objc.ID(rc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (r_ releasedProxies) Init() releasedProxies {
	rv := objc.Send[releasedProxies](r_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (r_ releasedProxies) Autorelease() releasedProxies {
	rv := objc.Send[releasedProxies](r_.ID, objc.Sel("autorelease"))
	return rv
}

// NewreleasedProxies creates a new releasedProxies instance.
func NewreleasedProxies() releasedProxies {
	return getreleasedProxiesClass().New()
}




