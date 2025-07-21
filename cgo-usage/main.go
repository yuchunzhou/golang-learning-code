package main

/*
   #include <stdio.h>

   void hello_world() { printf("hello world ...\n"); }
*/
import "C"

func main() {
	C.hello_world()
}
