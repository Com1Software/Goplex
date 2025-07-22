package main

import (
	"net"
	"net/http"
	"os/exec"
	"runtime"
	"strconv"

	"fmt"
	"os"

	"github.com/Com1Software/go-dbase/dbase"
	"golang.org/x/text/encoding/charmap"
)

type Apps struct {
	App string `dbase:"APP"`
}

// ----------------------------------------------------------------
func main() {
	fmt.Println("Goplex Server")
	fmt.Printf("Operating System : %s\n", runtime.GOOS)
	xip := fmt.Sprintf("%s", GetOutboundIP())
	port := "8080"
	switch {
	//-------------------------------------------------------------
	case len(os.Args) == 2:

		fmt.Println("Not")

		//-------------------------------------------------------------
	default:

		fmt.Println("Server running....")
		fmt.Println("Listening on " + xip + ":" + port)

		fmt.Println("")
		http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			xdata := InitPage(xip)
			fmt.Fprint(w, xdata)
		})
		//------------------------------------------------ About Page Handler
		http.HandleFunc("/about", func(w http.ResponseWriter, r *http.Request) {
			xdata := AboutPage(xip)
			fmt.Fprint(w, xdata)
		})
		//------------------------------------------------ App Table Display Page Handler
		http.HandleFunc("/apptabledisplay", func(w http.ResponseWriter, r *http.Request) {
			xdata := DisplayPage(xip, port)
			fmt.Fprint(w, xdata)

		})
		http.HandleFunc("/addapplication", func(w http.ResponseWriter, r *http.Request) {
			app := r.FormValue("app")
			fmt.Println(app)
			if len(app) > 0 {
				table, err := dbase.OpenTable(&dbase.Config{
					Filename:   "APPS.DBF",
					TrimSpaces: true,
					WriteLock:  true,
				})
				if err != nil {
					panic(err)
				}
				defer table.Close()
				row, err := table.Row()
				if err != nil {
					panic(err)
				}
				p := Apps{
					App: app,
				}

				row, err = table.RowFromStruct(p)
				if err != nil {
					panic(err)
				}
				//			fmt.Println(row)
				err = row.FieldByName("APP").SetValue(app)
				if err != nil {
					panic(err)
				}
				err = row.Write()
				if err != nil {
					panic(err)
				}
			}
			xdata := DisplayPage(xip, port)
			fmt.Fprint(w, xdata)

		})

		//------------------------------------------------- Start Server
		TableCheck()
		Openbrowser(xip + ":" + port)
		if err := http.ListenAndServe(xip+":"+port, nil); err != nil {
			panic(err)
		}
	}
}

// Openbrowser : Opens default web browser to specified url
func Openbrowser(url string) error {
	var cmd string
	var args []string
	switch runtime.GOOS {
	case "windows":
		cmd = "cmd"
		args = []string{"/c", "start msedge"}

	case "linux":
		cmd = "chromium-browser"
		args = []string{""}

	case "darwin":
		cmd = "open"
	default:
		cmd = "xdg-open"
	}
	args = append(args, url)
	return exec.Command(cmd, args...).Start()
}

func TableCheck() {
	tt := "APPS.DBF"

	if _, err := os.Stat(tt); err == nil {

	} else {

		file, err := dbase.NewTable(
			dbase.FoxProAutoincrement,
			&dbase.Config{
				Filename:   tt,
				Converter:  dbase.NewDefaultConverter(charmap.Windows1250),
				TrimSpaces: true,
			},
			tcolumns(),
			64,
			nil,
		)
		if err != nil {
			panic(err)
		}
		defer file.Close()

		row, err := file.RowFromStruct(&Apps{
			App: "APP",
		})
		if err != nil {
			panic(err)
		}

		err = row.Add()
		if err != nil {
			panic(err)
		}
		fmt.Printf(
			"Last modified: %v Columns count: %v Record count: %v File size: %v \n",
			file.Header().Modified(0),
			file.Header().ColumnsCount(),
			file.Header().RecordsCount(),
			file.Header().FileSize(),
		)

	}

}

