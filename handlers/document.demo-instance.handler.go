package handlers

import (
	"demo-shop-manager/models"
	"demo-shop-manager/responses"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func handle_demoInstance_index(w http.ResponseWriter, r *http.Request) {
	// Get the pagination parameters
	paginationOptions := models.GetPaginationParams(r, models.DemoInstanceModelMeta)

	// Load the available versions
	demoInstanceList := models.NewDemoInstanceModelList()
	err := demoInstanceList.Paginate(paginationOptions)
	if err != nil {
		responses.SendError(r, w, http.StatusInternalServerError, err.Error())
		return
	}

	// Load the versions for the models
	err = demoInstanceList.LoadVersions()
	if err != nil {
		responses.SendError(r, w, http.StatusInternalServerError, err.Error())
		return
	}

	// Load the total count
	totalCount, err := demoInstanceList.GetTotalCount()
	if err != nil {
		responses.SendError(r, w, http.StatusInternalServerError, err.Error())
		return
	}

	// Prepare the response
	response := responses.NewDocumentIndexResponse(
		totalCount,
		paginationOptions.Page,
		paginationOptions.Limit,
		models.DemoInstanceModelMeta.DocumentName,
		demoInstanceList,
	)

	// Send the response
	responses.SendResponse(r, w, http.StatusOK, response)
}

func handle_demoInstance_read(w http.ResponseWriter, r *http.Request) {
	// Get the ID from the URL
	params := mux.Vars(r)
	id := params["id"]

	// Load the demo instance
	demoInstance := &models.DemoInstanceModel{Id: id}
	err := demoInstance.Load()
	if err != nil {
		responses.SendError(r, w, http.StatusInternalServerError, err.Error())
		return
	}

	// Load the versions for the model
	_, err = demoInstance.LoadVersion()
	if err != nil {
		responses.SendError(r, w, http.StatusInternalServerError, err.Error())
		return
	}

	// Prepare the response
	response := responses.NewDocumentReadResponse(
		models.DemoInstanceModelMeta.DocumentName,
		demoInstance,
	)

	// Send the response
	responses.SendResponse(r, w, http.StatusOK, response)
}

func handle_demoInstance_delete(w http.ResponseWriter, r *http.Request) {
	// Get the ID from the URL
	params := mux.Vars(r)
	id := params["id"]

	// Load the demo instance
	demoInstance := &models.DemoInstanceModel{Id: id}
	err := demoInstance.Load()
	if err != nil {
		responses.SendError(r, w, http.StatusInternalServerError, err.Error())
		return
	}

	// Delete the demo instance
	err = demoInstance.Delete()
	if err != nil {
		responses.SendError(r, w, http.StatusInternalServerError, err.Error())
		return
	}

	// Prepare the response
	response := responses.NewDocumentDeleteResponse(
		models.DemoInstanceModelMeta.DocumentName,
		demoInstance.Id,
	)

	// Send the response
	responses.SendResponse(r, w, http.StatusOK, response)
}

func init() {
	DocumentRequestHandler{
		Path:          fmt.Sprintf("/%s", models.DemoInstanceModelMeta.DocumentName),
		IndexHandler:  handle_demoInstance_index,
		CreateHandler: nil,
		ReadHandler:   handle_demoInstance_read,
		UpdateHandler: nil,
		DeleteHandler: handle_demoInstance_delete,
	}.RegisterHandlers()
}
