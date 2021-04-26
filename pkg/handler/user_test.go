package handler

import (
	"fmt"
	"io/ioutil"
	"net/http/httptest"
	"testing"
)

func TestUserProfile(t *testing.T) {
	//
	// handler := func(w http.ResponseWriter, r *http.Request) {
	// 	// here we write our expected response, in this case, we return a
	// 	// JSON string which is typical when dealing with REST APIs
	// 	// io.WriteString(w, "{ \"status\": \"expected service response\"}")
	// }

	req := httptest.NewRequest("GET", "/user", nil)
	w := httptest.NewRecorder()
	UserProfile(w, req)

	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)
	// fmt.Println(resp.StatusCode)
	// fmt.Println(resp.Header.Get("Content-Type"))
	fmt.Println(string(body))

	if resp.StatusCode != 200 {
		t.Errorf("not working")
	}
}
