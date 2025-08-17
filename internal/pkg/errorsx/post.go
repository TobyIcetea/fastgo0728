package errorsx

import "net/http"

// ErrPostNotFound 表示未找到指定的博客.
var ErrorPostNotFound = &ErrorX{Code: http.StatusNotFound, Reason: "NotFound.PostNotFound", Message: "Post not found."}
