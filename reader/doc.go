// Package reader provides an io/fs.FS interface to Apple documentation JSON files.
//
// The package allows access to downloaded Apple documentation without embedding
// the entire 2GB+ documentation archive into binaries. It provides both low-level
// filesystem access and high-level query functions for working with frameworks,
// symbols, and documentation metadata.
//
// # Basic Usage
//
// Create a reader from a documentation cache directory:
//
//	fsys, err := reader.Open("output/tutorials/data/documentation")
//	if err != nil {
//		log.Fatal(err)
//	}
//
// Read a framework's documentation:
//
//	doc, err := reader.GetFramework(fsys, "Foundation")
//	if err != nil {
//		log.Fatal(err)
//	}
//
// List all available frameworks:
//
//	frameworks, err := reader.ListFrameworks(fsys)
//	if err != nil {
//		log.Fatal(err)
//	}
//
// # File System Structure
//
// The reader expects a directory structure matching Apple's documentation layout:
//
//	documentation/
//	├── Foundation.json
//	├── Foundation/
//	│   ├── NSString.json
//	│   └── NSString/
//	│       └── ...
//	├── AppKit.json
//	└── AppKit/
//	    └── ...
//
// Framework-level documentation is stored in JSON files at the root level
// (e.g., Foundation.json), while symbols and nested types are organized in
// subdirectories.
package reader
