package network

// AdAncestry Encapsulates the script ancestry and the root script filter list rule that caused the resource or element to be labeled as an ad.
type AdAncestry struct {
	// A chain of AdScriptIdentifiers representing the ancestry of an ad script that led to the creation of a resource or element. The chain is ordered from the script itself (lowest level) up to its root ancestor that was flagged by a filter list.
	AncestryChain []*AdScriptIdentifier `json:"ancestryChain"`
	// The filter list rule that caused the root (last) script in ancestryChain to be tagged as an ad.
	RootScriptFilterlistRule string `json:"rootScriptFilterlistRule"`
}
