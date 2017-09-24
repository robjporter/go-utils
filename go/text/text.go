package text

import (
	"crypto/md5"
	rand2 "crypto/rand"
	"encoding/hex"
	"fmt"
	"io"
	"math"
	"math/rand"
	"strings"
	"time"
	"unicode"

	"github.com/robjporter/go-utils/go/as"
)

const UPPERCASE = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
const LOWERCASE = "abcdefghijklmnopqrstuvwxyz"
const VOWELS = "aeoui"
const CONSONANTS = "bcdfghjklmnpqrstvwxyz"
const NUMBERS = "1234567890"
const SPECIALS = "!@#$%^&*-_"
const MAXPREFIXLENGTH = 4

var noop = func(a rune) rune { return a }
var sourceCharacters = map[rune]rune{
	'à': 'a', 'á': 'a', 'â': 'a', 'ã': 'a', 'ă': 'a', 'ạ': 'a', 'ả': 'a', 'ấ': 'a', 'ầ': 'a', 'ẩ': 'a', 'ẫ': 'a', 'ậ': 'a', 'ắ': 'a', 'ằ': 'a', 'ẳ': 'a', 'ẵ': 'a', 'ặ': 'a',
	'À': 'A', 'Á': 'A', 'Â': 'A', 'Ã': 'A', 'Ạ': 'A', 'Ầ': 'A', 'Ả': 'A', 'Ấ': 'A', 'Ẩ': 'A', 'Ẫ': 'A', 'Ậ': 'A', 'Ắ': 'A', 'Ằ': 'A', 'Ẳ': 'A', 'Ẵ': 'A', 'Ặ': 'A',
	'đ': 'd', 'Đ': 'D',
	'è': 'e', 'é': 'e', 'ê': 'e', 'ẹ': 'e', 'ẻ': 'e', 'ẽ': 'e', 'ế': 'e', 'ề': 'e', 'ể': 'e', 'ễ': 'e', 'ệ': 'e',
	'È': 'E', 'É': 'E', 'Ê': 'E', 'Ẹ': 'E', 'Ẻ': 'E', 'Ẽ': 'E', 'Ế': 'E', 'Ề': 'E', 'Ể': 'E', 'Ễ': 'E', 'Ệ': 'E',
	'ì': 'i', 'í': 'i', 'ĩ': 'i', 'ỉ': 'i', 'ị': 'i', 'î': 'i',
	'Ì': 'I', 'Í': 'I', 'Ĩ': 'I', 'Ỉ': 'I', 'Ị': 'I',
	'ò': 'o', 'ó': 'o', 'ô': 'o', 'õ': 'o', 'ơ': 'o', 'ọ': 'o', 'ỏ': 'o', 'ố': 'o', 'ồ': 'o', 'ổ': 'o', 'ỗ': 'o', 'ộ': 'o', 'ớ': 'o', 'ờ': 'o', 'ở': 'o', 'ỡ': 'o', 'ợ': 'o',
	'Ò': 'O', 'Ó': 'O', 'Ô': 'O', 'Õ': 'O', 'Ơ': 'O', 'Ọ': 'O', 'Ỏ': 'O', 'Ố': 'O', 'Ồ': 'O', 'Ổ': 'O', 'Ỗ': 'O', 'Ộ': 'O', 'Ớ': 'O', 'Ờ': 'O', 'Ở': 'O', 'Ỡ': 'O', 'Ợ': 'O',
	'ù': 'u', 'ú': 'u', 'ũ': 'u', 'ư': 'u', 'ụ': 'u', 'ủ': 'u', 'ứ': 'u', 'ừ': 'u', 'ử': 'u', 'ữ': 'u',
	'Ù': 'U', 'Ú': 'U', 'Ũ': 'U', 'Ư': 'U', 'Ự': 'U', 'Ủ': 'U', 'Ứ': 'U', 'Ừ': 'U', 'Ử': 'U', 'Ữ': 'U',
	'ý': 'y', 'ự': 'y', 'Ý': 'Y', 'Ụ': 'Y',
}

func HumanizeString(str string) string {
	var human []rune
	for i, l := range str {
		if i > 0 && isUppercase(byte(l)) {
			if (!isUppercase(str[i-1]) && str[i-1] != ' ') || (i+1 < len(str) && !isUppercase(str[i+1]) && str[i+1] != ' ' && str[i-1] != ' ') {
				human = append(human, rune(' '))
			}
		}
		human = append(human, l)
	}
	return strings.Title(string(human))
}

func isUppercase(char byte) bool {
	return 'A' <= char && char <= 'Z'
}

func WordCount(str string) int {
	return strings.Count(str, " ") + 1
}

func WordFrequency(str string) map[string]int {
	splits := strings.Split(str, " ")
	tmp := make(map[string]int)
	for i := 0; i < len(splits); i++ {
		tmp[splits[i]] += 1
	}
	return tmp
}

func LongestWord(str string) string {
	return ""
}

func LongestWordLength(str string) int {
	return 0
}

func AverageWordLength(str string) int {
	return 0
}

func MD5(str string) string {
	md := md5.New()
	md.Write([]byte(str))
	return hex.EncodeToString(md.Sum(nil))
}

