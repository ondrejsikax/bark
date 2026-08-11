//go:build darwin && cgo

package bark

import (
	"bytes"
	_ "embed"
	"log"
	"time"

	"github.com/ebitengine/oto/v3"
	"github.com/hajimehoshi/go-mp3"
)

//go:embed bark.mp3
var barkMP3 []byte

func Bark() {
	decoder, err := mp3.NewDecoder(bytes.NewReader(barkMP3))
	if err != nil {
		log.Fatal(err)
	}

	ctx, ready, err := oto.NewContext(&oto.NewContextOptions{
		SampleRate:   decoder.SampleRate(),
		ChannelCount: 2,
		Format:       oto.FormatSignedInt16LE,
	})
	if err != nil {
		log.Fatal(err)
	}
	<-ready

	player := ctx.NewPlayer(decoder)
	defer player.Close()

	player.Play()

	for player.IsPlaying() {
		time.Sleep(100 * time.Millisecond)
	}
}
