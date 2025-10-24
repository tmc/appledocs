// Code generated from Apple documentation for OpenDirectory. DO NOT EDIT.

package opendirectory

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class comment */


/* debug [class_header]: Header for comment */
// The class instance for the [comment] class.
var (
	CommentClass     _commentClass
	CommentClassOnce sync.Once
)

func getcommentClass() _commentClass {
	CommentClassOnce.Do(func() {
		CommentClass = _commentClass{objc.GetClass("comment")}
	})
	return CommentClass
}

type _commentClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for comment */
// An interface definition for the [comment] class.
type Icomment interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for comment */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for comment */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for comment */
// Alloc allocates a new instance without initialization.
func (cc _commentClass) Alloc() comment {
	rv := objc.Send[comment](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _commentClass) New() comment {
	rv := objc.Send[comment](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ comment) Init() comment {
	rv := objc.Send[comment](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ comment) Autorelease() comment {
	rv := objc.Send[comment](c_.ID, objc.Sel("autorelease"))
	return rv
}

// Newcomment creates a new comment instance.
func Newcomment() comment {
	return getcommentClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for comment */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODConfiguration/comment-c.ivar
type comment struct {
	objectivec.Object
}

// commentFrom constructs a [comment] from an unsafe.Pointer.
func commentFrom(ptr unsafe.Pointer) comment {
	return comment{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for comment *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for comment */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for comment */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for comment */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for comment */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class comment */



