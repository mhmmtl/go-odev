package main



//Vücut Kitle İndeksi

//Formül :   VKI = kg/(boy*boy)

import(
	"os"
	"bufio"
	"fmt"
	"strconv"
)



func main(){

	scanner:=bufio.NewScanner(os.Stdin)

	fmt.Print("Kilonuzu Giriniz:")
	scanner.Scan()

	kilo,_:=strconv.ParseFloat(scanner.Text(),64)

	fmt.Print("Boyunuzu Giriniz: ")
	scanner.Scan()

	boy,_:=strconv.ParseFloat(scanner.Text(),64)

	vki:=kilo/((boy/100)*(boy/100))

	fmt.Printf("Vücut Kitle İndeksiniz: %f",vki)












}
