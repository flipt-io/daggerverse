package main

import "fmt"

// Server returns a new Flipt container instance
func (f *Flipt) Server() *Container {
	ctr := dag.Container().From(fmt.Sprintf("docker.flipt.io/flipt/flipt:%s", f.Version)).
		WithExposedPort(f.HTTPPort).
		WithExposedPort(f.GRPCPort)

	return ctr
}
