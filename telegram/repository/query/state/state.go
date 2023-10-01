package stateQuery

const CreateDefaultStates = `
	INSERT INTO tbl_user_states (user_id, income_state, expense_state, income_category_state, expense_category_state)
	VALUES ($1, 0, 0, 0, 0)`
