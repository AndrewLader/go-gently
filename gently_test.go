package gently

import (
	"fmt"
	"log"
	"os"
	"testing"
)

type testStruct struct {
	name string
}

func (ts *testStruct) GetName() string {
	return fmt.Sprintf("testStruct %s", ts.name)
}

func (ts *testStruct) StopGently(signal os.Signal) {
	log.Printf("stopping test struct %s gently...", ts.name)
}

func TestRegisterSuccess(t *testing.T) {
	defer handleTestPanics(t)

	theTestStruct := createTestStruct()

	goodnight := New()
	goodnight.Register(theTestStruct)
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
