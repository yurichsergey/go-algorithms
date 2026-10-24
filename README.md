# go-algorithms

Go practice repo with three parts:

| Path | What it is |
|------|------------|
| [`go_neetcode/`](go_neetcode) | NeetCode-style algorithm problems solved in Go, each with tests |
| [`go_awesome/gophercies/`](go_awesome/gophercies) | Solutions to the [Gophercises](https://gophercises.com) exercises |
| [`site/`](site) | React + Vite site with interactive visualizers for some of the problems |

Root module: `github.com/yurichsergey/go-algorithms` (Go 1.25).

## Algorithm problems (`go_neetcode`)

Files are named `NN_problem_name.go` with a matching `NN_problem_name_test.go`.

| Package dir | Problems |
|-------------|----------|
| `01_arrays_hashing` | Contains Duplicate, Valid Anagram, Two Sum, Group Anagrams, Top K Frequent Elements, Encode and Decode Strings, Product of Array Except Self, Valid Sudoku, Longest Consecutive Sequence |
| `02_two_pointers` | Valid Palindrome, Two Sum II, 3Sum, Container With Most Water |

Run the tests:

```sh
go test ./go_neetcode/...
```

## Gophercises (`go_awesome/gophercies`)

| # | Exercise |
|---|----------|
| 01 | Quiz game |
| 02 | URL shortener |
| 03 | Choose Your Own Adventure |
| 04 | HTML link parser |
| 05 | Sitemap builder |
| 06 | HackerRank problems (Caesar cipher, camelCase) |
| 07 | CLI task manager |
| 08 | Phone number normalizer |
| 09 | Deck of cards |
| 10 | Blackjack |
| 11 | Blackjack AI |
| 12 | File renaming |
| 13 | Quiet HN |
| 14 | Panic recovery middleware |

Most exercises have their own `go.mod`, so run them from inside their directory:

```sh
cd go_awesome/gophercies/14_recover
go run ./cmd
```

Exercises 07 and 08 keep local databases (`tasks.db`, `phones.db`); `*.db` files are git-ignored.

## Visualizer site (`site`)

React 19 + Vite app (react-router, d3, framer-motion) that lists the problems by category and, for some of them, shows a step-by-step visualizer. Problems are registered in `site/src/problems/registry.js`; a problem with `Visualizer: null` has no visualizer yet.

Currently available visualizers: Two Sum, Container With Most Water.

```sh
cd site
npm install
npm run dev        # http://localhost:5173
npm run build      # output in site/dist
npm run lint       # oxlint
```

Or with Docker: `docker compose up` from `site/`.

### Adding a visualizer

1. Create a component under `site/src/problems/<category>/`.
2. Import it in `registry.js` and set it as the problem's `Visualizer`.

### Deployment

Pushes to `main` that touch `site/**` trigger `.github/workflows/deploy-site.yml`, which builds the site and publishes it to GitHub Pages (base path `/go-algorithms/`).
