package gee

import (
	"fmt"
	"reflect"
	"testing"
)

func newTestRouter() *router {
	r := newRouter()
	r.addRoute("GET", "/", nil)
	r.addRoute("GET", "/hello/:name", nil)
	r.addRoute("GET", "/hello/b/c", nil)
	r.addRoute("GET", "/hi/:name", nil)
	r.addRoute("GET", "/assets/*filepath", nil)
	return r
}

func TestParsePattern(t *testing.T) {
	ok := reflect.DeepEqual(parsePattern("/p/:name"), []string{"p", ":name"})
	ok = ok && reflect.DeepEqual(parsePattern("/p/*"), []string{"p", "*"})
	ok = ok && reflect.DeepEqual(parsePattern("/p/a/c"), []string{"p", "a", "c"})
	ok = ok && reflect.DeepEqual(parsePattern("/p/*name/*"), []string{"p", "*name"})
	if !ok {
		t.Fatal("test parsePattern failed")
	}
}

func TestGetRoute(t *testing.T) {
	r := newTestRouter()
	n, ps := r.getRoute("GET", "/hello/mike")
	if n == nil {
		t.Fatal("route nil")
	}

	if n.pattern != "/hello/:name" {
		t.Fatal(n.pattern, "not match", "/hello/:name")
	}

	if ps["name"] != "mike" {
		t.Fatal(ps["name"])
	}
}

func TestGetRoute2(t *testing.T) {
	r := newTestRouter()
	n, ps := r.getRoute("GET", "/assets/mike.txt")
	if n == nil {
		t.Fatal("route nil")
	}

	if n.pattern != "/assets/*filepath" {
		t.Fatal(n.pattern, "not match", "/assets/*")
	}

	if ps["filepath"] != "mike.txt" {
		t.Fatal(ps["name"])
	}

	n, ps = r.getRoute("GET", "/assets/css/mike.css")
	if n == nil {
		t.Fatal("route nil")
	}

	if n.pattern != "/assets/*filepath" {
		t.Fatal(n.pattern, "not match", "/assets/*")
	}

	if ps["filepath"] != "css/mike.css" {
		t.Fatal(ps["name"])
	}
}

func TestGetRoute3(t *testing.T) {
	r := newTestRouter()
	n, ps := r.getRoute("GET", "/")
	if n == nil {
		t.Fatal("route nil")
	}

	if n.pattern != "/" {
		t.Fatal(n.pattern, "not match", "/")
	}

	if len(ps) != 0 {
		t.Fatal(ps["name"])
	}
}

func TestRoutes(t *testing.T) {
	r := newTestRouter()
	nodes := r.getRoutes("GET")
	for i, n := range nodes {
		fmt.Println(i+1, n)
	}

	if len(nodes) != 5 {
		t.Fatal("wrong nodes")
	}
}
