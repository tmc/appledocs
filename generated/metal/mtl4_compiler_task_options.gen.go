// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTL4CompilerTaskOptions */


/* debug [class_header]: Header for MTL4CompilerTaskOptions */
// The class instance for the [MTL4CompilerTaskOptions] class.
var (
	MTL4CompilerTaskOptionsClass     _MTL4CompilerTaskOptionsClass
	MTL4CompilerTaskOptionsClassOnce sync.Once
)

func getMTL4CompilerTaskOptionsClass() _MTL4CompilerTaskOptionsClass {
	MTL4CompilerTaskOptionsClassOnce.Do(func() {
		MTL4CompilerTaskOptionsClass = _MTL4CompilerTaskOptionsClass{objc.GetClass("MTL4CompilerTaskOptions")}
	})
	return MTL4CompilerTaskOptionsClass
}

type _MTL4CompilerTaskOptionsClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTL4CompilerTaskOptions */
// An interface definition for the [MTL4CompilerTaskOptions] class.
type IMTL4CompilerTaskOptions interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MTL4CompilerTaskOptions */
	// properties:
	LookupArchives() []objc.ID
	SetLookupArchives(value []objc.ID)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTL4CompilerTaskOptions */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTL4CompilerTaskOptions */
// Alloc allocates a new instance without initialization.
func (mc _MTL4CompilerTaskOptionsClass) Alloc() MTL4CompilerTaskOptions {
	rv := objc.Send[MTL4CompilerTaskOptions](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTL4CompilerTaskOptionsClass) New() MTL4CompilerTaskOptions {
	rv := objc.Send[MTL4CompilerTaskOptions](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTL4CompilerTaskOptions) Init() MTL4CompilerTaskOptions {
	rv := objc.Send[MTL4CompilerTaskOptions](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTL4CompilerTaskOptions) Autorelease() MTL4CompilerTaskOptions {
	rv := objc.Send[MTL4CompilerTaskOptions](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTL4CompilerTaskOptions creates a new MTL4CompilerTaskOptions instance.
func NewMTL4CompilerTaskOptions() MTL4CompilerTaskOptions {
	return getMTL4CompilerTaskOptionsClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTL4CompilerTaskOptions */
// The configuration options that control the behavior of a compilation task for a Metal 4 compiler instance.
//
// You can configure task-specific settings that affect a compilation task by creating an instance of this class, setting its properties, and passing it to one of the applicable methods of an instance.


// The configuration options that control the behavior of a compilation task for a Metal 4 compiler instance.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4CompilerTaskOptions
type MTL4CompilerTaskOptions struct {
	objectivec.Object
}

// MTL4CompilerTaskOptionsFrom constructs a [MTL4CompilerTaskOptions] from an unsafe.Pointer.
//
// The configuration options that control the behavior of a compilation task for a Metal 4 compiler instance.
func MTL4CompilerTaskOptionsFrom(ptr unsafe.Pointer) MTL4CompilerTaskOptions {
	return MTL4CompilerTaskOptions{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTL4CompilerTaskOptions *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTL4CompilerTaskOptions */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTL4CompilerTaskOptions */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTL4CompilerTaskOptions */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTL4CompilerTaskOptions */

// An array of archive instances that can potentially accelerate a compilation task.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4CompilerTaskOptions/lookupArchives
func (m_ MTL4CompilerTaskOptions) LookupArchives() []objc.ID {
	rv := objc.Send[[]objc.ID](m_.ID, objc.Sel("lookupArchives"))
	return rv
}/* debug [instance_properties/getter]: lookupArchives */


// An array of archive instances that can potentially accelerate a compilation task.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4CompilerTaskOptions/lookupArchives
func (m_ MTL4CompilerTaskOptions) SetLookupArchives(value []objc.ID) {
	var nsArray objc.ID
	if len(value) > 0 {
		nsArray = objc.ID(objc.GetClass("NSMutableArray")).Send(objc.Sel("arrayWithCapacity:"), len(value))
		for _, item := range value {
			nsArray.Send(objc.Sel("addObject:"), item)
		}
	} else {
		nsArray = objc.ID(objc.GetClass("NSArray")).Send(objc.Sel("array"))
	}
	objc.Send[objc.ID](m_.ID, objc.Sel("setLookupArchives:"), nsArray)
}/* debug [instance_properties/setter]: lookupArchives */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTL4CompilerTaskOptions */



