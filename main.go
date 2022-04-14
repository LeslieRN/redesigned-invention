package main

import (
	"fmt"
	// "time"
	"bufio"
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"runtime"

	"github.com/LeslieRN/redesigned-invention/control"
	"github.com/LeslieRN/redesigned-invention/share"
)

// ---------------------------------------------------
// configuration contains the application settings
type configuration struct {
Server control.Server `json:"Server"`
User control.User `json:"User"`
}
func (c *configuration) ParseJSON(b []byte) error {
return json.Unmarshal(b, &c)
}
var config = &configuration{}
// var config configuration
func init() {
// Verbose logging with file name and line number
log.SetFlags(log.Lshortfile | log.LstdFlags)
// var err error
 file, err := os.OpenFile("biblos.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
if err != nil {
log.Fatalln("Failed to open log file", err)
}
 log.SetOutput(file)
runtime.GOMAXPROCS(runtime.NumCPU())
}
 func readFile( nameFile string)( tabla []string){
 file, err := os.Open(nameFile)
 if err != nil {
 log.Fatal(err)
 }
 defer file.Close()
 scanner := bufio.NewScanner(file)
 for scanner.Scan() {
 line := scanner.Text()
 fmt.Println(line)
 tabla = append(tabla, line)
 }
 return tabla
 }
func main() {
 share.Load("config.json", config)
 data, err := ioutil.ReadFile("./config.json")
 if err != nil {
 fmt.Print(err)
 }
 fmt.Println(" Datos A buscar")
 fmt.Println()
 fname := "./data.txt"
 tab := readFile(fname)
 err = json.Unmarshal([]byte(data), &config)
 if err != nil{
 log.Fatalln(err)
 }
 fmt.Println(" Reporte Libros ")
 fmt.Println(".................................")
 control.JLoginGET(config.Server, config.User)
 list := (control.JBook(config.Server, tab[0]))
 list_2 := control.JAuth(config.Server, tab[1])
 list_4 := control.JEdit(config.Server, tab[2])
 list_5 :=control.JLang(config.Server, tab[3])
 print("Libros por el título")
 fmt.Println()
	for i := 0; i < len(list); i++ {
		fmt.Println(list[i].Title)
		fmt.Println(list[i].Author)
		fmt.Println(list[i].Editor)
		fmt.Println(list[i].Language)
		fmt.Println(list[i].Comment)
		fmt.Println(list[i].Year)
		fmt.Println(".................................")
		fmt.Println()
}
print("Libros por autor")
fmt.Println()
	for i := 0; i < len(list_2); i++ {
		fmt.Println(list_2[i].Title)
		fmt.Println(list_2[i].Author)
		fmt.Println(list_2[i].Editor)
		fmt.Println(list_2[i].Language)
		fmt.Println(list_2[i].Comment)
		fmt.Println(list_2[i].Year)
		fmt.Println()
		fmt.Println(".................................")
		fmt.Println()
}
print("Libros por editorial")
for i := 0; i < len(list_4); i++ {
	fmt.Println(list_4[i].Title)
	fmt.Println(list_4[i].Author)
	fmt.Println(list_4[i].Editor)
	fmt.Println(list_4[i].Language)
	fmt.Println(list_4[i].Comment)
	fmt.Println(list_4[i].Year)
	fmt.Println(".................................")
	fmt.Println()
}

print("Libros por idioma")
for i := 0; i < len(list_5); i++ {
	fmt.Println(list_5[i].Title)
	fmt.Println(list_5[i].Author)
	fmt.Println(list_5[i].Editor)
	fmt.Println(list_5[i].Language)
	fmt.Println(list_5[i].Comment)
	fmt.Println(list_5[i].Year)
	fmt.Println(".................................")
	fmt.Println()
}
 //fmt.Println()
}
