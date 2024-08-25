package constants

type key int

func init() {
}

// internal constants
const ()

// http response headers
const (
	ListResultsTotalCountHeader    = "Pagination-Count"
	ListResultsPageSizeCountHeader = "Pagination-Limit"
)

// internal keys
const (
	ContextKeyPrincipal key = iota
)
