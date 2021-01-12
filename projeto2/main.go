package main

import(
	"fmt"
	"github.com/jeffotoni/organizando.project.go/projeto2/internal/crypt"
	"github.com/jeffotoni/organizando.project.go/projeto2/internal/fmts"
	//"github.com/jeffotoni/gcolorfake"
)

func main() {
	//yellow := gcolorfake.YellowCor("Aqui ira ficar yellow")
	//fmt.Println("boa.. vamos organizar:", yellow)

	fmt.Println("Crc32:", crypt.Crc32())
	fmt.Println("Crc32:", fmts.Concat("Go Go Dale Go 2021..."))
}

