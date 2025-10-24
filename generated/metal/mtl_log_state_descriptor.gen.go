// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTLLogStateDescriptor */


/* debug [class_header]: Header for MTLLogStateDescriptor */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for LogStateDescriptor */
// An interface definition for the [LogStateDescriptor] class.
type ILogStateDescriptor interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for LogStateDescriptor */
	// properties:
	BufferSize() int
	SetBufferSize(value int)
	Level() LogLevel
	SetLevel(value LogLevel)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for LogStateDescriptor */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for LogStateDescriptor */
// Alloc allocates a new instance without initialization.
func (lc _LogStateDescriptorClass) Alloc() LogStateDescriptor {
	rv := objc.Send[LogStateDescriptor](objc.ID(lc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for LogStateDescriptor */
// An interface that represents a log state configuration.
//
// Configure the descriptor to create an by calling . If you’ve set the environment variables or , then the system automatically enables logging. If any command buffer or command queue has an attached log state, then the system uses the log state’s settings instead of the environment variable values.


// An interface that represents a log state configuration.
//
// [Full Topic]
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for LogStateDescriptor *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for LogStateDescriptor */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for LogStateDescriptor */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for LogStateDescriptor */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for LogStateDescriptor */

// The size of the internal buffer the log state uses, specified in bytes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLLogStateDescriptor/bufferSize
func (l_ LogStateDescriptor) BufferSize() int {
	rv := objc.Send[int](l_.ID, objc.Sel("bufferSize"))
	return rv
}/* debug [instance_properties/getter]: bufferSize */


// The size of the internal buffer the log state uses, specified in bytes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLLogStateDescriptor/bufferSize
func (l_ LogStateDescriptor) SetBufferSize(value int) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setBufferSize:"), value)
}/* debug [instance_properties/setter]: bufferSize */


// The minimum level of messages that the shader can log.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLLogStateDescriptor/level
func (l_ LogStateDescriptor) Level() LogLevel {
	rv := objc.Send[LogLevel](l_.ID, objc.Sel("level"))
	return rv
}/* debug [instance_properties/getter]: level */


// The minimum level of messages that the shader can log.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLLogStateDescriptor/level
func (l_ LogStateDescriptor) SetLevel(value LogLevel) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setLevel:"), value)
}/* debug [instance_properties/setter]: level */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTLLogStateDescriptor */



