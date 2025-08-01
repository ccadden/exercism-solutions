// Package census simulates a system used to collect census data.
package census

// Resident represents a resident in this city.
type Resident struct {
	Name    string
	Age     int
	Address map[string]string
}

// NewResident registers a new resident in this city.
func NewResident(name string, age int, address map[string]string) *Resident {
	return &Resident{Name: name, Age: age, Address: address}
}

// HasRequiredInfo determines if a given resident has all of the required information.
func (r *Resident) HasRequiredInfo() bool {
// In order to be counted, a resident must provide a non-zero value for their name and an address that contains a non-zero value for the `street` key. All other information, such as the resident's age, is optional. Implement the `HasRequiredInfo` method that returns a boolean indicating if the resident has provided the required information.
	street, hasAddress := r.Address["street"]

	if !hasAddress || street == "" || r.Name == "" {
		return false
	}

	return true
}

// Delete deletes a resident's information.
func (r *Resident) Delete() {
	r.Name = ""
	r.Age = 0
	r.Address = nil
}

// Count counts all residents that have provided the required information.
func Count(residents []*Resident) int {
	count := 0

	for _,r := range(residents) {
		if r.HasRequiredInfo() {
			count ++
		}
	}

	return count
}
