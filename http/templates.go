package http

import "toast.cafe/x/brpaste/v2/template"

var (
	// CodeTemplate is the function to be used as the template for the code viewer
	// overwrite it before calling GenHandler
	CodeTemplate = template.Code
	// IndexTemplate is the function to be used as the template for the index page
	// overwrite it before calling GenHandler
	IndexTemplate = template.Index
)
