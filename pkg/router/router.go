package router

import (
	"context"
	"fmt"
	"net/http"
	"regexp"
	"strconv"
	"strings"
)

const (
	digitRegExp  = "([0-9]+)"
	uuidV4Regexp = "([a-fA-F0-9]{8}-[a-fA-F0-9]{4}-4[a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12})"
	emailRegexp  = "(\\S+@\\S+\\.\\S+)"
	tokenRegexp  = "([a-zA-Z0-9]+)"
)

type HandlerFunc func(w http.ResponseWriter, r *http.Request) error

func NopHandler(fn http.HandlerFunc) HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) error {
		fn(w, r)
		return nil
	}
}

type Route struct {
	method  string
	regex   *regexp.Regexp
	handler HandlerFunc
}

func newRoute(method, pattern string, handler HandlerFunc) Route {
	tags := map[string]string{
		":digit": digitRegExp,
		":uuid":  uuidV4Regexp,
		":email": emailRegexp,
		":token": tokenRegexp,
	}
	for k, v := range tags {
		pattern = strings.ReplaceAll(pattern, k, v)
	}
	return Route{method, regexp.MustCompile("^" + pattern + "$"), handler}
}

func Get(pattern string, handler HandlerFunc) Route {
	return newRoute(http.MethodGet, pattern, handler)
}
func Post(pattern string, handler HandlerFunc) Route {
	return newRoute(http.MethodPost, pattern, handler)
}
func Patch(pattern string, handler HandlerFunc) Route {
	return newRoute(http.MethodPatch, pattern, handler)
}
func Delete(pattern string, handler HandlerFunc) Route {
	return newRoute(http.MethodDelete, pattern, handler)
}

func Merge(sliceOfRoutes ...[]Route) []Route {
	totalNumberOfRoute := 0
	for _, routes := range sliceOfRoutes {
		totalNumberOfRoute += len(routes)
	}
	merged := make([]Route, 0, totalNumberOfRoute)
	for _, routes := range sliceOfRoutes {
		merged = append(merged, routes...)
	}
	return merged
}

type RoutingErrorHandler func(statusCode int, err error, w http.ResponseWriter, r *http.Request)

func ServeRoutesAndHandleErrorWith(routes []Route, errorHandler RoutingErrorHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var allow []string
		for _, route := range routes {
			matches := route.regex.FindStringSubmatch(r.URL.Path)
			if len(matches) > 0 {
				if r.Method != route.method {
					allow = append(allow, route.method)
					continue
				}
				ctx := context.WithValue(r.Context(), ctxKey{}, matches[1:])
				err := route.handler(w, r.WithContext(ctx))
				if err != nil {
					errorHandler(
						http.StatusInternalServerError,
						fmt.Errorf("handler failed %w", err), w, r,
					)
				}
				return
			}
		}
		if len(allow) > 0 {
			allowed := strings.Join(allow, ", ")
			w.Header().Set("Allow", allowed)
			err := fmt.Errorf("method not allowed : allow only %s", allowed)
			errorHandler(http.StatusMethodNotAllowed, err, w, r)
			return
		}
		err := fmt.Errorf("route not found")
		errorHandler(http.StatusNotFound, err, w, r)
	}
}

type ctxKey struct{}

func GetField(r *http.Request, index int) string {
	fields, _ := r.Context().Value(ctxKey{}).([]string)
	if len(fields) <= index {
		return ""
	}
	return fields[index]
}

func GetFieldAsInt(r *http.Request, index int) (int, error) {
	return strconv.Atoi(GetField(r, index))
}
