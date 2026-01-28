# 2.1 Middleware

This middleware will register the amount of time the people visit the app.

1. <main.go.> Crete a global struct type:
type apiConfig struct {
 fileserverHits atomic.Int32
}

2. inside main function, initialize a variable with apiConfig struct:
 apiCfg := apiConfig{
  fileserverHits: atomic.Int32{},
 }

3. <metrics.go>. Create a method to apiConfig to handle the metrics:
func (cfg *apiConfig) middlewareMetricsInc(next http.Handler) http.Handler {
 return http.HandlerFunc(func(w http.ResponseWriter, r*http.Request) {
  cfg.fileserverHits.Add(1)
  next.ServeHTTP(w, r)
 })
}

4. <main.go.> Wrap the second argument that is passed to the handle with the middleware:
 fsHandler := apiCfg.middlewareMetricsInc(http.StripPrefix("/app", http.FileServer(http.Dir(filepath))))

5. pass both arguments to the Handle method of the mux:
 mux.Handle("/app/", fsHandler)

# 2.2 Stateful Handlers

1. <metrics.go>. Create a method of apiConfig to show (handle) tha amount of hit to the app:
func (cfg *apiConfig) handlerMetrics(w http.ResponseWriter, r*http.Request) {
 w.Header().Add("Content-Type", "text/plain; charset=utf-8")
 w.WriteHeader(http.StatusOK)
 w.Write([]byte(fmt.Sprintf("Hits: %d", cfg.fileserverHits.Load())))
}

2. <main.go>. Call the HandleFunc method, passing the URL path and the method I create to show the metrics:
 mux.HandleFunc("/metrics", apiCfg.handlerMetrics)

 3. <reset.go>. Reset the count to cero
 func (cfg *apiConfig) handlerReset(w http.ResponseWriter, r*http.Request) {
 cfg.fileserverHits.Store(0)
 w.WriteHeader(http.StatusOK)
 w.Write([]byte("Hits reset to 0"))
}

4. <main.go>. Call the HandleFunc method, passing the URL path and the method I create to reset the metrics:
 mux.HandleFunc("/reset", apiCfg.handlerReset)

# 2.3 Routing
 Using HTTP methods (GET, POST)

  mux.HandleFunc("GET /healthz", handlerReadiness)
 mux.HandleFunc("POST /reset", apiCfg.handlerReset)
 mux.HandleFunc("GET /metrics", apiCfg.handlerMetrics)


* I create a collection in Postman to handle this requests!

