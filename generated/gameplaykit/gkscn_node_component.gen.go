// Code generated from Apple documentation for GameplayKit. DO NOT EDIT.

package gameplaykit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [SCNNodeComponent] class.
var (
	SCNNodeComponentClass     _SCNNodeComponentClass
	SCNNodeComponentClassOnce sync.Once
)

func getSCNNodeComponentClass() _SCNNodeComponentClass {
	SCNNodeComponentClassOnce.Do(func() {
		SCNNodeComponentClass = _SCNNodeComponentClass{objc.GetClass("GKSCNNodeComponent")}
	})
	return SCNNodeComponentClass
}

type _SCNNodeComponentClass struct {
	class objc.Class
}

// An interface definition for the [SCNNodeComponent] class.
type ISCNNodeComponent interface {
	IComponent
	Node() unsafe.Pointer
}

//
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKSCNNodeComponent
type SCNNodeComponent struct {
	Component
}

// SCNNodeComponentFrom constructs a [SCNNodeComponent] from an unsafe.Pointer.
func SCNNodeComponentFrom(ptr unsafe.Pointer) SCNNodeComponent {
	return SCNNodeComponent{
		Component: ComponentFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (nc _SCNNodeComponentClass) Alloc() SCNNodeComponent {
	rv := objc.Send[SCNNodeComponent](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (nc _SCNNodeComponentClass) New() SCNNodeComponent {
	rv := objc.Send[SCNNodeComponent](objc.ID(nc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (n_ SCNNodeComponent) Init() SCNNodeComponent {
	rv := objc.Send[SCNNodeComponent](n_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n_ SCNNodeComponent) Autorelease() SCNNodeComponent {
	rv := objc.Send[SCNNodeComponent](n_.ID, objc.Sel("autorelease"))
	return rv
}

// NewSCNNodeComponent creates a new SCNNodeComponent instance.
func NewSCNNodeComponent() SCNNodeComponent {
	return getSCNNodeComponentClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKSCNNodeComponent/init(node:)
func NewSCNNodeComponentWithNode(node unsafe.Pointer) SCNNodeComponent {
	instance := getSCNNodeComponentClass().Alloc()
	rv := objc.Send[SCNNodeComponent](instance.ID, objc.Sel("initWithNode:"), node)
	rv.Autorelease()
	return rv
}


//
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKSCNNodeComponent/componentWithNode:
func (nc _SCNNodeComponentClass) ComponentWithNode(node unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(nc.class), objc.Sel("componentWithNode:"), node)
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKSCNNodeComponent/node
func (n_ SCNNodeComponent) Node() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](n_.ID, objc.Sel("node"))
	return rv
}


