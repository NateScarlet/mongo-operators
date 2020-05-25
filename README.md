# mongo-operators

[![godev](https://img.shields.io/static/v1?label=godev&message=reference&color=00add8)](https://pkg.go.dev/github.com/NateScarlet/mongo-operators)
[![build status](https://github.com/NateScarlet/mongo-operators/workflows/go/badge.svg)](https://github.com/NateScarlet/mongo-operators/actions)

Utility functions to create [mongodb](https://www.mongodb.com/) operator.

## Why

- IDE hint.
- Avoid typo.
- Type check for literal options.
- Shortcut for common operator combinations.

## Usage

```shell
go get github.com/NateScarlet/mongo-operators
```

```Go
import (
    "context"
    "github.com/mongodb/mongo-go-driver/mongo"
    "github.com/mongodb/mongo-go-driver/mongo/bson/primitive"
    q "github.com/NateScarlet/mongo-operators/pkg/query"
    a "github.com/NateScarlet/mongo-operators/pkg/aggregation"
)

type M = primitive.M
type A = primitive.A

func find(ctx context.Context, col *mongo.Collection) {
    col.FindOne(ctx, M{"name": q.In(A{"foo", "bar"})})
}

func update(ctx context.Context, col *mongo.Collection) {
    col.UpdateOne(ctx, M{}, q.Update(
        q.Set(M{"name": "foo"}),
        q.Push(M{"tags": "bar"}),
    ))
}

func aggregation(ctx context.Context, col *mongo.Collection) {
    col.Aggregate(ctx, A{
        a.Match(M{"name": "foo"}),
        a.Unwind("tags"),
    })
}

```
