// Code generated from Apple documentation for GameplayKit. DO NOT EDIT.

package gameplaykit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class GKSCNNodeComponent */


/* debug [class_header]: Header for GKSCNNodeComponent */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for SCNNodeComponent */
// An interface definition for the [SCNNodeComponent] class.
type ISCNNodeComponent interface {
	IComponent
	
/* debug [class_interface_properties]: Properties for SCNNodeComponent */
	// properties:
	Node() NNode /* not a class type */
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for SCNNodeComponent */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for SCNNodeComponent */
// Alloc allocates a new instance without initialization.
func (nc _SCNNodeComponentClass) Alloc() SCNNodeComponent {
	rv := objc.Send[SCNNodeComponent](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for SCNNodeComponent */


// [Full Topic]
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for SCNNodeComponent */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKSCNNodeComponent/init(node:)
func NewSCNNodeComponentWithNode(node NNode /* not a class type */) SCNNodeComponent {
	instance := getSCNNodeComponentClass().Alloc()
	rv := objc.Send[SCNNodeComponent](instance.ID, objc.Sel("initWithNode:"), node)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewSCNNodeComponentWithNode */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for SCNNodeComponent */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKSCNNodeComponent/componentWithNode:
func (nc _SCNNodeComponentClass) ComponentWithNode(node NNode /* not a class type */) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(nc.class), objc.Sel("componentWithNode:"), node)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ComponentWithNode) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for SCNNodeComponent */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for SCNNodeComponent */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for SCNNodeComponent */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKSCNNodeComponent/node
func (n_ SCNNodeComponent) Node() NNode /* not a class type */ {
	rv := objc.Send[NNode](n_.ID, objc.Sel("node"))
	return rv
}/* debug [instance_properties/getter]: node */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class GKSCNNodeComponent */


