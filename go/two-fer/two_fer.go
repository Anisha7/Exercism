// Package twofer generates a simple string with the name given
package twofer

// ShareWith returns "One for /givenName, one for me."
// /givenName is 'you' by default
func ShareWith(name string) string {
	if len(name) == 0 {
		name = "you"
	}
	return "One for " + name + ", one for me."
}
