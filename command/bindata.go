// Code generated by go-bindata.
// sources:
// libraries/amzlinux/definition.yaml
// libraries/amzlinux/nodejs/install.yaml
// libraries/amzlinux/nodejs/nodejs.sh
// libraries/amzlinux/ruby/install.yaml
// libraries/amzlinux/ruby/ruby.sh
// libraries/ubuntu1604/nodejs.sh
// libraries/ubuntu1604/ruby.sh
// DO NOT EDIT!

package command

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

var _librariesAmzlinuxDefinitionYaml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xca\x4b\xcc\x4d\xb5\x52\x48\xcc\xad\xca\xc9\xcc\x2b\xad\xe0\xca\xc8\x07\x71\x95\xf4\x41\xb4\x7e\x6a\xb2\x91\x6e\x69\x71\x6a\x91\x12\x17\x20\x00\x00\xff\xff\xb3\x59\x76\x81\x26\x00\x00\x00")

func librariesAmzlinuxDefinitionYamlBytes() ([]byte, error) {
	return bindataRead(
		_librariesAmzlinuxDefinitionYaml,
		"libraries/amzlinux/definition.yaml",
	)
}

func librariesAmzlinuxDefinitionYaml() (*asset, error) {
	bytes, err := librariesAmzlinuxDefinitionYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "libraries/amzlinux/definition.yaml", size: 38, mode: os.FileMode(420), modTime: time.Unix(1483198862, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _librariesAmzlinuxNodejsInstallYaml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x2a\x28\x4a\xd5\x4d\xce\xcf\xcd\x4d\xcc\x4b\xb1\x52\xc8\x29\xe6\x82\xb3\xf3\x0a\x72\x15\x32\xf3\x8a\x4b\x12\x73\x72\xb8\x00\x01\x00\x00\xff\xff\x6c\x32\x79\xd0\x25\x00\x00\x00")

func librariesAmzlinuxNodejsInstallYamlBytes() ([]byte, error) {
	return bindataRead(
		_librariesAmzlinuxNodejsInstallYaml,
		"libraries/amzlinux/nodejs/install.yaml",
	)
}

func librariesAmzlinuxNodejsInstallYaml() (*asset, error) {
	bytes, err := librariesAmzlinuxNodejsInstallYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "libraries/amzlinux/nodejs/install.yaml", size: 37, mode: os.FileMode(420), modTime: time.Unix(1483236559, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _librariesAmzlinuxNodejsNodejsSh = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x5c\xcc\x4d\x0a\xc2\x30\x10\x06\xd0\x7d\x4e\xf1\x89\xcb\x32\xcd\x52\xf0\x32\x92\x8e\x83\x8d\xe6\xa7\x64\x26\xd0\x80\x87\x97\x6c\x5d\x3f\x78\xd7\x8b\xdf\x62\xf1\x5b\xd0\xdd\x39\x15\x03\x49\x3f\x9d\xe3\xde\x12\x88\x34\x26\x29\x06\xa2\x54\x39\x58\xac\x05\xbb\xd9\xa1\x77\xef\xdb\x91\xd7\x52\x9f\xa2\xb5\x37\x96\x95\x6b\xf6\x2a\xd6\x8f\xc7\x6d\x3d\xf1\xc5\xfc\x40\x6e\xf4\x8c\x58\xd4\x42\x4a\xa0\x81\x17\x33\xf1\xb2\x20\x87\x8f\x80\xc6\x3f\xcf\xef\xad\x13\x7e\x01\x00\x00\xff\xff\xe4\x84\xe2\x91\x96\x00\x00\x00")

func librariesAmzlinuxNodejsNodejsShBytes() ([]byte, error) {
	return bindataRead(
		_librariesAmzlinuxNodejsNodejsSh,
		"libraries/amzlinux/nodejs/nodejs.sh",
	)
}

func librariesAmzlinuxNodejsNodejsSh() (*asset, error) {
	bytes, err := librariesAmzlinuxNodejsNodejsShBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "libraries/amzlinux/nodejs/nodejs.sh", size: 150, mode: os.FileMode(420), modTime: time.Unix(1483141464, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _librariesAmzlinuxRubyInstallYaml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x2a\x28\x4a\x8d\x4f\xce\xcf\xcd\x4d\xcc\x4b\xb1\x52\x48\x4f\xcd\x55\xc8\xcc\x2b\x2e\x49\xcc\xc9\x51\x48\x2a\xcd\x4b\xc9\x49\x2d\xe2\x82\x4b\x42\x04\x60\xf2\x5c\x80\x00\x00\x00\xff\xff\x98\xd3\x76\xaf\x39\x00\x00\x00")

func librariesAmzlinuxRubyInstallYamlBytes() ([]byte, error) {
	return bindataRead(
		_librariesAmzlinuxRubyInstallYaml,
		"libraries/amzlinux/ruby/install.yaml",
	)
}

func librariesAmzlinuxRubyInstallYaml() (*asset, error) {
	bytes, err := librariesAmzlinuxRubyInstallYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "libraries/amzlinux/ruby/install.yaml", size: 57, mode: os.FileMode(420), modTime: time.Unix(1483311047, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _librariesAmzlinuxRubyRubySh = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x5c\x8e\x31\xae\xc2\x30\x10\x44\x7b\x9f\x62\x7f\x3e\xad\x31\x0a\x1d\x12\x91\xe8\x28\x29\xb8\x80\xed\x8c\x70\xa4\x38\x8e\x76\x6d\x84\x6f\x8f\x1c\x44\x43\xb3\xab\x99\x37\xc5\xfb\xff\x33\x6e\x5a\x8c\xb3\x12\x94\xaa\x25\x12\xd8\x0a\x88\x8b\xab\xfd\x61\x7b\xa4\xeb\x06\xa6\x45\xb2\x9d\xe7\x0f\x3a\xea\x11\x4f\x7c\x43\x9b\x3c\x10\xa9\xac\xa3\xcd\x20\xad\xa5\x4a\x46\x54\xf0\x21\x51\x87\xd7\x9a\x38\xd3\xed\x72\xbf\x9e\x77\xed\x9e\x4c\x48\x11\x06\xbe\xd7\x45\xc0\x4d\xa0\xa3\x61\xa0\x9f\x7a\xdf\xa4\xd8\xab\x77\x00\x00\x00\xff\xff\xba\x1e\x2a\x60\xa3\x00\x00\x00")

func librariesAmzlinuxRubyRubyShBytes() ([]byte, error) {
	return bindataRead(
		_librariesAmzlinuxRubyRubySh,
		"libraries/amzlinux/ruby/ruby.sh",
	)
}

func librariesAmzlinuxRubyRubySh() (*asset, error) {
	bytes, err := librariesAmzlinuxRubyRubyShBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "libraries/amzlinux/ruby/ruby.sh", size: 163, mode: os.FileMode(420), modTime: time.Unix(1483379470, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _librariesUbuntu1604NodejsSh = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x01\x00\x00\xff\xff\x00\x00\x00\x00\x00\x00\x00\x00")

func librariesUbuntu1604NodejsShBytes() ([]byte, error) {
	return bindataRead(
		_librariesUbuntu1604NodejsSh,
		"libraries/ubuntu1604/nodejs.sh",
	)
}

func librariesUbuntu1604NodejsSh() (*asset, error) {
	bytes, err := librariesUbuntu1604NodejsShBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "libraries/ubuntu1604/nodejs.sh", size: 0, mode: os.FileMode(420), modTime: time.Unix(1483118317, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _librariesUbuntu1604RubySh = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x52\x56\xd4\x4f\xca\xcc\xd3\x4f\x4a\x2c\xce\xe0\xe2\x2a\x2e\x4d\xc9\x57\xa8\x2c\xcd\x55\xc8\xcc\x2b\x2e\x49\xcc\xc9\x51\x28\x2a\x4d\xaa\xd4\x4d\x49\x2d\x4b\x85\x30\x15\xd2\xf3\x15\x2a\x92\x93\x74\x4b\x4b\x32\x73\x14\x74\x2b\xb9\x00\x01\x00\x00\xff\xff\xa6\x58\x98\xfc\x3d\x00\x00\x00")

func librariesUbuntu1604RubyShBytes() ([]byte, error) {
	return bindataRead(
		_librariesUbuntu1604RubySh,
		"libraries/ubuntu1604/ruby.sh",
	)
}

func librariesUbuntu1604RubySh() (*asset, error) {
	bytes, err := librariesUbuntu1604RubyShBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "libraries/ubuntu1604/ruby.sh", size: 61, mode: os.FileMode(420), modTime: time.Unix(1483038579, 0)}
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
	"libraries/amzlinux/definition.yaml": librariesAmzlinuxDefinitionYaml,
	"libraries/amzlinux/nodejs/install.yaml": librariesAmzlinuxNodejsInstallYaml,
	"libraries/amzlinux/nodejs/nodejs.sh": librariesAmzlinuxNodejsNodejsSh,
	"libraries/amzlinux/ruby/install.yaml": librariesAmzlinuxRubyInstallYaml,
	"libraries/amzlinux/ruby/ruby.sh": librariesAmzlinuxRubyRubySh,
	"libraries/ubuntu1604/nodejs.sh": librariesUbuntu1604NodejsSh,
	"libraries/ubuntu1604/ruby.sh": librariesUbuntu1604RubySh,
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
	"libraries": &bintree{nil, map[string]*bintree{
		"amzlinux": &bintree{nil, map[string]*bintree{
			"definition.yaml": &bintree{librariesAmzlinuxDefinitionYaml, map[string]*bintree{}},
			"nodejs": &bintree{nil, map[string]*bintree{
				"install.yaml": &bintree{librariesAmzlinuxNodejsInstallYaml, map[string]*bintree{}},
				"nodejs.sh": &bintree{librariesAmzlinuxNodejsNodejsSh, map[string]*bintree{}},
			}},
			"ruby": &bintree{nil, map[string]*bintree{
				"install.yaml": &bintree{librariesAmzlinuxRubyInstallYaml, map[string]*bintree{}},
				"ruby.sh": &bintree{librariesAmzlinuxRubyRubySh, map[string]*bintree{}},
			}},
		}},
		"ubuntu1604": &bintree{nil, map[string]*bintree{
			"nodejs.sh": &bintree{librariesUbuntu1604NodejsSh, map[string]*bintree{}},
			"ruby.sh": &bintree{librariesUbuntu1604RubySh, map[string]*bintree{}},
		}},
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
