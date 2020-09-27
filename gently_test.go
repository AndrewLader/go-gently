package gently

import (
	"fmt"
	"log"
	"os"
	"testing"
	"time"
)

type testStruct struct {
	name string
}

func (ts *testStruct) GetName() string {
	return fmt.Sprintf("testStruct '%s'", ts.name)
}

func (ts *testStruct) StopGently(signal os.Signal) {
	log.Printf("gently stopping test struct '%s'...", ts.name)
}

func TestInterruptSuccess(t *testing.T) {
	defer handleTestPanics(t)

	queueUpArtificialSignalAfterDelay(os.Interrupt, 500)
	t.Log("tested artificial interrupt signal")
}

func TestKillSuccess(t *testing.T) {
	defer handleTestPanics(t)

	queueUpArtificialSignalAfterDelay(os.Kill, 500)
	t.Log("tested artificial kill signal")
}

func createTestStruct() *testStruct {
	return &testStruct{
		name: "A Test",
	}
}

func handleTestPanics(t *testing.T) {
	recovered := recover()

	if recovered != nil {
		t.Errorf("Recovered from a panic: %v", recovered)
	} else {
		t.Log("No panics, all good...")
	}
}

func queueUpArtificialSignalAfterDelay(signal os.Signal, delay int) {
	theTestStruct := createTestStruct()

	goodnight := New()
	goodnight.Register(theTestStruct)

	go sendSignalAfterDelay(goodnight, signal, delay)

	// wait for all of the registered structs to be notified
	goodnight.Wait()
}

func sendSignalAfterDelay(goodnight *GoodNight, signal os.Signal, delay int) {
	// send specified signal after specified delay
	time.Sleep(time.Duration(delay) * time.Millisecond)
	goodnight.signalListener <- signal
}
