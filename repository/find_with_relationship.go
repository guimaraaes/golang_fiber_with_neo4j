package repository

import (
	"fmt"

	"github.com/guimaraaes/golang_fiber_with_neo4j/repository/utils"
)

func FindWithRelR(modelSource interface{}, relation interface{}, modelTarget interface{}) ([]string, string) {
	nSource, properSource := utils.Label_Properties_Node(modelSource, utils.Model_Properties_To_Params(modelSource))
	nTarget, properTarget := utils.Label_Properties_Node(modelTarget, utils.Model_Properties_To_Params(modelTarget))
	rel, properRel := utils.Label_Properties_Node(relation, utils.Model_Properties_To_Params(relation))

	query := "MATCH (t {" + properTarget + "})-[rel {" + properRel + "}]-(s {" + properSource + "}) " +
		"WHERE $node1 in labels(s) " +
		"and $node2 in labels(t) " +
		"WITH s as N " +
		"CALL apoc.path.subgraphAll(N, {maxLevel:2}) YIELD nodes WITH [N in nodes | N {.*, label:labels(N)[0]}] as nodes " +
		"RETURN apoc.convert.toJson(nodes[0]) "
	m := map[string]interface{}{"node1": nSource, "node2": nTarget, "rel": rel}
	fmt.Println(query)
	fmt.Println(m)

	c, err := utils.QueryCall(query, m)

	return c, err
}
