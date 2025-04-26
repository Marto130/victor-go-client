package routes

const (
	CreateIndexPath  = "/api/index/{indexName}"
	DestroyIndexPath = "/api/index/{indexName}"
	InsertVectorPath = "/api/vector/{indexName}"
	DeleteVectorPath = "/api/vector/{indexName}/{vectorID}"
	SearchVectorPath = "/api/vector/{indexName}/search"
)

const (
	CreateIndex  = "/api/index/%s"
	DestroyIndex = "/api/index/%s"
	InsertVector = "/api/vector/%s"
	DeleteVector = "/api/vector/%s/%v"
	SearchVector = "/api/vector/%s/search"
)
