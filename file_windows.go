package image_picker

import (
	"syscall"
	"unsafe"

	"github.com/pkg/errors"
)

func fileDialog(title string, fileType string) (string, error) {
	var ofn openfilenameW
	buf := make([]uint16, maxPath)

	t, _ := syscall.UTF16PtrFromString(title)

	ofn.lStructSize = uint32(unsafe.Sizeof(ofn))
	ofn.lpstrTitle = t
	ofn.lpstrFile = &buf[0]
	ofn.nMaxFile = uint32(len(buf))

	var filters string
	switch fileType {
	case "image":
		filters = `*.png *.jpg *.jpeg`
	case "video":
		filters = `*.webm *.mpeg *.mkv *.mp4 *.avi *.mov *.flv`
	default:
		return "", errors.New("unsupported fileType")
	}

	if filters != "" {
		ofn.lpstrFilter = utf16PtrFromString(filters)
	}

	flags := ofnExplorer | ofnFileMustExist | ofnHideReadOnly

	ofn.flags = uint32(flags)

	if getOpenFileName(&ofn) {
		return stringFromUtf16Ptr(ofn.lpstrFile), nil
	}

	return "", errors.New("failed to open file dialog")

}
