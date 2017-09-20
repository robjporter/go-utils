package main

import (
	"fmt"
	"os"
	"path/filepath"
	"syscall"
	"time"

	"./cron"
	"./encode/json"
	"./filesystem/files"
	"./filesystem/path"
	"./go/as"
	"./go/count"
	"./go/each"
	"./go/hashmap"
	"./go/iff"
	"./go/is"
	"./go/rwmap"
	"./go/safemap"
	"./go/text"
	"./maths"
	"./scheduler"
	"./sort"
	"./system"
	"./system/environment"
	"./system/ipfilter"
	"./system/signals"
	"./template"
	"./terminal/banner"
	"./terminal/colours"
	"./terminal/display"
	"./terminal/rainbow"
	"./terminal/table"
	"./timedate/dates"
	"./timedate/duration"
	"./timedate/times"
	"./timing"
	"./version"
)

var cronjob, cronjob2 *cron.CronJob
var s *signals.Signal

func test() {
	fmt.Println("\n*******************************************************")
	fmt.Println("**           TIMED EVENT TRIGGERED NOW               **")
	fmt.Println("*******************************************************")
}

func test2() {
	fmt.Println("\n*******************************************************")
	fmt.Println("**           TIMED EVENT TRIGGERED NOW 2             **")
	fmt.Println("*******************************************************")
}

func test3() {
	fmt.Println("\n*******************************************************")
	fmt.Println("**      SCHEDULER TIMED EVENT TRIGGERED NOW 3        **")
	fmt.Println("*******************************************************")
}

func clearscreenanddisplaybanner() {
	display.ClearScreen()
	banner.PrintNewFigure("LIBRARY", "varsity", true)
	time.Sleep(2 * time.Second)
}

func setupsignals() {
	s = signals.New()
	s.Bind(syscall.SIGHUP, hupHandler)
	s.Bind(syscall.SIGTERM, termHandler)
	s.Bind(syscall.SIGINT, intHandler)
	s.Start()
}

func hupHandler() {
	fmt.Println("HUP")
	os.Exit(0)
	return
}

func termHandler() {
	fmt.Println("TERM")
	os.Exit(0)
	return
}

func intHandler() {
	fmt.Println("INT")
	os.Exit(0)
	return
}

func setupschedulersandinitiate() {
	timing.Timer("main")
	cronjob = cron.New().SetName("Test").SetInterval(4).SetCallback(test)
	cronjob.Run()
	cronjob2 = cron.New().SetName("Test2").SetInterval(6).SetCallback(test2)
	cronjob2.Run()

	scheduler.Every(5).Seconds().NotImmediately().Run(test3)
	scheduler.Every().Day().Run(test3)
	scheduler.Every().Monday().At("08:30").Run(test3)

	s.Start()
}

func displaymemoryutilisation() {
	fmt.Println("")
	fmt.Println("MEMORY *******************************************************")
	fmt.Println("Memory Allocated:             >", system.MemAllocated())
	fmt.Println("Memory In Use:                >", system.MemInUse())
}

func displayfinishtimer() {
	fmt.Println("")
	fmt.Println("TIMER *******************************************************")
	mainTimer := timing.Timer("main")
	fmt.Println("Timer for Script:                 >", mainTimer)
}

func isempty() {
	fmt.Println("")
	fmt.Println("IS EMPTY *******************************************************")
	fmt.Println("IsEmpty(''):                  >", is.IsEmpty(""))
	fmt.Println("IsEmpty('test'):              >", is.IsEmpty("test"))
	fmt.Println("IsEmpty('  '):                >", is.IsEmpty("  "))
	fmt.Println("IsNotEmpty(''):               >", is.IsNotEmpty(""))
	fmt.Println("IsNotEmpty('test'):           >", is.IsNotEmpty("test"))
	fmt.Println("IsNotEmpty('  '):             >", is.IsNotEmpty("  "))
	fmt.Println("IsBlank(''):                  >", is.IsBlank(""))
	fmt.Println("IsBlank('test'):              >", is.IsBlank("test"))
	fmt.Println("IsUppercase('a'):             >", is.IsStringUppercase("a"))
	fmt.Println("IsUppercase('A'):             >", is.IsStringUppercase("A"))
	fmt.Println("IsStringAlpha('ab!'):         >", is.IsStringAllAlpha("ab!"))
	fmt.Println("IsStringAlpha('Ab'):          >", is.IsStringAllAlpha("Ab"))
	fmt.Println("IsStringAlpha('ab4'):         >", is.IsStringAllAlpha("ab4"))
	fmt.Println("IsStringContainAlpha('ab4'):  >", is.IsStringContainAlpha("ab4"))
	fmt.Println("IsStringContainAlpha('ab'):   >", is.IsStringContainAlpha("ab"))
	fmt.Println("IsStringContainAlpha('AB4'):  >", is.IsStringContainAlpha("AB4"))
	fmt.Println("IsStringContainNumber('ab4'): >", is.IsStringContainNumber("ab4"))
	fmt.Println("IsStringContainNumber('ab'):  >", is.IsStringContainNumber("ab"))
	fmt.Println("IsStringContainNumber('AB4'): >", is.IsStringContainNumber("AB4"))
	fmt.Println("IsAlphaNumeric('AB4'):        >", is.IsAlphaNumeric("AB4"))
	fmt.Println("IsAlphaNumeric('AB'):         >", is.IsAlphaNumeric("AB"))
	fmt.Println("IsAlphaNumeric('ab4'):        >", is.IsAlphaNumeric("ab4"))
	fmt.Println("IsAlphaNumeric('44'):         >", is.IsAlphaNumeric("44"))
	fmt.Println("IsBlank(''):                  >", is.IsBlank(""))
	fmt.Println("IsBlank('  '):                >", is.IsBlank("  "))
	fmt.Println("IsNotBlank(''):               >", is.IsNotBlank(""))
	fmt.Println("IsNotBlank('test'):           >", is.IsNotBlank("test"))
	fmt.Println("IsNotBlank('  '):             >", is.IsNotBlank("  "))
	fmt.Println("Reverse('test'):              >", is.Reverse("test"))
}

func isemail() {
	fmt.Println("")
	fmt.Println("IS EMAIL *******************************************************")

	fmt.Println("IsEmail(''):                  >", is.IsEmail(""))
	fmt.Println("IsEmail('test'):              >", is.IsEmail("test"))
	fmt.Println("IsEmail('test.test'):         >", is.IsEmail("test.test"))
	fmt.Println("IsEmail('test@test'):         >", is.IsEmail("test@test"))
	fmt.Println("IsEmail('test@test.test'):    >", is.IsEmail("test@test.com"))
}

func isint() {
	fmt.Println("")
	fmt.Println("IS INT *******************************************************")
	fmt.Println("IsInt('4'):                   >", is.IsInt("4"))
	fmt.Println("IsInt(4):                     >", is.IsInt(4))
	fmt.Println("IsInt(4.4):                   >", is.IsInt(4.4))
	fmt.Println("IsInt(true):                  >", is.IsInt(true))
	fmt.Println("IsInt(nil):                   >", is.IsInt(nil))
	fmt.Println("IsInt('07-06-2017'):          >", is.IsInt(time.Now()))
}

func isfloat() {
	fmt.Println("")
	fmt.Println("IS FLOAT *******************************************************")
	fmt.Println("IsFloat('4'):                 >", is.IsFloat("4"))
	fmt.Println("IsFloat(4):                   >", is.IsFloat(4))
	fmt.Println("IsFloat(4.4):                 >", is.IsFloat(4.4))
	fmt.Println("IsFloat(true):                >", is.IsFloat(true))
	fmt.Println("IsFloat(nil):                 >", is.IsFloat(nil))
	fmt.Println("IsInt('07-06-2017'):          >", is.IsFloat(time.Now()))
}

func isbool() {
	fmt.Println("")
	fmt.Println("IS BOOL *******************************************************")
	fmt.Println("IsBool('4'):                  >", is.IsBool("4"))
	fmt.Println("IsBool(4):                    >", is.IsBool(4))
	fmt.Println("IsBool(4.4):                  >", is.IsBool(4.4))
	fmt.Println("IsBool(true):                 >", is.IsBool(true))
	fmt.Println("IsBool(nil):                  >", is.IsBool(nil))
	fmt.Println("IsInt('07-06-2017'):          >", is.IsBool(time.Now()))
}

func isstring() {
	fmt.Println("")
	fmt.Println("IS STRING *******************************************************")
	fmt.Println("IsString('4'):                >", is.IsString("4"))
	fmt.Println("IsString(4):                  >", is.IsString(4))
	fmt.Println("IsString(4.4):                >", is.IsString(4.4))
	fmt.Println("IsString(true):               >", is.IsString(true))
	fmt.Println("IsString(nil):                >", is.IsString(nil))
	fmt.Println("IsString('07-06-2017'):       >", is.IsString(time.Now()))
}

func istime() {
	fmt.Println("")
	fmt.Println("IS TIME *******************************************************")
	fmt.Println("IsTime('4'):                  >", is.IsTime("4"))
	fmt.Println("IsTime(4):                    >", is.IsTime(4))
	fmt.Println("IsTime(4.4):                  >", is.IsTime(4.4))
	fmt.Println("IsTime(true):                 >", is.IsTime(true))
	fmt.Println("IsTime(nil):                  >", is.IsTime(nil))
	fmt.Println("IsTime(now()):                >", is.IsTime(time.Now()))
}

func isinintslice() {
	fmt.Println("")
	fmt.Println("IS IN INT SLICE *******************************************************")
	fmt.Println("IsInIntSlice():               >", is.IsInIntSlice(4, []int{1, 2, 3, 4}))
	fmt.Println("IsInIntSlice():               >", is.IsInIntSlice(4, []int{1, 2, 3, 5}))
}

func isinstringslice() {
	fmt.Println("")
	fmt.Println("IS IN STRING SLICE *******************************************************")
	fmt.Println("IsInStringSlice():            >", is.IsInStringSlice("a", []string{"a", "b", "c", "d"}))
	fmt.Println("IsInStringSlice():            >", is.IsInStringSlice("a", []string{"b"}))
}

