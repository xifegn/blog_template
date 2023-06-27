# blog_template
修改handler模板,使其返回格式统一

    package handler
    import (
    	"net/http"
      "blog_template/response"// 改成你的response包
    	"github.com/zeromicro/go-zero/rest/httpx"
    	{{.ImportPackages}}
    )
    
    
    func {{.HandlerName}}(svcCtx *svc.ServiceContext) http.HandlerFunc {
        return func(w http.ResponseWriter, r *http.Request) {
            {{if .HasRequest}}var req types.{{.RequestType}}
            if err := httpx.Parse(r, &req); err != nil {
                httpx.Error(w, err)
                return
            }{{end}}
    
            l := {{.PkgName}}.New{{.LogicType}}(r.Context(), svcCtx)
            {{if .HasResp}}resp, {{end}}err := l.{{.Call}}({{if .HasRequest}}&req{{end}})
            {{if .HasResp}}response.Response(w, resp, err){{else}}response.Response(w, nil, err){{end}}
        }
    }

response/response.go

      package response

      import (
      	"net/http"
      
      	"github.com/zeromicro/go-zero/rest/httpx"
      )
      
      type Body struct {
      	Code int         `json:"code"`
      	Msg  string      `json:"msg"`
      	Data interface{} `json:"data,omitempty"`
      }
      
      func Response(w http.ResponseWriter, resp interface{}, err error) {
      	var body Body
      	if err != nil {
      		body.Code = -1
      		body.Msg = err.Error()
      	} else {
      		body.Msg = "OK"
      		body.Data = resp
      	}
      	httpx.OkJson(w, body)
      }
