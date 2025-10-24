// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/metal"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MPSStateResourceList */


/* debug [class_header]: Header for MPSStateResourceList */
// The class instance for the [StateResourceList] class.
var (
	StateResourceListClass     _StateResourceListClass
	StateResourceListClassOnce sync.Once
)

func getStateResourceListClass() _StateResourceListClass {
	StateResourceListClassOnce.Do(func() {
		StateResourceListClass = _StateResourceListClass{objc.GetClass("MPSStateResourceList")}
	})
	return StateResourceListClass
}

type _StateResourceListClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for StateResourceList */
// An interface definition for the [StateResourceList] class.
type IStateResourceList interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for StateResourceList */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for StateResourceList */
	// methods:
	AppendTexture()
	AppendBuffer()
	AppendTextureWithDescriptor(descriptor metal.TextureDescriptor)
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for StateResourceList */
// Alloc allocates a new instance without initialization.
func (sc _StateResourceListClass) Alloc() StateResourceList {
	rv := objc.Send[StateResourceList](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (sc _StateResourceListClass) New() StateResourceList {
	rv := objc.Send[StateResourceList](objc.ID(sc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (s_ StateResourceList) Init() StateResourceList {
	rv := objc.Send[StateResourceList](s_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s_ StateResourceList) Autorelease() StateResourceList {
	rv := objc.Send[StateResourceList](s_.ID, objc.Sel("autorelease"))
	return rv
}

// NewStateResourceList creates a new StateResourceList instance.
func NewStateResourceList() StateResourceList {
	return getStateResourceListClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for StateResourceList */
// An interface for objects that define resources for Metal Performance Shaders state containers.


// An interface for objects that define resources for Metal Performance Shaders state containers.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSStateResourceList
type StateResourceList struct {
	objectivec.Object
}

// StateResourceListFrom constructs a [StateResourceList] from an unsafe.Pointer.
//
// An interface for objects that define resources for Metal Performance Shaders state containers.
func StateResourceListFrom(ptr unsafe.Pointer) StateResourceList {
	return StateResourceList{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for StateResourceList *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for StateResourceList */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsstateresourcelist/2947890-resourcelistwithtexturedescripto
func (sc _StateResourceListClass) ResourceListWithTextureDescriptors(d metal.TextureDescriptor) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(sc.class), objc.Sel("resourceListWithTextureDescriptors:"), d)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ResourceListWithTextureDescriptors) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsstateresourcelist/2947903-resourcelistwithbuffersizes
func (sc _StateResourceListClass) ResourceListWithBufferSizes(firstSize uint) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(sc.class), objc.Sel("resourceListWithBufferSizes:"), firstSize)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ResourceListWithBufferSizes) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsstateresourcelist/2947904-resourcelist
func (sc _StateResourceListClass) ResourceList() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(sc.class), objc.Sel("resourceList"))
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ResourceList) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for StateResourceList */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for StateResourceList */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsstateresourcelist/2947894-appendtexture
func (s_ StateResourceList) AppendTexture() {
	objc.Send[objc.ID](s_.ID, objc.Sel("appendTexture"))
}/* debug [instance_methods/method]: AppendTexture */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsstateresourcelist/2947905-appendbuffer
func (s_ StateResourceList) AppendBuffer() {
	objc.Send[objc.ID](s_.ID, objc.Sel("appendBuffer"))
}/* debug [instance_methods/method]: AppendBuffer */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSStateResourceList/appendTexture(_:)
func (s_ StateResourceList) AppendTextureWithDescriptor(descriptor metal.TextureDescriptor) {
	objc.Send[objc.ID](s_.ID, objc.Sel("appendTexture:"), descriptor)
}/* debug [instance_methods/method]: AppendTextureWithDescriptor */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for StateResourceList */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MPSStateResourceList */



