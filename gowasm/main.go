package main

import (
	"fmt"
	"os"

	"github.com/bytecodealliance/wasmtime-go"
)

func write(vote string) {
	// Open the file for writing, create it if it doesn't exist
	file, err := os.Create("output.txt")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	// Write a value to the file
	_, err = file.WriteString(vote)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}

	fmt.Println("Value written to file successfully.")

}
func main() {
	// Create an engine and store
	engine := wasmtime.NewEngine()
	store := wasmtime.NewStore(engine)

	// Load the WebAssembly module
	module, err := wasmtime.NewModuleFromFile(store.Engine,
		"../target/wasm32-unknown-unknown/debug/callback.wasm",
	)
	if err != nil {
		fmt.Print("Error 0")
		panic(err)
	}

	// Define the vote function and link it to the VotingSystem's Vote method

	linker := wasmtime.NewLinker(store.Engine)
	linker.DefineFunc(store, "env", "write", write)

	instance, err := linker.Instantiate(store, module)

	if err != nil {
		fmt.Print("Error 1")
		panic(err)
	}

	voteRed := instance.GetExport(store, "vote_red")
	voteBlue := instance.GetExport(store, "vote_blue")

	// Simulate voting process
	voteRed.Func().Call(store)
	voteRed.Func().Call(store)
	voteBlue.Func().Call(store)
	voteRed.Func().Call(store)

	// Get voting results

}
