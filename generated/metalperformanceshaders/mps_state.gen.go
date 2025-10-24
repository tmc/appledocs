// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)





// The class instance for the [State] class.
var (
	StateClass     _StateClass
	StateClassOnce sync.Once
)

func getStateClass() _StateClass {
	StateClassOnce.Do(func() {
		StateClass = _StateClass{objc.GetClass("MPSState")}
	})
	return StateClass
}

type _StateClass struct {
	class objc.Class
}





// An interface definition for the [State] class.
type IState interface {
	objectivec.IObject
	

	// properties:
	ReadCount() objectivec.IObject
	SetReadCount(value objectivec.IObject)
	IsTemporary() objectivec.IObject
	SetIsTemporary(value objectivec.IObject)
	Label() objectivec.IObject
	SetLabel(value objectivec.IObject)
	Resource() Resource get /* not a class type */
	SetResource(value Resource get /* not a class type */)
	ResourceCount() objectivec.IObject
	SetResourceCount(value objectivec.IObject)


	

	// methods:
	DestinationImageDescriptor()
	DestinationImageDescriptorForSourceImagesSourceStatesForKernelSuggestedDescriptor(sourceImages unsafe.Pointer, sourceStates unsafe.Pointer, kernel IKernel, inDescriptor IImageDescriptor) IImageDescriptor
	Synchronize()
	SynchronizeOnCommandBuffer(commandBuffer unsafe.Pointer)
	ResourceSize()
	TextureInfo()
	TextureInfoAtIndex(index uint) objc.IObject /* cross-framework: MPSStateTextureInfo */
	ResourceType()
	ResourceTypeAtIndex(index uint) StateResourceType
	BufferSize()
	BufferSizeAtIndex(index uint) uint
	ResourceAtIndexAllocateMemory(index uint, allocateMemory bool) unsafe.Pointer


}





