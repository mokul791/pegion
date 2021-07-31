package main

import (
	"database/sql"
	"flag"
	"log"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"swarmtree.com/pegion/pkg/models/mysql"
)

type application struct {
	erroLog *log.Logger
	infoLog *log.Logger
	pegions *mysql.PegionModel
}

func main() {

	// Define a new command-line flage with the name 'addr', a default value of '4000'
	// and some short help text explaining what the flage controls. The value of
	// flag will be stored in the addr variable at runtime.
	// func String(name, value, usage) *string
	// Parse flag.
	addr := flag.String("addr", ":4000", "HTTP network address")
	dsn := flag.String("dsn", "web:pass@/pegionbox?parseTime=true", "MySQL database source name")
	flag.Parse()

	//  Logging to an external file: go run cmd/web/* >>/tmp/info.log 2>>/tmp/error.log
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	// Connection to db
	db, err := openDB(*dsn)
	if err != nil {
		errorLog.Fatal(err)
	}

	defer db.Close()

	app := &application{
		erroLog: errorLog,
		infoLog: infoLog,
		pegions: &mysql.PegionModel{DB: db},
	}

	srv := &http.Server{
		Addr:     *addr,
		ErrorLog: errorLog,
		Handler:  app.routes(),
	}

	infoLog.Printf("Starting server on %s", *addr)
	// http.ListenAndServe() function starts a new web server.
	err = srv.ListenAndServe()
	errorLog.Fatal(err)

}

func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}
