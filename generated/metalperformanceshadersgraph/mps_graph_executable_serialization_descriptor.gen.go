// Code generated from Apple documentation for MetalPerformanceShadersGraph. DO NOT EDIT.

package metalperformanceshadersgraph

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

/* debug [class.gen.go]: Generating class MPSGraphExecutableSerializationDescriptor */


/* debug [class_header]: Header for MPSGraphExecutableSerializationDescriptor */
// The class instance for the [GraphExecutableSerializationDescriptor] class.
var (
	GraphExecutableSerializationDescriptorClass     _GraphExecutableSerializationDescriptorClass
	GraphExecutableSerializationDescriptorClassOnce sync.Once
)

func getGraphExecutableSerializationDescriptorClass() _GraphExecutableSerializationDescriptorClass {
	GraphExecutableSerializationDescriptorClassOnce.Do(func() {
		GraphExecutableSerializationDescriptorClass = _GraphExecutableSerializationDescriptorClass{objc.GetClass("MPSGraphExecutableSerializationDescriptor")}
	})
	return GraphExecutableSerializationDescriptorClass
}

type _GraphExecutableSerializationDescriptorClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for GraphExecutableSerializationDescriptor */
// An interface definition for the [GraphExecutableSerializationDescriptor] class.
type IGraphExecutableSerializationDescriptor interface {
	IGraphObject
	
/* debug [class_interface_properties]: Properties for GraphExecutableSerializationDescriptor */
	// properties:
	Append() bool
	SetAppend(value bool)
	DeploymentPlatform() GraphDeploymentPlatform
	SetDeploymentPlatform(value GraphDeploymentPlatform)
	MinimumDeploymentTarget() objc.IObject /* cross-framework: NSString */
	SetMinimumDeploymentTarget(value objc.IObject /* cross-framework: NSString */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for GraphExecutableSerializationDescriptor */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for GraphExecutableSerializationDescriptor */
// Alloc allocates a new instance without initialization.
func (gc _GraphExecutableSerializationDescriptorClass) Alloc() GraphExecutableSerializationDescriptor {
	rv := objc.Send[GraphExecutableSerializationDescriptor](objc.ID(gc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (gc _GraphExecutableSerializationDescriptorClass) New() GraphExecutableSerializationDescriptor {
	rv := objc.Send[GraphExecutableSerializationDescriptor](objc.ID(gc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (g_ GraphExecutableSerializationDescriptor) Init() GraphExecutableSerializationDescriptor {
	rv := objc.Send[GraphExecutableSerializationDescriptor](g_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (g_ GraphExecutableSerializationDescriptor) Autorelease() GraphExecutableSerializationDescriptor {
	rv := objc.Send[GraphExecutableSerializationDescriptor](g_.ID, objc.Sel("autorelease"))
	return rv
}

// NewGraphExecutableSerializationDescriptor creates a new GraphExecutableSerializationDescriptor instance.
func NewGraphExecutableSerializationDescriptor() GraphExecutableSerializationDescriptor {
	return getGraphExecutableSerializationDescriptorClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for GraphExecutableSerializationDescriptor */
// A class that consists of all the levers to serialize an executable.


// A class that consists of all the levers to serialize an executable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphExecutableSerializationDescriptor
type GraphExecutableSerializationDescriptor struct {
	GraphObject
}

// GraphExecutableSerializationDescriptorFrom constructs a [GraphExecutableSerializationDescriptor] from an unsafe.Pointer.
//
// A class that consists of all the levers to serialize an executable.
func GraphExecutableSerializationDescriptorFrom(ptr unsafe.Pointer) GraphExecutableSerializationDescriptor {
	return GraphExecutableSerializationDescriptor{
		GraphObject: GraphObjectFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for GraphExecutableSerializationDescriptor *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for GraphExecutableSerializationDescriptor */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for GraphExecutableSerializationDescriptor */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for GraphExecutableSerializationDescriptor */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for GraphExecutableSerializationDescriptor */

// Flag to append to an existing .mpsgraphpackage if found at provided url.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphExecutableSerializationDescriptor/append
func (g_ GraphExecutableSerializationDescriptor) Append() bool {
	rv := objc.Send[bool](g_.ID, objc.Sel("append"))
	return rv
}/* debug [instance_properties/getter]: append */


// Flag to append to an existing .mpsgraphpackage if found at provided url.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphExecutableSerializationDescriptor/append
func (g_ GraphExecutableSerializationDescriptor) SetAppend(value bool) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setAppend:"), value)
}/* debug [instance_properties/setter]: append */


// The deployment platform used to serialize the executable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphExecutableSerializationDescriptor/deploymentPlatform
func (g_ GraphExecutableSerializationDescriptor) DeploymentPlatform() GraphDeploymentPlatform {
	rv := objc.Send[GraphDeploymentPlatform](g_.ID, objc.Sel("deploymentPlatform"))
	return rv
}/* debug [instance_properties/getter]: deploymentPlatform */


// The deployment platform used to serialize the executable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphExecutableSerializationDescriptor/deploymentPlatform
func (g_ GraphExecutableSerializationDescriptor) SetDeploymentPlatform(value GraphDeploymentPlatform) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setDeploymentPlatform:"), value)
}/* debug [instance_properties/setter]: deploymentPlatform */


// The minimum deployment target to serialize the executable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphExecutableSerializationDescriptor/minimumDeploymentTarget
func (g_ GraphExecutableSerializationDescriptor) MinimumDeploymentTarget() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](g_.ID, objc.Sel("minimumDeploymentTarget"))
	return rv
}/* debug [instance_properties/getter]: minimumDeploymentTarget */


// The minimum deployment target to serialize the executable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphExecutableSerializationDescriptor/minimumDeploymentTarget
func (g_ GraphExecutableSerializationDescriptor) SetMinimumDeploymentTarget(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setMinimumDeploymentTarget:"), value)
}/* debug [instance_properties/setter]: minimumDeploymentTarget */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MPSGraphExecutableSerializationDescriptor */



