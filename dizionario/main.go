package main

import "fmt"

func main() {
	var a dizionario
	var b dizionario
	a = newDizionario()
	b = newDizionario()

	// a e b puntano alla stessa istanza
	a.Set("chiave", "valore")
	fmt.Println(b.Get("chiave")) // stampa "valore"

	// resettiamo il contenuto del dizionario
	a.Reset()
	fmt.Println(b.Get("chiave")) // stampa ""
}
