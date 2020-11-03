package repository_algorithms

import (
	"github.com/guimaraaes/golang_fiber_with_neo4j/repository/utils"
	_ "github.com/guimaraaes/golang_fiber_with_neo4j/repository/utils"
)

func Centrality(node, relationship string) ([]string, string) {
	res, err := utils.QueryCall("", nil)
	return res, err
}

func Community(node, relationship string) ([]string, string) {
	res, err := utils.QueryCall("", nil)
	return res, err
}

func Path(node, relationship string) ([]string, string) {
	res, err := utils.QueryCall("", nil)
	return res, err
}

func PageRank(node, relationship string) ([]string, string) {
	query := "CALL gds.pageRank.stream({ nodeProjection: [$node], " +
		"relationshipProjection: [$rel], " +
		"maxIterations: 1, dampingFactor: 0.85}) " +
		"YIELD nodeId, score " +
		"RETURN apoc.convert.toJson(gds.util.asNode(nodeId).name) ORDER BY score DESC"
	m := map[string]interface{}{"node": node, "rel": relationship}
	res, err := utils.QueryCall(query, m)
	return res, err
}