func isinslice() {
	fmt.Println("")
	fmt.Println("IS IN SLICE *******************************************************")
	/*a := []interface{"1", "2", "3", "4"}
	fmt.Println("IsInSlice():                  >", is.IsInSlice(4, a))
	a = []interface{1, 2, 3, 5}
	fmt.Println("IsInSlice():                  >", is.IsInSlice(4, a))
	fmt.Println("IsInSlice():                  >", is.IsInSlice("a", []string{"a", "b", "c", "d"}))
	fmt.Println("IsInSlice():                  >", is.IsInSlice("a", []string{"b"}))
	*/
}

func isipaddress() {
	fmt.Println("")
	fmt.Println("IS IP ADDRESS *******************************************************")
	fmt.Println("IS IP (1111.1111.1111.1111):  >", is.IsIPAddress("1111.1111.1111.1111"))
	fmt.Println("IS IP (111.111.111.111):      >", is.IsIPAddress("111.111.111.111"))
	fmt.Println("IS IP (11.11.11.11):          >", is.IsIPAddress("11.11.11.11"))
	fmt.Println("IS IP (1.1.1.1):              >", is.IsIPAddress("1.1.1.1"))
}

func isfuncs() {
	time.Sleep(1 * time.Second)
	isempty()
	time.Sleep(1 * time.Second)
	isemail()
	time.Sleep(1 * time.Second)
	isint()
	time.Sleep(1 * time.Second)
	isfloat()
	time.Sleep(1 * time.Second)
	isbool()
	time.Sleep(1 * time.Second)
	isstring()
	time.Sleep(1 * time.Second)
	istime()
	time.Sleep(1 * time.Second)
	isinintslice()
	time.Sleep(1 * time.Second)
	isinstringslice()
	time.Sleep(1 * time.Second)
	isinslice()
	time.Sleep(1 * time.Second)
	isipaddress()
}

func ifthen() {
	fmt.Println("")
	fmt.Println("IF THEN *******************************************************")
	fmt.Println("IfThen(4>2):                  >", iff.IfThen(4 > 2, "4 is greater than 2"))
	fmt.Println("IfThen(4>2):                  >", iff.IfThen(4 > 2, true))
	fmt.Println("IfThenElse(4>2):              >", iff.IfThenElse(4 > 2, "4 is greater than 2", "4 is less than 2"))
	fmt.Println("IfThenElse(4>2):              >", iff.IfThenElse(4 > 2, true, false))
	fmt.Println("IfThen(4>2):                  >", iff.IfThen(1 == 1, "Yes"))            // "Yes"
	fmt.Println("IfThen(4>2):                  >", iff.IfThen(1 != 1, "Woo"))            // nil
	fmt.Println("IfThen(4>2):                  >", iff.IfThen(1 < 2, "Less"))            // "Less"
	fmt.Println("IfThen(4>2):                  >", iff.IfThenElse(1 == 1, "Yes", false)) // "Yes"
	fmt.Println("IfThen(4>2):                  >", iff.IfThenElse(1 != 1, nil, 1))       // 1
	fmt.Println("IfThen(4>2):                  >", iff.IfThenElse(1 < 2, nil, "No"))     // nil
}

func ifequals() {
	fmt.Println("")
	fmt.Println("IF EQUALS *******************************************************")
	fmt.Println("IfEquals(4>2):                >", iff.IfEquals(4 > 2))
	fmt.Println("IfEquals(4<2):                >", iff.IfEquals(4 < 2))
}

func ifequalsmultipleint() {
	fmt.Println("")
	fmt.Println("IF EQUALS MULTIPLE INT *******************************************************")
	fmt.Println("IfEqualsMultipleInt(4,0,2,4): >", iff.IfEqualsMultipleInt(4, 0, 2, 4))
	fmt.Println("IfEqualsMultipleInt(4,0,1):   >", iff.IfEqualsMultipleInt(4, 0, 1))
}

func ifequalsmultiplestring() {
	fmt.Println("")
	fmt.Println("IF EQUALS MULTIPLE STRING *******************************************************")
	fmt.Println("IfEqualsMultipleString():     >", iff.IfEqualsMultipleString("a", "a", "b"))
	fmt.Println("IfEqualsMultipleString():     >", iff.IfEqualsMultipleString("a", "b"))
}

func iffuncs() {
	time.Sleep(1 * time.Second)
	ifthen()
	time.Sleep(1 * time.Second)
	ifequals()
	time.Sleep(1 * time.Second)
	ifequalsmultipleint()
	time.Sleep(1 * time.Second)
	ifequalsmultiplestring()
}

func mathsdivide() {
	fmt.Println("")
	fmt.Println("MATHS DIVIDE *******************************************************")
	fmt.Println("DivideDetail(44,6):           >", maths.DivideDetail(44, 6))
	fmt.Println("RoundPlus_Divide(44,6):       >", maths.RoundPrecise(maths.DivideDetail(44, 6), 0.5, 5))
	fmt.Println("RoundPlus_Divide(44,6):       >", maths.Round(maths.DivideDetail(44, 6), 4))
	fmt.Println("DivideDetailRound(44,6,2):    >", maths.DivideDetailRound(44, 6, 2))
	fmt.Println("Round_DivideDetail(44,6):     >", maths.Round(maths.DivideDetail(44, 6), 2))
	fmt.Println("Divide(44,6):                 >", maths.Divide(44, 6))
	fmt.Println("Remainder(44,6):              >", maths.Remainder(44, 6))
}

func mathsoperators() {
	fmt.Println("")
	fmt.Println("MATHS *******************************************************")
	fmt.Println("Maths('*',5,2):               >", maths.Operator("*", 5, 2))
	fmt.Println("Maths('+',5,2):               >", maths.Operator("+", 5, 2))
	fmt.Println("Maths('-',5,2):               >", maths.Operator("-", 5, 2))
	fmt.Println("Maths('/',5,2):               >", maths.Operator("/", 5, 2))
	fmt.Println("Maths('%',5,2):               >", maths.Operator("%", 5, 2))
	fmt.Println("Maths(1):                     >", maths.ToRomanNumeral(1))
	fmt.Println("Maths(10):                    >", maths.ToRomanNumeral(10))
	fmt.Println("Maths(100):                   >", maths.ToRomanNumeral(100))
	fmt.Println("Maths(1000):                  >", maths.ToRomanNumeral(1000))
	fmt.Println("Maths(10000):                 >", maths.ToRomanNumeral(10000))
}

func mathfuncs() {
	time.Sleep(1 * time.Second)
	mathsdivide()
	time.Sleep(1 * time.Second)
	mathsoperators()
}

func countup() {
	fmt.Println("")
	fmt.Println("COUNT UP *******************************************************")
	for i := range count.CountUp(10) {
		fmt.Println(i)
	}
}

func countdown() {
	fmt.Println("")
	fmt.Println("COUNT DOWN *******************************************************")
	for i := range count.CountDown(10) {
		fmt.Println(i)
	}
}

func countupstep() {
	fmt.Println("")
	fmt.Println("COUNT UP STEP *******************************************************")
	for i := range count.CountUpS(1, 10) {
		fmt.Println(i)
	}
}

func countdownstep() {
	fmt.Println("")
	fmt.Println("COUNT DOWN STEP *******************************************************")
	for i := range count.CountDownS(10, 1) {
		fmt.Println(i)
	}
}

func countloop() {
	fmt.Println("")
	fmt.Println("COUNT LOOP *******************************************************")
	a := []string{"a", "b", "c", "d"}
	fmt.Println(count.Loop(a))

}

func eachloop() {
	fmt.Println("")
	fmt.Println("EACH *******************************************************")
	fn := func(s, i interface{}) {
		fmt.Println(s.(string))
	}
	s := []string{"a", "b", "c"}

	each.Each(s, fn)
}

func loopfuncs() {
	time.Sleep(1 * time.Second)
	countup()
	time.Sleep(1 * time.Second)
	countdown()
	time.Sleep(1 * time.Second)
	countupstep()
	time.Sleep(1 * time.Second)
	countdownstep()
	time.Sleep(1 * time.Second)
	countloop()
	time.Sleep(1 * time.Second)
	eachloop()
}

func getfilepaths() {
	fmt.Println("")
	fmt.Println("PATH *******************************************************")
	dir, err := filepath.Abs("./")
	if err == nil {
		fmt.Println("Path:                          >", dir)
		fmt.Println("Splitpath:                     >", path.SplitPath(dir))
		fmt.Println("Get Path:                      >", path.GetPath(dir))
		fmt.Println("Get File MD5:                  >", path.GetFileMd5("/Users/roporter/Documents/Development/Work/MyDevelopment/is/main.go"))
		fmt.Println("Parentpath:                    >", path.ParentPath(dir))
		fmt.Println("Relativepath:                  >", path.RelativePath(dir, path.ParentPath(path.ParentPath(dir))))
		fmt.Println("BaseName:                      >", path.BaseName(dir))
		fmt.Println("ListFiles:                     >", path.ListFilesRecursive("", dir, false))
		tmp := path.ListFilesRecursive("", dir, false)
		for i := 0; i < len(tmp); i++ {
			pather := dir + "/" + tmp[i]
			tmp2, _ := path.FileMode(pather)
			tmp3, _ := path.FileSize(pather)
			fmt.Println("ListFilesInfo:                 >", pather, " - ", tmp2)
			fmt.Println("ListFilesSize:                 >", pather, " - ", tmp3)
		}
	}
}

func filepathfuncs() {
	time.Sleep(1 * time.Second)
	getfilepaths()
}

func humanizestrings() {
	fmt.Println("")
	fmt.Println("TEXT *******************************************************")
	fmt.Println("Humanize 'this Is a test'         >", text.HumanizeString("this Is a test"))
	fmt.Println("Humanize 'ThisIsATest'            >", text.HumanizeString("ThisIsATest"))
	fmt.Println("Humanize 1                        >", text.ToOrdinise(1))
	fmt.Println("Humanize 2                        >", text.ToOrdinise(2))
	fmt.Println("Humanize 3                        >", text.ToOrdinise(3))
	fmt.Println("Humanize 4                        >", text.ToOrdinise(4))
	fmt.Println("Humanize 11                       >", text.ToOrdinise(11))
	fmt.Println("Humanize 15                       >", text.ToOrdinise(15))
	fmt.Println("Humanize 21                       >", text.ToOrdinise(21))
	fmt.Println("Humanize 25                       >", text.ToOrdinise(25))
	fmt.Println("Humanize 31                       >", text.ToOrdinise(31))
}

