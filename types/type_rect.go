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

func main() {
    rect := &Rect{20, 30}
    fmt.Println(rect)
}
