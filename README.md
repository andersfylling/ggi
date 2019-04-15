# ggi
GGI or Generate General Interface, is a package that helps you generate common methods by just stating a implementation.

Example (Reset interface):
```go
package main

import "github.com/andersfylling/ggi"

type C struct{
    Cool string  
}

var _ ggi.Reseter = (*C)(nil)

type B struct{
	Okay bool 
}

type A struct {
	Something int 
	OtherThing *B
	AndThenThis *C 
}

var _ ggi.Reseter = (*A)(nil)
```

Will generate the file `ggi_reseter.go` with the following contents:
```go 
func (c *C) Reset() {
    c.Cool = ""
}
func (a *A) Reset() {
    a.Something = 0
    a.OtherThing = &B{}
    a.AndThenThis.Reset() // detects Reseter implementation and resuses memory
}
```

## Interfaces
 - Reseter (TODO)
 - Hasher (TODO)
 - URLQueryStringer (TODO)