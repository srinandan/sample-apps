// Code generated by go-bindata. DO NOT EDIT.
// sources:
// ../db/items.json (1.135kB)

package inventorydata

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

var _DbItemsJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xd4\x92\x31\x4b\xc7\x30\x10\xc5\xf7\x7c\x8a\x70\xae\x7f\x8a\x6d\x29\xad\xae\x0a\xe2\xa4\x22\x4e\xe2\x10\xd3\x43\x82\xc7\x5d\x49\x93\x82\x88\xdf\x5d\x9a\x4a\x8d\x0e\x22\xfc\x97\x36\xe3\xbd\x77\x3f\xde\x3d\xf2\xa8\xb4\x7e\x57\x3a\x3d\x70\x3c\x21\x07\xf1\x6f\x70\xbe\x0e\xb5\x86\xd1\x10\xde\x7a\x67\xf1\xc7\x58\x6b\xb0\xd1\x7b\x64\x3b\xdb\xe1\xe1\xfe\x12\x0e\xb9\x38\x19\x8a\xf3\x06\x54\x55\x53\xb4\x15\xac\xda\xc7\x21\x43\x23\xd1\x0d\x5f\x89\xbc\x10\xde\x45\xc3\xc1\x85\x44\x2b\x33\x16\x98\xc9\x38\x32\xcf\x8e\x16\xf1\x34\x93\x5e\x1d\xf7\xb3\xdf\x0a\x07\xe4\x70\xf2\x7d\x41\x66\x1a\x8e\xc8\x5e\x77\x45\x59\x67\xd9\xd5\xaf\x1b\x60\xf0\xd2\x47\x1b\xae\x53\x8c\xb3\xf4\x56\x16\x8c\x41\x3c\x5e\x48\x9f\x58\xc2\xe4\x18\x17\xd6\xd7\xfe\xce\x9a\x6f\x36\xde\x7c\xb7\xeb\xe6\x9b\x3f\x9a\x2f\xb7\xfe\xe9\xdb\xff\x54\xaf\x9e\x3e\x03\x00\x00\xff\xff\x95\xe5\xeb\x3b\x6f\x04\x00\x00")

func DbItemsJsonBytes() ([]byte, error) {
	return bindataRead(
		_DbItemsJson,
		"../db/items.json",
	)
}

func DbItemsJson() (*asset, error) {
	bytes, err := DbItemsJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "../db/items.json", size: 1135, mode: os.FileMode(0644), modTime: time.Unix(1577468874, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x8c, 0x23, 0x81, 0x5f, 0xaf, 0xad, 0xf7, 0x3e, 0x31, 0x23, 0xec, 0xba, 0xe1, 0x2b, 0xce, 0xf5, 0xd8, 0xa, 0xab, 0x29, 0x36, 0xa8, 0x54, 0xca, 0xe4, 0xa, 0xc5, 0x1a, 0xad, 0x34, 0x3e, 0xc9}}
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
	"../db/items.json": DbItemsJson,
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
		"db": {nil, map[string]*bintree{
			"items.json": {DbItemsJson, map[string]*bintree{}},
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
