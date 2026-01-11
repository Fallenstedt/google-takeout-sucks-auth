package handlers

import (
	_ "embed"
)

//go:embed home-content.md
var homeContent []byte

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
		<div class="container">
			<div>{{.Content}}</div>
		</div>
	</body>
</html>`
