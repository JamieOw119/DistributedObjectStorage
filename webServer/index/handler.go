package index

import (
	"bufio"
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
)

type Metadata struct {
	Name    string
	Version int
	Size    int64
	Hash    string
}

func Handler(w http.ResponseWriter, r *http.Request) {
	req, e := http.Get("http://" + os.Getenv("API_SERVER") + "/versions/")
	if e != nil {
		log.Println(e)
		return
	}
	s := bufio.NewScanner(req.Body)

	t, e := template.ParseFiles("view/filelist.html")
	if e != nil {
		log.Println(e)
		return
	}

	var start map[string]int = map[string]int{}
	var metas []Metadata
	var lastname string
	var lastindex int

	for s.Scan() {
		var meta Metadata
		json.Unmarshal([]byte(s.Text()), &meta)
		if lastname != meta.Name {
			lastname = meta.Name
			lastindex = 1
			start[lastname] = lastindex
		}
		if meta.Hash == "" {
			start[lastname] = -1
		} else {
			if start[lastname] == -1 {
				start[lastname] = lastindex
			}
		}
		metas = append(metas, meta)
		lastindex++
	}

	if len(metas) != 0 {
		var build strings.Builder
		lastname = metas[0].Name
		lastindex = 1
		for _, meta := range metas {
			if lastname != meta.Name {
				lastname = meta.Name
				lastindex = 1
			}
			if lastindex >= start[lastname] && start[lastname] != -1 {
				realname, _ := url.QueryUnescape(meta.Name)
				l1 := fmt.Sprintf("<tr><td><a href=/download?name=%s>%s</a></td><td>%d</td><td>%d</td><td>%s</td>", meta.Name, realname, meta.Version, meta.Size, meta.Hash)
				l2 := fmt.Sprintf(`<td><a class="layui-btn layui-btn-xs" href=/download?name=%s&version=%d>下载</a>
					<a class="layui-btn layui-btn-xs layui-btn-danger" href="/delete?name=%s" onclick="javascript:return del(%s)">删除</a></td></tr>`, meta.Name, meta.Version, meta.Name, meta.Name)
				build.WriteString(l1)
				build.WriteString(l2)
			}
			lastindex++
		}
		t.Execute(w, template.HTML(build.String()))
	} else {
		t.Execute(w, template.HTML(""))
	}
}
