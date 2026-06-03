// Copyright 2026 joelcodess. Licensed under Apache-2.0. See LICENSE.

package cli

import (
	"bytes"
	"strings"
	"testing"

	"github.com/spf13/cobra"
)

// findCmd walks the command tree following the given path of Use-names
// (first word of each command's Use) and returns the leaf, or nil.
func findCmd(root *cobra.Command, path ...string) *cobra.Command {
	cur := root
	for _, name := range path {
		var next *cobra.Command
		for _, c := range cur.Commands() {
			if strings.Fields(c.Use)[0] == name {
				next = c
				break
			}
		}
		if next == nil {
			return nil
		}
		cur = next
	}
	return cur
}

// The endpoint path/method annotations are load-bearing: a typo here sends
// the request to the wrong URL. Pin them so a future edit can't silently
// drift them.
func TestODevicesCommands_Annotations(t *testing.T) {
	root := newRootCmd(&rootFlags{})

	cases := []struct {
		path       []string
		wantMethod string
		wantPath   string
	}{
		{[]string{"o", "devices", "commands", "run"}, "POST", "/o/{organization_id}/devices/commands/run"},
		{[]string{"o", "devices", "commands", "output"}, "GET", "/o/{organization_id}/devices/{agent_id}/commands/{action_id}/output"},
	}
	for _, tc := range cases {
		c := findCmd(root, tc.path...)
		if c == nil {
			t.Fatalf("command %q not wired into the tree", strings.Join(tc.path, " "))
		}
		if got := c.Annotations["pp:method"]; got != tc.wantMethod {
			t.Errorf("%s: pp:method = %q, want %q", strings.Join(tc.path, " "), got, tc.wantMethod)
		}
		if got := c.Annotations["pp:path"]; got != tc.wantPath {
			t.Errorf("%s: pp:path = %q, want %q", strings.Join(tc.path, " "), got, tc.wantPath)
		}
	}
}

// run requires --agent-ids and --command. The check fires before any client
// or config is constructed, so this is deterministic regardless of the host's
// ~/.config/addigy-cli state.
func TestODevicesCommandsRun_RequiresFlags(t *testing.T) {
	root := newRootCmd(&rootFlags{})
	var out bytes.Buffer
	root.SetOut(&out)
	root.SetErr(&out)
	root.SetArgs([]string{"o", "devices", "commands", "run", "org-123"})

	err := root.Execute()
	if err == nil {
		t.Fatalf("expected an error when required flags are missing, got nil")
	}
	if !strings.Contains(err.Error(), "agent-ids") {
		t.Errorf("error should name the missing agent-ids flag, got: %v", err)
	}
}

// output requires --agent-id and --action-id for the same reason.
func TestODevicesCommandsOutput_RequiresFlags(t *testing.T) {
	root := newRootCmd(&rootFlags{})
	var out bytes.Buffer
	root.SetOut(&out)
	root.SetErr(&out)
	root.SetArgs([]string{"o", "devices", "commands", "output", "org-123"})

	err := root.Execute()
	if err == nil {
		t.Fatalf("expected an error when required flags are missing, got nil")
	}
	if !strings.Contains(err.Error(), "agent-id") {
		t.Errorf("error should name a missing id flag, got: %v", err)
	}
}
