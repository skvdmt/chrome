package network

// AdProvenance Represents the provenance of an ad resource or element. Only one of filterlistRule or adScriptAncestry can be set. If filterlistRule is provided, the resource URL directly matches a filter list rule. If adScriptAncestry is provided, an ad script initiated the resource fetch or appended the element to the DOM. If neither is provided, the entity is known to be an ad, but provenance tracking information is unavailable.
type AdProvenance struct {
	// The filterlist rule that matched, if any.
	FilterlistRule string `json:"filterlistRule"`
	// The script ancestry that created the ad, if any.
	AdScriptAncestry *AdAncestry `json:"adScriptAncestry"`
}
