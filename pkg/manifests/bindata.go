// Code generated by go-bindata. DO NOT EDIT.
// sources:
// assets/dns/cluster-role-binding.yaml (223B)
// assets/dns/cluster-role.yaml (397B)
// assets/dns/daemonset.yaml (6.385kB)
// assets/dns/metrics/cluster-role-binding.yaml (279B)
// assets/dns/metrics/cluster-role.yaml (246B)
// assets/dns/metrics/role-binding.yaml (293B)
// assets/dns/metrics/role.yaml (284B)
// assets/dns/namespace.yaml (369B)
// assets/dns/service-account.yaml (85B)
// assets/dns/service.yaml (520B)

package manifests

import (
	"bytes"
	"compress/gzip"
	"crypto/sha256"
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

var _assetsDnsClusterRoleBindingYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x6c\xce\x31\x8e\x83\x40\x0c\x05\xd0\x7e\x4e\xe1\x0b\xc0\x6a\xbb\xd5\x74\x9b\xdc\x80\x48\xe9\xcd\x8c\x09\x0e\x60\xa3\xb1\x87\x22\xa7\x8f\x10\x4a\x45\x3a\x17\xfe\xff\xfd\x89\x25\x47\xb8\xce\xd5\x9c\x4a\xa7\x33\x5d\x58\x32\xcb\x23\xe0\xca\x77\x2a\xc6\x2a\x11\x4a\x8f\xa9\xc5\xea\xa3\x16\x7e\xa1\xb3\x4a\x3b\xfd\x59\xcb\xfa\xb3\xfd\x86\x85\x1c\x33\x3a\xc6\x00\x00\x20\xb8\x50\x04\x5d\x49\x6c\xe4\xc1\x9b\x2c\x16\xac\xf6\x4f\x4a\x6e\x31\x34\x70\x78\x37\x2a\x1b\x27\xfa\x4f\x49\xab\x78\xf8\xc4\xf6\xe7\xe3\xb6\x15\xd3\xa9\xa7\xe8\x4c\x1d\x0d\x3b\x74\x9a\x1d\xbe\xd3\xef\x00\x00\x00\xff\xff\xfa\x62\xe7\x50\xdf\x00\x00\x00")

func assetsDnsClusterRoleBindingYamlBytes() ([]byte, error) {
	return bindataRead(
		_assetsDnsClusterRoleBindingYaml,
		"assets/dns/cluster-role-binding.yaml",
	)
}

func assetsDnsClusterRoleBindingYaml() (*asset, error) {
	bytes, err := assetsDnsClusterRoleBindingYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "assets/dns/cluster-role-binding.yaml", size: 223, mode: os.FileMode(420), modTime: time.Unix(1, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xd9, 0xf6, 0x2a, 0x3b, 0x84, 0xd7, 0x3e, 0xc4, 0xe1, 0x70, 0x66, 0x31, 0xda, 0xc4, 0x2f, 0x53, 0x27, 0x29, 0x13, 0xfe, 0x80, 0x36, 0xc5, 0xa1, 0x70, 0xdc, 0x2d, 0xef, 0xcf, 0xe0, 0xc4, 0xeb}}
	return a, nil
}

var _assetsDnsClusterRoleYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x7c\x90\xb1\x4e\xc4\x30\x10\x44\x7b\x7f\x85\x75\xfd\x05\xd1\xa1\xb4\x14\xf4\x14\xf4\x1b\x67\x50\x96\xe4\x76\xad\xdd\x75\x4e\xe2\xeb\x51\x2e\x57\xa0\x8b\xa0\x9b\x19\xd9\xf3\x3c\x9e\x59\xc6\x3e\xbf\x2e\xcd\x03\xf6\xae\x0b\x12\x55\xfe\x80\x39\xab\xf4\xd9\x06\x2a\x1d\xb5\x98\xd4\xf8\x9b\x82\x55\xba\xf9\xc5\x3b\xd6\xa7\xf5\x39\x5d\x10\x34\x52\x50\x9f\x72\x16\xba\xa0\xcf\x5a\x21\x3e\xf1\x67\x9c\x47\xf1\x64\x6d\x81\xf7\xe9\x9c\xa9\xf2\x9b\x69\xab\xbe\x9d\x3c\xe7\xd3\x29\xe5\x6c\x70\x6d\x56\x70\xcf\x20\x63\x55\x96\xf0\x9b\x73\xd8\xca\x05\xbb\xa9\x3a\xee\x62\x63\x78\xa5\x3d\x5f\x61\xc3\xfd\xee\xc2\x1e\x37\x71\xa5\x28\x53\x3a\x02\xb7\x01\x90\xe0\xf2\x7b\xc1\xf1\x0d\xa1\x33\xc4\xb0\x32\xae\x0f\x84\x62\xa0\xc0\x1f\xcd\x8f\x5f\x73\x2c\xf6\x36\x7c\xa1\x04\x95\x02\xf7\xff\x00\x3f\x01\x00\x00\xff\xff\x76\x1b\x55\x2e\x8d\x01\x00\x00")

func assetsDnsClusterRoleYamlBytes() ([]byte, error) {
	return bindataRead(
		_assetsDnsClusterRoleYaml,
		"assets/dns/cluster-role.yaml",
	)
}

func assetsDnsClusterRoleYaml() (*asset, error) {
	bytes, err := assetsDnsClusterRoleYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "assets/dns/cluster-role.yaml", size: 397, mode: os.FileMode(420), modTime: time.Unix(1, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x84, 0xae, 0xd1, 0xba, 0xfa, 0x6b, 0xf8, 0x6e, 0x8d, 0x28, 0xc2, 0xa7, 0xaf, 0xc9, 0x3b, 0xc7, 0xcd, 0x80, 0xbe, 0xec, 0x98, 0xb4, 0x61, 0xa0, 0x9, 0xae, 0xa, 0xd8, 0xb2, 0x2e, 0x16, 0xf2}}
	return a, nil
}

var _assetsDnsDaemonsetYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xcc\x59\x6d\x57\x1b\xb7\xf2\x7f\xcf\xa7\x98\x2e\xfc\x43\xd2\xb0\x60\x92\x90\xe6\xbf\x09\xbd\x75\xc1\x14\x4e\x03\xf8\x60\xa7\x79\xc1\xe1\xfa\xc8\xda\xb1\x57\xd7\x5a\x49\x95\xb4\x6b\xf6\x80\xbf\xfb\x3d\xd2\xfa\x61\xd7\x6b\x68\x73\x9f\x4e\xf3\xc2\xb1\x35\x33\x3f\xcd\x8c\xe6\x49\x62\xc2\x44\x1c\xc1\x29\xc1\x54\x8a\x1e\xda\x2d\xa2\xd8\x6f\xa8\x0d\x93\x22\x02\xa2\x94\x39\xc8\x0f\xb7\xb6\x41\x90\x14\xf7\xfc\xa7\x51\x84\x22\x10\x11\x03\x27\x43\xe4\x06\x88\x46\x30\x68\x81\x58\xd0\x99\xb0\x2c\xc5\x2d\xa3\x90\x46\x5b\x00\x16\x53\xc5\x89\x45\xf7\x1d\x60\xb1\xea\xbf\xa3\xce\x19\xc5\x36\xa5\x32\x13\xf6\x8a\xa4\x18\x41\x2c\xcc\x9c\xaa\x34\x93\x9a\xd9\xe2\x84\x13\x63\x4a\xa2\x29\x8c\xc5\x34\x14\x32\xc6\x90\x6a\x66\x19\x25\x7c\xce\x4d\xa5\xb0\x84\x09\xd4\x66\x81\x1e\x7a\x4d\xab\x88\x00\xdb\xc0\x52\x32\x46\x60\x66\x5d\xdb\x05\x87\xa7\x77\x33\xce\xbb\x92\x33\x5a\x44\x70\x31\xba\x92\xb6\xab\xd1\xa0\xb0\x4b\x2e\x8b\x3a\x65\x82\x58\x26\xc5\x25\x1a\xe3\x44\xe6\xec\x67\x84\xf3\x21\xa1\x93\xbe\xfc\x2c\xc7\xe6\x5a\x74\xb4\x96\x7a\x29\x47\x65\x9a\x12\xe7\xea\x5b\x08\xa8\xd4\x18\x0b\x13\xc0\xdd\x92\x4c\xf4\xd8\x78\x5a\x48\xa5\x18\x05\x7b\x10\x1c\xa0\xa5\x07\x73\xce\x83\x13\xa9\x71\xc4\x38\x56\x45\x72\xc9\xb3\x14\x2f\x9d\x03\x97\x96\xaf\x6c\x77\x30\x6c\x1c\x96\x4c\x4b\x2a\x40\xea\xf8\xbb\xc4\x26\x11\x54\x77\xa8\x70\x68\x24\xf1\xb5\xe0\x45\x04\x56\x67\x2b\x51\x25\x75\x7d\x9f\xa5\xdf\xbb\x52\xdb\x08\x8e\xde\x1e\xbd\xad\xa0\x34\x4f\xc0\x9d\xab\xb4\x92\x4a\x1e\xc1\x97\xd3\xee\xb7\x23\x85\x96\xaa\x8d\x68\xfd\x93\x15\x9a\xd3\x9e\x09\x34\xa6\xab\xe5\x10\xa3\x0a\x7f\x62\xad\xfa\x05\x6d\x75\x09\x40\x95\x9e\x48\x90\x70\x9b\xd4\x29\x5e\x97\x0f\xad\x0f\xad\xda\xb2\xa1\x09\x3a\x7d\xce\xfb\xfd\x6e\x85\xc0\x04\xb3\x8c\xf0\x53\xe4\xa4\xe8\x21\x95\x22\x36\x11\x1c\x56\x45\x15\x6a\x26\xe3\x25\xad\x6a\xa1\xc9\x28\x45\x63\xfa\x89\x46\x93\x48\x1e\x47\x70\x58\xa1\x8e\x08\xe3\x99\xc6\x0a\xb5\x2a\xeb\x42\x58\x66\x76\x03\x2e\x67\x39\xfe\x45\x1c\xf1\xbe\xf5\x8c\xc6\x47\xff\x86\x27\x8e\x2a\xe7\x6e\x64\xa6\x29\x9a\xa8\x16\xca\xbf\x67\x68\xac\xa9\x9b\x4a\x55\x16\xc1\x51\x2b\xad\x2d\xa6\x98\x4a\x5d\x44\xf0\x43\xeb\x92\xad\x95\x91\x49\x36\xc4\x50\x0f\x09\x0d\x95\x96\xf7\xc5\x37\x94\x14\x9f\xd5\x95\x40\x0f\x43\x2e\xc7\x56\x1a\x1b\xa3\xd6\xb5\x75\x83\x34\xd3\x18\x72\x66\x2c\x8a\x90\xc4\xb1\x46\x63\x8e\xa3\xff\x3f\x3c\x7a\x57\xe3\xb3\xdc\x84\x94\xa9\x04\x75\x68\x32\x66\xd1\x1c\xf7\x3f\xf7\x06\x9d\x93\xd3\xf3\xce\xe0\xa6\xd7\x1e\x7c\xbd\xe8\x9f\x0f\xda\x9d\xde\xe0\xf0\xcd\x87\xc1\x2f\x27\x97\x83\xde\x79\xfb\xcd\xd1\xfb\xbd\x15\x57\xe7\xe4\xf4\x0f\xf8\x1a\x38\x27\x3f\x9f\xfc\x29\x9c\x8d\x7c\xcf\xa0\xd5\x2c\xcb\x94\xb1\x1a\x49\x7a\xec\xc2\x33\x3a\x38\x38\x7c\xf3\xc3\x7e\x6b\xbf\xb5\x7f\xe8\x9c\xf0\xf6\xa0\xe9\x05\xd4\x36\x74\x35\xf1\xd8\xd7\x31\xcb\xcd\x81\xd2\x2c\x27\x16\xdd\xf7\x7d\xaa\x6d\x43\x64\x4e\x0f\x27\x58\x3c\x23\x39\xc1\xe2\x4f\x17\xbd\xda\xf9\x2c\x4a\x55\x8a\x56\x33\x6a\xfe\xe5\xd0\x3c\x7c\x22\x34\xdf\xad\x42\xf3\xe9\xea\xbf\x5e\xdf\x2b\xd6\x3d\xa5\xa8\xf3\xcd\x1f\xd5\xff\x4a\x4b\x2d\x9b\xb0\x33\x8a\xe7\xa8\xff\x32\x0d\xd6\x67\x90\x1b\x1a\xa4\xb0\x78\x5f\xab\x6e\xce\x7e\xc6\x71\x8c\xf1\x5a\x4f\x7b\xbe\x85\x26\xd2\x58\xe3\x03\xe5\x99\xfe\xe9\x99\x2a\x4e\x40\x91\xc3\x55\xfb\xb2\xd3\xeb\xdc\xfc\xd6\xb9\xf1\x83\xd2\xc9\xe7\x2f\xbd\x7e\xe7\x66\x70\x7a\x7d\xd9\xbe\xb8\xda\x34\x30\x2d\xc4\x51\xe4\x4d\x35\x1c\xd2\xc5\x49\xa7\x57\x51\x62\x1b\x4e\xdc\x38\x01\x52\x43\x39\x8f\x19\x54\x44\x13\x8b\x31\xb8\x0a\x02\x72\xb4\x98\xb0\x4c\x4d\xea\xea\xba\xdf\x89\xe0\x4c\x6a\x10\x72\xba\x07\x28\x4c\xa6\x11\x6c\x82\x06\xbd\x5a\x1a\x39\xb1\x2c\xc7\x72\xd2\xfb\x08\x23\xa9\x01\x09\x4d\xea\x84\xbd\x1a\x26\x11\x40\x38\x23\x06\xa6\xcc\x26\x0e\x6b\xdd\x5e\x93\x8d\x46\xec\x1e\xa6\x8c\x73\x20\xdc\x48\x18\x22\x90\x38\xc6\x78\xbf\x82\x93\x13\x9e\x61\x04\x81\x8f\x91\x50\xe3\x98\x19\xab\x8b\x7d\xa9\x50\x98\x84\x8d\x6c\xb8\x46\x30\x39\x0d\x1a\xb3\x55\xc5\x75\x07\x43\x26\x0e\x86\xc4\x24\xd5\x22\x40\x2b\x3f\x1e\xab\x46\x7c\xd7\x64\x07\x7f\x46\x61\x26\x41\x31\x85\xae\xf3\x6c\x55\x7b\x98\x26\x0a\x76\xff\x21\x87\x06\x42\x05\x8f\x70\xef\x2a\x3d\x4c\x9c\x89\x8f\x8f\x3e\xc6\x3e\xc2\x94\x30\xfb\x11\xf0\x9e\x59\x68\xed\x42\xbf\x73\x73\x59\x45\xb8\xee\x76\xae\x7a\xe7\x17\x67\xfd\xc1\x65\xfb\xe6\xd7\xce\xcd\x71\xb0\xb2\x75\x8c\x02\xfd\x69\xd6\x53\x2d\xa8\x88\x9f\x5f\xf7\xfa\xbd\xc1\xd9\xc5\xe7\xce\x71\xb0\x8a\xc3\x2a\x47\xbf\x73\xd9\x6d\x30\xec\xdb\x54\x05\x55\x35\x2e\xce\x7a\xc7\xbb\x7b\xb0\xeb\xb3\x1e\x42\x0d\x21\x59\x86\x0e\x7c\xfa\xf4\x09\x82\x9d\x87\x45\x00\xce\x6a\x92\xdb\x70\x49\x26\x08\xc4\x4f\xf9\x52\x13\x5d\x80\x4b\x95\x55\x18\x48\x1e\x97\x29\xe4\xd7\x77\x0d\x10\x6b\x35\x1b\x66\x16\x4d\xf5\xe4\xa9\x82\x70\x04\x61\xb8\xa2\x86\x52\xf0\xc2\x6d\xbc\x32\x72\x16\xb8\xdf\x4b\x93\xea\x9a\x4c\x13\xb7\x6f\xe9\xf4\x58\xd6\x4a\x67\x8c\x94\xbb\xc0\x0e\xdb\x60\x72\x3a\x60\xca\xd4\xc8\x2e\xbe\x4d\x4e\x81\x09\x07\xbf\xb0\xfb\xf6\xa7\xbb\x59\xd0\x80\x72\x16\x9f\xa1\xa5\xc9\xc2\x3f\x70\xd1\x85\x91\x96\x29\x50\x9e\x19\x8b\xda\xd5\x46\x60\x23\x50\x65\x41\xdb\x87\xaf\x08\xa9\x73\x91\xc1\x1c\x35\xe1\x60\x35\x43\xd3\xc0\xb4\x12\x62\x09\xcc\x46\x70\xd1\xcd\xdf\xed\xb9\xcf\xf7\xfe\xf3\x1d\xc8\x1c\xb5\x1b\x6e\x7d\x15\x71\xeb\xcb\x95\x7d\xe8\x27\x08\x76\x2a\x81\x13\x97\xef\x62\x03\xb0\xb3\xdb\x19\x18\xa3\xe2\xb2\x48\x51\xd8\x79\x8e\xfe\x9a\xe9\x42\x83\x14\xee\x84\x50\xc3\xb5\x42\xd1\xb3\x84\x4e\xe0\xe5\x75\xaf\x7b\xf8\xf6\x15\x84\x60\x13\x69\xd0\xe9\x25\xa4\x6d\x00\x9b\x4c\xb9\xbe\xe8\x86\x78\xe0\x92\xc4\x43\xc2\x89\xa0\xa8\x8d\xd7\xd3\x35\x36\xe6\x6b\x09\xa1\x09\x13\x63\x38\xbd\xea\x81\x4d\xb4\xcc\xc6\x89\x57\x7d\x0d\x8f\xa6\xb1\x39\x7e\xb9\x1b\xb3\x31\x84\x16\xda\xf0\x53\xb0\xf3\xb0\x2a\xa0\xb3\x00\x5e\x9b\xc4\xed\xe6\x0e\x28\xa7\xb3\xfd\x9d\x87\x7a\x7d\x99\x05\x8f\x63\x8d\x0a\xc2\x1c\x82\xbf\x7f\x0c\x76\xd7\xe0\xcb\x7f\x4b\xf8\x76\xfb\xbf\xbd\x03\xbc\xb6\x54\xc1\x6b\x8d\x56\x17\xc7\xad\xff\x81\x39\xff\xd9\xfd\x5e\xad\x6d\xe8\x22\x88\xb9\x04\xd9\x79\xf8\xce\x1d\xd5\xed\xf7\x77\xb3\x35\x96\x46\xa2\x00\x30\x65\x8e\x5f\xee\xbc\xc4\x9c\x70\xb7\xb3\x17\x64\x77\xb3\xe0\xd5\x3a\x3c\xb8\x8c\xb9\xbd\x85\x60\xe7\x6f\x01\x84\xf8\x3b\xb4\xe0\xc5\x0b\x27\xb2\xcd\x54\x99\x88\x10\x0a\x84\x16\xdc\xdd\x7d\x74\x55\x45\x6c\xf0\xc7\x3c\xb3\x6f\xe7\x26\x06\x77\xc7\xc1\xce\xc3\x42\x7c\x03\xff\x50\x23\x99\x34\xd6\x47\xac\x61\x96\xc0\xad\xc6\x42\x6d\x65\x1b\xbe\xa8\x98\x58\xac\x8c\x02\xe0\x8b\x17\x1b\xc1\x14\x61\x8c\xd6\x35\x36\x16\x57\x4a\x86\x59\x03\xf8\x8a\x65\x67\x14\xd2\x42\xd6\x00\x9b\x26\x28\x9c\xd9\xda\xcf\x55\xf3\xbb\xfa\x12\x4d\x66\xd6\x4d\x5c\x52\x03\x51\x0c\x32\x41\x72\xc2\x38\x19\x32\xce\x6c\xb1\xb6\x4d\xcf\x12\x8e\x80\xc2\xd7\x20\xa0\x32\xe3\xb1\x6b\x4d\xc6\xba\xa3\xad\x6c\xc8\x46\xbe\x76\x2f\x76\x60\x06\x62\xe4\x68\x31\xde\x6a\x9e\x59\x28\xe6\x51\xe5\xbd\xff\xfd\x5d\x38\x0b\x9e\x3a\xa6\x6d\xf8\x39\x63\x3c\x06\x02\x02\xa7\x95\xae\x50\x16\xd0\xaa\xc1\xae\x40\xc9\x4c\x03\xcd\x8c\x95\xe9\x52\xe3\x11\xe3\x16\x35\xc6\xce\xe6\x35\xec\x65\xf8\x6e\xc3\xce\xc3\x7a\x5b\x2d\x1b\x47\xad\x91\xfc\xf8\x4c\x2b\x29\x75\x6d\x2b\x85\xbe\x92\x95\x7d\x77\xa5\x84\x6b\x17\xcd\xb9\x0a\x1a\x9d\xe4\xbb\x85\x53\x9e\xe8\x24\xf3\xb4\x52\x65\x5e\x2d\x98\xcb\xf0\xbd\x9b\x6d\x14\x00\x40\x9a\x48\xf0\x91\x3d\x2b\x85\x16\xff\x35\x73\x1a\x9e\x70\xc5\x8f\x0d\xdb\xd7\x37\x69\x04\xfd\xa6\xb0\x77\x3e\xea\x5f\x9f\x5e\x47\x1b\xc2\x9f\x58\x99\x32\x4a\x38\x2f\x5c\x67\x23\xb9\x64\x31\x10\x51\x00\x13\x54\x0a\xe3\xaf\xb7\x16\x86\x98\x90\x9c\x55\x86\xf7\x05\xea\x0d\x2a\xee\xe6\xd9\x4d\x11\x91\xca\x98\x8d\x18\xc6\x90\x97\xef\x93\x2e\x0a\x05\x62\xbc\x16\x9b\xae\xa3\xa8\x35\x33\x1b\x31\xf0\xf8\x38\x9f\x3b\x9e\xe7\x6b\x5a\xbd\xe0\x75\x99\xe1\x52\x56\x63\x2a\x73\x8c\x57\xb6\xfa\xa8\xa6\x1a\xdd\x6d\xb2\x4c\x1d\xdf\x15\x57\xd3\x0d\x50\xa9\x0a\xa0\x49\xa6\xeb\x49\xb2\x56\x7f\x0c\x47\x54\xf0\xbe\x05\x2f\xfc\x20\x59\xa3\x65\xc2\xcd\xa6\xcd\x81\xa6\x76\x78\xdf\xfc\x20\xb2\xf9\xd2\xf9\xe6\x70\x79\xe9\x8c\x85\x59\x5c\xc5\x4e\x71\x44\x32\xbe\xd0\xca\x4d\xa9\x3d\xe4\x48\xad\xd4\x2b\xe4\x49\x36\x44\x2d\xd0\x8d\x7b\x4c\x1e\x48\x13\x01\x67\x22\xbb\x2f\x89\x73\xae\xf2\x02\xd6\x78\xb8\xdd\xfc\x78\x59\xae\x5e\x12\x15\x55\xee\x5b\x57\x24\x7d\xee\xce\x09\xc0\x2c\xa6\x35\x7b\x43\x98\x60\x11\xc1\xe2\x49\x75\xc3\x2b\xd8\x1a\xe9\x99\xfb\xa0\x5b\xf2\x97\xc1\xad\x75\x8c\x0d\x97\x43\x00\x5b\x28\x8c\xe0\xac\x09\xbd\xe9\x26\xbe\xed\xae\xb4\x1a\xed\xb3\x16\x5a\xc9\xdd\x55\x81\x49\xb1\xb4\x71\xdb\x4f\x5c\x2e\x33\x8c\x0b\x4b\x9d\x09\x70\x03\x68\x31\x75\x6d\x64\x1f\xfa\xa5\x04\x02\xe1\x1c\x2c\x61\x62\xa9\x61\x08\x52\x39\x92\xd4\x11\x74\x5c\x6f\x70\x84\xb2\x27\xf5\xac\x13\x19\x17\xe5\x1e\xa5\x19\x37\x92\x73\x26\xc6\x65\x09\xf0\xeb\xba\xba\xb2\x52\xe7\x4a\x5a\x8c\xfc\xc0\x1a\xfb\x3f\x33\xf8\x47\x14\xc7\x8b\x1a\xb4\xcc\x84\xd3\x33\x41\x50\xa8\x29\x0a\xdf\xd1\x32\xb5\x14\x7e\x99\x09\xce\x26\xfe\x92\x5a\x99\x64\x2b\x10\x7b\x6e\xfe\x77\x57\xd4\x12\x29\x96\x53\xf1\x6a\x31\x63\xa6\xe4\xfe\xcb\xa2\x2b\x72\x8c\xe0\xb0\xf5\x7f\x5b\xff\x0c\x00\x00\xff\xff\x6f\x9b\x9e\xf3\xf1\x18\x00\x00")

