package blackjack

type Move int

const (
	Hit Move = iota
	Stand
)

type AI interface {
	Bet(balance int) int
	Play(hand Hand, dealer Card) Move
	Result(hand Hand, result Result)
}
