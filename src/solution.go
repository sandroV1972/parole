package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var d *dizionario

var letters = "abcdefghijklmnopqrstuvwxyz"

type dizionario struct {
	Parole      map[string]struct{}
	Schemi      map[string]struct{}
	GrafoCatena map[string][]string
}

func main() {
	var comando string
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		comando = scanner.Text()
		esegui(&d, comando)
	}
}

func newDizionario() *dizionario {
	if d == nil {
		d = &dizionario{}
	}
	d.Parole = make(map[string]struct{})
	d.Schemi = make(map[string]struct{})
	d.GrafoCatena = make(map[string][]string)
	return d
}

func esegui(d **dizionario, comando string) {
	token := strings.Fields(comando)

	if len(token) == 0 {
		os.Exit(-1)
	}
	// TODO
	switch token[0] {
	case "c":
		if len(token) == 1 {
			*d = newDizionario()
			return
		} else if len(token) == 2 {
			carica(token[1])
		} else if len(token) > 2 {
			catena(token[1], token[2])
		}
	case "p":
		stampa_parole()
	case "s":
		stampa_schemi()
	case "i":
		(*d).inserisci(token[1])
	case "e":
		elimina(token[1])
	case "r":
		ricerca(token[1])
	case "d":
		distanza(token[1], token[2])
	case "g":
		gruppo(token[1])
	case "t":
		os.Exit(0)
	}
}

func carica(filename string) {
	// legge le parole in un file di input e le carica nel dizionario come parole o schemi
	file, err := os.Open(filename)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Errore nell'apertura del file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		w := scanner.Text()
		d.inserisci(w)
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "Errore nella lettura del file:", err)
	}
}

func (d dizionario) inserisci(w string) {
	if len(w) == 0 {
		return
	}
	if w == strings.ToLower(w) {
		if _, ok := d.Parole[w]; ok {
			return // La parola è già presente nel dizionario
		}
		d.Parole[w] = struct{}{}
		// inserisci la parola nel albero di ricerca e tutte le parole del dizioonario con
		// distanza 1
		aggiornaGrafo(w)

	} else {
		d.Schemi[w] = struct{}{}
	}
}

func aggiornaGrafo(w string) {
	if _, ok := d.GrafoCatena[w]; !ok {
		d.GrafoCatena[w] = []string{}
	}
	for _, k := range d.generaDistanza1(w) {
		d.GrafoCatena[w] = append(d.GrafoCatena[w], k)
		d.GrafoCatena[k] = append(d.GrafoCatena[k], w)
	}
}

// generaDistanza1 restituisce tutte le parole già presenti in d.Parole
// che sono a distanza esattamente 1 da w
func (d *dizionario) generaDistanza1(w string) []string {
	var res []string

	// substitution
	for i := 0; i < len(w); i++ {
		orig := w[i]
		for j := range letters {
			ch := letters[j]
			if byte(ch) == orig {
				continue
			}
			cand := w[:i] + string(ch) + w[i+1:]
			if _, ok := d.Parole[cand]; ok {
				res = append(res, cand)
			}
		}
	}
	// insertion
	for i := 0; i <= len(w); i++ {
		for j := range letters {
			cand := w[:i] + string(letters[j]) + w[i:]
			if _, ok := d.Parole[cand]; ok {
				res = append(res, cand)
			}
		}
	}
	// deletion
	for i := 0; i < len(w); i++ {
		cand := w[:i] + w[i+1:]
		if _, ok := d.Parole[cand]; ok {
			res = append(res, cand)
		}
	}
	// transposition
	for i := 0; i < len(w)-1; i++ {
		cand := w[:i] + string(w[i+1]) + string(w[i]) + w[i+2:]
		if _, ok := d.Parole[cand]; ok {
			res = append(res, cand)
		}
	}
	return res
}

func stampa_parole() {
	fmt.Println("[")
	for w := range d.Parole {
		fmt.Println(w)
	}
	fmt.Print("]\n")
}

func stampa_schemi() {
	fmt.Println("[")
	for w := range d.Schemi {
		fmt.Println(w)
	}
	fmt.Print("]\n")
}

func elimina(w string) {
	if len(w) == 0 {
		return
	}
	if w == strings.ToLower(w) {
		delete(d.Parole, w)
	} else {
		delete(d.Schemi, w)
	}
}