func assetsDnsDaemonsetYamlBytes() ([]byte, error) {
	return bindataRead(
		_assetsDnsDaemonsetYaml,
		"assets/dns/daemonset.yaml",
	)
}

func assetsDnsDaemonsetYaml() (*asset, error) {
	bytes, err := assetsDnsDaemonsetYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "assets/dns/daemonset.yaml", size: 6385, mode: os.FileMode(420), modTime: time.Unix(1, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x18, 0x97, 0xef, 0xef, 0x5d, 0xc1, 0xe9, 0x24, 0x25, 0x22, 0x38, 0x44, 0xb6, 0xab, 0x59, 0xe1, 0x68, 0xa6, 0x24, 0x3e, 0x82, 0xd4, 0x45, 0x40, 0xe4, 0xe6, 0x8b, 0x9f, 0x96, 0xed, 0x9d, 0x68}}
	return a, nil
}

var _assetsDnsMetricsClusterRoleBindingYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x7c\x8f\xb1\x4a\x04\x41\x0c\x86\xfb\x79\x8a\xbc\xc0\xae\xd8\x1d\xd3\xa9\x85\xfd\x09\xf6\xb9\x99\x9c\x1b\x77\x27\x19\x92\xcc\x16\x3e\xbd\x2c\x8a\x08\xe2\xb5\x81\x7c\xdf\xff\xad\x2c\x35\xc3\xd3\x36\x3c\xc8\xce\xba\xd1\x23\x4b\x65\x79\x4b\xd8\xf9\x95\xcc\x59\x25\x83\x5d\xb0\xcc\x38\x62\x51\xe3\x0f\x0c\x56\x99\xd7\x93\xcf\xac\x77\xfb\x7d\x6a\x14\x58\x31\x30\x27\x00\xc1\x46\x19\xaa\xf8\xd4\x54\x38\xd4\x0e\x92\x8f\xcb\x3b\x95\xf0\x9c\x26\xf8\xd2\xbd\x90\xed\x5c\xe8\xa1\x14\x1d\x12\x3f\x7f\xdd\xb4\x51\x2c\x34\x7c\x5a\x4f\xfe\x7d\xf6\x8e\x85\x32\x68\x27\xf1\x85\xaf\xf1\x9b\x6c\xba\xd1\x99\xae\x87\xf9\x4f\xc7\x7f\x6b\x00\xb0\xf3\xb3\xe9\xe8\x37\xba\xd2\x67\x00\x00\x00\xff\xff\x5b\x52\x00\xaa\x17\x01\x00\x00")

