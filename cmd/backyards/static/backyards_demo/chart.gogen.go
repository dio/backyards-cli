// Code generated by vfsgen; DO NOT EDIT.

package backyards_demo

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	pathpkg "path"
	"time"
)

// Chart statically implements the virtual filesystem provided to vfsgen.
var Chart = func() http.FileSystem {
	fs := vfsgen۰FS{
		"/": &vfsgen۰DirInfo{
			name:    "/",
			modTime: time.Date(2019, 1, 1, 0, 1, 0, 0, time.UTC),
		},
		"/.helmignore": &vfsgen۰CompressedFileInfo{
			name:             ".helmignore",
			modTime:          time.Date(2019, 1, 1, 0, 1, 0, 0, time.UTC),
			uncompressedSize: 342,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x4c\x8e\xc1\x6a\x23\x31\x0c\x86\xef\x7a\x8a\x7f\x99\xcb\xee\xb0\x78\x1e\x22\xd9\xc3\x9e\x5a\x48\xc9\xb5\x78\x66\x14\x5b\x89\xc7\x36\x96\x26\x69\x7b\xe8\xb3\x97\x24\x84\xf6\xf2\x81\x3e\x24\xf1\x75\x78\xf6\x66\xdc\xb2\xc2\x0a\x24\xe4\xd2\x18\x97\xc8\x19\xe3\x2a\x69\x96\x1c\x50\xfd\x74\xf2\x81\xd5\x51\x87\x97\x28\x0a\x5d\x6b\x2d\xcd\x14\x1a\x39\x25\x84\x54\x46\x2c\xde\xa6\x28\x39\xfc\x45\xe3\xe4\x4d\xce\x8c\xea\x2d\xfe\xf0\x3e\xcf\xd4\x21\x73\xf0\x26\x25\xe3\x77\x6d\x7c\x90\x37\x9e\x71\x11\x8b\xf8\xf5\xc7\xe1\x29\xa7\x77\x94\x7c\xbb\xbc\x26\xa1\x72\x43\x92\xcc\x8e\xdc\x76\xf7\xba\xb3\xd2\x98\x3a\x6c\xca\xb2\x94\x8c\xfd\x66\x87\x59\x9a\x92\x0b\x62\xc3\x8d\xf7\x7c\x72\xe3\x47\x1b\x6e\x7c\x88\x18\x86\x2b\x1e\xa3\x9e\xf3\xf0\xfd\x68\xf4\xd3\x69\xad\x38\x48\x62\xa5\xde\xe9\xa5\x52\xef\x46\x7f\xa2\xde\xd9\x52\xa9\xff\xa4\x0e\x7b\xdf\xa4\xac\x8a\xff\xdb\x7f\x4a\xae\xb6\x72\xe4\xc9\xc8\xc9\xcc\x7e\xb8\xef\xb5\x72\x24\x77\xd6\xa9\xcc\x3c\xd0\x57\x00\x00\x00\xff\xff\xf5\x89\xaa\x2d\x56\x01\x00\x00"),
		},
		"/Chart.yaml": &vfsgen۰CompressedFileInfo{
			name:             "Chart.yaml",
			modTime:          time.Date(2019, 1, 1, 0, 1, 0, 0, time.UTC),
			uncompressedSize: 255,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x74\xcf\x31\x4e\xc5\x30\x0c\x06\xe0\xdd\xa7\xb0\xba\x27\x6d\xd7\x4c\xa8\xdc\x81\xdd\x4d\x42\x6b\xd1\xc4\x56\x12\x2a\xc1\xe9\x51\x0a\x65\x78\xd2\x1b\xfd\xfd\xb2\xfc\x9b\x94\xdf\x62\xa9\x2c\xd9\xe1\x39\x03\xa9\xfe\x8f\xc3\x3c\x40\x88\xd5\x17\xd6\x76\xc1\x42\xfe\xe3\x8b\x4a\xa8\x18\x62\x12\x24\x55\xf4\x3b\x95\x06\xbb\xa4\xe8\x70\x6f\x4d\xab\x1b\xc7\x95\xf2\x37\xb1\x3f\xe4\x33\x58\x2f\x09\xd8\xf7\xe5\x27\xe9\xc8\x69\xfb\x33\x73\xa1\x39\x64\x13\xab\x79\x83\x44\x9c\x1b\x71\x8e\xa5\x3a\x30\x18\x13\xf1\xe1\x90\xf3\xbb\xbc\x3c\x9e\x40\xcc\xd4\x2b\x2c\x97\xe3\x6b\x0f\xe0\x97\xd6\xbb\xb4\xe9\xa5\xe1\xbc\xbf\x9b\xec\x6c\x27\xf8\x09\x00\x00\xff\xff\x84\x30\xc1\x03\xff\x00\x00\x00"),
		},
		"/templates": &vfsgen۰DirInfo{
			name:    "templates",
			modTime: time.Date(2019, 1, 1, 0, 1, 0, 0, time.UTC),
		},
		"/templates/_helpers.tpl": &vfsgen۰CompressedFileInfo{
			name:             "_helpers.tpl",
			modTime:          time.Date(2019, 1, 1, 0, 1, 0, 0, time.UTC),
			uncompressedSize: 1522,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x9c\x94\x4f\x6b\xdb\x40\x10\xc5\xef\xfa\x14\x83\x68\xa0\x4d\x91\x7c\x28\xf4\x60\xc8\xa1\xa4\x3d\x94\x42\x0a\x0d\xa4\xc7\x32\x92\x66\xa3\x21\xab\xdd\xcd\xfe\x71\x63\x1c\x7f\xf7\xb2\xbb\xb2\x6c\x37\x76\xed\xf4\xb6\xc8\x6f\xde\xbc\xf9\xcd\xae\x57\xab\xd9\x25\x2c\x78\x98\x83\x23\x0f\x82\x25\xf9\xa5\xa1\xab\x21\x38\x8f\x6d\x4f\x73\xb8\x9c\xad\xd7\x45\x54\x15\x5f\x9e\x0c\xaa\x0e\x7c\x4f\xa0\x70\x20\xd0\x22\x9d\xdb\x1e\xad\xaf\x8b\x51\x57\x41\x47\x82\x15\x41\xd9\x60\xfb\xb0\x44\xdb\xb9\xaa\xa3\x41\xd7\xb1\xa4\x84\x6a\x2b\xc2\x20\x3d\xd4\xd7\xa9\xfa\x26\xfa\xd5\x77\x28\x03\xb9\xa4\xfc\xbe\x20\x6b\xb9\x23\x78\x06\x6f\x83\x6a\xe1\xe3\x87\x74\xe4\xe1\x36\x08\xc1\x4f\x50\x56\x5b\x33\x52\x5d\x3a\xe7\x98\xd7\x96\xd0\x13\xe0\xd4\x43\x04\x29\x97\xf0\x18\x50\xb2\x60\xea\x00\x8d\x49\x03\xd4\xc5\x4f\xca\xee\x49\xef\x63\x8f\x38\x8c\x83\x86\x5a\x0c\x8e\xc0\xe9\x81\xe0\x5b\x68\xc8\x2a\xf2\xe4\xf2\xd8\x82\x49\x76\x0e\xd0\x12\x48\x1e\xd8\x53\x07\x5e\x83\xef\xd9\xc1\xdb\x66\x99\x90\x7c\xbe\xb9\x8d\x5a\x56\xf7\xe0\x0c\xb5\xef\xea\xe2\xab\x00\x4b\x92\xd0\x8d\xec\x5a\xad\x3c\xb2\x72\x99\x5e\xfe\xc6\x1e\x7e\xb3\x94\xd0\x10\x04\x17\x73\x3a\xc0\x14\x7e\x4c\x7b\x8a\x70\x94\xee\x53\x66\x31\x41\xdd\xfc\x38\x81\xdd\x68\x8e\x0a\xce\x22\x2f\xdd\xd6\xe9\x4d\x9a\x62\x7e\x75\xfe\x72\x77\x72\x4e\x40\xb2\x4b\xfd\x23\xd3\xca\xc5\x53\xd6\xbd\xaf\xaf\x0e\x68\x2c\x2b\x2f\xa0\xbc\x70\xd5\x85\x2b\xff\x72\xcb\x7d\x5f\x73\xdd\x8e\x9d\xf7\xae\xe1\xce\x7e\xe3\xe3\x59\x90\x75\xac\x55\xdc\x6d\xda\xf1\x78\x61\xb2\x4a\x62\x43\xf2\xf4\x9e\x93\xb8\x3c\x3a\xd5\x2e\xf4\x7c\xbe\x1b\x9b\x3e\x83\x25\x23\xb1\x25\x28\xdf\x97\x50\xfe\x2a\xff\xe7\x75\xe9\x61\xd0\x2a\x47\x75\x27\xa3\x66\x59\xf6\x42\x63\xea\x87\xe9\x2d\xd5\xac\x67\x91\xca\x1c\x56\x2b\x60\xd5\xca\xd0\x1d\xfb\xcf\xa8\x61\xbd\x2e\x7a\x92\x43\xed\xfa\x59\x1a\xfe\x9f\x45\x23\x9e\x54\xf5\xb2\x27\x2b\xe7\x51\xb5\xb9\xef\xfe\x0d\xd8\x79\x34\x99\xdb\x27\x63\x36\xe8\x0e\x7a\x8d\xcb\xcc\x56\x2f\x4a\x9e\xe1\x31\x68\x3f\xd9\x46\x8c\x07\x5d\x06\x54\x78\x4f\x5d\xd5\x2c\xf7\x33\xdd\x92\x5d\x70\x4b\x87\x8b\x0c\x5a\x5f\x69\x71\x2e\xbd\xdd\x3d\xfe\x09\x00\x00\xff\xff\xd0\x29\xe6\x90\xf2\x05\x00\x00"),
		},
		"/templates/analytics-v1.yaml": &vfsgen۰CompressedFileInfo{
			name:             "analytics-v1.yaml",
			modTime:          time.Date(2019, 1, 1, 0, 1, 0, 0, time.UTC),
			uncompressedSize: 803,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x74\x92\x41\x6f\x9c\x30\x10\x85\xef\xfc\x8a\x27\xee\x4b\x41\xea\x21\xf2\xad\x6a\x2e\x51\xab\x2c\xda\x76\xd3\x63\x35\x0b\xd3\x95\x15\x63\x5b\xf6\x2c\x0a\xda\xf2\xdf\x2b\xaf\xb3\x84\x64\xd5\x39\xc1\x9b\xe1\xf3\x9b\x87\xcf\x67\xe8\x3f\xa8\x9e\xc8\x9c\x38\x56\x64\xc9\x4c\xa2\xbb\x88\x79\x2e\xc8\xeb\x27\x0e\x51\x3b\xab\xc0\x2f\xc2\x36\x3d\xc6\x4f\x63\x73\x60\xa1\xa6\x78\xd6\xb6\x57\xb8\x67\x6f\xdc\x34\xb0\x95\x62\x60\xa1\x9e\x84\x54\x01\x58\x1a\x58\x61\xc1\x6d\xc6\xe6\x55\x8c\x9e\x3a\x56\x38\x9f\x51\xed\xd8\x30\x45\xae\x1e\xaf\x72\x3a\x14\x30\x74\x60\x13\x13\x04\x20\xef\x57\x94\x8b\x34\x5e\x2d\x8d\x4d\x91\xcc\xdb\xce\x9c\x7a\x46\x79\xa0\xee\x79\xa2\xd0\xc7\x4d\xcf\x83\xab\x32\xa5\x44\x85\xbf\xd0\xb6\x67\x2b\xf8\x9c\xf8\xd1\x73\x97\xd8\x81\xbd\xd1\x1d\xc5\x6c\xe5\x75\xfd\xab\x98\x8d\x44\x09\x24\x7c\x9c\xb2\x95\xe0\x8c\xd1\xf6\xb8\xf7\x3d\x09\x67\x09\x18\xe8\xe5\xc7\x29\x1c\x59\xa1\x79\x53\xf6\x96\x46\xd2\x86\x0e\x86\x15\xea\x8b\x2e\x93\x67\x85\xdd\x1a\x51\x00\xc2\x83\x37\x0b\x6d\x1d\x5f\xaa\x75\x0e\xff\xc9\xe2\x63\x1e\xe9\xfd\xba\x60\xaa\xce\x59\x21\x6d\x39\x2c\x98\xcd\xc7\x3f\xb3\x70\xf4\x40\x69\x8f\x72\x15\xc7\x45\x4a\xa1\xb8\xa8\xc5\x85\x09\xf3\xac\x6e\xda\x42\x47\xcc\x73\xf9\x9e\xd3\x9e\x8c\x69\x9d\xd1\xdd\xf4\x2e\xdf\xfc\x85\x5f\x9a\x39\xe7\x5c\xde\x05\x59\x6d\xbb\x79\x33\xdf\xba\x20\x0a\x77\xf5\x5d\xbd\x74\xd9\x8e\xeb\xd1\xbc\xd3\xaf\xed\xee\xdb\xf7\xed\x97\xfb\xa5\x01\x8c\xe9\x5c\x85\xf6\xe1\x66\xb8\x7d\xf8\xfd\x75\xbb\x7f\xfc\x79\x3b\x5c\x36\x75\x5d\x97\xe9\x6e\xb1\xed\x93\xc3\x7f\x01\x00\x00\xff\xff\x42\xc9\x71\x8c\x23\x03\x00\x00"),
		},
		"/templates/bookings-v1.yaml": &vfsgen۰CompressedFileInfo{
			name:             "bookings-v1.yaml",
			modTime:          time.Date(2019, 1, 1, 0, 1, 0, 0, time.UTC),
			uncompressedSize: 817,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x6c\x92\x41\x6f\xdb\x30\x0c\x85\xef\xfe\x15\x84\x77\x8e\x53\x03\x3b\x14\xba\x0d\x6b\x80\x01\x1b\x96\xcc\xe9\xba\xe3\x40\xcb\x5c\x2a\x44\x96\x04\x89\x31\x66\x78\xfe\xef\x83\xa2\xd8\x75\xeb\xea\x94\x3c\xd2\x8f\xdf\x23\x38\x0c\xa0\xfe\x40\xf1\x84\xfa\x42\xa1\xa8\xad\x3d\x2b\x73\x0a\x30\x8e\x19\x3a\xf5\x44\x3e\x28\x6b\x04\xd0\x5f\x26\x13\x7f\x86\x6d\x57\xd6\xc4\x58\x66\x67\x65\x1a\x01\x0f\xe4\xb4\xed\x5b\x32\x9c\xb5\xc4\xd8\x20\xa3\xc8\x00\x0c\xb6\x24\x60\x72\xdb\x74\xe5\x4d\x0b\x0e\x25\x09\x18\x06\x28\x2a\xd2\x84\x81\x8a\xef\x93\x1c\x67\x02\x68\xac\x49\x87\xe8\x01\x80\xce\xbd\x98\x5c\x95\x6e\x02\xea\xca\x2c\x92\x1b\xa9\x2f\x0d\x41\x5e\xa3\x3c\xf7\xe8\x9b\xb0\x69\xa8\xb5\x45\x32\xc9\xa1\x80\x7f\xa0\x4c\x43\x86\xe1\x63\xb4\x0f\x8e\x64\xb4\xf6\xe4\xb4\x92\x18\x12\xc9\x2d\xfb\x24\x26\x0e\xa6\xd6\x69\x64\x4a\x24\xcb\x6c\xf1\x2d\x29\xdf\x27\x7d\x4b\x1b\xff\x4f\xe3\xe3\x93\xd6\x30\x2a\x43\x7e\x76\xd9\xbc\xd9\xda\x6c\xa3\x5a\x3c\x91\x80\x7c\xc1\x7a\x95\x22\xb1\x0d\x8a\xad\xef\x61\x1c\xc5\xaa\xcc\x78\x82\x71\xcc\x5f\xfb\x1c\x2e\x5a\x1f\xac\x56\xb2\x7f\x15\x3e\x7d\xe1\xe6\x62\x5a\x42\x7a\xce\x7a\x5e\x64\xdd\xbc\xb0\x1f\xac\x67\x01\xf7\x77\xf7\x77\x73\x95\x4c\xb7\x6c\x4d\x91\x7e\xed\xab\xaf\xdf\xf6\x9f\x1e\xe6\x02\x40\x17\xe7\x0a\xd8\xc9\x67\xbb\x6a\xdf\x7d\xfe\xb2\xff\x7d\x7c\xac\xd6\xed\xf9\x7c\xa0\x81\x7c\xa7\x24\x81\xa7\xe0\xac\x09\x94\xaf\x5c\xaa\xdd\x8f\x9f\xbb\xe3\xe3\xf1\x1d\x97\x67\x66\x27\xb6\x5b\x34\xa8\x7b\x56\x32\x88\x18\x61\xfb\xa1\x84\x5b\xc1\xe1\xf5\xa4\x67\x3d\x8f\xc7\x46\xa6\x89\x5b\xf9\x1f\x00\x00\xff\xff\x99\x82\xf9\x9e\x31\x03\x00\x00"),
		},
		"/templates/catalog-v1.yaml": &vfsgen۰CompressedFileInfo{
			name:             "catalog-v1.yaml",
			modTime:          time.Date(2019, 1, 1, 0, 1, 0, 0, time.UTC),
			uncompressedSize: 880,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x6c\x92\x4f\x8f\xda\x30\x10\xc5\xef\xf9\x14\xa3\xf4\x4c\x20\x52\x0f\x2b\xdf\xaa\x2e\x52\xa5\x56\x85\x86\x65\x7b\xac\x06\x67\xca\x5a\xeb\x7f\xb2\x87\x68\x23\xca\x77\xaf\x8c\x49\x36\x2c\xf8\x94\xbc\x71\x7e\x7e\xef\xc5\xc7\x23\xa8\xbf\x50\x3d\xa3\x3e\x50\xac\x24\x32\x6a\xb7\x87\xd3\xa9\x40\xaf\x9e\x29\x44\xe5\xac\x00\x7a\x63\xb2\xe9\x31\xce\xbb\x7a\x47\x8c\x75\xf1\xaa\x6c\x2b\xe0\x91\xbc\x76\xbd\x21\xcb\x85\x21\xc6\x16\x19\x45\x01\x60\xd1\x90\x80\x0b\x6c\xd6\xd5\x17\x29\x7a\x94\x24\xe0\x78\x84\xaa\x21\x4d\x18\xa9\xfa\x39\xc8\xe9\x48\x00\x8d\x3b\xd2\x31\x21\x00\xd0\xfb\x91\x71\x16\xba\xc1\x4e\x57\x17\xc9\xb6\x95\xfa\xd0\x12\x94\x3b\x94\xaf\x3d\x86\x36\xce\x5a\x32\xae\xca\x8c\x12\x2a\xf8\x07\xca\xb6\x64\x19\x3e\x27\x7a\xf4\x24\x13\x39\x90\xd7\x4a\x62\xcc\x46\x2e\xc1\x07\x31\xdb\x88\x1c\x90\x69\xdf\x67\x23\xc1\x69\xad\xec\x7e\xeb\x5b\x64\xca\x12\x80\xc1\xb7\xcd\x21\xec\x49\x40\xfd\xae\x6c\x2d\x76\xa8\x34\xee\x34\x09\x58\x9c\x75\xee\x3d\x09\x68\xa6\x88\x02\x80\xc9\x78\x3d\xd2\xa6\xd5\xa5\x35\x6d\xe1\x6e\x13\x1f\xdb\x48\xef\x43\xbc\xb4\xa4\xb3\x8c\xca\x52\x18\x21\xb3\xeb\x7f\x32\x52\x94\xc1\x94\xa1\x9c\x54\x71\x96\x52\x21\x2e\x2a\x76\xa1\x87\xd3\x49\xdc\x8c\x19\xd3\x25\x29\xaf\x39\xeb\x83\xd6\x6b\xa7\x95\xec\xaf\xba\xcd\x5f\xf8\x71\x98\x3b\xce\xcb\xbb\xc0\x93\xa4\xb3\x77\xeb\x6b\x17\x58\xc0\xc3\xe2\x61\x31\x4e\xc9\x76\xd3\xad\x39\xd1\xef\x55\xf3\xfd\xc7\xea\xcb\xe3\x38\x00\xe8\xd2\xb9\x02\x96\xf2\xc5\xdd\x6c\x5f\x7e\xfd\xb6\xfa\xb3\x79\x6a\x6e\xb7\x97\x52\x59\x32\x38\x54\x04\x81\xa2\x77\x36\x52\x79\xc3\x68\x96\xbf\xb6\xcb\xcd\xd3\xe6\x0e\xe3\x85\xd9\x8b\xf9\xdc\xb8\x4e\x51\x14\xc9\xfd\xfc\x53\x5d\xa6\xeb\x4a\xb6\x4d\xc1\xff\x07\x00\x00\xff\xff\x7a\xce\xb4\xe3\x70\x03\x00\x00"),
		},
		"/templates/frontpage-v1.yaml": &vfsgen۰CompressedFileInfo{
			name:             "frontpage-v1.yaml",
			modTime:          time.Date(2019, 1, 1, 0, 1, 0, 0, time.UTC),
			uncompressedSize: 900,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x74\x92\x41\x6f\xdb\x3e\x0c\xc5\xef\xfe\x14\x84\x7b\x8e\x13\xff\xf1\x3f\x14\xba\x0d\x6b\x80\x01\x1b\x96\xcc\x69\xba\xe3\x40\xcb\xac\x2b\x44\x96\x04\x89\x31\x6a\x64\xfe\xee\x83\xa2\xc4\x75\x9b\x4d\x27\xfb\x91\xfc\xe9\xf1\xd9\xa7\x13\xa8\x67\x28\x9e\x50\x1f\x29\x14\xcf\xde\x1a\x76\xd8\x12\x8c\x63\x86\x4e\x3d\x91\x0f\xca\x1a\x01\xf4\xca\x64\xe2\x63\x58\xf6\x65\x4d\x8c\x65\x76\x50\xa6\x11\xf0\x40\x4e\xdb\xa1\x23\xc3\x59\x47\x8c\x0d\x32\x8a\x0c\xc0\x60\x47\x02\x26\xdc\xa2\x2f\x2f\x62\x70\x28\x49\xc0\xe9\x04\x45\x45\x9a\x30\x50\xf1\xfd\x2a\xc7\x4b\x01\x34\xd6\xa4\x43\x84\x00\xa0\x73\x33\xca\x59\xea\xaf\x96\xfa\x32\x8b\xe6\x8d\xd4\xc7\x86\x20\xaf\x51\x1e\x06\xf4\x4d\x58\x34\xd4\xd9\x22\x51\x72\x28\xe0\x37\x28\xd3\x90\x61\xf8\x3f\xf2\x83\x23\x19\xd9\x9e\x9c\x56\x12\x43\xb2\x72\x59\xff\x2a\x26\x23\x81\x3d\x32\xb5\x43\xb2\xe2\xad\xd6\xca\xb4\x7b\xd7\x20\x53\x92\x00\x3a\x7c\xdd\x1d\x7d\x4b\x02\xca\x37\x65\x6f\xb0\x47\xa5\xb1\xd6\x24\x60\x75\xd6\x79\x70\x24\xa0\x9a\x23\x32\x00\xa6\xce\xe9\x89\x36\x8f\x2f\x9e\x79\x0e\xff\xc8\xe2\x63\x1e\xf1\xfd\xba\x60\x3c\xd2\x1a\x46\x65\xc8\x4f\x98\xc5\xc7\x2f\x33\x71\x54\x87\x71\x8f\x7c\x16\xc7\x59\x8a\xa1\xd8\xa0\xd8\xfa\x01\xc6\x51\xdc\x94\x19\x5b\x18\xc7\xfc\x3d\x67\x7b\xd4\x7a\x6b\xb5\x92\xc3\xbb\x7c\xd3\x84\x9b\x8a\x29\xe7\x74\x9c\xf5\x3c\xdb\x76\xf1\x66\x7e\x6b\x3d\x0b\xb8\x5f\xdd\xaf\xa6\x2a\x99\x7e\xde\x9a\x76\xfa\xb9\xa9\xbe\x7e\xdb\x7c\x7a\x98\x0a\x00\x7d\xbc\x57\xc0\x5a\xbe\xd8\x9b\xf6\xf5\xe7\x2f\x9b\x5f\xbb\xc7\xea\xb6\x3d\x9f\xd2\xc9\x6f\xa6\xaa\xf5\x8f\xfd\x7a\xf7\xb8\xfb\xcb\xd4\x0b\xb3\x13\xcb\xa5\x44\x46\x6d\x5b\x11\x0d\x2f\xef\xfe\x83\x8b\x5c\x5b\x7b\x50\xa6\x0d\x67\xfd\xae\xcc\xe3\xcf\x4b\xa6\x89\x11\xfc\x09\x00\x00\xff\xff\x3a\x97\xb5\x0c\x84\x03\x00\x00"),
		},
		"/templates/istio.yaml": &vfsgen۰CompressedFileInfo{
			name:             "istio.yaml",
			modTime:          time.Date(2019, 1, 1, 0, 1, 0, 0, time.UTC),
			uncompressedSize: 788,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xac\x91\x4d\x6a\xc3\x40\x0c\x85\xf7\x3e\x85\x2e\x60\x97\xc4\x5d\xcd\xae\xd0\x4d\xa0\x7f\xe4\x6f\x5b\x14\x5b\x89\x45\xc6\x33\x66\x24\x3b\x14\x93\xbb\x17\xdb\x71\x12\x5a\x08\xb4\xc4\x2b\xcf\xe8\xbd\x37\xd2\xa7\xb6\x05\xde\x42\xb2\x46\x5b\x93\x24\x2c\xca\x3e\x90\xf8\x3a\x64\x24\x70\x3c\x46\x58\xf1\x9a\x82\xb0\x77\x06\x1c\xe9\xc1\x87\x3d\xbb\xdd\x20\x4c\xd8\x3f\x34\x13\xb4\x55\x81\x69\xb4\x67\x97\x1b\x78\x26\x51\x76\xa8\xec\xdd\xbc\xb6\x14\x95\xa4\x98\xa3\xa2\x89\x00\x1c\x96\x64\xa0\xf4\x0d\x93\x9c\x8e\x52\x61\x46\x06\xda\x16\x92\x39\x59\x42\xa1\xe4\x6d\xbc\xee\x5e\x97\x8a\xb2\xce\x5a\x78\xd1\x2b\xab\x06\xdc\x6e\x39\xfb\xf0\x96\xb3\xaf\xae\x0e\xa0\x56\x86\x1f\x80\xd2\xe7\x64\x60\xb6\x58\xce\xde\x3f\x5f\x57\xcb\xd5\xd3\x4b\x04\x20\xf5\x46\x48\x7b\x4d\x7c\xea\xa4\x99\xf4\x06\x8b\x1b\xba\x98\x9b\x71\xd8\xbe\x7a\x96\x4e\x6f\x4a\xa7\xd7\xd2\xf4\xa6\x34\x8d\xe2\x38\xfe\x2b\xd5\x35\x07\xad\xd1\x2e\x28\x34\x9c\xdd\x15\xea\xa9\xc3\xf8\x92\x50\xa8\x56\x03\xa5\xe0\x6b\xa5\xb1\x9e\x5f\x16\x3b\x0e\xf5\x6b\x2d\xc3\x37\x90\x3e\xe3\x05\x38\x10\xef\x0a\x35\x90\xa6\xff\xcf\x9a\xde\x31\x2b\xfd\x99\xf5\xd8\xef\xa4\x6d\x81\x5c\xde\x01\xfa\x0e\x00\x00\xff\xff\xb4\x1e\xb3\x30\x14\x03\x00\x00"),
		},
		"/templates/movies-v1.yaml": &vfsgen۰CompressedFileInfo{
			name:             "movies-v1.yaml",
			modTime:          time.Date(2019, 1, 1, 0, 1, 0, 0, time.UTC),
			uncompressedSize: 787,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x6c\x92\x41\x8f\xd3\x30\x10\x85\xef\xf9\x15\x4f\xb9\x37\x6c\x24\x0e\x2b\xdf\x10\xbb\x12\x12\x88\x56\x5d\x76\x39\xa2\x69\x32\x14\x6b\x1d\xdb\xb2\xa7\xd6\x46\x25\xff\x1d\xb9\x6e\x43\x4a\xf1\x29\x7e\x33\xfe\x66\xde\x4c\x8e\x47\xe8\x9f\x68\x5e\xc8\x1c\x38\x36\x83\x4b\x9a\x63\x6a\x31\x4d\x15\x79\xfd\xc2\x21\x6a\x67\x15\xf8\x4d\xd8\xe6\xcf\xf8\x2e\xb5\x3b\x16\x6a\xab\x57\x6d\x7b\x85\x07\xf6\xc6\x8d\x03\x5b\xa9\x06\x16\xea\x49\x48\x55\x80\xa5\x81\x15\x0a\x6d\x95\xda\xb3\x12\x3d\x75\xac\x70\x3c\xa2\xd9\xb2\x61\x8a\xdc\x7c\xbd\xc8\xb9\x22\x60\x68\xc7\x26\x66\x02\x40\xde\x5f\x10\xa7\x7b\xba\x34\x93\xda\x2a\x77\x6d\x3b\x73\xe8\x19\xf5\x8e\xba\xd7\x91\x42\x1f\x57\x3d\x0f\xae\x29\x88\x1a\x0d\x7e\x43\xdb\x9e\xad\xe0\x7d\x86\x47\xcf\x5d\x06\x07\xf6\x46\x77\x14\x15\x72\x5b\x51\x02\x09\xef\xc7\x52\x32\x38\x63\xb4\xdd\x3f\xfb\x9e\x84\x8b\x04\x0c\xf4\xf6\x74\x08\x7b\x2e\x0f\xce\xca\xb3\xa5\x44\xda\xd0\xce\xb0\xc2\xdd\x49\x97\xd1\xb3\xc2\x76\x89\xa8\x00\xe1\xc1\x9b\x99\xb6\x9c\x51\x3e\x4b\xbf\xff\xf3\xfc\xaf\xef\x7c\xbf\x18\xc9\xa7\x73\x56\x48\x5b\x0e\x33\x63\x75\x35\xfb\x19\xa2\x07\xca\x0e\xea\x3c\xfb\xf3\xae\x4f\x52\x13\xd8\xbb\xa8\xc5\x85\x11\xd3\xa4\x6e\xc2\x42\x7b\x4c\x53\x7d\xcd\xd9\x1c\x8c\xd9\x38\xa3\xbb\xb1\x2c\xf3\xea\x85\x9f\x83\x65\xa5\xe5\x78\x17\x64\xe1\x73\xf5\xb7\xf3\x8d\x0b\xa2\x70\x7f\x77\x7f\x37\x47\xd9\xa6\x65\x6a\x31\xf4\x7d\xbd\xfd\xfc\x65\xfd\xe1\x61\x0e\x00\x29\xd7\x55\x78\xec\x7e\xb9\x9b\xf4\xc7\x8f\x9f\xd6\x3f\x9e\xbe\x6d\x6f\xd3\xeb\x32\x1a\x18\x1d\x05\x81\xa3\x77\x36\x72\x9d\x7f\x29\xb6\x7d\xee\xf9\x4f\x00\x00\x00\xff\xff\x29\xb3\x87\x6a\x13\x03\x00\x00"),
		},
		"/templates/movies-v2.yaml": &vfsgen۰CompressedFileInfo{
			name:             "movies-v2.yaml",
			modTime:          time.Date(2019, 1, 1, 0, 1, 0, 0, time.UTC),
			uncompressedSize: 790,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x6c\x92\x41\x8f\xd3\x30\x10\x85\xef\xf9\x15\x4f\xb9\x37\x6c\x2b\x0e\x2b\xdf\x10\xbb\x12\x12\x88\x56\x5d\x76\x39\xa2\x69\x32\x14\x6b\x1d\xdb\xb2\xa7\xd6\x46\x25\xff\x1d\xb9\x6e\x43\x4a\xf1\x29\x7e\x33\xfe\x66\xde\x4c\x8e\x47\xe8\x9f\x68\x5e\xc8\x1c\x38\x36\xbd\x4b\x9a\x63\x5a\x61\x1c\x2b\xf2\xfa\x85\x43\xd4\xce\x2a\xf0\x9b\xb0\xcd\x9f\xf1\x5d\x5a\xee\x58\x68\x59\xbd\x6a\xdb\x29\x3c\xb0\x37\x6e\xe8\xd9\x4a\xd5\xb3\x50\x47\x42\xaa\x02\x2c\xf5\xac\x50\x68\x8b\xb4\x3a\x2b\xd1\x53\xcb\x0a\xc7\x23\x9a\x2d\x1b\xa6\xc8\xcd\xd7\x8b\x9c\x2b\x02\x86\x76\x6c\x62\x26\x00\xe4\xfd\x05\x71\xba\xa7\x4b\x33\x69\x55\xe5\xae\x6d\x6b\x0e\x1d\xa3\xde\x51\xfb\x3a\x50\xe8\xe2\xa2\xe3\xde\x35\x05\x51\xa3\xc1\x6f\x68\xdb\xb1\x15\xbc\xcf\xf0\xe8\xb9\xcd\xe0\xc0\xde\xe8\x96\xa2\xc2\xb2\x02\xa2\x04\x12\xde\x0f\xa5\x64\x70\xc6\x68\xbb\x7f\xf6\x1d\x09\x17\x09\xe8\xe9\xed\xe9\x10\xf6\x5c\x1e\x9c\x95\x67\x4b\x89\xb4\xa1\x9d\x61\x85\xbb\x93\x2e\x83\x67\x85\xed\x1c\x51\x01\xc2\xbd\x37\x13\x6d\x3e\xa3\x7c\xe6\x7e\xff\xe7\xf9\x5f\xdf\xf9\x7e\x31\x92\x4f\xeb\xac\x90\xb6\x1c\x26\xc6\xe2\x6a\xf6\x13\x44\xf7\x94\x1d\xd4\x79\xf6\xe7\x5d\x9f\xa4\x26\xb0\x77\x51\x8b\x0b\x03\xc6\x51\xdd\x84\x85\xf6\x18\xc7\xfa\x9a\xb3\x39\x18\xb3\x71\x46\xb7\x43\x59\xe6\xd5\x0b\x3f\x05\xcb\x4a\xcb\xf1\x2e\xc8\xcc\xe7\xe2\x6f\xe7\x1b\x17\x44\xe1\xfe\xee\xfe\x6e\x8a\xb2\x4d\xf3\xd4\x62\xe8\xfb\x7a\xfb\xf9\xcb\xfa\xc3\xc3\x14\x00\x52\xae\xab\xf0\xd8\xfe\x72\x37\xe9\x8f\x1f\x3f\xad\x7f\x3c\x7d\xdb\xde\xa6\xd7\x65\x34\x30\x3a\x0a\x02\x47\xef\x6c\x64\xa4\x55\x9d\xff\x2a\xb6\x5d\x6e\xfb\x4f\x00\x00\x00\xff\xff\xdb\xc6\xc9\xb1\x16\x03\x00\x00"),
		},
		"/templates/movies-v3.yaml": &vfsgen۰CompressedFileInfo{
			name:             "movies-v3.yaml",
			modTime:          time.Date(2019, 1, 1, 0, 1, 0, 0, time.UTC),
			uncompressedSize: 790,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x6c\x92\x41\x8f\xd3\x30\x10\x85\xef\xf9\x15\x4f\xb9\x37\x6c\x55\x0e\x2b\xdf\x10\xbb\x12\x12\x88\x56\x5d\x76\x39\xa2\x69\x32\x14\x6b\x1d\xdb\xb2\xa7\xd6\x46\x25\xff\x1d\xb9\x6e\x43\x4a\xf1\x29\x7e\x33\xfe\x66\xde\x4c\x8e\x47\xe8\x9f\x68\x5e\xc8\x1c\x38\x36\xbd\x4b\x9a\x63\x5a\x61\x1c\x2b\xf2\xfa\x85\x43\xd4\xce\x2a\xf0\x9b\xb0\xcd\x9f\xf1\x5d\x5a\xee\x58\x68\x59\xbd\x6a\xdb\x29\x3c\xb0\x37\x6e\xe8\xd9\x4a\xd5\xb3\x50\x47\x42\xaa\x02\x2c\xf5\xac\x50\x68\x8b\xb4\x3a\x2b\xd1\x53\xcb\x0a\xc7\x23\x9a\x2d\x1b\xa6\xc8\xcd\xd7\x8b\x9c\x2b\x02\x86\x76\x6c\x62\x26\x00\xe4\xfd\x05\x71\xba\xa7\x4b\x33\x69\x55\xe5\xae\x6d\x6b\x0e\x1d\xa3\xde\x51\xfb\x3a\x50\xe8\xe2\xa2\xe3\xde\x35\x05\x51\xa3\xc1\x6f\x68\xdb\xb1\x15\xbc\xcf\xf0\xe8\xb9\xcd\xe0\xc0\xde\xe8\x96\xa2\xc2\xb2\x02\xa2\x04\x12\xde\x0f\xa5\x64\x70\xc6\x68\xbb\x7f\xf6\x1d\x09\x17\x09\xe8\xe9\xed\xe9\x10\xf6\x5c\x1e\x9c\x95\x67\x4b\x89\xb4\xa1\x9d\x61\x85\xbb\x93\x2e\x83\x67\x85\xed\x1c\x51\x01\xc2\xbd\x37\x13\x6d\x3e\xa3\x7c\xe6\x7e\xff\xe7\xf9\x5f\xdf\xf9\x7e\x31\x92\x4f\xeb\xac\x90\xb6\x1c\x26\xc6\xe2\x6a\xf6\x13\x44\xf7\x94\x1d\xd4\x79\xf6\xe7\x5d\x9f\xa4\x26\xb0\x77\x51\x8b\x0b\x03\xc6\x51\xdd\x84\x85\xf6\x18\xc7\xfa\x9a\xb3\x39\x18\xb3\x71\x46\xb7\x43\x59\xe6\xd5\x0b\x3f\x05\xcb\x4a\xcb\xf1\x2e\xc8\xcc\xe7\xe2\x6f\xe7\x1b\x17\x44\xe1\xfe\xee\xfe\x6e\x8a\xb2\x4d\xf3\xd4\x62\xe8\xfb\x7a\xfb\xf9\xcb\xfa\xc3\xc3\x14\x00\x52\xae\xab\xf0\xd8\xfe\x72\x37\xe9\x8f\x1f\x3f\xad\x7f\x3c\x7d\xdb\xde\xa6\xd7\x65\x34\x30\x3a\x0a\x02\x47\xef\x6c\x64\xa4\x55\x9d\xff\x2a\xb6\x5d\x6e\xfb\x4f\x00\x00\x00\xff\xff\xec\xa7\x9f\x7e\x16\x03\x00\x00"),
		},
		"/templates/namespace.yaml": &vfsgen۰CompressedFileInfo{
			name:             "namespace.yaml",
			modTime:          time.Date(2019, 1, 1, 0, 1, 0, 0, time.UTC),
			uncompressedSize: 214,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x3c\xcf\xb1\x6a\xc3\x40\x10\x04\xd0\xfe\xbe\x62\x70\xaf\x83\x40\xaa\xfb\x88\x14\x2e\xdc\xaf\x6e\x27\xb0\xf1\x69\x65\xb4\xa7\x40\xb8\xf8\xdf\x83\x62\x50\xff\x66\x98\x19\x63\x82\x7d\x22\xdf\xa4\xed\x8c\xbc\x07\x3f\x64\x61\x3c\xa4\xf2\xca\x58\xf7\xad\x12\xcf\x67\x92\x87\xdd\xb8\x85\xad\x5e\xf0\xfd\x96\xee\xe6\x5a\x70\xca\xb4\xb0\x8b\x4a\x97\x92\x00\x97\x85\x05\x63\x20\x5f\xd9\x28\xc1\x7c\xba\xa3\x09\x68\x32\xb3\x45\x49\x63\xc0\xbc\xb6\x5d\x89\xcb\x2c\xf5\xfe\x23\x9b\xc6\xa4\x5c\xd6\xfc\x22\x17\x64\xfc\xc2\x5c\xe9\x1d\xef\xaf\x30\x60\xd1\x6d\x9d\xcc\xbf\x58\xfb\xff\x1e\xba\xcc\x8d\x9a\x8e\x2b\x74\x3d\xdc\x5f\x00\x00\x00\xff\xff\xb7\xa8\xfd\xf7\xd6\x00\x00\x00"),
		},
		"/templates/notifications-v1.yaml": &vfsgen۰CompressedFileInfo{
			name:             "notifications-v1.yaml",
			modTime:          time.Date(2019, 1, 1, 0, 1, 0, 0, time.UTC),
			uncompressedSize: 838,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x7c\x92\x41\x6f\xdb\x30\x0c\x85\xef\xfe\x15\x84\xef\xf1\x62\x60\x87\x42\xb7\x61\x2d\x30\x60\xc3\x12\xa4\x6b\x77\x1c\x18\x89\xcd\x84\xca\x94\x20\x31\x46\x8d\xcc\xff\x7d\x50\x54\xbb\x4e\x33\x8c\xa7\xe4\x91\x7e\xe4\xfb\xec\xd3\x09\xec\x13\x34\x8f\xe8\x8e\x94\x1a\xf6\x62\x9f\xac\x46\xb1\x9e\x13\x8c\x63\x85\xc1\x3e\x52\x4c\xd6\xb3\x02\x7a\x11\xe2\xfc\x33\x7d\xe8\xdb\x3d\x09\xb6\xd5\xb3\x65\xa3\xe0\x96\x82\xf3\x43\x47\x2c\x55\x47\x82\x06\x05\x55\x05\xc0\xd8\x91\x82\x0b\xcb\x55\xdf\xbe\x36\x52\x40\x4d\x0a\x4e\x27\x68\x76\xe4\x08\x13\x35\xdf\x27\x39\x2f\x06\x70\xb8\x27\x97\xb2\x11\x00\x86\xf0\xce\xe9\x2c\xf7\xd3\x69\x7d\x5b\xe5\x20\xac\xdd\xd1\x10\xd4\x7b\xd4\xcf\x03\x46\x93\x56\x86\x3a\xdf\x14\xa7\x1a\x1a\xf8\x03\x96\x0d\xb1\xc0\xc7\xbc\x23\x05\xd2\xd9\x3f\x52\x70\x56\x63\x2a\xe7\xbc\xa2\x98\xc4\x72\x4c\x92\x88\x42\x87\xa1\x9c\x13\xbd\x73\x96\x0f\x0f\xc1\xa0\x50\x91\x00\x3a\x7c\xb9\x3f\xc6\x03\x29\x68\xdf\x94\x07\xc6\x1e\xad\xc3\xbd\x23\x05\xeb\xb3\x2e\x43\x20\x05\xbb\xa5\x45\x05\x20\xd4\x05\x37\xbb\x2d\x31\xe6\x5a\xb2\xf8\x0f\x8f\xf7\x4c\xf2\xff\x29\x64\x2e\xed\x59\xd0\x32\xc5\xd9\x6a\xf5\xaf\xb7\x34\x7b\xd9\x0e\x73\x9e\x7a\x81\xe5\x2c\x65\x38\x3e\x59\xf1\x71\x80\x71\x54\x57\x6d\xc1\x03\x8c\x63\x7d\xe9\xb3\x3d\x3a\xb7\xf5\xce\xea\xe1\x82\x73\x79\x22\xcc\xcd\xc2\xbb\x54\xf0\x51\x16\xa9\x57\x6f\x01\xb6\x3e\x8a\x82\x9b\xf5\xcd\x7a\xee\x12\xf7\xcb\xd1\x92\xeb\xe7\x66\xf7\xf5\xdb\xe6\xd3\xed\xdc\x00\xe8\xf3\x5e\x05\x77\xfa\xb7\xbf\x1a\xbf\xfb\xfc\x65\xf3\xeb\xfe\xc7\xee\x7a\xbc\x5e\x12\x82\x44\x2c\x75\xfe\xe4\x88\x4d\x3e\xf8\x6f\x00\x00\x00\xff\xff\x95\x5a\xa2\x73\x46\x03\x00\x00"),
		},
		"/templates/payments-v1.yaml": &vfsgen۰CompressedFileInfo{
			name:             "payments-v1.yaml",
			modTime:          time.Date(2019, 1, 1, 0, 1, 0, 0, time.UTC),
			uncompressedSize: 884,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x6c\x53\xc1\x6e\xdb\x30\x0c\xbd\xfb\x2b\x08\xef\x1c\x27\x06\x76\x28\x74\x1b\xd6\x00\x03\x36\x2c\x59\xd2\x74\xc7\x81\x91\xd8\x54\xa8\x2c\x09\x12\x63\xd4\xc8\xfc\xef\x85\xa2\xd8\x71\x9b\xf2\x64\x3f\x92\x4f\xef\x3d\x5b\xa7\x13\xe8\x27\xa8\x1e\xd1\x1c\x29\x56\x1e\xbb\x86\x2c\x47\xe8\xfb\x02\xbd\x7e\xa4\x10\xb5\xb3\x02\xe8\x95\xc9\xa6\xc7\x38\x6f\xeb\x3d\x31\xd6\xc5\x8b\xb6\x4a\xc0\x3d\x79\xe3\xce\x3b\x45\x43\x8c\x0a\x19\x45\x01\x60\xb1\x21\x01\x03\xdb\xac\xad\x2f\x58\xf4\x28\x49\xc0\xe9\x04\xd5\x86\x0c\x61\xa4\xea\xf7\x00\xa7\x33\x01\x0c\xee\xc9\xc4\xc4\x01\x80\xde\x5f\x49\xce\x48\x3b\x08\x6a\xeb\x22\x29\xb7\xd2\x1c\x15\x41\xb9\x47\xf9\xd2\x61\x50\x71\xa6\xa8\x71\x55\x26\x29\xa1\x82\xff\xa0\xad\x22\xcb\xf0\x35\xd1\x47\x4f\x32\x51\x07\xf2\x46\x4b\x8c\x59\xc9\xc5\xfb\x00\x66\x1d\x91\x03\x32\x1d\xba\xac\x24\x38\x63\xb4\x3d\xec\xbc\x42\xa6\x0c\x01\x34\xf8\xba\x3d\x86\x03\x09\xa8\xaf\xc8\xce\x62\x8b\xda\xe0\xde\x90\x80\xc5\x19\xe7\xce\x93\x80\xcd\x94\xa2\x00\x60\x6a\xbc\x19\xd9\xa6\xe1\xa5\x9a\xc6\xf0\x79\x14\x1f\xe3\x48\xef\x83\xbf\x54\xd2\x59\x46\x6d\x29\x8c\x2c\xb3\x0f\x9f\x65\xa4\xd1\x0d\x26\x17\xe5\x24\x8c\x33\x94\x22\x71\x51\xb3\x0b\x1d\xf4\xbd\xb8\x69\x33\x1e\xa0\xef\xcb\xf7\x3c\xeb\xa3\x31\x6b\x67\xb4\xec\xde\xa5\x9b\x37\xfc\xd8\xcc\x29\xe7\xf2\x2e\xf0\xc4\xeb\xec\xaa\x7d\xed\x02\x0b\xb8\x5b\xdc\x2d\xc6\x2e\xd9\x76\x3a\x9a\x2d\xfd\x5d\x6d\x7e\xfe\x5a\x7d\xbb\x1f\x1b\x00\x6d\x3a\x57\xc0\x52\x3e\xbb\x9b\xf1\xe5\xf7\x1f\xab\x7f\xdb\x87\xcd\xed\x78\x79\x09\x07\x02\x49\xd2\x2d\xa9\xf2\x66\x79\xb3\xfc\xb3\x5b\x6e\x1f\xb6\x9f\x2c\x3f\x33\x7b\x31\x9f\x5b\xc7\xfa\x49\x4b\xe4\x74\x61\x44\x52\xff\xa5\x2e\xd3\xff\x4a\x56\x25\xdf\x6f\x01\x00\x00\xff\xff\x82\x14\xf4\xaf\x74\x03\x00\x00"),
		},
		"/templates/services.yaml": &vfsgen۰CompressedFileInfo{
			name:             "services.yaml",
			modTime:          time.Date(2019, 1, 1, 0, 1, 0, 0, time.UTC),
			uncompressedSize: 1970,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xc4\xd3\xc1\x6a\xfa\x40\x10\x06\xf0\x7b\x9e\x62\xf0\x9e\xe0\x1f\xfe\x07\xc9\x43\xf4\xd0\x82\xf7\x71\x77\xb4\x83\x9b\xd9\x25\x3b\x0a\x92\xfa\xee\x65\x35\xc6\xb5\xf4\xd2\xcd\x21\x37\xf9\x4c\xbe\x6f\xf8\x41\x86\x01\x78\x0f\xcd\x16\xdd\x89\x62\x13\xa9\x3f\xb3\xa1\x08\xd7\x6b\x85\x81\xb7\xd4\x47\xf6\xd2\xc2\xf9\x5f\x75\x64\xb1\x2d\x7c\xdc\x1f\xa8\x3a\x52\xb4\xa8\xd8\x56\x00\x82\x1d\xb5\x80\x82\xee\xa2\x6c\xe2\x98\xc4\x80\x86\x5a\x18\x06\x68\xde\xc9\x11\x46\x6a\xde\x1e\x71\xaa\x07\x70\xb8\x23\x17\x53\x03\x00\x86\xf0\x5a\x01\x30\xde\x92\xc7\xe9\x58\x31\xee\x64\x09\x56\x3b\x34\xc7\x0b\xf6\x36\xd6\x96\x3a\xdf\xdc\xcb\x56\xd0\xc0\x17\xb0\x58\x12\x85\xff\x69\x26\x06\x32\x69\x22\xf8\x5e\x6f\x5b\xf5\xed\x67\x0b\x9b\xf5\x66\x7d\xdb\xb9\x9f\xff\xa9\x1a\xaa\x34\xea\xc8\xa8\xef\x7f\xbd\xaa\xae\xeb\xbf\xaa\xec\xbc\x3f\xb2\x1c\x66\xa0\x64\x0d\x99\xc9\x94\x2e\x42\x32\xad\x17\x88\x18\x54\x74\xfe\x50\x0e\xf2\x2c\xc8\x3c\x1e\xe1\x22\x1c\x8f\xf1\x02\x8d\x7d\xef\x45\x03\x1e\xa8\xdc\x23\xaf\xc8\x44\x9e\xf1\x22\x26\xcf\xf9\x02\x95\xce\x9f\x99\x66\x7c\x33\xd3\xfb\x99\xc7\x98\x2d\x82\x31\x6e\x17\x48\x88\x57\xde\xb3\x41\x65\x2f\x33\x40\x7e\xd6\x64\x2e\xaf\x7f\x2d\xc2\xf3\x7a\x42\x81\x52\xc0\x4b\x47\xa2\x33\x80\xb2\x86\xcc\x66\x4a\x17\x61\x99\xd6\x93\xc8\x30\x00\x89\x4d\xc5\xdf\x01\x00\x00\xff\xff\xfe\xd4\xa7\x37\xb2\x07\x00\x00"),
		},
		"/values.yaml": &vfsgen۰CompressedFileInfo{
			name:             "values.yaml",
			modTime:          time.Date(2019, 1, 1, 0, 1, 0, 0, time.UTC),
			uncompressedSize: 402,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x54\xce\x31\x6b\xf3\x40\x0c\xc6\xf1\xfd\x3e\x85\x48\xe6\xf8\xcd\x9b\x6e\xde\x0a\x59\x0a\x6d\xc9\xd0\xa5\xa3\x7c\x96\x5d\x11\xf9\x74\x48\xb2\xc1\xfd\xf4\x25\x34\x86\x66\xfc\xfd\x11\xe2\xd9\xc3\x99\x06\x9c\x25\x60\x41\x99\xc9\x61\x50\x83\x0e\xf3\x75\x45\xeb\xfd\xd0\xd3\xa4\x4d\xda\xc3\xc7\x17\x3b\xb0\x03\xc2\xe7\xf3\xdb\xeb\x61\x50\x9b\x30\x82\x7a\x18\x58\xe8\x76\x70\xa6\x2c\x68\x04\x0b\x1a\x63\x27\xe4\x10\x0a\x1d\x41\x45\x77\xea\x81\x4b\x28\xac\x3a\x1b\x04\x4d\x55\x30\xc8\x9b\x94\x8c\xaa\x70\x46\x6f\xe1\x94\x12\x4f\x38\x52\x9b\x00\x8c\xaa\x3a\x87\xda\xda\x42\x87\xe5\x1b\x39\x8b\xce\xfd\x3f\x14\xf1\x8a\x76\x4d\x00\x81\x63\x0b\xbb\x63\x73\x6c\x4e\xbb\x04\x50\x67\x91\x8b\x0a\xe7\xb5\x85\x97\xe1\x5d\xe3\x62\xe4\x54\x22\x25\x27\x5b\x38\x93\xb7\x10\x36\x53\x4a\xec\xc1\x6a\xe4\x3a\xdb\x9f\x8a\x05\x65\x0d\xce\x5b\xe8\x54\xaf\x5c\xc6\x8d\x19\x03\x45\xc7\xbb\x06\xd3\x12\xf5\xb6\xf5\xd7\x93\x2e\x4c\xbe\xfc\x7f\xe4\xe9\x91\x4f\x77\x16\x0d\x1e\x38\x63\xb0\x96\xed\x7d\xc5\x75\xa2\x12\x1b\x7f\x02\x00\x00\xff\xff\x66\x71\xc1\xe1\x92\x01\x00\x00"),
		},
	}
	fs["/"].(*vfsgen۰DirInfo).entries = []os.FileInfo{
		fs["/.helmignore"].(os.FileInfo),
		fs["/Chart.yaml"].(os.FileInfo),
		fs["/templates"].(os.FileInfo),
		fs["/values.yaml"].(os.FileInfo),
	}
	fs["/templates"].(*vfsgen۰DirInfo).entries = []os.FileInfo{
		fs["/templates/_helpers.tpl"].(os.FileInfo),
		fs["/templates/analytics-v1.yaml"].(os.FileInfo),
		fs["/templates/bookings-v1.yaml"].(os.FileInfo),
		fs["/templates/catalog-v1.yaml"].(os.FileInfo),
		fs["/templates/frontpage-v1.yaml"].(os.FileInfo),
		fs["/templates/istio.yaml"].(os.FileInfo),
		fs["/templates/movies-v1.yaml"].(os.FileInfo),
		fs["/templates/movies-v2.yaml"].(os.FileInfo),
		fs["/templates/movies-v3.yaml"].(os.FileInfo),
		fs["/templates/namespace.yaml"].(os.FileInfo),
		fs["/templates/notifications-v1.yaml"].(os.FileInfo),
		fs["/templates/payments-v1.yaml"].(os.FileInfo),
		fs["/templates/services.yaml"].(os.FileInfo),
	}

	return fs
}()

