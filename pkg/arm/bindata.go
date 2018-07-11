// Code generated by go-bindata. DO NOT EDIT.
// sources:
// data/azuredeploy.json
// data/startup.sh
package arm

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

var _azuredeployJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x1b\x5d\x6f\xdb\x38\xf2\x3d\xbf\x82\xd0\x15\x48\xbb\x88\xbf\x92\xb4\x57\x04\xb8\x07\x37\xd9\x5e\x8d\x6e\x5a\x23\x6a\xf7\x1e\x16\x41\x97\xa6\xc6\x36\x37\x12\xa9\x23\x29\x37\xae\xcf\xff\xfd\xa0\x4f\x4b\x32\x65\x52\x8e\xdb\x04\x77\x65\x1f\xdc\x8a\x33\x43\xce\xf7\x70\xc8\xae\x8e\x10\x42\xc8\x79\x26\xc9\x1c\x02\xec\x5c\x20\x67\xae\x54\x28\x2f\x7a\xbd\xf4\x4b\x37\xc0\x0c\xcf\x20\x00\xa6\xba\xf8\x5b\x24\xa0\x4b\x78\x90\xcd\xc9\xde\x69\x7f\xf0\xb2\xd3\x1f\x74\xfa\x83\x9e\x07\xa1\xcf\x97\x31\xdc\x27\x08\x42\x1f\x2b\xe8\xfe\x25\x39\xfb\x9b\x73\x92\xae\x40\x38\x53\xc0\xd4\xef\x20\x24\xe5\x2c\x5e\x68\xd0\xed\xc7\x7f\x72\x80\x10\x0b\x1c\x80\x02\x21\x9d\x0b\xb4\x5a\x67\x5f\x17\x58\x50\x3c\xf1\xa1\xf2\x51\x80\xe4\x91\x20\xc9\xc7\x3f\x8e\x56\x2b\x44\xa7\x88\x71\x85\x9e\x23\xca\x3c\xb8\x47\xdd\x4b\xce\x14\xa6\x0c\x84\x0b\x62\x41\x09\x74\xc7\x82\x87\x20\x14\x05\xd9\x1d\xce\x80\xa9\x31\xe7\xfe\x58\xf0\x29\xf5\x41\xa2\xfe\x8b\xee\xef\x0c\x94\x1b\x4d\x18\xa8\xd1\x15\x5a\xaf\x51\x65\x1c\xe5\x7f\x59\x1d\x95\x3f\x3b\x6a\x19\x42\xcc\xc8\x35\x25\x82\x4b\x3e\x55\xdd\x0f\xa0\xbe\x72\x71\xd7\x5b\x50\xa1\x22\xec\x67\xff\x94\x19\x8b\x05\x22\x0e\x69\x49\x0e\xa7\xfd\xc1\xab\x4e\xff\xac\x73\xd6\xaf\xc3\xf9\x9c\x60\x95\x41\xad\x56\x1a\xb6\x7e\xcb\x00\xd0\x7f\xd0\x5f\x12\xad\xd7\x75\x02\x0c\x07\xc9\x0e\x17\x0c\x54\x7d\x2e\x2c\x44\x12\x8b\xf6\x08\xd5\x86\x83\x3d\x4f\x80\x94\x6e\x88\x09\x68\x21\xca\x50\x63\x01\x53\x7a\x9f\x29\x44\x07\x98\x00\x0f\xfa\xa9\xca\x7b\xaf\x1d\x2d\xd0\xed\xd6\xd7\xf5\xc9\xf6\xc6\x64\xa2\xa7\xe6\xa5\xf4\x3b\x45\x65\x79\x78\x30\xc5\x91\x5f\x17\x49\x05\xd4\x20\x9e\x0a\x6c\x45\x08\x89\x69\xe7\x8c\x9e\x9e\xeb\x39\x4d\x78\xd3\xce\x6c\x7f\xad\x4a\x65\x33\xbf\x3e\x89\x4d\x1f\x98\x87\xd6\xeb\xf6\x26\x1a\x46\x13\x9f\x92\xd1\x78\x98\xee\x1d\x1e\xdf\x48\x69\xd8\x11\x3c\x52\x20\xda\x5a\x6a\xc1\x8a\x9f\xef\xe5\x1a\xd4\x9c\x7b\x31\x51\x57\x61\x45\x89\x46\xd1\x0e\xf5\x7c\xf8\x44\x03\xe0\x91\x1a\xb1\x6b\xca\x22\x95\xd0\x1f\xbc\xac\x8a\xbb\xb6\x17\x79\x17\xe9\x37\x91\x73\xf1\x06\x4b\x4a\x9c\x66\x95\xb5\x56\x94\xcf\xb1\xf7\x06\xfb\x98\x91\x38\x34\xee\xa9\x24\x0f\x42\x60\x9e\xfc\xc8\xb4\x5e\xe3\xfc\x91\x47\xd4\x91\xf7\xfc\xd8\xc2\x56\x8e\x4f\xd0\x71\xa1\xae\xe3\x17\xb7\x55\x7e\x6f\xbf\x97\x85\xf8\x93\x3d\x2d\x64\x82\xc9\x1d\x30\x2f\xdb\x7f\x1c\xfc\x1f\x14\x3e\x32\x72\x7a\xd7\xd6\x38\xb0\xc6\xfa\xa6\x22\xc9\x87\xde\x68\x7c\xc9\xd9\x94\xce\x22\x91\x48\xe0\x41\xdb\xca\x69\x1e\x2a\xac\x85\x82\x2e\xb0\x02\xbd\x63\x5d\x2d\x19\x0e\xb4\x9e\x55\xa5\x51\x35\x1d\xe3\xa2\x28\x75\xcd\x78\x85\xc3\x1a\x65\x7d\xe8\xa3\x6f\xf3\x8c\x9d\x5a\x29\x9b\xf0\x88\x79\x1f\xb0\x2a\x8c\x6c\x37\xd8\x4d\x94\xd6\x36\x5a\xb0\x8d\xeb\x53\x36\x2b\x20\xf7\x35\x8f\x90\x0b\xd5\x79\x5d\x0f\x0d\x15\xd0\x36\xd6\xb1\xed\x53\xad\x94\x4b\x38\x23\x58\x3d\xdf\xad\xe3\x4a\xec\x8b\xf5\x5b\x44\x80\xe3\x17\x27\xe8\xb8\xa7\xf1\xeb\xfc\x9b\xd9\x00\x0c\x96\x9b\xd1\x19\x73\xa1\x9c\x0b\xf4\xba\x6f\x00\x07\x16\x97\xa9\x6f\x7d\x8e\x15\x65\xb3\xd1\xd8\xb9\x40\x53\xec\x4b\x30\xa0\x35\xc4\x81\x1f\x2f\xca\xa6\x80\x54\x4c\x3c\x58\xa0\x39\x21\x5b\x89\x36\x26\x68\x03\x5e\xcc\xe9\x15\x95\x4a\xd0\x49\x94\xa7\x9d\x2b\x63\xc1\x87\x32\xfb\x9f\x34\x17\xbb\xb5\xdd\x1d\x4e\xf6\xc9\xb2\xb2\x97\x39\xe8\x83\x05\x1d\x0a\xae\x38\x49\xfc\xd1\xf9\x44\xc2\xd6\x15\xa8\x9e\xbc\x6d\x84\x39\x3f\x3f\xfb\x19\x62\xd2\xd1\x2e\xc4\x9c\x9f\x9f\xfd\x8c\x31\x87\x8d\x31\x66\x91\xfe\x7f\x06\x99\xf3\xf3\xb3\x47\x8f\x32\x56\xf5\x14\x8f\x94\x55\xa5\x94\x32\xf7\x84\xaa\x23\xca\x14\x88\x05\xf6\x47\xcc\x05\xc2\x99\x17\xa3\x98\x6c\x8a\x45\xc1\x04\xc4\xc7\xe9\x38\x67\xe6\xd4\xa4\x02\xcb\x4c\xfa\x3f\x93\x10\x7e\xa4\x54\xcd\xb1\xe3\xf0\x1e\xf0\x44\x5a\x0f\xae\xe2\x02\xcf\xa0\x27\xd3\xdf\x21\x21\x3c\x62\xca\xdc\x7c\x78\xd9\xe9\xbf\xea\x0c\x5e\x7e\xb7\x0e\x51\x86\x3c\xa5\xb3\xee\x0d\xcc\xe2\xe8\xbb\x74\x2b\x5b\x6c\x22\x60\xec\x75\xa6\xe8\x9f\x32\x49\xb8\x0a\x33\x0f\x0b\xef\xcb\x6f\x37\x6e\xa3\x3c\x8f\x56\xab\x0e\x12\x98\xcd\x00\x3d\xc3\x61\x88\x2e\xfe\xd1\xb6\xef\xbc\xde\xe8\xc4\xba\x1b\xc4\xd2\x5f\x17\x48\x24\xa8\x5a\xfe\x53\xf0\x28\x3c\x54\xeb\xee\xd9\x0e\xcd\x34\x2b\x85\xc9\x59\x27\x46\xc6\x61\xd8\xfd\x80\x03\xd8\x43\xfc\x32\xe3\xe6\xc1\x47\x5c\xec\xfb\xfc\xeb\x17\x29\xe7\x07\xeb\xec\x12\x92\x76\x2d\x9c\x61\x4c\xda\x94\xda\x3d\x90\x44\xd0\x30\x97\x69\x82\x83\x5c\xf7\x1d\x52\x02\x4f\xa7\xe6\x6e\x89\x07\x52\x51\x96\x48\x7c\x58\xef\x29\xff\xd2\x02\x39\x2e\x7f\x6e\x62\xd3\x4c\xf4\x7f\xda\x39\x3d\x35\x22\x53\x01\x24\xdf\xf7\x28\x6d\x4f\x98\x0b\x19\xca\x63\xb5\xc5\xf5\x52\x7f\xd0\x32\x58\x1a\xc0\xd3\x4a\xa7\xbd\x10\x52\xbc\x0a\xff\xbf\xd8\xf6\xe9\xd6\xd9\x55\x12\xfc\x3b\xb5\xe6\x1b\xee\x03\xfa\x93\xb2\xa9\xc0\x7f\x26\x97\x42\x5a\x3a\x27\xd6\x76\x39\x57\x6a\x17\xdf\x3f\xdc\x30\xe3\xfd\x3c\xaa\x65\xee\x2c\xb7\xd0\x41\xcc\xd2\x98\xec\x9f\x9a\x59\xa2\xef\x50\x86\x6d\xcc\xaf\x9e\x2a\x2a\xe0\x8f\x61\x7f\xf2\x51\x0d\x70\x77\x69\x8a\x0e\x62\x81\x6d\xab\xc8\xa7\x6c\x81\x9a\x0b\xc7\x7c\xec\xb8\xa5\xb4\x2f\x3c\x2f\x79\x10\x46\x0a\xf2\xfb\xf3\x6b\x4c\xe6\x94\x81\x4b\xb0\x0f\x2e\x58\x14\xa0\x7f\xef\x0c\x4e\x3b\xfd\xc1\x03\xea\x9c\x86\x02\xb2\x7a\x7f\x66\x48\x13\xd5\xb1\x5d\xee\x18\x2e\x36\x76\x1d\xe4\x6f\x9d\xf2\xa5\xef\xe6\xe5\x43\xb2\x91\xda\x23\x86\xd6\x0b\xd7\x5e\x2d\xc4\x4b\x2f\x18\xa8\xfa\xaa\x6d\xc9\x6a\x4b\xd6\x98\xb8\xa6\x72\x34\xde\x25\xe6\x21\x4d\x4a\x53\xd1\xd9\x78\x5c\x52\x14\x44\xb9\xca\xd7\x5d\x0e\x13\x1c\x62\x92\x7a\xb0\x93\xaf\x73\xb9\xe3\x84\x81\x6a\xc7\x94\x54\x1f\xd7\x2e\xfd\x06\x05\x46\xfd\x70\x17\x9f\x1f\x32\x33\xca\x4f\x35\xa3\x00\xcf\xe0\x26\x93\x66\xc2\x97\xe3\xd4\x45\xee\x84\x3e\xd6\xf7\xee\x74\xe7\xa4\x84\xa2\xfb\xfe\xf3\xae\x7d\x27\x37\x69\x72\x9e\x4a\xa5\x8e\x3b\xce\x27\x77\x52\x10\xdc\x8b\x88\xd2\xe1\x7f\x9c\x4e\x4b\xb8\x3a\x19\x68\xcc\xca\x78\x62\xe0\x0b\x10\xa1\xe0\x0b\x9a\x39\x7f\x43\x1b\xd4\x89\xc2\x99\xc0\x1e\x8c\xb9\x4f\xc9\xb2\xf9\x15\x4b\xc0\xbd\x34\x12\x61\x16\x61\x7f\x3b\x2c\xea\x9e\xa1\x54\x63\x54\x76\xa8\x6b\x5e\x82\x4b\x13\x48\x02\x86\xbd\x80\xb2\xcf\x12\x44\xae\x4b\xe2\xf3\xc8\xeb\x44\x72\xeb\xd2\xbb\x82\x46\xd2\xc0\x29\x62\x9b\xd9\x24\x85\x9a\x87\x74\x76\x51\xf0\x29\x8b\xee\xdb\xf5\x87\x1d\x8f\x4a\x3c\xf1\x61\x8c\xa5\xfc\xca\x85\x37\x8c\xd4\x1c\x98\xa2\x45\xb0\x55\x22\x32\xf5\xa6\xe3\xb3\x9a\x55\x2b\x34\xbd\xee\x7d\x0f\xcb\xdd\x6f\x8c\xca\xc3\x4c\xb5\xa0\x7e\x07\xcb\x2b\xac\x70\x26\x34\xd7\x7d\x37\xce\x97\x1b\x4a\x57\x09\xca\x66\x1b\x27\x75\xdd\x77\xef\x61\xd9\x2d\x20\xf4\x3e\xd1\xcc\x08\x56\x31\xcb\x4e\x6f\xce\x03\xe8\x6d\xd4\xdb\xeb\x4a\x39\xef\xe1\x48\xcd\xb9\xa0\xdf\xc0\xfb\x72\x17\xf3\x6a\x45\xb7\xf9\x2e\x3b\x1f\xdb\x2f\xab\xec\xf0\x5b\x55\xa2\x4e\xd6\x2f\xb2\x32\x74\x9a\x86\xba\x29\x08\x60\xd9\x03\xb3\x3d\x03\xe2\x16\x69\x1e\x87\x9c\x72\xae\x6f\x08\x46\x36\x4f\x17\x4a\x61\xf1\x59\xeb\xb8\x58\x15\x4e\x92\x8f\xb6\x09\xed\x0e\xce\x15\x12\x8b\x4d\xb1\xb3\x45\x26\x2b\x84\x36\x91\x36\x09\xad\xbe\x04\xa3\xb8\x34\xaf\x2d\xb6\xa8\xe7\x2a\x48\x92\x77\xbe\x46\x9c\xc3\xb7\x6b\xb7\x44\xb1\x49\x7e\x6f\xa4\x92\x28\x32\x27\x12\x67\xfc\x86\x34\x50\x1e\x3b\xae\x23\x1c\x2e\xaf\xa8\xbc\x33\xc7\x2b\x92\xc4\xea\x59\xcc\xee\x0d\x60\xef\x5f\x82\x2a\x30\xc9\x9c\x08\xc0\x0a\x3e\x16\x47\x97\xb7\x82\x07\x09\x33\x26\xc4\xf4\xe1\xac\x67\xb5\x33\x54\xf2\x9e\x61\xb5\x17\x39\x16\x10\xd0\x28\xd8\x6e\x45\xd6\xc7\xa1\x9c\x38\x2b\xd7\xac\x9c\x38\x83\x1d\x31\x05\x62\x8a\x09\x58\xbe\x65\xca\x87\x85\x50\x8a\x2e\xa3\xf1\x84\x88\x5a\x1e\x61\x4b\x38\x34\xc0\x62\x69\x95\xaf\x0a\xa4\xf4\x2a\x76\x34\x7e\xcb\xc5\x57\x2c\xbc\xd4\xaa\x5a\xe0\xeb\x8a\x62\xeb\x2d\xa3\x06\xb7\x3d\x70\xed\xdd\x34\x0c\x37\x83\x9b\x1d\x86\x2d\xed\xa1\x3c\xec\x25\x81\x50\xe5\x11\x29\x49\xd6\x6c\x91\x91\xd1\x9e\x86\x53\xc3\x6f\x6f\x44\x15\x02\xe9\xab\xe6\x2c\x0f\xc6\x69\xd0\xea\x48\x67\x45\x3a\xb5\x94\xe2\x54\x52\x25\xe9\x24\x27\x3b\x8b\x3c\x61\x20\x6f\x77\x49\xdd\x78\xc4\x3c\x41\xc7\xbd\xec\x61\x77\x2f\x7b\x9d\x9d\x65\x06\x53\x62\x68\x1a\x96\x56\x5a\x61\x25\x7b\x53\x18\x66\xed\x95\xf6\xcf\x25\xb4\x54\x37\xb6\xd9\xd2\x2a\x37\xfb\x7a\x98\x75\x16\x74\xec\xde\x3c\xdb\x8e\x3d\x74\xb2\x57\x87\x7d\xd7\x38\x71\xca\xed\x92\x37\x2d\xde\xf9\x9a\xc6\x83\xc4\xfc\x34\xde\x1f\x35\x8d\xf6\x8a\xbb\xdd\xd3\x11\xed\xa1\xed\x20\x77\x9f\x63\xcc\x74\x9a\x67\xf5\x94\x9b\x8a\x24\xb8\x57\xc0\xe2\x7a\xdb\xaa\x4c\x2a\xa0\x0f\x5a\x12\x11\x69\xaa\x41\xd1\xbe\x25\x11\x8e\x14\xff\x9c\xb6\x4e\xae\x29\xe3\x62\xd3\x65\x6d\x51\xe2\x84\x82\x2b\x20\x0a\x3c\x17\x94\xa2\x6c\xd6\x2e\x76\x39\xe9\x85\x41\x96\xb9\xde\x60\x09\xaf\xce\x7f\x65\x84\x7b\x80\x9e\xbb\x0a\x0b\x15\x85\xa5\xea\xe5\xc5\x56\x73\xa9\x69\xd8\x16\x2f\x95\xd3\xdf\xc6\x6d\x87\xc9\xff\x81\xfb\x75\xa3\x50\x4b\x72\xb2\x24\x03\xdb\x2d\xe4\xfd\xf1\xcb\x48\x2a\x1e\xb8\xa9\x3c\x5a\xe0\xbe\xc3\xcc\xf3\x41\x94\x5b\xe4\xdd\xbe\x59\x4a\x07\xf6\xa0\xed\x66\x5a\xd3\x4d\x41\x2d\xc6\x64\xad\x5f\x87\x47\x2a\x8c\x54\x2a\xb9\xa3\xf5\xd1\x7f\x03\x00\x00\xff\xff\x6e\xef\x6d\xcb\xb2\x38\x00\x00")

