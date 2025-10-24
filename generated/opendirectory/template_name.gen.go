// Code generated from Apple documentation for OpenDirectory. DO NOT EDIT.

package opendirectory

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [templateName] class.
var (
	TemplateNameClass     _templateNameClass
	TemplateNameClassOnce sync.Once
)

func gettemplateNameClass() _templateNameClass {
	TemplateNameClassOnce.Do(func() {
		TemplateNameClass = _templateNameClass{objc.GetClass("templateName")}
	})
	return TemplateNameClass
}

type _templateNameClass struct {
	class objc.Class
}

// An interface definition for the [templateName] class.
type ItemplateName interface {
	objectivec.IObject
	// properties:
	// methods:
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODConfiguration/templateName-c.ivar
type templateName struct {
	objectivec.Object
}

// templateNameFrom constructs a [templateName] from an unsafe.Pointer.
func templateNameFrom(ptr unsafe.Pointer) templateName {
	return templateName{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (tc _templateNameClass) Alloc() templateName {
	rv := objc.Send[templateName](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (tc _templateNameClass) New() templateName {
	rv := objc.Send[templateName](objc.ID(tc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (t_ templateName) Init() templateName {
	rv := objc.Send[templateName](t_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t_ templateName) Autorelease() templateName {
	rv := objc.Send[templateName](t_.ID, objc.Sel("autorelease"))
	return rv
}

// NewtemplateName creates a new templateName instance.
func NewtemplateName() templateName {
	return gettemplateNameClass().New()
}




