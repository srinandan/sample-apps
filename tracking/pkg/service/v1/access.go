// Code generated by go-bindata. DO NOT EDIT.
// sources:
// ../../../db/tracking.json (1.422kB)

package v1

import (
	"bytes"
	"compress/gzip"
	"crypto/sha256"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func bindataRead(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("read %q: %w", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("read %q: %w", name, err)
	}
	if clErr != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

type asset struct {
	bytes  []byte
	info   os.FileInfo
	digest [sha256.Size]byte
}

type bindataFileInfo struct {
	name    string
	size    int64
	mode    os.FileMode
	modTime time.Time
}

func (fi bindataFileInfo) Name() string {
	return fi.name
}
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}
func (fi bindataFileInfo) IsDir() bool {
	return false
}
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _DbTrackingJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xdc\x93\xcf\x4a\x03\x31\x10\xc6\xef\x7d\x8a\x90\x73\x2b\xc9\xd2\x5a\x9b\x9b\x8a\x22\xe2\x41\x68\x7b\x51\xa4\x2c\xc9\xd0\x0e\x5d\x93\x25\x99\xad\x56\xe9\xbb\x4b\x52\xad\xd1\xee\xc1\x93\x62\x87\x65\x09\xf9\x7e\xf3\x87\xf9\xc8\x7d\x87\x31\xc6\x5e\xd3\x3f\x06\x27\x5f\xea\x25\xda\xf9\x0c\x0d\x57\x8c\x17\x22\x06\xef\x7e\x02\x81\x4a\x6a\x42\xd4\x0c\x54\xb8\x02\x0f\x26\x97\xb5\x87\x92\x60\x46\xf8\x08\xdb\x7c\xd9\xef\x49\xd9\x93\x27\x13\x29\xd4\x40\xaa\x41\xff\x2e\xc7\x9b\xda\xb4\xe2\xa3\x76\x3c\xe0\xdc\x42\x1a\xec\xda\x2d\x2c\x9b\x40\x20\xf0\x39\xf0\x04\x38\x5f\x50\x04\xe4\xf0\xe8\x38\x57\x20\xd0\xec\x7d\xe2\xf5\x5e\xbf\x62\x38\x11\x42\xa5\xef\x4b\x3f\x5d\x7a\x8f\xe0\x23\x3a\xbd\x1d\xe7\xca\xc7\x9e\x6e\x9c\x2e\x09\x9d\xe5\x2a\xdb\xe2\x36\x17\x69\x1d\x13\xcf\x60\x05\xbe\x5a\xb3\x2b\xac\xaa\x90\x95\xd8\x6d\x33\x4d\x72\x7e\xfa\x5d\xd2\xae\xb1\xe4\x53\x89\xe9\x78\x4f\x7d\xc1\x9a\x2b\x36\x12\x85\x14\xf9\xfd\x0e\xda\xa4\xd3\xa6\xfb\x13\x7f\x65\xbb\xbf\xce\x32\xef\x1a\x82\xbf\xb0\xf7\x77\x3c\xbd\x04\x73\xf1\x7c\xa8\xae\x16\xed\xae\xd6\xa8\x97\x60\x58\x53\x1f\xae\xad\xd3\xf1\xbf\x7c\xab\x9d\x87\xb7\x00\x00\x00\xff\xff\x11\x60\xf2\x56\x8e\x05\x00\x00")

func DbTrackingJsonBytes() ([]byte, error) {
	return bindataRead(
		_DbTrackingJson,
		"../../../db/tracking.json",
	)
}

func DbTrackingJson() (*asset, error) {
	bytes, err := DbTrackingJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "../../../db/tracking.json", size: 1422, mode: os.FileMode(0644), modTime: time.Unix(1592765401, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x3d, 0xce, 0x2a, 0xe7, 0x7e, 0xa5, 0x31, 0x52, 0xe7, 0xe9, 0xf2, 0x3e, 0x46, 0xdb, 0xb3, 0xf8, 0x9e, 0x15, 0x96, 0x7b, 0xb0, 0xfd, 0x95, 0x79, 0xcd, 0x96, 0x2f, 0xc8, 0x25, 0x3c, 0x78, 0x72}}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[canonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// AssetString returns the asset contents as a string (instead of a []byte).
func AssetString(name string) (string, error) {
	data, err := Asset(name)
	return string(data), err
}

// MustAsset is like Asset but panics when Asset would return an error.
// It simplifies safe initialization of global variables.
func MustAsset(name string) []byte {
	a, err := Asset(name)
	if err != nil {
		panic("asset: Asset(" + name + "): " + err.Error())
	}

	return a
}

// MustAssetString is like AssetString but panics when Asset would return an
// error. It simplifies safe initialization of global variables.
func MustAssetString(name string) string {
	return string(MustAsset(name))
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[canonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetDigest returns the digest of the file with the given name. It returns an
// error if the asset could not be found or the digest could not be loaded.
func AssetDigest(name string) ([sha256.Size]byte, error) {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[canonicalName]; ok {
		a, err := f()
		if err != nil {
			return [sha256.Size]byte{}, fmt.Errorf("AssetDigest %s can't read by error: %v", name, err)
		}
		return a.digest, nil
	}
	return [sha256.Size]byte{}, fmt.Errorf("AssetDigest %s not found", name)
}

// Digests returns a map of all known files and their checksums.
func Digests() (map[string][sha256.Size]byte, error) {
	mp := make(map[string][sha256.Size]byte, len(_bindata))
	for name := range _bindata {
		a, err := _bindata[name]()
		if err != nil {
			return nil, err
		}
		mp[name] = a.digest
	}
	return mp, nil
}

// AssetNames returns the names of the assets.
func AssetNames() []string {
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

// _bindata is a table, holding each asset generator, mapped to its name.
var _bindata = map[string]func() (*asset, error){
	"../../../db/tracking.json": DbTrackingJson,
}

// AssetDebug is true if the assets were built with the debug flag enabled.
const AssetDebug = false

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//
//	data/
//	  foo.txt
//	  img/
//	    a.png
//	    b.png
//
// then AssetDir("data") would return []string{"foo.txt", "img"},
// AssetDir("data/img") would return []string{"a.png", "b.png"},
// AssetDir("foo.txt") and AssetDir("notexist") would return an error, and
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		canonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(canonicalName, "/")
		for _, p := range pathList {
			node = node.Children[p]
			if node == nil {
				return nil, fmt.Errorf("Asset %s not found", name)
			}
		}
	}
	if node.Func != nil {
		return nil, fmt.Errorf("Asset %s not found", name)
	}
	rv := make([]string, 0, len(node.Children))
	for childName := range node.Children {
		rv = append(rv, childName)
	}
	return rv, nil
}

type bintree struct {
	Func     func() (*asset, error)
	Children map[string]*bintree
}

var _bintree = &bintree{nil, map[string]*bintree{
	"..": {nil, map[string]*bintree{
		"..": {nil, map[string]*bintree{
			"..": {nil, map[string]*bintree{
				"db": {nil, map[string]*bintree{
					"tracking.json": {DbTrackingJson, map[string]*bintree{}},
				}},
			}},
		}},
	}},
}}

// RestoreAsset restores an asset under the given directory.
func RestoreAsset(dir, name string) error {
	data, err := Asset(name)
	if err != nil {
		return err
	}
	info, err := AssetInfo(name)
	if err != nil {
		return err
	}
	err = os.MkdirAll(_filePath(dir, filepath.Dir(name)), os.FileMode(0755))
	if err != nil {
		return err
	}
	err = os.WriteFile(_filePath(dir, name), data, info.Mode())
	if err != nil {
		return err
	}
	return os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
}

// RestoreAssets restores an asset under the given directory recursively.
func RestoreAssets(dir, name string) error {
	children, err := AssetDir(name)
	// File
	if err != nil {
		return RestoreAsset(dir, name)
	}
	// Dir
	for _, child := range children {
		err = RestoreAssets(dir, filepath.Join(name, child))
		if err != nil {
			return err
		}
	}
	return nil
}

func _filePath(dir, name string) string {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(canonicalName, "/")...)...)
}