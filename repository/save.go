package repository

import "github.com/guimaraaes/golang_fiber_with_neo4j/repository/utils"

func SaveR(model interface{}, info map[string]interface{}) ([]string, string) {
	c, _ := FindR(model, info)
	if c == nil {
		return nil, "não existe"
	}
	node, propertiesSet := utils.Label_and_Properties(model)
	propertiesParams := utils.Properties(info)

	query := "MATCH (m {" + propertiesParams + "  }) WHERE $node in labels(m) " +
		"SET m += { " + propertiesSet + " } WITH m " +
		"CALL apoc.path.subgraphAll(m, {maxLevel:0}) " +
		"YIELD nodes WITH [node in nodes | node {.*, label:labels(node)[0]}] as nodes " +
		"RETURN apoc.convert.toJson(nodes[0])"
	m := map[string]interface{}{"node": node}
	c, err := utils.QueryCall(query, m)
	return c, err

}
