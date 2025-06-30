
# Laboratorii di Algoritmi e Strutture Dati
# Relazione del Progetto: Gestione di un dizionario di parole e schemi (Go)



## Introduzione

Il progetto "parole e schemi" gestisce un dizionario composto da **parole** (solo lettere minuscole) e **schemi** (che contengono almeno una lettera maiuscola). Il progetto permette la creazione di un **dizionario** che potrà essere popolato di parole e schemi attraverso operzioni bulk con file o dirette. Diverse operazioni sono permesse sugli elementi del dizionario. Il programma 41319A_valenti_alessandro.go contiene strutture, operazioni (e algoritmi) tutti in un unico file. viene poi sfruttato il listato fornito dalla Prof. V. Lonati per eseguire test supplementari descrtitti in questo documento.


## Descrizione di 41319A_valenti_alessandro.go
### Risposta al problema della traccia
Il problema proposto rupta intorno ad una struttura *dizionario* che mantiene la Prole e gli Schemi e una sottostruttura GrafoCatena che viene aggiornato all'inserimento e alla cancellazione delle parole. Il GrafoCatena è una lista di adiacenza che mantiene per ogni parola del dizionario la lista delle parole di distanza 1 (secondo le specifiche del problema) [Aggiorna Grafo](#aggiorna-grafo).
Il programma legge una serie di righe che contengono comandi definiti dal progetto che consentono di effettuare operazioni sul **dizionario**. 

```go
type dizionario struct {
	Parole      map[string]struct{}
	Schemi      map[string]struct{}
	GrafoCatena map[string][]string
}
```

Il dizionario contiene una mappa per le **Parole** una per gli **Schemi** e una mappa **GrafoCatena** che rappresenta la lista di connessioni a distanza 1 di ogni parola. Un percorso tra due parole nel **GrafoCatena** rappresenta una **catena**. 
La scelta di modellare Parole e Schemi con mappe perchè in Go, con *e elemento* della mappa (chiave), *trova(e)* in O(1) ammortizzato, *elimina(e)* in O(1) ammortizzato, la mappa mi garantisce che non vi siano duplicati.
In tutte le operazioni chiave (inserisci, elimina, ricerca, compatibilità) serve soprattutto test di appartenenza e aggiornamenti rapidi.
Il GrafoCatena viene aggiornato a ogni inserimento di nuove parole nel dizionario. Ricavo quindi facilmente una catena(x, y) e un gruppo(x).

### Grafo Catena 
Il grafo modellato è un grafo non orientato (non pesato) con componenti connesse multiple (alcune parti possono non essere mutualmente ragiungibili). La struttura dati scelta è una lista di adiacenza implementata con una mappa di mappe. Ogni chiave è un nodo e il valore è una (mappa di [string]struct{}) che rappresenta la lista di adiacenza di vicini. La lista di adiacenza è una map[string]struct{} per permetterci ricerca e aggiornamenti in O(1). 
----
Vale la pena soffermarsi sulla scelta implementativa. L'aggiornamento del Grafo puo essere fatto seguendo due approcci:

1) Scansione del dizionario
Per ogni parola u di lunghezza L:
	•	Confronti u con tutte le N parole del dizionario usando Damerau–Levenshtein in O(L²).
	•	Costo per inserimento (o per ogni passo di BFS): O(N × L²).
Qui la complessità cresce linearmente con la dimensione del dizionario N e quadraticamente con la lunghezza L della parola.

2) Generazione on-the-fly dei vicini (genero tutti i possibili vicini senza cercarli nel dizionario)
Per la stessa parola u di lunghezza L si generano circa:
	•	L cancellazioni
	•	L trasposizioni
	•	L×26 sostituzioni
	•	L×26 inserzioni in O(L × |Σ|) operazioni (qui |Σ| = 26).
	•	Per ciascuno dei ~2L + 52L candidati fai un lookup O(1) in map[string]struct{}.
	•	Costo per inserimento (o per passo di BFS): O(L × |Σ|).