func azuredeployJsonBytes() ([]byte, error) {
	return bindataRead(
		_azuredeployJson,
		"azuredeploy.json",
	)
}

func azuredeployJson() (*asset, error) {
	bytes, err := azuredeployJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "azuredeploy.json", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _startupSh = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xac\x96\xe1\x4e\xe3\x46\x10\xc7\xbf\xef\x53\x4c\x43\xd5\xb4\xa7\xda\x86\xeb\x09\xb5\x1c\x20\x41\x08\x08\xc1\x91\x88\xa4\x27\x9d\xaa\x0a\xad\x77\xc7\x78\x15\x7b\xd7\xcc\x8e\x23\x7c\x21\xef\x5e\x39\x4e\x42\x02\x39\x0e\x4e\xfd\x06\xb3\x33\xbf\xff\xec\xcc\x7a\x26\x5b\x3f\x45\xb1\xb1\x51\x2c\x7d\x0a\x01\xde\x0b\xb1\x05\xc3\xde\x49\x6f\x0f\x22\x64\x15\x69\xeb\x73\xe9\xef\x42\x1d\x39\x32\xb7\xc6\x06\x65\xe1\x99\x50\xe6\x81\xb6\x3e\x54\xce\x26\x60\x3c\xa8\x92\x08\x2d\x67\x15\xa4\x92\xb4\x72\x1a\xf5\x47\x30\x2c\xb6\xa0\x20\x17\xcb\x38\xab\xc0\xa7\xae\xcc\xb4\x6d\x33\xc4\x28\xc4\xa0\x7b\xfd\xf9\xbc\xd3\xbd\x19\x7e\xe9\x77\x0f\x1a\xb2\x30\x09\xfc\x03\x41\x02\xad\x99\xb0\xaf\x7c\x4d\x37\xb7\x91\x64\x97\x1b\x15\xb8\x02\xad\x4f\x4d\xc2\x81\x75\x1a\x5b\xf0\xef\x47\xe0\x14\xad\x00\x00\x58\xc3\x3d\xf5\x17\x89\x11\xc2\x57\x9e\x31\x57\x9c\x81\x67\x57\x80\x76\x6a\x84\x14\x7a\xa4\xb1\x51\x28\xca\xdc\x95\x96\x21\xca\x2d\x47\x84\xde\x95\xa4\x10\x1e\x1e\x80\xa9\x44\x91\x8f\x12\x1f\xde\x27\xbe\xce\x2d\xd2\x38\x8e\xbc\x8e\x77\x04\xaa\xd4\x41\x7b\xf9\x3f\x40\x34\x96\x14\x65\x26\x8e\x1a\x36\x40\x1d\x02\xb7\x54\xdc\x95\x8e\x25\xc0\x36\x6c\xb7\xe1\xf0\x70\x76\xb9\xc4\xb3\x8c\xc5\x5c\x74\x3d\x4e\x10\x7a\x76\x84\xca\x59\x08\xae\x9f\x9d\xae\xde\x43\x12\x3f\xbd\x88\x98\x4c\x4c\x02\x78\x07\x61\xf7\x9e\x49\x86\xd7\x2e\x43\x68\x19\x9b\x90\x6c\x4d\xa7\x4d\xd2\xad\xe3\x5e\x6f\x38\x18\x5e\x1f\xf5\x6f\x3a\xbd\xab\xd3\xf3\xb3\x9b\xab\xa3\x4f\xdd\x83\xba\xaa\x41\x53\xf2\xa0\x09\x58\x64\xfb\xd8\x8a\x9f\x27\xab\x95\x9e\xce\x3a\x21\x26\x13\xcc\x3c\xbe\x81\xae\x5c\x5e\x94\x8c\x6f\xe0\x5b\x3d\x9d\x0a\xe1\x51\x43\x60\x20\x40\x68\xf9\xad\x93\xee\xf1\xdf\x67\x37\x97\xbd\xb3\xcb\xee\xe7\xee\xe5\xc1\xfb\xa7\x86\x0f\x5b\x2d\x78\x0d\x5d\x50\x0e\x01\x25\x8d\x2f\xb2\xd2\xd1\xbb\xe6\xef\xe6\x55\x46\xb9\xf4\x8c\x14\xbd\x13\x22\x96\x1e\x77\x3f\x40\xa0\x61\x7f\x7f\x1f\x26\x13\x38\x9e\x19\xba\xb6\x7e\xef\xf0\xeb\x17\x99\x67\x9f\x24\xf9\x54\x66\x10\x76\x66\x8a\xe1\x95\xd3\x78\xec\x1c\x7b\x26\x59\x5c\x94\x31\x36\x99\xfc\x06\xd3\x29\x1c\xae\xaa\xd4\xa9\x44\xf1\xc2\x33\x1c\x2d\x5d\xbf\xa7\xda\x41\xe2\x23\x7f\x5c\x31\xfa\xcd\xaa\xb5\xc3\x37\xf4\x66\x3d\x59\x8a\x16\x48\xa1\x22\xfe\x9e\x60\x9f\xcc\x58\x32\x5e\x60\xf5\xa2\xec\x05\x56\xaf\x56\x1d\x61\x25\x54\x9a\x3b\x0d\xdb\xbb\xdb\xdb\xf0\xba\x88\xe7\x6e\x1b\xcb\xf7\x23\xf5\xeb\xc8\x17\x8a\xa6\xe4\xac\x4a\xaa\x78\x9e\x40\x73\xd4\xd8\x8b\x91\x89\x94\x0c\x98\x4a\xcf\x51\x33\x50\x22\x69\x55\xea\xc8\x47\x8f\x93\x6c\x0e\x2b\x0b\x2d\x19\x83\x85\xbf\x98\x0f\x17\x2b\x73\xac\x3f\x6c\x24\xd8\xd9\xfd\x33\xdc\xfd\x23\xdc\x79\xff\x57\xb8\xb3\xdb\xde\x90\x56\x3d\xb5\xb2\xf1\x6c\x20\x8b\x7c\xa4\x0d\x41\xb0\x9e\xa1\xca\x5c\xa9\x0b\x72\x63\xa3\x91\x1e\x87\x3c\xa7\xc6\xd7\x13\x5c\x97\x45\x66\x94\x64\xd4\x60\x6c\x3d\x58\x21\xc5\x2c\x07\x95\x4a\x62\x0f\x89\xa3\x99\xad\xf9\x18\xa0\x70\xda\xff\x0e\xd2\x6a\xe0\x54\x72\xdb\x83\x75\x0c\x46\xa3\xcc\x84\x92\xbc\x9e\xdd\x9a\x6e\x24\xbf\x96\x84\xcd\xda\xd8\xdf\x6f\x77\x7b\xa7\x6d\xc1\x68\xa5\xe5\x73\xbd\x57\x37\xa7\x6e\x01\x4b\x63\x91\x06\xcd\x44\x0b\xfb\xe4\x0a\x24\x36\xe8\xc3\xa3\xaf\x7d\x72\x89\xc9\x30\x1c\x36\x31\x27\xf0\x00\xf5\x74\x45\x98\x4e\x85\x2f\x63\xaf\xc8\x14\x6c\x9c\x7d\x1b\x6d\xb0\x1a\xb9\xc6\x94\x52\x77\x32\x83\xaf\x4b\x6f\x6e\xea\x93\xb1\xca\x14\x32\x5b\xe0\xe7\x84\xcd\xe0\x01\x2a\x42\xfe\x71\x78\x13\xff\x04\x3d\xfc\x9f\x4a\xba\x58\x85\x67\xe4\xca\xe2\x2d\xb0\xeb\xd5\xc0\x55\x62\xe6\x94\xac\xcb\xfc\x0d\xd8\xe5\xfc\x78\xad\xaf\xa8\x4a\x32\x5c\xcd\x58\x57\x32\xc7\x3d\xb0\x7e\xb9\x44\x44\x41\x26\x97\x54\x0d\x94\xcc\x70\x80\xdc\x38\x78\xbf\x3c\x1f\xe7\xc3\xaa\xc0\x3d\x18\xe7\xde\x8b\x6e\xef\xb4\x7e\xfa\xd6\x31\xee\xc1\xa6\x75\x00\x8a\xea\xdf\x41\x99\x73\x85\x87\xd2\xb2\xc9\x16\x6f\xde\x78\x28\x8b\x95\xf5\x8b\x56\xc6\x19\x6e\x84\x2c\xb7\xf1\xd3\x65\xfd\x92\x33\xfc\x22\xfe\x0b\x00\x00\xff\xff\xb5\xd4\x96\x92\x8a\x09\x00\x00")

func startupShBytes() ([]byte, error) {
	return bindataRead(
		_startupSh,
		"startup.sh",
	)
}

func startupSh() (*asset, error) {
	bytes, err := startupShBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "startup.sh", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
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
	"azuredeploy.json": azuredeployJson,
	"startup.sh":       startupSh,
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
	"azuredeploy.json": {azuredeployJson, map[string]*bintree{}},
	"startup.sh":       {startupSh, map[string]*bintree{}},
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