type vfsgen۰FS map[string]interface{}

func (fs vfsgen۰FS) Open(path string) (http.File, error) {
	path = pathpkg.Clean("/" + path)
	f, ok := fs[path]
	if !ok {
		return nil, &os.PathError{Op: "open", Path: path, Err: os.ErrNotExist}
	}

	switch f := f.(type) {
	case *vfsgen۰CompressedFileInfo:
		gr, err := gzip.NewReader(bytes.NewReader(f.compressedContent))
		if err != nil {
			// This should never happen because we generate the gzip bytes such that they are always valid.
			panic("unexpected error reading own gzip compressed bytes: " + err.Error())
		}
		return &vfsgen۰CompressedFile{
			vfsgen۰CompressedFileInfo: f,
			gr:                        gr,
		}, nil
	case *vfsgen۰DirInfo:
		return &vfsgen۰Dir{
			vfsgen۰DirInfo: f,
		}, nil
	default:
		// This should never happen because we generate only the above types.
		panic(fmt.Sprintf("unexpected type %T", f))
	}
}

// vfsgen۰CompressedFileInfo is a static definition of a gzip compressed file.
type vfsgen۰CompressedFileInfo struct {
	name              string
	modTime           time.Time
	compressedContent []byte
	uncompressedSize  int64
}

