package option

type Release struct {
	Namespace   string
	ReleaseName string
}

type GlobalRootOptions struct {
	Namespaces    []string
	Releases      []string
	AllNamespaces bool
}
