# Exercise 16: Twitter Contest CLI — plan

## Main idea

Giveaway contests on Twitter work like this: "retweet this tweet to enter". The organiser needs two things:

1. **A list of everyone who retweeted.** The API only returns the ~100 most recent retweets, so on a popular tweet you can't wait for the contest to end and ask once. You must poll regularly and **accumulate** users in your own file.
2. **A fair random winner** picked from that accumulated list.

So the program is a small pipeline:

```
 Twitter API ──► fetch retweeters ──► merge with file ──► save file ──► (optional) pick N random winners
   (auth)         (last ≤100)         (set union)        (persist)
```

What you practise in Go: HTTP client with auth, JSON decoding, file I/O, sets (`map[string]struct{}`), `flag`, random selection, interfaces for testing, and (bonus) `time.Ticker` + `context`.

## Heads-up before you start: API access

The exercise was written for Twitter API v1.1 and app-only auth. Twitter is now X and the API has changed a lot:

- Free access may not include "who retweeted" at all; it may need a paid tier. **I haven't verified the current tiers, so check the developer portal before step 1.**
- The v1.1 endpoint `GET statuses/retweets/:id` may be retired or restricted. The modern equivalent is v2 `GET /2/tweets/:id/retweeted_by`. It returns users directly and supports pagination, so it is better than the 100-limit of v1.1.

To keep the exercise doable either way, we **hide the API behind an interface** and test against a fake server (`httptest`). Everything except the final "talk to the real X" step works without paying. If real access works for you, great, you plug it in at the end.

## Layout

```
16_twitter/
  go.mod                 module 16_twitter
  PLAN.md
  cmd/main.go            flags, wiring, run loop
  twitter/               API client package (bonus: reusable)
    client.go            Client, New(), do-request helper
    auth.go              app-only token (bearer) exchange
    retweets.go          Retweeters(ctx, tweetID) ([]string, error)
    client_test.go       tests against httptest.Server
  store/                 persistence of the users file
    store.go             Load(path), Save(path, users)
    store_test.go
  contest/               pure logic, no I/O
    contest.go           Merge(old, new), Pick(users, n)
    contest_test.go
```

Why `cmd/main.go`: same as exercise 14, `go build ./...` leaves no binary in the root. Run with `go run ./cmd`.

Secrets: credentials come from env vars (`TWITTER_KEY`, `TWITTER_SECRET`), never flags and never committed. The repo `.gitignore` already ignores dotfiles, so a `.env` file is safe, but the users file (`retweeters.json`) is not ignored. That's fine, it contains only public handles.

## Steps

Each step ends with a checkpoint. Try it yourself first; ask when stuck.

### Step 0. Project skeleton
- `go mod init 16_twitter`, create the folders above.
- ✅ `go build ./...` passes with an empty `main`.

### Step 1. Pure logic first: merge and pick (`contest`)
No network, easiest to get right, and the rest depends on it.
- `Merge(old, fresh []string) []string`: union, no duplicates, **normalised** (Twitter handles are case-insensitive, so lowercase them and trim a leading `@`), sorted so the file is stable.
- `Pick(users []string, n int) ([]string, error)`: n distinct random users. Idea: shuffle a copy (`math/rand/v2`, `rand.Shuffle`) and take the first n. Decide what happens when `n > len(users)` (error is cleanest) and when `n <= 0`.
- ✅ Table tests: duplicates, case differences, empty input, n larger than the list, picked winners are unique and come from the list.

### Step 2. Persistence (`store`)
- `Load(path) ([]string, error)`: a missing file is **not** an error (first run returns an empty list); a corrupt file is.
- `Save(path, users) error`: JSON array of strings is the easy choice. Safer write: write to `path + ".tmp"`, then `os.Rename`, so a crash mid-write can't destroy the accumulated list. That list is the one thing you can't re-fetch.
- ✅ Tests with `t.TempDir()`: round trip, missing file, corrupt file.

### Step 3. Authenticate (`twitter/auth.go`)
Application-only auth, v1.1 style:
- `POST https://api.twitter.com/oauth2/token`
- Header `Authorization: Basic base64(url-encoded key ":" url-encoded secret)`. In Go, `req.SetBasicAuth(key, secret)` does the base64 part.
- Header `Content-Type: application/x-www-form-urlencoded;charset=UTF-8`, body `grant_type=client_credentials`.
- Response JSON has `token_type: "bearer"` and `access_token`. Check the type, keep the token.
- All later calls send `Authorization: Bearer <token>`.
- For v2 the portal gives you a Bearer Token directly, so this step is skipped; keep the code anyway if you follow the exercise literally. Decide which you target and keep the base URL configurable (needed for the fake server).
- ✅ Test against `httptest.NewServer` that checks the headers and body it receives and returns a token.

