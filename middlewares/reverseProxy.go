package middlewares

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"golang.org/x/net/http2"
	helper "github.com/RaihanMalay21/helper_TB_Berkah_Jaya"
)

func ReverseProxy(target string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
		url, err := url.Parse(target)
		if err != nil {
			log.Println("Error Parsing URL: ", err)
			message := map[string]interface{}{"message": err}
			helper.Response(w, message, http.StatusInternalServerError)
			return
		}

		proxy := httputil.NewSingleHostReverseProxy(url)

        // Set custom Transport to support HTTP/2
        proxy.Transport = &http2.Transport{}

		// Modify the request
        r.URL.Host = url.Host
        r.URL.Scheme = url.Scheme
        r.Header.Set("X-Forwarded-Host", r.Host)
        r.Host = url.Host

		proxy.ServeHTTP(w, r)
	})
}