func (f *vfsgen۰CompressedFileInfo) Readdir(count int) ([]os.FileInfo, error) {
	return nil, fmt.Errorf("cannot Readdir from file %s", f.name)
}
func (f *vfsgen۰CompressedFileInfo) Stat() (os.FileInfo, error) { return f, nil }

func (f *vfsgen۰CompressedFileInfo) GzipBytes() []byte {
	return f.compressedContent
}

func (f *vfsgen۰CompressedFileInfo) Name() string       { return f.name }
func (f *vfsgen۰CompressedFileInfo) Size() int64        { return f.uncompressedSize }
func (f *vfsgen۰CompressedFileInfo) Mode() os.FileMode  { return 0444 }
func (f *vfsgen۰CompressedFileInfo) ModTime() time.Time { return f.modTime }
func (f *vfsgen۰CompressedFileInfo) IsDir() bool        { return false }
func (f *vfsgen۰CompressedFileInfo) Sys() interface{}   { return nil }

// vfsgen۰CompressedFile is an opened compressedFile instance.
type vfsgen۰CompressedFile struct {
	*vfsgen۰CompressedFileInfo
	gr      *gzip.Reader
	grPos   int64 // Actual gr uncompressed position.
	seekPos int64 // Seek uncompressed position.
}

func (f *vfsgen۰CompressedFile) Read(p []byte) (n int, err error) {
	if f.grPos > f.seekPos {
		// Rewind to beginning.
		err = f.gr.Reset(bytes.NewReader(f.compressedContent))
		if err != nil {
			return 0, err
		}
		f.grPos = 0
	}
	if f.grPos < f.seekPos {
		// Fast-forward.
		_, err = io.CopyN(ioutil.Discard, f.gr, f.seekPos-f.grPos)
		if err != nil {
			return 0, err
		}
		f.grPos = f.seekPos
	}
	n, err = f.gr.Read(p)
	f.grPos += int64(n)
	f.seekPos = f.grPos
	return n, err
}
func (f *vfsgen۰CompressedFile) Seek(offset int64, whence int) (int64, error) {
	switch whence {
	case io.SeekStart:
		f.seekPos = 0 + offset
	case io.SeekCurrent:
		f.seekPos += offset
	case io.SeekEnd:
		f.seekPos = f.uncompressedSize + offset
	default:
		panic(fmt.Errorf("invalid whence value: %v", whence))
	}
	return f.seekPos, nil
}
func (f *vfsgen۰CompressedFile) Close() error {
	return f.gr.Close()
}

