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

table "weights" {
	schema = schema.main

	column "id" {
		null = false
		type = integer
		auto_increment = true
	}
	column "user_id" {
		type = integer
		null = false
	}
	column "weight" {
		type = decimal
		null = false
	}
	column "creation_date" {
		type = datetime
		null = false
	}

	primary_key {
		columns = [column.id]
	}

	foreign_key "user_id" {
		columns     = [column.user_id]
		ref_columns = [table.users.column.id]
		on_update   = NO_ACTION
		on_delete   = CASCADE
	}
}

schema "main" {
}
