package utils

// ErrorJson formats Error Json Response
func ErrorJson(err string, code string) map[string]interface{} {

	var errorsSlice []map[string]interface{}
	var initError = map[string]interface{}{
		"code":    code,
		"message": err,
	}

	errorsSlice = append(errorsSlice, initError)
	newError := map[string]interface{}{
		"data":   map[string]interface{}{},
		"errors": errorsSlice,
	}
	return newError
}

// JsonResp Json formats single object responses
func JsonResp(resp interface{}) map[string]interface{} {

	newResp := map[string]interface{}{
		"data": resp,
	}
	return newResp
}
