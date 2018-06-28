package helper

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Response struct {
	Status int
}

func GetMd5(s string) string {
	w := md5.New()
	io.WriteString(w, s)
	md5str := fmt.Sprintf("%x", w.Sum(nil))
	return md5str
}

func NewResponse() *Response {
	return &Response{Status: 200}
}

func (resp *Response) WriteJson(w http.ResponseWriter) {
	b, err := json.Marshal(resp)
	if err != nil {
		fmt.Println("json marshal error ", err)
		w.Write([]byte(`{Status:-1}`))
	} else {
		w.Write(b)
	}
}