### Step 4. Fetch retweeters (`twitter/retweets.go`)
- Define the small interface where it's *used* (in `cmd` or `contest`): `type Fetcher interface { Retweeters(ctx context.Context, tweetID string) ([]string, error) }`.
- v1.1: `GET /1.1/statuses/retweets/{id}.json?count=100`, a JSON array of tweets; the handle is `user.screen_name`. Decode only the fields you need.
- v2: `GET /2/tweets/{id}/retweeted_by?max_results=100`, a JSON object with `data[].username`; follow `meta.next_token` via `pagination_token` until it's empty.
- Handle non-200 properly: read the error body into the error message; on `429` read the `x-rate-limit-reset` header (unix seconds) and return a descriptive error.
- Always use `http.NewRequestWithContext` and a client with a timeout.
- ✅ Tests against a fake server: happy path, empty list, 401, 429, malformed JSON.

### Step 5. Wire it up (`cmd/main.go`)
Flags:
- `-tweet` (required): tweet ID.
- `-file` (default `retweeters.json`): the users file.
- `-pick N` (default 0 = don't pick): how many winners.

Flow: load file → fetch → merge → save → print how many new users were added → if `-pick > 0`, pick and print winners.
Decision for you: update first, then pick (what the video does), or pick from the file only. Updating first is better; make it fail loudly if the fetch fails, but still allow `-pick` from the existing file (maybe via a `-offline` flag).
- ✅ Manual run against your fake server (small `main` or test helper) and, if you have access, the real API.

### Step 6. Polish
- Exit codes: non-zero on error, messages to stderr.
- Run `go vet ./...` and `go test -race ./...`.

## Bonus A: reusable API client package

Make `twitter` a proper client library:
- `twitter.New(key, secret, opts ...Option)` with options `WithHTTPClient`, `WithBaseURL`.
- Token fetched lazily and cached, protected by a mutex (several goroutines may call it); refetch once on 401.
- A shared `do(ctx, method, path, query, out any) error` helper doing auth, status checks, JSON decode and rate-limit errors, so each endpoint is ~10 lines.
- Add 2 or 3 more endpoints, e.g. `Tweet(ctx, id)` (single tweet), `UserByUsername(ctx, name)`, `Likers(ctx, id)`.
- A typed `RateLimitError{ResetAt time.Time}` the caller can inspect with `errors.As`.

## Bonus B: poll automatically, pick at the end

Run once, walk away.
- New flags: `-every 5m` (poll interval) and `-for 24h` (total duration). With `-pick N` the winners are picked when the time is up.
- Use `context.WithTimeout` for the total time and `signal.NotifyContext(ctx, os.Interrupt)` so Ctrl+C stops cleanly.
- Loop with `time.NewTicker(every)`, `select` on `ctx.Done()` and `ticker.C`. Do one fetch immediately at start (a ticker waits a full interval for its first tick). Always `defer ticker.Stop()`.
- A failed poll shouldn't kill the run: log it and try on the next tick. On `RateLimitError`, sleep until `ResetAt` (or skip that tick).
- On exit (timeout *or* Ctrl+C): one last fetch, save, then pick if `-pick` was given. Think about whether Ctrl+C should pick winners; the safer default is no, only on timeout.
- ✅ Test the loop with `-every 1s -for 5s` against the fake server.

## Pitfalls to remember
- Tweet IDs are bigger than JS-safe integers; in Go keep them as **strings**.
- Never log the token or secret.
- Close response bodies (`defer resp.Body.Close()`) and check `StatusCode` before decoding.
- Don't write the users file from several goroutines at once; the polling loop is sequential, keep it that way.
- Random: use `math/rand/v2`; no manual seeding needed in Go 1.20+.

## Definition of done
- [ ] `go vet ./...` and `go test -race ./...` clean
- [ ] Users file accumulates across runs and never loses names
- [ ] `-pick N` returns N unique winners
- [ ] Bonus A: client is a separate package with extra endpoints
- [ ] Bonus B: polling + automatic winner works with Ctrl+C handling
