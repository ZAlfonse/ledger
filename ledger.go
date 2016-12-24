package ledger

import (
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/zalfonse/lumber"
)

type Line struct {
	gorm.Model
	Action    string
	Quantity  float32
	Price     float32
	Ticker    string
	Placed    time.Time
	Completed time.Time
	Cancelled time.Time
}

type Ledger struct {
	name   string
	logger lumber.Logger
	gorm.DB
}

func (l *Ledger) AddTransaction(action string, quantity float32, price float32, ticker string) Line {
	line := Line{
		Action:   action,
		Quantity: quantity,
		Price:    price,
		Ticker:   ticker,
	}
	l.DB.Create(&line)
	return line
}

func (l *Ledger) SummarizePosition(ticker string) {
	var history []Line
	l.DB.Where("ticker = ?", ticker).Find(&history)
	for _, line := range history {
		l.logger.Info(fmt.Sprintf("%v order of %f shares of %v priced @ %f executed on %v.", line.Action, line.Quantity, line.Ticker, line.Price, line.Placed.Format(time.RFC822Z)))
	}
}

func (l *Ledger) Close() {
	l.DB.Close()
}

func OpenLedger(name string) *Ledger {
	db, err := gorm.Open("sqlite3", fmt.Sprintf("%v.db", name))
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&Line{})

	logger := lumber.NewLogger(lumber.TRACE)

	return &Ledger{
		name,
		*logger,
		*db,
	}
}
