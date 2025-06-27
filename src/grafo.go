package main

type grafo map[string][]string

func newGrafo() grafo {
	return make(map[string][]string)
}
