//
// Code generated by go-jet DO NOT EDIT.
//
// WARNING: Changes to this file may cause incorrect behavior
// and will be lost if the code is regenerated
//

package table

// UseSchema sets a new schema name for all generated table SQL builder types. It is recommended to invoke
// this method only once at the beginning of the program.
func UseSchema(schema string) {
	AuditLogEntries = AuditLogEntries.FromSchema(schema)
	FlowState = FlowState.FromSchema(schema)
	Identities = Identities.FromSchema(schema)
	Instances = Instances.FromSchema(schema)
	MfaAmrClaims = MfaAmrClaims.FromSchema(schema)
	MfaChallenges = MfaChallenges.FromSchema(schema)
	MfaFactors = MfaFactors.FromSchema(schema)
	RefreshTokens = RefreshTokens.FromSchema(schema)
	SamlProviders = SamlProviders.FromSchema(schema)
	SamlRelayStates = SamlRelayStates.FromSchema(schema)
	SchemaMigrations = SchemaMigrations.FromSchema(schema)
	Sessions = Sessions.FromSchema(schema)
	SsoDomains = SsoDomains.FromSchema(schema)
	SsoProviders = SsoProviders.FromSchema(schema)
	Users = Users.FromSchema(schema)
}