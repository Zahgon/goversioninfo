package goversioninfo

import (
	"bytes"
	"os"

	"github.com/akavel/rsrc/coff"
	"github.com/akavel/rsrc/ico"
)

// *****************************************************************************
/*
Code from https://github.com/akavel/rsrc

The MIT License (MIT)

Copyright (c) 2013-2014 The rsrc Authors.

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
// *****************************************************************************

const (
	rtIcon      = coff.RT_ICON
	rtGroupIcon = coff.RT_GROUP_ICON
	rtManifest  = coff.RT_MANIFEST
)

// on storing icons, see: http://blogs.msdn.com/b/oldnewthing/archive/2012/07/20/10331787.aspx
type gRPICONDIR struct {
	ico.ICONDIR
	Entries []gRPICONDIRENTRY
}

func (group gRPICONDIR) Size() int64 { _ = "STUB: not implemented"; return 0 }

type gRPICONDIRENTRY struct {
	ico.IconDirEntryCommon
	ID uint16
}

func addIcon(coff *coff.Coff, fnames string, newID func() uint16) error {
	_ = "STUB: not implemented"
	return nil
}

func addOneIcon(coff *coff.Coff, fname string, newID func() uint16) error {
	_ = "STUB: not implemented"
	return nil
}

func addIconWithGroupID(coff *coff.Coff, fname string, newID func() uint16, groupID uint16) error {
	_ = "STUB: not implemented"
	return nil
}

func addOneIconWithGroupID(coff *coff.Coff, fname string, newID func() uint16, groupID uint16) error {
	_ = "STUB: not implemented"
	return nil
}

// RT_ICONs

// magic num.
// magic num.

func bufferIcon(f *os.File, offset int64, size int) (*bytes.Reader, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
