package repository

import (
	"github.com/guimaraaes/golang_fiber_with_neo4j/repository/utils"
)

func SaveWithRel(modelSource interface{}, relation interface{}, modelTarget interface{}, modelSourceUpdate interface{}, relationUpdate interface{}, modelTargetUpdate interface{}) ([]string, string) {

	nSource, properSource := utils.Label_and_Properties(modelSource)
	nTarget, properTarget := utils.Label_and_Properties(modelTarget)
	rel, properRel := utils.Label_and_Properties(relation)

	_, properSourceUpdate := utils.Label_and_Properties(modelSourceUpdate)
	_, properTargetUpdate := utils.Label_and_Properties(modelTargetUpdate)
	_, properRelUpdate := utils.Label_and_Properties(relationUpdate)

	query := "MERGE (t {" + properTarget + "})-[rel {" + properRel + "}]-(s {" + properSource + "}) " +
		"WHERE $nSource IN labels(s) AND $nTarget IN labels(t)" +
		"SET t += { " + properTargetUpdate + " }, rel += { " + properRelUpdate + " }, s += { " + properSourceUpdate + " } " +
		"WITH s AS N " +
		"CALL apoc.path.subgraphAll(N, {maxLevel:2}) YIELD nodes WITH [N in nodes | N {.*, label:labels(N)[0]}] as nodes " +
		"RETURN apoc.convert.toJson(nodes[0]) "
	m := map[string]interface{}{"nSource": nSource, "nTarget": nTarget, "rel": rel}

	c, err := utils.QueryCall(query, m)

	return c, err
}
