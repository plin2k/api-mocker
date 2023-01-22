package faker

import (
	"bytes"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"fmt"
	"math/rand"
	"strings"
	"text/template"
	"time"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func init() {
	rand.Seed(time.Now().UnixMicro())
}

func Process(data []byte) ([]byte, error) {
	tmpl, err := template.New("faker").Funcs(funcFakerMap).Parse(string(data))
	if err != nil {
		return nil, err
	}

	var buf bytes.Buffer
	err = tmpl.Execute(&buf, nil)
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

var funcFakerMap = template.FuncMap{
	"formatDate": formatDate,
	//
	"randInt":         randInt,
	"randString":      randString,
	"randFloat":       randFloat,
	"randBool":        randBool,
	"randArrayInt":    randArrayInt,
	"randArrayString": randArrayString,
	//
	"randHash":   randHash,
	"randCookie": randCookie,
	"randJWT":    randJWT,
	"randUUID":   randUUID,
	"toLower":    toLower,
	"toUpper":    toUpper,
	"toTitle":    toTitle,
	//
	"randCity":    randCity,
	"randCountry": randCountry,
	//
	//"randPhone":  randPhone,
	"randEmail":    randEmail,
	"randDomain":   randDomain,
	"randLangCode": randLangCode,
	"randDate":     randDate,
	//
	"randFirstName": randFirstName,
	"randLastName":  randLastName,
	"randSex":       randSex,
}

func formatDate(timeStamp time.Time) string {
	return timeStamp.Format("Mon, 02 Jan 2006")

}

func randJWT() string {
	return fmt.Sprintf("%s.%s.%s", randLetters(90, letterRunesAndNumbers), randLetters(128, letterRunesAndNumbers), randLetters(43, letterRunesAndNumbers))
}

func randUUID() string {
	return strings.ToLower(fmt.Sprintf("%s-%s-%s-%s-%s",
		randLetters(8, letterRunesAndNumbers),
		randLetters(4, letterRunesAndNumbers),
		randLetters(4, letterRunesAndNumbers),
		randLetters(4, letterRunesAndNumbers),
		randLetters(12, letterRunesAndNumbers)))
}

func randCookie() string {
	return fmt.Sprintf("id=%s; Expires=%s;", randLetters(5, letterRunesAndNumbers), time.Now().AddDate(0, 0, randInt(0, 365)).Format(time.RFC1123))
}

func randFirstName(sex int) string {
	if sex == 0 {
		return firstMaleNames[rand.Intn(len(firstMaleNames)-1)]
	}

	return firstFemaleNames[rand.Intn(len(firstMaleNames)-1)]
}

func randLastName() string {
	return lastNames[rand.Intn(len(lastNames)-1)]
}

func randLangCode() string {
	return langCodes[rand.Intn(len(langCodes)-1)]
}

func randInt(min, max int) int {
	return rand.Intn(max-min) + min
}

func randBool() bool {
	if randInt(0, 1) == 1 {
		return true
	}

	return false
}

func randSex() string {
	if randInt(0, 1) == 1 {
		return "male"
	}

	return "female"
}

func randFloat(min, max int) float64 {
	return float64(rand.Intn(max-min)+min) / 100
}

func randArrayInt(n int, min, max int) (arr []int) {
	for i := 0; i < n; i++ {
		arr = append(arr, randInt(min, max))
	}

	return arr
}

func randArrayString(n int, nc int) (arr []string) {
	for i := 0; i < n; i++ {
		arr = append(arr, randString(nc))
	}

	return arr
}

func randString(n int) string {
	return randLetters(n, letterRunes)
}

func randCountry() string {
	return countries[rand.Intn(len(countries)-1)]
}

func randCity() string {
	return cities[rand.Intn(len(cities)-1)]
}

func randDomain() string {
	return strings.ToLower(domains[rand.Intn(len(domains)-1)])
}

func randEmail() string {
	var (
		str    string
		choice = randInt(0, 3)
	)

	if choice == 0 {
		n := rand.Intn(1)
		str = fmt.Sprintf("%s.%s", randFirstName(n), randLastName())
	} else if choice == 1 {
		n := rand.Intn(1)
		str = fmt.Sprintf("%d%s.%s", randInt(1960, time.Now().Year()), randFirstName(n), randLastName())
	} else {
		str = fmt.Sprintf("%s", randString(randInt(4, 10)))
	}

	return strings.ToLower(fmt.Sprintf("%s@%s", str, domains[rand.Intn(len(domains)-1)]))
}

func randLetters(n int, lr []rune) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = lr[rand.Intn(len(lr)-1)]
	}

	return string(b)
}

func randDate(min, max string) string {
	minTime, err := time.Parse("02-01-2006", min)
	if err != nil {
		panic(err)
	}
	maxTime, err := time.Parse("02-01-2006", max)
	if err != nil {
		panic(err)
	}

	return time.UnixMicro(int64(randInt(int(minTime.Unix()), int(maxTime.Unix())))).Format("02-01-2006")
}

func randHash(hashName string) string {
	var (
		rndStr = []byte(randString(10))
	)

	switch strings.ToLower(hashName) {
	case "md5":
		return fmt.Sprintf("%x", md5.Sum(rndStr))
	case "sha1":
		return fmt.Sprintf("%x", sha1.Sum(rndStr))
	case "sha256":
		return fmt.Sprintf("%x", sha256.Sum256(rndStr))
	case "sha512":
		return fmt.Sprintf("%x", sha512.Sum512(rndStr))
	default:
		return randHash("md5")
	}
}

func toLower(str string) string {
	return strings.ToLower(str)
}

func toUpper(str string) string {
	return strings.ToUpper(str)
}

func toTitle(str string) string {
	return cases.Title(language.English, cases.Compact).String(str)
}
