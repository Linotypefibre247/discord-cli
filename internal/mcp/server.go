package mcp

import (
	"fmt"

	"github.com/mark3labs/mcp-go/server"
)

// Serve starts the MCP server over stdio.
func Serve(storeDir string) error {
	s := server.NewMCPServer(
		"discocli",
		"1.0.0",
		server.WithToolCapabilities(false),
		server.WithRecovery(),
	)

	RegisterTools(s, storeDir)

	if err := server.ServeStdio(s); err != nil {
		return fmt.Errorf("MCP server error: %w", err)
	}
	return nil
}
