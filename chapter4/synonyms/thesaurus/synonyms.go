package thesaurus

type Synonyms interface {
	Synonyms(term string) ([]string, error)
}
