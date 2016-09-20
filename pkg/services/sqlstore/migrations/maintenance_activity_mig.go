package migrations

import  ."github.com/grafana/grafana/pkg/services/sqlstore/migrator"

func addMaintenanceActivityMigrations(mg *Migrator) {

  maintenance_activityV1 := Table{
    Name: "maintenance_activity",
    Columns: []*Column{
      {Name: "id", Type: DB_BigInt, Length: 20, IsPrimaryKey: true, IsAutoIncrement: true},
      {Name: "name", Type: DB_NVarchar, Length: 255, Nullable: false},
      {Name: "org_id", Type: DB_BigInt, Length: 20, Nullable: false},
      {Name: "message", Type: DB_NVarchar, Length: 255, Nullable: false},
      {Name: "updated", Type: DB_DateTime, Nullable: false},
    },
    Indices: []*Index{
      {Cols: []string{"id"}, Type: IndexType},
      {Cols: []string{"org_id"}, Type: IndexType},
      {Cols: []string{"name"}, Type: IndexType},
      {Cols: []string{"message"}, Type: IndexType},
    },

  }
  mg.AddMigration("create process  table v1-7", NewAddTableMigration(maintenance_activityV1))
  addTableIndicesMigrations(mg, "v1-7", maintenance_activityV1)

}
