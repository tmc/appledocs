// Code generated from Apple documentation for OpenDirectory. DO NOT EDIT.

package opendirectory

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [nodeName] class.
var (
	NodeNameClass     _nodeNameClass
	NodeNameClassOnce sync.Once
)

func getnodeNameClass() _nodeNameClass {
	NodeNameClassOnce.Do(func() {
		NodeNameClass = _nodeNameClass{objc.GetClass("nodeName")}
	})
	return NodeNameClass
}

type _nodeNameClass struct {
	class objc.Class
}

// An interface definition for the [nodeName] class.
type InodeName interface {
	objectivec.IObject
	// properties:
	// methods:
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODConfiguration/nodeName-c.ivar
type nodeName struct {
	objectivec.Object
}

// nodeNameFrom constructs a [nodeName] from an unsafe.Pointer.
func nodeNameFrom(ptr unsafe.Pointer) nodeName {
	return nodeName{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (nc _nodeNameClass) Alloc() nodeName {
	rv := objc.Send[nodeName](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (nc _nodeNameClass) New() nodeName {
	rv := objc.Send[nodeName](objc.ID(nc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (n_ nodeName) Init() nodeName {
	rv := objc.Send[nodeName](n_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n_ nodeName) Autorelease() nodeName {
	rv := objc.Send[nodeName](n_.ID, objc.Sel("autorelease"))
	return rv
}

// NewnodeName creates a new nodeName instance.
func NewnodeName() nodeName {
	return getnodeNameClass().New()
}




