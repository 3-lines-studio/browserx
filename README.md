# browserx

<p align="center"><img src=".github/ax.svg" width="96" height="96" alt="AX ecosystem"></p>

Read-only visible Chromium tool for AX.

## Configure

Start Chromium with a dedicated profile and CDP port:

```sh
chromium-browser --remote-debugging-address=127.0.0.1 --remote-debugging-port=9222 --user-data-dir="$HOME/snap/chromium/common/browser-profile"
```

Set the artifact directory and register Browserx:

```sh
export BROWSERX_ARTIFACT_DIR="$HOME/.local/share/browserx/artifacts"
export AX_TOOLS=browserx
```

Override the CDP endpoint with `BROWSERX_CDP_URL`.

## Protocol

```sh
browserx describe
printf '{"url":"https://example.com"}' | browserx run browser_open
printf '{}' | browserx run browser_read
printf '{}' | browserx run browser_links
printf '{"name":"example"}' | browserx run browser_screenshot
```

Browserx only opens HTTP and HTTPS URLs, reads the current page, lists links, and saves screenshots.

## Test

```sh
go test ./...
```
