// Code generated by ent, DO NOT EDIT.

package workflow

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the workflow type in the database.
	Label = "workflow"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldProjectOld holds the string denoting the project_old field in the database.
	FieldProjectOld = "project_old"
	// FieldTeam holds the string denoting the team field in the database.
	FieldTeam = "team"
	// FieldRunsCount holds the string denoting the runs_count field in the database.
	FieldRunsCount = "runs_count"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldDeletedAt holds the string denoting the deleted_at field in the database.
	FieldDeletedAt = "deleted_at"
	// FieldPublic holds the string denoting the public field in the database.
	FieldPublic = "public"
	// FieldOrganizationID holds the string denoting the organization_id field in the database.
	FieldOrganizationID = "organization_id"
	// FieldProjectID holds the string denoting the project_id field in the database.
	FieldProjectID = "project_id"
	// FieldDescription holds the string denoting the description field in the database.
	FieldDescription = "description"
	// EdgeRobotaccounts holds the string denoting the robotaccounts edge name in mutations.
	EdgeRobotaccounts = "robotaccounts"
	// EdgeWorkflowruns holds the string denoting the workflowruns edge name in mutations.
	EdgeWorkflowruns = "workflowruns"
	// EdgeOrganization holds the string denoting the organization edge name in mutations.
	EdgeOrganization = "organization"
	// EdgeContract holds the string denoting the contract edge name in mutations.
	EdgeContract = "contract"
	// EdgeIntegrationAttachments holds the string denoting the integration_attachments edge name in mutations.
	EdgeIntegrationAttachments = "integration_attachments"
	// EdgeProject holds the string denoting the project edge name in mutations.
	EdgeProject = "project"
	// EdgeReferrers holds the string denoting the referrers edge name in mutations.
	EdgeReferrers = "referrers"
	// Table holds the table name of the workflow in the database.
	Table = "workflows"
	// RobotaccountsTable is the table that holds the robotaccounts relation/edge.
	RobotaccountsTable = "robot_accounts"
	// RobotaccountsInverseTable is the table name for the RobotAccount entity.
	// It exists in this package in order to avoid circular dependency with the "robotaccount" package.
	RobotaccountsInverseTable = "robot_accounts"
	// RobotaccountsColumn is the table column denoting the robotaccounts relation/edge.
	RobotaccountsColumn = "workflow_robotaccounts"
	// WorkflowrunsTable is the table that holds the workflowruns relation/edge.
	WorkflowrunsTable = "workflow_runs"
	// WorkflowrunsInverseTable is the table name for the WorkflowRun entity.
	// It exists in this package in order to avoid circular dependency with the "workflowrun" package.
	WorkflowrunsInverseTable = "workflow_runs"
	// WorkflowrunsColumn is the table column denoting the workflowruns relation/edge.
	WorkflowrunsColumn = "workflow_workflowruns"
	// OrganizationTable is the table that holds the organization relation/edge.
	OrganizationTable = "workflows"
	// OrganizationInverseTable is the table name for the Organization entity.
	// It exists in this package in order to avoid circular dependency with the "organization" package.
	OrganizationInverseTable = "organizations"
	// OrganizationColumn is the table column denoting the organization relation/edge.
	OrganizationColumn = "organization_id"
	// ContractTable is the table that holds the contract relation/edge.
	ContractTable = "workflows"
	// ContractInverseTable is the table name for the WorkflowContract entity.
	// It exists in this package in order to avoid circular dependency with the "workflowcontract" package.
	ContractInverseTable = "workflow_contracts"
	// ContractColumn is the table column denoting the contract relation/edge.
	ContractColumn = "workflow_contract"
	// IntegrationAttachmentsTable is the table that holds the integration_attachments relation/edge.
	IntegrationAttachmentsTable = "integration_attachments"
	// IntegrationAttachmentsInverseTable is the table name for the IntegrationAttachment entity.
	// It exists in this package in order to avoid circular dependency with the "integrationattachment" package.
	IntegrationAttachmentsInverseTable = "integration_attachments"
	// IntegrationAttachmentsColumn is the table column denoting the integration_attachments relation/edge.
	IntegrationAttachmentsColumn = "workflow_id"
	// ProjectTable is the table that holds the project relation/edge.
	ProjectTable = "workflows"
	// ProjectInverseTable is the table name for the Project entity.
	// It exists in this package in order to avoid circular dependency with the "project" package.
	ProjectInverseTable = "projects"
	// ProjectColumn is the table column denoting the project relation/edge.
	ProjectColumn = "project_id"
	// ReferrersTable is the table that holds the referrers relation/edge. The primary key declared below.
	ReferrersTable = "referrer_workflows"
	// ReferrersInverseTable is the table name for the Referrer entity.
	// It exists in this package in order to avoid circular dependency with the "referrer" package.
	ReferrersInverseTable = "referrers"
)

