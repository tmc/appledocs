// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [SearchField] class.
var (
	searchFieldClass     _SearchFieldClass
	searchFieldClassOnce sync.Once
)

func getSearchFieldClass() _SearchFieldClass {
	searchFieldClassOnce.Do(func() {
		searchFieldClass = _SearchFieldClass{objc.GetClass("NSSearchField")}
	})
	return searchFieldClass
}

type _SearchFieldClass struct {
	class objc.Class
}

// An interface definition for the [SearchField] class.
type ISearchField interface {
	ITextField
}

// A text field optimized for performing text-based searches. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSearchField

type SearchField struct {
	TextField
}

// SearchFieldFrom constructs a [SearchField] from an unsafe.Pointer.
//
// A text field optimized for performing text-based searches.
func SearchFieldFrom(ptr unsafe.Pointer) SearchField {
	return SearchField{
		TextField: TextFieldFrom(ptr),
	}
}
// Alloc allocates a new instance without initialization.
func (sc _SearchFieldClass) Alloc() SearchField {
	rv := objc.Send[SearchField](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new instance with a +1 retain count.
func (sc _SearchFieldClass) New() SearchField {
	rv := objc.Send[SearchField](objc.ID(sc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (s_ SearchField) Init() SearchField {
	rv := objc.Send[SearchField](s_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s_ SearchField) Autorelease() SearchField {
	rv := objc.Send[SearchField](s_.ID, objc.Sel("autorelease"))
	return rv
}

// NewSearchField creates a new SearchField instance.
func NewSearchField() SearchField {
	return getSearchFieldClass().New()
}




