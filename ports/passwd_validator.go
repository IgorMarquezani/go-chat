package ports

type PasswdValidator interface {
	Validate(string) []error
}