func cleantext() {
	fmt.Println("")
	fmt.Println("TEXT CLEAN TEXT *******************************************************")
	fmt.Println("CLEAN TEXT OF FOREIGN CHARS       >", text.CleanTextForeignCharacters("testing"))
	fmt.Println("CLEAN TEXT OF FOREIGN CHARS       >", text.CleanTextForeignCharacters("Je prends une thé chaud, s'il vous plaît"))
}

func uuid() {
	fmt.Println("")
	fmt.Println("TEXT UUID *******************************************************")
	fmt.Println("UUID                              >", text.UUID())
	fmt.Println("UUID                              >", text.UUID())
	fmt.Println("UUID                              >", text.UUID())
	fmt.Println("UUID                              >", text.UUID())
}

func wordfrequency() {
	fmt.Println("")
	fmt.Println("TEXT WORDCOUNT *******************************************************")
	fmt.Println("WordCount                         >", text.WordCount("This is a sample. To count the number of words."))
	fmt.Println("WordFrequency                     >", text.WordFrequency("This is a sample. To count the number of words."))
	fmt.Println("WordFrequency                     >", text.WordFrequency("This is a sample. To count the number of words.  Plus this is another test to add some more word frequency."))
}

func md5encode() {
	fmt.Println("")
	fmt.Println("TEXT MD5 *******************************************************")
	fmt.Println("MD5 ('Hellow World')              >", text.MD5("Hellow World"))
	fmt.Println("MD5 ('Hello World 1234!')         >", text.MD5("Hello World 1234!"))
}

func plurals() {
	fmt.Println("")
	fmt.Println("TEXT PLURALS *******************************************************")
	fmt.Println("Pluralise ('test')                >", text.Pluralise("test"))
	fmt.Println("Pluralise ('testy')               >", text.Pluralise("testy"))
}

func announcements() {
	fmt.Println("")
	fmt.Println("TEXT ANNOUNCEMENTS *******************************************************")
	fmt.Println("Announce (message)                >", text.Announcement("test"))
	fmt.Println("Make Announce (Message)           >", text.MakeAnnouncement("testy"))
}

func textfuncs() {
	time.Sleep(1 * time.Second)
	humanizestrings()
	time.Sleep(1 * time.Second)
	cleantext()
	time.Sleep(1 * time.Second)
	uuid()
	time.Sleep(1 * time.Second)
	wordfrequency()
	time.Sleep(1 * time.Second)
	md5encode()
	time.Sleep(1 * time.Second)
	plurals()
	time.Sleep(1 * time.Second)
	announcements()
}

func templatestring() {
	fmt.Println("")
	fmt.Println("TEMPLATE *******************************************************")
	values := make(map[string]interface{})
	in := "Hello {{ name|capfirst }}!\n\nPongo2 Version = {{pongo2.version}}"
	values["name"] = "florian"
	fmt.Println("Template 'Hello <?>'              >", template.TemplateToString(in, values))
	in = "sorted int map '{% for key in intmap sorted %}{{ key }} {% endfor %}'"
	values["intmap"] = map[int]string{
		1: "one",
		5: "five",
		2: "two",
	}
	fmt.Println("Template Sort '1,5,2'             >", template.TemplateToString(in, values))
}

func templatefuncs() {
	time.Sleep(1 * time.Second)
	templatestring()
}

func rainbowtext() {
	fmt.Println("")
	fmt.Println("RAINBOW TEXT *******************************************************")
	fmt.Println("Rainbow - 'Rainbow'               >", rainbow.Foreground("Rainbow", 255, 0, 0))
	fmt.Println("Rainbow - 'Rainbow'               >", rainbow.ForegroundGradient("Rainbow", 100, 0, 0, 20, 0, 0))
	fmt.Println("Rainbow - 'Rainbow'               >", rainbow.Background("Rainbow", 255, 0, 0, 255, 255, 255))
	fmt.Println("Rainbow - 'Rainbow'               >", rainbow.BackgroundGradient("Rainbow", 255, 0, 0, 100, 100, 100, 20, 20, 20))
}

func rainbowtextloop() {
	fmt.Println("")
	fmt.Println("RAINBOW LOOP *******************************************************")
	for r := 0; r < 255; r += 50 {
		for g := 0; g < 255; g += 50 {
			for b := 0; b < 255; b += 50 {
				fmt.Println("Rainbow - 'Rainbow'               >", rainbow.Foreground("This is a test message to see the effect of different colours.", uint(r), uint(g), uint(b)))
			}
		}
	}
}

func tabledisplay() {
	t, e := table.NewTableWriter([]string{"ID", "Name", "Description", "Cost"}, []int{10, 35, 50, 20})
	if e != nil {
		fmt.Println(e)
	}
	t.PrintHeader()
	t.PrintRow([]string{"1", "Test name", "Test description", "cost"}, table.AlignLeft)
	t.PrintFooter()
	summaryRow1 := "Id's: 3"
	summaryRow2 := "Count: 1000022"
	t.PrintRowAsOneColumn(summaryRow1, table.AlignLeft)
	t.PrintRowAsOneColumn(summaryRow2, table.AlignLeft)
	t.PrintFooter()
}

func terminalfuncs() {
	time.Sleep(1 * time.Second)
	rainbowtext()
	time.Sleep(1 * time.Second)
	rainbowtextloop()
	time.Sleep(1 * time.Second)
	tabledisplay()
}

func bubblesort() {
	fmt.Println("")
	fmt.Println("SORT *******************************************************")
	fmt.Println("Bubble Sort ('askingedik')        >", sort.BubbleSort([]byte("askingedik")))
	fmt.Println("Bubble Sort (5, 1, 4, 2, 8)       >", sort.BubbleSort([]byte{5, 1, 4, 2, 8}))
}

func sortfuncs() {
	time.Sleep(1 * time.Second)
	bubblesort()
}

func jsonprettyprint() {
	fmt.Println("")
	fmt.Println("JSON *******************************************************")
	x := map[string]interface{}{"number": 1, "string": "cool", "bool": true, "float": 1.5}
	fmt.Println("Non pretty JSON                   >", x)
	output, _ := json.PrettyJson(x)
	fmt.Println("Pretty JSON                       >", output)
}

func jsonfuncs() {
	time.Sleep(1 * time.Second)
	jsonprettyprint()
}

func rwmapfuncs() {
	fmt.Println("")
	fmt.Println("RWMAP *******************************************************")
	m := rwmap.RwMap(1000)
	m.Store(1, "a")
	m.Store(2, "b")
	m.Store(3, "c")
	m.Store(4, "d")
	m.Store(5, "e")
	m.Store(6, "f")
	fmt.Println("RWMap Length                      >", m.Len())
	val, valid := m.Load(1)
	fmt.Println("RWMap Valid position 1            >", valid)
	fmt.Println("RWMap Value position 1            >", val)
}

func hashfuncs() {
	fmt.Println("")
	fmt.Println("HASH MAP *******************************************************")
	hm := hashmap.NewHashMap()
	hm.Set("foo", "123")
	hm.Set("bar", "456")
	fmt.Println(hm)
}

func safemapfuncs() {
	fmt.Println("")
	fmt.Println("SAFE MAP *******************************************************")
	safe := safemap.NewSafeMap()
	fmt.Println("SAFE MAP ISEMPTY                  >", safe.IsEmpty())
	safe.Set("item1", "Item1")
	safe.Set("item2", "Item2")
	safe.Set("item3", "Item3")
	safe.Set("item4", "Item4")
	fmt.Println("SAFE MAP ISEMPTY                  >", safe.IsEmpty())
	fmt.Println("SAFE MAP LENGTH                   >", safe.Len())
	fmt.Println("SAFE MAP LIST                     >", safe.List())
	fmt.Println("SAFE MAP HAS 'item1'              >", safe.IsExists("item1"))
	fmt.Println("SAFE MAP GET 'item1'              >", safe.Get("item1"))
	safe.Delete("item1")
	fmt.Println("SAFE MAP LENGTH                   >", safe.Len())
	fmt.Println("SAFE MAP LIST                     >", safe.List())
	fmt.Println("SAFE MAP HAS 'item1'              >", safe.IsExists("item1"))
}

func mapfuncs() {
	time.Sleep(1 * time.Second)
	rwmapfuncs()
	time.Sleep(1 * time.Second)
	hashfuncs()
	time.Sleep(1 * time.Second)
	safemapfuncs()
	time.Sleep(1 * time.Second)
}

func parsestringversion() {
	fmt.Println("")
	fmt.Println("PARSE STRING VERSION *******************************************************")
	fmt.Println("Parse Version 1.02.3.             >", version.ParseVersionString("1.02.3."))
	fmt.Println("Parse Version 1.2..3              >", version.ParseVersionString("1.2..3"))
	fmt.Println("Parse Version 1.00.2              >", version.ParseVersionString("1.00.2"))
	fmt.Println("Parse Version 1.02.a              >", version.ParseVersionString("1.02.a"))
	fmt.Println("Parse Version 1.02.a.b            >", version.ParseVersionString("1.02.a.b"))
	fmt.Println("Parse Version 1.02(a)             >", version.ParseVersionString("1.02(a)"))
	fmt.Println("Parse Version 1.02(4a)            >", version.ParseVersionString("1.02(4a)"))
	fmt.Println("Parse Version 1.02(4b)            >", version.ParseVersionString("1.02(4b)"))
	fmt.Println("Parse Version 1.02(5a)            >", version.ParseVersionString("1.02(5a)"))
}

func parsescompareversion() {
	fmt.Println("")
	fmt.Println("PARSE COMPARE VERSION *******************************************************")
	fmt.Println("Compare 3.1(4b) v 2.02(3a)        >", version.CompareStrings("3.1(4b)", "2.02(3a)"))
	fmt.Println("Compare 3.1(4b) v 3.1(4b)         >", version.CompareStrings("3.1(4b)", "3.1(4b)"))
	fmt.Println("Compare 2.02(3a) v 3.1(4b)        >", version.CompareStrings("2.02(3a)", "3.1(4b)"))
}

func parseversionfuncs() {
	time.Sleep(1 * time.Second)
	parsestringversion()
	time.Sleep(1 * time.Second)
	parsescompareversion()
}

