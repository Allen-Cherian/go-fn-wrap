package main

import (
	"fmt"
	"os"
	"unsafe"

	"github.com/bytecodealliance/wasmtime-go"
)

func write(inputPtr unsafe.Pointer, inputLen uintptr) {
	// Open the file for writing, create it if it doesn't exist
	file, err := os.OpenFile("output.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0600)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()
	//	inputBytes := ((*[1 << 30]byte)(unsafe.Pointer(inputPtr)))[:inputLen:inputLen]
	inputBytes := (*[1 << 30]byte)(inputPtr)[:inputLen:inputLen]
	inputString := string(inputBytes)
	//voteString := strconv.Itoa(int(vote))
	// Write a value to the file
	//write int32 to file
	_, err = file.WriteString(inputString)

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
	// module, err := wasmtime.NewModuleFromFile(store.Engine,  //error unknown import: `__wbindgen_placeholder__::__wbindgen_describe` has not been defined
	// 	"../target/wasm32-unknown-unknown/debug/callback.wasm",
	// )
	module, err := wasmtime.NewModuleFromFile(store.Engine,
		"../wasm/callback_bg.wasm") //error unknown import: `wbg::__wbg_write_648d82bb3db5137e` has not been defined
	if err != nil {
		fmt.Print("Error 0")
		panic(err)
	}

	// Define the vote function and link it to the VotingSystem's Vote method

	// linker := wasmtime.NewLinker(store.Engine)
	// linker.DefineFunc(store, "env", "write", write)
	write_item := wasmtime.WrapFunc(store, write)
	instance, err := wasmtime.NewInstance(store, module, []wasmtime.AsExtern{write_item})

	//instance, err := linker.Instantiate(store, module)

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
