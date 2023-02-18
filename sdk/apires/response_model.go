// apires package helps writing API-RESponses in json format using *gin.Context.JSON()
// to keep it consistent and readable.
/*
	success		= all went well  (2xx)
	fail		= api is not satisfied (user error, 4xx)
	error		= server error (5xx)
*/
// Success() to write a success response.
// FailOrError() to write a fail or an error response
package apires

type responseSuccess struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type responseFail struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type responseError struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}
