package analysis

// Tier 0 Analysis test cases - should never fail.
// List of applications with expected analysis outputs.
var Tier0TestCases = []TC{
	TackleTestappPublicWithDeps,
	TackleTestappPublicPackageFilter,
	Tomcat,
	AcmeairWebappBinary,	// Binary upload from binary/acmeair-webapp-1.0-SNAPSHOT.war
	TackleTestappPublicBinary,	// Binary fetched from tackle-testapp-public maven registry
}

// Tier 1 Analysis test cases - should work.
// List of applications with expected analysis outputs.
var Tier1TestCases = []TC{
	// Setting empty until have working applications.
}

// Tier 2 Analysis test cases - great if works.
// List of applications with expected analysis outputs.
var Tier2TestCases = []TC{
	AdministracionEfectivo,
	Daytrader,
	PetclinicHazelcast,
	ApacheWicket,
	SeamBooking,
}

// Tier 3 Analysis with credentials test cases - should work
// List of applications with expected analysis outputs.
var Tier3TestCases = []TC{
}
