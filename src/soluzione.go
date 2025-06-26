package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var d dizionario

func main() {

	d = newDizionario()

	for {
		var comando string
		_, err := fmt.Scan(&comando)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Errore nella lettura del comando:", err)
			break
		}
		esegui(d, comando)
	}
}

func esegui(d dizionario, comando string) {
	token := split(comando)

	if len(token) == 0 {
		os.Exit(-1)
	}
	// TODO
	switch comando {
	case "c":
		if len(token) == 1 {
			d = newDizionario()
		} else if len(token) == 2 {
			carica(d, token[1])
		} else if len(token) > 2 {
			catena(d, token[1], token[2])
		}
	case "p":
		stampa_parole(d)
	case "s":
		stampa_schemi(d)
	case "i":
		inserisci(d, token[1])
	case "e":
		elimina(d, token[1])
	case "r":
		ricerca(d, token[1])
	case "d":
		distanza(d, token[1], token[2])
	case "g":
		gruppi(d)
	case "t":
		os.Exit(0)
	}
}

func split(comando string) []string {
	return strings.Fields(comando)
}

func carica(d dizionario, filename string) {
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
		inserisci(d, w)
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "Errore nella lettura del file:", err)
	}
}

func inserisci(d dizionario, w string) {
	if len(w) == 0 {
		return
	}
	if w == strings.ToLower(w) {
		d.Parole[w] = struct{}{}
	} else {
		d.Schemi[w] = struct{}{}
	}
}

func stampa_parole(d dizionario) {
	fmt.Print("[")
	for w := range d.Parole {
		fmt.Println(w)
	}
	fmt.Print("]\n")
}

func stampa_schemi(d dizionario) {
	fmt.Println("[")
	for w := range d.Schemi {
		fmt.Println(w)
	}
	fmt.Print("]\n")
}

func elimina(d dizionario, w string) {
	if len(w) == 0 {
		return
	}
	if w == strings.ToLower(w) {
		delete(d.Parole, w)
	} else {
		delete(d.Schemi, w)
	}
}

func ricerca(d dizionario, S string) {
	if len(S) == 0 {
		return
	}
	if S == strings.ToLower(S) {
		return
	} else {
		for _, k := range d.Parole {
			if compatibile(d, S, k) {
				fmt.Println(k)
			}
		}
	}
}

func compatibile(d, S string, w string) bool {
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

func catena(d dizionario, w1 string, w2 string) {
	fmt.Println("non esiste")
}

func gruppo(d dizionario) {
	fmt.Println("non esiste")
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
				dp[i-1][j]+1,    // Deletion
				dp[i][j-1]+1,    // Insertion
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

func min(a, b, c int) int {
	if a < b && a < c {
		return a
	}
	if b < c {
		return b
	}
	return c
}
}