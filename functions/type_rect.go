/** ***********************************************
 * File Name : type_rect.go
 * Author    : rootkiter
 * E-mail    : rootkiter@rootkiter.com
 * Created   : 2022-07-20 20:15:48 CST
************************************************* */

package main
import (
    "fmt"
)

type Rect struct {
    Length int
    Width  int
}

func (this *Rect) Area () (int) {
    return this.Length * this.Width
}

func main() {
    rect := &Rect{20, 30}
    fmt.Println(rect)
    fmt.Printf(
        "%d * %d = %d\n",
        rect.Length,
        rect.Width,
        rect.Area(),
    )
}

// go build -gcflags "-N -l" type_rect.go
