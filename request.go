package godo

type Request map[string]interface{}

func (r Request) GetHeaders() map[string]string {
	return r["__ow_headers"].(map[string]string)
}

func (r Request) GetPath() string {
	return r["__ow_path"].(string)
}

func (r Request) GetMethod() string {
	return r["__ow_method"].(string)
}

func (r Request) GetBody() string {
	return r["__ow_body"].(string)
}

func (r Request) GetQuery() string {
	return r["__ow_query"].(string)
}