func parsedirs() {
	fmt.Println("")
	fmt.Println("FILES DIRECTORIES *******************************************************")
	r, err := files.NewDirReader("/Users/roporter/Documents", files.N_RECURSE)
	if err != nil {
		fmt.Println("ERROR", err)
		return
	}
	r.Filter = files.MultiFilter(files.DirFilter())
	for {
		info, err := r.Next()
		if err != nil {
			fmt.Println("ERROR", err)
			return
		}
		if info == nil {
			break
		}
		fmt.Println(info.Path)
	}
}

func parsefiles() {
	fmt.Println("")
	fmt.Println("FILES *******************************************************")
	r, err := files.NewDirReader("./filesystem/files", files.N_RECURSE)
	if err != nil {
		fmt.Println("ERROR", err)
		return
	}
	r.Filter = files.MultiFilter(files.FileFilter())
	for {
		info, err := r.Next()
		if err != nil {
			fmt.Println("ERROR", err)
			return
		}
		if info == nil {
			break
		}
		fmt.Println(info.Path)
		fmt.Println(files.GetFileMd5(info.Path))
	}
}

func dirandfilefuncs() {
	time.Sleep(1 * time.Second)
	parsedirs()
	time.Sleep(1 * time.Second)
	parsefiles()
}

func datefuncs() {
	fmt.Println("")
	fmt.Println("DATES *******************************************************")
	d := dates.MakeDuration(time.Minute * 72)
	fmt.Println("Dates - Set future date           >", "72 minutes from now")
	fmt.Println("How many hours from now           >", d.RemainingHoursStr(), "hours")
	fmt.Println("How many minutes                  >", d.RemainingMinutesStr(), "minutes")
}

func addtoallsimple() {
	fmt.Println("")
	fmt.Println("DATES ADD SIMPLE *******************************************************")
	t := times.NewTodayAuto()
	t6 := t.Copy()
	fmt.Println("COPY:                             >", t6)
	result, _ := t.AddDecade()
	fmt.Println("ADD DECADE:                       >", result)
	result, _ = t.AddYear()
	fmt.Println("ADD YEAR:                         >", result)
	result, _ = t.AddMonth()
	fmt.Println("ADD MONTH:                        >", result)
	result, _ = t.AddWeek()
	fmt.Println("ADD WEEK:                         >", result)
	result, _ = t.AddDay()
	fmt.Println("ADD DAY:                          >", result)
	result, _ = t.AddHour()
	fmt.Println("ADD HOUR:                         >", result)
	result, _ = t.AddMinute()
	fmt.Println("ADD MINUTE:                       >", result)
	result, _ = t.AddSecond()
	fmt.Println("ADD SECOND:                       >", result)
}

func addtoalladvanced() {
	fmt.Println("")
	fmt.Println("DATES ADD ADVANCED *******************************************************")
	t := times.NewTodayAuto()
	result, _ := t.AddDecades(2)
	fmt.Println("ADD DECADES:                   (2)>", result)
	result, _ = t.AddYears(2)
	fmt.Println("ADD YEARS:                     (2)>", result)
	result, _ = t.AddMonths(2)
	fmt.Println("ADD MONTHS:                    (2)>", result)
	result, _ = t.AddWeeks(2)
	fmt.Println("ADD WEEKS:                     (2)>", result)
	result, _ = t.AddDays(2)
	fmt.Println("ADD DAYS:                      (2)>", result)
	result, _ = t.AddHours(2)
	fmt.Println("ADD HOURS:                     (2)>", result)
	result, _ = t.AddMinutes(2)
	fmt.Println("ADD MINUTES:                   (2)>", result)
	result, _ = t.AddSeconds(2)
	fmt.Println("ADD SECONDS:                   (2)>", result)
}

func subtoallsimple() {
	fmt.Println("")
	fmt.Println("DATES SUB SIMPLE *******************************************************")
	t := times.NewTodayAuto()
	result, _ := t.SubDecade()
	fmt.Println("SUB DECADE:                     ()>", result)
	result, _ = t.SubYear()
	fmt.Println("SUB YEAR:                       ()>", result)
	result, _ = t.SubMonth()
	fmt.Println("SUB MONTH:                      ()>", result)
	result, _ = t.SubWeek()
	fmt.Println("SUB WEEK:                       ()>", result)
	result, _ = t.SubDay()
	fmt.Println("SUB DAY:                        ()>", result)
	result, _ = t.SubHour()
	fmt.Println("SUB HOUR:                       ()>", result)
	result, _ = t.SubMinute()
	fmt.Println("SUB MINUTE:                     ()>", result)
	result, _ = t.SubSecond()
	fmt.Println("SUB SECOND:                     ()>", result)
}

func subtoalladvanced() {
	fmt.Println("")
	fmt.Println("DATES SUB ADVANCED *******************************************************")
	t := times.NewTodayAuto()
	result, _ := t.SubDecades(2)
	fmt.Println("SUB DECADES:                   (2)>", result)
	result, _ = t.SubYears(2)
	fmt.Println("SUB YEARS:                     (2)>", result)
	result, _ = t.SubMonths(5)
	fmt.Println("SUB MONTHS:                    (5)>", result)
	result, _ = t.SubWeeks(2)
	fmt.Println("SUB WEEKS:                     (2)>", result)
	result, _ = t.SubDays(6)
	fmt.Println("SUB DAYS:                      (6)>", result)
	result, _ = t.SubHours(2)
	fmt.Println("SUB HOURS:                     (2)>", result)
	result, _ = t.SubMinutes(2)
	fmt.Println("SUB MINUTES:                   (2)>", result)
	result, _ = t.SubSeconds(2)
	fmt.Println("SUB SECONDS:                   (2)>", result)
}

func startofall() {
	fmt.Println("")
	fmt.Println("DATES START OF ALL *******************************************************")
	t := times.NewTodayAuto()
	result, _ := t.StartOfCentury()
	fmt.Println("START OF CENTURY:              ()>", result)
	result, _ = t.StartOfDecade()
	fmt.Println("START OF DECADE:               ()>", result)
	result, _ = t.StartOfYear()
	fmt.Println("START OF YEAR:                 ()>", result)
	result, _ = t.StartOfMonth()
	fmt.Println("START OF MONTH:                ()>", result)
	result, _ = t.StartOfWorkWeek()
	fmt.Println("START OF WORK WEEK:            ()>", result)
	result, _ = t.StartOfWeek()
	fmt.Println("START OF WEEK:                 ()>", result)
	result, _ = t.StartOfDay()
	fmt.Println("START OF DAY:                  ()>", result)
	result, _ = t.StartOfHour()
	fmt.Println("START OF HOUR:                 ()>", result)
	result, _ = t.StartOfMinute()
	fmt.Println("START OF MINUTE:               ()>", result)
}

func endofall() {
	fmt.Println("")
	fmt.Println("DATES END OF ALL *******************************************************")
	t := times.NewTodayAuto()
	result, _ := t.EndOfCentury()
	fmt.Println("END OF CENTURY:               ()>", result)
	result, _ = t.EndOfDecade()
	fmt.Println("END OF DECADE:                ()>", result)
	result, _ = t.EndOfYear()
	fmt.Println("END OF YEAR:                  ()>", result)
	result, _ = t.EndOfMonth()
	fmt.Println("END OF MONTH:                 ()>", result)
	result, _ = t.EndOfWorkWeek()
	fmt.Println("END OF WORK WEEK:             ()>", result)
	result, _ = t.EndOfWeek()
	fmt.Println("END OF WEEK:                  ()>", result)
	result, _ = t.EndOfDay()
	fmt.Println("END OF DAY:                   ()>", result)
	result, _ = t.EndOfHour()
	fmt.Println("END OF HOUR:                  ()>", result)
	result, _ = t.EndOfMinute()
	fmt.Println("END OF MINUTE:                ()>", result)
}

func quarterinfo() {
	fmt.Println("")
	fmt.Println("DATES QUARTER INFO *******************************************************")
	t := times.NewTodayAuto()
	fmt.Println("Is 1st QUARTER:               ()>", t.Is1stQuarter())
	fmt.Println("Is 2nd QUARTER:               ()>", t.Is2ndQuarter())
	fmt.Println("Is 3rd QUARTER:               ()>", t.Is3rdQuarter())
	fmt.Println("Is 4th QUARTER:               ()>", t.Is4thQuarter())
	result, _ := t.StartOfQuarter()
	fmt.Println("START OF QUARTER:             ()>", result)
	result, _ = t.EndOfQuarter()
	fmt.Println("END OF QUARTER:               ()>", result)
	fmt.Println("QUARTER:                      ()>", t.Quarter())
	fmt.Println("QUARTER NUMBER:               ()>", t.QuarterNumber())
}

func seasoninfo() {
	fmt.Println("")
	fmt.Println("DATES SEASON INFO *******************************************************")
	t := times.NewTodayAuto()
	fmt.Println("SEASON:                       ()>", t.Season())
	fmt.Println("TIME TO SPRING:               ()>", t.TimeToSpring())
	fmt.Println("TIME TO SPRING:               ()>", t.TimeToSpringDiff())
	fmt.Println("TIME TO SUMMER:               ()>", t.TimeToSummer())
	fmt.Println("TIME TO SUMMER:               ()>", t.TimeToSummerDiff())
	fmt.Println("TIME TO AUTUMN:               ()>", t.TimeToAutumn())
	fmt.Println("TIME TO AUTUMN:               ()>", t.TimeToAutumnDiff())
	fmt.Println("TIME TO WINTER:               ()>", t.TimeToWinter())
	fmt.Println("TIME TO WINTER:               ()>", t.TimeToWinterDiff())
}

func taxyearinfo() {
	fmt.Println("")
	fmt.Println("DATES TAX YEAR INFO *******************************************************")
	t := times.NewTodayAuto()
	fmt.Println("TAX YEAR:                     ()>", t.TaxYear())
	result, _ := t.StartOfTaxYear()
	fmt.Println("START OF UK TAX YEAR:         ()>", result)
	result, _ = t.EndOfTaxYear()
	fmt.Println("END OF UK TAX YEAR:           ()>", result)
	fmt.Println("NEXT UK TAX YEAR:             ()>", t.TimeToTaxYear())
	fmt.Println("NEXT UK TAX YEAR DIFF         ()>", t.TimeToTaxYearDiff())
}

