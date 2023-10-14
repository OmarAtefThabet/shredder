## Installation

```bash
go get gitlab.com/zeyad.y.g/shred
```

```go
import (
	"fmt"
	"os"
    "github.com/OmarAtefThabet/shredder"
)

func main() {

	if len(os.Args) != 2 {
        os.Exit(1)
    }
    Input := shredder.Config{Iterations: 3, Remove: false}
    Path := os.Args[1]
    err := Input.File(Path)
    if err != nil {
        fmt.Println("Error: %v", err)
    } else {
        fmt.Println("File shredded")
    }
}
```

## Testing

```bash
go test -v github.com/OmarAtefThabet/shredder
```