La complessità non dipende da N (la dimensione del dizionario) se non nel più che trascurabile fattore dei lookup O(1), ma cresce solo con L e con la dimensione dell’alfabeto.

Quando conviene quale?
	•	Se N molto grande (milioni di parole) e L moderato (poche decine), l’approccio “genera vicini” è decisamente più veloce, perché O(L·|Σ|) ≪ O(N·L²).
	•	Se L molto grande (centinaia/molti caratteri) ma N piccolo (pochi elementi), si potrebbe favorire la scansione del dizionario, ma in pratica con L ≤ 50 N può essere grande, quindi “genera vicini” è quasi sempre preferibile.

In sintesi:
	•	Scansione → complessità O(N·L²) per parola
	•	Generazione → complessità O(L·|Σ|) per parola
----
Per rispondere alla richiesta del problema di un entità dizionario unica (singleton), che può essere creata se non esistente, o resettata nei contenuti se già esistente, si crea una istanza in *main()* ma viene utilizzato sempre e solo un puntatore a quella istanza in tutti i metodi.

```go
var d *dizionario
...
func main() {
	...
	dict := newDizionario()
	d = &dict
    ...
}
```
Vengono quindi esguiti insequenza i comandi inseriti in *stdin* fino a quando non viene inserito il comando **'t'**.

#### Crea
- Crea un nuovo dizionario se non esistente o ricrea le strutture del dizionario con nuove strutture vuote.
- La crezione della struttura richiede O(1). 
- Popolare da file (*c nomefile.txt'*) richiede O(n) con n=numero di parole/schemi caricate nel dizionario.

### Inserisci
- Inserisce una parola o schema nel dizionario in base alla presenza di lettere maiuscole nella stringa caricata. Questa operazione richiede di convertire in minuscole la parola che richiede O(n) con n lunghezza della parola e poi confrontarla con l'originale che richiede O(n) quindi con una complessità di O(2n) semplificato a O(n) 
- L'inserimento richiede O(1)
- Verifica duplicati in tempo costante O(1)
- Ad ogni inserimento viene aggiornato il grafo del dizionario chiamando **aggiornaGrafo(w, ADD)**
    #### Aggiorna Grafo
    - Aggiungi parola:
        - Se la parola w non esiste in GrafoCatena del dizionario la aggiunge [O(1)]
        - Se la parola esiste calcola le possibili permutazioni di distanza 1 della parola w, esegue un lookup nel dizionario e se la permutazione esiste la aggiunge alla lista di adiacenza (vicini) di w in GrafoCatena [Vedi [Grafo Catena](#grafo-catena)]
    - Rimuovi parola:
        - Elimina la parola dal dizionario in O(1)
        - Elimina la chiave (parola) in GrafoCatena dopo aver rimoss la parola dalla lista di adiacenza d tutte le parola nella sua stessa lista di adiacenza. Il costo di questa operazione è O(n) con n lunghezza della lista di adiacenza della parola da eliminare [caso peggiore la parola dista 1 da tutte le altre parole del dizionario] 


### Elimina
- Rimuove la parola dal dizionario in O(1)
- Aggiorna il GrafoCatena, itera la lista di adiacenza della parola da cancellare per eliminare la parola nelle rispettive liste di adiacenza in O(n) con n lunghezza della lista di adiacenza della parola da eliminare.


### Carica
- Legge da file parole e schemi. Richiama inserisci. Ogni operazione viene eseguita in tempo costante O(n) con n numero di parole/schemi nel file


### Compatibile
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

## Test del programma 



---

## Considerazioni finali

- La funzione `crea()` **sovrascrive** il dizionario esistente.
- L'algoritmo Damerau-Levenshtein gestisce tutti i tipi di operazioni.
- Le strutture e funzioni sono progettate per essere scalabili e compatibili con Go.