func iscertaininfo() {
	fmt.Println("")
	fmt.Println("DATES IS CERTAIN INFO *******************************************************")
	t := times.NewTodayAuto()
	fmt.Println("IS WEEKEND:                   ()>", t.IsWeekend())
	fmt.Println("IS WORKDAY:                   ()>", t.IsWorkday())
	fmt.Println("IS MONDAY:                    ()>", t.IsMonday())
	fmt.Println("IS TUESDAY:                   ()>", t.IsTuesday())
	fmt.Println("IS WEDNESDAY:                 ()>", t.IsWednesday())
	fmt.Println("IS THURSDAY:                  ()>", t.IsThursday())
	fmt.Println("IS FRIDAY:                    ()>", t.IsFriday())
	fmt.Println("IS SATURDAY:                  ()>", t.IsSaturday())
	fmt.Println("IS SUNDAY:                    ()>", t.IsSunday())
	date := time.Now()
	t4 := times.New(date.Year()-1, times.MonthToNumber(date.Month()), date.Day(), date.Hour(), date.Minute(), date.Second(), "Europe/London")
	t5 := times.New(date.Year()+1, times.MonthToNumber(date.Month()), date.Day(), date.Hour(), date.Minute(), date.Second(), "Europe/London")
	fmt.Println("IS BETWEEN:              (t4,t5)>", t.IsBetween(t4, t5))
	fmt.Println("IS SPRING:                    ()>", t.IsSpring())
	fmt.Println("IS SUMMER:                    ()>", t.IsSummer())
	fmt.Println("IS AUTUMN:                    ()>", t.IsAutumn())
	fmt.Println("IS WINTER:                    ()>", t.IsWinter())
}

func differenceinfo() {
	fmt.Println("")
	fmt.Println("DATES DIFFERENCE INFO *******************************************************")
	t := times.NewTodayAuto()
	date := time.Now()
	t3 := times.New(date.Year()+1, times.MonthToNumber(date.Month())+1, date.Day()+8, date.Hour()+1, date.Minute()+1, date.Second()+4, "Europe/London")
	fmt.Println("DIFFERENCE:                   ()>", t.DifferenceDiff(t3))
	fmt.Println("DIFFERENCE:                   ()>", t.Difference(t3))
	fmt.Println("DIFFINYEARS:                  ()>", t.DiffInYears(t3))
	fmt.Println("DIFFINMONTHS:                 ()>", t.DiffInMonths(t3))
	fmt.Println("DIFFINWEEKS:                  ()>", t.DiffInWeeks(t3))
	fmt.Println("DIFFINDAYS:                   ()>", t.DiffInDays(t3))
	fmt.Println("DIFFINHOURS:                  ()>", t.DiffInHours(t3))
	fmt.Println("DIFFINMINUTES:                ()>", t.DiffInMinutes(t3))
	fmt.Println("DIFFINSECONDS:                ()>", t.DiffInSeconds(t3))
}

func formatinfo() {
	fmt.Println("")
	fmt.Println("DATES FORMAT INFO *******************************************************")
	t := times.NewTodayAuto()
	result2, _ := t.Format822()
	fmt.Println("FORMAT 822:                   ()>", result2)
	result2, _ = t.Format1123()
	fmt.Println("FORMAT 1123:                  ()>", result2)
	result2, _ = t.Format1123z()
	fmt.Println("FORMAT 1123z:                 ()>", result2)
	result2, _ = t.Format3339()
	fmt.Println("FORMAT 3339:                  ()>", result2)
	result2, _ = t.Format3339nano()
	fmt.Println("FORMAT 3339 Nano:             ()>", result2)
	result2, _ = t.Format8222z()
	fmt.Println("FORMAT 8222z:                 ()>", result2)
	result2, _ = t.Format850()
	fmt.Println("FORMAT 850:                   ()>", result2)
	result2, _ = t.Format1()
	fmt.Println("FORMAT 1:                     ()>", result2)
	result2, _ = t.Format2()
	fmt.Println("FORMAT 2:                     ()>", result2)
	result2, _ = t.Format3()
	fmt.Println("FORMAT 3:                     ()>", result2)
	result2, _ = t.Format4()
	fmt.Println("FORMAT 4:                     ()>", result2)
	result2, _ = t.Format5()
	fmt.Println("FORMAT 5:                     ()>", result2)
	result2, _ = t.Format6()
	fmt.Println("FORMAT 6:                     ()>", result2)
	result2, _ = t.Format7()
	fmt.Println("FORMAT 7:                     ()>", result2)
	result2, _ = t.Format8()
	fmt.Println("FORMAT 8:                     ()>", result2)
	result2, _ = t.Format9()
	fmt.Println("FORMAT 9:                     ()>", result2)
	result2, _ = t.Format10()
	fmt.Println("FORMAT 10:                    ()>", result2)
	result2, _ = t.Format11()
	fmt.Println("FORMAT 11:                    ()>", result2)
	result2, _ = t.Format12()
	fmt.Println("FORMAT 12:                    ()>", result2)
	result2, _ = t.Format13()
	fmt.Println("FORMAT 13:                    ()>", result2)
	result2, _ = t.Format14()
	fmt.Println("FORMAT 14:                    ()>", result2)
	result2, _ = t.Format15()
	fmt.Println("FORMAT 15:                    ()>", result2)
	result2, _ = t.Format16()
	fmt.Println("FORMAT 16:                    ()>", result2)
	result2, _ = t.Format17()
	fmt.Println("FORMAT 17:                    ()>", result2)
	result2, _ = t.Format18()
	fmt.Println("FORMAT 18:                    ()>", result2)
	result2, _ = t.Format19()
	fmt.Println("FORMAT 19:                    ()>", result2)
	result2, _ = t.Format20()
	fmt.Println("FORMAT 20:                    ()>", result2)
}

func functioninfo() {
	fmt.Println("")
	fmt.Println("DATES FORMAT INFO *******************************************************")
	t := times.NewTodayAuto()
	date := time.Now()
	fmt.Println("IS LEAP YEAR:                 ()>", t.IsLeapYear())
	fmt.Println("NEXT LEAP YEAR:               ()>", t.NextLeapYear())
	result, _ := t.TimeNext(time.Wednesday)
	fmt.Println("TIME NEXT:           (Wednesday)>", result)
	result, _ = t.TimePrevious(time.Wednesday)
	fmt.Println("TIME LAST:           (Wednesday)>", result)
	t2 := times.New(date.Year(), times.MonthToNumber(date.Month()), date.Day()+1, date.Hour(), date.Minute(), date.Second(), "Europe/London")
	fmt.Println("IS FUTURE:                    ()>", t.IsFuture(t2))
	t2.SubDay()
	fmt.Println("IS PAST:                      ()>", t.IsPast(t2))
	result2, _ := t.ISOWeek()
	fmt.Println("ISO WEEK:                     ()>", result2)
	fmt.Println("MONTH NUMBER:              (jan)>", times.MonthNameToNumber("jan"))
	fmt.Println("MONTH NAME:                (jan)>", times.MonthNameToFullName("jan"))
	fmt.Println("MONTH NAME:              (April)>", times.MonthToNumber(date.Month()))
	fmt.Println("MONTH NAME:              (April)>", times.MonthToName(date.Month()))
	fmt.Println("LOCATION:                     ()>", t.Location())
}

func timefuncs() {
	time.Sleep(1 * time.Second)
	addtoallsimple()
	time.Sleep(1 * time.Second)
	addtoalladvanced()
	time.Sleep(1 * time.Second)
	subtoallsimple()
	time.Sleep(1 * time.Second)
	subtoalladvanced()
	time.Sleep(1 * time.Second)
	startofall()
	time.Sleep(1 * time.Second)
	endofall()
	time.Sleep(1 * time.Second)
	quarterinfo()
	time.Sleep(1 * time.Second)
	seasoninfo()
	time.Sleep(1 * time.Second)
	taxyearinfo()
	time.Sleep(1 * time.Second)
	iscertaininfo()
	time.Sleep(1 * time.Second)
	differenceinfo()
	time.Sleep(1 * time.Second)
	formatinfo()
	time.Sleep(1 * time.Second)
	functioninfo()
}

func durationfuncs() {
	fmt.Println("")
	fmt.Println("DATES FORMAT INFO *******************************************************")
	t1 := time.Now()
	time.Sleep(500 * time.Millisecond)
	t2 := time.Now()
	t3 := t1.Sub(t2)
	fmt.Println("HUMAN TIME DURATION               >", duration.HumanDuration(t3))
}

func timeanddatefuncs() {
	time.Sleep(1 * time.Second)
	datefuncs()
	time.Sleep(1 * time.Second)
	timefuncs()
	time.Sleep(1 * time.Second)
	durationfuncs()
}

func systemusername() {
	fmt.Println("")
	fmt.Println("SYSTEM USERNAME *******************************************************")
	fmt.Println("Current Username:                 >", system.GetUsername())
}

func systemenvironments() {
	fmt.Println("")
	fmt.Println("SYSTEM ENVIRONMENTS *******************************************************")
	fmt.Println("ENVIRONMENT GO PATH BIN:          >", environment.GOPATHBIN())
	fmt.Println("ENVIRONMENT GO PATH:              >", environment.GOPATH())
	fmt.Println("ENVIRONMENT PATH SEPARATOR:       >", environment.PathSeparator())
	fmt.Println("ENVIRONMENT LIST SEPARATOR:       >", environment.ListSeparator())
	fmt.Println("ENVIRONMENT IS COMPILED:          >", environment.IsCompiled())
	fmt.Println("ENVIRONMENT BUILD DEBUG:          >", environment.BuildDebug())
	fmt.Println("ENVIRONMENT CHECK ARCH:           >", environment.CheckArchitecture())
	fmt.Println("ENVIRONMENT BUILD STAMP:          >", environment.BuildStamp())
	fmt.Println("ENVIRONMENT BUILD HOST:           >", environment.BuildHost())
	fmt.Println("ENVIRONMENT COMPILER:             >", environment.Compiler())
	fmt.Println("ENVIRONMENT GO ARCH:              >", environment.GOARCH())
	fmt.Println("ENVIRONMENT GO OS:                >", environment.GOOS())
	fmt.Println("ENVIRONMENT GO ROOT:              >", environment.GOROOT())
	fmt.Println("ENVIRONMENT GO VERSION:           >", environment.GOVER())
	fmt.Println("ENVIRONMENT NUMBER CPU:           >", environment.NumCPU())
	fmt.Println("ENVIRONMENT FORMATTED TIME:       >", environment.GetFormattedTime())
}

