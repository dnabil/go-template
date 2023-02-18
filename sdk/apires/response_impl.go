package apires

import (
	"log"

	"github.com/gin-gonic/gin"
)

var (
	statusSuccess = "success"
	statusFail    = "fail"
	statusError   = "error"
)

// Called in controller layer after a service method called without returning errors
// to return a success response.
//
// for 'data' parameter, make sure the type is map[string]interface{} or a struct.
func Success(c *gin.Context, httpCode int, message string, data interface{}) {
	response :=  responseSuccess{
		Status: statusSuccess,
		Message: message,
		Data: data,
	}
	c.JSON(httpCode, response)
}

// Called in controller layer to return a fail or an error response 
//
// errs receveived will be parsed to a more readable format
func FailOrError(c *gin.Context, httpCode int, message string, errs ...error){
	httpCode, parsed := parseError(httpCode, errs...)
	
	switch httpCode / 100{
		case 4: // 4xx
			c.JSON(httpCode, responseFail{
				Status: statusFail,
				Message: message,
				Data: parsed,
			})

		case 5: // 5xx
			for i, val := range errs {
				log.Printf("[%d]%T: %v, ", i, val, val)
			}
			
			c.AbortWithStatusJSON(httpCode, responseError{
				Status: statusError,
				Message: message,
			})

		default:
			c.AbortWithStatusJSON(500, responseError{
				Status: statusError,
				Message: "internal server error",
			})
	}
}

// // Called in controller layer to return a fail response 
// // (with custom body)
// func FailCustom(c *gin.Context, httpCode int, message string, data interface{}){
// 	switch httpCode / 100{
// 		case 4: // 4xx
// 			c.JSON(httpCode, responseFail{
// 				Status: statusFail,
// 				Message: message,
// 				Data: data,
// 			})
		
// 		default: 
// 			log.Fatalln("invalid http code for fail operation")
// 	}
// }

// func Error(c *gin.Context, httpCode int, message string){
// 	switch httpCode / 100{
// 		case 5: // 5xx
// 			c.JSON(httpCode, responseError{
// 				Status: statusFail,
// 				Message: message,
// 			})
		
// 		default: 
// 			log.Fatalln("invalid http code for error operation")
// 	}
// }
