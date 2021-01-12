package main

import(
	"fmt"
	"internal/crypt"
	"internal/fmts"
	//"github.com/jeffotoni/gcolorfake"
)

func main() {
	//yellow := gcolorfake.YellowCor("Aqui ira ficar yellow")
	//fmt.Println("boa.. vamos organizar:", yellow)

	fmt.Println("Crc32:", crypt.Crc32())
	fmt.Println("Crc32:", fmts.Concat("Go Go Dale Go 2021..."))
}