func ToTrain(s string) string {
	return snaker(s, '-', unicode.ToUpper, unicode.ToUpper, noop)
}

func ToSpinal(s string) string {
	return snaker(s, '-', unicode.ToLower, unicode.ToLower, unicode.ToLower)
}

func ToSnake(s string) string {
	return snaker(s, '_', unicode.ToLower, unicode.ToLower, unicode.ToLower)
}

func ToSnakeUpper(s string) string {
	return snaker(s, '_', unicode.ToUpper, unicode.ToUpper, unicode.ToUpper)
}

func ToCamel(s string) string {
	return snaker(s, rune(0), unicode.ToUpper, unicode.ToUpper, noop)
}

func ToCamelLower(s string) string {
	return snaker(s, rune(0), unicode.ToLower, unicode.ToUpper, noop)
}

func GenerateRandomString(length int) string {
	return generateRandom(length, false)
}

func GenerateRandomStringSpecial(length int) string {
	return generateRandom(length, true)
}

func generateRandomContains(str string, num int) bool {
	l := strings.ContainsAny(LOWERCASE, str)
	u := strings.ContainsAny(UPPERCASE, str)
	n := strings.ContainsAny(NUMBERS, str)
	s := strings.ContainsAny(SPECIALS, str)

	if num == 3 {
		return l && u && n
	} else if num == 4 {
		return l && u && n && s
	}
	return false
}

func generateRandom(length int, specials bool) string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	str := make([]string, length)
	choiceset := ""
	if specials {
		choiceset = LOWERCASE + UPPERCASE + NUMBERS + SPECIALS
	} else {
		choiceset = LOWERCASE + UPPERCASE + NUMBERS
	}

	for i := 0; i < length; i++ {
		index := r.Intn(len(choiceset))
		str[i] = choiceset[index : index+1]
	}

	tmpStr := strings.Join(str, "")

	num := 3
	if specials {
		num = 4
	}

	if generateRandomContains(tmpStr, num) {
		return tmpStr
	} else {
		generateRandom(length, specials)
	}
	return ""
}

func ordinise(number int) string {
	switch int(math.Abs(float64(number))) % 100 {
	case 11, 12, 13:
		return "th"
	default:
		switch int(math.Abs(float64(number))) % 10 {
		case 1:
			return "st"
		case 2:
			return "nd"
		case 3:
			return "rd"
		}
	}
	return "th"
}

func ToOrdinise(number int) string {
	return as.ToString(number) + ordinise(number)
}

func snaker(s string, wordSeparator rune, firstRune func(rune) rune, firstRuneOfWord func(rune) rune, otherRunes func(rune) rune) string {
	useWordSeperator := wordSeparator != rune(0)
	newS := []rune{}

	// pops a rune off newS
	lastRuneIsWordSeparator := func() bool {
		if len(newS) > 0 {
			return newS[len(newS)-1] == wordSeparator
		}
		return true
	}

	prev := wordSeparator
	for _, cur := range s {
		isWordBoundary := (unicode.IsLower(prev) && unicode.IsUpper(cur)) || unicode.IsSpace(prev)

		if !unicode.IsLetter(cur) {
			// ignore
		} else if isWordBoundary {
			if useWordSeperator && !lastRuneIsWordSeparator() {
				newS = append(newS, wordSeparator)
			}
			newS = append(newS, firstRuneOfWord(cur))
		} else {
			newS = append(newS, otherRunes(cur))
		}

		prev = cur
	}

	if len(newS) > 0 {
		newS[0] = firstRune(newS[0])
	}

	return string(newS)
}

func CleanTextForeignCharacters(str string) string {
	arr := []rune(str)
	for i, r := range arr {
		if v, ok := sourceCharacters[r]; ok {
			arr[i] = v
		}
	}
	return string(arr)
}

// UUID4 generates a random UUID according to RFC 4122
func UUID4() (string, bool) {
	uuid := make([]byte, 16)
	n, err := io.ReadFull(rand2.Reader, uuid)
	if n != len(uuid) || err != nil {
		return "", false
	}
	// variant bits; see section 4.1.1
	uuid[8] = uuid[8]&^0xc0 | 0x80
	// version 4 (pseudo-random); see section 4.1.3
	uuid[6] = uuid[6]&^0xf0 | 0x40
	return fmt.Sprintf("%x-%x-%x-%x-%x", uuid[0:4], uuid[4:6], uuid[6:8], uuid[8:10], uuid[10:]), true
}

func Pluralise(s string) string {
	if strings.HasSuffix(s, "s") {
		return s + "es"
	}
	if (len(s) > 1) && strings.HasSuffix(s, "y") && !isOneOf(s[(len(s)-2):], "ay", "ey", "oy", "uy", "iy") {
		return s[0:(len(s)-1)] + "ies"
	}
	return s + "s"
}

func isOneOf(s string, all ...string) bool {
	for _, a := range all {
		if s == a {
			return true
		}
	}
	return false
}

func MakeAnnouncement(message string) {
	fmt.Println(Announcement(message))
}

func Announcement(message string) string {
	length := len(message)
	t := strings.Repeat("=", length)
	mess := t + "\n" + message + "\n" + t
	return mess
}
