// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (

	"github.com/tmc/appledocs/generated/objc"

	"github.com/tmc/appledocs/generated/corefoundation"

	"github.com/tmc/appledocs/generated/foundation"

	"github.com/tmc/appledocs/generated/vision"
)

// PAccessibility is the NSAccessibility protocol interface.
//
// The complete list of properties and methods for accessible elements.
//
// Availability:
//   - macOS +
//
// See: doc://com.apple.appkit/documentation/AppKit/NSAccessibilityProtocol
type PAccessibility interface {
	// Required methods
	AccessibilityActivationPoint()/* debug [protocol_interface/required_method]: AccessibilityActivationPoint */
	AccessibilityAllowedValues()/* debug [protocol_interface/required_method]: AccessibilityAllowedValues */
	AccessibilityApplicationFocusedUIElement()/* debug [protocol_interface/required_method]: AccessibilityApplicationFocusedUIElement */
	AccessibilityAttributedStringForRange(range_ corefoundation.Range) foundation.AttributedString/* debug [protocol_interface/required_method]: AccessibilityAttributedStringForRange */
	AccessibilityAttributedUserInputLabels()/* debug [protocol_interface/required_method]: AccessibilityAttributedUserInputLabels */
	AccessibilityCancelButton()/* debug [protocol_interface/required_method]: AccessibilityCancelButton */
	AccessibilityCellForColumnRow(column int, row int) objc.ID/* debug [protocol_interface/required_method]: AccessibilityCellForColumnRow */
	AccessibilityChildren()/* debug [protocol_interface/required_method]: AccessibilityChildren */
	AccessibilityChildrenInNavigationOrder()/* debug [protocol_interface/required_method]: AccessibilityChildrenInNavigationOrder */
	AccessibilityClearButton()/* debug [protocol_interface/required_method]: AccessibilityClearButton */
	AccessibilityCloseButton()/* debug [protocol_interface/required_method]: AccessibilityCloseButton */
	AccessibilityColumnCount()/* debug [protocol_interface/required_method]: AccessibilityColumnCount */
	AccessibilityColumnHeaderUIElements()/* debug [protocol_interface/required_method]: AccessibilityColumnHeaderUIElements */
	AccessibilityColumnIndexRange()/* debug [protocol_interface/required_method]: AccessibilityColumnIndexRange */
	AccessibilityColumnTitles()/* debug [protocol_interface/required_method]: AccessibilityColumnTitles */
	AccessibilityColumns()/* debug [protocol_interface/required_method]: AccessibilityColumns */
	AccessibilityContents()/* debug [protocol_interface/required_method]: AccessibilityContents */
	AccessibilityCriticalValue()/* debug [protocol_interface/required_method]: AccessibilityCriticalValue */
	AccessibilityCustomActions()/* debug [protocol_interface/required_method]: AccessibilityCustomActions */
	AccessibilityCustomRotors()/* debug [protocol_interface/required_method]: AccessibilityCustomRotors */
	AccessibilityDecrementButton()/* debug [protocol_interface/required_method]: AccessibilityDecrementButton */
	AccessibilityDefaultButton()/* debug [protocol_interface/required_method]: AccessibilityDefaultButton */
	AccessibilityDisclosedByRow()/* debug [protocol_interface/required_method]: AccessibilityDisclosedByRow */
	AccessibilityDisclosedRows()/* debug [protocol_interface/required_method]: AccessibilityDisclosedRows */
	AccessibilityDisclosureLevel()/* debug [protocol_interface/required_method]: AccessibilityDisclosureLevel */
	AccessibilityDocument()/* debug [protocol_interface/required_method]: AccessibilityDocument */
	AccessibilityExtrasMenuBar()/* debug [protocol_interface/required_method]: AccessibilityExtrasMenuBar */
	AccessibilityFilename()/* debug [protocol_interface/required_method]: AccessibilityFilename */
	AccessibilityFocusedWindow()/* debug [protocol_interface/required_method]: AccessibilityFocusedWindow */
	AccessibilityFrame()/* debug [protocol_interface/required_method]: AccessibilityFrame */
	AccessibilityFrameForRange(range_ corefoundation.Range) Rect/* debug [protocol_interface/required_method]: AccessibilityFrameForRange */
	AccessibilityFullScreenButton()/* debug [protocol_interface/required_method]: AccessibilityFullScreenButton */
	AccessibilityGrowArea()/* debug [protocol_interface/required_method]: AccessibilityGrowArea */
	AccessibilityHandles()/* debug [protocol_interface/required_method]: AccessibilityHandles */
	AccessibilityHeader()/* debug [protocol_interface/required_method]: AccessibilityHeader */
	AccessibilityHelp()/* debug [protocol_interface/required_method]: AccessibilityHelp */
	AccessibilityHorizontalScrollBar()/* debug [protocol_interface/required_method]: AccessibilityHorizontalScrollBar */
	AccessibilityHorizontalUnitDescription()/* debug [protocol_interface/required_method]: AccessibilityHorizontalUnitDescription */
	AccessibilityHorizontalUnits()/* debug [protocol_interface/required_method]: AccessibilityHorizontalUnits */
	AccessibilityIdentifier()/* debug [protocol_interface/required_method]: AccessibilityIdentifier */
	AccessibilityIncrementButton()/* debug [protocol_interface/required_method]: AccessibilityIncrementButton */
	AccessibilityIndex()/* debug [protocol_interface/required_method]: AccessibilityIndex */
	AccessibilityInsertionPointLineNumber()/* debug [protocol_interface/required_method]: AccessibilityInsertionPointLineNumber */
	AccessibilityLabel()/* debug [protocol_interface/required_method]: AccessibilityLabel */
	AccessibilityLabelUIElements()/* debug [protocol_interface/required_method]: AccessibilityLabelUIElements */
	AccessibilityLabelValue()/* debug [protocol_interface/required_method]: AccessibilityLabelValue */
	AccessibilityLayoutPointForScreenPoint(point vision.Point) vision.Point/* debug [protocol_interface/required_method]: AccessibilityLayoutPointForScreenPoint */
	AccessibilityLayoutSizeForScreenSize(size Size /* not a class type */) Size/* debug [protocol_interface/required_method]: AccessibilityLayoutSizeForScreenSize */
	AccessibilityLineForIndex(index int) int/* debug [protocol_interface/required_method]: AccessibilityLineForIndex */
	AccessibilityLinkedUIElements()/* debug [protocol_interface/required_method]: AccessibilityLinkedUIElements */
	AccessibilityMainWindow()/* debug [protocol_interface/required_method]: AccessibilityMainWindow */
	AccessibilityMarkerGroupUIElement()/* debug [protocol_interface/required_method]: AccessibilityMarkerGroupUIElement */
	AccessibilityMarkerTypeDescription()/* debug [protocol_interface/required_method]: AccessibilityMarkerTypeDescription */
	AccessibilityMarkerUIElements()/* debug [protocol_interface/required_method]: AccessibilityMarkerUIElements */
	AccessibilityMarkerValues()/* debug [protocol_interface/required_method]: AccessibilityMarkerValues */
	AccessibilityMaxValue()/* debug [protocol_interface/required_method]: AccessibilityMaxValue */
	AccessibilityMenuBar()/* debug [protocol_interface/required_method]: AccessibilityMenuBar */
	AccessibilityMinValue()/* debug [protocol_interface/required_method]: AccessibilityMinValue */
	AccessibilityMinimizeButton()/* debug [protocol_interface/required_method]: AccessibilityMinimizeButton */
	AccessibilityNextContents()/* debug [protocol_interface/required_method]: AccessibilityNextContents */
	AccessibilityNumberOfCharacters()/* debug [protocol_interface/required_method]: AccessibilityNumberOfCharacters */
	AccessibilityOrientation()/* debug [protocol_interface/required_method]: AccessibilityOrientation */
	AccessibilityOverflowButton()/* debug [protocol_interface/required_method]: AccessibilityOverflowButton */
	AccessibilityParent()/* debug [protocol_interface/required_method]: AccessibilityParent */
	AccessibilityPerformCancel() bool/* debug [protocol_interface/required_method]: AccessibilityPerformCancel */
	AccessibilityPerformConfirm() bool/* debug [protocol_interface/required_method]: AccessibilityPerformConfirm */
	AccessibilityPerformDecrement() bool/* debug [protocol_interface/required_method]: AccessibilityPerformDecrement */
	AccessibilityPerformDelete() bool/* debug [protocol_interface/required_method]: AccessibilityPerformDelete */
	AccessibilityPerformIncrement() bool/* debug [protocol_interface/required_method]: AccessibilityPerformIncrement */
	AccessibilityPerformPick() bool/* debug [protocol_interface/required_method]: AccessibilityPerformPick */
	AccessibilityPerformPress() bool/* debug [protocol_interface/required_method]: AccessibilityPerformPress */
	AccessibilityPerformRaise() bool/* debug [protocol_interface/required_method]: AccessibilityPerformRaise */
	AccessibilityPerformShowAlternateUI() bool/* debug [protocol_interface/required_method]: AccessibilityPerformShowAlternateUI */
	AccessibilityPerformShowDefaultUI() bool/* debug [protocol_interface/required_method]: AccessibilityPerformShowDefaultUI */
	AccessibilityPerformShowMenu() bool/* debug [protocol_interface/required_method]: AccessibilityPerformShowMenu */
	AccessibilityPlaceholderValue()/* debug [protocol_interface/required_method]: AccessibilityPlaceholderValue */
	AccessibilityPreviousContents()/* debug [protocol_interface/required_method]: AccessibilityPreviousContents */
	AccessibilityProxy()/* debug [protocol_interface/required_method]: AccessibilityProxy */
	AccessibilityRTFForRange(range_ corefoundation.Range) foundation.Data/* debug [protocol_interface/required_method]: AccessibilityRTFForRange */
	AccessibilityRangeForPosition(point vision.Point) corefoundation.Range/* debug [protocol_interface/required_method]: AccessibilityRangeForPosition */
	AccessibilityRangeForIndex(index int) corefoundation.Range/* debug [protocol_interface/required_method]: AccessibilityRangeForIndex */
	AccessibilityRangeForLine(line int) corefoundation.Range/* debug [protocol_interface/required_method]: AccessibilityRangeForLine */
	AccessibilityRole()/* debug [protocol_interface/required_method]: AccessibilityRole */
	AccessibilityRoleDescription()/* debug [protocol_interface/required_method]: AccessibilityRoleDescription */
	AccessibilityRowCount()/* debug [protocol_interface/required_method]: AccessibilityRowCount */
	AccessibilityRowHeaderUIElements()/* debug [protocol_interface/required_method]: AccessibilityRowHeaderUIElements */
	AccessibilityRowIndexRange()/* debug [protocol_interface/required_method]: AccessibilityRowIndexRange */
	AccessibilityRows()/* debug [protocol_interface/required_method]: AccessibilityRows */
	AccessibilityRulerMarkerType()/* debug [protocol_interface/required_method]: AccessibilityRulerMarkerType */
	AccessibilityScreenPointForLayoutPoint(point vision.Point) vision.Point/* debug [protocol_interface/required_method]: AccessibilityScreenPointForLayoutPoint */
	AccessibilityScreenSizeForLayoutSize(size Size /* not a class type */) Size/* debug [protocol_interface/required_method]: AccessibilityScreenSizeForLayoutSize */
	AccessibilitySearchButton()/* debug [protocol_interface/required_method]: AccessibilitySearchButton */
	AccessibilitySearchMenu()/* debug [protocol_interface/required_method]: AccessibilitySearchMenu */
	AccessibilitySelectedCells()/* debug [protocol_interface/required_method]: AccessibilitySelectedCells */
	AccessibilitySelectedChildren()/* debug [protocol_interface/required_method]: AccessibilitySelectedChildren */
	AccessibilitySelectedColumns()/* debug [protocol_interface/required_method]: AccessibilitySelectedColumns */
	AccessibilitySelectedRows()/* debug [protocol_interface/required_method]: AccessibilitySelectedRows */
	AccessibilitySelectedText()/* debug [protocol_interface/required_method]: AccessibilitySelectedText */
	AccessibilitySelectedTextRange()/* debug [protocol_interface/required_method]: AccessibilitySelectedTextRange */
	AccessibilitySelectedTextRanges()/* debug [protocol_interface/required_method]: AccessibilitySelectedTextRanges */
	AccessibilityServesAsTitleForUIElements()/* debug [protocol_interface/required_method]: AccessibilityServesAsTitleForUIElements */
	AccessibilitySharedCharacterRange()/* debug [protocol_interface/required_method]: AccessibilitySharedCharacterRange */
	AccessibilitySharedFocusElements()/* debug [protocol_interface/required_method]: AccessibilitySharedFocusElements */
	AccessibilitySharedTextUIElements()/* debug [protocol_interface/required_method]: AccessibilitySharedTextUIElements */
	AccessibilityShownMenu()/* debug [protocol_interface/required_method]: AccessibilityShownMenu */
	AccessibilitySortDirection()/* debug [protocol_interface/required_method]: AccessibilitySortDirection */
	AccessibilitySplitters()/* debug [protocol_interface/required_method]: AccessibilitySplitters */
	AccessibilityStringForRange(range_ corefoundation.Range) foundation.String/* debug [protocol_interface/required_method]: AccessibilityStringForRange */
	AccessibilityStyleRangeForIndex(index int) corefoundation.Range/* debug [protocol_interface/required_method]: AccessibilityStyleRangeForIndex */
	AccessibilitySubrole()/* debug [protocol_interface/required_method]: AccessibilitySubrole */
	AccessibilityTabs()/* debug [protocol_interface/required_method]: AccessibilityTabs */
	AccessibilityTitle()/* debug [protocol_interface/required_method]: AccessibilityTitle */
	AccessibilityTitleUIElement()/* debug [protocol_interface/required_method]: AccessibilityTitleUIElement */
	AccessibilityToolbarButton()/* debug [protocol_interface/required_method]: AccessibilityToolbarButton */
	AccessibilityTopLevelUIElement()/* debug [protocol_interface/required_method]: AccessibilityTopLevelUIElement */
	AccessibilityURL()/* debug [protocol_interface/required_method]: AccessibilityURL */
	AccessibilityUnitDescription()/* debug [protocol_interface/required_method]: AccessibilityUnitDescription */
	AccessibilityUnits()/* debug [protocol_interface/required_method]: AccessibilityUnits */
	AccessibilityUserInputLabels()/* debug [protocol_interface/required_method]: AccessibilityUserInputLabels */
	AccessibilityValue()/* debug [protocol_interface/required_method]: AccessibilityValue */
	AccessibilityValueDescription()/* debug [protocol_interface/required_method]: AccessibilityValueDescription */
	AccessibilityVerticalScrollBar()/* debug [protocol_interface/required_method]: AccessibilityVerticalScrollBar */
	AccessibilityVerticalUnitDescription()/* debug [protocol_interface/required_method]: AccessibilityVerticalUnitDescription */
	AccessibilityVerticalUnits()/* debug [protocol_interface/required_method]: AccessibilityVerticalUnits */
	AccessibilityVisibleCells()/* debug [protocol_interface/required_method]: AccessibilityVisibleCells */
	AccessibilityVisibleCharacterRange()/* debug [protocol_interface/required_method]: AccessibilityVisibleCharacterRange */
	AccessibilityVisibleChildren()/* debug [protocol_interface/required_method]: AccessibilityVisibleChildren */
	AccessibilityVisibleColumns()/* debug [protocol_interface/required_method]: AccessibilityVisibleColumns */
	AccessibilityVisibleRows()/* debug [protocol_interface/required_method]: AccessibilityVisibleRows */
	AccessibilityWarningValue()/* debug [protocol_interface/required_method]: AccessibilityWarningValue */
	AccessibilityWindow()/* debug [protocol_interface/required_method]: AccessibilityWindow */
	AccessibilityWindows()/* debug [protocol_interface/required_method]: AccessibilityWindows */
	AccessibilityZoomButton()/* debug [protocol_interface/required_method]: AccessibilityZoomButton */
	IsAccessibilityAlternateUIVisible()/* debug [protocol_interface/required_method]: IsAccessibilityAlternateUIVisible */
	IsAccessibilityDisclosed()/* debug [protocol_interface/required_method]: IsAccessibilityDisclosed */
	IsAccessibilityEdited()/* debug [protocol_interface/required_method]: IsAccessibilityEdited */
	IsAccessibilityElement()/* debug [protocol_interface/required_method]: IsAccessibilityElement */
	IsAccessibilityEnabled()/* debug [protocol_interface/required_method]: IsAccessibilityEnabled */
	IsAccessibilityExpanded()/* debug [protocol_interface/required_method]: IsAccessibilityExpanded */
	IsAccessibilityFocused()/* debug [protocol_interface/required_method]: IsAccessibilityFocused */
	IsAccessibilityFrontmost()/* debug [protocol_interface/required_method]: IsAccessibilityFrontmost */
	IsAccessibilityHidden()/* debug [protocol_interface/required_method]: IsAccessibilityHidden */
	IsAccessibilityMain()/* debug [protocol_interface/required_method]: IsAccessibilityMain */
	IsAccessibilityMinimized()/* debug [protocol_interface/required_method]: IsAccessibilityMinimized */
	IsAccessibilityModal()/* debug [protocol_interface/required_method]: IsAccessibilityModal */
	IsAccessibilityOrderedByRow()/* debug [protocol_interface/required_method]: IsAccessibilityOrderedByRow */
	IsAccessibilityProtectedContent()/* debug [protocol_interface/required_method]: IsAccessibilityProtectedContent */
	IsAccessibilityRequired()/* debug [protocol_interface/required_method]: IsAccessibilityRequired */
	IsAccessibilitySelected()/* debug [protocol_interface/required_method]: IsAccessibilitySelected */
	IsAccessibilitySelectorAllowed(selector objc.SEL) bool/* debug [protocol_interface/required_method]: IsAccessibilitySelectorAllowed */
	SetAccessibilityActivationPoint()/* debug [protocol_interface/required_method]: SetAccessibilityActivationPoint */
	SetAccessibilityAllowedValues()/* debug [protocol_interface/required_method]: SetAccessibilityAllowedValues */
	SetAccessibilityAlternateUIVisible()/* debug [protocol_interface/required_method]: SetAccessibilityAlternateUIVisible */
	SetAccessibilityApplicationFocusedUIElement()/* debug [protocol_interface/required_method]: SetAccessibilityApplicationFocusedUIElement */
	SetAccessibilityAttributedUserInputLabels()/* debug [protocol_interface/required_method]: SetAccessibilityAttributedUserInputLabels */
	SetAccessibilityCancelButton()/* debug [protocol_interface/required_method]: SetAccessibilityCancelButton */
	SetAccessibilityChildren()/* debug [protocol_interface/required_method]: SetAccessibilityChildren */
	SetAccessibilityChildrenInNavigationOrder()/* debug [protocol_interface/required_method]: SetAccessibilityChildrenInNavigationOrder */
	SetAccessibilityClearButton()/* debug [protocol_interface/required_method]: SetAccessibilityClearButton */
	SetAccessibilityCloseButton()/* debug [protocol_interface/required_method]: SetAccessibilityCloseButton */
	SetAccessibilityColumnCount()/* debug [protocol_interface/required_method]: SetAccessibilityColumnCount */
	SetAccessibilityColumnHeaderUIElements()/* debug [protocol_interface/required_method]: SetAccessibilityColumnHeaderUIElements */
	SetAccessibilityColumnIndexRange()/* debug [protocol_interface/required_method]: SetAccessibilityColumnIndexRange */
	SetAccessibilityColumnTitles()/* debug [protocol_interface/required_method]: SetAccessibilityColumnTitles */
	SetAccessibilityColumns()/* debug [protocol_interface/required_method]: SetAccessibilityColumns */
	SetAccessibilityContents()/* debug [protocol_interface/required_method]: SetAccessibilityContents */
	SetAccessibilityCriticalValue()/* debug [protocol_interface/required_method]: SetAccessibilityCriticalValue */
	SetAccessibilityCustomActions()/* debug [protocol_interface/required_method]: SetAccessibilityCustomActions */
	SetAccessibilityCustomRotors()/* debug [protocol_interface/required_method]: SetAccessibilityCustomRotors */
	SetAccessibilityDecrementButton()/* debug [protocol_interface/required_method]: SetAccessibilityDecrementButton */
	SetAccessibilityDefaultButton()/* debug [protocol_interface/required_method]: SetAccessibilityDefaultButton */
	SetAccessibilityDisclosed()/* debug [protocol_interface/required_method]: SetAccessibilityDisclosed */
	SetAccessibilityDisclosedByRow()/* debug [protocol_interface/required_method]: SetAccessibilityDisclosedByRow */
	SetAccessibilityDisclosedRows()/* debug [protocol_interface/required_method]: SetAccessibilityDisclosedRows */
	SetAccessibilityDisclosureLevel()/* debug [protocol_interface/required_method]: SetAccessibilityDisclosureLevel */
	SetAccessibilityDocument()/* debug [protocol_interface/required_method]: SetAccessibilityDocument */
	SetAccessibilityEdited()/* debug [protocol_interface/required_method]: SetAccessibilityEdited */
	SetAccessibilityElement()/* debug [protocol_interface/required_method]: SetAccessibilityElement */
	SetAccessibilityEnabled()/* debug [protocol_interface/required_method]: SetAccessibilityEnabled */
	SetAccessibilityExpanded()/* debug [protocol_interface/required_method]: SetAccessibilityExpanded */
	SetAccessibilityExtrasMenuBar()/* debug [protocol_interface/required_method]: SetAccessibilityExtrasMenuBar */
	SetAccessibilityFilename()/* debug [protocol_interface/required_method]: SetAccessibilityFilename */
	SetAccessibilityFocused()/* debug [protocol_interface/required_method]: SetAccessibilityFocused */
	SetAccessibilityFocusedWindow()/* debug [protocol_interface/required_method]: SetAccessibilityFocusedWindow */
	SetAccessibilityFrame()/* debug [protocol_interface/required_method]: SetAccessibilityFrame */
	SetAccessibilityFrontmost()/* debug [protocol_interface/required_method]: SetAccessibilityFrontmost */
	SetAccessibilityFullScreenButton()/* debug [protocol_interface/required_method]: SetAccessibilityFullScreenButton */
	SetAccessibilityGrowArea()/* debug [protocol_interface/required_method]: SetAccessibilityGrowArea */
	SetAccessibilityHandles()/* debug [protocol_interface/required_method]: SetAccessibilityHandles */
	SetAccessibilityHeader()/* debug [protocol_interface/required_method]: SetAccessibilityHeader */
	SetAccessibilityHelp()/* debug [protocol_interface/required_method]: SetAccessibilityHelp */
	SetAccessibilityHidden()/* debug [protocol_interface/required_method]: SetAccessibilityHidden */
	SetAccessibilityHorizontalScrollBar()/* debug [protocol_interface/required_method]: SetAccessibilityHorizontalScrollBar */
	SetAccessibilityHorizontalUnitDescription()/* debug [protocol_interface/required_method]: SetAccessibilityHorizontalUnitDescription */
	SetAccessibilityHorizontalUnits()/* debug [protocol_interface/required_method]: SetAccessibilityHorizontalUnits */
	SetAccessibilityIdentifier()/* debug [protocol_interface/required_method]: SetAccessibilityIdentifier */
	SetAccessibilityIncrementButton()/* debug [protocol_interface/required_method]: SetAccessibilityIncrementButton */
	SetAccessibilityIndex()/* debug [protocol_interface/required_method]: SetAccessibilityIndex */
	SetAccessibilityInsertionPointLineNumber()/* debug [protocol_interface/required_method]: SetAccessibilityInsertionPointLineNumber */
	SetAccessibilityLabel()/* debug [protocol_interface/required_method]: SetAccessibilityLabel */
	SetAccessibilityLabelUIElements()/* debug [protocol_interface/required_method]: SetAccessibilityLabelUIElements */
	SetAccessibilityLabelValue()/* debug [protocol_interface/required_method]: SetAccessibilityLabelValue */
	SetAccessibilityLinkedUIElements()/* debug [protocol_interface/required_method]: SetAccessibilityLinkedUIElements */
	SetAccessibilityMain()/* debug [protocol_interface/required_method]: SetAccessibilityMain */
	SetAccessibilityMainWindow()/* debug [protocol_interface/required_method]: SetAccessibilityMainWindow */
	SetAccessibilityMarkerGroupUIElement()/* debug [protocol_interface/required_method]: SetAccessibilityMarkerGroupUIElement */
	SetAccessibilityMarkerTypeDescription()/* debug [protocol_interface/required_method]: SetAccessibilityMarkerTypeDescription */
	SetAccessibilityMarkerUIElements()/* debug [protocol_interface/required_method]: SetAccessibilityMarkerUIElements */
	SetAccessibilityMarkerValues()/* debug [protocol_interface/required_method]: SetAccessibilityMarkerValues */
	SetAccessibilityMaxValue()/* debug [protocol_interface/required_method]: SetAccessibilityMaxValue */
	SetAccessibilityMenuBar()/* debug [protocol_interface/required_method]: SetAccessibilityMenuBar */
	SetAccessibilityMinValue()/* debug [protocol_interface/required_method]: SetAccessibilityMinValue */
	SetAccessibilityMinimizeButton()/* debug [protocol_interface/required_method]: SetAccessibilityMinimizeButton */
	SetAccessibilityMinimized()/* debug [protocol_interface/required_method]: SetAccessibilityMinimized */
	SetAccessibilityModal()/* debug [protocol_interface/required_method]: SetAccessibilityModal */
	SetAccessibilityNextContents()/* debug [protocol_interface/required_method]: SetAccessibilityNextContents */
	SetAccessibilityNumberOfCharacters()/* debug [protocol_interface/required_method]: SetAccessibilityNumberOfCharacters */
	SetAccessibilityOrderedByRow()/* debug [protocol_interface/required_method]: SetAccessibilityOrderedByRow */
	SetAccessibilityOrientation()/* debug [protocol_interface/required_method]: SetAccessibilityOrientation */
	SetAccessibilityOverflowButton()/* debug [protocol_interface/required_method]: SetAccessibilityOverflowButton */
	SetAccessibilityParent()/* debug [protocol_interface/required_method]: SetAccessibilityParent */
	SetAccessibilityPlaceholderValue()/* debug [protocol_interface/required_method]: SetAccessibilityPlaceholderValue */
	SetAccessibilityPreviousContents()/* debug [protocol_interface/required_method]: SetAccessibilityPreviousContents */
	SetAccessibilityProtectedContent()/* debug [protocol_interface/required_method]: SetAccessibilityProtectedContent */
	SetAccessibilityProxy()/* debug [protocol_interface/required_method]: SetAccessibilityProxy */
	SetAccessibilityRequired()/* debug [protocol_interface/required_method]: SetAccessibilityRequired */
	SetAccessibilityRole()/* debug [protocol_interface/required_method]: SetAccessibilityRole */
	SetAccessibilityRoleDescription()/* debug [protocol_interface/required_method]: SetAccessibilityRoleDescription */
	SetAccessibilityRowCount()/* debug [protocol_interface/required_method]: SetAccessibilityRowCount */
	SetAccessibilityRowHeaderUIElements()/* debug [protocol_interface/required_method]: SetAccessibilityRowHeaderUIElements */
	SetAccessibilityRowIndexRange()/* debug [protocol_interface/required_method]: SetAccessibilityRowIndexRange */
	SetAccessibilityRows()/* debug [protocol_interface/required_method]: SetAccessibilityRows */
	SetAccessibilityRulerMarkerType()/* debug [protocol_interface/required_method]: SetAccessibilityRulerMarkerType */
	SetAccessibilitySearchButton()/* debug [protocol_interface/required_method]: SetAccessibilitySearchButton */
	SetAccessibilitySearchMenu()/* debug [protocol_interface/required_method]: SetAccessibilitySearchMenu */
	SetAccessibilitySelected()/* debug [protocol_interface/required_method]: SetAccessibilitySelected */
	SetAccessibilitySelectedCells()/* debug [protocol_interface/required_method]: SetAccessibilitySelectedCells */
	SetAccessibilitySelectedChildren()/* debug [protocol_interface/required_method]: SetAccessibilitySelectedChildren */
	SetAccessibilitySelectedColumns()/* debug [protocol_interface/required_method]: SetAccessibilitySelectedColumns */
	SetAccessibilitySelectedRows()/* debug [protocol_interface/required_method]: SetAccessibilitySelectedRows */
	SetAccessibilitySelectedText()/* debug [protocol_interface/required_method]: SetAccessibilitySelectedText */
	SetAccessibilitySelectedTextRange()/* debug [protocol_interface/required_method]: SetAccessibilitySelectedTextRange */
	SetAccessibilitySelectedTextRanges()/* debug [protocol_interface/required_method]: SetAccessibilitySelectedTextRanges */
	SetAccessibilityServesAsTitleForUIElements()/* debug [protocol_interface/required_method]: SetAccessibilityServesAsTitleForUIElements */
	SetAccessibilitySharedCharacterRange()/* debug [protocol_interface/required_method]: SetAccessibilitySharedCharacterRange */
	SetAccessibilitySharedFocusElements()/* debug [protocol_interface/required_method]: SetAccessibilitySharedFocusElements */
	SetAccessibilitySharedTextUIElements()/* debug [protocol_interface/required_method]: SetAccessibilitySharedTextUIElements */
	SetAccessibilityShownMenu()/* debug [protocol_interface/required_method]: SetAccessibilityShownMenu */
	SetAccessibilitySortDirection()/* debug [protocol_interface/required_method]: SetAccessibilitySortDirection */
	SetAccessibilitySplitters()/* debug [protocol_interface/required_method]: SetAccessibilitySplitters */
	SetAccessibilitySubrole()/* debug [protocol_interface/required_method]: SetAccessibilitySubrole */
	SetAccessibilityTabs()/* debug [protocol_interface/required_method]: SetAccessibilityTabs */
	SetAccessibilityTitle()/* debug [protocol_interface/required_method]: SetAccessibilityTitle */
	SetAccessibilityTitleUIElement()/* debug [protocol_interface/required_method]: SetAccessibilityTitleUIElement */
	SetAccessibilityToolbarButton()/* debug [protocol_interface/required_method]: SetAccessibilityToolbarButton */
	SetAccessibilityTopLevelUIElement()/* debug [protocol_interface/required_method]: SetAccessibilityTopLevelUIElement */
	SetAccessibilityURL()/* debug [protocol_interface/required_method]: SetAccessibilityURL */
	SetAccessibilityUnitDescription()/* debug [protocol_interface/required_method]: SetAccessibilityUnitDescription */
	SetAccessibilityUnits()/* debug [protocol_interface/required_method]: SetAccessibilityUnits */
	SetAccessibilityUserInputLabels()/* debug [protocol_interface/required_method]: SetAccessibilityUserInputLabels */
	SetAccessibilityValue()/* debug [protocol_interface/required_method]: SetAccessibilityValue */
	SetAccessibilityValueDescription()/* debug [protocol_interface/required_method]: SetAccessibilityValueDescription */
	SetAccessibilityVerticalScrollBar()/* debug [protocol_interface/required_method]: SetAccessibilityVerticalScrollBar */
	SetAccessibilityVerticalUnitDescription()/* debug [protocol_interface/required_method]: SetAccessibilityVerticalUnitDescription */
	SetAccessibilityVerticalUnits()/* debug [protocol_interface/required_method]: SetAccessibilityVerticalUnits */
	SetAccessibilityVisibleCells()/* debug [protocol_interface/required_method]: SetAccessibilityVisibleCells */
	SetAccessibilityVisibleCharacterRange()/* debug [protocol_interface/required_method]: SetAccessibilityVisibleCharacterRange */
	SetAccessibilityVisibleChildren()/* debug [protocol_interface/required_method]: SetAccessibilityVisibleChildren */
	SetAccessibilityVisibleColumns()/* debug [protocol_interface/required_method]: SetAccessibilityVisibleColumns */
	SetAccessibilityVisibleRows()/* debug [protocol_interface/required_method]: SetAccessibilityVisibleRows */
	SetAccessibilityWarningValue()/* debug [protocol_interface/required_method]: SetAccessibilityWarningValue */
	SetAccessibilityWindow()/* debug [protocol_interface/required_method]: SetAccessibilityWindow */
	SetAccessibilityWindows()/* debug [protocol_interface/required_method]: SetAccessibilityWindows */
	SetAccessibilityZoomButton()/* debug [protocol_interface/required_method]: SetAccessibilityZoomButton */
}
