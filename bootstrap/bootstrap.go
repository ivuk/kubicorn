// Code generated by go-bindata.
// sources:
// bootstrap/1.7.0_ubuntu_16.04_master.sh
// bootstrap/1.7.0_ubuntu_16.04_node.sh
// DO NOT EDIT!

package bootstrap

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func bindataRead(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}
	if clErr != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

type asset struct {
	bytes []byte
	info  os.FileInfo
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

var _bootstrap170_ubuntu_1604_masterSh = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\x90\x31\x6f\xe4\x20\x14\x84\x7b\xff\x8a\xb9\xbd\x62\x2b\xa0\xdf\xe2\xba\xfb\x17\xd7\x60\x78\xb2\x91\x31\xcf\xe2\x3d\x4e\xb1\x14\xe5\xb7\x47\xf6\xb2\xca\x26\x69\x52\xd9\x33\xa3\xf9\x18\xf8\xfd\xcb\x35\xa9\x6e\x4c\xc5\x51\xf9\x8f\xd1\xcb\x3c\x08\x29\x0c\x0d\x21\xe2\x6d\x18\xa4\x45\x46\x68\x35\xc3\x08\x66\xd5\x4d\x6e\xce\x6d\x3e\x2c\x7e\x22\xb1\x21\x73\x8b\x76\x62\x9e\x32\xd9\xc0\xab\xf3\x9b\xba\xc8\xe1\xf8\x9a\x85\x76\x3b\x6d\x13\x5e\x71\x42\xba\x05\x1f\x23\xcc\x9d\xab\xdc\xc2\x0c\x47\x7a\x16\x9c\x70\xab\x81\xc4\xe6\x24\x6a\xa3\x5b\xda\x48\xb5\x90\x76\xe7\x5e\x91\x19\x26\xe0\x4a\x61\x66\x5c\x22\x8d\xe7\xa6\x9b\x3b\xfa\xf6\xa9\x90\xd8\xe1\x43\x9a\x17\x2a\xc9\x67\xac\x3e\x95\x0b\xfe\xfc\xf4\xc4\x6b\xbf\xfe\xb1\x7c\x22\x45\xdb\xa2\x57\x82\xd9\x3f\xdb\xa9\x88\xfa\x9c\x61\x76\xfc\x1b\x00\x40\x38\x78\xed\xff\x34\xaa\x1f\x33\x49\x97\x91\xc3\x42\xd5\x26\xee\xfa\x60\x68\xf5\x45\x36\xae\x6a\xce\xf7\xed\xc9\xb1\x25\x93\x3e\x29\x1f\xd7\x3e\x48\x76\x51\x5a\x83\x66\x50\x39\xe8\x1d\xfb\x35\x14\xf5\x55\x1f\xd9\x3d\x34\x7f\x1f\x28\x54\x12\xd2\x6f\x6e\x2a\x49\xdf\x03\x00\x00\xff\xff\xff\x5a\x2a\x62\x15\x02\x00\x00")

func bootstrap170_ubuntu_1604_masterShBytes() ([]byte, error) {
	return bindataRead(
		_bootstrap170_ubuntu_1604_masterSh,
		"bootstrap/1.7.0_ubuntu_16.04_master.sh",
	)
}

func bootstrap170_ubuntu_1604_masterSh() (*asset, error) {
	bytes, err := bootstrap170_ubuntu_1604_masterShBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "bootstrap/1.7.0_ubuntu_16.04_master.sh", size: 533, mode: os.FileMode(420), modTime: time.Unix(1499612499, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _bootstrap170_ubuntu_1604_nodeSh = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\x90\x31\x6f\xe4\x20\x14\x84\x7b\xff\x8a\xb9\xbd\x62\x2b\xa0\xdf\xe2\xba\xfb\x17\xd7\x60\x78\xb2\x91\x31\xcf\xe2\x3d\x4e\xb1\x14\xe5\xb7\x47\xf6\xb2\xca\x26\x69\x52\xd9\x33\xa3\xf9\x18\xf8\xfd\xcb\x35\xa9\x6e\x4c\xc5\x51\xf9\x8f\xd1\xcb\x3c\x08\x29\x0c\x0d\x21\xe2\x6d\x18\xa4\x45\x46\x68\x35\xc3\x08\x66\xd5\x4d\x6e\xce\x6d\x3e\x2c\x7e\x22\xb1\x21\x73\x8b\x76\x62\x9e\x32\xd9\xc0\xab\xf3\x9b\xba\xc8\xe1\xf8\x9a\x85\x76\x3b\x6d\x13\x5e\x71\x42\xba\x05\x1f\x23\xcc\x9d\xab\xdc\xc2\x0c\x47\x7a\x16\x9c\x70\xab\x81\xc4\xe6\x24\x6a\xa3\x5b\xda\x48\xb5\x90\x76\xe7\x5e\x91\x19\x26\xe0\x4a\x61\x66\x5c\x22\x8d\xe7\xa6\x9b\x3b\xfa\xf6\xa9\x90\xd8\xe1\x43\x9a\x17\x2a\xc9\x67\xac\x3e\x95\x0b\xfe\xfc\xf4\xc4\x6b\xbf\xfe\xb1\x7c\x22\x45\xdb\xa2\x57\x82\xd9\x3f\xdb\xa9\x88\xfa\x9c\x61\x76\xfc\x1b\x00\x40\x38\x78\xed\xff\x34\xaa\x1f\x33\x49\x97\x91\xc3\x42\xd5\x26\xee\xfa\x60\x68\xf5\x45\x36\xae\x6a\xce\xf7\xed\xc9\xb1\x25\x93\x3e\x29\x1f\xd7\x3e\x48\x76\x51\x5a\x83\x66\x50\x39\xe8\x1d\xfb\x35\x14\xf5\x55\x1f\xd9\x3d\x34\x7f\x1f\x28\x54\x12\xd2\x6f\x6e\x2a\x49\xdf\x03\x00\x00\xff\xff\xff\x5a\x2a\x62\x15\x02\x00\x00")

func bootstrap170_ubuntu_1604_nodeShBytes() ([]byte, error) {
	return bindataRead(
		_bootstrap170_ubuntu_1604_nodeSh,
		"bootstrap/1.7.0_ubuntu_16.04_node.sh",
	)
}

func bootstrap170_ubuntu_1604_nodeSh() (*asset, error) {
	bytes, err := bootstrap170_ubuntu_1604_nodeShBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "bootstrap/1.7.0_ubuntu_16.04_node.sh", size: 533, mode: os.FileMode(420), modTime: time.Unix(1499612499, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
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

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
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
	"bootstrap/1.7.0_ubuntu_16.04_master.sh": bootstrap170_ubuntu_1604_masterSh,
	"bootstrap/1.7.0_ubuntu_16.04_node.sh": bootstrap170_ubuntu_1604_nodeSh,
}

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//     data/
//       foo.txt
//       img/
//         a.png
//         b.png
// then AssetDir("data") would return []string{"foo.txt", "img"}
// AssetDir("data/img") would return []string{"a.png", "b.png"}
// AssetDir("foo.txt") and AssetDir("notexist") would return an error
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		cannonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(cannonicalName, "/")
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
	"bootstrap": &bintree{nil, map[string]*bintree{
		"1.7.0_ubuntu_16.04_master.sh": &bintree{bootstrap170_ubuntu_1604_masterSh, map[string]*bintree{}},
		"1.7.0_ubuntu_16.04_node.sh": &bintree{bootstrap170_ubuntu_1604_nodeSh, map[string]*bintree{}},
	}},
}}

// RestoreAsset restores an asset under the given directory
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
	err = ioutil.WriteFile(_filePath(dir, name), data, info.Mode())
	if err != nil {
		return err
	}
	err = os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
	if err != nil {
		return err
	}
	return nil
}

// RestoreAssets restores an asset under the given directory recursively
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
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(cannonicalName, "/")...)...)
}