func tcolumns() []*dbase.Column {

	appCol, err := dbase.NewColumn("APP", dbase.Varchar, 80, 0, false)
	if err != nil {
		panic(err)
	}
	return []*dbase.Column{
		appCol,
	}
}

func vcolumns() []*dbase.Column {

	appCol, err := dbase.NewColumn("App", dbase.Varchar, 80, 0, false)
	userCol, err := dbase.NewColumn("User", dbase.Varchar, 254, 0, false)
	pwCol, err := dbase.NewColumn("Password", dbase.Varchar, 254, 0, false)

	if err != nil {
		panic(err)
	}
	return []*dbase.Column{
		appCol,
		userCol,
		pwCol,
	}
}

func DateTimeDisplay(xdata string) string {
	//------------------------------------------------------------------------
	xdata = xdata + "<script>"
	xdata = xdata + "function startTime() {"
	xdata = xdata + "  var today = new Date();"
	xdata = xdata + "  var d = today.getDay();"
	xdata = xdata + "  var h = today.getHours();"
	xdata = xdata + "  var m = today.getMinutes();"
	xdata = xdata + "  var s = today.getSeconds();"
	xdata = xdata + "  var ampm = h >= 12 ? 'pm' : 'am';"
	xdata = xdata + "  var mo = today.getMonth();"
	xdata = xdata + "  var dm = today.getDate();"
	xdata = xdata + "  var yr = today.getFullYear();"
	xdata = xdata + "  m = checkTimeMS(m);"
	xdata = xdata + "  s = checkTimeMS(s);"
	xdata = xdata + "  h = checkTimeH(h);"
	//------------------------------------------------------------------------
	xdata = xdata + "  switch (d) {"
	xdata = xdata + "    case 0:"
	xdata = xdata + "       day = 'Sunday';"
	xdata = xdata + "    break;"
	xdata = xdata + "    case 1:"
	xdata = xdata + "    day = 'Monday';"
	xdata = xdata + "        break;"
	xdata = xdata + "    case 2:"
	xdata = xdata + "        day = 'Tuesday';"
	xdata = xdata + "        break;"
	xdata = xdata + "    case 3:"
	xdata = xdata + "        day = 'Wednesday';"
	xdata = xdata + "        break;"
	xdata = xdata + "    case 4:"
	xdata = xdata + "        day = 'Thursday';"
	xdata = xdata + "        break;"
	xdata = xdata + "    case 5:"
	xdata = xdata + "        day = 'Friday';"
	xdata = xdata + "        break;"
	xdata = xdata + "    case 6:"
	xdata = xdata + "        day = 'Saturday';"
	xdata = xdata + "}"
	//------------------------------------------------------------------------------------
	xdata = xdata + "  switch (mo) {"
	xdata = xdata + "    case 0:"
	xdata = xdata + "       month = 'January';"
	xdata = xdata + "       break;"
	xdata = xdata + "    case 1:"
	xdata = xdata + "       month = 'Febuary';"
	xdata = xdata + "       break;"
	xdata = xdata + "    case 2:"
	xdata = xdata + "       month = 'March';"
	xdata = xdata + "       break;"
	xdata = xdata + "    case 3:"
	xdata = xdata + "       month = 'April';"
	xdata = xdata + "       break;"
	xdata = xdata + "    case 4:"
	xdata = xdata + "       month = 'May';"
	xdata = xdata + "       break;"
	xdata = xdata + "    case 5:"
	xdata = xdata + "       month = 'June';"
	xdata = xdata + "       break;"
	xdata = xdata + "    case 6:"
	xdata = xdata + "       month = 'July';"
	xdata = xdata + "       break;"
	xdata = xdata + "    case 7:"
	xdata = xdata + "       month = 'August';"
	xdata = xdata + "       break;"
	xdata = xdata + "    case 8:"
	xdata = xdata + "       month = 'September';"
	xdata = xdata + "       break;"
	xdata = xdata + "    case 9:"
	xdata = xdata + "       month = 'October';"
	xdata = xdata + "       break;"
	xdata = xdata + "    case 10:"
	xdata = xdata + "       month = 'November';"
	xdata = xdata + "       break;"
	xdata = xdata + "    case 11:"
	xdata = xdata + "       month = 'December';"
	xdata = xdata + "       break;"
	xdata = xdata + "}"
	//  -------------------------------------------------------------------
	xdata = xdata + "  document.getElementById('txtdt').innerHTML = day+', '+month+' '+dm+', '+yr+' - '+h + ':' + m + ':' + s+' '+ampm;"

	xdata = xdata + "  var t = setTimeout(startTime, 500);"
	xdata = xdata + "}"
	//----------
	xdata = xdata + "function checkTimeMS(i) {"
	xdata = xdata + "  if (i < 10) {i = '0' + i};"
	xdata = xdata + "  return i;"
	xdata = xdata + "}"
	//----------
	xdata = xdata + "function checkTimeH(i) {"
	xdata = xdata + "  if (i > 12) {i = i -12};"
	xdata = xdata + "  return i;"
	xdata = xdata + "}"
	xdata = xdata + "</script>"
	return xdata

}
func LoopDisplay(xdata string) string {
	//------------------------------------------------------------------------
	xdata = xdata + "<script>"
	xdata = xdata + "function startLoop() {"
	//  -------------------------------------------------------------------
	xdata = xdata + "  document.getElementById('txtloop').innerHTML = Math.random();"
	xdata = xdata + "  var t = setTimeout(startLoop, 500);"
	xdata = xdata + "}"
	xdata = xdata + "</script>"
	return xdata

}

