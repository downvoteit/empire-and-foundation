// Code generated by entc, DO NOT EDIT.

package migrate

import (
	"github.com/pdeguing/empire-and-foundation/ent/planet"

	"github.com/facebookincubator/ent/dialect/sql/schema"
	"github.com/facebookincubator/ent/schema/field"
)

var (
	// PlanetsColumns holds the columns for the "planets" table.
	PlanetsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "metal", Type: field.TypeInt64, Default: planet.DefaultMetal},
		{Name: "metal_last_update", Type: field.TypeTime},
		{Name: "metal_rate", Type: field.TypeInt, Default: planet.DefaultMetalRate},
		{Name: "metal_prod_level", Type: field.TypeInt, Default: planet.DefaultMetalProdLevel},
		{Name: "metal_storage_level", Type: field.TypeInt, Default: planet.DefaultMetalStorageLevel},
		{Name: "hydrogen", Type: field.TypeInt64, Default: planet.DefaultHydrogen},
		{Name: "hydrogen_last_update", Type: field.TypeTime},
		{Name: "hydrogen_rate", Type: field.TypeInt, Default: planet.DefaultHydrogenRate},
		{Name: "hydrogen_prod_level", Type: field.TypeInt, Default: planet.DefaultHydrogenProdLevel},
		{Name: "hydrogen_storage_level", Type: field.TypeInt, Default: planet.DefaultHydrogenStorageLevel},
		{Name: "silica", Type: field.TypeInt64, Default: planet.DefaultSilica},
		{Name: "silica_last_update", Type: field.TypeTime},
		{Name: "silica_rate", Type: field.TypeInt, Default: planet.DefaultSilicaRate},
		{Name: "silica_prod_level", Type: field.TypeInt, Default: planet.DefaultSilicaProdLevel},
		{Name: "silica_storage_level", Type: field.TypeInt, Default: planet.DefaultSilicaStorageLevel},
		{Name: "population", Type: field.TypeInt64, Default: planet.DefaultPopulation},
		{Name: "population_last_update", Type: field.TypeTime},
		{Name: "population_rate", Type: field.TypeInt, Default: planet.DefaultPopulationRate},
		{Name: "population_prod_level", Type: field.TypeInt, Default: planet.DefaultPopulationProdLevel},
		{Name: "population_storage_level", Type: field.TypeInt, Default: planet.DefaultPopulationStorageLevel},
		{Name: "energy_cons", Type: field.TypeInt64, Default: planet.DefaultEnergyCons},
		{Name: "energy_prod", Type: field.TypeInt64, Default: planet.DefaultEnergyProd},
		{Name: "solar_prod_level", Type: field.TypeInt, Default: planet.DefaultSolarProdLevel},
		{Name: "name", Type: field.TypeString, Default: planet.DefaultName},
		{Name: "owner_id", Type: field.TypeInt, Nullable: true},
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
		{Name: "action", Type: field.TypeEnum, Enums: []string{"upgrade_metal_mine"}},
		{Name: "group", Type: field.TypeEnum, Enums: []string{"building"}},
		{Name: "end_time", Type: field.TypeTime},
		{Name: "planet_id", Type: field.TypeInt, Nullable: true},
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
				Name:    "timer_group_planet_id",
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
