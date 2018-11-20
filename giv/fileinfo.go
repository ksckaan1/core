// Copyright (c) 2018, The GoKi Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package giv

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/alecthomas/chroma/lexers"
	"github.com/c2h5oh/datasize"
	"github.com/goki/gi/filecat"
	"github.com/goki/gi/gi"
	"github.com/goki/ki"
	"github.com/goki/ki/kit"
)

// FileInfo represents the information about a given file / directory,
// including icon, mimetype, etc
type FileInfo struct {
	Ic      gi.IconName `tableview:"no-header" desc:"icon for file"` // tableview:"no-header"
	Name    string      `width:"40" desc:"name of the file, without any path"`
	Size    FileSize    `desc:"size of the file in bytes"`
	Kind    string      `width:"20" max-width:"20" desc:"type of file / directory -- shorter, more user-friendly version of mime type, based on category"`
	Mime    string      `tableview:"-" desc:"full official mime type of the contents"`
	Cat     filecat.Cat `tableview:"-" desc:"functional category of the file, based on mime data etc"`
	Mode    os.FileMode `desc:"file mode bits"`
	ModTime FileTime    `desc:"time that contents (only) were last modified"`
	Path    string      `view:"-" tableview:"-" desc:"full path to file, including name -- for file functions"`
}

var KiT_FileInfo = kit.Types.AddType(&FileInfo{}, FileInfoProps)

// NewFileInfo returns a new FileInfo based on a filename -- directly returns
// filepath.Abs or os.Stat error on the given file.  filename can be anything
// that works given current directory -- Path will contain the full
// filepath.Abs path, and Name will be just the filename.
func NewFileInfo(fname string) (*FileInfo, error) {
	fi := &FileInfo{}
	err := fi.InitFile(fname)
	return fi, err
}

// InitFile initializes a FileInfo based on a filename -- directly returns
// filepath.Abs or os.Stat error on the given file.  filename can be anything
// that works given current directory -- Path will contain the full
// filepath.Abs path, and Name will be just the filename.
func (fi *FileInfo) InitFile(fname string) error {
	path, err := filepath.Abs(fname)
	if err != nil {
		return err
	}
	fi.Path = path
	_, fi.Name = filepath.Split(path)
	return fi.Stat()
}

// Stat runs os.Stat on file, returns any error directly but otherwise updates
// file info, including mime type, which then drives Kind and Icon -- this is
// the main function to call to update state.
func (fi *FileInfo) Stat() error {
	info, err := os.Stat(fi.Path)
	if err != nil {
		return err
	}
	fi.Size = FileSize(info.Size())
	fi.Mode = info.Mode()
	fi.ModTime = FileTime(info.ModTime())
	if info.IsDir() {
		fi.Kind = "Folder"
		fi.Cat = filecat.Folder
	} else {
		fi.Cat = filecat.Unknown
		fi.Kind = ""
		mtyp, _, err := filecat.MimeFromFile(fi.Path)
		if err == nil {
			fi.Mime = mtyp
			fi.Cat = filecat.CatFromMime(fi.Mime)
			if fi.Cat != filecat.Unknown {
				fi.Kind = fi.Cat.String() + ": "
			}
			fi.Kind += FileKindFromMime(fi.Mime)
		}
		if fi.Cat == filecat.Unknown {
			if fi.IsExec() {
				fi.Cat = filecat.Exe
			}
		}
	}
	icn, _ := fi.FindIcon()
	fi.Ic = icn
	return nil
}

// IsDir returns true if file is a directory (folder)
func (fi *FileInfo) IsDir() bool {
	return fi.Mode.IsDir()
}

// IsExec returns true if file is an executable file
func (fi *FileInfo) IsExec() bool {
	if fi.Mode&0111 != 0 {
		return true
	}
	ext := filepath.Ext(fi.Path)
	if ext == ".exe" {
		return true
	}
	return false
}

// IsSymLink returns true if file is a symbolic link
func (fi *FileInfo) IsSymlink() bool {
	return fi.Mode&os.ModeSymlink != 0
}

//////////////////////////////////////////////////////////////////////////////
//    File ops

// Duplicate creates a copy of given file -- only works for regular files, not
// directories.
func (fi *FileInfo) Duplicate() error {
	if fi.IsDir() {
		err := fmt.Errorf("giv.Duplicate: cannot copy directory: %v", fi.Path)
		log.Println(err)
		return err
	}
	ext := filepath.Ext(fi.Path)
	noext := strings.TrimSuffix(fi.Path, ext)
	dst := noext + "_Copy" + ext
	return CopyFile(dst, fi.Path, fi.Mode)
}

