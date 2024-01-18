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
	column "date" {
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

table "assistances" {
	schema = schema.main

	column "id" {
		null = false
		type = integer
		auto_increment = true
	}
	column "type" {
		null = false
		type = text
	}
	column "date" {
		type = datetime
		null = false
	}

	primary_key {
		columns = [column.id]
	}
}

table "user_assistances" {
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
	column "assistance_id" {
		type = integer
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
	foreign_key "assistance_id" {
		columns     = [column.assistance_id]
		ref_columns = [table.assistances.column.id]
		on_update   = NO_ACTION
		on_delete   = CASCADE
	}
}

schema "main" {
}
