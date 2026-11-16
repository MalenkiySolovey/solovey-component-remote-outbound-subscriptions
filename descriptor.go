package remotesubscriptions

import _ "embed"

//go:embed component.json
var manifest []byte

func ManifestJSON() []byte { return append([]byte(nil), manifest...) }
