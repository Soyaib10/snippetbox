package forms

// Define a new errors type, which we will use to hold the validation error messages for forms.
type errors map[string][]string

// Add() method adds error messages for a given field to the map
func (e errors) Add(field, message string) {
	e[field] = append(e[field], message)
}

// Get() method to retrieve the first error message for a given field from the map.
func (e errors) Get(field string) string {
	es := e[field]
	if len(es) == 0 {
		return ""
	}
	return es[0]
}
