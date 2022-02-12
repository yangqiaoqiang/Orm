package ORM

import (
	"ORM/dialect"
	"ORM/log"
	"ORM/session"
	"database/sql"
)

// 负责数据库交互前准备(连接/测试数据库)，交互后收尾(关闭连接)交给Engine
//Engine是ORM交互入口
type Engine struct {
	db      *sql.DB
	dialect dialect.Dialect
}

func NewEngine(driver, source string) (e *Engine, err error) {
	db, err := sql.Open(driver, source)
	if err != nil {
		log.Error(err)
		return
	}
	//使用ping函数确定databese连接成功
	if err = db.Ping(); err != nil {
		log.Error(err)
		return
	}
	// make sure the specific dialect exists
	dial, ok := dialect.GetDialect(driver)
	if !ok {
		log.Errorf("dialect %s Not Found", driver)
		return
	}
	e = &Engine{db: db, dialect: dial}
	log.Info("connect database success")
	return
}
func (engine *Engine) Close() {
	if err := engine.db.Close(); err != nil {
		log.Error("Failure to close database")
	}
	log.Info("Close database success")
}
func (engine *Engine) NewSession() *session.Session {
	return session.New(engine.db, engine.dialect)
}
