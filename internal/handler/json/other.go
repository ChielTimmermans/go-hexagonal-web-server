package json

import (
	"go-hexagonal/internal"
	"go-hexagonal/internal/domain/errormessage"
	"go-hexagonal/internal/domain/language"
	"log"
	"strings"

	jsoniter "github.com/json-iterator/go"
	"github.com/valyala/fasthttp"
)

type Answer struct {
	CodeToError    map[language.Code]map[string]errormessage.Error
	jsonStreamPool jsoniter.StreamPool
}

type StatusOKAnswer struct {
	Data interface{} `json:"data"`
}

//nolint will be long bc off error messages
func InitErrors(jsp jsoniter.StreamPool) *Answer {
	err := Answer{
		jsonStreamPool: jsp,
		CodeToError: map[language.Code]map[string]errormessage.Error{
			language.ENG: map[string]errormessage.Error{
				"name_required": errormessage.Error{
					Title:  "Name required",
					Detail: "The name is required.",
				},
				"name_length": errormessage.Error{
					Title:  "Name length incorrect",
					Detail: "The name length must be between 3 and 255 characters.",
				},
				"email_required": errormessage.Error{
					Title:  "Email required",
					Detail: "The email is required.",
				},
				"email_length": errormessage.Error{
					Title:  "Email length incorrect",
					Detail: "The email length must be between 3 and 255 characters.",
				},
				"email_wrong_format": errormessage.Error{
					Title:  "Email incorrect format",
					Detail: "The email must be of format email@example.com.",
				},
				"role_required": errormessage.Error{
					Title:  "Role required",
					Detail: "The role is required.",
				},
				"role_not_found": errormessage.Error{
					Title:  "Role not found",
					Detail: "The role must be user, customeradmin or admin",
				},
				"password_required": errormessage.Error{
					Title:  "Password required",
					Detail: "The password is required.",
				},
				"password_length": errormessage.Error{
					Title:  "Password length incorrect",
					Detail: "The password length must be between 3 and 255 characters.",
				},
			},
		},
	}
	return &err
}

func (a *Answer) SetAnswer(ctx *fasthttp.RequestCtx, statusCode int, data interface{}) {
	if statusCode == fasthttp.StatusOK {
		a.setStatusOK(ctx, data)
	} else if statusCode == fasthttp.StatusUnprocessableEntity {
		if err, ok := data.(error); ok {
			a.setValidationErrors(ctx, err)
		}
	} else {
		log.Println(data)
		ctx.SetStatusCode(statusCode)
	}
}

func (a *Answer) setStatusOK(ctx *fasthttp.RequestCtx, data interface{}) {
	jsonStream := a.jsonStreamPool.BorrowStream(nil)
	defer a.jsonStreamPool.ReturnStream(jsonStream)
	jsonStream.WriteVal(StatusOKAnswer{
		Data: data,
	})

	ctx.SetBody(jsonStream.Buffer())
	ctx.SetStatusCode(fasthttp.StatusOK)
}

func (a *Answer) setValidationErrors(ctx *fasthttp.RequestCtx, err error) {
	statusCode := fasthttp.StatusUnprocessableEntity
	requestID := getRequestID(ctx)
	lang := getLanguage(ctx)

	errors := &errormessage.Errors{
		ID:         requestID,
		StatusCode: statusCode,
	}

	errors.Errors = a.makeErrors(err.Error(), lang)

	jsonStream := a.jsonStreamPool.BorrowStream(nil)
	defer a.jsonStreamPool.ReturnStream(jsonStream)
	jsonStream.WriteVal(errors)

	ctx.SetBody(jsonStream.Buffer())
	ctx.SetStatusCode(statusCode)
}

// TODO fix this mess
func (a *Answer) makeErrors(err string, lang language.Code) *[]errormessage.Error {
	splittedError := SplitErr(err)
	errors := make([]errormessage.Error, len(splittedError))
	for k, v := range splittedError {
		key, value := getKeyValue(v)
		errors[k] = a.getError(value, lang, errormessage.Source{
			Pointer: key,
		})
	}
	return &errors
}

func SplitErr(err string) []string {
	s := err
	sep := ";"
	n := strings.Count(s, sep)
	var a []string

	i := 0
	for i < n {
		m := strings.Index(s, sep)
		if m < 0 {
			break
		}
		i++
		x := strings.Index(s, "(")
		y := strings.Index(s, ")")
		if m > x && m < y {
			continue
		}
		v := s[:m+0]
		a = appendBasedOnType(a, v)

		s = s[m+len(sep):]
	}
	a = appendBasedOnType(a, s)
	return a
}

func appendBasedOnType(slice []string, s string) []string {
	if !strings.Contains(s, "(") && !strings.Contains(s, ")") {
		slice = append(slice, s)
	} else {
		slice = append(slice, getErrorsBasedOnNesting(s)...)
	}
	return slice
}

func getErrorsBasedOnNesting(s string) []string {
	data := strings.SplitN(s, ":", 2)
	mainKey := data[0]
	errors := data[1]
	splittedErrors := strings.Split(errors, ";")
	a := make([]string, len(splittedErrors))
	for k, v := range splittedErrors {
		a[k] = strings.ReplaceAll(strings.ReplaceAll(strings.ReplaceAll(internal.Concat(mainKey, "/", v), " ", ""), "(", ""), ").", "")
	}
	return a
}

func (a *Answer) getError(code string, lg language.Code, source errormessage.Source) errormessage.Error {
	err := a.CodeToError[lg][code]
	return errormessage.Error{
		Title:  err.Title,
		Detail: err.Detail,
		Source: source,
	}
}

func getKeyValue(v string) (key, value string) {
	data := strings.Split(v, ":")

	// removing artifacts from err.error()
	if len(data) < 2 {
		return data[0], ""
	}
	key = strings.ToLower(strings.ReplaceAll(strings.ReplaceAll(data[0], ".", ""), " ", ""))
	value = strings.ReplaceAll(strings.ReplaceAll(data[1], ".", ""), " ", "")
	return key, value
}

func getRequestID(ctx *fasthttp.RequestCtx) []byte {
	var requestID []byte
	var ok bool

	if requestID, ok = ctx.UserValue("requestId").([]byte); !ok {
		requestID = []byte("\"request ID not found\"")
	}
	return requestID
}

func getLanguage(ctx *fasthttp.RequestCtx) language.Code {
	var lang string
	var ok bool

	if lang, ok = ctx.UserValue("language").(string); !ok {
		lang = "ENG"
	}
	return language.Code(lang)
}