func networking() {
	fmt.Println("")
	fmt.Println("SYSTEM IS PRIVATE IP *******************************************************")
	fmt.Println("IS PRIVATE IP 1.1.1.1:            >", system.IsPrivateIp("1.1.1.1"))
	fmt.Println("IS PRIVATE IP 10.52.208.1:        >", system.IsPrivateIp("10.52.208.1"))
	fmt.Println("IS PUBLIC IP 1.1.1.1:             >", system.IsPublicIp("1.1.1.1"))
	fmt.Println("IS PUBLIC IP 10.52.208.1:         >", system.IsPublicIp("10.52.208.1"))
}

func httpget() {
	fmt.Println("")
	fmt.Println("SYSTEM HTTP GET *******************************************************")
	e := DownloadFile("https://github.com/metaleap/go-util-net/raw/master/net.go", "./file.txt")
	fmt.Println("HTTP GET ON FILE:                 >", e)
}

func systemfuncs() {
	time.Sleep(1 * time.Second)
	systemusername()
	time.Sleep(1 * time.Second)
	systemenvironments()
	time.Sleep(1 * time.Second)
	networking()
	time.Sleep(1 * time.Second)
	httpget()
}

func astostring() {
	fmt.Println("")
	fmt.Println("AS TO STRING *******************************************************")
	fmt.Println("STRING: (32)                        >", `"`+as.ToString(32)+`"`)
	fmt.Println("STRING: (true)                      >", `"`+as.ToString(bool(true))+`"`)
	fmt.Println("STRING: ('mayonegg')                >", `"`+as.ToString("mayonegg")+`"`)         // "mayonegg"
	fmt.Println("STRING: (8)                         >", `"`+as.ToString(8)+`"`)                  // "8"
	fmt.Println("STRING: (8.31)                      >", `"`+as.ToString(8.31)+`"`)               // "8.31"
	fmt.Println("STRING: ([]byte('one time'))        >", `"`+as.ToString([]byte("one time"))+`"`) // "one time"
	fmt.Println("STRING: (nil)                       >", `"`+as.ToString(nil)+`"`)                // ""
	var foo interface{} = "one more time"
	fmt.Println("STRING: (interface{'one more time}) >", `"`+as.ToString(foo)) // "one more time"
}
func astrimmed() {
	fmt.Println("")
	fmt.Println("AS TRIMMED *******************************************************")
	fmt.Println("TRIMMED: ('    TEST      ')         >", `"`+as.Trimmed("    TEST      ")+`"`)
}
func astofloat() {
	fmt.Println("")
	fmt.Println("AS TO FLOAT *******************************************************")
	fmt.Println("FLOAT: (32.4400)                    >", as.ToFloat(32.4400))
	fmt.Println("FLOAT32: (32.4400)                  >", as.ToFloat32(32.4400))
}
func astorunelength() {
	fmt.Println("")
	fmt.Println("AS TO RUNE LENGTH*******************************************************")
	fmt.Println("RUNELENGTH: ('test')                >", as.ToRuneLength("test"))
	fmt.Println("RUNELENGTH: ('TEST')                >", as.ToRuneLength("TEST"))
	fmt.Println("RUNELENGTH: ('iiii')                >", as.ToRuneLength("iiii"))
	fmt.Println("RUNELENGTH: ('QQKK')                >", as.ToRuneLength("QQKK"))
	fmt.Println("RUNELENGTH: ('Lllm')                >", as.ToRuneLength("Lllm"))
}
func astobool() {
	fmt.Println("")
	fmt.Println("AS TO BOOL *******************************************************")
	fmt.Println("BOOL: (1)                           >", as.ToBool(1))
	fmt.Println("BOOL: (0)                           >", as.ToBool(0))
	fmt.Println("BOOL: ('1')                         >", as.ToBool("1"))
	fmt.Println("BOOL: ('true')                      >", as.ToBool("true"))
	fmt.Println("BOOL: ('down')                      >", as.ToBool("down"))
}
func astobytes() {
	fmt.Println("")
	fmt.Println("AS TO BYTES *******************************************************")
	fmt.Println("BYTES: ('Testing')                  >", as.ToBytes("Testing"))
}
func astoslice() {
	fmt.Println("")
	fmt.Println("AS TO SLICE *******************************************************")
	var foo2 []interface{}
	foo2 = append(foo2, "one") //more time"
	fmt.Println("SLICE: ('one')                      >", as.ToSlice(foo2))
}
func astoint() {
	fmt.Println("")
	fmt.Println("TO INT *******************************************************")
	fmt.Println("INT: ('1')                          >", as.ToInt("1"))
	fmt.Println("INT64: ('1')                        >", as.ToInt64("1"))
	fmt.Println("INT32: ('1')                        >", as.ToInt32("1"))
	fmt.Println("INT16: ('1')                        >", as.ToInt16("1"))
	fmt.Println("INT8: ('1')                         >", as.ToInt8("1"))
}

func astoip() {
	fmt.Println("")
	fmt.Println("TO IP *******************************************************")
	fmt.Println("IP ADDRESS: ('192.168.0.1')          >", as.ToIP("192.168.0.1"))   // "one more time"
	fmt.Println("IP ADDRESS: ('one more time')        >", as.ToIP("one more time")) //
	fmt.Println("IP ADDRESS: ('1')                    >", as.ToIP("1"))             // "one more time"
	fmt.Println("IP ADDRESS: ('1.0')                  >", as.ToIP("1.0"))           // "one more time"
	fmt.Println("IP ADDRESS: ('1.0.0')                >", as.ToIP("1.0.0"))         // "one more time"
	fmt.Println("IP ADDRESS: ('1.0.0.0/8')            >", as.ToIP("1.0.0.0/8"))     // "one more time"
}

func astobase64() {
	fmt.Println("")
	fmt.Println("TO BASE64 *******************************************************")
	fmt.Println("TOBASE64: ('This is a test')         >", as.ToBase64("This is a test"))
}

func asfrombase64() {
	fmt.Println("")
	fmt.Println("FROM BASE64 *******************************************************")
	fmt.Println("FROMBASE64: ('VGhpcyBpcyBhIHRlc3Q=') >", as.FromBase64("VGhpcyBpcyBhIHRlc3Q="))
}

func asisempty() {
	fmt.Println("")
	fmt.Println("AS IS EMPTY *******************************************************")
	fmt.Println("IP EMPTY: ('0')                      >", as.IsEmpty(0))
	fmt.Println("IP EMPTY: ('1')                      >", as.IsEmpty(1))
	fmt.Println("IP EMPTY: ('')                       >", as.IsEmpty(""))
	fmt.Println("IP EMPTY: ('sdasdass')               >", as.IsEmpty("sdasdass"))
	fmt.Println("IP EMPTY: ('[]string{}')             >", as.IsEmpty([]string{}))
}

func asiskind() {
	fmt.Println("")
	fmt.Println("AS IS KIND *******************************************************")
	fmt.Println("IS KIND: (string,0)                  >", as.IsKind("string", 0))
	fmt.Println("IS KIND: (string,'')                 >", as.IsKind("string", ""))
	fmt.Println("IS KIND: (int,0)                     >", as.IsKind("int", 0))
	fmt.Println("IS KIND: (int,'test')                >", as.IsKind("int", "test"))
}

func asofkind() {
	fmt.Println("")
	fmt.Println("AS OF KIND *******************************************************")
	fmt.Println("KIND OF: ('string')                  >", as.OfKind("string"))
	fmt.Println("KIND OF: ([]string{})                >", as.OfKind([]string{}))
	fmt.Println("KIND OF: (nil)                       >", as.OfKind(nil))
	fmt.Println("KIND OF: ([]byte('one time))         >", as.OfKind([]byte("one time")))
	fmt.Println("KIND OF: (bool(true))                >", as.OfKind(bool(true)))
	fmt.Println("KIND OF: (32)                        >", as.OfKind(32))
}
func asoftype() {
	fmt.Println("")
	fmt.Println("AS OF TYPE *******************************************************")
	fmt.Println("TYPE: (32)                           >", as.OfType(32))
	fmt.Println("TYPE: ('')                           >", as.OfType(""))
	fmt.Println("TYPE: ([]string{}])                  >", as.OfType([]string{}))
	fmt.Println("TYPE: (true)                         >", as.OfType(true))
	fmt.Println("TYPE: (1.0f)                         >", as.OfType(1.00))
	fmt.Println("TYPE: (int64(22))                    >", as.OfType(int64(22)))
}

func astotime() {
	fmt.Println("")
	fmt.Println("AS TO TIME *******************************************************")
	fmt.Println("TIME: ('2016-04-04')                 >", as.ToTime(false, "2016-04-04"))
	fmt.Println("TIME: ('04-04-2016')                 >", as.ToTime(false, "04-04-2016"))
	fmt.Println("TIME: ('2016-04-04 16:20:40')        >", as.ToTime(false, "2016-04-04 16:20:40"))
	fmt.Println("TIME: ('2016-04-04 16:20:40 +1 BST') >", as.ToTime(false, "2016-04-04 16:20:40 +1 BST"))
	t1 := time.Now()
	fmt.Println("TIME: NOW TO INT                     >", as.FromTime(t1))
	fmt.Println("TIME: INT TO TIME                    >", as.ToTime(true, as.FromTime(t1)))
}
func astoduration() {
	fmt.Println("")
	fmt.Println("AS TO DURATION *******************************************************")
	fmt.Println("DURATION: (1h44m)                    >", as.ToDuration("1h44m"))
	fmt.Println("DURATION: (44)                       >", as.ToDuration("44"))
	fmt.Println("DURATION: (44s)                      >", as.ToDuration("44s"))
	fmt.Println("DURATION: (444h)                     >", as.ToDuration("444h"))
	fmt.Println("DURATION: (88m)                      >", as.ToDuration("88m"))
}

