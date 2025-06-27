package main

type dizionario struct {
	Parole      map[string]struct{}
	Schemi      map[string]struct{}
	GrafoCatena map[string][]string
}

func newDizionario() *dizionario {
	return &dizionario{
		Parole: make(map[string]struct{}),
		Schemi: make(map[string]struct{}),
	}
}
