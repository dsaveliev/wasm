package main

import (
	"fmt"
	"syscall/js"
)

var (
	beforeUnloadCh = make(chan struct{})
	renderCanvas   = js.Global().Get("renderProcessed")
	data           = make([]uint8, 1228800, 1228800)
)

func beforeUnload(event js.Value) {
	beforeUnloadCh <- struct{}{}
}

func process(args []js.Value) {
	// pixelData := args[0]

	for i := 0; i < 1228800; i += 4 {
		// pixelData.SetIndex(i, pixelData.Index(i))
		args[0].SetIndex(i, args[0].Index(i+3).Int()-40)
		// data[i] = (uint8)(pixelData.Index(i).Int() + 20)
		// data[i+1] = (uint8)(pixelData.Index(i + 1).Int())
		// data[i+2] = (uint8)(pixelData.Index(i + 2).Int())
		// data[i+3] = (uint8)(pixelData.Index(i + 3).Int())
	}

	// data := pixelData.Get("data")
	// length := data.Length()
	// fmt.Printf("ARGS >>> %#v \n", length)
	// for i := 0; i < length; i++ {
	// 	// p := data.Index(i)
	// 	data.SetIndex(i, 1)
	// }

	// tData := js.TypedArrayOf(data)
	// fmt.Println("CONVERT SLICE <<<<<<")
	// pixelData.Set("data", jsData)
	// renderCanvas.Invoke(pixelData)

	//renderCanvas.Invoke(pixelData)
}

func main() {
	for i := 0; i < 1228800; i += 4 {
		data[i] = 255
		data[i+1] = 0
		data[i+2] = 0
		data[i+3] = 255
	}

	callback := js.NewCallback(process)
	defer callback.Release()
	setPrintMessage := js.Global().Get("setDSPCallback")
	setPrintMessage.Invoke(callback)

	// fillRect := js.Global().Get("fillRect")
	// fillRect.Invoke(50, 25, 150, 100)

	// drawImage := js.Global().Get("drawImage")
	// drawImage.Invoke()

	beforeUnloadCb := js.NewEventCallback(0, beforeUnload)
	defer beforeUnloadCb.Release()
	addEventListener := js.Global().Get("addEventListener")
	addEventListener.Invoke("beforeunload", beforeUnloadCb)

	<-beforeUnloadCh
	fmt.Println("Bye Wasm !")
}
