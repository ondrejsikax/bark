//go:build !(darwin && cgo)

package bark

import "log"

func Bark() {
	log.Println("bark: MP3 playback requires macOS with CGO_ENABLED=1")
}