func GetOutboundIP() net.IP {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		fmt.Println(err)
	}
	defer conn.Close()

	localAddr := conn.LocalAddr().(*net.UDPAddr)

	return localAddr.IP
}

func InitPage(xip string) string {
	//---------------------------------------------------------------------------
	//----------------------------------------------------------------------------
	xxip := ""
	xdata := "<!DOCTYPE html>"
	xdata = xdata + "<html>"
	xdata = xdata + "<head>"
	//------------------------------------------------------------------------
	xdata = xdata + "<title>Goplex Server</title>"
	xdata = DateTimeDisplay(xdata)
	//------------------------------------------------------------------------
	xdata = xdata + "</head>"
	//------------------------------------------------------------------------

	xdata = xdata + "<body>"
	xdata = xdata + "<center>"
	xdata = xdata + "<H1>Goplex Server</H1>"
	//---------
	host, _ := os.Hostname()
	addrs, _ := net.LookupIP(host)
	for _, addr := range addrs {
		if ipv4 := addr.To4(); ipv4 != nil {
			xxip = fmt.Sprintf("%s", ipv4)
		}
	}
	xdata = xdata + "<body onload='startTime()'>"
	xdata = xdata + "<div id='txtdt'></div>"

	xdata = xdata + "Host Port IP : " + xip
	xdata = xdata + "<BR> Machine IP : " + xxip + "</p>"

	xdata = xdata + "  <A HREF='http://" + xip + ":8080/about'> [ About ] </A>  "
	xdata = xdata + "  <A HREF='http://" + xip + ":8080/apptabledisplay'> [ Aplication Table ] </A>  "
	xdata = xdata + "<BR><BR>Goplex Server"

	//------------------------------------------------------------------------

	xdata = xdata + "</center>"
	xdata = xdata + " </body>"
	xdata = xdata + " </html>"
	return xdata
}

