package migration

func upgradeDatabaseToVersion3_2_0(fromVersion string) error {
	return updateSchemaVersion(version3_2_0)
}

func upgradeDatabaseToVersion3_2_1(fromVersion string) error {
	return updateSchemaVersion(version3_2_1)
}
