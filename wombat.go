/*
  Read options, call lexical scanner to get back a stream (?) of tokens
  consisting of: token name; the text; space after (none or just-space or
  a newline with an indent number). Then pass the stream of tokens to the
  compiler, getting back an ast. Then compile that...
*/

package main

import "fmt"
import "wLex"
import "wToAst"
import "wCompile"

func main() {
	flag.Parse()
	fromLex   := make(chan wLex.LexToks)
	wLex.Scan(fromLex, flag.Arg(0), "%")
	ast :=  wToAst.Build(fromLex)
	fmt.Printf("Hello, world.\n")
}
