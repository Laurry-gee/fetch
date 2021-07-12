package config

type Request struct {
	Method string
	URL    string
	Body   string

	Header map[string]string
}

type Assert struct {
	Status string

	StatusIn []string `mapstructure:"status_in"`

	StatusCode    int
	StatusCodeIn  []int `mapstructure:"statusCode_in"`
	StatusCodeLt  int   `mapstructure:"statusCode_lt"`
	StatusCodeLte int   `mapstructure:"statusCode_lte"`
	StatusCodeGt  int   `mapstructure:"statusCode_gt"`
	StatusCodeGte int   `mapstructure:"statusCode_gte"`

	ContentLength    int64
	ContentLengthLt  int64 `mapstructure:"contentLength_lt"`
	ContentLengthLte int64 `mapstructure:"contentLength_lte"`
	ContentLengthGt  int64 `mapstructure:"contentLength_gt"`
	ContentLengthGte int64 `mapstructure:"contentLength_gte"`

	// TODO: header
	ContentType string

	// latency
	LatencyLt  int64 `mapstructure:"latency_lt"`
	LatencyLte int64 `mapstructure:"latency_lte"`
	LatencyGt  int64 `mapstructure:"latency_gt"`
	LatencyGte int64 `mapstructure:"latency_gte"`

	Body string

	BodyContains    string `mapstructure:"body_contains"`
	BodyNotContains string `mapstructure:"body_not_contains"`
	BodyStartsWith  string `mapstructure:"body_startswith"`
	BodyEndsWith    string `mapstructure:"body_endswith"`

	//Json []map[string]interface{} `mapstructure:"json"`
	Json []AssertJson `mapstructure:"json"`
	// TODO: json body assert?
}

type Case struct {
	Title       string
	Description string

	Request Request
	Assert  Assert
}

type AssertJson struct {
	Path  string
	Value interface{}
}