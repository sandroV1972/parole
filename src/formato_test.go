package main

import (
	"testing"
)

func TestBase(t *testing.T) {
	LanciaGenericaConFileInOutAtteso(
		t,
		prog,
		"base-in",
		"base-out",
		verbose,
	)

}

func TestFormatoCaricamento(t *testing.T) {
	casiTest := []CasoTest{
		{"caricamento dizionario con file",
			`c
c base_dizionario
p`,
			`[
abba
abca
]
`},
	}
	for _, ct := range casiTest {
		t.Run(ct.nome, func(t *testing.T) {
			output := eseguiTest(ct.input)
			if output != ct.atteso {
				t.Errorf("\nInput:\n%s \n\nESECUZIONE:\n<<<<<\n%s\n>>>>\n\nATTESO:\n<<<<<\n%s\n>>>>", ct.input, output, ct.atteso)
			}
		})
	}
}

func TestFormatoInserisci(t *testing.T) {
	casiTest := []CasoTest{
		{"inserimento parole",
			`c
i a
i b
i a
p`,
			`[
a
b
]
`},
		{"inserimento schemi",
			`c
i aB
i Aa
i Aa
s`,
			`[
Aa
aB
]
`},
	}
	for _, ct := range casiTest {
		t.Run(ct.nome, func(t *testing.T) {
			output := eseguiTest(ct.input)
			if output != ct.atteso {
				t.Errorf("\nInput:\n%s \n\nESECUZIONE:\n<<<<<\n%s\n>>>>\n\nATTESO:\n<<<<<\n%s\n>>>>", ct.input, output, ct.atteso)
			}
		})
	}
}

func TestFormatoElimina(t *testing.T) {
	casiTest := []CasoTest{
		{"eliminazione parole",
			`c
i a
i b
e a
p`,
			`[
b
]
`},
		{"eliminazione schemi",
			`c
i Aa
i aB
e Aa
s`,
			`[
aB
]
`},
	}
	for _, ct := range casiTest {
		t.Run(ct.nome, func(t *testing.T) {
			output := eseguiTest(ct.input)
			if output != ct.atteso {
				t.Errorf("\nInput:\n%s \n\nESECUZIONE:\n<<<<<\n%s\n>>>>\n\nATTESO:\n<<<<<\n%s\n>>>>", ct.input, output, ct.atteso)
			}
		})
	}
}

func TestFormatoRicerca(t *testing.T) {
	casiTest := []CasoTest{
		{"ricerca schema",
			`c
i aa
i ab
r aC`,
			`aC:[
aa
ab
]
`},
	}

	for _, ct := range casiTest {
		t.Run(ct.nome, func(t *testing.T) {
			output := eseguiTest(ct.input)
			if output != ct.atteso {
				t.Errorf("\nInput:\n%s \n\nESECUZIONE:\n<<<<<\n%s\n>>>>\n\nATTESO:\n<<<<<\n%s\n>>>>", ct.input, output, ct.atteso)
			}
		})
	}
}

func TestFormatoDistanza(t *testing.T) {
	casiTest := []CasoTest{
		{"distanza 1",
			`c
d aa aba`,
			"1\n"},
		{"distanza 0",
			`c
d aa aa`,
			"0\n"},
	}

	for _, ct := range casiTest {
		t.Run(ct.nome, func(t *testing.T) {
			output := eseguiTest(ct.input)
			if output != ct.atteso {
				t.Errorf("\nInput:\n%s \n\nESECUZIONE:\n<<<<<\n%s\n>>>>\n\nATTESO:\n<<<<<\n%s\n>>>>", ct.input, output, ct.atteso)
			}
		})
	}
}

func TestFormatoCatena(t *testing.T) {
	casiTest := []CasoTest{
		{"catena esistente",
			`c
i aa
i aaa
i aba
i bba
c aa bba`,
			`(
aa
aba
bba
)
`},
		{"catena non esistente",
			`c
i aa
c aa bb`,
			"non esiste\n"},
		{"catena lunga 0",
			`c
i aa
c aa aa`,
			`(
aa
)
`},
	}

	for _, ct := range casiTest {
		t.Run(ct.nome, func(t *testing.T) {
			output := eseguiTest(ct.input)
			if output != ct.atteso {
				t.Errorf("\nInput:\n%s \n\nESECUZIONE:\n<<<<<\n%s\n>>>>\n\nATTESO:\n<<<<<\n%s\n>>>>", ct.input, output, ct.atteso)
			}
		})
	}
}

