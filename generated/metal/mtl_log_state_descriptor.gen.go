// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [LogStateDescriptor] class.
var (
	LogStateDescriptorClass     _LogStateDescriptorClass
	LogStateDescriptorClassOnce sync.Once
)

func getLogStateDescriptorClass() _LogStateDescriptorClass {
	LogStateDescriptorClassOnce.Do(func() {
		LogStateDescriptorClass = _LogStateDescriptorClass{objc.GetClass("MTLLogStateDescriptor")}
	})
	return LogStateDescriptorClass
}

type _LogStateDescriptorClass struct {
	class objc.Class
}

// An interface definition for the [LogStateDescriptor] class.
type ILogStateDescriptor interface {
	objectivec.IObject
}

// An interface that represents a log state configuration.
//
// Configure the descriptor to create an by calling . If you’ve set the environment variables or , then the system automatically enables logging. If any command buffer or command queue has an attached log state, then the system uses the log state’s settings instead of the environment variable values.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLLogStateDescriptor
type LogStateDescriptor struct {
	objectivec.Object
}

// LogStateDescriptorFrom constructs a [LogStateDescriptor] from an unsafe.Pointer.
//
// An interface that represents a log state configuration.
func LogStateDescriptorFrom(ptr unsafe.Pointer) LogStateDescriptor {
	return LogStateDescriptor{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (lc _LogStateDescriptorClass) Alloc() LogStateDescriptor {
	rv := objc.Send[LogStateDescriptor](objc.ID(lc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (lc _LogStateDescriptorClass) New() LogStateDescriptor {
	rv := objc.Send[LogStateDescriptor](objc.ID(lc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (l_ LogStateDescriptor) Init() LogStateDescriptor {
	rv := objc.Send[LogStateDescriptor](l_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (l_ LogStateDescriptor) Autorelease() LogStateDescriptor {
	rv := objc.Send[LogStateDescriptor](l_.ID, objc.Sel("autorelease"))
	return rv
}

// NewLogStateDescriptor creates a new LogStateDescriptor instance.
func NewLogStateDescriptor() LogStateDescriptor {
	return getLogStateDescriptorClass().New()
}


// The size of the internal buffer the log state uses, specified in bytes.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLLogStateDescriptor/bufferSize
func (l_ LogStateDescriptor) BufferSize() int {
	rv := objc.Send[int](l_.ID, objc.Sel("bufferSize"))
	return rv
}


// SetBufferSize sets the value of the bufferSize property.
// The size of the internal buffer the log state uses, specified in bytes.

//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLLogStateDescriptor/bufferSize
func (l_ LogStateDescriptor) SetBufferSize(value int) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setBufferSize:"), value)
}

// The minimum level of messages that the shader can log.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLLogStateDescriptor/level
func (l_ LogStateDescriptor) Level() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](l_.ID, objc.Sel("level"))
	return rv
}


// SetLevel sets the value of the level property.
// The minimum level of messages that the shader can log.

//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLLogStateDescriptor/level
func (l_ LogStateDescriptor) SetLevel(value unsafe.Pointer) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setLevel:"), value)
}