func assetsDnsMetricsClusterRoleBindingYamlBytes() ([]byte, error) {
	return bindataRead(
		_assetsDnsMetricsClusterRoleBindingYaml,
		"assets/dns/metrics/cluster-role-binding.yaml",
	)
}

func assetsDnsMetricsClusterRoleBindingYaml() (*asset, error) {
	bytes, err := assetsDnsMetricsClusterRoleBindingYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "assets/dns/metrics/cluster-role-binding.yaml", size: 279, mode: os.FileMode(420), modTime: time.Unix(1, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x79, 0x95, 0x6f, 0xa4, 0xd5, 0xed, 0x48, 0x27, 0x41, 0x56, 0x5c, 0xea, 0x5c, 0x89, 0xdc, 0xc1, 0x44, 0x91, 0xd4, 0xb, 0x18, 0x85, 0x79, 0x75, 0xaa, 0x6e, 0xb5, 0x98, 0xbe, 0xc6, 0x33, 0x43}}
	return a, nil
}

var _assetsDnsMetricsClusterRoleYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x34\xcd\x31\x4b\x34\x41\x0c\x87\xf1\x7e\x3e\x45\xe0\xad\x77\x5f\xec\x64\x5a\x05\x3b\x0b\x05\xfb\xec\xce\xdf\xdb\x70\x3b\xc9\x90\x64\x0e\xf4\xd3\x8b\x70\xb6\x0f\x3f\x78\xfe\xd1\xd3\x39\x23\xe1\xe4\x76\x22\x48\x81\x86\x46\xdb\x17\x0d\xb7\x8e\x3c\x30\x83\xd2\x28\x76\xe7\x01\x7a\x7e\x7d\xa7\x8e\x74\xd9\x83\xa0\x6d\x98\x68\x16\x1e\xf2\x01\x0f\x31\xad\xe4\x1b\xef\x2b\xcf\x3c\xcc\xe5\x9b\x53\x4c\xd7\xeb\x63\xac\x62\xff\x6f\x0f\xe5\x2a\xda\xea\xdf\xf0\xcd\x4e\x94\x8e\xe4\xc6\xc9\xb5\x10\x29\x77\x54\x6a\x1a\x4b\x37\x95\x34\x17\xbd\x14\x9f\x27\xa2\x96\x85\x78\xc8\x8b\xdb\x1c\xf1\x4b\x17\xb2\x01\xe7\x34\x5f\x6d\x40\xe3\x90\xcf\x5c\xc5\x0a\x91\x23\x6c\xfa\x8e\x3b\x6b\x1a\x88\x42\x74\x83\x6f\xf7\x74\x41\x96\x9f\x00\x00\x00\xff\xff\x9f\xa8\x4d\x6c\xf6\x00\x00\x00")

