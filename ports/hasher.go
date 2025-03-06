package ports

type PasswdHasher interface {
	HashPasswd(string) (string, error)
}

type HashComparator interface {
}
