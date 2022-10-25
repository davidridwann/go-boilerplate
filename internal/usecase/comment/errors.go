package commentUseCase

import "errors"

var ErrPostNotFound = errors.New("Post not found")
var ErrCommentNotFound = errors.New("Comment not found")
var ErrUnexpected = errors.New("Unexpected Error")