func assetsDnsMetricsClusterRoleYamlBytes() ([]byte, error) {
	return bindataRead(
		_assetsDnsMetricsClusterRoleYaml,
		"assets/dns/metrics/cluster-role.yaml",
	)
}

func assetsDnsMetricsClusterRoleYaml() (*asset, error) {
	bytes, err := assetsDnsMetricsClusterRoleYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "assets/dns/metrics/cluster-role.yaml", size: 246, mode: os.FileMode(420), modTime: time.Unix(1, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x64, 0xdb, 0xe0, 0x95, 0x65, 0xae, 0x53, 0x96, 0x3a, 0x5f, 0x5e, 0x8b, 0x69, 0xe2, 0x7d, 0x5, 0xbf, 0x1f, 0x3a, 0xf, 0xff, 0xd0, 0x6b, 0x23, 0x4f, 0xfd, 0x11, 0x7f, 0x57, 0xd4, 0x4a, 0x8b}}
	return a, nil
}

var _assetsDnsMetricsRoleBindingYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x94\xce\xb1\x4e\xc4\x40\x0c\x04\xd0\x7e\xbf\xc2\x3f\x90\x20\xba\xd3\x76\xd0\xd0\x1f\x12\xbd\x6f\xd7\x97\x98\x64\xed\x95\xed\x4d\xc1\xd7\x23\xa4\x48\x54\x20\x5d\x3b\x9a\xd1\x1b\xec\xfc\x41\xe6\xac\x92\xc1\x6e\x58\x66\x1c\xb1\xaa\xf1\x17\x06\xab\xcc\xdb\xc5\x67\xd6\xa7\xe3\x39\x6d\x2c\x35\xc3\x55\x77\x7a\x65\xa9\x2c\x4b\x6a\x14\x58\x31\x30\x27\x00\xc1\x46\x19\xba\x69\xa3\x58\x69\xf8\xb4\x5d\xfc\x8c\xbd\x63\xa1\x0c\xda\x49\x7c\xe5\x7b\x4c\x55\x3c\x99\xee\x74\xa5\xfb\xcf\x14\x3b\xbf\x99\x8e\xfe\x8f\x9f\x00\x7e\xf9\xbf\x34\x1f\xb7\x4f\x2a\xe1\x39\x4d\x67\xfb\x9d\xec\xe0\x42\x2f\xa5\xe8\x90\x78\xf0\x65\x53\xe1\x50\x63\x59\x20\x7d\x07\x00\x00\xff\xff\xb9\xd9\xab\x8d\x25\x01\x00\x00")

