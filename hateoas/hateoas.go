package hateoas

import (
	"fmt"
	"reflect"
)

type link struct {
	Href   string `json:"href" example:"http(s)://<DOMAIN_OR_IP>/item/{id}"`
	Method string `json:"method" example:"GET"`
	Doc    string `json:"doc" example:"http(s)://<DOMAIN_OR_IP>/doc/index.html"`
	Key    string `json:"key,omitempty"`
}

type Hateoas struct {
	Links   []link `json:"_links" example:"http(s)://<DOMAIN_OR_IP>/item/{id}"`
	Embeded []link `json:"_embeded" example:"http(s)://<DOMAIN_OR_IP>/item/{id}"`
}

func Generate(item interface{}, resource, baseUrl string) {
	id := reflect.ValueOf(item).Elem().FieldByName("ID").Int()
	next := id + 1
	var prev int64 = 0

	embeded := []link{
		{
			Href:   fmt.Sprintf("%s/%s/%d", baseUrl, resource, id),
			Method: "GET",
			Doc:    fmt.Sprintf("%s/doc/index.html#/%s/get_%s", baseUrl, resource, resource),
		},
		{
			Href:   fmt.Sprintf("%s/%s/%d", baseUrl, resource, id),
			Method: "PUT",
			Doc:    fmt.Sprintf("%s/doc/index.html#/%s/put_%s__id_", baseUrl, resource, resource),
		},
		{
			Href:   fmt.Sprintf("%s/%s/%d", baseUrl, resource, id),
			Method: "DELETE",
			Doc:    fmt.Sprintf("%s/doc/index.html#/%s/delete_%s__id_", baseUrl, resource, resource),
		},
	}

	links := []link{
		{
			Key:    "new",
			Href:   fmt.Sprintf("%s/%s", baseUrl, resource),
			Method: "POST",
			Doc:    fmt.Sprintf("%s/doc/index.html#/%s/post_%s", baseUrl, resource, resource),
		},
		{
			Key:    "next",
			Href:   fmt.Sprintf("%s/%s/%d", baseUrl, resource, next),
			Method: "GET",
			Doc:    fmt.Sprintf("%s/doc/index.html#/%s/get_%s", baseUrl, resource, resource),
		},
	}

	if id > 0 {
		prev = id - 1

		links = append(links, link{
			Key:    "prev",
			Href:   fmt.Sprintf("%s/%s/%d", baseUrl, resource, prev),
			Method: "GET",
			Doc:    fmt.Sprintf("%s/doc/index.html#/%s/get_%s", baseUrl, resource, resource),
		})
	}

	reflectLinks := reflect.ValueOf(links)
	reflectEmbeded := reflect.ValueOf(embeded)
	reflect.ValueOf(item).Elem().FieldByName("Links").Set(reflectLinks)
	reflect.ValueOf(item).Elem().FieldByName("Embeded").Set(reflectEmbeded)
}