// vfsgen۰DirInfo is a static definition of a directory.
type vfsgen۰DirInfo struct {
	name    string
	modTime time.Time
	entries []os.FileInfo
}

func (d *vfsgen۰DirInfo) Read([]byte) (int, error) {
	return 0, fmt.Errorf("cannot Read from directory %s", d.name)
}
func (d *vfsgen۰DirInfo) Close() error               { return nil }
func (d *vfsgen۰DirInfo) Stat() (os.FileInfo, error) { return d, nil }

func (d *vfsgen۰DirInfo) Name() string       { return d.name }
func (d *vfsgen۰DirInfo) Size() int64        { return 0 }
func (d *vfsgen۰DirInfo) Mode() os.FileMode  { return 0755 | os.ModeDir }
func (d *vfsgen۰DirInfo) ModTime() time.Time { return d.modTime }
func (d *vfsgen۰DirInfo) IsDir() bool        { return true }
func (d *vfsgen۰DirInfo) Sys() interface{}   { return nil }

// vfsgen۰Dir is an opened dir instance.
type vfsgen۰Dir struct {
	*vfsgen۰DirInfo
	pos int // Position within entries for Seek and Readdir.
}

func (d *vfsgen۰Dir) Seek(offset int64, whence int) (int64, error) {
	if offset == 0 && whence == io.SeekStart {
		d.pos = 0
		return 0, nil
	}
	return 0, fmt.Errorf("unsupported Seek in directory %s", d.name)
}

func (d *vfsgen۰Dir) Readdir(count int) ([]os.FileInfo, error) {
	if d.pos >= len(d.entries) && count > 0 {
		return nil, io.EOF
	}
	if count <= 0 || count > len(d.entries)-d.pos {
		count = len(d.entries) - d.pos
	}
	e := d.entries[d.pos : d.pos+count]
	d.pos += count
	return e, nil
}