func assetsDnsMetricsRoleBindingYamlBytes() ([]byte, error) {
	return bindataRead(
		_assetsDnsMetricsRoleBindingYaml,
		"assets/dns/metrics/role-binding.yaml",
	)
}

func assetsDnsMetricsRoleBindingYaml() (*asset, error) {
	bytes, err := assetsDnsMetricsRoleBindingYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "assets/dns/metrics/role-binding.yaml", size: 293, mode: os.FileMode(420), modTime: time.Unix(1, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xc, 0x7d, 0xc7, 0x45, 0x33, 0xc4, 0xd8, 0xf, 0x8d, 0x89, 0x8d, 0x6, 0x47, 0xa7, 0xa, 0x6b, 0x17, 0xf5, 0x5f, 0x5a, 0x2f, 0xd8, 0xf9, 0x6, 0x71, 0xaa, 0x78, 0x8d, 0xb5, 0x7a, 0xf6, 0x99}}
	return a, nil
}

var _assetsDnsMetricsRoleYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x4c\x8e\xb1\x4e\xec\x40\x0c\x45\xfb\xf9\x0a\x6b\x5f\x9d\x7d\xa2\x5b\x4d\x8d\x44\x47\x01\x12\xbd\x77\xe6\x42\xac\x24\xe3\x91\xed\x04\xc1\xd7\xa3\xcd\x06\x89\xca\xf7\x1e\x59\x3e\xfe\x47\x2f\x3a\xc3\xa9\x01\x15\x95\xae\x5f\xd4\x4d\x17\xc4\x88\xd5\x29\x94\xbc\x18\x77\xd0\xe3\xf3\x2b\x2d\x08\x93\xe2\x84\x56\xbb\x4a\x8b\xc4\x5d\xde\x60\x2e\xda\x32\xd9\x95\xcb\x99\xd7\x18\xd5\xe4\x9b\x43\xb4\x9d\xa7\x8b\x9f\x45\xff\x6f\x0f\x69\x92\x56\xf3\x2e\x4a\x0b\x82\x2b\x07\xe7\x44\xd4\x78\x41\xfe\xe3\x1b\xa6\x8b\x1f\xd8\x3b\x17\x64\xd2\x8e\xe6\xa3\xbc\xc7\x50\x9b\x27\x5b\x67\x78\x4e\x03\x71\x97\x27\xd3\xb5\xfb\xed\xca\x40\xa7\x53\x22\x32\xb8\xae\x56\x70\x30\x87\x6d\x52\xe0\x7b\xf9\xfd\xf8\xde\xba\xd6\x5b\xd8\x60\xd7\x63\xf9\x03\xb1\xcf\x59\xfc\x1e\x3e\x39\xca\x98\x7e\x02\x00\x00\xff\xff\x29\x39\xda\x05\x1c\x01\x00\x00")

