package api

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"regexp"
	"strings"

	"maps"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	pb "github.com/muharik19/boiler-plate-grpc/api/grpc/api/pb/v1/role"
	"github.com/muharik19/boiler-plate-grpc/configs"
	"github.com/muharik19/boiler-plate-grpc/internal/constant"
	"github.com/muharik19/boiler-plate-grpc/internal/domain/entities"
	"github.com/muharik19/boiler-plate-grpc/internal/pkg/utils"
	"github.com/muharik19/boiler-plate-grpc/pkg/logger"
	global "github.com/muharik19/boiler-plate-grpc/pkg/utils"
	"github.com/rs/cors"
	"go.elastic.co/apm/module/apmgrpc"
	"go.elastic.co/apm/module/apmhttp"
	"google.golang.org/grpc"
)

func NewHttpServer() error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	// Connect to the GRPC server
	addr := "0.0.0.0:" + *global.Getenv("GRPC_PORT")
	// addr := "0.0.0.0:" + "9090"
	conn, err := grpc.Dial(addr,
		grpc.WithInsecure(),
		grpc.WithStreamInterceptor(apmgrpc.NewStreamClientInterceptor()),
		grpc.WithUnaryInterceptor(apmgrpc.NewUnaryClientInterceptor()))
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
		panic(err)
	}
	defer conn.Close()

	// Create new grpc-gateway
	rmux := runtime.NewServeMux(runtime.WithIncomingHeaderMatcher(customMatcher))

	// register gateway endpoints
	for _, f := range []func(ctx context.Context, mux *runtime.ServeMux, conn *grpc.ClientConn) error{
		// register grpc service handler
		pb.RegisterRoleHandler,
	} {
		if err = f(ctx, rmux, conn); err != nil {
			log.Fatal(err)
			panic(err)
		}
	}

	// create http server mux
	mux := http.NewServeMux()
	mux.Handle("/", rmux)

	corsMiddleware := cors.New(cors.Options{
		AllowedOrigins:   global.GetEnvCors("CORS_ORIGIN_ALLOWED"),
		AllowedMethods:   global.GetEnvCors("CORS_METHOD_ALLOWED"),
		AllowedHeaders:   global.GetEnvCors("CORS_HEADER_ALLOWED"),
		AllowCredentials: false,
	})

	// running rest http server
	logger.Infof("Serving Rest Http on 0.0.0.0: %v", *global.Getenv("HTTP_PORT"))
	err = http.ListenAndServe("0.0.0.0:"+*global.Getenv("HTTP_PORT"), corsMiddleware.Handler(http.HandlerFunc(func(w http.ResponseWriter, request *http.Request) {
		wHead := w.Header()
		rHead := request.Header
		wHead.Set(configs.ACAM, strings.Join(global.GetEnvCors("CORS_METHOD_ALLOWED"), ", "))
		wHead.Set(configs.ACAH, configs.ACAH_VALUE)
		wHead.Set(configs.ACAC, configs.ACAC_VALUE)
		wHead.Set(configs.HSTS, configs.HSTS_VALUE)
		wHead.Set(configs.CC, configs.CC_VALUE)
		wHead.Set(configs.XCTO, configs.XCTO_VALUE)
		wHead.Set("Content-Security-Policy", "default-src 'self'")

		rHead.Set(configs.GRPC_METHOD, strings.Join(global.GetEnvCors("CORS_METHOD_ALLOWED"), ", "))
		var bodyString string
		if request.Body != nil {
			buf := new(strings.Builder)
			_, err := io.Copy(buf, request.Body)
			if err != nil {
				http.Error(w, "can't read body", http.StatusBadRequest)
				return
			}

			bodyString, err = validateRequestBody(buf.String())
			if err != nil {
				http.Error(w, "can't read body", http.StatusBadRequest)
				return
			}

			request.Body = io.NopCloser(strings.NewReader(bodyString))
		}

		rec := httptest.NewRecorder()
		mux.ServeHTTP(rec, request)
		apmhttp.Wrap(mux)

		maps.Copy(w.Header(), rec.Header())
		// grab the captured response body
		data := rec.Body.Bytes()

		response := entities.Response{}
		json.Unmarshal(data, &response)

		// var bodyInterface any
		// json.Unmarshal([]byte(bodyString), &bodyInterface)
		// bodyMarshal, _ := json.Marshal(bodyInterface)

		// REPOSITORY LOGGER
		// logger.Infof("[boiler-plate-grpc:log] [RequestURL] : %s, [RequestMethod] : %s, [RequestBody] : %s, [ResponseData] : %s", request.RequestURI, request.Method, string(bodyMarshal), string(data))

		if request.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
		} else {
			if utils.ConvertStatusResponseCode(response.ResponseCode) == 500 {
				response = entities.Response{
					ResponseCode: constant.FAILED_INTERNAL,
					ResponseDesc: http.StatusText(http.StatusInternalServerError),
				}
				data, _ = json.Marshal(response)
			}
			w.WriteHeader(int(utils.ConvertStatusResponseCode(response.ResponseCode)))
		}

		i, err := w.Write(data)
		if err != nil {
			logger.Infof("Error Write Data %d : %v", i, err.Error())
		}
	})))
	if err != nil {
		log.Fatal(err)
		panic(err)
	}
	return nil
}

// CustomMatcher is a Mapping from HTTP request headers to gRPC client metadata
func customMatcher(key string) (string, bool) {
	switch key {
	case configs.ACCEPT:
		return key, true
	case configs.AUTHORIZATION:
		return key, true
	default:
		return runtime.DefaultHeaderMatcher(key)
	}
}

func validateRequestBody(bodyString string) (string, error) {
	// validate if thre is some special character
	positiveinput := regexp.MustCompile(`^[a-zA-Z0-9_.,:/()\t\n\v\f\r]*$`).MatchString
	inputan := positiveinput(bodyString)
	if !strings.Contains(bodyString, "http://") && !strings.Contains(bodyString, "https://") && !strings.Contains(bodyString, "://") && !strings.Contains(bodyString, "./") {
		if strings.Contains(bodyString, "data:image/jpeg;base64") {
			return bodyString, nil
		} else if !inputan {
			output := replaceinputhelp(bodyString)
			return output, nil
		} else {
			return bodyString, nil
		}
	} else {
		return "", fmt.Errorf("body contain unknown type")
	}
}

func replaceinputhelp(input string) (output string) {
	request := "{\"body\":" + input + "}"
	output = request
	return output
}
