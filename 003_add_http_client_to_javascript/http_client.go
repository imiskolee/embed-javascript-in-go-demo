package main

type HttpClient struct {

}

func (h *HttpClient) Get(url string) (string,error) {
	return "todo",nil
}
