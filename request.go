package godo

type Request map[string]interface{}

func (r Request) GetHeaders() map[string]interface{} {
	return r["__ow_headers"].(map[string]interface{})
}

func (r Request) GetPath() string {
	return r["__ow_path"].(string)
}

func (r Request) GetMethod() Method {
	return r["__ow_method"].(Method)
}

func (r Request) GetBody() string {
	return r["__ow_body"].(string)
}

func (r Request) GetQuery() string {
	return r["__ow_query"].(string)
}

type Method string

const (
	MethodHead    Method = "head"
	MethodGet     Method = "get"
	MethodPost    Method = "post"
	MethodPut     Method = "put"
	MethodDelete  Method = "delete"
	MethodConnect Method = "connect"
	MethodOptions Method = "options"
	MethodTrace   Method = "trace"
	MethodPatch   Method = "patch"
)
