package dnac

// TaskResponse is the Response definition
type TaskResponse struct {
	TaskID string `json:"taskId,omitempty"` //
	URL    string `json:"url,omitempty"`    //
}

// CountResult is the CountResult definition
type CountResult struct {
	Response int    `json:"response,omitempty"` //
	Version  string `json:"version,omitempty"`  //
}