func astofixedlengthafter() {
	fmt.Println("")
	fmt.Println("AS TO FIXED LENGTH AFTER *******************************************************")
	fmt.Println("FIXED LENGTH AFTER (*,20):           >", as.ToFixedLengthAfter("Test String", "*", 20))
	fmt.Println("FIXED LENGTH AFTER (-,50):           >", as.ToFixedLengthAfter("Test String", "-", 50))
	fmt.Println("FIXED LENGTH AFTER (*,10):           >", as.ToFixedLengthAfter("Test String", "*", 10))
	fmt.Println("FIXED LENGTH AFTER (*,8):            >", as.ToFixedLengthAfter("Test String", "*", 8))
}

func astofixedlengthbefore() {
	fmt.Println("")
	fmt.Println("AS TO FIXED LENGTH BEFORE *******************************************************")
	fmt.Println("FIXED LENGTH BEFORE (*,20):          >", as.ToFixedLengthBefore("Test String", "*", 20))
	fmt.Println("FIXED LENGTH BEFORE (-,50):          >", as.ToFixedLengthBefore("Test String", "-", 50))
	fmt.Println("FIXED LENGTH BEFORE (*,10):          >", as.ToFixedLengthBefore("Test String", "*", 10))
	fmt.Println("FIXED LENGTH BEFORE (*,8):           >", as.ToFixedLengthBefore("Test String", "*", 8))
}

func astofixedlengthcenter() {
	fmt.Println("")
	fmt.Println("AS TO FIXED LENGTH CENTER *******************************************************")
	fmt.Println("FIXED LENGTH CENTER (*,20):          >", as.ToFixedLengthCenter("Test String", "*", 20))
	fmt.Println("FIXED LENGTH CENTER (-,50):          >", as.ToFixedLengthCenter("Test String", "-", 50))
	fmt.Println("FIXED LENGTH CENTER (*,10):          >", as.ToFixedLengthCenter("Test String", "*", 10))
	fmt.Println("FIXED LENGTH CENTER (*,8):           >", as.ToFixedLengthCenter("Test String", "*", 8))

}

func asisint() {
	fmt.Println("")
	fmt.Println("AS IS INT *******************************************************")
	fmt.Println("INT: (44)                           >", as.IsInt(44))
	fmt.Println("INT: (true)                         >", as.IsInt(true))
	fmt.Println("INT: (44.44)                        >", as.IsInt(44.44))
	fmt.Println("INT: ('test')                       >", as.IsInt("test"))
	fmt.Println("INT: ('14:14:14')                   >", as.IsInt(as.ToTime(false, "14:14:14")))
}

func asisbool() {
	fmt.Println("")
	fmt.Println("AS IS BOOL *******************************************************")
	fmt.Println("BOOL: (44)                           >", as.IsBool(44))
	fmt.Println("BOOL: (true)                         >", as.IsBool(true))
	fmt.Println("BOOL: (44.44)                        >", as.IsBool(44.44))
	fmt.Println("BOOL: ('test')                       >", as.IsBool("test"))
	fmt.Println("BOOL: ('14:14:14')                   >", as.IsBool(as.ToTime(false, "14:14:14")))
}

func asisfloat() {
	fmt.Println("")
	fmt.Println("AS IS FLOAT *******************************************************")
	fmt.Println("FLOAT: (44)                           >", as.IsFloat(44))
	fmt.Println("FLOAT: (true)                         >", as.IsFloat(true))
	fmt.Println("FLOAT: (44.44)                        >", as.IsFloat(44.44))
	fmt.Println("FLOAT: ('test')                       >", as.IsFloat("test"))
	fmt.Println("FLOAT: ('14:14:14')                   >", as.IsFloat(as.ToTime(false, "14:14:14")))
}

func asisstring() {
	fmt.Println("")
	fmt.Println("AS IS STRING *******************************************************")
	fmt.Println("STRING: (44)                         >", as.IsString(44))
	fmt.Println("STRING: (true)                       >", as.IsString(true))
	fmt.Println("STRING: (44.44)                      >", as.IsString(44.44))
	fmt.Println("STRING: ('test')                     >", as.IsString("test"))
	fmt.Println("STRING: ('14:14:14')                 >", as.IsString(as.ToTime(false, "14:14:14")))
}

func asistime() {
	fmt.Println("")
	fmt.Println("AS IS TIME *******************************************************")
	fmt.Println("TIME: (44)                           >", as.IsTime(44))
	fmt.Println("TIME: (true)                         >", as.IsTime(true))
	fmt.Println("TIME: (44.44)                        >", as.IsTime(44.44))
	fmt.Println("TIME: ('test')                       >", as.IsTime("test"))
	fmt.Println("TIME: ('14:14:14')                   >", as.IsTime(as.ToTime(false, "14:14:14")))
}

func asisnillable() {
	fmt.Println("")
	fmt.Println("AS IS NILLABLE *******************************************************")
	fmt.Println("NILLABLE: ('')                       >", as.IsNillable(""))
	fmt.Println("NILLABLE: ([]string{})               >", as.IsNillable([]string{}))
}

func astoformattedbytes() {
	fmt.Println("")
	fmt.Println("AS TO FORMATTED BYTES *******************************************************")
	fmt.Println("FORMAT: (44)                       >", as.FormatIntToByte(44))
	fmt.Println("FORMAT: (444)                      >", as.FormatIntToByte(444))
	fmt.Println("FORMAT: (4444)                     >", as.FormatIntToByte(4444))
	fmt.Println("FORMAT: (44444)                    >", as.FormatIntToByte(44444))
	fmt.Println("FORMAT: (444444)                   >", as.FormatIntToByte(444444))
	fmt.Println("FORMAT: (4444444)                  >", as.FormatIntToByte(4444444))
	fmt.Println("FORMAT: (44444444)                 >", as.FormatIntToByte(44444444))
	fmt.Println("FORMAT: (444444444)                >", as.FormatIntToByte(444444444))
	fmt.Println("FORMAT: (4444444444)               >", as.FormatIntToByte(4444444444))
	fmt.Println("FORMAT: (44444444444)              >", as.FormatIntToByte(44444444444))
	fmt.Println("FORMAT: (444444444444)             >", as.FormatIntToByte(444444444444))
	fmt.Println("FORMAT: (4444444444444)            >", as.FormatIntToByte(4444444444444))
	fmt.Println("FORMAT: (44444444444444)           >", as.FormatIntToByte(44444444444444))
	fmt.Println("FORMAT: (444444444444444)          >", as.FormatIntToByte(444444444444444))
	fmt.Println("FORMAT: (4444444444444444)         >", as.FormatIntToByte(4444444444444444))
	fmt.Println("FORMAT: (44444444444444444)        >", as.FormatIntToByte(44444444444444444))
	fmt.Println("FORMAT: (444444444444444444)       >", as.FormatIntToByte(444444444444444444))
	fmt.Println("FORMAT: (999999999999999999)       >", as.FormatIntToByte(999999999999999999))
	fmt.Println("FORMAT: (1000000000000000000)      >", as.FormatIntToByte(1152921504606846976))
}

func asfuncs() {
	time.Sleep(1 * time.Second)
	astostring()
	time.Sleep(1 * time.Second)
	astrimmed()
	time.Sleep(1 * time.Second)
	astofloat()
	time.Sleep(1 * time.Second)
	astorunelength()
	time.Sleep(1 * time.Second)
	astobool()
	time.Sleep(1 * time.Second)
	astobytes()
	time.Sleep(1 * time.Second)
	astoslice()
	time.Sleep(1 * time.Second)
	astoint()
	time.Sleep(1 * time.Second)
	astoip()
	time.Sleep(1 * time.Second)
	astobase64()
	time.Sleep(1 * time.Second)
	asfrombase64()
	time.Sleep(1 * time.Second)
	asisempty()
	time.Sleep(1 * time.Second)
	asiskind()
	time.Sleep(1 * time.Second)
	asofkind()
	time.Sleep(1 * time.Second)
	asoftype()
	time.Sleep(1 * time.Second)
	astotime()
	time.Sleep(1 * time.Second)
	astoduration()
	time.Sleep(1 * time.Second)
	astofixedlengthafter()
	time.Sleep(1 * time.Second)
	astofixedlengthbefore()
	time.Sleep(1 * time.Second)
	astofixedlengthcenter()
	time.Sleep(1 * time.Second)
	asisint()
	time.Sleep(1 * time.Second)
	asisbool()
	time.Sleep(1 * time.Second)
	asisfloat()
	time.Sleep(1 * time.Second)
	asisstring()
	time.Sleep(1 * time.Second)
	asistime()
	time.Sleep(1 * time.Second)
	asisnillable()
	time.Sleep(1 * time.Second)
	astoformattedbytes()
}

func terminalsize() {
	fmt.Println("")
	fmt.Println("TERMINAL DISPLAY **********************************************")
	width, height, err := display.GetTerminalSize()
	fmt.Println("TERMINAL SIZE                      >", width, height, err)
}

func lines() {
	fmt.Println("")
	fmt.Println("TERMINAL LINES **********************************************")
	fmt.Println("TERMINAL THINK LINE                >", display.ThickLine(40))
	fmt.Println("TERMINAL THIN LINE                 >", display.ThinLine(40))
	fmt.Println("TERMINAL SPECIAL LINE              >", display.Specialine(40))
	fmt.Println("TERMINAL SPACE LINE                >", display.SpaceLine(40))
}

func displayfuncs() {
	time.Sleep(1 * time.Second)
	terminalsize()
	time.Sleep(1 * time.Second)
	lines()
}

func coloureffects() {
	fmt.Println("")
	fmt.Println("COLOUR EFFECTS **********************************************")
	fmt.Println("COLOURS BOLD                       >", colours.Bold("BOLD"))
	fmt.Println("COLOURS ITALIC                     >", colours.Italic("ITALIC"))
	fmt.Println("COLOURS UNDERLINE                  >", colours.Underline("UNDERLINE"))
	fmt.Println("COLOURS STRIKETHROUGH              >", colours.StrikeThrough("STRIKETHROUGH"))
	fmt.Println("COLOURS BLINK                      >", colours.Blink("BLINK"))
	fmt.Println("COLOURS REVERSED                   >", colours.Reversed("REVERSED"))
}

