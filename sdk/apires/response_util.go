package apires

import (
	"encoding/json"
	"fmt"
	"regexp"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-sql-driver/mysql" // for error handling
)

// return type is string
func parseError(code int, errs... error) (int, map[string]any){
	parsed := map[string]any{}
	errorSlice := make([]string, 0, len(errs))

	for _, err := range errs{

		switch errorType := any(err).(type) {
			// json unmarshalling-error
			case *json.UnmarshalTypeError:
				val := *errorType
				errorSlice = append(errorSlice, fmt.Sprintf("(json-err) %s must be %s", val.Field, val.Type.String()))
		
			// json syntax error
			case *json.SyntaxError:
				errorSlice = append(errorSlice, fmt.Sprintf("(json-err) %s (%d)", errorType.Error(), errorType.Offset))
			
			// *MYSQL* related error
			case *mysql.MySQLError:
				switch errorType.Number{
				
				// unique constraint error 	
				case 1062: // *note: delete this if https://github.com/go-gorm/gorm/issues/5651
					// Error format (mysql) = Duplicate entry '<value>' for key '<field>' 
					strArr := regexp.MustCompile("([\"'])(.*?)([\"'])").FindAllString(errorType.Message, 2)
					field := strArr[1][1:len(strArr[1])-1]
					code = 409 //StatusConflict
					errorSlice = append(errorSlice, fmt.Sprintf("%s already exists", field))
					
					default: 
					errorSlice = append(errorSlice, errorType.Message)
				}
				
			case validation.Errors:
				errorSlice = append(errorSlice, "validation error")
				for i, val := range errorType{
					parsed[i] = val.Error()
				}
				
			case *validation.InternalError:
				code = 500 //internal server error
				break //out of loop
				
			default:
				var defaultError string 
				
				switch err.Error(){
					case "EOF":
						defaultError = "End Of File"
					case "unexpected EOF":
						defaultError = "unexpected End Of File" 
					default:
						defaultError = err.Error()
				}
		
				errorSlice = append(errorSlice, defaultError)
				
		} // end of switch
	} // end of for loop

	parsed["error"] = errorSlice
	return code, parsed
}