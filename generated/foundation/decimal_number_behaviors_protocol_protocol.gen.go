// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// PDecimalNumberBehaviors is the NSDecimalNumberBehaviors protocol interface.
//
// A protocol that declares three methods that control the discretionary aspects of working with decimal numbers.
//
// Availability:
//   - Mac Catalyst 13.0+
//   - iOS 2.0+
//   - iPadOS 2.0+
//   - macOS 10.0+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 2.0+
//
// See: doc://com.apple.foundation/documentation/Foundation/NSDecimalNumberBehaviors
type PDecimalNumberBehaviors interface {
	// Required methods
	ExceptionDuringOperationErrorLeftOperandRightOperand(operation objc.SEL, error_ CalculationError, leftOperand IDecimalNumber, rightOperand IDecimalNumber) DecimalNumber
	RoundingMode() RoundingMode
	Scale() unsafe.Pointer
}
