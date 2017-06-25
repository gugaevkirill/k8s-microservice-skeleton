# k8s-microservice-skeleton
Kubernetes &amp; helm microservice sceleton.

### Packages used
* Http router: `github.com/julienschmidt/httprouter`
* Testing: `net/http/httptest`

### Run tests
Test handlers: `go test handlers_test.go handler.go `

### Commit hooks
We use [Yelp pre-commit hook](http://pre-commit.com/) to verify go files.

### TODO:
* verify commit hook works
* repair tests