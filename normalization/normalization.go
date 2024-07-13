package normalization

import (
	constants "goimporttodb/utils"
	"goimporttodb/validators"
	"strings"
)

// The normalization function validates all data read in each line of the file,
// as well as removes white spaces and line breaks. If the validation of all fields is true,
// the function returns a populated map to save the data in the database.
// If the result of any validation is false, an empty map will be returned and the data will not be persisted in the database.

func NormalizationLine(colums_table map[int]string) map[int]string {
	map_colums_clients_table_normalized := make(map[int]string)
	if validators.IsCpfOrCnpj(colums_table[constants.COLUMN_CPF]) &&
		validators.IsZeroOrOne(colums_table[constants.COLUMN_PRIVATE]) &&
		validators.IsZeroOrOne(colums_table[constants.COLUMN_INCOMPLETO]) &&
		validators.IsDate(colums_table[constants.COLUMN_DATA_ULTIMA_COMPRA]) &&
		validators.IsMonetaryReal(colums_table[constants.COLUMN_TICKET_MEDIO]) &&
		validators.IsMonetaryReal(colums_table[constants.COLUMN_TICKET_ULTIMA_COMPRA]) &&
		validators.IsCpfOrCnpj(colums_table[constants.COLUMN_LOJA_MAIS_FREQUENTE]) &&
		validators.IsCpfOrCnpj(colums_table[constants.COLUMN_LOJA_ULTIMA_COMPRA]) {
		map_colums_clients_table_normalized[constants.COLUMN_CPF] = strings.TrimSuffix(colums_table[constants.COLUMN_CPF], " ")
		map_colums_clients_table_normalized[constants.COLUMN_PRIVATE] = colums_table[constants.COLUMN_PRIVATE]
		map_colums_clients_table_normalized[constants.COLUMN_INCOMPLETO] = colums_table[constants.COLUMN_INCOMPLETO]
		map_colums_clients_table_normalized[constants.COLUMN_DATA_ULTIMA_COMPRA] = colums_table[constants.COLUMN_DATA_ULTIMA_COMPRA]
		map_colums_clients_table_normalized[constants.COLUMN_TICKET_MEDIO] = strings.Replace(colums_table[constants.COLUMN_TICKET_MEDIO], ",", ".", -1)
		map_colums_clients_table_normalized[constants.COLUMN_TICKET_ULTIMA_COMPRA] = strings.Replace(colums_table[constants.COLUMN_TICKET_ULTIMA_COMPRA], ",", ".", -1)
		map_colums_clients_table_normalized[constants.COLUMN_LOJA_MAIS_FREQUENTE] = strings.TrimSuffix(colums_table[constants.COLUMN_LOJA_MAIS_FREQUENTE], " ")
		colums_table[constants.COLUMN_LOJA_ULTIMA_COMPRA] = strings.TrimSuffix(colums_table[constants.COLUMN_LOJA_ULTIMA_COMPRA], " ")
		map_colums_clients_table_normalized[constants.COLUMN_LOJA_ULTIMA_COMPRA] = strings.TrimSuffix(colums_table[constants.COLUMN_LOJA_ULTIMA_COMPRA], "\n")
		return map_colums_clients_table_normalized
	}
	return map_colums_clients_table_normalized
}
