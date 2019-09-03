package main

import (
	"fmt"
	"math"
	"math/rand"
	"sync"
	"time"
)

const (
	maxID        = 100
	maxSpeed     = 20 // mph
	maxLatitude  = 90
	maxLongitude = 180
	speedLimit   = 10 // mph
	sleepTime    = 1  // seconds
)

// BirdTrack data on 1 Bird scooter
type BirdTrack struct {
	id        int
	speed     float64
	latitude  float64
	longitude float64
}

var wg = sync.WaitGroup{}

// round to nearest 2 decimal places
func round(x float64) float64 {
	return math.Round(x*100) / 100
}

// random float between min and max integers
func randFloat(min, max int) float64 {
	return float64(min) + rand.Float64()*float64(max-min)
}

func main() {
	ch := make(chan BirdTrack)
	wg.Add(1)

	go makeRandomBird(ch)
	go printBirdTooFast(ch)

	wg.Wait()
}

// Makes infinite # of random Bird scooters.
// If want certain # of loops (like 10), uncomment the 3 lines in this function.
func makeRandomBird(ch chan<- BirdTrack) {
	for {
		// for i := 0; i < 10; i++ {
		birdTrack := BirdTrack{
			id:        rand.Intn(maxID) + 1,
			speed:     round(rand.Float64() * maxSpeed),
			latitude:  round(randFloat(-maxLatitude, maxLatitude)),
			longitude: round(randFloat(-maxLongitude, maxLongitude)),
		}
		ch <- birdTrack
		time.Sleep(sleepTime * time.Second)
	}
	// close(ch)
	// wg.Done()
}

func printBirdTooFast(ch <-chan BirdTrack) {
	for bird := range ch {
		if bird.speed > speedLimit {
			fmt.Printf("%+v\n", bird)
		}
	}
}

// Output:
// {id:82 speed:18.81 latitude:29.62 longitude:-22.42}
// {id:82 speed:13.74 latitude:-78.19 longitude:-123.65}
// {id:91 speed:13.93 latitude:4.29 longitude:-169.81}
// {id:88 speed:12.15 latitude:85.54 longitude:-151.4}
// {id:14 speed:10.82 latitude:7.95 longitude:-79.74}
// {id:34 speed:10.61 latitude:-44.36 longitude:-78.45}
// {id:89 speed:13.62 latitude:-46.53 longitude:-67.85}
// {id:52 speed:14.84 latitude:54.19 longitude:82.88}
// {id:84 speed:18.44 latitude:-73.65 longitude:-2.47}
// {id:3 speed:19.1 latitude:-27.37 longitude:68.7}
// {id:78 speed:11.28 latitude:26.91 longitude:18.64}
// {id:92 speed:13.39 latitude:22.09 longitude:-46.91}
// {id:47 speed:10.71 latitude:-56.3 longitude:-94.02}
// {id:26 speed:12.5 latitude:9.03 longitude:44.5}
// {id:88 speed:16.61 latitude:-89.91 longitude:84.98}
// {id:68 speed:11.19 latitude:56.77 longitude:136.08}
// {id:27 speed:12 latitude:-85.27 longitude:124.5}
// {id:67 speed:12.84 latitude:-45.46 longitude:-117.48}
// etc...
