package main

import (
	"fmt"
	"log"
	"math"
	"time"

	"./rtaudio"
	//"github.com/yanel/go-rtaudio"
)

func showCompiledAPI() {

	log.Println("RtAudio version : ", rtaudio.Version())

	for _, api := range rtaudio.CompiledAPI() {
		log.Println("Compiled API: ", api)
	}
}

func showDeviceList() {
	audio, err := rtaudio.Create(rtaudio.APIUnspecified)
	//audio, err := rtaudio.Create(rtaudio.APIWindowsASIO)
	if err != nil {
		log.Fatal(err)
	}
	defer audio.Destroy()

	dev, err := audio.Devices()
	if err != nil {
		log.Fatal(err)
	}

	for id, d := range dev {
		fmt.Printf("%v ---> %#v\n", id, d)
	}

}

func play() {
	const (
		sampleRate = 48000
		bufSz      = 1024
		freq       = 440.0
	)
	phase := 0.0
	audio, err := rtaudio.Create(rtaudio.APIWindowsASIO)
	if err != nil {
		log.Fatal(err)
	}
	defer audio.Destroy()

	params := rtaudio.StreamParams{
		DeviceID:     uint(audio.DefaultOutputDevice()),
		NumChannels:  2,
		FirstChannel: 0,
	}
	options := rtaudio.StreamOptions{
		Flags: rtaudio.FlagsScheduleRealtime & rtaudio.FlagsMinimizeLatency,
	}
	cb := func(out, in rtaudio.Buffer, dur time.Duration, status rtaudio.StreamStatus) int {
		samples := out.Float32()
		for i := 0; i < len(samples)/2; i++ {
			sample := float32(math.Sin(2 * math.Pi * phase))
			phase += freq / sampleRate

			samples[i*2] = sample
			samples[i*2+1] = sample
		}
		return 0
	}
	err = audio.Open(&params, nil, rtaudio.FormatFloat32, sampleRate, bufSz, cb, &options)
	if err != nil {
		log.Fatal(err)
	}
	defer audio.Close()
	audio.Start()
	defer audio.Stop()

}

func main() {

	showCompiledAPI()
	showDeviceList()
	//play()

}
