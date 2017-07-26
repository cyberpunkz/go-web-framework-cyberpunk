# Cyberpunk Web Framework on Golang
alpha stage

## Simple example

We suppose you have an **src** directory in root of your project.

### Create src/view/html/page/index.html

```html
<p>{{ .Title }} world</p>
```
### Create src/main.go

```go
package main

import (
	"github.com/cyberpunkz/cyberpunk"
	"net/http"
)

func main() {
	config := cyberpunk.GetSystemConfig()
	config.SetFrontendDir("src/view")

	templateRegister := cyberpunk.GetTemplateRegister()
	templateRegister.Register("index")
	templateRegister.AddTemplate("index", "page/index")

	cyberpunk.AddNewRoute("/", indexPage)

	cyberpunk.StartNewHttpServer("1234")
}

func indexPage(w http.ResponseWriter, r *http.Request) {
	type Page struct {
		Title string
	}

	page := Page{Title:"hello"}
	cyberpunk.RenderLayout(w, "index", page)
}
```

### Running

With command ```go run main.go``` you can see the webpage on your localhost: http://localhost:1234/
