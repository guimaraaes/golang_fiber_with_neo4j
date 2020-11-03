package repository

import (
	"github.com/guimaraaes/golang_fiber_with_neo4j/repository/utils"
)

func CreateWithRElR(modelSource interface{}, relation interface{}, modelTarget interface{}) ([]string, string) {

	nSource, properSource := utils.Label_and_Properties(modelSource)
	nTarget, properTarget := utils.Label_and_Properties(modelTarget)
	rel, properRel := utils.Label_and_Properties(relation)

	query := "MERGE (t {" + properTarget + "}) WITH t " +
		"MERGE (s{" + properSource + "}) WITH s, t " +
		"CALL apoc.create.addLabels(t, [$node2])YIELD node " +
		"CALL apoc.create.addLabels(s, [$node1]) YIELD node AS N " +
		"WITH s, t, N " +
		"CALL apoc.create.relationship(s, $rel,{" + properRel + "}, t)YIELD rel " +
		"WITH s, t, N " +
		"CALL apoc.path.subgraphAll(N, {maxLevel:0}) YIELD nodes WITH [N in nodes | N {.*, label:labels(N)[0]}] as nodes " +
		"RETURN apoc.convert.toJson(nodes[0]) "
	m := map[string]interface{}{"node1": nSource, "node2": nTarget, "rel": rel}

	c, err := utils.QueryCall(query, m)

	return c, err
}
