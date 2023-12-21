table "users" {
	schema = schema.main
	column "id" {
		null = false
		type = integer
		auto_increment = true
	}
	column "first_name" {
		null = false
		type = text
	}
	column "last_name" {
		null = false
		type = text
	}
	primary_key {
		columns = [column.id]
	}
}
schema "main" {
}
