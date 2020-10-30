package repository

import "github.com/guimaraaes/golang_fiber_with_neo4j/repository/utils"

func CreateR(model interface{}) ([]string, string) {
	properties := utils.Model_Properties_To_Params(model)
	c, _ := FindR(model, properties)
	if c != nil {
		return nil, "já existe"
	}

	nodeCreate, propertiesCreate := utils.Label_Properties_Node(model, utils.Model_Properties_To_Params(model))
	query := "MERGE (m {" + propertiesCreate + "}) WITH m " +
		"CALL apoc.create.addLabels(m, [$node]) YIELD node " +
		"CALL apoc.path.subgraphAll(node, {maxLevel:0}) " +
		"YIELD nodes WITH [node in nodes | node {.*, label:labels(node)[0]}] as nodes " +
		"RETURN apoc.convert.toJson(nodes[0])"
	m := map[string]interface{}{"node": nodeCreate}
	c, err := utils.QueryCall(query, m)
	return c, err
}
