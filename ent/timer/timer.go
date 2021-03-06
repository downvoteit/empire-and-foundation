// Code generated by entc, DO NOT EDIT.

package timer

import (
	"fmt"
)

const (
	// Label holds the string label denoting the timer type in the database.
	Label = "timer"
	// FieldID holds the string denoting the id field in the database.
	FieldID      = "id"     // FieldAction holds the string denoting the action vertex property in the database.
	FieldAction  = "action" // FieldGroup holds the string denoting the group vertex property in the database.
	FieldGroup   = "group"  // FieldEndTime holds the string denoting the end_time vertex property in the database.
	FieldEndTime = "end_time"

	// EdgePlanet holds the string denoting the planet edge name in mutations.
	EdgePlanet = "planet"

	// Table holds the table name of the timer in the database.
	Table = "timers"
	// PlanetTable is the table the holds the planet relation/edge.
	PlanetTable = "timers"
	// PlanetInverseTable is the table name for the Planet entity.
	// It exists in this package in order to avoid circular dependency with the "planet" package.
	PlanetInverseTable = "planets"
	// PlanetColumn is the table column denoting the planet relation/edge.
	PlanetColumn = "planet_timers"
)

// Columns holds all SQL columns for timer fields.
var Columns = []string{
	FieldID,
	FieldAction,
	FieldGroup,
	FieldEndTime,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the Timer type.
var ForeignKeys = []string{
	"planet_timers",
}

// Action defines the type for the action enum field.
type Action string

// Action values.
const (
	ActionUpgradeMetalProd       Action = "upgrade_metal_prod"
	ActionUpgradeHydrogenProd    Action = "upgrade_hydrogen_prod"
	ActionUpgradeSilicaProd      Action = "upgrade_silica_prod"
	ActionUpgradeSolarProd       Action = "upgrade_solar_prod"
	ActionUpgradeUrbanism        Action = "upgrade_urbanism"
	ActionUpgradeMetalStorage    Action = "upgrade_metal_storage"
	ActionUpgradeHydrogenStorage Action = "upgrade_hydrogen_storage"
	ActionUpgradeSilicaStorage   Action = "upgrade_silica_storage"
	ActionUpgradeResearchCenter  Action = "upgrade_research_center"
	ActionUpgradeShipFactory     Action = "upgrade_ship_factory"
	ActionTest                   Action = "test"
)

func (s Action) String() string {
	return string(s)
}

// ActionValidator is a validator for the "a" field enum values. It is called by the builders before save.
func ActionValidator(a Action) error {
	switch a {
	case ActionUpgradeMetalProd, ActionUpgradeHydrogenProd, ActionUpgradeSilicaProd, ActionUpgradeSolarProd, ActionUpgradeUrbanism, ActionUpgradeMetalStorage, ActionUpgradeHydrogenStorage, ActionUpgradeSilicaStorage, ActionUpgradeResearchCenter, ActionUpgradeShipFactory, ActionTest:
		return nil
	default:
		return fmt.Errorf("timer: invalid enum value for action field: %q", a)
	}
}

// Group defines the type for the group enum field.
type Group string

// Group values.
const (
	GroupBuilding Group = "building"
)

func (s Group) String() string {
	return string(s)
}

// GroupValidator is a validator for the "gr" field enum values. It is called by the builders before save.
func GroupValidator(gr Group) error {
	switch gr {
	case GroupBuilding:
		return nil
	default:
		return fmt.Errorf("timer: invalid enum value for group field: %q", gr)
	}
}
