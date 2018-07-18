package main

import (
	"fmt"
	"syscall/js"
	"unsafe"
)

const (
	width  = 320
	height = 240
	size   = 307200
)

var (
	beforeUnloadCh   = make(chan struct{})
	rawData          = js.Value{}
	data             = [size]uint8{}
	saveDataCallback = js.Global().Get("saveData")
	typedArray       = js.Global().Get("Uint8Array")
	buffer           = js.Memory.Get("buffer")
)

func beforeUnload(event js.Value) {
	beforeUnloadCh <- struct{}{}
}

func process(args []js.Value) {

	rawData = args[0]

	for i := size - 1; i >= 3; i -= 4 {
		data[i-3] = 255 - rawData.IndexUint8(i-3)
		data[i-2] = 255 - rawData.IndexUint8(i-2)
		data[i-1] = 255 - rawData.IndexUint8(i-1)
	}

	saveDataCallback.Invoke(typedArray.New(buffer, unsafe.Pointer(&data), size))
}

func main() {
	for i := 0; i < size; i += 4 {
		data[i] = 0
		data[i+1] = 0
		data[i+2] = 0
		data[i+3] = 255
	}

	callback := js.NewCallback(process)
	defer callback.Release()
	setPrintMessage := js.Global().Get("setDSPCallback")
	setPrintMessage.Invoke(callback)

	beforeUnloadCb := js.NewEventCallback(0, beforeUnload)
	defer beforeUnloadCb.Release()
	addEventListener := js.Global().Get("addEventListener")
	addEventListener.Invoke("beforeunload", beforeUnloadCb)

	<-beforeUnloadCh
	fmt.Println("Bye Wasm !")
}
