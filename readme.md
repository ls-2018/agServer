func (s *APIVersionHandler) AddToWebService(ws *restful.WebService) {
	mediaTypes, _ := negotiation.MediaTypesForSerializer(s.serializer)
	ws.Route(ws.GET("/").To(s.handle).
		Doc("get available resources").
		Operation("getAPIResources").
		Produces(mediaTypes...).
		Consumes(mediaTypes...).
		Writes(metav1.APIResourceList{}))
}




func DefaultBuildHandlerChain(apiHandler http.Handler, c *Config) http.Handler {

func (d director) ServeHTTP(w http.ResponseWriter, req *http.Request)
func withAuthorization(handler http.Handler, a authorizer.Authorizer, s runtime.NegotiatedSerializer, metrics recordAuthorizationMetricsFunc) http.Handler {
func WithImpersonation(handler http.Handler, a authorizer.Authorizer, s runtime.NegotiatedSerializer) http.Handler {
func WithAudit(handler http.Handler, sink audit.Sink, policy audit.PolicyRuleEvaluator, longRunningCheck request.LongRunningRequestCheck) http.Handler {
func withAuthentication(handler http.Handler, auth authenticator.Request, failed http.Handler, apiAuds authenticator.Audiences, requestHeaderConfig *authenticatorfactory.RequestHeaderConfig, metrics authenticationRecordMetricsFunc) http.Handler {
func WithCORS(handler http.Handler, allowedOriginPatterns []string, allowedMethods []string, allowedHeaders []string, exposedHeaders []string, allowCredentials string) http.Handler {
func (t *timeoutHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
func withRequestDeadline(handler http.Handler, sink audit.Sink, policy audit.PolicyRuleEvaluator, longRunning request.LongRunningRequestCheck,
func withWaitGroup(handler http.Handler, longRunning apirequest.LongRunningRequestCheck, wg RequestWaitGroup, isRequestExemptFn isRequestExemptFunc) http.Handler {






