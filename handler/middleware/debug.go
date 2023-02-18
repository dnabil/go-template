package middleware

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/gin-gonic/gin"
)

var (
	startSign string = "=====Debug======"
	endSign string = "====EndDebug===="
)

// middleware printout request (BODY)
func PrintBody() gin.HandlerFunc{
	return func(c *gin.Context) {
		body, _ := ioutil.ReadAll(c.Request.Body)
    	bodyParsed, err := debug(string(body))
		if err == nil {
			fmt.Println(startSign)
			fmt.Println(bodyParsed)
			fmt.Println(endSign)
		}
		//putting it back to body
    	c.Request.Body = ioutil.NopCloser(bytes.NewReader(body))
		c.Next()
	}
}

// middleware printout request (HEADER)
func PrintHeader() gin.HandlerFunc{
	return func(c *gin.Context) {
		headerParsed, err := debug(c.Request.Header)
		if err == nil {
			fmt.Println(startSign)
			fmt.Println(headerParsed)
			fmt.Println(endSign)
		}
		c.Next()
	}
}



func debug(i interface{}) (string, error) {
	result, err := json.MarshalIndent(i, "", "\t")
	if err != nil {
		return "", err
	}
	return string(result), err
}