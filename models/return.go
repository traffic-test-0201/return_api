package models


type Payload struct{
	RESULT          bool `json:"result"`
	Title       string `json:"message"`
	Description string `json:"data"`
}
type Response struct{
	ID          string `json:"id"`
	ORDER_ID       string `json:"order_id"`
	ORDER_NUMBER string `json:"order_number"`
	ORDER_ITEM_ID string `json:"order_item_id"`
	ORDER_ITEM_NUMBER string `json:"order_item_number"`
	BARCODE string `json:"barcode"`
	RETURN_QUANTITY string `json:"return_quantity"`
	SHIPMENT_NUMBER string `json:"shipment_number"`
	SENDER_NAME string `json:"sender_name"`
	SENDER_PHONE string `json:"sender_phone"`
	SENDER_ADDRESS1 string `json:"sender_address1"`
	SENDER_ADDRESS2 string `json:"sender_address2"`
	SENDER_ZIPCODE string `json:"sender_zipcode"`
}

