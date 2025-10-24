// Code generated from Apple documentation for OpenDirectory. DO NOT EDIT.

package opendirectory

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

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

// An interface definition for the [comment] class.
type Icomment interface {
	objectivec.IObject
	// properties:
	// methods:
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODConfiguration/comment-c.ivar
type comment struct {
	objectivec.Object
}

// commentFrom constructs a [comment] from an unsafe.Pointer.
func commentFrom(ptr unsafe.Pointer) comment {
	return comment{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (cc _commentClass) Alloc() comment {
	rv := objc.Send[comment](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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




