package http

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/codespade/stream-server/api"
	"github.com/go-chi/render"
)

var repo api.Repository

func Init(r api.Repository) {
	repo = r
}

func VerifyHash(w http.ResponseWriter, r *http.Request) {
	resp := api.Response{}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err)
	}

	req := api.Request{}
	err = json.Unmarshal(body, &req)
	if err != nil {
		fmt.Println(err)
	}

	hasher := md5.New()
	hasher.Write([]byte(req.Id))
	data := hex.EncodeToString(hasher.Sum(nil))

	if req.Hash == string(data) {
		fmt.Println("Hash for ID ", req.Id, "is VALID")
		resp.Id = req.Id
		resp.Status = "success"
	} else {
		fmt.Println("Hash for ID ", req.Id, "is INVALID")
		resp.Id = req.Id
		resp.Status = "failed"
	}

	render.Status(r, http.StatusOK)
	render.JSON(w, r, resp)
}

func BlockID(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err)
	}

	req := api.Request{}
	err = json.Unmarshal(body, &req)
	if err != nil {
		fmt.Println(err)
	}

	resp, err := repo.BlockID(r.Context(), req.Id)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("ID: ", resp.Id, "|Status: ", resp.Status)

	render.Status(r, http.StatusOK)
	render.JSON(w, r, resp)
}
