package main

import (
	"bytes"
	_ "embed"
	"log"
	"time"

	"github.com/hajimehoshi/go-mp3"
	"github.com/hajimehoshi/oto/v2"
)

//go:embed bark.mp3
var barkMP3 []byte

func main() {
	decoder, err := mp3.NewDecoder(bytes.NewReader(barkMP3))
	if err != nil {
		log.Fatal(err)
	}

	ctx, ready, err := oto.NewContext(decoder.SampleRate(), 2, 2)
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
