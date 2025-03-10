package routers

import (
	"testing"

	"github.com/gorilla/mux"
	corev2 "github.com/sensu/sensu-go/api/core/v2"
	"github.com/sensu/sensu-go/testing/mockstore"
)

func TestNamespacesRouter(t *testing.T) {
	// Setup the router
	s := &mockstore.MockStore{}
	router := NewNamespacesRouter(s)
	parentRouter := mux.NewRouter().PathPrefix(corev2.URLPrefix).Subrouter()
	router.Mount(parentRouter)

	empty := &corev2.Namespace{}
	fixture := corev2.FixtureNamespace("foo")

	tests := []routerTestCase{}
	tests = append(tests, getTestCases(fixture)...)
	tests = append(tests, listTestCases(empty)...)
	tests = append(tests, createTestCases(empty)...)
	tests = append(tests, updateTestCases(fixture)...)
	tests = append(tests, deleteTestCases(fixture)...)
	for _, tt := range tests {
		run(t, tt, parentRouter, s)
	}
}
