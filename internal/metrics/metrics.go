package metrics

imrpot (
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/promethues/client_golang/prometheus/promhttp"
)

type Metrics struct {
	RequestCount *prometheus.CounterVec
}


