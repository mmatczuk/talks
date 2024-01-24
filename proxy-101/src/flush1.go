func shouldFlush(res *http.Response) bool {
	if res.Request.Method == http.MethodHead {
		return false
	}
	if res.StatusCode == http.StatusNoContent || res.StatusCode == http.StatusNotModified {
		return false
	}

	return isTextEventStream(res) || res.ContentLength == -1
}

func isTextEventStream(res *http.Response) bool {
	// The MIME type is defined in https://www.w3.org/TR/eventsource/#text-event-stream
	resCT := res.Header.Get("Content-Type")
	baseCT, _, _ := mime.ParseMediaType(resCT)
	return baseCT == "text/event-stream"
}
