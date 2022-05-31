// The safe_http package consists of wrappers around go's built-in HTTP request
// functionality. The purpose of the wrappers is to add default time-out conditions
// and to prevent leaks to/from the requested server.

package safe_http

import (
	//"encoding/json"
	"io"
	"log"
	"net/http"
	"time"
)

// Safe_http_get creates a net client with a max wait time passed by the caller.
// It also closes the request at completion of the function and adds some error handling.
func Safe_http_get(url string, wait_time int) ([]byte, error) {
	
	var netClient = &http.Client {
		Timeout: time.Second * time.Duration(wait_time), // Custom timeout
	} // TODO : Communictate get status to server, log meta data
	
	req, _ := http.NewRequest("GET", url, nil)

	response, err := netClient.Do(req)
	if err != nil {
		log.Fatal(err)
	} // Checks if request returned error, prints error and ends execution otherwise

	defer response.Body.Close() // Close the request to prevent leaks

	return io.ReadAll(response.Body)
}