func ricerca(S string) {
	if len(S) == 0 {
		return
	}
	if S == strings.ToLower(S) {
		return
	} else {
		fmt.Printf("%s:[\n", S)
		for k := range d.Parole {
			if compatibile(S, k) {
				fmt.Println(k)
			}
		}
		fmt.Println("]")
	}
}

func compatibile(S string, w string) bool {
	M := make(map[byte]byte)
	if len(S) != len(w) {
		return false
	}
	for i := 0; i < len(S)-1; i++ {
		if strings.ToLower(string(S[i])) == string(S[i]) {
			if S[i] != w[i] {
				return false
			}
		} else {
			if M[S[i]] == 0 {
				M[S[i]] = w[i]
			} else {
				if M[S[i]] != w[i] {
					return false
				}
			}
		}
	}
	return true
}

func distanza(w1 string, w2 string) {
	fmt.Printf("%d\n", distDL(w1, w2))
}

func catena(w1 string, w2 string) {
	_, ok1 := d.Parole[w1]
	_, ok2 := d.Parole[w2]
	if !ok1 || !ok2 {
		fmt.Println("non esiste")
		return
	}
	catena := generaCatenaBFS(w1, w2)
	if len(catena) == 0 {
		fmt.Println("non esiste")
		return
	}
	if len(w1) > 0 && len(w2) > 0 {
		fmt.Println("(")
		for _, c := range catena {
			fmt.Println(c)
		}
		fmt.Println(")")
	} else {
		fmt.Println("non esiste")
	}
}

// usa l'albero di ricerca per calcolare la catena cioè la sequenza minima
// di parole a distanza 1 l'una dall'altra per arrivare da w1 a w2
func generaCatenaBFS(source string, dest string) []string {

	percorso := make([]string, 0)
	// prepara
	visited := map[string]bool{}
	parent := map[string]string{}

	queue := []string{source}
	visited[source] = true

	// BFS
	for len(queue) > 0 {
		u := queue[0]
		queue = queue[1:]

		if u == dest {
			break
		}

		for _, v := range d.GrafoCatena[u] {
			if !visited[v] {
				visited[v] = true
				parent[v] = u
				queue = append(queue, v)
			}
		}
	}

	// ricostruci il percorso
	if !visited[dest] {
		return []string{} // non esiste percorso
	}
	curr := dest
	for ; ; curr = parent[curr] {
		percorso = append([]string{curr}, percorso...)
		if curr == source {
			break
		}
	}

	return percorso
}

func gruppo(w string) {
	_, ok := d.Parole[w]
	if !ok {
		fmt.Println("non esiste")
		return
	}
	g := ricavaGruppo(w)
	if len(g) == 0 || len(w) == 0 {
		fmt.Println("non esiste")
	} else {
		fmt.Println("[")
		for i := 0; i < len(g); i++ {
			fmt.Println(g[i])
		}
		fmt.Println("]")
	}
}

func ricavaGruppo(w string) []string {
	gruppo := []string{}
	for k := range d.Parole {
		if strings.Contains(k, w) {
			gruppo = append(gruppo, k)
		}
	}
	return gruppo
}

func distDL(w1 string, w2 string) int {
	n := len(w1)
	m := len(w2)

	// Create a 2D slice to store distances
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, m+1)
	}

	// Initialize base cases
	for i := 0; i <= n; i++ {
		dp[i][0] = i
	}
	for j := 0; j <= m; j++ {
		dp[0][j] = j
	}

	// Fill the DP table
	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			cost := 0
			if w1[i-1] != w2[j-1] {
				cost = 1
			}

			dp[i][j] = min(
				dp[i-1][j]+1,      // Deletion
				dp[i][j-1]+1,      // Insertion
				dp[i-1][j-1]+cost, // Substitution
			)

			// Check for transposition
			if i > 1 && j > 1 && w1[i-1] == w2[j-2] && w1[i-2] == w2[j-1] {
				dp[i][j] = min(dp[i][j], dp[i-2][j-2]+1)
			}
		}
	}

	return dp[n][m]
}

func min(vals ...int) int {
	min := vals[0]
	for _, v := range vals[1:] {
		if v < min {
			min = v
		}
	}
	return min
}
