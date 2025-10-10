package groupino

// Clear resets the Daterelation map to an empty state.
// It is used to remove all existing relation data from the struct.

func (r *Relation) Clear() {
	r.Daterelation = make(map[string][]string)
}
