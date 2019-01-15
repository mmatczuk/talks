package x

// START OMIT
type RequestContext struct {
	SessionID string `http:"sid"`
}

type SearchRequest struct {
	RequestContext
	Labels     []string `http:"l"`
	MaxResults int      `http:"max"`
	Exact      bool     `http:"x"`
}

// Other Request types...
// END OMIT
