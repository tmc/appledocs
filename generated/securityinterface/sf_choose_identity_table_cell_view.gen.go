// Code generated from Apple documentation for SecurityInterface. DO NOT EDIT.

package securityinterface

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/appkit"
)

// The class instance for the [SFChooseIdentityTableCellView] class.
var (
	SFChooseIdentityTableCellViewClass     _SFChooseIdentityTableCellViewClass
	SFChooseIdentityTableCellViewClassOnce sync.Once
)

func getSFChooseIdentityTableCellViewClass() _SFChooseIdentityTableCellViewClass {
	SFChooseIdentityTableCellViewClassOnce.Do(func() {
		SFChooseIdentityTableCellViewClass = _SFChooseIdentityTableCellViewClass{objc.GetClass("SFChooseIdentityTableCellView")}
	})
	return SFChooseIdentityTableCellViewClass
}

type _SFChooseIdentityTableCellViewClass struct {
	class objc.Class
}

// An interface definition for the [SFChooseIdentityTableCellView] class.
type ISFChooseIdentityTableCellView interface {
	appkit.ITableCellView
	IssuerTextField() appkit.TextField
	SetIssuerTextField(value appkit.ITextField)
}

//
// [Full Topic]: https://developer.apple.com/documentation/SecurityInterface/SFChooseIdentityTableCellView
type SFChooseIdentityTableCellView struct {
	appkit.TableCellView
}

// SFChooseIdentityTableCellViewFrom constructs a [SFChooseIdentityTableCellView] from an unsafe.Pointer.
func SFChooseIdentityTableCellViewFrom(ptr unsafe.Pointer) SFChooseIdentityTableCellView {
	return SFChooseIdentityTableCellView{
		TableCellView: appkit.TableCellViewFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (sc _SFChooseIdentityTableCellViewClass) Alloc() SFChooseIdentityTableCellView {
	rv := objc.Send[SFChooseIdentityTableCellView](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (sc _SFChooseIdentityTableCellViewClass) New() SFChooseIdentityTableCellView {
	rv := objc.Send[SFChooseIdentityTableCellView](objc.ID(sc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (s_ SFChooseIdentityTableCellView) Init() SFChooseIdentityTableCellView {
	rv := objc.Send[SFChooseIdentityTableCellView](s_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s_ SFChooseIdentityTableCellView) Autorelease() SFChooseIdentityTableCellView {
	rv := objc.Send[SFChooseIdentityTableCellView](s_.ID, objc.Sel("autorelease"))
	return rv
}

// NewSFChooseIdentityTableCellView creates a new SFChooseIdentityTableCellView instance.
func NewSFChooseIdentityTableCellView() SFChooseIdentityTableCellView {
	return getSFChooseIdentityTableCellViewClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/SecurityInterface/SFChooseIdentityTableCellView/issuerTextField-swift.property
func (s_ SFChooseIdentityTableCellView) IssuerTextField() appkit.TextField {
	rv := objc.Send[appkit.TextField](s_.ID, objc.Sel("issuerTextField"))
	return rv
}


// SetIssuerTextField sets the value of the issuerTextField property.
//
// [Full Topic]: https://developer.apple.com/documentation/SecurityInterface/SFChooseIdentityTableCellView/issuerTextField-swift.property
func (s_ SFChooseIdentityTableCellView) SetIssuerTextField(value appkit.ITextField) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setIssuerTextField:"), value)
}



