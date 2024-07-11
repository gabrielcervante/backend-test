package database

var (
	insert = "INSERT INTO breeds (id, species, pet_size, `name`, `weight`, average_male_adult_weight, average_female_adult_weight) VALUES (?, ?, ?, ?, ?, ?, ?)"

	SelectById = "SELECT id, species, pet_size, `name`, `weight`, average_male_adult_weight, average_female_adult_weight FROM breeds WHERE id = ?"

	SelectBySpecies = "SELECT id, species, pet_size, `name`, `weight`, average_male_adult_weight, average_female_adult_weight FROM breeds WHERE species = ?"

	SelectByPetSize = "SELECT id, species, pet_size, `name`, `weight`, average_male_adult_weight, average_female_adult_weight FROM breeds WHERE pet_size = ?"

	SelectByName = "SELECT id, species, pet_size, `name`, `weight`, average_male_adult_weight, average_female_adult_weight FROM breeds WHERE `name` = ?"

	SelectByWeight = "SELECT id, species, pet_size, `name`, `weight`, average_male_adult_weight, average_female_adult_weight FROM breeds WHERE `weight` = ?"

	update = "UPDATE breeds SET species=?, pet_size=?, `name`=?, `weight`=?, average_male_adult_weight=?, average_female_adult_weight=? WHERE id=?"

	delete = `DELETE FROM breeds WHERE id = ?`
)
