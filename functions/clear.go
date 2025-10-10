package groupino

func ( r *Relation) Clear() {
	r.Daterelation = make(map[string][]string)
}