func TestFormatoGruppo(t *testing.T) {
	casiTest := []CasoTest{
		{"gruppo esistente",
			`c
i aa
i aba
i aaa
i bba
i cca
g aa`,
			`[
aa
aaa
aba
bba
]
`},
	}

	for _, ct := range casiTest {
		t.Run(ct.nome, func(t *testing.T) {
			output := eseguiTest(ct.input)
			if output != ct.atteso {
				t.Errorf("\nInput:\n%s \n\nESECUZIONE:\n<<<<<\n%s\n>>>>\n\nATTESO:\n<<<<<\n%s\n>>>>", ct.input, output, ct.atteso)
			}
		})
	}
}

func TestParolaSchemaInesistente(t *testing.T) {
	casiTest := []CasoTest{
		{"parola e schema inesistenti",
			`c
i a
i A
e b
e B`,
			""},
	}
	for _, ct := range casiTest {
		t.Run(ct.nome, func(t *testing.T) {
			output := eseguiTest(ct.input)
			if output != ct.atteso {
				t.Errorf("\nInput:\n%s \n\nESECUZIONE:\n<<<<<\n%s\n>>>>\n\nATTESO:\n<<<<<\n%s\n>>>>", ct.input, output, ct.atteso)
			}
		})
	}
}
func TestDataFormatErrato(t *testing.T) {
	casiTest := []CasoTest{
		{"parola e schema con caratteri errati",
			`c
i a$
i A&`,
			""},
	}
	for _, ct := range casiTest {
		t.Run(ct.nome, func(t *testing.T) {
			output := eseguiTest(ct.input)
			if output != ct.atteso {
				t.Errorf("\nInput:\n%s \n\nESECUZIONE:\n<<<<<\n%s\n>>>>\n\nATTESO:\n<<<<<\n%s\n>>>>", ct.input, output, ct.atteso)
			}
		})
	}
}

func TestInsertDuplicatiParole(t *testing.T) {
	casiTest := []CasoTest{
		{"parola duplicati",
			`c
i a
i a
p`,
			`[
a
]
`},
	}
	for _, ct := range casiTest {
		t.Run(ct.nome, func(t *testing.T) {
			output := eseguiTest(ct.input)
			if output != ct.atteso {
				t.Errorf("\nInput:\n%s \n\nESECUZIONE:\n<<<<<\n%s\n>>>>\n\nATTESO:\n<<<<<\n%s\n>>>>", ct.input, output, ct.atteso)
			}
		})
	}
}

func TestInsertDuplicatiSchemi(t *testing.T) {
	casiTest := []CasoTest{
		{"schema duplicati",
			`c
i A
i A
s`,
			`[
A
]
`},
	}
	for _, ct := range casiTest {
		t.Run(ct.nome, func(t *testing.T) {
			output := eseguiTest(ct.input)
			if output != ct.atteso {
				t.Errorf("\nInput:\n%s \n\nESECUZIONE:\n<<<<<\n%s\n>>>>\n\nATTESO:\n<<<<<\n%s\n>>>>", ct.input, output, ct.atteso)
			}
		})
	}
}

func TestCatenaVuota(t *testing.T) {
	casiTest := []CasoTest{
		{"catena tra due parole del dizionario non collegate",
			`c
i a
i aa
i cc
c a cc`,
			`non esiste
`},
	}
	for _, ct := range casiTest {
		t.Run(ct.nome, func(t *testing.T) {
			output := eseguiTest(ct.input)
			if output != ct.atteso {
				t.Errorf("\nInput:\n%s \n\nESECUZIONE:\n<<<<<\n%s\n>>>>\n\nATTESO:\n<<<<<\n%s\n>>>>", ct.input, output, ct.atteso)
			}
		})
	}
}

/* func TestFormatoFamiglia(t *testing.T) {
	casiTest := []CasoTest{
		{"famiglia esistente",
			`c
i ab
i bb
i aC
i Ab
i AA
i Ba
f Ab`,
			`[
aC
Ab
AA
]
`},
	}

	for _, ct := range casiTest {
		t.Run(ct.nome, func(t *testing.T) {
			output := eseguiTest(ct.input)
			if output != ct.atteso {
				t.Errorf("\nInput:\n%s \n\nESECUZIONE:\n<<<<<\n%s\n>>>>\n\nATTESO:\n<<<<<\n%s\n>>>>", ct.input, output, ct.atteso)
			}
		})
	}
}
*/
