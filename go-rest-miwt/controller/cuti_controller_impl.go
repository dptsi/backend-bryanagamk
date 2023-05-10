package controller

import (
	"bryanagamk/go-rest-miwt/helper"
	"bryanagamk/go-rest-miwt/service"
	"bryanagamk/go-rest-miwt/web"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type CutiControllerImpl struct {
	CutiService service.CutiService
}

func NewCutiController(cutiService service.CutiService) CutiController {
	return &CutiControllerImpl{CutiService: cutiService}
}

func (controller *CutiControllerImpl) Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	cutiCreateRequest := web.CutiCreateRequest{}
	helper.ReadFromRequestBody(request, &cutiCreateRequest)

	cutiResponse := controller.CutiService.Create(request.Context(), cutiCreateRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   cutiResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *CutiControllerImpl) Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	cutiUpdateRequest := web.CutiUpdateRequest{}
	helper.ReadFromRequestBody(request, &cutiUpdateRequest)

	cutiUpdateRequest.IdRiwayatCuti = params.ByName("cutiId")

	cutiResponse := controller.CutiService.Update(request.Context(), cutiUpdateRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   cutiResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *CutiControllerImpl) Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	cutiId := params.ByName("cutiId")

	controller.CutiService.Delete(request.Context(), cutiId)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *CutiControllerImpl) FindById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	cutiId := params.ByName("cutiId")

	cutiResponse := controller.CutiService.FindById(request.Context(), cutiId)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   cutiResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *CutiControllerImpl) FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {

	cutiResponses := controller.CutiService.FindAll(request.Context())
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   cutiResponses,
	}

	helper.WriteToResponseBody(writer, webResponse)
}
