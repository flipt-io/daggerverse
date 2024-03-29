package main

import (
	"context"
	"fmt"
)

type CLI struct {
	Ctr  *Container
	addr string
}

// CLI returns a new Flipt CLI instance
func (f *Flipt) CLI(
	ctx context.Context,
	// The service to bind to the CLI
	//
	// +optional
	svc *Service) (*CLI, error) {
	cli := &CLI{Ctr: f.Server()}

	return cli, nil
}

// Validate runs the Flipt validation command on the given directory
func (c *CLI) Validate(ctx context.Context, dir *Directory) (*Container, error) {
	// TODO: this doesn't work until we can run validate against a directory, as we get permission denied errors because it tries to scan the entire filesystem
	return c.Ctr.WithMountedDirectory("/tmp", dir).WithExec([]string{"./flipt", "validate"}).Sync(ctx)
}

// Import runs the Flipt import command on the given file
func (c *CLI) Import(ctx context.Context,
	// The file to import
	//
	// +optional
	file *File) (*Container, error) {
	if file == nil && c.addr == "" {
		return nil, fmt.Errorf("either file or addr must be provided")
	}

	if c.addr != "" {
		return c.Ctr.WithExec([]string{"./flipt", "import", "--addr", c.addr}).Sync(ctx)
	}

	ctr, err := c.Ctr.WithFile("/tmp/features.yml", file).Sync(ctx)
	if err != nil {
		return nil, err
	}

	return ctr.WithExec([]string{"./flipt", "import", "/tmp/features.yml"}).Sync(ctx)
}

// Export runs the Flipt export command and returns the exported file
func (c *CLI) Export(ctx context.Context) (*File, error) {
	var (
		ctr *Container
		err error
	)

	if c.addr != "" {
		ctr, err = c.Ctr.WithExec([]string{"./flipt", "export", "--addr", c.addr, "-o", "/tmp/out/features.yml"}).Sync(ctx)
	} else {
		ctr, err = c.Ctr.WithExec([]string{"./flipt", "export", "-o", "/tmp/out/features.yml"}).Sync(ctx)
	}

	if err != nil {
		return nil, err
	}

	return ctr.File("/tmp/out/features.yml"), nil
}