// Delete deletes this file -- does not work on directories (todo: fix)
func (fi *FileInfo) Delete() error {
	if fi.IsDir() {
		err := fmt.Errorf("giv.Delete: cannot deleted directory: %v", fi.Path)
		log.Println(err)
		return err
	}
	return os.Remove(fi.Path)
	// note: we should be deleted now!
}

// Rename renames file to new name
func (fi *FileInfo) Rename(newpath string) error {
	if newpath == "" {
		err := fmt.Errorf("giv.Rename: new name is empty")
		log.Println(err)
		return err
	}
	if newpath == fi.Path {
		return nil
	}
	ndir, np := filepath.Split(newpath)
	if ndir == "" {
		if np == fi.Name {
			return nil
		}
		dir, _ := filepath.Split(fi.Path)
		newpath = filepath.Join(dir, newpath)
	}
	err := os.Rename(fi.Path, newpath)
	if err == nil {
		fi.InitFile(newpath)
	}
	return err
}

// FileKindFromMime returns simplified Kind description based on the given full
// mime type string.  Strips out application/ prefix, and converts all the
// chroma-based mime-types to their basic names
func FileKindFromMime(mime string) string {
	if CustomMimeToKindMap != nil {
		if kind, ok := CustomMimeToKindMap[mime]; ok {
			return kind
		}
	}
	if csidx := strings.Index(mime, ";"); csidx > 0 {
		mime = mime[:csidx]
	}
	if mt, has := filecat.AvailMimes[mime]; has {
		if mt.Support != filecat.NoSupport {
			return mt.Support.String()
		}
	}
	MimeToKindMapInit()
	if kind, ok := MimeToKindMap[mime]; ok {
		return kind
	}
	if sidx := strings.Index(mime, "/"); sidx > 0 {
		mime = mime[sidx+1:]
	}
	return mime
}

// MimeToKindMapInit makes sure the MimeToKindMap is initialized from
// InitMimeToKindMap plus chroma lexer types.
func MimeToKindMapInit() {
	if MimeToKindMap != nil {
		return
	}
	MimeToKindMap = InitMimeToKindMap
	for _, l := range lexers.Registry.Lexers {
		config := l.Config()
		nm := strings.ToLower(config.Name)
		if len(config.MimeTypes) > 0 {
			mtyp := config.MimeTypes[0]
			MimeToKindMap[mtyp] = nm
		} else {
			MimeToKindMap["application/"+nm] = nm
		}
	}
}

// FindIcon uses file info to find an appropriate icon for this file -- uses
// Kind string first to find a correspondingly-named icon, and then tries the
// extension.  Returns true on success.
func (fi *FileInfo) FindIcon() (gi.IconName, bool) {
	kind := fi.Kind
	icn := gi.IconName(kind)
	if icn.IsValid() {
		return icn, true
	}
	kind = strings.ToLower(kind)
	icn = gi.IconName(kind)
	if icn.IsValid() {
		return icn, true
	}
	if fi.IsDir() {
		return gi.IconName("folder"), true
	}
	if icn = "file-" + gi.IconName(kind); icn.IsValid() {
		return icn, true
	}
	if ms, ok := KindToIconMap[kind]; ok {
		if icn = gi.IconName(ms); icn.IsValid() {
			return icn, true
		}
	}
	if strings.Contains(kind, "/") {
		si := strings.IndexByte(kind, '/')
		typ := kind[:si]
		subtyp := kind[si+1:]
		if icn = "file-" + gi.IconName(subtyp); icn.IsValid() {
			return icn, true
		}
		if icn = gi.IconName(subtyp); icn.IsValid() {
			return icn, true
		}
		if ms, ok := KindToIconMap[string(subtyp)]; ok {
			if icn = gi.IconName(ms); icn.IsValid() {
				return icn, true
			}
		}
		if icn = "file-" + gi.IconName(typ); icn.IsValid() {
			return icn, true
		}
		if icn = gi.IconName(typ); icn.IsValid() {
			return icn, true
		}
		if ms, ok := KindToIconMap[string(typ)]; ok {
			if icn = gi.IconName(ms); icn.IsValid() {
				return icn, true
			}
		}
	}
	ext := filepath.Ext(fi.Name)
	if ext != "" {
		if icn = gi.IconName(ext[1:]); icn.IsValid() {
			return icn, true
		}
	}

	icn = gi.IconName("none")
	return icn, false
}

var FileInfoProps = ki.Props{
	"CtxtMenu": ki.PropSlice{
		{"Duplicate", ki.Props{
			"updtfunc": ActionUpdateFunc(func(fii interface{}, act *gi.Action) {
				fi := fii.(*FileInfo)
				act.SetInactiveState(fi.IsDir())
			}),
		}},
		{"Delete", ki.Props{
			"desc":    "Ok to delete this file?  This is not undoable and is not moving to trash / recycle bin",
			"confirm": true,
			"updtfunc": ActionUpdateFunc(func(fii interface{}, act *gi.Action) {
				fi := fii.(*FileInfo)
				act.SetInactiveState(fi.IsDir())
			}),
		}},
		{"Rename", ki.Props{
			"desc": "Rename file to new file name",
			"Args": ki.PropSlice{
				{"New Name", ki.Props{
					"default-field": "Name",
				}},
			},
		}},
	},
}

