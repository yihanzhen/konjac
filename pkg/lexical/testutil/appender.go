package testutil

import (
	"testing"

	"github.com/yihanzhen/konjac/pkg/lexical"
	"github.com/yihanzhen/konjac/pkg/lexical/adjective"
	"github.com/yihanzhen/konjac/pkg/lexical/verb"
)

func AppenderWithVerb(t *testing.T, v string) *lexical.Appender {
	t.Helper()
	vb, err := verb.NewVerb(v)
	if err != nil {
		t.Fatalf("NewVerb(%q) has unexpected error: %v", v, err)
	}
	app := lexical.NewAppender()
	if err := app.Append(vb); err != nil {
		t.Fatalf("Append(%v) has unexpected error: %v", vb, err)
	}
	return app
}

func AppenderWithAdjective(t *testing.T, a string) *lexical.Appender {
	t.Helper()
	adj, err := adjective.NewAdjective(a)
	if err != nil {
		t.Fatalf("NewVerb(%q) has unexpected error: %v", a, err)
	}
	app := lexical.NewAppender()
	if err := app.Append(adj); err != nil {
		t.Fatalf("Append(%v) has unexpected error: %v", adj, err)
	}
	return app
}