// Alloc allocates a new instance without initialization.
func (sc _StateClass) Alloc() State {
	rv := objc.Send[State](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (sc _StateClass) New() State {
	rv := objc.Send[State](objc.ID(sc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (s_ State) Init() State {
	rv := objc.Send[State](s_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s_ State) Autorelease() State {
	rv := objc.Send[State](s_.ID, objc.Sel("autorelease"))
	return rv
}

// NewState creates a new State instance.
func NewState() State {
	return getStateClass().New()
}





// An opaque data container for large storage in MPS CNN filters.
//
// Some MPS CNN kernels produce additional information beyond an . These may be pooling indices where the result came from, convolution weights, or other information not contained in the usual result from a . An object typically contains one or more expensive objects such as textures or buffers to store this information. It provides a base class with interfaces for managing this storage. Child classes may add additional functionality specific to their contents. Some objects are temporary. Temporary state objects, for example, and , are for very short lived storage, perhaps just a few lines of code within the scope of a single . They are very efficient for storage, as several temporary objects can share the same memory over the course of a command buffer. This can improve both memory usage and time spent in the kernel wiring down memory and such. You may find that some large CNN tasks can not be computed without them, as nontemporary storage would simply take up too much memory. In exchange, the lifetime of the underlying storage in temporary objects needs to be carefully managed. ARC often waits until the end of scope to release objects. Temporary storage often needs to be released sooner than that. Consequently the lifetime of the data in the underlying Metal resources is managed by a property. Each time a reads a temporary object the is automatically decremented. When it reaches 0, the underlying storage is recycled for use by other MPS temporary objects, and the data is becomes undefined. If you need to consume the data multiple times, you should set the to a larger number to prevent the data from becoming undefined. You may set the to 0 yourself to return the storage to MPS, if for any reason, you realize that the object will no longer be used. The contents of a temporary object are only valid from creation to the time the reaches 0. The data is only valid for the on which it was created. Nontemporary objects are valid on any on the same device until they are released.


// An opaque data container for large storage in MPS CNN filters.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSState
type State struct {
	objectivec.Object
}

// StateFrom constructs a [State] from an unsafe.Pointer.
//
// An opaque data container for large storage in MPS CNN filters.
func StateFrom(ptr unsafe.Pointer) State {
	return State{objectivec.Object{objc.ID(ptr)}}
}






// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsstate/2942392-initwithdevice
func NewStateWithDeviceBufferSize(device unsafe.Pointer, bufferSize uintptr /* not a class type */) State {
	instance := getStateClass().Alloc()
	rv := objc.Send[State](instance.ID, objc.Sel("initWithDevice:bufferSize:"), device, bufferSize)
	rv.Autorelease()
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsstate/2947908-initwithdevice
func NewStateWithDeviceResourceList(device unsafe.Pointer, resourceList IStateResourceList) State {
	instance := getStateClass().Alloc()
	rv := objc.Send[State](instance.ID, objc.Sel("initWithDevice:resourceList:"), device, resourceList)
	rv.Autorelease()
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsstate/2942400-initwithdevice
func NewStateWithDeviceTextureDescriptor(device unsafe.Pointer, descriptor metal.TextureDescriptor) State {
	instance := getStateClass().Alloc()
	rv := objc.Send[State](instance.ID, objc.Sel("initWithDevice:textureDescriptor:"), device, descriptor)
	rv.Autorelease()
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsstate/2942390-initwithresource
func NewStateWithResource(resource unsafe.Pointer) State {
	instance := getStateClass().Alloc()
	rv := objc.Send[State](instance.ID, objc.Sel("initWithResource:"), resource)
	rv.Autorelease()
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsstate/2947895-initwithresources
func NewStateWithResources(resources unsafe.Pointer) State {
	instance := getStateClass().Alloc()
	rv := objc.Send[State](instance.ID, objc.Sel("initWithResources:"), resources)
	rv.Autorelease()
	return rv
}







// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsstate/2942391-temporarystate
func (sc _StateClass) TemporaryState() {
	objc.Send[objc.ID](objc.ID(sc.class), objc.Sel("temporaryState"))
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsstate/2942391-temporarystatewithcommandbuffer
func (sc _StateClass) TemporaryStateWithCommandBufferBufferSize(cmdBuf unsafe.Pointer, bufferSize uintptr /* not a class type */) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(sc.class), objc.Sel("temporaryStateWithCommandBuffer:bufferSize:"), cmdBuf, bufferSize)
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsstate/2942393-temporarystatewithcommandbuffer
func (sc _StateClass) TemporaryStateWithCommandBuffer(cmdBuf unsafe.Pointer) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(sc.class), objc.Sel("temporaryStateWithCommandBuffer:"), cmdBuf)
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsstate/2942395-temporarystatewithcommandbuffer
func (sc _StateClass) TemporaryStateWithCommandBufferTextureDescriptor(cmdBuf unsafe.Pointer, descriptor metal.TextureDescriptor) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(sc.class), objc.Sel("temporaryStateWithCommandBuffer:textureDescriptor:"), cmdBuf, descriptor)
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsstate/2947915-temporarystatewithcommandbuffer
func (sc _StateClass) TemporaryStateWithCommandBufferResourceList(commandBuffer unsafe.Pointer, resourceList IStateResourceList) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(sc.class), objc.Sel("temporaryStateWithCommandBuffer:resourceList:"), commandBuffer, resourceList)
	return rv
}












// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsstate/2942394-destinationimagedescriptor
func (s_ State) DestinationImageDescriptor() {
	objc.Send[objc.ID](s_.ID, objc.Sel("destinationImageDescriptor"))
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsstate/2942394-destinationimagedescriptorforsou
func (s_ State) DestinationImageDescriptorForSourceImagesSourceStatesForKernelSuggestedDescriptor(sourceImages unsafe.Pointer, sourceStates unsafe.Pointer, kernel IKernel, inDescriptor IImageDescriptor) IImageDescriptor {
	rv := objc.Send[ImageDescriptor](s_.ID, objc.Sel("destinationImageDescriptorForSourceImages:sourceStates:forKernel:suggestedDescriptor:"), sourceImages, sourceStates, kernel, inDescriptor)
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsstate/2942396-synchronize
func (s_ State) Synchronize() {
	objc.Send[objc.ID](s_.ID, objc.Sel("synchronize"))
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsstate/2942396-synchronizeoncommandbuffer
func (s_ State) SynchronizeOnCommandBuffer(commandBuffer unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("synchronizeOnCommandBuffer:"), commandBuffer)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsstate/2942397-resourcesize
func (s_ State) ResourceSize() {
	objc.Send[objc.ID](s_.ID, objc.Sel("resourceSize"))
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsstate/2947899-textureinfo
func (s_ State) TextureInfo() {
	objc.Send[objc.ID](s_.ID, objc.Sel("textureInfo"))
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsstate/2947899-textureinfoatindex
func (s_ State) TextureInfoAtIndex(index uint) objc.IObject /* cross-framework: MPSStateTextureInfo */ {
	rv := objc.Send[objc.ID](s_.ID, objc.Sel("textureInfoAtIndex:"), index)
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsstate/2947902-resourcetype
func (s_ State) ResourceType() {
	objc.Send[objc.ID](s_.ID, objc.Sel("resourceType"))
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsstate/2947902-resourcetypeatindex
func (s_ State) ResourceTypeAtIndex(index uint) StateResourceType {
	rv := objc.Send[StateResourceType](s_.ID, objc.Sel("resourceTypeAtIndex:"), index)
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsstate/2947913-buffersize
func (s_ State) BufferSize() {
	objc.Send[objc.ID](s_.ID, objc.Sel("bufferSize"))
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsstate/2947913-buffersizeatindex
func (s_ State) BufferSizeAtIndex(index uint) uint {
	rv := objc.Send[uint](s_.ID, objc.Sel("bufferSizeAtIndex:"), index)
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsstate/2947916-resourceatindex
func (s_ State) ResourceAtIndexAllocateMemory(index uint, allocateMemory bool) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("resourceAtIndex:allocateMemory:"), index, allocateMemory)
	return rv
}







// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsstate/2867042-readcount
func (s_ State) ReadCount() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](s_.ID, objc.Sel("readCount"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsstate/2867042-readcount
func (s_ State) SetReadCount(value objectivec.IObject) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setReadCount:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsstate/2867114-istemporary
func (s_ State) IsTemporary() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](s_.ID, objc.Sel("isTemporary"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsstate/2867114-istemporary
func (s_ State) SetIsTemporary(value objectivec.IObject) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setIsTemporary:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsstate/2867179-label
func (s_ State) Label() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](s_.ID, objc.Sel("label"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsstate/2867179-label
func (s_ State) SetLabel(value objectivec.IObject) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setLabel:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsstate/2942398-resource
func (s_ State) Resource() Resource get /* not a class type */ {
	rv := objc.Send[objc.ID](s_.ID, objc.Sel("resource"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsstate/2942398-resource
func (s_ State) SetResource(value Resource get /* not a class type */) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setResource:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsstate/2947900-resourcecount
func (s_ State) ResourceCount() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](s_.ID, objc.Sel("resourceCount"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsstate/2947900-resourcecount
func (s_ State) SetResourceCount(value objectivec.IObject) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setResourceCount:"), value)
}







