// Package goversioninfo creates a syso file which contains Microsoft Version Information and an optional icon.
package goversioninfo

import (
	"bytes"
	"io"

	"github.com/akavel/rsrc/coff"
)

// *****************************************************************************
// JSON and Config
// *****************************************************************************

// ParseJSON parses the given bytes as a VersionInfo JSON.
func (vi *VersionInfo) ParseJSON(jsonBytes []byte) error { _ = "STUB: not implemented"; return nil }

// VersionInfo data container
type VersionInfo struct {
	FixedFileInfo       `json:"FixedFileInfo"`
	StringFileInfo      `json:"StringFileInfo"`
	VarFileInfo         `json:"VarFileInfo"`
	Timestamp           bool
	Buffer              bytes.Buffer
	Structure           VSVersionInfo
	IconPath            string `json:"IconPath"`
	ManifestPath        string `json:"ManifestPath"`
	ApplicationIconPath string `json:"ApplicationIconPath"`
}

// Translation with langid and charsetid.
type Translation struct {
	LangID    `json:"LangID"`
	CharsetID `json:"CharsetID"`
}

// FileVersion with 3 parts.
type FileVersion struct {
	Major int
	Minor int
	Patch int
	Build int
}

// FixedFileInfo contains file characteristics - leave most of them at the defaults.
type FixedFileInfo struct {
	FileVersion    `json:"FileVersion"`
	ProductVersion FileVersion
	FileFlagsMask  string
	FileFlags      string
	FileOS         string
	FileType       string
	FileSubType    string
}

// VarFileInfo is the translation container.
type VarFileInfo struct {
	Translation `json:"Translation"`
}

// StringFileInfo is what you want to change.
type StringFileInfo struct {
	Comments         string
	CompanyName      string
	FileDescription  string
	FileVersion      string
	InternalName     string
	LegalCopyright   string
	LegalTrademarks  string
	OriginalFilename string
	PrivateBuild     string
	ProductName      string
	ProductVersion   string
	SpecialBuild     string
}

// *****************************************************************************
// Helpers
// *****************************************************************************

// SizedReader is a *bytes.Buffer.
type SizedReader struct {
	*bytes.Buffer
}

// Size returns the length of the buffer.
func (s SizedReader) Size() int64 { _ = "STUB: not implemented"; return 0 }

func str2Uint32(s string) uint32 { _ = "STUB: not implemented"; return 0 }

func padString(s string, zeros int) []byte { _ = "STUB: not implemented"; return nil }

func padBytes(i int) []byte { _ = "STUB: not implemented"; return nil }

// NewFileVersion parses semver version string into a FileVersion object
func NewFileVersion(version string) (FileVersion, error) {
	_ = "STUB: not implemented"
	return *new(FileVersion), nil
}

// First match group is a whole matched string.

func (f FileVersion) getVersionHighString() string { _ = "STUB: not implemented"; return "" }

func (f FileVersion) getVersionLowString() string { _ = "STUB: not implemented"; return "" }

// IsZero returns true if all version components are zero.
func (f FileVersion) IsZero() bool { _ = "STUB: not implemented"; return false }

// GetVersionString returns a string representation of the version
func (f FileVersion) GetVersionString() string { _ = "STUB: not implemented"; return "" }

// fillVersions syncs version info between FixedFileInfo and StringFileInfo.
// If one section has version data and the other doesn't, the missing section
// is populated automatically. Warnings are logged when StringFileInfo version
// strings cannot be parsed or when the two sections have conflicting values.
func (vi *VersionInfo) fillVersions() { _ = "STUB: not implemented"; return }

func (vi *VersionInfo) fillVersion(name string, fixed *FileVersion, str *string) {
	_ = "STUB: not implemented"
	return
}

func (t Translation) getTranslationString() string { _ = "STUB: not implemented"; return "" }

func (t Translation) getTranslation() string { _ = "STUB: not implemented"; return "" }

// *****************************************************************************
// IO Methods
// *****************************************************************************

// Walk writes the data buffer with hexadecimal data from the structs
func (vi *VersionInfo) Walk() {
	_ = "STUB: not implemented"
	// Create a buffer
	return
}

// Write to the buffer

// WriteSyso creates a resource file from the version info and optionally an icon.
// arch must be an architecture string accepted by coff.Arch, like "386" or "amd64"
func (vi *VersionInfo) WriteSyso(filename string, arch string) error {
	_ = "STUB: not implemented"
	return nil
}

// Create a new RSRC section

// Set the architecture

// ID 16 is for Version Information

// If manifest is enabled

// If icon is enabled

// IDI_APPLICATION (32512) is the icon shown in the window title bar.
// Default to IconPath if not explicitly set.

// Write to file

// WriteHex creates a hex file for debugging version info
func (vi *VersionInfo) WriteHex(filename string) error { _ = "STUB: not implemented"; return nil }

// WriteGo creates a Go file that contains the version info so you can access
// it in the application
func (vi *VersionInfo) WriteGo(filename, packageName string) error {
	_ = "STUB: not implemented"
	return nil
}

func writeCoff(coff *coff.Coff, fnameout string) error { _ = "STUB: not implemented"; return nil }

func writeCoffTo(w io.WriteCloser, coff *coff.Coff) error { _ = "STUB: not implemented"; return nil }

// write the resulting file to disk
