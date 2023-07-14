package kana

import (
	"fmt"

	"github.com/yihanzhen/konjac/pkg/errors"
)

func NotHiraganaErrorf(r rune) error {
	return fmt.Errorf("Not a hiragana: %v: %w", r, errors.IllegalArgError)
}

var hiraganaTable [][]rune = [][]rune{
	{'あ', 'い', 'う', 'え', 'お'},
	{'か', 'き', 'く', 'け', 'こ'},
	{'が', 'ぎ', 'ぐ', 'げ', 'ご'},
	{'さ', 'し', 'す', 'せ', 'そ'},
	{'ざ', 'じ', 'ず', 'ぜ', 'ぞ'},
	{'た', 'ち', 'つ', 'て', 'と'},
	{'だ', 'ぢ', 'づ', 'で', 'ど'},
	{'な', 'に', 'ぬ', 'ね', 'の'},
	{'は', 'ひ', 'ふ', 'へ', 'ほ'},
	{'ば', 'び', 'ぶ', 'べ', 'ぼ'},
	{'ぱ', 'ぴ', 'ぷ', 'ぺ', 'ぽ'},
	{'ま', 'み', 'む', 'め', 'も'},
	{'や', ' ', 'ゆ', ' ', 'よ'},
	{'ら', 'り', 'る', 'れ', 'ろ'},
	{'わ', ' ', ' ', ' ', 'を'},
	{'ん', 'っ'},
	{'ゃ', ' ', 'ゅ', ' ', 'ょ'},
}

var hiraganaMap map[rune][]int = map[rune][]int{
	'あ': {0, 0}, 'い': {0, 1}, 'う': {0, 2}, 'え': {0, 3}, 'お': {0, 4},
	'か': {1, 0}, 'き': {1, 1}, 'く': {1, 2}, 'け': {1, 3}, 'こ': {1, 4},
	'が': {2, 0}, 'ぎ': {2, 1}, 'ぐ': {2, 2}, 'げ': {2, 3}, 'ご': {2, 4},
	'さ': {3, 0}, 'し': {3, 1}, 'す': {3, 2}, 'せ': {3, 3}, 'そ': {3, 4},
	'ざ': {4, 0}, 'じ': {4, 1}, 'ず': {4, 2}, 'ぜ': {4, 3}, 'ぞ': {4, 4},
	'た': {5, 0}, 'ち': {5, 1}, 'つ': {5, 2}, 'て': {5, 3}, 'と': {5, 4},
	'だ': {6, 0}, 'ぢ': {6, 1}, 'づ': {6, 2}, 'で': {6, 3}, 'ど': {6, 4},
	'な': {7, 0}, 'に': {7, 1}, 'ぬ': {7, 2}, 'ね': {7, 3}, 'の': {7, 4},
	'は': {8, 0}, 'ひ': {8, 1}, 'ふ': {8, 2}, 'へ': {8, 3}, 'ほ': {8, 4},
	'ば': {9, 0}, 'び': {9, 1}, 'ぶ': {9, 2}, 'べ': {9, 3}, 'ぼ': {9, 4},
	'ぱ': {10, 0}, 'ぴ': {10, 1}, 'ぷ': {10, 2}, 'ぺ': {10, 3}, 'ぽ': {10, 4},
	'ま': {11, 0}, 'み': {11, 1}, 'む': {11, 2}, 'め': {11, 3}, 'も': {11, 4},
	'や': {12, 0}, 'ゆ': {12, 2}, 'よ': {12, 4},
	'ら': {13, 0}, 'り': {13, 1}, 'る': {13, 2}, 'れ': {13, 3}, 'ろ': {13, 4},
	'わ': {14, 0}, 'を': {14, 4},
	'ん': {15, 0}, 'っ': {15, 1},
	'ゃ': {16, 0}, 'ゅ': {16, 0}, 'ょ': {16, 0},
}

func IsHiragana(r rune) bool {
	_, ok := hiraganaMap[r]
	return ok
}

func IsHiraganaString(str string) bool {
	for _, r := range str {
		if !IsHiragana(r) {
			return false
		}
	}
	return true
}

func IsCol(r rune, col int) bool {
	pos, ok := hiraganaMap[r]
	if !ok {
		return false
	}
	return pos[1] == col
}

func ToCol(r rune, col int) (rune, error) {
	pos, ok := hiraganaMap[r]
	if !ok {
		return ' ', fmt.Errorf("invalid rune %v, not in HiraganaMap", r)
	}

	if pos[0] < 0 || pos[0] >= len(hiraganaTable) {
		panic(fmt.Sprintf("internal: invalid index %v for HiraganaTable", pos))
	}

	if pos[1] < 0 || pos[1] >= len(hiraganaTable[0]) {
		panic(fmt.Sprintf("internal: invalid index %v for HiraganaTable", pos))
	}
	if col < 0 || col >= len(hiraganaTable[0]) {
		return ' ', fmt.Errorf("Col: invalid col, got %v, want >= 0 and < 5: %w", col, errors.IllegalArgError)
	}

	rt := hiraganaTable[pos[0]][col]
	if rt == ' ' {
		return ' ', fmt.Errorf("invalid rune lookup: %v in HiraganaTable is empty: %w", pos, errors.IllegalArgError)
	}

	return rt, nil
}

func LastRuneToCol(word string, col int) (string, error) {
	wordRunes := []rune(word)
	if len(wordRunes) == 0 {
		return "", fmt.Errorf("LastRuneToCol: input word is empty: %w", errors.IllegalArgError)
	}
	updatedRune, err := ToCol(wordRunes[len(wordRunes)-1], col)
	if err != nil {
		return "", fmt.Errorf("LastRuneToCol: %w", err)
	}

	wordRunes = append(wordRunes[0:len(wordRunes)-1], updatedRune)
	return string(wordRunes), nil
}

func TrimLastRune(word string) (string, error) {
	wordRunes := []rune(word)
	if len(wordRunes) == 0 {
		return "", fmt.Errorf("LastRuneToCol: input word is empty: %w", errors.IllegalArgError)
	}
	return string(wordRunes[0 : len(wordRunes)-1]), nil
}
