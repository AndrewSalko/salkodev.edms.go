package controller

// For GET requests (single object) URL part contains uid (last part of get url)
const UIDParam = "uid"

// For simple requests, with only one uid field
type UIDRequest struct {
	UID string `json:"uid" binding:"required"`
}
