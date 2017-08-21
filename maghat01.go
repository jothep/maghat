//go get golang.org/x/sys/windows/registry
package main
import (
	"fmt"
	"log"
	"golang.org/x/sys/windows/registry"
)

func main() {
	key, err := registry.OpenKey(registry.CURRENT_USER, `SOFTWARE\Microsoft\Windows\CurrentVersion\Internet Settings`, registry.ALL_ACCESS)
	if err != nil {
		log.Fatal(err)
	}
	key.SetDWordValue("ProxyEnable", 0x00000000)
	defer key.Close()
	
	fmt.Printf("Windows proxy is closed.")

}