// Columns holds all SQL columns for workflow fields.
var Columns = []string{
	FieldID,
	FieldName,
	FieldProjectOld,
	FieldTeam,
	FieldRunsCount,
	FieldCreatedAt,
	FieldDeletedAt,
	FieldPublic,
	FieldOrganizationID,
	FieldProjectID,
	FieldDescription,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "workflows"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"workflow_contract",
}

var (
	// ReferrersPrimaryKey and ReferrersColumn2 are the table columns denoting the
	// primary key for the referrers relation (M2M).
	ReferrersPrimaryKey = []string{"referrer_id", "workflow_id"}
)

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultRunsCount holds the default value on creation for the "runs_count" field.
	DefaultRunsCount int
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultPublic holds the default value on creation for the "public" field.
	DefaultPublic bool
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() uuid.UUID
)

// OrderOption defines the ordering options for the Workflow queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByName orders the results by the name field.
func ByName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldName, opts...).ToFunc()
}

// ByProjectOld orders the results by the project_old field.
func ByProjectOld(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldProjectOld, opts...).ToFunc()
}

// ByTeam orders the results by the team field.
func ByTeam(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldTeam, opts...).ToFunc()
}

// ByRunsCount orders the results by the runs_count field.
func ByRunsCount(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldRunsCount, opts...).ToFunc()
}

// ByCreatedAt orders the results by the created_at field.
func ByCreatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedAt, opts...).ToFunc()
}

// ByDeletedAt orders the results by the deleted_at field.
func ByDeletedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDeletedAt, opts...).ToFunc()
}

// ByPublic orders the results by the public field.
func ByPublic(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPublic, opts...).ToFunc()
}

// ByOrganizationID orders the results by the organization_id field.
func ByOrganizationID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldOrganizationID, opts...).ToFunc()
}

// ByProjectID orders the results by the project_id field.
func ByProjectID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldProjectID, opts...).ToFunc()
}

// ByDescription orders the results by the description field.
func ByDescription(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDescription, opts...).ToFunc()
}

// ByRobotaccountsCount orders the results by robotaccounts count.
func ByRobotaccountsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newRobotaccountsStep(), opts...)
	}
}

// ByRobotaccounts orders the results by robotaccounts terms.
func ByRobotaccounts(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newRobotaccountsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByWorkflowrunsCount orders the results by workflowruns count.
func ByWorkflowrunsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newWorkflowrunsStep(), opts...)
	}
}

// ByWorkflowruns orders the results by workflowruns terms.
func ByWorkflowruns(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newWorkflowrunsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByOrganizationField orders the results by organization field.
func ByOrganizationField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newOrganizationStep(), sql.OrderByField(field, opts...))
	}
}

// ByContractField orders the results by contract field.
func ByContractField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newContractStep(), sql.OrderByField(field, opts...))
	}
}

// ByIntegrationAttachmentsCount orders the results by integration_attachments count.
func ByIntegrationAttachmentsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newIntegrationAttachmentsStep(), opts...)
	}
}

// ByIntegrationAttachments orders the results by integration_attachments terms.
func ByIntegrationAttachments(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newIntegrationAttachmentsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByProjectField orders the results by project field.
func ByProjectField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newProjectStep(), sql.OrderByField(field, opts...))
	}
}

// ByReferrersCount orders the results by referrers count.
func ByReferrersCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newReferrersStep(), opts...)
	}
}

// ByReferrers orders the results by referrers terms.
func ByReferrers(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newReferrersStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newRobotaccountsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(RobotaccountsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, RobotaccountsTable, RobotaccountsColumn),
	)
}
func newWorkflowrunsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(WorkflowrunsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, WorkflowrunsTable, WorkflowrunsColumn),
	)
}
func newOrganizationStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(OrganizationInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, OrganizationTable, OrganizationColumn),
	)
}
func newContractStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(ContractInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, false, ContractTable, ContractColumn),
	)
}
func newIntegrationAttachmentsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(IntegrationAttachmentsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, true, IntegrationAttachmentsTable, IntegrationAttachmentsColumn),
	)
}
func newProjectStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(ProjectInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, ProjectTable, ProjectColumn),
	)
}
func newReferrersStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(ReferrersInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2M, true, ReferrersTable, ReferrersPrimaryKey...),
	)
}
