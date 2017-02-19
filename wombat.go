/*
  Read options, call lexical scanner to get back a stream (?) of tokens
  consisting of: token name; the text; space after (none or just-space or
  a newline with an indent number). Then pass the stream of tokens to the
  compiler, getting back an ast. Then compile that...
*/

package main

import "fmt"

func main() {
	fmt.Printf("Hello, world.\n")
}
