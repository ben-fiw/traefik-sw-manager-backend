package handlers

import (
	"demo-shop-manager/models"
	"demo-shop-manager/responses"
	"fmt"
	"net/http"
)

func handle_availableVersion_index(w http.ResponseWriter, r *http.Request) {
	// Get the pagination parameters
	paginationOptions := models.GetPaginationParams(r, models.AvailableVersionModelMeta)

	// Load the available versions
	availableVersionList := models.NewAvailableVersionModelList()
	err := availableVersionList.Paginate(paginationOptions)
	if err != nil {
		responses.SendError(r, w, http.StatusInternalServerError, err.Error())
		return
	}

	// Load the total count
	totalCount, err := availableVersionList.GetTotalCount()
	if err != nil {
		responses.SendError(r, w, http.StatusInternalServerError, err.Error())
		return
	}

	// Prepare the response
	response := responses.NewDocumentIndexResponse(
		totalCount,
		paginationOptions.Page,
		paginationOptions.Limit,
		models.AvailableVersionModelMeta.DocumentName,
		availableVersionList,
	)

	// Send the response
	responses.SendResponse(r, w, http.StatusOK, response)
}

func init() {
	DocumentRequestHandler{
		Path:          fmt.Sprintf("/%s", models.AvailableVersionModelMeta.DocumentName),
		IndexHandler:  handle_availableVersion_index,
		CreateHandler: nil,
		ReadHandler:   nil,
		UpdateHandler: nil,
		DeleteHandler: nil,
	}.RegisterHandlers()
}
