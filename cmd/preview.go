package cmd

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/cobra"

	"github.com/miladmahmoodi/forge/internal/generator"
)

func jsonEscape(s string) string {
	b, _ := json.Marshal(s)
	return string(b)
}

var (
	previewConfig string
	previewPort   string
)

var previewCmd = &cobra.Command{
	Use:   "preview",
	Short: "Serve a live preview at localhost:4000",
	Long: `Launches a local HTTP server that renders your profile in real time.
Changes to config.yml are detected automatically and the preview updates.`,
	Example: `  forge preview
  forge preview --port 8080`,
	RunE: func(cmd *cobra.Command, args []string) error {
		tmpFile := filepath.Join(os.TempDir(), "forge-preview.md")

		if err := runBuild(previewConfig, tmpFile); err != nil {
			return err
		}

		watcher, err := fsnotify.NewWatcher()
		if err != nil {
			return fmt.Errorf("file watcher: %w", err)
		}
		defer watcher.Close()

		if err := watcher.Add(previewConfig); err != nil {
			return fmt.Errorf("watching %s: %w", previewConfig, err)
		}

		go func() {
			debounce := time.NewTimer(0)
			<-debounce.C
			for {
				select {
				case _, ok := <-watcher.Events:
					if !ok {
						return
					}
					debounce.Reset(300 * time.Millisecond)
				case <-debounce.C:
					if err := runBuild(previewConfig, tmpFile); err != nil {
						fmt.Fprintf(os.Stderr, "  [rebuild error] %v\n", err)
					} else {
						fmt.Println("  rebuilt")
					}
				case err, ok := <-watcher.Errors:
					if !ok {
						return
					}
					fmt.Fprintf(os.Stderr, "  [watcher] %v\n", err)
				}
			}
		}()

		addr := ":" + previewPort
		fmt.Printf("  forge preview\n")
		fmt.Printf("  serving http://localhost:%s\n", previewPort)
		fmt.Printf("  watching %s\n", previewConfig)
		fmt.Println("  press Ctrl+C to stop")

		http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			content, err := os.ReadFile(tmpFile)
			if err != nil {
				http.Error(w, "build failed", 500)
				return
			}
			w.Header().Set("Content-Type", "text/html; charset=utf-8")
			fmt.Fprintf(w, previewHTML, jsonEscape(string(content)))
		})

		// /raw returns the current markdown for the auto-reload poller.
		http.HandleFunc("/raw", func(w http.ResponseWriter, r *http.Request) {
			content, _ := os.ReadFile(tmpFile)
			w.Header().Set("Content-Type", "text/plain; charset=utf-8")
			w.Write(content)
		})

		// Serve theme assets (SVGs, etc.) relative to CWD.
		http.Handle("/themes/", http.StripPrefix("/themes/", http.FileServer(http.Dir("themes"))))

		return http.ListenAndServe(addr, nil)
	},
}

func runBuild(configPath, outputPath string) error {
	_, err := generator.Generate(generator.Options{
		ConfigPath: configPath,
		OutputPath: outputPath,
		ThemesFS:   ThemesFS,
	})
	return err
}

const previewHTML = `<!DOCTYPE html>
<html lang="en">
<head>
  <meta charset="UTF-8">
  <meta name="viewport" content="width=device-width, initial-scale=1.0">
  <title>Forge Preview</title>
  <link rel="stylesheet"
    href="https://cdnjs.cloudflare.com/ajax/libs/github-markdown-css/5.6.1/github-markdown-dark.min.css">
  <style>
    * { box-sizing: border-box; }
    html, body { margin: 0; padding: 0; background: #0d1117; }
    .markdown-body {
      max-width: 980px;
      margin: 0 auto;
      padding: 40px 24px;
    }
    .forge-bar {
      position: fixed; top: 0; left: 0; right: 0; z-index: 100;
      background: #161b22; border-bottom: 1px solid #30363d;
      padding: 8px 20px; display: flex; align-items: center; gap: 12px;
      font-family: "SFMono-Regular", Consolas, monospace; font-size: 12px;
      color: #8b949e;
    }
    .forge-bar .dot { width: 8px; height: 8px; border-radius: 50%; background: #3fb950; }
    .forge-bar .live { color: #3fb950; }
    body { padding-top: 37px; }
  </style>
</head>
<body>
  <div class="forge-bar">
    <span class="dot"></span>
    <span>⚡ forge preview</span>
    <span class="live">● live</span>
    <span style="margin-left:auto; color:#484f58">auto-reloads on config.yml changes</span>
  </div>
  <article class="markdown-body" id="content">Loading...</article>

  <script src="https://cdnjs.cloudflare.com/ajax/libs/marked/12.0.0/marked.min.js"></script>
  <script>
    const raw = %s;
    document.getElementById('content').innerHTML = marked.parse(raw);

    // Auto-reload by polling for changes every 1.5s.
    let lastLen = raw.length;
    setInterval(async () => {
      try {
        const r = await fetch('/raw');
        const text = await r.text();
        if (text.length !== lastLen) {
          lastLen = text.length;
          document.getElementById('content').innerHTML = marked.parse(text);
        }
      } catch(_) {}
    }, 1500);
  </script>
</body>
</html>`

func init() {
	previewCmd.Flags().StringVarP(&previewConfig, "config", "c", "config.yml", "path to config.yml")
	previewCmd.Flags().StringVarP(&previewPort, "port", "p", "4000", "local port to serve on")
	rootCmd.AddCommand(previewCmd)
}
