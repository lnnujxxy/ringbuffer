package ringbuffer

import "fmt"

func ExampleRingBuffer() {
	rb := NewRingBuffer(1024)
	rb.Write([]byte("abcd"))
	fmt.Println(rb.Length())
	fmt.Println(rb.Available())
	buf := make([]byte, 4)

	rb.Read(buf)
	fmt.Println(string(buf))
	// Output: 4
	// 1020
	// abcd
}