func assetsDnsMetricsRoleYamlBytes() ([]byte, error) {
	return bindataRead(
		_assetsDnsMetricsRoleYaml,
		"assets/dns/metrics/role.yaml",
	)
}

func assetsDnsMetricsRoleYaml() (*asset, error) {
	bytes, err := assetsDnsMetricsRoleYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "assets/dns/metrics/role.yaml", size: 284, mode: os.FileMode(420), modTime: time.Unix(1, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x8c, 0xf2, 0x4e, 0x40, 0x91, 0xd8, 0x5e, 0x1c, 0x98, 0xb6, 0x2f, 0x11, 0x2a, 0x15, 0x8f, 0xe4, 0x7c, 0xfe, 0xc6, 0x31, 0xf3, 0xb2, 0xa0, 0x38, 0xb2, 0x3f, 0x15, 0x5a, 0x33, 0x12, 0xd2, 0x88}}
	return a, nil
}

var _assetsDnsNamespaceYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x64\x90\xcd\x4e\xc4\x30\x0c\x84\xef\x79\x8a\x51\x38\x2f\x3f\xd7\xbc\x03\x5c\x90\xb8\xbb\x8d\x97\x35\x4d\xed\x2a\x76\xcb\xeb\xa3\xb2\x15\xac\xb4\xc7\x68\x46\xf3\x7d\xf1\x24\x5a\x0b\xde\x68\x66\x5f\x68\xe4\x44\x8b\x7c\x70\x77\x31\x2d\xd8\x5e\xd2\xcc\x41\x95\x82\x4a\x02\x48\xd5\x82\x42\x4c\x7d\x7f\x02\xb6\xb0\xfa\x45\xce\xf1\x28\xf6\xa4\x56\xf9\xe4\xdc\x78\x0c\xeb\x05\x39\x27\x40\x69\xe6\xf2\x5f\x3b\x55\xf5\x04\x34\x1a\xb8\x1d\x13\x0f\x70\x0e\x6c\xd4\x56\x46\x18\x68\x33\xa9\xa8\xbc\xb0\x56\xd1\x4f\x98\x62\x5a\x07\x06\xd5\x59\x7c\x97\x42\x5c\x28\x8e\x82\xef\xf1\xdf\x38\x68\x11\xbf\xd7\xea\xab\x9e\x1a\x6f\xdc\x0a\xf2\x73\x3e\x98\xd4\x9a\x7d\xdf\x78\xcd\xa6\x12\xd6\x77\x62\x18\x9a\xd9\x84\xb3\x75\xbc\x73\xdf\x64\xe4\xd7\x6b\x0a\x1b\xbe\x78\x0c\x87\xec\x16\xe2\xbf\xbf\xbb\x1e\xed\x8e\x3a\xb6\xd5\x83\xfb\xcd\x70\x41\x8e\xbe\x72\x4e\x3f\x01\x00\x00\xff\xff\x82\x6d\x29\x03\x71\x01\x00\x00")

func assetsDnsNamespaceYamlBytes() ([]byte, error) {
	return bindataRead(
		_assetsDnsNamespaceYaml,
		"assets/dns/namespace.yaml",
	)
}

func assetsDnsNamespaceYaml() (*asset, error) {
	bytes, err := assetsDnsNamespaceYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "assets/dns/namespace.yaml", size: 369, mode: os.FileMode(420), modTime: time.Unix(1, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xe, 0xab, 0x50, 0x84, 0x61, 0x5f, 0x41, 0xf4, 0x17, 0x3b, 0x6, 0x84, 0xc0, 0x5f, 0x4f, 0xbb, 0xd8, 0x1d, 0xae, 0x26, 0x3e, 0x1f, 0x29, 0x2c, 0x84, 0x6d, 0x5e, 0xc1, 0x87, 0x97, 0x5f, 0xc9}}
	return a, nil
}

var _assetsDnsServiceAccountYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x2c\xc9\xb1\x09\xc4\x30\x0c\x05\xd0\xde\x53\x68\x81\x2b\xae\x55\x77\x33\x1c\xa4\x17\xf2\x0f\x11\xc1\xb2\xb1\x14\xcf\x1f\x02\xe9\x1e\xbc\xd3\xbc\x32\xfd\x31\x97\x29\x7e\xaa\xfd\xf2\x2c\x32\x6c\xc3\x0c\xeb\xce\xb4\xbe\xa5\x21\xa5\x4a\x0a\x17\x22\x97\x06\xa6\xea\xf1\x3a\x86\x28\x98\xfa\x80\xc7\x61\x7b\x7e\x9e\xba\x03\x00\x00\xff\xff\x8e\x2c\xf1\x2e\x55\x00\x00\x00")

func assetsDnsServiceAccountYamlBytes() ([]byte, error) {
	return bindataRead(
		_assetsDnsServiceAccountYaml,
		"assets/dns/service-account.yaml",
	)
}

func assetsDnsServiceAccountYaml() (*asset, error) {
	bytes, err := assetsDnsServiceAccountYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "assets/dns/service-account.yaml", size: 85, mode: os.FileMode(420), modTime: time.Unix(1, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x57, 0x12, 0x50, 0x4d, 0x67, 0x2f, 0x1b, 0x74, 0xa0, 0xa4, 0xbb, 0xa7, 0x59, 0xe9, 0x5a, 0xc6, 0xc1, 0x1a, 0xf8, 0x5f, 0xff, 0x5, 0xdb, 0xc, 0x10, 0x8b, 0xc1, 0x0, 0xcc, 0xf, 0x9f, 0x3a}}
	return a, nil
}

