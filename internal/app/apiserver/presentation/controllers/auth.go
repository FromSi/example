package controllers

import (
	"github.com/fromsi/example/internal/app/apiserver/application/cqrs"
	"github.com/fromsi/example/internal/app/apiserver/application/cqrs/queries"
	"github.com/fromsi/example/internal/app/apiserver/application/cqrs/responses"
	"github.com/fromsi/example/internal/app/apiserver/presentation/mappers"
	"github.com/fromsi/example/internal/app/apiserver/presentation/requests"
	presentationresponses "github.com/fromsi/example/internal/app/apiserver/presentation/responses"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type GinAuthController struct {
	Engine      *gin.Engine
	CommandCQRS *cqrs.CommandCQRS
	QueryCQRS   *cqrs.QueryCQRS
}

func (controller GinAuthController) ShowLogin(context *gin.Context) {
	_, err := requests.NewGinShowLoginAuthRequest(context)

	if err != nil {
		context.JSON(http.StatusBadRequest, presentationresponses.Response{
			Data: presentationresponses.ErrorResponse{
				Message: err.Error(),
			},
		})

		log.Println(err.Error())

		return
	}

	mnemonicAuthQueryResponse, err := (*controller.QueryCQRS).Ask(queries.GetMnemonicAuthQuery{})

	if err != nil {
		context.Status(http.StatusInternalServerError)

		log.Println("something went wrong")

		return
	}

	mnemonicAuthQueryResponseImplementation, exists := mnemonicAuthQueryResponse.(*responses.GetMnemonicAuthQueryResponse)

	if !exists {
		context.Status(http.StatusNotFound)

		log.Println("invalid query type")

		return
	}

	response, err := mappers.ToGinShowLoginAuthResponse(mnemonicAuthQueryResponseImplementation)

	if err != nil {
		context.JSON(http.StatusBadRequest, presentationresponses.Response{
			Data: presentationresponses.ErrorResponse{
				Message: err.Error(),
			},
		})

		log.Println(err.Error())

		return
	}

	context.JSON(http.StatusOK, response)
}
