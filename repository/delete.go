package repository

import "github.com/guimaraaes/golang_fiber_with_neo4j/repository/utils"

func DeleteR(model interface{}, info map[string]interface{}) (string, string) {
	c, _ := FindR(model, info)
	if c == nil {
		return "", "não existe"
	}
	node, propertiesParams := utils.Label_Properties_Node(model, info)
	query := "MATCH (m {" + propertiesParams + "  }) WHERE $node in labels(m) " +
		"DETACH DELETE m "
	m := map[string]interface{}{"node": node}
	_, err := utils.QueryCall(query, m)
	return "excluído", err
}
