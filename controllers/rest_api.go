package controllers

import (
	"github.com/dongdonjuhahaha/traffic/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

var returns = models.Response{
	ID: "dd",
	ORDER_ID: "123",
	ORDER_NUMBER: "1",
	ORDER_ITEM_ID: "2",
	ORDER_ITEM_NUMBER: "3",
	BARCODE: "4",
	RETURN_QUANTITY: "5",
	SHIPMENT_NUMBER: "6",
	SENDER_NAME: "7",
	SENDER_PHONE: "8",
	SENDER_ADDRESS1: "9",
	SENDER_ADDRESS2: "00",
	SENDER_ZIPCODE: "123",

}
func GetReturn(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, returns )
}
