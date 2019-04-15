package slave

import (
	"fmt"

	"github.com/swapbyt3s/zenit/common/log"
	"github.com/swapbyt3s/zenit/common/mysql"
	"github.com/swapbyt3s/zenit/config"
	"github.com/swapbyt3s/zenit/plugins/inputs"
	"github.com/swapbyt3s/zenit/plugins/lists/metrics"
)

const query = "SHOW SLAVE STATUS"

type MySQLSlave struct{}

func (l *MySQLSlave) Collect() {
	defer func() {
		if err := recover(); err != nil {
			log.Debug(fmt.Sprintf("Plugin - MySQLSlave - Panic (code %d) has been recover from somewhere.\n", err))
		}
	}()

	if !config.File.Inputs.MySQL.Status {
		return
	}

	var a = metrics.Load()
	var m = mysql.GetInstance("mysql")
	m.Connect(config.File.Inputs.MySQL.DSN)

	rows := m.Query(query)

	for column := range rows[0] {
		if value, ok := mysql.ParseValue(rows[0][column]); ok {
			a.Add(metrics.Metric{
				Key:    "zenit_mysql_slave",
				Tags:   []metrics.Tag{{"name", column}},
				Values: value,
			})

			log.Debug(fmt.Sprintf("Plugin - InputMySQLSlave - %s=%d", column, value))
		}
	}
}

func init() {
	inputs.Add("InputMySQLSlave", func() inputs.Input { return &MySQLSlave{} })
}
