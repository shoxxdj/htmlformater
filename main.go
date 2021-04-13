package main 

import (
 "bufio"
 "fmt"
 "os"
 "github.com/yosssi/gohtml"
)

func main(){
 scanner := bufio.NewScanner(os.Stdin)
 for scanner.Scan(){
  fmt.Println(gohtml.Format(scanner.Text()))
 }
}