//////////////////////////////////////////////////////////////////////////////
//    FileTime, FileSize

// Note: can get all the detailed birth, access, change times from this package
// 	"github.com/djherbis/times"

// FileTime provides a default String format for file modification times, and
// other useful methods -- will plug into ValueView with date / time editor.
type FileTime time.Time

// Int satisfies the ints.Inter interface for sorting etc
func (ft FileTime) Int() int64 {
	return (time.Time(ft)).Unix()
}

// FromInt satisfies the ints.Inter interface
func (ft *FileTime) FromInt(val int64) {
	*ft = FileTime(time.Unix(val, 0))
}

func (ft FileTime) String() string {
	return (time.Time)(ft).Format("Mon Jan  2 15:04:05 MST 2006")
}

func (ft FileTime) MarshalBinary() ([]byte, error) {
	return time.Time(ft).MarshalBinary()
}

func (ft FileTime) MarshalJSON() ([]byte, error) {
	return time.Time(ft).MarshalJSON()
}

func (ft FileTime) MarshalText() ([]byte, error) {
	return time.Time(ft).MarshalText()
}

func (ft *FileTime) UnmarshalBinary(data []byte) error {
	return (*time.Time)(ft).UnmarshalBinary(data)
}

func (ft *FileTime) UnmarshalJSON(data []byte) error {
	return (*time.Time)(ft).UnmarshalJSON(data)
}

func (ft *FileTime) UnmarshalText(data []byte) error {
	return (*time.Time)(ft).UnmarshalText(data)
}

type FileSize datasize.ByteSize

// Int satisfies the kit.Inter interface for sorting etc
func (fs FileSize) Int() int64 {
	return int64(fs) // note: is actually uint64
}

// FromInt satisfies the ints.Inter interface
func (fs *FileSize) FromInt(val int64) {
	*fs = FileSize(val)
}

func (fs FileSize) String() string {
	return (datasize.ByteSize)(fs).HumanReadable()
}

//////////////////////////////////////////////////////////////////////////////
//    CopyFile

// here's all the discussion about why CopyFile is not in std lib:
// https://old.reddit.com/r/golang/comments/3lfqoh/why_golang_does_not_provide_a_copy_file_func/
// https://github.com/golang/go/issues/8868

// CopyFile copies the contents from src to dst atomically.
// If dst does not exist, CopyFile creates it with permissions perm.
// If the copy fails, CopyFile aborts and dst is preserved.
func CopyFile(dst, src string, perm os.FileMode) error {
	in, err := os.Open(src)
	if err != nil {
		return err
	}
	defer in.Close()
	tmp, err := ioutil.TempFile(filepath.Dir(dst), "")
	if err != nil {
		return err
	}
	_, err = io.Copy(tmp, in)
	if err != nil {
		tmp.Close()
		os.Remove(tmp.Name())
		return err
	}
	if err = tmp.Close(); err != nil {
		os.Remove(tmp.Name())
		return err
	}
	if err = os.Chmod(tmp.Name(), perm); err != nil {
		os.Remove(tmp.Name())
		return err
	}
	return os.Rename(tmp.Name(), dst)
}

//////////////////////////////////////////////////////////////////////////////
//    Kind, Icon Maps

// MimeToKindMap maps from mime type names to kind names.  Add any standard
// manual cases to InitMimeToKindMap, which will be used here along with the
// chroma lexer mime to name mapping.
var MimeToKindMap map[string]string

// InitMimeToKindMap maps from mime type names to kind names.  Add any
// standard manual cases here -- will be used as start of MimeToKindMap, which
// is kept empty as trigger for initialization.
var InitMimeToKindMap = map[string]string{}

// CustomMimeToKindMap maps from mime type names to kind names, and can be set
// by user for any special cases.  This is used before the standard one.
var CustomMimeToKindMap map[string]string

// KindToIconMap has special cases for mapping mime type to icon, for those
// that basic string doesn't work
var KindToIconMap = map[string]string{
	"svg+xml":           "svg",
	"msword":            "file-word",
	"postscript":        "file-pdf",
	"vnd.ms-excel":      "file-excel",
	"vnd.ms-powerpoint": "file-powerpoint",
	"x-apple-diskimage": "file-zip",
	"octet-stream":      "file-binary",
	"gzip":              "file-zip",
}
