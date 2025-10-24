// Code generated from Apple documentation for Accessibility. DO NOT EDIT.

package accessibility

// PAXMathExpressionProvider is the AXMathExpressionProvider protocol interface.
//
// Availability:
//   - Mac Catalyst 14.0+
//   - iOS 14.0+
//   - iPadOS 14.0+
//   - macOS 11.0+
//   - tvOS 14.0+
//   - visionOS 1.0+
//   - watchOS 7.0+
//
// See: doc://com.apple.Accessibility/documentation/Accessibility/AXMathExpressionProvider
type PAXMathExpressionProvider interface {
	// Required methods
	AccessibilityMathExpression() AXMathExpression/* debug [protocol_interface/required_method]: AccessibilityMathExpression */
}