var _assetsDnsServiceYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x84\x91\x31\x6f\xe2\x40\x10\x85\x7b\xff\x8a\x27\xe8\x4e\xc0\x09\xdd\x51\x9c\xdb\xa3\x89\x52\x80\x14\x48\x3f\x5e\x4f\xcc\x8a\xf5\x8c\xb5\x33\x06\xf1\xef\x23\x4c\x42\x80\x14\x69\x56\xda\x7d\x9f\x3e\x3d\xbd\xdd\x47\xa9\x4b\xbc\x70\x3e\xc4\xc0\x05\x75\xf1\x95\xb3\x45\x95\x12\x87\x79\x31\x86\x50\xcb\x93\xe1\xb4\x8e\x02\x4f\x12\x55\x9c\x0c\x24\x35\x48\x44\x9d\x3c\xaa\x18\x28\x33\x8c\x1d\xe4\xc8\xbd\x78\x6c\xb9\xb0\x8e\x43\x59\x00\x63\x84\xd4\x9b\x73\x7e\x5a\xe3\x18\x53\x42\xc5\xa0\xde\xb5\x25\x8f\x81\x52\x3a\xa1\x25\xa1\x86\xeb\xd9\x00\x1b\x27\x0e\xae\x19\xd1\x1e\x8d\x40\xa7\xd9\xed\x2c\x9d\x0e\x95\x4a\xd4\x62\x05\x70\x09\x4a\x2c\xfe\x0c\x17\xa7\xdc\xb0\xaf\x87\xa7\x2b\x90\xd5\x35\x68\x2a\xb1\x5d\xae\xef\x05\x53\x0f\xdd\x8f\x92\x2f\xe8\x2a\xda\xfc\xbf\x15\xb5\xec\x39\x86\xdb\x36\xff\xe6\x8b\xbf\xdf\x54\x77\xd8\x83\x6a\x8c\xcd\x6a\xb9\x2a\xb1\x95\xa0\x6d\xcb\xe2\x38\xee\x58\x60\x97\xbf\x81\x6b\xa7\x49\x9b\x13\xde\x98\xbc\xcf\x8c\x86\x9c\xcf\x33\xb1\x50\x95\x3e\xf6\xfb\x84\x9e\xf9\x64\x97\xf5\x31\xc5\x68\xdf\x57\x9c\x85\x9d\x6d\x16\xf5\xf7\x4e\xcd\xcf\xa5\x47\xd7\xfc\xd7\xa8\x78\x0f\x00\x00\xff\xff\x82\x42\x75\xa4\x08\x02\x00\x00")

func assetsDnsServiceYamlBytes() ([]byte, error) {
	return bindataRead(
		_assetsDnsServiceYaml,
		"assets/dns/service.yaml",
	)
}

func assetsDnsServiceYaml() (*asset, error) {
	bytes, err := assetsDnsServiceYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "assets/dns/service.yaml", size: 520, mode: os.FileMode(420), modTime: time.Unix(1, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x18, 0x69, 0xc5, 0xf1, 0xe, 0xc, 0x77, 0xe5, 0x78, 0xce, 0xfc, 0xc2, 0x41, 0xf8, 0x21, 0x87, 0x8a, 0xb7, 0x67, 0xdd, 0x48, 0x94, 0x63, 0x79, 0x69, 0x4e, 0x38, 0x53, 0x3c, 0xdb, 0xc7, 0x13}}
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
	"assets/dns/cluster-role-binding.yaml": assetsDnsClusterRoleBindingYaml,

	"assets/dns/cluster-role.yaml": assetsDnsClusterRoleYaml,

	"assets/dns/daemonset.yaml": assetsDnsDaemonsetYaml,

	"assets/dns/metrics/cluster-role-binding.yaml": assetsDnsMetricsClusterRoleBindingYaml,

	"assets/dns/metrics/cluster-role.yaml": assetsDnsMetricsClusterRoleYaml,

	"assets/dns/metrics/role-binding.yaml": assetsDnsMetricsRoleBindingYaml,

	"assets/dns/metrics/role.yaml": assetsDnsMetricsRoleYaml,

	"assets/dns/namespace.yaml": assetsDnsNamespaceYaml,

	"assets/dns/service-account.yaml": assetsDnsServiceAccountYaml,

	"assets/dns/service.yaml": assetsDnsServiceYaml,
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
	"assets": {nil, map[string]*bintree{
		"dns": {nil, map[string]*bintree{
			"cluster-role-binding.yaml": {assetsDnsClusterRoleBindingYaml, map[string]*bintree{}},
			"cluster-role.yaml":         {assetsDnsClusterRoleYaml, map[string]*bintree{}},
			"daemonset.yaml":            {assetsDnsDaemonsetYaml, map[string]*bintree{}},
			"metrics": {nil, map[string]*bintree{
				"cluster-role-binding.yaml": {assetsDnsMetricsClusterRoleBindingYaml, map[string]*bintree{}},
				"cluster-role.yaml":         {assetsDnsMetricsClusterRoleYaml, map[string]*bintree{}},
				"role-binding.yaml":         {assetsDnsMetricsRoleBindingYaml, map[string]*bintree{}},
				"role.yaml":                 {assetsDnsMetricsRoleYaml, map[string]*bintree{}},
			}},
			"namespace.yaml":       {assetsDnsNamespaceYaml, map[string]*bintree{}},
			"service-account.yaml": {assetsDnsServiceAccountYaml, map[string]*bintree{}},
			"service.yaml":         {assetsDnsServiceYaml, map[string]*bintree{}},
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
	err = ioutil.WriteFile(_filePath(dir, name), data, info.Mode())
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
