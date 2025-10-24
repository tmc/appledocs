//go:build ios || tvos || visionos

// Test for appledocs-531: Array return type type checking fix
// This test verifies that the type conversion fixes compile correctly
// We don't actually call the methods because base NSObject doesn't implement them

package objectivec_test

import (
	"testing"
	"unsafe"

	"github.com/tmc/appledocs/generated/objectivec"
)

// TestAppledocs531TypeConversionsCompile verifies the fix for appledocs-531
// Before the fix, these type assignments would fail with:
// "cannot use rv (variable of type []objc.ID) as IObject value in return statement"
func TestAppledocs531TypeConversionsCompile(t *testing.T) {
	// This test validates that the type conversions compile.
	// We create a variable to ensure the compiler checks the types,
	// but we don't actually call the methods since base NSObject doesn't implement them.

	// Get an object (we need it for method signatures)
	var obj objectivec.Object

	// Suppress "declared and not used" errors by marking _ = expr
	_ = obj

	// Test 1: []objc.ID -> []IObject conversion compiles
	// The method signature should return []IObject, not []objc.ID
	var _ func() []objectivec.IObject = obj.AccessibilityAttributedUserInputLabels

	// Test 2: objc.ID -> IObject conversion compiles
	// The method signature should return IObject, not objc.ID
	var _ func() objectivec.IObject = obj.AccessibilityFrame

	// Test 3: []IObject parameter type compiles
	// The method signature should accept []IObject parameter
	var _ func([]objectivec.IObject) = obj.SetAccessibilityAttributedUserInputLabels

	// Test 4: IObject parameter type compiles
	// The method signature should accept IObject parameter
	var _ func(objectivec.IObject) = obj.SetAccessibilityFrame

	// Test 5: Primitives are not converted
	// The method signature should return uint, not some wrapped type
	var _ func() uint = obj.ImageVersion

	// Test 6: Function types are preserved
	// The method signature should return the function type, not convert it
	var _ func() func() unsafe.Pointer = obj.AccessibilityExpandedStatusBlock

	// Test 7: Cross-framework type fallback
	// Methods returning types from other frameworks should fall back to []IObject
	var _ func(objectivec.IObject) []objectivec.IObject = obj.IndicesOfObjectsByEvaluatingObjectSpecifier

	t.Log("✓ All type conversions from appledocs-531 fix compile correctly!")
	t.Log("  - []objc.ID → []IObject wrapping works")
	t.Log("  - objc.ID → IObject wrapping works")
	t.Log("  - Primitive types preserved")
	t.Log("  - Function types preserved")
	t.Log("  - Cross-framework fallback works")
}

// TestObjectStructure verifies basic Object structure
func TestObjectStructure(t *testing.T) {
	obj := objectivec.NewObject()

	// Verify Object embeds objc.ID
	if obj.ID == 0 {
		t.Error("Object.ID should not be zero after NewObject()")
	}

	// Verify Object implements IObject interface
	var _ objectivec.IObject = obj

	t.Logf("✓ Object structure is correct (ID: %d)", obj.ID)
}
