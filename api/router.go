package api

import (
	"api/api/auth"
	authCheck "api/api/auth-check"
	"api/api/health"
	"api/api/token"
	"api/openapi"
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	middleware "github.com/deepmap/oapi-codegen/pkg/chi-middleware"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/getkin/kin-openapi/openapi3filter"
	"github.com/getkin/kin-openapi/routers"
	"github.com/getkin/kin-openapi/routers/gorillamux"
	"github.com/go-chi/chi/v5"
)

type Server struct {
	*auth.AuthServer
	*authCheck.AuthCheckServer
	*health.HealthCheckServer
	*token.TokenServer
}

func OapiResponseValidator(swagger *openapi3.T) func(next http.Handler) http.Handler {
	router, err := gorillamux.NewRouter(swagger)
	if err != nil {
		panic(err)
	}

	return func(next http.Handler) http.Handler {
		errorMessage := http.StatusText(http.StatusInternalServerError)
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// serve
			next.ServeHTTP(w, r)

			// validate response
			if statusCode, err := validateResponse(w, r, router); err != nil {
				http.Error(w, errorMessage, statusCode)
				return
			}
		})
	}
}

func validateResponse(w http.ResponseWriter, r *http.Request, router routers.Router) (int, error) {
	// Find route
	route, pathParams, err := router.FindRoute(r)
	if err != nil {
		return http.StatusBadRequest, err // We failed to find a matching route for the request.
	}
	// Validate request
	requestValidationInput := &openapi3filter.RequestValidationInput{
		Request:    r,
		PathParams: pathParams,
		Route:      route,
	}

	responseHeaders := http.Header{"Content-Type": []string{"application/json"}}
	responseCode, _ := strconv.Atoi(w.Header().Get("status"))
	responseBody := []byte(w.Header().Get("body"))

	// Validate response
	responseValidationInput := &openapi3filter.ResponseValidationInput{
		RequestValidationInput: requestValidationInput,
		Status:                 responseCode,
		Header:                 responseHeaders,
	}
	responseValidationInput.SetBodyBytes(responseBody)

	if err := openapi3filter.ValidateResponse(context.Background(), responseValidationInput); err != nil {
		fmt.Println(err)
		return 500, err
	}
	return 0, nil
}

func Init() {
	swagger, err := openapi.GetSwagger()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error loading swagger spec¥n: %s", err)
		os.Exit(1)
	}
	swagger.Servers = nil
	r := chi.NewRouter()
	r.Use(middleware.OapiRequestValidatorWithOptions(swagger, &middleware.Options{
		ErrorHandler: func(w http.ResponseWriter, message string, statusCode int) {
			if message != "no matching operation was found" {
				fmt.Println(struct {
					message    string
					statusCode int
				}{
					message:    message,
					statusCode: statusCode,
				})
			}
			http.Error(w, "", statusCode)
		},
	}))
	// r.Use(OapiResponseValidator(swagger))
	// r.Use(chi_middleware.RealIP)
	// r.Use(chi_middleware.Logger)
	openapi.HandlerFromMux(&Server{}, r)

	port := "3000"
	log.Printf("port: %s", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), r))
}
