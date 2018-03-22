package shared

// StringOps :
type StringOps interface {
	Reverse(in *Request, out *Response) error
	ToUpper(in *Request, out *Response) error
	ToLower(in *Request, out *Response) error
}
