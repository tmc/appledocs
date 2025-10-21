// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation_test

import (
	"github.com/tmc/appledocs/generated/foundation"
)

// Suppress unused import errors
var _ = foundation.NewAttributedStringMarkdownSourcePosition

// ExampleNewAttributedStringMarkdownSourcePositionWithStartLineStartColumnEndLineEndColumn demonstrates how to create a AttributedStringMarkdownSourcePosition instance using NewAttributedStringMarkdownSourcePositionWithStartLineStartColumnEndLineEndColumn.
// Creates a Markdown source position instance from its start and end line and column.
func ExampleNewAttributedStringMarkdownSourcePositionWithStartLineStartColumnEndLineEndColumn() {
	_ = foundation.NewAttributedStringMarkdownSourcePositionWithStartLineStartColumnEndLineEndColumn(
		0, // startLine int
		0, // startColumn int
		0, // endLine int
		0, // endColumn int
	)
	// Output:
}

