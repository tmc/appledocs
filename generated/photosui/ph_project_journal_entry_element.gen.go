// Code generated from Apple documentation for PhotosUI. DO NOT EDIT.

package photosui

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [PHProjectJournalEntryElement] class.
var (
	PHProjectJournalEntryElementClass     _PHProjectJournalEntryElementClass
	PHProjectJournalEntryElementClassOnce sync.Once
)

func getPHProjectJournalEntryElementClass() _PHProjectJournalEntryElementClass {
	PHProjectJournalEntryElementClassOnce.Do(func() {
		PHProjectJournalEntryElementClass = _PHProjectJournalEntryElementClass{objc.GetClass("PHProjectJournalEntryElement")}
	})
	return PHProjectJournalEntryElementClass
}

type _PHProjectJournalEntryElementClass struct {
	class objc.Class
}

// An interface definition for the [PHProjectJournalEntryElement] class.
type IPHProjectJournalEntryElement interface {
	IPHProjectElement
	AssetElement() PHProjectAssetElement
	Date() foundation.NSDate
	TextElement() PHProjectTextElement
}

// An element that represents a journal entry within project section content.
//
// [Full Topic]: https://developer.apple.com/documentation/PhotosUI/PHProjectJournalEntryElement
type PHProjectJournalEntryElement struct {
	PHProjectElement
}

// PHProjectJournalEntryElementFrom constructs a [PHProjectJournalEntryElement] from an unsafe.Pointer.
//
// An element that represents a journal entry within project section content.
func PHProjectJournalEntryElementFrom(ptr unsafe.Pointer) PHProjectJournalEntryElement {
	return PHProjectJournalEntryElement{
		PHProjectElement: PHProjectElementFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (pc _PHProjectJournalEntryElementClass) Alloc() PHProjectJournalEntryElement {
	rv := objc.Send[PHProjectJournalEntryElement](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (pc _PHProjectJournalEntryElementClass) New() PHProjectJournalEntryElement {
	rv := objc.Send[PHProjectJournalEntryElement](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ PHProjectJournalEntryElement) Init() PHProjectJournalEntryElement {
	rv := objc.Send[PHProjectJournalEntryElement](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ PHProjectJournalEntryElement) Autorelease() PHProjectJournalEntryElement {
	rv := objc.Send[PHProjectJournalEntryElement](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewPHProjectJournalEntryElement creates a new PHProjectJournalEntryElement instance.
func NewPHProjectJournalEntryElement() PHProjectJournalEntryElement {
	return getPHProjectJournalEntryElementClass().New()
}


// An optional asset to represent the date in the journal entry.
//
// [Full Topic]: https://developer.apple.com/documentation/PhotosUI/PHProjectJournalEntryElement/assetElement
func (p_ PHProjectJournalEntryElement) AssetElement() PHProjectAssetElement {
	rv := objc.Send[PHProjectAssetElement](p_.ID, objc.Sel("assetElement"))
	return rv
}

// The date of the journal entry.
//
// [Full Topic]: https://developer.apple.com/documentation/PhotosUI/PHProjectJournalEntryElement/date
func (p_ PHProjectJournalEntryElement) Date() foundation.NSDate {
	rv := objc.Send[foundation.NSDate](p_.ID, objc.Sel("date"))
	return rv
}

// Descriptive text for the date of the entry.
//
// [Full Topic]: https://developer.apple.com/documentation/PhotosUI/PHProjectJournalEntryElement/textElement
func (p_ PHProjectJournalEntryElement) TextElement() PHProjectTextElement {
	rv := objc.Send[PHProjectTextElement](p_.ID, objc.Sel("textElement"))
	return rv
}



