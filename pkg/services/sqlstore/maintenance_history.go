package sqlstore

import (
	"github.com/grafana/grafana/pkg/bus"
	m "github.com/grafana/grafana/pkg/models"

	"github.com/go-xorm/xorm"
)

func init()  {

	bus.AddHandler("sql",GetMaintenanceHistory)
	bus.AddHandler("sql",GetMaintenanceHistoryByInterval)
	bus.AddHandler("sql",UpdateMaintenanceHistory)

}

func GetMaintenanceHistory(query *m.GetMaintenanceHistoryQuery) error {

	query.Result = make([]*m.MaintenanceHistoryDTO, 0)
	sess := x.Table("maintenance_history")
	sess.Where("maintenance_history.org=?", query.Org)
	sess.Cols("maintenance_history.id","maintenance_history.org","maintenance_history.message","maintenance_history.sended","maintenance_history.interval")

	err := sess.Find(&query.Result)
	return err
}

func GetMaintenanceHistoryByInterval(query *m.GetMaintenanceHistoryByIntervalQuery) error {

	query.Result = make([]*m.MaintenanceHistoryDTO, 0)
	sess := x.Table("maintenance_history")
	sess.Where("maintenance_history.interval=?", query.Interval)
	sess.Cols("maintenance_history.id","maintenance_history.org","maintenance_history.status","maintenance_history.message","maintenance_history.sended","maintenance_history.interval")

	err := sess.Find(&query.Result)
	return err
}

func UpdateMaintenanceHistory(cmd *m.UpdateMaintenanceHistory) error {
	return inTransaction(func(sess *xorm.Session) error {
		var stutus="Done"
		var rawSql = "UPDATE maintenance_history SET  status=? WHERE Id=?"
		_, err := sess.Exec(rawSql,stutus,cmd.Id)
		if err != nil {
			return err
		}

		return validateOneAdminLeftInOrg(cmd.Org, sess)
	})
}