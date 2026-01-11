package handlers

import (
	"bytes"
	"crypto/sha256"
	_ "embed"
	"fmt"
	"html/template"

	"github.com/Fallenstedt/google-takeout-sucks-auth/internal/logging"
	"github.com/yuin/goldmark"
)

var (
	//go:embed content/home-content.md
	homeContent []byte
	//go:embed content/privacy-policy.md
	privacyPolicyContent []byte
	//go:embed content/terms-of-service.md
	termsOfServiceContent []byte

	homeETag           string
	privacyETag        string
	termsOfServiceETag string
	pageTmpl           *template.Template
	homeHTML           template.HTML
	privacyHTML        template.HTML
	termsOfServiceHTML template.HTML
)

const page = `<!doctype html>
<html lang="en">
	<head>
		<meta charset="utf-8">
		<meta name="viewport" content="width=device-width,initial-scale=1">
		<title>Google Takeout Sucks</title>
		<style>
			body { font-family: -apple-system, BlinkMacSystemFont, "Segoe UI", Roboto, "Helvetica Neue", Arial; margin: 2rem; color: #111 }
			.container { max-width: 900px; margin: 0 auto; }
			pre { white-space: pre-wrap; background: #f7f7f8; padding: 1rem; border-radius: 6px; overflow: auto; }
			h1,h2,h3 { color: #0b5fff }
			img { max-width: 100%; height: auto }
		</style>
	</head>
	<body>
		<main class="container">
			<div>{{.Content}}</div>
		</main>
		<footer class="container">
			<a href="/">Home</a>
			<a href="/privacy-policy">Privacy Policy</a>
			<a href="/terms-of-service">Terms of Service</a>
		</footer>
	</body>
</html>`

func init() {
	// compute ETags for embedded content
	var sum = sha256.Sum256(homeContent)
	homeETag = fmt.Sprintf(`"%x"`, sum[:])

	sum = sha256.Sum256(privacyPolicyContent)
	privacyETag = fmt.Sprintf(`"%x"`, sum[:])

	sum = sha256.Sum256(termsOfServiceContent)
	privacyETag = fmt.Sprintf(`"%x"`, sum[:])

	// parse page template once
	pageTmpl = template.Must(template.New("page").Parse(page))

	// Pre-render embedded markdown to HTML to avoid per-request conversion
	md := goldmark.New()
	var buf bytes.Buffer
	if err := md.Convert(homeContent, &buf); err != nil {
		logging.ErrorLog.Printf("failed to convert homeContent during init: %v", err)
	} else {
		homeHTML = template.HTML(buf.String())
	}

	buf.Reset()
	if err := md.Convert(privacyPolicyContent, &buf); err != nil {
		logging.ErrorLog.Printf("failed to convert privacyPolicyContent during init: %v", err)
	} else {
		privacyHTML = template.HTML(buf.String())
	}

	buf.Reset()
	if err := md.Convert(termsOfServiceContent, &buf); err != nil {
		logging.ErrorLog.Printf("failed to convert termsOfServiceContent during init: %v", err)
	} else {
		termsOfServiceHTML = template.HTML(buf.String())
	}
}
