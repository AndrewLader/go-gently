package gently

import (
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

// Gently is the interface a struct must implement if it wants to be registered
// to notified as to when to stop gently
type Gently interface {
	GetName() string
	StopGently(sginal os.Signal)
}

// GoodNight is the struct that will notify other structs when they should
// stop gently
type GoodNight struct {
	signalListener chan os.Signal
	toBeNotified   []Gently
	waiter         sync.WaitGroup
}

// New initializes a new instance of the GoodNight struct
func New() *GoodNight {
	var signalsToListenOn = []os.Signal{syscall.SIGTERM, syscall.SIGINT, syscall.SIGQUIT}

	goodNight := &GoodNight{
		signalListener: make(chan os.Signal),
		toBeNotified:   make([]Gently, 0),
	}

	// watch for SIGTERM, SIGINT and SIGQUIT from the operating system, and notify the app on
	// the signalListener channel
	signal.Notify(goodNight.signalListener, signalsToListenOn...)

	// add 1 to the WaitGroup
	goodNight.waiter.Add(1)

	go waitForSignal(goodNight)

	return goodNight
}

// Register is used to register a struct that implements the Gently interface
// with the GoodNight struct so it can be notified when to stop gently
func (goodNight *GoodNight) Register(toBeRegistered Gently) {
	goodNight.toBeNotified = append(goodNight.toBeNotified, toBeRegistered)
}

// Wait will wait for the GoodNight instance to signal all of its registered users
func (goodNight *GoodNight) Wait() {
	goodNight.waiter.Wait()
}

func waitForSignal(goodNight *GoodNight) {
	signalRecevied := <-goodNight.signalListener

	for _, itemToBeNotified := range goodNight.toBeNotified {
		log.Printf("Notifying { %s } to stop gently...", itemToBeNotified.GetName())
		itemToBeNotified.StopGently(signalRecevied)
	}

	signal.Stop(goodNight.signalListener)

	goodNight.waiter.Done()
}
