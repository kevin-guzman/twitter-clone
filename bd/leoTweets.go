package bd

import (
	"context"
	"log"
	"time"
	"twitter/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func LeoTweets(ID string, pagina int64, limite int64) ([]*models.DevuelvoTweets, bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	db := MongoCN.Database("twittor")
	col := db.Collection("tweet")

	var resultados []*models.DevuelvoTweets

	condicion:=bson.M{"userid":ID}
	opciones:=options.Find()
	opciones.SetLimit(limite)
	opciones.SetSort(bson.D{{Key:"fecha", Value: -1}})
	opciones.SetSkip((pagina-1)*limite)

	cursor,err:=col.Find(ctx,condicion,opciones)
	if err != nil {
		log.Fatal(err.Error())
		return resultados, false,err
	}

	for cursor.Next(context.TODO()){
		var registro models.DevuelvoTweets
		err:=cursor.Decode(&registro)
		if err != nil {
			return resultados, false, nil
		}
		resultados=append(resultados, &registro)
	}
	
	return resultados, true, nil
}
