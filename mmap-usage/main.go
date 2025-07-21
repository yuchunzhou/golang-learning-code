package main

import (
	"fmt"
	"os"
	"syscall"

	"golang.org/x/sys/unix"
)

func main() {
	file, err := os.OpenFile("test.txt", os.O_CREATE|os.O_TRUNC|os.O_RDWR, 0644)
	if err != nil {
		fmt.Printf("open file failed: %s\n", err.Error())
		return
	}

	n, err := file.WriteString("hello world")
	if err != nil {
		fmt.Printf("write to file failed: %s\n", err.Error())
		return
	}
	fmt.Printf("write %d bytes data into file\n", n)

	err = file.Sync()
	if err != nil {
		fmt.Printf("call sync failed: %s\n", err.Error())
		return
	}

	stat, err := file.Stat()
	if err != nil {
		fmt.Printf("call stat failed: %s\n", err.Error())
		return
	}

	data, err := syscall.Mmap(int(file.Fd()), 0, int(stat.Size()), syscall.PROT_READ|syscall.PROT_WRITE, syscall.MAP_SHARED)
	if err != nil {
		fmt.Printf("call mmap failed: %s\n", err.Error())
		return
	}
	fmt.Printf("read %d bytes data: %s\n", len(data), string(data))

	for i := 0; i < len(data); i++ {
		if data[i] >= 97 && data[i] <= 122 {
			data[i] -= 32
		}
	}

	err = unix.Msync(data, unix.MS_SYNC)
	if err != nil {
		fmt.Printf("call msync failed: %s\n", err.Error())
	}

	err = syscall.Munmap(data)
	if err != nil {
		fmt.Printf("call munmap failed: %s\n", err.Error())
		return
	}

	if err = file.Close(); err != nil {
		fmt.Printf("close file failed: %s\n", err.Error())
		return
	}
}
