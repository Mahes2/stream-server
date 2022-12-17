package http

import (
	"encoding/json"
	"fmt"
	"github.com/codespade/stream-server/entity"
	"github.com/codespade/stream-server/service"
	"io/ioutil"
	"net/http"

	"github.com/go-chi/render"
)

var svc service.Service

func VerifyHash(w http.ResponseWriter, r *http.Request) {
	resp := entity.Response{}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err)
	}

	req := entity.Request{}
	err = json.Unmarshal(body, &req)
	if err != nil {
		fmt.Println(err)
	}

	md5Hash := svc.HashToMD5(req.Id)

	if req.Hash == md5Hash {
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

	req := entity.Request{}
	err = json.Unmarshal(body, &req)
	if err != nil {
		fmt.Println(err)
	}

	resp, err := svc.BlockID(r.Context(), req.Id)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("ID: ", resp.Id, "|Status: ", resp.Status)

	render.Status(r, http.StatusOK)
	render.JSON(w, r, resp)
}
