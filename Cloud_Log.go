package Cloud_Log
import(
	"net/http"
	"appengine"
)

func AppEngineLog( r *http.Request, log string){
	c := appengine.NewContext(r)
	c.Infof("%s",log)
}