// ----------------------------------------------------------------
func AboutPage(xip string) string {
	//----------------------------------------------------------------------------
	xdata := "<!DOCTYPE html>"
	xdata = xdata + "<html>"
	xdata = xdata + "<head>"
	//------------------------------------------------------------------------
	xdata = xdata + "<title>About Page</title>"
	xdata = LoopDisplay(xdata)
	//------------------------------------------------------------------------
	xdata = DateTimeDisplay(xdata)
	xdata = xdata + "<style>"
	xdata = xdata + "body {"
	xdata = xdata + "    background-color: white;"
	xdata = xdata + "}"
	xdata = xdata + "	h1 {"
	xdata = xdata + "	color: white;"
	xdata = xdata + "	text-align: center;"
	xdata = xdata + "}"
	xdata = xdata + "	p {"
	xdata = xdata + "font-family: verdana;"
	xdata = xdata + "	font-size: 20px;"
	xdata = xdata + "}"
	xdata = xdata + "</style>"
	xdata = xdata + "</head>"
	//------------------------------------------------------------------------
	xdata = xdata + "<body onload='startTime()'>"
	xdata = xdata + "<p>Goplex Server</p>"
	xdata = xdata + "<div id='txtdt'></div>"
	//---------
	xdata = xdata + "<BR><BR>"
	xdata = xdata + "  <A HREF='https://github.com/Com1Software/Goplex'> [ Goplex Server GitHub Repository ] </A>  "
	xdata = xdata + "<BR><BR>"
	//------------------------------------------------------------------------
	xdata = xdata + "  <A HREF='http://" + xip + ":8080'> [ Return to Start Page ] </A>  "
	xdata = xdata + "<BR><BR>"

	xdata = xdata + "Goplex Server"
	//------------------------------------------------------------------------

	//------------------------------------------------------------------------
	xdata = xdata + " </body>"
	xdata = xdata + " </html>"
	return xdata

}

// ----------------------------------------------------------------
func DisplayPage(xip string, port string) string {
	//----------------------------------------------------------------------------
	xdata := "<!DOCTYPE html>"
	xdata = xdata + "<html>"
	xdata = xdata + "<head>"
	xdata = xdata + "<meta name='viewport' content='width=device-width, initial-scale=1'>"
	xdata = xdata + "<style>"
	xdata = xdata + "div.scroll-container {"
	xdata = xdata + "background-color: #333;"
	xdata = xdata + "overflow: auto;"
	xdata = xdata + "white-space: nowrap;"
	xdata = xdata + "padding: 10px;"
	xdata = xdata + "}"
	xdata = xdata + "div.scroll-container img {"
	xdata = xdata + "padding: 10px;"
	xdata = xdata + "}"
	xdata = xdata + "</style>"

	xdata = xdata + "</head>"
	//------------------------------------------------------------------------
	xdata = xdata + "<body onload='startTime()'>"
	xdata = xdata + "<center>"
	xdata = xdata + "<H1>Application Table</H1>"
	xdata = xdata + "<div id='txtdt'></div>"
	xdata = xdata + "<p2>Application Table</p2>"
	xdata = xdata + "<BR>"
	xdata = xdata + "</center>"
	//------------------------------------------------------------------------
	//------------------------------------------------------------------------
	xdata = xdata + "<center>"
	xdata = xdata + "<BR><BR>"
	xdata = xdata + " <A HREF='http://" + xip + ":8080'> [ Return to Start Page ] </A>  "
	//------------------------------------------------------------------------
	xdata = xdata + "<BR><BR>"
	xdata = xdata + "<form action='/addapplication' method='post'>"
	xdata = xdata + "<textarea id='app' name='app' rows='1' cols='20'></textarea>"
	xdata = xdata + "<BR><BR>"
	xdata = xdata + "<input type='submit' value='Add Application'/>"
	xdata = xdata + "</form>"
	xdata = xdata + "<BR><BR>"
	table, err := dbase.OpenTable(&dbase.Config{
		Filename:   "APPS.DBF",
		TrimSpaces: true,
	})
	if err != nil {
		panic(err)
	}
	defer table.Close()
	recno := 0
	for !table.EOF() {
		row, err := table.Next()
		if err != nil {
			panic(err)
		}
		//field := row.FieldByName("tag")
		field := row.Field(0)
		if field == nil {
			panic("Field not found")
		}
		s := fmt.Sprintf("%v", field.GetValue())
		xdata = xdata + "  <A HREF='http://" + xip + ":8080/tagedit?recno=" + strconv.Itoa(recno) + "'> [ " + s + " ] </A>  "
		xdata = xdata + "<BR>"
		recno++

	}
	xdata = xdata + "</center>"
	//------------------------------------------------------------------------
	xdata = xdata + " </body>"
	xdata = xdata + " </html>"
	return xdata

}
