package api

type Handler = func(Context)

type Context interface {
	// Executes the next method in the stack that matches the current route.
	Next()

	// Set the status code for the response.
	Status(status int) Context
	// Sets 'key: value' header to the response.
	SetHeader(key string, value string) Context

	// Writes a raw bytes to the response body using the argument in-place,
	// meaning the encoder shares the same data from the given argument. Any
	// changes to the argument after calling this function is reflected.
	WriteInplace(body []byte)
	// Writes a string to the response body and set the Content-Type header
	// value to text/plain.
	WriteString(body string)
	// Writes a raw bytes to the response body.
	WriteBytes(body []byte)
	// Encode the interface with JSON Encoder to the response body and set the
	// Content-Type header value to application/json.
	WriteJSON(body interface{})
}