func colourcolours() {
	fmt.Println("")
	fmt.Println("COLOURS COLORS **********************************************")
	fmt.Println("COLOURS BACKGROUND                 >", colours.Background("BACKGROUND", colours.RED))
	fmt.Println("COLOURS YELLOW                     >", colours.Color("COLOR", colours.YELLOW))
	fmt.Println("COLOURS HIGHLIGHT                  >", colours.Highlight("HIGHLIGHT", "GHLI", colours.BRIGHTYELLOW))
	fmt.Println("COLOURS BOLD YELLO                 >", colours.Color(colours.Bold("BOLD YELLOW"), colours.YELLOW))
	fmt.Println("COLOURS BACKGROOUND                >", colours.Background(colours.Color(colours.Bold("BOLD YELLOW WITH RED BACKGROUND"), colours.YELLOW), colours.RED))
	fmt.Println("COLOURS BLACK                      >", colours.Color("BLACK", colours.BLACK))
	fmt.Println("COLOURS BRIGHTBLACK                >", colours.Color("BRIGHTBLACK", colours.BRIGHTBLACK))
	fmt.Println("COLOURS RED                        >", colours.Color("RED", colours.RED))
	fmt.Println("COLOURS BRIGHTRED                  >", colours.Color("BRIGHTRED", colours.BRIGHTRED))
	fmt.Println("COLOURS GREEN                      >", colours.Color("GREEN", colours.GREEN))
	fmt.Println("COLOURS BRIGHTGREEN                >", colours.Color("BRIGHTGREEN", colours.BRIGHTGREEN))
	fmt.Println("COLOURS YELLOW                     >", colours.Color("YELLOW", colours.YELLOW))
	fmt.Println("COLOURS BRIGHTYELLOW               >", colours.Color("BRIGHTYELLOW", colours.BRIGHTYELLOW))
	fmt.Println("COLOURS BLUE                       >", colours.Color("BLUE", colours.BLUE))
	fmt.Println("COLOURS BRIGHTBLUE                 >", colours.Color("BRIGHTBLUE", colours.BRIGHTBLUE))
	fmt.Println("COLOURS MAGENTA                    >", colours.Color("MAGENTA", colours.MAGENTA))
	fmt.Println("COLOURS BRIGHTMAGENTA              >", colours.Color("BRIGHTMAGENTA", colours.BRIGHTMAGENTA))
	fmt.Println("COLOURS WHITE                      >", colours.Color("WHITE", colours.WHITE))
	fmt.Println("COLOURS BRIGHTWHITE                >", colours.Color("BRIGHTWHITE", colours.BRIGHTWHITE))
}

func colourpanels() {
	fmt.Println("")
	fmt.Println("COLOUR PANELS ********************************************")
	fmt.Println("COLOURS BLACK SMALL PANEL          >", colours.BlackSmallPanel("Here is some text in a black panel."))
	fmt.Println("COLOURS RED SMALL PANEL            >", colours.RedSmallPanel("Here is some text in a red panel."))
	fmt.Println("COLOURS GREEN SMALL PANEL          >", colours.GreenSmallPanel("Here is some text in a green panel."))
	fmt.Println("COLOURS YELLOW SMALL PANEL         >", colours.YellowSmallPanel("Here is some text in a yellow panel."))
	fmt.Println("COLOURS BLUE SMALL PANEL           >", colours.BlueSmallPanel("Here is some text in a blue panel."))
	fmt.Println("COLOURS MAGENTA SMALL PANEL        >", colours.MagentaSmallPanel("Here is some text in a magenta panel."))
	fmt.Println("COLOURS CYAN SMALL PANEL           >", colours.CyanSmallPanel("Here is some text in a cyan panel."))
	fmt.Println("COLOURS WHITE SMALL PANEL          >", colours.WhiteSmallPanel("Here is some text in a white panel."))
	fmt.Println("COLOURS BLACK PANEL                >", colours.BlackPanel("Here is some text in a black panel."))
	fmt.Println("COLOURS RED PANEL                  >", colours.RedPanel("Here is some text in a red panel."))
	fmt.Println("COLOURS GREEN PANEL                >", colours.GreenPanel("Here is some text in a green panel."))
	fmt.Println("COLOURS YELLOW PANEL               >", colours.YellowPanel("Here is some text in a yellow panel."))
	fmt.Println("COLOURS BLUE PANEL                 >", colours.BluePanel("Here is some text in a blue panel."))
	fmt.Println("COLOURS MAGENTA PANEL              >", colours.MagentaPanel("Here is some text in a magenta panel."))
	fmt.Println("COLOURS CYAN PANEL                 >", colours.CyanPanel("Here is some text in a cyan panel."))
	fmt.Println("COLOURS WHITE PANEL                >", colours.WhitePanel("Here is some text in a white panel."))
}

func colourcommands() {
	fmt.Println("")
	fmt.Println("COLOUR COMMANDS ********************************************")
	fmt.Println("COLOURS TITLE                      >", colours.Title("Title"))
	fmt.Println("COLOURS CUSTOM TITLE               >", colours.CustomTitle("Title", colours.BRIGHTWHITE, colours.BRIGHTYELLOW))
	fmt.Println("COLOURS INFO                       >", colours.Info("This is an info message"))
	fmt.Println("COLOURS SUCCESS                    >", colours.Success("This is a success message"))
	fmt.Println("COLOURS WARNING                    >", colours.Warning("This is an warning message"))
	fmt.Println("COLOURS ERROR                      >", colours.Error("This is an error message"))
}

func colouricons() {
	fmt.Println("")
	fmt.Println("COLOUR ASCII ICONS **********************************************")
	fmt.Println("COLOURS ICON TICK                  >", colours.TICK, colours.Green(colours.TICK), colours.BrightGreen(colours.TICK))
	fmt.Println("COLOURS ICON CROSS                 >", colours.CROSS, colours.Red(colours.CROSS), colours.BrightRed(colours.CROSS))
	fmt.Println("COLOURS ICON COPYRIGHT             >", colours.COPYRIGHT)
	fmt.Println("COLOURS ICON REGISTREDTM           >", colours.REGISTREDTM)
	fmt.Println("COLOURS ICON TRADEMARK             >", colours.TRADEMARK)
	fmt.Println("COLOURS ICON BULLET                >", colours.BULLET)
	fmt.Println("COLOURS ICON ARROWLEFT             >", colours.ARROWLEFT)
	fmt.Println("COLOURS ICON ARROWRIGHT            >", colours.ARROWRIGHT)
	fmt.Println("COLOURS ICON ARROWUP               >", colours.ARROWUP)
	fmt.Println("COLOURS ICON ARROWDOWN             >", colours.ARROWDOWN)
	fmt.Println("COLOURS ICON ARROWLEFTRIGHT        >", colours.ARROWLEFTRIGHT)
	fmt.Println("COLOURS ICON INFINITY              >", colours.INFINITY)
	fmt.Println("COLOURS ICON CELSIUS               >", colours.CELSIUS)
	fmt.Println("COLOURS ICON FAHRENHEIT            >", colours.FAHRENHEIT)
	fmt.Println("COLOURS ICON SUNSHINE              >", colours.SUNSHINE)
	fmt.Println("COLOURS ICON CLOUDY                >", colours.CLOUDY)
	fmt.Println("COLOURS ICON RAIN                  >", colours.RAIN)
	fmt.Println("COLOURS ICON SNOW                  >", colours.SNOW)
	fmt.Println("COLOURS ICON STAR BLACK            >", colours.STARBLACK)
	fmt.Println("COLOURS ICON STAR WHITE            >", colours.STARWHITE)
	fmt.Println("COLOURS ICON PHONE BLACK           >", colours.PHONEBLACK)
	fmt.Println("COLOURS ICON PHONE WHITE           >", colours.PHONEWHITE)
	fmt.Println("COLOURS ICON POINT LEFT            >", colours.POINTLEFT)
	fmt.Println("COLOURS ICON POINT RIGHT           >", colours.POINTRIGHT)
	fmt.Println("COLOURS ICON POINT UP              >", colours.POINTUP)
	fmt.Println("COLOURS ICON POINT DOWN            >", colours.POINTDOWN)
	fmt.Println("COLOURS ICON DEATH                 >", colours.DEATH)
	fmt.Println("COLOURS ICON SMILEY                >", colours.SMILEY)
	fmt.Println("COLOURS ICON HEART                 >", colours.HEART)
	fmt.Println("COLOURS ICON DIAMOND               >", colours.DIAMOND)
	fmt.Println("COLOURS ICON SPADE                 >", colours.SPADE)
	fmt.Println("COLOURS ICON CLUB                  >", colours.CLUB)
}

func colourfuncs() {
	time.Sleep(1 * time.Second)
	coloureffects()
	time.Sleep(1 * time.Second)
	colourcolours()
	time.Sleep(1 * time.Second)
	colourpanels()
	time.Sleep(1 * time.Second)
	colourcommands()
	time.Sleep(1 * time.Second)
	colouricons()
}

func ipfilters() {
	fmt.Println("")
	fmt.Println("SYSTEM IP FILTER **********************************************")
	options := ipfilter.Options{
		AllowedIPs:     []string{"10.52.208.1", "10.0.0.0/16"},
		BlockedIPs:     []string{},
		BlockByDefault: true,
	}
	filter := ipfilter.New(options)

	fmt.Println("IP FILTER - 10.0.0.1               >", filter.Allowed("10.0.0.1"))
	fmt.Println("IP FILTER - 10.0.42.1              >", filter.Allowed("10.0.42.1"))
	fmt.Println("IP FILTER - 10.42.0.1              >", filter.Allowed("10.42.0.1"))
	fmt.Println("IP FILTER - 10.52.208.1            >", filter.Allowed("10.52.208.1"))
	fmt.Println("IP FILTER - 10.52.208.10           >", filter.Allowed("10.52.208.10"))
}

func securityfuncs() {
	time.Sleep(1 * time.Second)
	ipfilters()
}

func main() {
	clearscreenanddisplaybanner()
	setupsignals()
	setupschedulersandinitiate()
	displaymemoryutilisation()

	isfuncs()
	cronjob2.Stop()
	mathfuncs()
	loopfuncs()
	filepathfuncs()
	textfuncs()
	templatefuncs()
	terminalfuncs()
	sortfuncs()
	jsonfuncs()
	mapfuncs()
	parseversionfuncs()
	dirandfilefuncs()
	timeanddatefuncs()
	systemfuncs()
	asfuncs()
	displayfuncs()
	colourfuncs()
	securityfuncs()

	displaymemoryutilisation()
	displayfinishtimer()
}
