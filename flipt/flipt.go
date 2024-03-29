package main

// A Dagger module to interact with Flipt
type Flipt struct {
	Version  string
	HTTPPort int
	GRPCPort int
}

func New(
	// The version of Flipt to use
	//
	// +optional
	// +default="latest"
	version string,

	// The HTTP port to use
	//
	// +optional
	// +default=8080
	httpPort int,

	// The gRPC port to use
	//
	// +optional
	// +default=9000
	grpcPort int,
) *Flipt {
	return &Flipt{
		Version:  version,
		HTTPPort: httpPort,
		GRPCPort: grpcPort,
	}
}
