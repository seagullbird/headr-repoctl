// Package servicetest is a unittest package against the service.Service interface.
// Tested implementations include the basicService struct and all application level middlewares.
// This package is inspired by the test mechanism of golang.org/x/net/nettest which tests a net.Conn implementation satisfies its interface,
// but modified the structured to fit the chained middleware style of gokit.
// It is intended to cover every line of code in package service.
// The only implementation of service.Service that is not covered by this package is endpoint.Set, which I think should be covered on a integration test level.
package servicetest
