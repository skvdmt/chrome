package network

import rnt "github.com/skvdmt/chrome/internal/devtools/types/runtime"

type AdScriptIdentifier struct {
	// The script's V8 identifier.
	scriptId rnt.ScriptId
	// V8's debugging ID for the v8::Context.
	debuggerId rnt.UniqueDebuggerId
	// The script's url (or generated name based on id if inline script).
	name string
}
