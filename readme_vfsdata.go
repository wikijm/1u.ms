// Code generated by vfsgen; DO NOT EDIT.

package main

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

// Readme statically implements the virtual filesystem provided to vfsgen.
var Readme = func() http.FileSystem {
	fs := vfsgen۰FS{
		"/": &vfsgen۰DirInfo{
			name:    "/",
			modTime: time.Date(2020, 7, 5, 19, 36, 54, 431386289, time.UTC),
		},
		"/index.html": &vfsgen۰CompressedFileInfo{
			name:             "index.html",
			modTime:          time.Date(2020, 7, 5, 19, 36, 54, 575387461, time.UTC),
			uncompressedSize: 7641,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xc4\x59\xdd\x72\xdb\x38\xb2\xbe\xf7\x53\xf4\x51\x9d\xca\xd8\x35\x26\x69\x49\xb6\x93\xd8\xb2\xe6\x28\x76\x52\x51\x9d\x24\xf6\x8c\x95\x49\xb2\x77\x10\xd9\x24\x61\x83\x00\x83\x1f\xfd\xf8\x6a\x1f\x62\x9f\x70\x9f\x64\x0b\x00\x49\x51\x3f\xce\xda\x33\xb3\xb5\x17\x89\x24\xa0\xd1\xe8\x6e\x34\xbe\xfe\xd0\x1e\xfc\xcf\xd5\xf5\xe5\xe4\xdb\xcd\x5b\xc8\x75\xc1\xe0\xe6\xf3\x9b\x0f\xe3\x4b\xe8\x04\x51\xf4\xa5\x7f\x19\x45\x57\x93\x2b\xf8\xfa\x7e\xf2\xf1\x03\x74\xc3\x23\x98\x48\xc2\x15\xd5\x54\x70\xc2\xa2\xe8\xed\xa7\x0e\x74\x72\xad\xcb\xb3\x28\x9a\xcf\xe7\xe1\xbc\x1f\x0a\x99\x45\x93\xdf\xa2\x85\xd5\xd5\xb5\x8b\xab\xaf\x81\x6e\xad\x0c\x13\x9d\x74\x86\x7b\x03\xb7\xe1\xa2\x60\x5c\x5d\xec\x50\xd3\x7d\xfd\xfa\xb5\x5f\xed\x64\x91\x24\xc3\x3d\x80\x41\x81\x9a\x80\x95\x0e\xf0\xbb\xa1\xb3\x8b\xce\xa5\xe0\x1a\xb9\x0e\x26\xcb\x12\x3b\x10\xfb\x5f\x17\x1d\x8d\x0b\x1d\xd9\xd5\xe7\x10\xe7\x44\x2a\xd4\x17\x46\xa7\xc1\xab\x0e\x44\x3f\xd6\x73\xab\x97\x0c\x77\x6a\x8b\x95\x5a\x5b\xcd\x49\x81\x17\x9d\x0c\x39\x4a\xa2\x85\x6c\x89\x97\x84\x27\x22\xae\x85\x35\xd5\x0c\x87\x5d\x13\x16\x6a\x10\xf9\x1f\x76\x98\x51\x7e\x0f\x12\xd9\x45\x47\xd9\x2d\x55\x8e\xa8\x3b\x90\x4b\x4c\x7d\x38\xd4\x59\x14\xc5\x09\xbf\x53\x61\xcc\x84\x49\x52\x46\x24\x86\xb1\x28\x22\x72\x47\x16\x11\xa3\x53\x15\x65\x54\xe7\x66\x1a\x14\x44\xde\x27\x62\xce\x83\x58\xa9\xa8\x1f\x1e\x85\xdd\xcd\x99\xb0\xa0\x3c\x74\xf6\xeb\x65\x89\xdb\xfe\xa8\x58\xd2\x52\x83\x92\xf1\x93\xf7\xce\x69\x96\x33\x9a\xe5\x3a\xbc\x53\xd1\xeb\xb0\x7b\x12\x76\x8f\x5a\x83\x76\xc3\x3b\xd5\x01\xca\x35\x66\x92\xea\xe5\x45\x47\xe5\xa4\x77\x72\x1a\x74\x1f\xcc\xcf\xfd\x37\xfc\xc3\xb7\xdf\x5f\x7f\x48\xe8\xb7\x57\x27\xe6\xeb\xc7\x07\x4a\xfb\xd3\x44\xde\x23\x5b\x96\xfd\x97\x78\xf4\xb7\xe5\x64\xf4\x6b\x7e\x74\xd1\x81\x58\x0a\xa5\x84\xa4\x19\xe5\x17\x1d\xc2\x05\x5f\x16\xc2\xa8\xce\x70\x10\x79\x9b\xff\x52\xf3\x19\xe1\x99\x21\x19\xaa\x48\xe5\xc8\xd8\x0f\x9c\xe0\xf3\xeb\x8f\xfd\xc5\xdb\xf8\xf4\xf2\x5d\x2a\x6f\x3e\x5d\x7d\xea\x9a\xf2\xeb\xcf\x27\x4b\x7e\xf7\xe5\xff\x47\x5f\x55\xf6\xf3\xf4\xcb\x69\xff\xf6\xf6\x41\xe3\xd3\x9c\xf8\x2b\xb2\x61\xa7\x4b\x5e\x5b\x94\x60\x4a\x0c\xd3\xab\x3c\xd8\x76\xe9\x21\x36\xfc\xfb\x2d\xef\x32\x96\x8d\xae\xc8\xf8\xe6\xdd\xf2\x41\xfe\xfa\xea\xf3\xed\xf8\xee\x6b\xef\xf7\xd2\x2c\xde\x3f\x7c\x9e\x7f\xa3\xea\x7a\x2e\x5e\x3d\xee\x12\x44\x43\x58\x9d\xc8\x30\x67\x77\x2a\xa4\x9c\xea\xf7\xb5\x6d\x94\x67\xd7\xfc\x83\x20\xc9\xfe\xc1\x79\xcb\xfd\xc8\x5f\xf0\xc1\x54\x24\x4b\x88\x19\x51\xea\xa2\xd3\xa4\xb5\x1d\x74\x20\xd0\xad\x6f\x51\xde\xb5\x3f\x7b\xc3\x2f\x39\xd1\x40\x15\xe8\x9c\xda\xd1\xde\x70\x6f\x50\x0e\x27\x39\x55\x76\x90\x80\x2a\x08\x63\xa0\x50\x83\x48\xe1\x01\xa5\x08\x62\xc1\x53\x9a\x19\x49\x2c\x14\xc1\xd5\xa7\x5b\x30\x9a\x32\xaa\x29\x2a\x48\x85\x04\xa2\x14\x55\xd6\x48\xa0\x1c\x12\xd4\x18\x3b\x41\xc2\x13\xc0\x45\xc9\x04\xd5\x7e\xa5\x48\xe1\xf6\xf6\xb7\x77\x81\x44\x46\x34\x26\x30\x33\xcc\xc2\xc0\xb4\xd2\x15\xc2\x58\x43\x29\xc5\x8c\x26\xa8\x00\x89\x5a\x82\x16\x60\x14\xba\x2d\x25\x4e\x29\x4f\xec\x26\x7e\xf3\xe5\x21\x10\x05\x73\x64\xcc\x7e\x12\x98\x13\x27\x9e\xa1\x06\x89\x4a\xb0\x19\x99\x32\x74\x5f\x8d\x8c\xed\x97\x58\xc8\x44\xc1\x9c\xea\x1c\x08\x5f\x42\x46\x67\xc8\x6b\xf4\x51\xe1\x20\x2a\xab\x38\x20\x68\x21\x18\x24\x02\x15\x70\xa1\x01\x8b\x92\x89\xa5\x5b\xc3\xc5\x0c\x19\x68\x8c\x73\x4e\xbf\x1b\x54\xce\x45\xea\xc5\x8c\x1b\xb2\x11\xb0\x92\x0a\xb9\xc2\x10\x46\x8c\x41\x8a\x44\x1b\x69\x85\x25\x82\x96\x74\x46\x09\xb3\x96\xd2\xa2\x64\x58\x20\xd7\x4e\x4b\x4c\x38\x4c\xd1\x7a\x4d\xd9\x12\x52\x61\xac\x66\x0e\x42\xe7\x28\x41\xd1\x82\x32\x22\x9d\x61\xeb\xa6\x2a\x94\x33\x1a\xa3\x35\x22\x36\x52\x22\xd7\x6c\x09\xd2\x70\x20\x1a\x74\x8e\x30\x88\x45\xd2\xc0\xa8\xfb\x0e\x89\x28\x08\xe5\xb0\xef\x6c\xd7\xea\x27\x50\x66\xea\xc7\xd4\x41\xa5\x3b\xef\x0d\x3f\x2b\x92\x61\x95\x1d\xf9\xf1\x70\x14\xf8\x00\x0e\xa2\xfc\xd8\x6d\x3e\x4e\x61\x29\x0c\xcc\x09\xd7\x75\xd8\x49\x15\x64\xd0\x36\xc1\xfc\x21\xa0\xb2\xb3\x84\xc3\xf8\xe6\xd0\x1d\xa5\x35\x2a\x15\x8c\x89\xb9\x3d\xca\x66\xeb\xb3\xda\x29\x6f\x70\x41\xee\x31\x78\xc1\xf4\xf9\xf8\xe6\x45\xa6\xcf\x03\x29\xc3\xb6\x0f\xb5\xf0\x3b\x21\x01\x17\xc4\x06\xf2\xb0\xf6\xab\xa5\xa0\x1b\xf6\xc2\x7e\x78\xbc\xb9\x7a\xcd\xb4\x2a\x40\x5e\xb2\x12\xa8\x8d\x91\x58\xdf\x2a\x87\x6c\x9d\xca\xb8\xff\x85\x5c\x28\x0d\x81\x86\x11\xec\xdc\x67\x6f\xe7\x28\xe4\x36\x4f\x93\x44\xa2\x52\x50\x4d\xee\xad\x1c\x92\xe8\x5c\xfa\x26\x8c\xcb\x05\x1b\xab\x84\xa8\x1c\x15\x50\xae\x34\x92\xc4\x5e\x9f\x44\x68\x65\xd3\x9d\x09\x9e\xd9\x4f\x1b\xcd\xf1\x8d\x3d\xfd\x19\x61\x34\x79\xb6\xe1\x41\x2f\xe8\x07\x5b\x86\x6f\x8c\x3e\xcf\xf0\x92\x91\x18\x41\x89\x02\xeb\x2b\x51\x4a\x4c\xe9\x22\x52\x26\x4d\xe9\x02\xa6\x98\x0a\x89\xad\x63\xaa\x4f\xc5\x02\x49\xaa\x51\x56\x53\x52\xd6\x13\xfb\xde\x6b\x89\x40\x6c\xde\x60\x72\xf0\x0c\x3f\x49\xe8\xb7\x0f\x36\x5c\x0b\x08\x4f\x42\x12\x78\xa3\x2a\xef\x9f\x21\xfb\xb4\x98\x7c\x34\x4c\xd3\x92\xad\xd0\xa7\xba\xe5\x0a\x4b\x22\x1d\x04\x4e\x97\x95\xbb\x76\x8f\xe0\x8f\xe6\x5f\x65\xa9\x55\x71\x12\x9c\x06\x2f\x83\x57\x8f\x1c\xe9\x2e\x89\x9d\xae\x3c\x6f\xd5\x49\x78\x1a\xbe\x0c\x5f\x6d\x06\x20\x3f\x1e\xae\x21\xf7\x0a\x3c\xb8\xcb\x5c\x07\xbe\x0b\x5d\xd7\x05\x98\x9a\x4c\x1d\x6e\x80\xbd\x2b\x49\x16\xe5\x83\x7b\x2e\xe6\x7c\x05\xbe\xa0\x89\xcc\xd0\x55\x9d\xc9\xf5\xe5\xe4\xfa\xb3\x63\x69\x56\xd7\x46\x59\x81\xc4\x48\x2b\x35\xbe\x81\x29\x23\xf1\x3d\xab\x6a\x95\x90\x30\xcf\xa9\xc6\xea\xb7\xab\x3e\x54\x41\x89\x32\x15\xb2\xc0\x04\x8c\xb2\x62\xa4\x06\x96\x75\x54\xb3\x38\x0f\x0c\x33\xaa\xad\xe2\x6a\x0b\x07\x6e\x54\x2a\x2b\x67\xeb\x83\x86\xfd\x38\xc7\xf8\xfe\xc0\x41\xbc\x16\x15\xfa\xc9\x29\x4d\x12\xe4\x20\x38\xb6\x57\x2a\x8c\x05\x4f\x56\x4b\x8d\xc2\x83\x15\xd8\x5b\x8c\x75\xa4\x19\x57\x26\x59\x32\xed\x2b\x9a\xad\xe5\x30\xc5\x9c\xcc\xa8\x90\x3b\x91\x76\xc9\x35\x59\x3c\x0e\xb3\x5d\x8f\xb3\x2e\xee\xd5\x50\xef\xc9\xd0\x6b\xb7\x6a\x9b\xb4\x0b\x83\xbd\xe6\xee\xe9\xeb\xb0\x77\x72\xdc\x7c\x6e\x42\xf3\x9c\xda\x8a\x59\x85\xd0\x85\xfa\x31\x90\xf6\x21\xcd\x91\xb7\x04\xbc\xd2\xb0\xfa\xfc\xb3\x68\xfe\x63\x8b\xf7\x9e\x23\xbc\xf3\x8e\xfd\x17\x76\x5e\x0f\xd0\x36\x60\x59\x4a\xc1\x44\x46\x63\x9b\x4b\xd4\x07\xb8\x26\x30\xee\x2a\xaa\x2a\xa5\x54\x15\x55\xc3\xea\xa0\x52\x1e\x4b\x47\x67\x88\x7b\x7b\x32\x3a\xa4\xa9\x5d\x2e\x11\xe6\xf6\x3f\x2e\xea\xcc\x56\xfe\x26\x50\x55\xe7\x4c\x75\x07\x18\x51\x1a\x4e\xaa\x6b\xa0\x0e\x81\xea\x9f\x54\x9d\x05\xab\xdb\xe3\x72\x63\x7c\x73\x3e\x88\x18\xf5\xfb\x38\x8e\x34\xa7\x0a\x1f\x59\x52\xdd\x2b\xc1\x31\xac\x16\x45\x86\x6d\x15\x5b\x0f\xfd\xa8\xb6\x0b\x54\xd0\xce\xb8\xaa\x8a\xb5\x2b\x55\xb0\x2a\x55\x96\xb1\xc9\x04\xa5\x23\xae\x16\xa6\x08\xa3\x0f\xb8\x75\x3f\xf6\x31\xcc\xc2\x6a\xf5\x5a\xc9\xf9\x77\x67\xba\x56\x83\xaa\x4d\x0f\x42\x70\xa4\x95\x16\x28\x4c\x55\x28\xeb\x02\xe3\xe8\x39\x92\x38\x6f\x6f\xdf\x60\x4a\xc5\xe1\x38\xfa\x58\xc5\x39\xe1\x59\x65\xac\x7f\xf9\xac\x4e\xa3\x56\xff\x67\xa1\x25\x15\xd2\x0e\xda\x87\x94\x9c\x11\xe6\x66\x9e\x80\x36\x73\x97\x47\x5e\xf3\xe6\xfa\x26\xf6\xca\xb1\x0e\x9d\xbb\x5c\xa2\xf7\x0d\xf9\x3d\x6a\xd0\x65\xbf\x7b\x54\x3b\x74\x60\xf1\xdf\x0b\x9c\x14\xcd\xfc\x09\x14\x94\x1b\x8d\xea\x60\x77\x90\x6c\x09\x78\xf1\xdd\x08\x7d\xde\x94\x0e\x4c\xfc\x80\x2d\x03\xfb\xf3\x9c\xc6\xb9\xb5\x64\x7c\xd3\x75\xc9\x60\x1a\x94\x54\x07\xb6\xf2\x4b\xd4\x46\x72\x4c\xa0\xa8\xa9\x81\x0d\x6c\x93\x74\x4d\xdd\xfb\x0f\xc4\xd9\xa5\x6c\xcf\x6d\xf7\x6c\x80\x7f\x14\xd0\x53\x21\xfb\x47\xaa\xad\xba\xdb\x7b\x19\x1e\x85\x47\x61\x77\x37\xb6\xd7\xa8\xde\x30\xf5\x0d\x54\xf7\xf7\x5b\xcf\x85\x0f\xcc\xe1\x0a\xe6\x1f\x59\x5e\x6f\xd7\x28\x10\x12\xb8\x65\x15\xfd\xe6\xac\xeb\x57\xcd\xf1\x70\x34\x1a\x6d\xbd\x64\x26\xc2\xa1\x2f\x98\xf2\xb1\x5a\x2f\x38\x5b\x36\xcf\x98\xd9\x69\x0d\xa7\xcf\x3b\x23\x5a\x06\xb3\xd3\x27\xbc\x69\x2e\x05\x13\x5c\x41\x61\x94\xf6\x09\xe3\x08\x75\xe2\xcb\x3c\x43\xbd\xc2\x9d\xb8\x5a\x19\xc2\x48\x01\x61\x73\xb2\x54\x87\x20\x09\x4f\x44\x51\xa1\x59\x1b\xb1\x2a\xea\x69\x14\x3e\xe3\x99\x30\x1a\x8d\xa0\x0d\x50\xde\x89\x6e\xdc\x8b\xe3\xfe\x26\x1c\xed\x3d\x55\xd0\x95\xa4\x76\x20\xa1\x7b\xd6\x3b\x3b\xeb\xef\xa0\x8f\x97\x9f\x46\x1f\xdf\xaa\xe6\xa8\xde\x2c\x6b\x64\x3a\x04\xc3\x4b\x22\x95\x7b\xe4\x57\x6a\xaa\x27\x76\x2c\xb8\xa2\x09\x4a\x4c\x6c\xb9\xf2\x1a\x9e\xe6\xb1\xb3\xbc\xca\xf9\x30\x16\xc5\x46\xc5\xdd\x35\xe3\x8a\x22\x07\xc2\x28\xf1\xbd\x90\x96\x50\xb8\x17\x86\xe1\x8e\x1a\x2b\xac\x60\xdc\xa2\x71\x5a\xd8\x93\x21\xde\xd6\x43\xeb\x4e\x7d\xc2\x16\xae\x9b\xf2\xe3\xe3\xfb\x5c\x3a\xe3\x75\x3c\xee\xd5\xe3\xf3\x7f\xc0\xb7\x4f\x42\xa3\xbf\x3c\x3e\xee\x2e\x6b\x85\xd1\xd0\x3c\xe1\xd6\x4f\x67\xd5\x7a\xf0\xf9\x6d\x78\xdd\x1b\xf2\xde\x28\x33\x0d\x55\x1e\xde\xe1\x66\xf3\x42\x69\x53\xd2\x64\xd5\x43\xf1\x0d\xa6\x04\x35\xa1\xec\xe0\x10\x1a\x1c\x2e\xa5\x98\x92\x29\x5b\xba\x16\xcd\xdc\x1a\x56\x77\x2d\x9e\x1b\x47\x26\x62\xc2\xec\xd8\x46\xfc\xb6\xc7\xb7\xe2\xd6\x88\x84\x8d\x3f\x5b\xb1\xcb\x8f\x87\xd7\xae\xd7\x53\xf7\x50\x96\x25\xaa\x76\xb7\xc5\xe2\x8c\xaf\x6d\x53\xd4\x73\x44\xfe\x18\x43\xd9\x26\x25\xf6\x09\xd5\x5c\x16\xaf\xdf\x52\x25\x3b\xd1\xd4\x23\xd7\xc7\xe3\xcb\xe6\x11\x55\x71\xb5\xf0\xa9\x51\x9a\x7c\x9d\xf8\x38\x4d\x19\xc9\xed\xbf\x8d\x30\x6d\x0d\x43\x82\xbe\x91\x49\x2d\xc7\xb7\x80\xed\xeb\x68\x2d\xe8\x7f\xed\x08\xd3\x7b\x5c\x00\xf2\x58\xac\xbd\x28\x6b\x26\xe7\x26\x70\xf5\xba\xe4\x5a\x59\x6f\x9a\xd6\x14\xe5\x90\xa3\xc7\x45\x7b\xcd\x48\x15\xad\x1c\x17\x1b\xd7\x6c\x8d\xe1\xb5\x63\xfc\xdc\xbc\xb1\xaa\xfb\xdd\x1e\xf6\x7b\x3d\xec\xf7\x7b\xd8\xdf\x24\xf2\x3f\x10\x78\x4a\x97\x21\x3f\x1e\x7e\x10\x19\xcc\x28\xce\xdb\x01\xa9\x98\xbc\x73\x9e\xb1\xea\x51\x5d\xd1\x6f\x7b\x2f\xcc\x94\xd1\xd8\x31\x47\x89\xbe\x33\xb9\x56\xc7\x90\x27\xa5\xa0\x5c\x3f\x85\xe7\x0f\x48\xab\xe5\x7e\x16\x45\xce\xf4\xc8\xb2\xf9\x4e\xbd\xc8\x48\xda\x19\x6e\x4d\x0f\x22\x32\x84\x7f\xfe\xfd\x1f\xae\x17\xab\x3c\xff\xef\x1e\x1d\x35\x86\xae\x28\xfe\xfa\x52\x91\xb9\x55\x84\x03\xe5\x29\xe5\x54\x5b\x57\x89\x6b\x17\x94\x24\xab\x5e\xc4\x55\x3b\xd4\x05\xa1\x6a\xbf\x1c\xb6\x69\xa1\x45\x0a\x08\xd2\xa6\x86\x8e\x6d\xb6\x24\xee\xd9\x6f\x75\xd8\xbb\x61\x38\xb7\x3a\x2b\x2c\x36\x92\xc1\xa6\x1d\x2d\xde\x4f\x40\xa3\x2c\x28\x27\xcc\xa2\x0f\xb3\x84\xdf\x2e\x5e\x7a\x0a\x58\xa2\xa4\x45\xd5\x68\x6e\x3d\x41\xf2\xde\xf0\x52\x70\x4d\x62\xad\xe0\x05\x29\xca\x73\x78\x37\xfa\xb5\xe9\xc9\x57\xd4\x33\x27\x33\x74\x77\xd3\x45\x85\x5a\x76\x20\x24\x28\x93\x65\xf5\x4f\xca\x2d\x6f\x4d\x0e\x21\x45\x64\x90\x4a\x74\xaf\xe6\xd8\x6b\x86\x02\x61\x46\x09\xac\x1d\x93\x3a\x8b\x22\x1d\x16\x18\x71\xc4\x05\x16\x94\x75\x86\xff\x57\x7f\x75\xc7\x22\x38\x4c\x90\x61\x26\x49\xb1\xa2\x4f\x63\xff\x47\x03\xdf\x20\x77\xdb\x6c\xc2\x86\x0b\xde\x2f\x4d\x16\xde\x1a\x89\xab\xe5\x6f\x4c\x85\xc1\x34\x85\xb1\x23\x4f\x54\xd7\x8f\x3f\xd7\x30\xa4\x8c\x61\x46\x18\x90\x64\x86\xdc\xbd\x39\x85\x84\xab\x2b\x71\x6b\x05\x7d\xe7\x1e\x72\x93\x21\x90\x42\x18\xee\x1a\x47\x5a\x92\x34\xa5\xf1\x2f\xed\xc4\x57\x3e\xa3\xc9\x3c\x35\x0c\x68\x82\x44\x55\x05\x61\x0c\x89\xe0\x3f\x69\x97\x06\x95\x59\xd1\x54\x24\xcb\xe1\xc0\xfd\xd5\x73\xb8\xf7\xaf\x00\x00\x00\xff\xff\x21\x67\x85\xda\xd9\x1d\x00\x00"),
		},
	}
	fs["/"].(*vfsgen۰DirInfo).entries = []os.FileInfo{
		fs["/index.html"].(os.FileInfo),
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