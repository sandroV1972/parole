
# Progetto: Gestione di un dizionario di parole e schemi (Go)

## Descrizione generale

Il progetto gestisce un dizionario composto da **parole** (solo lettere minuscole) e **schemi** (che contengono almeno una lettera maiuscola). Le operazioni includono inserimento, rimozione, ricerca di compatibilità, distanza di editing tra parole, e la scoperta di gruppi e famiglie.

---

## Strutture Dati

```go
type Dizionario struct {
    Parole map[string]struct{}
    Schemi map[string]struct{}
}
```

Le mappe Go `map[string]struct{}` sono usate per rappresentare insiemi (set), permettendo lookup e inserimenti in tempo **O(1)** ammortizzato.

---

## Funzioni principali e complessità

### `func crea() *Dizionario`
- Crea un nuovo dizionario vuoto.
- Alla creazione di un dizionario il dizionario esistente viene perso.
- **Tempo:** O(1)

---

### `func Inserisci(d *Dizionario, w string)`
- Inserisce una parola o schema in base alla presenza di lettere maiuscole.
- Verifica duplicati in tempo costante.
- **Tempo:** O(1)

---

### `func Elimina(d *Dizionario, w string)`
- Rimuove da entrambi i set, se presente.
- **Tempo:** O(1)

---

### `func Carica(d *Dizionario, path string)`
- Legge da file e inserisce ogni token.
- **Tempo:** O(k), con k numero di parole/schemi nel file

---

### `func Compatibile(schema, parola string) bool`
- Verifica se esiste un’assegnazione coerente di lettere per rendere schema == parola.
- **Tempo:** O(L), con L lunghezza dello schema/parola

---

### `func DamerauLevenshtein(a, b string) int`
- Calcola la distanza minima considerando inserzione, cancellazione, sostituzione, e scambio.
- **Algoritmo:** programmazione dinamica con matrice
- **Tempo:** O(n × m), dove n = len(a), m = len(b)

---

### `func Catena(d *Dizionario, x, y string) []string`
- BFS su parole collegate da distanza di editing 1.
- **Tempo:** O(N × L²) nel caso peggiore (con N parole e confronto Damerau-Levenshtein tra parole di lunghezza L)

---

### `func Gruppo(d *Dizionario, x string) []string`
- BFS per trovare la componente connessa in cui tutte le parole sono raggiungibili tramite distanza 1.
- **Tempo:** O(N × L²)

---

### `func CostruisciGrafoSchemi(schemi, parole)`
- Crea un grafo degli schemi, collegando due schemi se esiste almeno una parola compatibile con entrambi.
- **Tempo:** O(S² × P × L), dove S = numero schemi, P = numero parole, L = lunghezza media

---

### `func Famiglia(S string, schemi, parole)`
- BFS sul grafo degli schemi per trovare la famiglia che contiene `S`.
- **Tempo:** O(S + E), dove E = numero di archi compatibilità (dipende da compatibilità tra schemi)

---

## Output previsto

- Parole/schemi stampati in blocchi `[...]`
- Risultati di ricerca `schema:[...]`
- Catene stampate in `(...)`
- Errori come `non esiste` se input non valido

---

## Considerazioni finali

- La funzione `crea()` **sovrascrive** il dizionario esistente.
- L'algoritmo Damerau-Levenshtein gestisce tutti i tipi di operazioni.
- Le strutture e funzioni sono progettate per essere scalabili e compatibili con Go.
