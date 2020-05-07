// Code generated by entc, DO NOT EDIT.

package migrate

import (
	"github.com/facebookincubator/ent/dialect/sql/schema"
	"github.com/facebookincubator/ent/schema/field"
)

var (
	// PlanetsColumns holds the columns for the "planets" table.
	PlanetsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "metal", Type: field.TypeInt64},
		{Name: "metal_prod_level", Type: field.TypeInt},
		{Name: "metal_storage_level", Type: field.TypeInt},
		{Name: "hydrogen", Type: field.TypeInt64},
		{Name: "hydrogen_prod_level", Type: field.TypeInt},
		{Name: "hydrogen_storage_level", Type: field.TypeInt},
		{Name: "silica", Type: field.TypeInt64},
		{Name: "silica_prod_level", Type: field.TypeInt},
		{Name: "silica_storage_level", Type: field.TypeInt},
		{Name: "population", Type: field.TypeInt64},
		{Name: "population_prod_level", Type: field.TypeInt},
		{Name: "population_storage_level", Type: field.TypeInt},
		{Name: "solar_prod_level", Type: field.TypeInt},
		{Name: "ship_factory_level", Type: field.TypeInt},
		{Name: "research_center_level", Type: field.TypeInt},
		{Name: "region_code", Type: field.TypeInt},
		{Name: "system_code", Type: field.TypeInt},
		{Name: "orbit_code", Type: field.TypeInt},
		{Name: "suborbit_code", Type: field.TypeInt},
		{Name: "position_code", Type: field.TypeInt, Unique: true},
		{Name: "name", Type: field.TypeString, Default: "Unnamed"},
		{Name: "planet_type", Type: field.TypeEnum, Enums: []string{"habitable", "mineral", "gas_giant", "ice_giant"}},
		{Name: "planet_skin", Type: field.TypeString},
		{Name: "last_resource_update", Type: field.TypeTime},
		{Name: "user_planets", Type: field.TypeInt, Nullable: true},
	}
	// PlanetsTable holds the schema information for the "planets" table.
	PlanetsTable = &schema.Table{
		Name:       "planets",
		Columns:    PlanetsColumns,
		PrimaryKey: []*schema.Column{PlanetsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:  "planets_users_planets",
				Columns: []*schema.Column{PlanetsColumns[27]},

				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// SessionsColumns holds the columns for the "sessions" table.
	SessionsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "token", Type: field.TypeString, Unique: true},
		{Name: "data", Type: field.TypeBytes},
		{Name: "expiry", Type: field.TypeTime},
	}
	// SessionsTable holds the schema information for the "sessions" table.
	SessionsTable = &schema.Table{
		Name:        "sessions",
		Columns:     SessionsColumns,
		PrimaryKey:  []*schema.Column{SessionsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// TimersColumns holds the columns for the "timers" table.
	TimersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "action", Type: field.TypeEnum, Enums: []string{"upgrade_metal_prod", "upgrade_hydrogen_prod", "upgrade_silica_prod", "upgrade_solar_prod", "upgrade_urbanism", "upgrade_metal_storage", "upgrade_hydrogen_storage", "upgrade_silica_storage", "upgrade_research_center", "upgrade_ship_factory", "test"}},
		{Name: "group", Type: field.TypeEnum, Enums: []string{"building"}},
		{Name: "end_time", Type: field.TypeTime},
		{Name: "planet_timers", Type: field.TypeInt, Nullable: true},
	}
	// TimersTable holds the schema information for the "timers" table.
	TimersTable = &schema.Table{
		Name:       "timers",
		Columns:    TimersColumns,
		PrimaryKey: []*schema.Column{TimersColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:  "timers_planets_timers",
				Columns: []*schema.Column{TimersColumns[4]},

				RefColumns: []*schema.Column{PlanetsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "timer_group_planet_timers",
				Unique:  true,
				Columns: []*schema.Column{TimersColumns[2], TimersColumns[4]},
			},
		},
	}
	// UsersColumns holds the columns for the "users" table.
	UsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "username", Type: field.TypeString, Unique: true},
		{Name: "email", Type: field.TypeString, Unique: true},
		{Name: "password", Type: field.TypeString},
		{Name: "verify_token", Type: field.TypeString, Nullable: true},
	}
	// UsersTable holds the schema information for the "users" table.
	UsersTable = &schema.Table{
		Name:        "users",
		Columns:     UsersColumns,
		PrimaryKey:  []*schema.Column{UsersColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		PlanetsTable,
		SessionsTable,
		TimersTable,
		UsersTable,
	}
)

func init() {
	PlanetsTable.ForeignKeys[0].RefTable = UsersTable
	TimersTable.ForeignKeys[0].RefTable = PlanetsTable
}
